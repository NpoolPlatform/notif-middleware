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

func (h *createHandler) createSMSTemplate(ctx context.Context, cli *ent.Client) error {
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*h.AppID,
		*h.LangID,
		*h.UsedFor,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &smstemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistSMSTemplateConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("smstemplate exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := smstemplatecrud.CreateSet(
		cli.SMSTemplate.Create(),
		&smstemplatecrud.Req{
			ID:      h.ID,
			AppID:   h.AppID,
			LangID:  h.LangID,
			UsedFor: h.UsedFor,
			Subject: h.Subject,
			Message: h.Message,
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
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createSMSTemplate(ctx, cli); err != nil {
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

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.AppID = req.AppID
			handler.LangID = req.LangID
			if err := handler.createSMSTemplate(ctx, cli); err != nil {
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
	h.Limit = uint32(len(ids))

	infos, _, err := h.GetSMSTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
