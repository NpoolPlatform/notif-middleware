package sms

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/sms"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createSMSTemplate(ctx context.Context, tx *ent.Tx, req *smstemplatecrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid lang")
	}
	if req.LangID == nil {
		return fmt.Errorf("invalid logo")
	}
	if req.UsedFor == nil {
		return fmt.Errorf("invalid usedFor")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*req.AppID,
		*req.LangID,
		req.UsedFor,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &smstemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *req.UsedFor},
	}
	exist, err := h.ExistSMSTemplateConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("smstemplate exist")
	}

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}

	info, err := smstemplatecrud.CreateSet(
		tx.SMSTemplate.Create(),
		&smstemplatecrud.Req{
			ID:      req.ID,
			AppID:   req.AppID,
			LangID:  req.LangID,
			UsedFor: req.UsedFor,
			Subject: req.Subject,
			Message: req.Message,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateSMSTemplate(ctx context.Context) (*npool.SMSTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &smstemplatecrud.Req{
		ID:      handler.ID,
		AppID:   handler.AppID,
		LangID:  handler.LangID,
		UsedFor: handler.UsedFor,
		Subject: handler.Subject,
		Message: handler.Message,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createSMSTemplate(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSMSTemplate(ctx)
}

func (h *Handler) CreateSMSTemplates(ctx context.Context) ([]*npool.SMSTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createSMSTemplate(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &smstemplatecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetSMSTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
