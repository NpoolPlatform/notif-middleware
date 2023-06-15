package channel

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/channel"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entchannel "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notifchannel"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifChannelSelect
	infos []*npool.Channel
	total uint32
}

func (h *queryHandler) selectChannel(stm *ent.NotifChannelQuery) {
	h.stm = stm.Select(
		entchannel.FieldID,
		entchannel.FieldAppID,
		entchannel.FieldChannel,
		entchannel.FieldEventType,
		entchannel.FieldCreatedAt,
		entchannel.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryChannel(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid notif/channel id")
	}
	h.selectChannel(
		cli.NotifChannel.
			Query().
			Where(
				entchannel.ID(*h.ID),
				entchannel.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryChannelsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.NotifChannel.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectChannel(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetChannels(ctx context.Context) ([]*npool.Channel, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannelsByConds(_ctx, cli); err != nil {
			return err
		}

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

func (h *Handler) GetChannel(ctx context.Context) (info *npool.Channel, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannel(cli); err != nil {
			return err
		}
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

func (h *Handler) GetChannelConds(ctx context.Context) ([]*npool.Channel, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChannelsByConds(_ctx, cli); err != nil {
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
