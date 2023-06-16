package channel

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/channel"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) CreateChannel(ctx context.Context) (info *npool.Channel, err error) {
	h.Conds = &crud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		Channel:   &cruder.Cond{Op: cruder.EQ, Val: basetypes.NotifChannel(uint32(*h.Channel))},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: basetypes.UsedFor(uint32(*h.EventType))},
	}

	exist, err := h.ExistChannelConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("channel exist")
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
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			h.AppID = req.AppID
			h.Channel = req.Channel
			h.EventType = req.EventType

			h.Conds = &crud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				Channel:   &cruder.Cond{Op: cruder.EQ, Val: basetypes.NotifChannel(uint32(*h.Channel))},
				EventType: &cruder.Cond{Op: cruder.EQ, Val: basetypes.UsedFor(uint32(*h.EventType))},
			}

			exist, err := h.ExistChannelConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				continue
			}

			info, err := h.CreateChannel(ctx)
			if err != nil {
				return err
			}
			infos = append(infos, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
