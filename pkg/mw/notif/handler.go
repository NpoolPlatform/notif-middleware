package notif

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
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
	EventType   *string
	UseTemplate *bool
	Title       *string
	Content     *string
	Channel     *basetypes.Channel
	Extra       *string
	Type        *string
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

func WithChannel(_channel *basetypes.Channel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _channel == nil {
			return nil
		}
		switch *_channel {
		case basetypes.NotifChannel_ChannelFrontend:
		case basetypes.NotifChannel_ChannelEmail:
		case basetypes.NotifChannel_ChannelSMS:
		default:
			return fmt.Errorf("Invalid channel")
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
		case basetypes.UsedFor_Signup:
		case basetypes.UsedFor_Signin:
		case basetypes.UsedFor_Update:
		case basetypes.UsedFor_Contact:
		case basetypes.UsedFor_SetWithdrawAddress:
		case basetypes.UsedFor_Withdraw:
		case basetypes.UsedFor_CreateInvitationCode:
		case basetypes.UsedFor_SetCommission:
		case basetypes.UsedFor_SetTransferTargetUser:
		case basetypes.UsedFor_Transfer:
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		default:
			return fmt.Errorf("Invalid eventtype")
		}
		h.EventType = eventtype
		return nil
	}
}

func WithType(_type *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case basetypes.UsedFor_Signup:
		case basetypes.UsedFor_Signin:
		case basetypes.UsedFor_Update:
		case basetypes.UsedFor_Contact:
		case basetypes.UsedFor_SetWithdrawAddress:
		case basetypes.UsedFor_Withdraw:
		case basetypes.UsedFor_CreateInvitationCode:
		case basetypes.UsedFor_SetCommission:
		case basetypes.UsedFor_SetTransferTargetUser:
		case basetypes.UsedFor_Transfer:
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		default:
			return fmt.Errorf("Invalid type")
		}
		h.Type = _type
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
			if req.LangID != nil {
				id, err := uuid.Parse(req.GetLangID())
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.Type != nil {
				switch req.GetType() {
				case basetypes.UsedFor_Signup:
				case basetypes.UsedFor_Signin:
				case basetypes.UsedFor_Update:
				case basetypes.UsedFor_Contact:
				case basetypes.UsedFor_SetWithdrawAddress:
				case basetypes.UsedFor_Withdraw:
				case basetypes.UsedFor_CreateInvitationCode:
				case basetypes.UsedFor_SetCommission:
				case basetypes.UsedFor_SetTransferTargetUser:
				case basetypes.UsedFor_Transfer:
				case basetypes.UsedFor_WithdrawalRequest:
				case basetypes.UsedFor_WithdrawalCompleted:
				case basetypes.UsedFor_DepositReceived:
				case basetypes.UsedFor_KYCApproved:
				case basetypes.UsedFor_KYCRejected:
				case basetypes.UsedFor_Announcement:
				default:
					return fmt.Errorf("Invalid Type")
				}
				_req.Type = req.Type
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &notifcrud.Conds{}
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
		if conds.Type != nil {
			switch conds.GetUsedFor().GetValue() {
			case int32(basetypes.UsedFor_Signup):
			case int32(basetypes.UsedFor_Signin):
			case int32(basetypes.UsedFor_Update):
			case int32(basetypes.UsedFor_Contact):
			case int32(basetypes.UsedFor_SetWithdrawAddress):
			case int32(basetypes.UsedFor_Withdraw):
			case int32(basetypes.UsedFor_CreateInvitationCode):
			case int32(basetypes.UsedFor_SetCommission):
			case int32(basetypes.UsedFor_SetTransferTargetUser):
			case int32(basetypes.UsedFor_Transfer):
			case int32(basetypes.UsedFor_WithdrawalRequest):
			case int32(basetypes.UsedFor_WithdrawalCompleted):
			case int32(basetypes.UsedFor_DepositReceived):
			case int32(basetypes.UsedFor_KYCApproved):
			case int32(basetypes.UsedFor_KYCRejected):
			case int32(basetypes.UsedFor_Announcement):
			default:
				return fmt.Errorf("Invalid Type")
			}
			h.Conds.Type = &cruder.Cond{
				Op:  conds.GetType().GetOp(),
				Val: conds.GetType().GetValue(),
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
		h.Limit = limit
		return nil
	}
}
