package announcement

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
)

type queryHandler struct {
	*Handler
	stm   *ent.AnnouncementSelect
	infos []*npool.Announcement
	total uint32
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
		entamt.FieldEndAt,
		entamt.FieldCreatedAt,
		entamt.FieldUpdatedAt,
	)
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

func (h *Handler) GetAnnouncementConds(ctx context.Context) ([]*npool.Announcement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAnnouncementsByConds(_ctx, cli); err != nil {
			return err
		}

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
