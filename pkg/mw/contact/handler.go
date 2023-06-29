package contact

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/contact"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	UsedFor     *basetypes.UsedFor
	AccountType *basetypes.SignMethod
	Account     *string
	Sender      *string
	Conds       *crud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			return fmt.Errorf("app id is empty")
		}
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		h.AppID = &_appID
		return nil
	}
}

func WithAccount(account *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == nil {
			return nil
		}
		if *account == "" {
			return fmt.Errorf("account is empty")
		}
		h.Account = account
		return nil
	}
}

func WithSender(sender *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if sender == nil {
			return nil
		}
		if *sender == "" {
			return fmt.Errorf("sender is empty")
		}
		h.Sender = sender
		return nil
	}
}

func WithUsedFor(usedFor *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch *usedFor {
		case basetypes.UsedFor_Contact:
		default:
			return fmt.Errorf("used for %v invalid", *usedFor)
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithAccountType(_type *basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case basetypes.SignMethod_Email:
		default:
			return fmt.Errorf("type %v invalid", *_type)
		}
		h.AccountType = _type
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.AccountType != nil {
			accountType := conds.GetAccountType().GetValue()
			h.Conds.AccountType = &cruder.Cond{Op: conds.GetAccountType().GetOp(), Val: basetypes.SignMethod(accountType)}
		}
		if conds.UsedFor != nil {
			usedFor := conds.GetUsedFor().GetValue()
			h.Conds.UsedFor = &cruder.Cond{Op: conds.GetUsedFor().GetOp(), Val: basetypes.UsedFor(usedFor)}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
