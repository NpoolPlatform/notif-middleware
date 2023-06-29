package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	EventType *basetypes.UsedFor
	Reqs      []*usercrud.Req
	Conds     *usercrud.Conds
	Offset    int32
	Limit     int32
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

func WithAppID(appid *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appid == nil {
			return nil
		}
		_appid, err := uuid.Parse(*appid)
		if err != nil {
			return err
		}
		h.AppID = &_appid
		return nil
	}
}

func WithUserID(userid *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userid == nil {
			return nil
		}
		_userid, err := uuid.Parse(*userid)
		if err != nil {
			return err
		}
		h.UserID = &_userid
		return nil
	}
}

func WithEventType(eventtype *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventtype == nil {
			return nil
		}
		switch *eventtype {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		case basetypes.UsedFor_GoodBenefit:
		default:
			return fmt.Errorf("invalid eventtype")
		}
		h.EventType = eventtype
		return nil
	}
}

//nolint
func WithReqs(reqs []*npool.NotifUserReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*usercrud.Req{}
		for _, req := range reqs {
			_req := &usercrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(req.GetID())
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(req.GetAppID())
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(req.GetUserID())
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.EventType != nil {
				switch req.GetEventType() {
				case basetypes.UsedFor_WithdrawalRequest:
				case basetypes.UsedFor_WithdrawalCompleted:
				case basetypes.UsedFor_DepositReceived:
				case basetypes.UsedFor_KYCApproved:
				case basetypes.UsedFor_KYCRejected:
				case basetypes.UsedFor_Announcement:
				case basetypes.UsedFor_GoodBenefit:
				default:
					return fmt.Errorf("invalid usedfor")
				}
				_req.EventType = req.EventType
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &usercrud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.EventType != nil {
			switch conds.GetEventType().GetValue() {
			case uint32(basetypes.UsedFor_WithdrawalRequest):
			case uint32(basetypes.UsedFor_WithdrawalCompleted):
			case uint32(basetypes.UsedFor_DepositReceived):
			case uint32(basetypes.UsedFor_KYCApproved):
			case uint32(basetypes.UsedFor_KYCRejected):
			case uint32(basetypes.UsedFor_Announcement):
			case uint32(basetypes.UsedFor_GoodBenefit):
			default:
				return fmt.Errorf("invalid usedfor")
			}
			_type := conds.GetEventType().GetValue()
			h.Conds.EventType = &cruder.Cond{
				Op:  conds.GetEventType().GetOp(),
				Val: basetypes.UsedFor(_type),
			}
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
