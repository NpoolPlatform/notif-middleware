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

func (h *createHandler) createUser(ctx context.Context, tx *ent.Tx, req *usercrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.UserID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.NotifID == nil {
		return fmt.Errorf("invalid eventid")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*req.AppID,
		*req.UserID,
		*req.NotifID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &usercrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:  &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		NotifID: &cruder.Cond{Op: cruder.EQ, Val: *req.NotifID},
	}
	exist, err := h.ExistUserConds(ctx)
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
			ID:      req.ID,
			AppID:   req.AppID,
			UserID:  req.UserID,
			NotifID: req.NotifID,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateUser(ctx context.Context) (*npool.UserNotif, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &usercrud.Req{
		ID:      handler.ID,
		AppID:   handler.AppID,
		UserID:  handler.UserID,
		NotifID: handler.NotifID,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createUser(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}

func (h *Handler) CreateUsers(ctx context.Context) ([]*npool.UserNotif, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createUser(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &usercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
