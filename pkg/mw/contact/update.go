package contact

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) UpdateContact(ctx context.Context) (info *npool.Contact, err error) {
	h.Conds = &crud.Conds{
		ID: &cruder.Cond{
			Op:  cruder.EQ,
			Val: *h.ID,
		},
		AppID: &cruder.Cond{
			Op:  cruder.EQ,
			Val: *h.AppID,
		},
	}
	exist, err := h.ExistContactConds(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid id or app id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Contact.UpdateOneID(*h.ID),
			&crud.Req{
				Account:     h.Account,
				AccountType: h.AccountType,
				Sender:      h.Sender,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetContact(ctx)
}
