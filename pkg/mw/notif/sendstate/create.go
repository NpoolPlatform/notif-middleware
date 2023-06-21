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

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createSendState(ctx context.Context, tx *ent.Tx, req *sendstatecrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if req.UserID == nil {
		return fmt.Errorf("invalid userid")
	}
	if req.EventID == nil {
		return fmt.Errorf("invalid eventid")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*req.AppID,
		*req.UserID,
		*req.EventID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()
	h.Conds = &sendstatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		UserID:  &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
		EventID: &cruder.Cond{Op: cruder.EQ, Val: *req.EventID},
		Channel: &cruder.Cond{Op: cruder.EQ, Val: *req.Channel},
	}

	exist, err := h.ExistSendStateConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("sendstate exist")
	}

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}

	info, err := sendstatecrud.CreateSet(
		tx.SendNotif.Create(),
		&sendstatecrud.Req{
			ID:      req.ID,
			AppID:   req.AppID,
			UserID:  req.UserID,
			EventID: req.EventID,
			Channel: req.Channel,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateSendState(ctx context.Context) (*npool.SendState, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &sendstatecrud.Req{
		ID:      handler.ID,
		AppID:   handler.AppID,
		UserID:  handler.UserID,
		EventID: handler.EventID,
		Channel: handler.Channel,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createSendState(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSendState(ctx)
}

func (h *Handler) CreateSendStates(ctx context.Context) ([]*npool.SendState, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createSendState(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &sendstatecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetSendStates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
