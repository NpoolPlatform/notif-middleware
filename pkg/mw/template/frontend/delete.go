package frontend

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/frontend"
)

func (h *Handler) DeleteFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetFrontendTemplate(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &frontendtemplatecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistFrontendTemplateConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("frontendtemplate not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := frontendtemplatecrud.UpdateSet(
			cli.FrontendTemplate.UpdateOneID(*h.ID),
			&frontendtemplatecrud.Req{
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
