package email

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"
)

func (h *Handler) DeleteEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetEmailTemplate(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &emailtemplatecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistEmailTemplateConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("emailtemplate not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := emailtemplatecrud.UpdateSet(
			cli.EmailTemplate.UpdateOneID(*h.ID),
			&emailtemplatecrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
