package user

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createNotifUser(ctx context.Context, tx *ent.Tx, req *usercrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.UserID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.EventType == nil {
		return fmt.Errorf("invalid eventtype")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*req.AppID,
		*req.UserID,
		*req.EventType,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &usercrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:    &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *req.EventType},
	}
	exist, err := h.ExistNotifUserConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user notif exist")
	}

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}

	info, err := usercrud.CreateSet(
		tx.UserNotif.Create(),
		&usercrud.Req{
			ID:        req.ID,
			AppID:     req.AppID,
			UserID:    req.UserID,
			EventType: req.EventType,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &usercrud.Req{
		ID:        handler.ID,
		AppID:     handler.AppID,
		UserID:    handler.UserID,
		EventType: handler.EventType,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createNotifUser(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotifUser(ctx)
}
