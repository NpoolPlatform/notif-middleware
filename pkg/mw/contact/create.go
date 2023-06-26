package contact

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("app id is empty")
	}
	if h.Account == nil {
		return fmt.Errorf("account is empty")
	}
	if h.Sender == nil {
		return fmt.Errorf("sender is empty")
	}
	if h.AccountType == nil {
		return fmt.Errorf("account type is empty")
	}
	if h.UsedFor == nil {
		return fmt.Errorf("used for is empty")
	}
	return nil
}

func (h *Handler) CreateContact(ctx context.Context) (info *npool.Contact, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	h.Conds = &crud.Conds{
		AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		AccountType: &cruder.Cond{Op: cruder.EQ, Val: basetypes.SignMethod(basetypes.SignMethod_value[h.AccountType.String()])},
		UsedFor:     &cruder.Cond{Op: cruder.EQ, Val: basetypes.UsedFor(basetypes.UsedFor_value[h.UsedFor.String()])},
	}
	exist, err := h.ExistContactConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("contact exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.Contact.Create(),
			&crud.Req{
				ID:          h.ID,
				AppID:       h.AppID,
				Account:     h.Account,
				AccountType: h.AccountType,
				UsedFor:     h.UsedFor,
				Sender:      h.Sender,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetContact(ctx)
}
