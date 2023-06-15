package channel

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/channel"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
	Reqs []*crud.Req
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateChannel(ctx context.Context) (info *npool.Channel, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.NotifChannel.Create(),
			&crud.Req{
				ID:        h.ID,
				AppID:     h.AppID,
				Channel:   h.Channel,
				EventType: h.EventType,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetChannel(ctx)
}

func (h *Handler) CreateChannels(ctx context.Context) (infos []*npool.Channel, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		rows := []*ent.NotifChannelCreate{}
		for _, req := range handler.Reqs {
			rows = append(rows, crud.CreateSet(cli.NotifChannel.Create(), req))
		}

		_infos, err := cli.NotifChannel.CreateBulk(rows...).Save(_ctx)
		if err != nil {
			return err
		}

		for _, row := range _infos {
			infos = append(infos, &npool.Channel{
				ID:           row.ID.String(),
				AppID:        row.AppID.String(),
				EventType:    basetypes.UsedFor(basetypes.UsedFor_value[row.EventType]),
				EventTypeStr: row.EventType,
				Channel:      basetypes.NotifChannel(basetypes.NotifChannel_value[row.Channel]),
				ChannelStr:   row.EventType,
				CreatedAt:    row.CreatedAt,
				UpdatedAt:    row.UpdatedAt,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
