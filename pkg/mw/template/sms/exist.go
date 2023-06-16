package sms

import (
	"context"
	"fmt"

	smstemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/sms"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entsmstemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/smstemplate"
)

func (h *Handler) ExistSMSTemplate(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			SMSTemplate.
			Query().
			Where(
				entsmstemplate.ID(*h.ID),
				entsmstemplate.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistSMSTemplateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := smstemplatecrud.SetQueryConds(cli.SMSTemplate.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
