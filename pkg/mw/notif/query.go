package notif

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	entnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notif"
	entreadnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readnotif"
	entsendnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendnotif"
	entusernotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifSelect
	infos []*npool.Notif
	total uint32
}

func (h *queryHandler) selectNotif(stm *ent.NotifQuery) {
	h.stm = stm.Select(
		entnotif.FieldID,
		entnotif.FieldAppID,
		entnotif.FieldLangID,
		entnotif.FieldUserID,
		entnotif.FieldEventID,
		entnotif.FieldNotified,
		entnotif.FieldEventType,
		entnotif.FieldUseTemplate,
		entnotif.FieldTitle,
		entnotif.FieldContent,
		entnotif.FieldChannel,
		entnotif.FieldExtra,
		entnotif.FieldType,
		entnotif.FieldCreatedAt,
		entnotif.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryNotif(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid notifid")
	}

	h.selectNotif(
		cli.Notif.
			Query().
			Where(
				entnotif.ID(*h.ID),
				entnotif.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinNotifUser(s)
		h.queryJoinSendNotif(s)
		h.queryJoinReadNotif(s)
	})
}

func (h *queryHandler) queryJoinNotifUser(s *sql.Selector) {
	t := sql.Table(entusernotif.Table)
	s.LeftJoin(t).
		On(
			t.C(entusernotif.FieldEventType),
			s.C(entnotif.FieldEventType),
		).
		AppendSelect(
			sql.As(t.C(entusernotif.FieldUserID), "notif_user_id"),
		)
}

func (h *queryHandler) queryJoinSendNotif(s *sql.Selector) {
	t := sql.Table(entsendnotif.Table)
	s.LeftJoin(t).
		On(
			t.C(entsendnotif.FieldEventID),
			s.C(entnotif.FieldEventID),
		).
		On(
			t.C(entsendnotif.FieldChannel),
			s.C(entnotif.FieldChannel),
		).
		AppendSelect(
			sql.As(t.C(entsendnotif.FieldUserID), "send_user_id"),
		)
}

func (h *queryHandler) queryJoinReadNotif(s *sql.Selector) {
	t := sql.Table(entreadnotif.Table)
	s.LeftJoin(t).
		On(
			t.C(entreadnotif.FieldNotifID),
			s.C(entnotif.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entreadnotif.FieldUserID), "read_user_id"),
		)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
		info.NotifType = basetypes.NotifType(basetypes.NotifType_value[info.NotifTypeStr])
		switch info.NotifType {
		case basetypes.NotifType_NotifMulticast:
			info.UserID = info.NotifUserID
			switch info.NotifType {
			case basetypes.NotifType(basetypes.NotifChannel_ChannelEmail):
				fallthrough //nolint
			case basetypes.NotifType(basetypes.NotifChannel_ChannelSMS):
				if info.UserID == info.SendUserID {
					info.Notified = true
				}
			case basetypes.NotifType(basetypes.NotifChannel_ChannelFrontend):
				if info.UserID == info.ReadUserID {
					info.Notified = true
				}
			}
		case basetypes.NotifType_NotifUnicast:
		}
	}
}

func (h *queryHandler) queryNotifs(ctx context.Context, cli *ent.Client) error {
	stm, err := notifcrud.SetQueryConds(cli.Notif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectNotif(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetNotif(ctx context.Context) (*npool.Notif, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotif(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetNotifs(ctx context.Context) ([]*npool.Notif, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(ctx, cli); err != nil {
			return err
		}
		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetNotifOnly(ctx context.Context) (info *npool.Notif, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(_ctx, cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	handler.formalize()

	return handler.infos[0], nil
}
