package email

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	entemailtemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/emailtemplate"

	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"
)

type queryHandler struct {
	*Handler
	stm   *ent.EmailTemplateSelect
	infos []*npool.EmailTemplate
	total uint32
}

func (h *queryHandler) selectEmailTemplate(stm *ent.EmailTemplateQuery) {
	h.stm = stm.Select(
		entemailtemplate.FieldID,
		entemailtemplate.FieldAppID,
		entemailtemplate.FieldLangID,
		entemailtemplate.FieldUsedFor,
		entemailtemplate.FieldSender,
		entemailtemplate.FieldReplyTos,
		entemailtemplate.FieldCcTos,
		entemailtemplate.FieldSubject,
		entemailtemplate.FieldBody,
		entemailtemplate.FieldCreatedAt,
		entemailtemplate.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryEmailTemplate(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid emailtemplateid")
	}

	h.selectEmailTemplate(
		cli.EmailTemplate.
			Query().
			Where(
				entemailtemplate.ID(*h.ID),
				entemailtemplate.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryEmailTemplates(ctx context.Context, cli *ent.Client) error {
	stm, err := emailtemplatecrud.SetQueryConds(cli.EmailTemplate.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectEmailTemplate(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplate(cli); err != nil {
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

func (h *Handler) GetEmailTemplates(ctx context.Context) ([]*npool.EmailTemplate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplates(ctx, cli); err != nil {
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

func (h *Handler) GetEmailTemplateOnly(ctx context.Context) (info *npool.EmailTemplate, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEmailTemplates(_ctx, cli); err != nil {
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
