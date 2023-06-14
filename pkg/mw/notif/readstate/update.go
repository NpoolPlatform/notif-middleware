package readstate

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/readstate"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateReadState(ctx context.Context, cli *ent.Client) error {
	if _, err := readstatecrud.UpdateSet(
		cli.ReadNotif.UpdateOneID(*h.ID),
		&readstatecrud.Req{},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateReadState(ctx context.Context) (*npool.ReadState, error) {
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

	h.Conds = &readstatecrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	h.Offset = 0
	h.Limit = 2

	email, err := h.GetReadStateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if email != nil {
		return nil, fmt.Errorf("readstate exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateReadState(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetReadState(ctx)
}
