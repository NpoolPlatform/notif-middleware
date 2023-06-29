package contact

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entcontact "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/contact"
)

func (h *Handler) ExistContactConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(cli.Contact.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistContact(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Contact.
			Query().
			Where(
				entcontact.ID(*h.ID),
				entcontact.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
