package sendstate

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
	entsendamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendannouncement"
)

type queryHandler struct {
	*Handler
	stm   *ent.SendAnnouncementSelect
	infos []*npool.SendState
	total uint32
}

func (h *queryHandler) selectSendState(stm *ent.SendAnnouncementQuery) {
	h.stm = stm.Select(
		entsendamt.FieldID,
		entsendamt.FieldAppID,
		entsendamt.FieldUserID,
		entsendamt.FieldAnnouncementID,
		entsendamt.FieldChannel,
		entsendamt.FieldCreatedAt,
		entsendamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinAnnouncement(s *sql.Selector) {
	t := sql.Table(entamt.Table)
	s.LeftJoin(t).
		On(
			s.C(entsendamt.FieldAnnouncementID),
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

func (h *queryHandler) querySendState(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid user announcement id")
	}
	h.selectSendState(
		cli.SendAnnouncement.
			Query().
			Where(
				entsendamt.ID(*h.ID),
				entsendamt.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) querySendStatesByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.SendAnnouncement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectSendState(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSendStates(ctx context.Context) ([]*npool.SendState, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendStatesByConds(_ctx, cli); err != nil {
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

func (h *Handler) GetSendState(ctx context.Context) (info *npool.SendState, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendState(cli); err != nil {
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
