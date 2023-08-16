package sms

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	smstemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/sms"

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

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
