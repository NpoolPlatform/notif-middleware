package announcement

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
	entread "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readannouncement"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm    *ent.AnnouncementSelect
	infos  []*npool.Announcement
	total  uint32
	UserID *uuid.UUID
}

func (h *queryHandler) selectAnnouncement(stm *ent.AnnouncementQuery) {
	h.stm = stm.Select(
		entamt.FieldID,
		entamt.FieldAppID,
		entamt.FieldLangID,
		entamt.FieldTitle,
		entamt.FieldContent,
		entamt.FieldChannel,
		entamt.FieldType,
		entamt.FieldStartAt,
		entamt.FieldEndAt,
		entamt.FieldCreatedAt,
		entamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinReadState(s *sql.Selector) {
	t := sql.Table(entread.Table)
	s.LeftJoin(t).
		On(
			s.C(entread.FieldAppID),
			t.C(entamt.FieldAppID),
		).
		OnP(
			sql.EQ(t.C(entread.FieldUserID), *h.UserID),
		).
		OnP(
			sql.EQ(s.C(entamt.FieldLangID), *h.LangID),
		).
		AppendSelect(
			sql.As(t.C(entread.FieldUserID), "user_id"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		if h.UserID != nil && h.LangID != nil {
			h.queryJoinReadState(s)
		}
	})
}

func (h *queryHandler) queryAnnouncement(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid announcement id")
	}
	h.selectAnnouncement(
		cli.Announcement.
			Query().
			Where(
				entamt.ID(*h.ID),
				entamt.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.AnnouncementType = basetypes.NotifType(basetypes.NotifType_value[info.AnnouncementTypeStr])
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
	}
}

func (h *queryHandler) queryAnnouncementsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.Announcement.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAnnouncement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetAnnouncements(ctx context.Context) ([]*npool.Announcement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAnnouncementsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.queryJoin()
		handler.
			stm.
			Offset(int(h.Offset)).
			Order(ent.Desc(entamt.FieldUpdatedAt)).
			Limit(int(h.Limit))

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAnnouncement(cli); err != nil {
			return err
		}

		if err := handler.scan(_ctx); err != nil {
			return err
		}

		handler.queryJoin()
		handler.formalize()
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
