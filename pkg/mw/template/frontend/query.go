package frontend

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	entfrontendtemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/frontendtemplate"

	frontendtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/frontend"
)

type queryHandler struct {
	*Handler
	stm   *ent.FrontendTemplateSelect
	infos []*npool.FrontendTemplate
	total uint32
}

func (h *queryHandler) selectFrontendTemplate(stm *ent.FrontendTemplateQuery) {
	h.stm = stm.Select(
		entfrontendtemplate.FieldID,
		entfrontendtemplate.FieldAppID,
		entfrontendtemplate.FieldLangID,
		entfrontendtemplate.FieldUsedFor,
		entfrontendtemplate.FieldTitle,
		entfrontendtemplate.FieldContent,
	)
}

func (h *queryHandler) queryFrontendTemplate(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid frontendtemplateid")
	}

	h.selectFrontendTemplate(
		cli.FrontendTemplate.
			Query().
			Where(
				entfrontendtemplate.ID(*h.ID),
				entfrontendtemplate.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryFrontendTemplates(ctx context.Context, cli *ent.Client) error {
	stm, err := frontendtemplatecrud.SetQueryConds(cli.FrontendTemplate.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectFrontendTemplate(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetFrontendTemplate(ctx context.Context) (*npool.FrontendTemplate, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFrontendTemplate(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetFrontendTemplates(ctx context.Context) ([]*npool.FrontendTemplate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFrontendTemplates(ctx, cli); err != nil {
			return err
		}
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetFrontendTemplateOnly(ctx context.Context) (info *npool.FrontendTemplate, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFrontendTemplates(_ctx, cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
