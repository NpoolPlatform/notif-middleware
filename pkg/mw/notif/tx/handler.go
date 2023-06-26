package tx

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/tx"
	"github.com/google/uuid"
)

type Handler struct {
	ID         *uuid.UUID
	TxID       *uuid.UUID
	NotifState *npool.TxState
	TxType     *basetypes.TxType
	Conds      *crud.Conds
	Offset     int32
	Limit      int32
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
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithTxID(txID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if txID == nil {
			return fmt.Errorf("invalid txID")
		}
		txID, err := uuid.Parse(*txID)
		if err != nil {
			return err
		}
		h.TxID = &txID
		return nil
	}
}

func WithTxType(txType *basetypes.TxType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if txType == nil {
			return fmt.Errorf("invalid tx type")
		}
		switch *txType {
		case basetypes.TxType_TxWithdraw:
		default:
			return fmt.Errorf("tx type %v invalid", *txType)
		}
		h.TxType = txType
		return nil
	}
}

func WithNotifState(state *npool.TxState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case npool.TxState_WaitSuccess:
		case npool.TxState_WaitNotified:
		case npool.TxState_Notified:
		default:
			return fmt.Errorf("tx state is invalid")
		}

		h.NotifState = state
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
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.TxID != nil {
			appID, err := uuid.Parse(conds.GetTxID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TxID = &cruder.Cond{Op: conds.GetTxID().GetOp(), Val: appID}
		}
		if conds.TxType != nil {
			_type := conds.GetTxType().GetValue()
			h.Conds.TxType = &cruder.Cond{Op: conds.GetTxType().GetOp(), Val: basetypes.TxType(_type)}
		}
		if conds.NotifState != nil {
			_state := conds.GetNotifState().GetValue()
			h.Conds.NotifState = &cruder.Cond{Op: conds.GetNotifState().GetOp(), Val: npool.TxState(_state)}
		}
		return nil
	}
}
