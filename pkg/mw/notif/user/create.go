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

func (h *createHandler) createUser(ctx context.Context, cli *ent.Client) error {
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*h.AppID,
		*h.UserID,
		*h.NotifID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &usercrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:  &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		NotifID: &cruder.Cond{Op: cruder.EQ, Val: *h.NotifID},
	}
	exist, err := h.ExistUserConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := usercrud.CreateSet(
		cli.UserNotif.Create(),
		&usercrud.Req{
			ID:      h.ID,
			AppID:   h.AppID,
			UserID:  h.UserID,
			NotifID: h.NotifID,
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
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createUser(ctx, cli); err != nil {
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

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.AppID = req.AppID
			handler.UserID = req.UserID
			if err := handler.createUser(ctx, cli); err != nil {
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
