package contact

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/contact"
)

type queryHandler struct {
	*Handler
	stm   *ent.ContactSelect
	infos []*npool.Contact
	total uint32
}

func (h *queryHandler) selectContact(stm *ent.ContactQuery) {
	h.stm = stm.Select(
		entamt.FieldID,
		entamt.FieldAppID,
		entamt.FieldAccount,
		entamt.FieldAccountType,
		entamt.FieldUsedFor,
		entamt.FieldSender,
		entamt.FieldCreatedAt,
		entamt.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryContact(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid contact id")
	}
	h.selectContact(
		cli.Contact.
			Query().
			Where(
				entamt.ID(*h.ID),
				entamt.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryContactsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.Contact.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectContact(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetContacts(ctx context.Context) ([]*npool.Contact, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContactsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetContact(ctx context.Context) (info *npool.Contact, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContact(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetContactOnly(ctx context.Context) (*npool.Contact, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContactsByConds(_ctx, cli); err != nil {
			return err
		}

		_, err := handler.stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 0 {
		return nil, fmt.Errorf("to many record")
	}

	return handler.infos[0], nil
}
