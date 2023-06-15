package frontend

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/frontend"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

	h.Conds = &frontendtemplatecrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
	}
	h.Offset = 0
	h.Limit = 2

	frontend, err := h.GetFrontendTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if frontend != nil {
		return nil, fmt.Errorf("frontendtemplate exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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