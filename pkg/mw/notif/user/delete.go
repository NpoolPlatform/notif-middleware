package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"
)

func (h *Handler) DeleteUser(ctx context.Context) (*npool.User, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &usercrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistUserConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("user not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := usercrud.UpdateSet(
			cli.UserNotif.UpdateOneID(*h.ID),
			&usercrud.Req{
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
