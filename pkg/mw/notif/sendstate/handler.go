package sendstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	sendstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/sendstate"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uuid.UUID
	AppID   *uuid.UUID
	UserID  *uuid.UUID
	EventID *uuid.UUID
	Channel *basetypes.NotifChannel
	Reqs    []*sendstatecrud.Req
	Conds   *sendstatecrud.Conds
	Offset  int32
	Limit   int32
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

func WithChannel(_channel *basetypes.NotifChannel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _channel == nil {
			return nil
		}
		switch *_channel {
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		case basetypes.NotifChannel_ChannelFrontend:
		default:
			return fmt.Errorf("invalid channel")
		}
		h.Channel = _channel
		return nil
	}
}

func WithEventID(eventid *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if eventid == nil {
			return nil
		}
		_eventid, err := uuid.Parse(*eventid)
		if err != nil {
			return err
		}
		h.EventID = &_eventid
		return nil
	}
}

func WithReqs(reqs []*npool.SendStateReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*sendstatecrud.Req{}
		for _, req := range reqs {
			_req := &sendstatecrud.Req{}
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
			if req.EventID != nil {
				id, err := uuid.Parse(req.GetEventID())
				if err != nil {
					return err
				}
				_req.EventID = &id
			}
			if req.Channel != nil {
				switch req.GetChannel() {
				case basetypes.NotifChannel_ChannelEmail:
				case basetypes.NotifChannel_ChannelSMS:
				case basetypes.NotifChannel_ChannelFrontend:
				default:
					return fmt.Errorf("invalid Channel")
				}
				_req.Channel = req.Channel
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &sendstatecrud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op: conds.GetAppID().GetOp(), Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op: conds.GetUserID().GetOp(), Val: id,
			}
		}
		if conds.EventID != nil {
			id, err := uuid.Parse(conds.GetEventID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EventID = &cruder.Cond{
				Op: conds.GetEventID().GetOp(), Val: id,
			}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{
				Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel),
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
