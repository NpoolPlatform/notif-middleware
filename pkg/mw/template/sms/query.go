package sms

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/sms"
	entsmstemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/smstemplate"

	smstemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/sms"
)

type queryHandler struct {
	*Handler
	stm   *ent.SMSTemplateSelect
	infos []*npool.SMSTemplate
	total uint32
}

func (h *queryHandler) selectSMSTemplate(stm *ent.SMSTemplateQuery) {
	h.stm = stm.Select(
		entsmstemplate.FieldID,
		entsmstemplate.FieldAppID,
		entsmstemplate.FieldLangID,
		entsmstemplate.FieldUsedFor,
		entsmstemplate.FieldSubject,
		entsmstemplate.FieldMessage,
	)
}

func (h *queryHandler) querySMSTemplate(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid smstemplateid")
	}

	h.selectSMSTemplate(
		cli.SMSTemplate.
			Query().
			Where(
				entsmstemplate.ID(*h.ID),
				entsmstemplate.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) querySMSTemplates(ctx context.Context, cli *ent.Client) error {
	stm, err := smstemplatecrud.SetQueryConds(cli.SMSTemplate.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectSMSTemplate(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSMSTemplate(ctx context.Context) (*npool.SMSTemplate, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySMSTemplate(cli); err != nil {
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

func (h *Handler) GetSMSTemplates(ctx context.Context) ([]*npool.SMSTemplate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySMSTemplates(ctx, cli); err != nil {
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

func (h *Handler) GetSMSTemplateOnly(ctx context.Context) (info *npool.SMSTemplate, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySMSTemplates(_ctx, cli); err != nil {
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
