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

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createFrontendTemplate(ctx context.Context, tx *ent.Tx, req *frontendtemplatecrud.Req) error {
	if req.AppID == nil {
		return fmt.Errorf("invalid lang")
	}
	if req.LangID == nil {
		return fmt.Errorf("invalid logo")
	}
	if req.UsedFor == nil {
		return fmt.Errorf("invalid create usedFor")
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

	h.Conds = &frontendtemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *req.UsedFor},
	}
	exist, err := h.ExistFrontendTemplateConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("frontendtemplate exist")
	}

	id := uuid.New()
	if req.ID == nil {
		req.ID = &id
	}

	info, err := frontendtemplatecrud.CreateSet(
		tx.FrontendTemplate.Create(),
		&frontendtemplatecrud.Req{
			ID:      req.ID,
			AppID:   req.AppID,
			LangID:  req.LangID,
			UsedFor: req.UsedFor,
			Title:   req.Title,
			Content: req.Content,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	req := &frontendtemplatecrud.Req{
		ID:      handler.ID,
		AppID:   handler.AppID,
		LangID:  handler.LangID,
		UsedFor: handler.UsedFor,
		Title:   handler.Title,
		Content: handler.Content,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createFrontendTemplate(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFrontendTemplate(ctx)
}

func (h *Handler) CreateFrontendTemplates(ctx context.Context) ([]*npool.FrontendTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.createFrontendTemplate(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &frontendtemplatecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetFrontendTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
