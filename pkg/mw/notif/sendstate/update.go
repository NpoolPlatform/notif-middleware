package sendstate

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	sendstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/sendstate"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateSendState(ctx context.Context, cli *ent.Client) error {
	if _, err := sendstatecrud.UpdateSet(
		cli.SendNotif.UpdateOneID(*h.ID),
		&sendstatecrud.Req{
			NotifID: h.NotifID,
			Channel: h.Channel,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSendState(ctx context.Context) (*npool.SendState, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}

	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixSetFiat,
		*h.ID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &sendstatecrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	h.Offset = 0
	h.Limit = 2

	email, err := h.GetSendStateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if email != nil {
		return nil, fmt.Errorf("sendstate exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateSendState(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSendState(ctx)
}
