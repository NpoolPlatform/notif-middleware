package goodbenefit

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/goodbenefit"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	GoodID      *uuid.UUID
	GoodName    *string
	Amount      *string
	State       *basetypes.Result
	Message     *string
	BenefitDate *uint32
	TxID        *uuid.UUID
	Notified    *bool
	DeletedAt   *uint32
	Conds       *crud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
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

func WithGoodID(goodID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_goodID, err := uuid.Parse(*goodID)
		if err != nil {
			return err
		}
		h.GoodID = &_goodID
		return nil
	}
}

func WithGoodName(goodName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodName == nil || *goodName == "" {
			return fmt.Errorf("invalid good name")
		}
		h.GoodName = goodName
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		h.Amount = amount
		return nil
	}
}

func WithState(state *basetypes.Result) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return fmt.Errorf("state is empty")
		}
		switch *state {
		case basetypes.Result_Fail:
		case basetypes.Result_Success:
		default:
			return fmt.Errorf("invalid state %s", *state)
		}

		h.State = state
		return nil
	}
}

func WithMessage(message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil || *message == "" {
			return fmt.Errorf("invalid message")
		}
		h.Message = message
		return nil
	}
}

func WithNotified(notified *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Notified = notified
		return nil
	}
}

func WithTxID(txID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if txID == nil {
			return nil
		}
		_txID, err := uuid.Parse(*txID)
		if err != nil {
			return err
		}
		h.TxID = &_txID
		return nil
	}
}

// nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.GoodID != nil {
			goodID, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: goodID}
		}

		if conds.Notified != nil {
			h.Conds.Notified = &cruder.Cond{Op: conds.GetNotified().GetOp(), Val: conds.GetNotified().GetValue()}
		}

		if conds.BenefitDate != nil {
			h.Conds.BenefitDate = &cruder.Cond{Op: conds.GetBenefitDate().GetOp(), Val: conds.GetNotified().GetValue()}
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
