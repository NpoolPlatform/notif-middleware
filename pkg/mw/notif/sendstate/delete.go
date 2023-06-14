package sendstate

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	sendstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/sendstate"
)

func (h *Handler) DeleteSendState(ctx context.Context) (*npool.SendState, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetSendState(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &sendstatecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistSendStateConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("sendstate not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := sendstatecrud.UpdateSet(
			cli.SendNotif.UpdateOneID(*h.ID),
			&sendstatecrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
