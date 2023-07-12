package notif

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	templatemwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	UserID      *uuid.UUID
	LangID      *uuid.UUID
	EventID     *uuid.UUID
	Notified    *bool
	EventType   *basetypes.UsedFor
	UseTemplate *bool
	Title       *string
	Content     *string
	Channel     *basetypes.NotifChannel
	Extra       *string
	NotifType   *basetypes.NotifType
	Vars        *templatemwpb.TemplateVars
	IDs         *[]uuid.UUID
	Reqs        []*notifcrud.Req
	Conds       *notifcrud.Conds
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
		if userid == nil || *userid == "" {
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

func WithLangID(langid *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if langid == nil {
			return nil
		}
		_langid, err := uuid.Parse(*langid)
		if err != nil {
			return err
		}
		h.LangID = &_langid
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

func WithNotified(notified *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Notified = notified
		return nil
	}
}

func WithUseTemplate(usetemplate *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UseTemplate = usetemplate
		return nil
	}
}

func WithTitle(title *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if title == nil {
			return nil
		}
		if *title == "" {
			return fmt.Errorf("invalid title")
		}
		h.Title = title
		return nil
	}
}

func WithContent(content *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if content == nil {
			return nil
		}
		if *content == "" {
			return fmt.Errorf("invalid content")
		}
		h.Content = content
		return nil
	}
}

func WithChannel(_channel *basetypes.NotifChannel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _channel == nil {
			return nil
		}
		switch *_channel {
		case basetypes.NotifChannel_ChannelFrontend:
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		default:
			return fmt.Errorf("invalid channel")
		}
		h.Channel = _channel
		return nil
	}
}

func WithExtra(extra *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			return nil
		}
		if *extra == "" {
			return fmt.Errorf("invalid extra")
		}
		h.Extra = extra
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
		case basetypes.UsedFor_GoodBenefit1:
		default:
			return fmt.Errorf("invalid eventtype")
		}
		h.EventType = eventtype
		return nil
	}
}

func WithNotifType(_type *basetypes.NotifType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case basetypes.NotifType_NotifMulticast:
		case basetypes.NotifType_NotifUnicast:
		default:
			return fmt.Errorf("invalid type")
		}
		h.NotifType = _type
		return nil
	}
}

func WithVars(vars *templatemwpb.TemplateVars) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Vars = vars
		return nil
	}
}

func WithIDs(ids *[]string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ids == nil {
			return nil
		}
		if len(*ids) == 0 {
			return fmt.Errorf("invalid ids")
		}
		_reqs := []*notifcrud.Req{}
		for _, id := range *ids {
			_req := &notifcrud.Req{}
			if id != "" {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_req.ID = &_id
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

// nolint:gocyclo
func WithReqs(reqs []*npool.NotifReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*notifcrud.Req{}
		for _, req := range reqs {
			_req := &notifcrud.Req{}
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
			if req.LangID != nil {
				id, err := uuid.Parse(req.GetLangID())
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.EventID != nil {
				id, err := uuid.Parse(req.GetEventID())
				if err != nil {
					return err
				}
				_req.EventID = &id
			}
			if req.NotifType != nil {
				switch req.GetNotifType() {
				case basetypes.NotifType_NotifMulticast:
				case basetypes.NotifType_NotifUnicast:
				default:
					return fmt.Errorf("invalid Type")
				}
				_req.NotifType = req.NotifType
			}
			if req.EventType != nil {
				switch req.GetEventType() {
				case basetypes.UsedFor_WithdrawalRequest:
				case basetypes.UsedFor_WithdrawalCompleted:
				case basetypes.UsedFor_DepositReceived:
				case basetypes.UsedFor_KYCApproved:
				case basetypes.UsedFor_KYCRejected:
				case basetypes.UsedFor_Announcement:
				case basetypes.UsedFor_GoodBenefit1:
				default:
					return fmt.Errorf("invalid EventType")
				}
				_req.EventType = req.EventType
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
			if req.UseTemplate != nil {
				_req.UseTemplate = req.UseTemplate
			}
			if req.Title != nil {
				_req.Title = req.Title
			}
			if req.Content != nil {
				_req.Content = req.Content
			}
			if req.Extra != nil {
				_req.Extra = req.Extra
			}
			if req.Notified != nil {
				_req.Notified = req.Notified
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

// nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &notifcrud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{Op: conds.GetLangID().GetOp(), Val: id}
		}
		if conds.EventID != nil {
			id, err := uuid.Parse(conds.GetEventID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EventID = &cruder.Cond{Op: conds.GetEventID().GetOp(), Val: id}
		}
		if conds.NotifType != nil {
			switch conds.GetNotifType().GetValue() {
			case uint32(basetypes.NotifType_NotifMulticast):
			case uint32(basetypes.NotifType_NotifUnicast):
			default:
				return fmt.Errorf("invalid Type")
			}
			h.Conds.Type = &cruder.Cond{Op: conds.GetNotifType().GetOp(), Val: conds.GetNotifType().GetValue()}
		}
		if conds.EventType != nil {
			switch conds.GetEventType().GetValue() {
			case uint32(basetypes.UsedFor_WithdrawalRequest):
			case uint32(basetypes.UsedFor_WithdrawalCompleted):
			case uint32(basetypes.UsedFor_DepositReceived):
			case uint32(basetypes.UsedFor_KYCApproved):
			case uint32(basetypes.UsedFor_KYCRejected):
			case uint32(basetypes.UsedFor_Announcement):
			case uint32(basetypes.UsedFor_GoodBenefit1):
			default:
				return fmt.Errorf("invalid EventType")
			}
			_type := conds.GetEventType().GetValue()
			h.Conds.EventType = &cruder.Cond{Op: conds.GetEventType().GetOp(), Val: basetypes.UsedFor(_type)}
		}
		if conds.Channel != nil {
			switch conds.GetChannel().GetValue() {
			case uint32(basetypes.NotifChannel_ChannelFrontend):
			case uint32(basetypes.NotifChannel_ChannelEmail):
			case uint32(basetypes.NotifChannel_ChannelSMS):
			default:
				return fmt.Errorf("invalid channel")
			}
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel)}
		}
		if conds.IDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: ids}
		}
		if conds.Notified != nil {
			h.Conds.Notified = &cruder.Cond{Op: conds.GetNotified().GetOp(), Val: conds.GetNotified().GetValue()}
		}
		if conds.Extra != nil {
			h.Conds.Extra = &cruder.Cond{Op: conds.GetExtra().GetOp(), Val: conds.GetExtra().GetValue()}
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
