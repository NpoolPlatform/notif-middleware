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
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateSMSTemplate(ctx context.Context, cli *ent.Client) error {
	if _, err := smstemplatecrud.UpdateSet(
		cli.SMSTemplate.UpdateOneID(*h.ID),
		&smstemplatecrud.Req{
			Subject: h.Subject,
			Message: h.Message,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSMSTemplate(ctx context.Context) (*npool.SMSTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.LangID == nil {
		return nil, fmt.Errorf("invalid langid")
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

	h.Conds = &smstemplatecrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
	}
	h.Offset = 0
	h.Limit = 2

	sms, err := h.GetSMSTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if sms != nil {
		return nil, fmt.Errorf("smstemplate exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateSMSTemplate(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSMSTemplate(ctx)
}
