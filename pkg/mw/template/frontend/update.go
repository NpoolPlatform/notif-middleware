package frontend

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/frontend"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateFrontendTemplate(ctx context.Context, cli *ent.Client) error {
	if _, err := frontendtemplatecrud.UpdateSet(
		cli.FrontendTemplate.UpdateOneID(*h.ID),
		&frontendtemplatecrud.Req{
			Title:   h.Title,
			Content: h.Content,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateFrontendTemplate(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFrontendTemplate(ctx)
}
