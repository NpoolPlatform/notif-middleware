package email

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"

	"github.com/google/uuid"
)

type Handler struct {
	ID                *uuid.UUID
	AppID             *uuid.UUID
	LangID            *uuid.UUID
	DefaultToUsername *string
	UsedFor           *basetypes.UsedFor
	Sender            *string
	ReplyTos          *[]string
	CcTos             *[]string
	Subject           *string
	Body              *string
	Reqs              []*emailtemplatecrud.Req
	Conds             *emailtemplatecrud.Conds
	Offset            int32
	Limit             int32
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

func WithDefaultToUsername(defaultToUsername *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if defaultToUsername == nil {
			return nil
		}
		if *defaultToUsername == "" {
			return fmt.Errorf("invalid defaultToUsername")
		}
		h.DefaultToUsername = defaultToUsername
		return nil
	}
}

// nolint:gocyclo
func WithUsedFor(_usedFor *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _usedFor == nil {
			return nil
		}
		switch *_usedFor {
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
			return fmt.Errorf("invalid UsedFor")
		}
		h.UsedFor = _usedFor
		return nil
	}
}

func WithSender(sender *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if sender == nil {
			return nil
		}
		if *sender == "" {
			return fmt.Errorf("invalid sender")
		}
		h.Sender = sender
		return nil
	}
}

func WithReplyTos(replyTos *[]string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if replyTos == nil {
			return nil
		}
		if len(*replyTos) == 0 {
			return fmt.Errorf("invalid replyTos")
		}
		h.ReplyTos = replyTos
		return nil
	}
}

func WithCcTos(ccTos *[]string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ccTos == nil {
			return nil
		}
		if len(*ccTos) == 0 {
			return fmt.Errorf("invalid ccTos")
		}
		h.CcTos = ccTos
		return nil
	}
}

func WithSubject(subject *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if subject == nil {
			return nil
		}
		if *subject == "" {
			return fmt.Errorf("invalid subject")
		}
		h.Subject = subject
		return nil
	}
}

func WithBody(body *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if body == nil {
			return nil
		}
		if *body == "" {
			return fmt.Errorf("invalid body")
		}
		h.Body = body
		return nil
	}
}

// nolint:gocyclo
func WithReqs(reqs []*npool.EmailTemplateReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*emailtemplatecrud.Req{}
		for _, req := range reqs {
			_req := &emailtemplatecrud.Req{}
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
			if req.UsedFor != nil {
				switch req.GetUsedFor() {
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
					return fmt.Errorf("invalid UsedFor")
				}
				_req.UsedFor = req.UsedFor
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
		h.Conds = &emailtemplatecrud.Conds{}
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
		if conds.UsedFor != nil {
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
				return fmt.Errorf("invalid UsedFor")
			}
			h.Conds.UsedFor = &cruder.Cond{
				Op:  conds.GetUsedFor().GetOp(),
				Val: conds.GetUsedFor().GetValue(),
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
