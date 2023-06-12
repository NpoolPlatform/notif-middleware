package read

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/read"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
	entreadamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readannouncement"
)

type queryHandler struct {
	*Handler
	stm   *ent.ReadAnnouncementSelect
	infos []*npool.ReadAnnouncement
	total uint32
}

func (h *queryHandler) selectReadAnnouncement(stm *ent.ReadAnnouncementQuery) {
	h.stm = stm.Select(
		entreadamt.FieldID,
		entreadamt.FieldAppID,
		entreadamt.FieldUserID,
		entreadamt.FieldAnnouncementID,
		entreadamt.FieldCreatedAt,
		entreadamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinAnnouncement(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.LeftJoin(t).
		On(
			s.C(entreadamt.FieldAnnouncementID),
			t.C(entamt.FieldID),
		).
		AppendSelect(
			t.C(entamt.FieldLangID),
			t.C(entamt.FieldTitle),
			t.C(entamt.FieldContent),
			t.C(entamt.FieldType),
			t.C(entamt.FieldChannel),
			t.C(entamt.FieldEndAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAnnouncement(s)
	})
}

func (h *queryHandler) queryReadAnnouncement(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid user announcement id")
	}
	h.selectReadAnnouncement(
		cli.ReadAnnouncement.
			Query().
			Where(
				entreadamt.ID(*h.ID),
				entreadamt.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryReadAnnouncementsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.ReadAnnouncement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectReadAnnouncement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetReadAnnouncements(ctx context.Context) ([]*npool.ReadAnnouncement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadAnnouncementsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.queryJoin()
		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetReadAnnouncement(ctx context.Context) (info *npool.ReadAnnouncement, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadAnnouncement(cli); err != nil {
			return err
		}

		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
