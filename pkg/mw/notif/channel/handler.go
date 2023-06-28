package channel

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/channel"
	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	Channel   *basetypes.NotifChannel
	EventType *basetypes.UsedFor
	Reqs      []*crud.Req
	Conds     *crud.Conds
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
			return fmt.Errorf("invalid app id")
		}
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		h.AppID = &_appID
		return nil
	}
}

func WithChannel(channel *basetypes.NotifChannel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if channel == nil {
			return fmt.Errorf("invalid channel")
		}
		switch *channel {
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		case basetypes.NotifChannel_ChannelFrontend:
		default:
			return fmt.Errorf("channel %v invalid", *channel)
		}
		h.Channel = channel
		return nil
	}
}

func WithEventType(_type *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return fmt.Errorf("invalid event type")
		}
		switch *_type {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		case basetypes.UsedFor_GoodBenefit:
		default:
			return fmt.Errorf("EventType is invalid")
		}

		h.EventType = _type
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
		if conds.AppID != nil {
			appID, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: appID}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel)}
		}
		if conds.EventType != nil {
			_type := conds.GetEventType().GetValue()
			h.Conds.EventType = &cruder.Cond{Op: conds.GetEventType().GetOp(), Val: basetypes.UsedFor(_type)}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.ChannelReq) func(context.Context, *Handler) error {
	return func(_ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range _reqs {
			if req.AppID == nil || req.Channel == nil || req.EventType == nil {
				continue
			}
			// AppID
			_req := &crud.Req{}
			appID, err := uuid.Parse(req.AppID.String())
			if err != nil {
				return err
			}
			_req.AppID = &appID

			// EventType
			switch *req.EventType {
			case basetypes.UsedFor_WithdrawalRequest:
			case basetypes.UsedFor_WithdrawalCompleted:
			case basetypes.UsedFor_DepositReceived:
			case basetypes.UsedFor_KYCApproved:
			case basetypes.UsedFor_KYCRejected:
			case basetypes.UsedFor_Announcement:
			default:
				return fmt.Errorf("EventType is invalid %v", *req.EventType)
			}
			_req.EventType = req.EventType

			switch *req.Channel {
			case basetypes.NotifChannel_ChannelEmail:
			case basetypes.NotifChannel_ChannelSMS:
			case basetypes.NotifChannel_ChannelFrontend:
			default:
				return fmt.Errorf("channel %v invalid", *req.Channel)
			}
			_req.Channel = req.Channel

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
