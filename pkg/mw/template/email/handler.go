package email

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	templatemwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"

	"github.com/google/uuid"
)

type Handler struct {
	ID                *uint32
	EntID             *uuid.UUID
	AppID             *uuid.UUID
	LangID            *uuid.UUID
	DefaultToUsername *string
	UsedFor           *basetypes.UsedFor
	Sender            *string
	ReplyTos          *[]string
	CcTos             *[]string
	Subject           *string
	Body              *string
	UserID            *uuid.UUID
	Vars              *templatemwpb.TemplateVars
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(appid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appid == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
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

func WithLangID(langid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if langid == nil {
			if must {
				return fmt.Errorf("invalid langid")
			}
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

func WithDefaultToUsername(defaultToUsername *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if defaultToUsername == nil {
			if must {
				return fmt.Errorf("invalid defaulttousername")
			}
			return nil
		}
		if *defaultToUsername == "" {
			return fmt.Errorf("invalid defaultToUsername")
		}
		h.DefaultToUsername = defaultToUsername
		return nil
	}
}

// nolint
func WithUsedFor(_usedFor *basetypes.UsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
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
		case basetypes.UsedFor_GoodBenefit1:
		case basetypes.UsedFor_UpdateEmail:
		case basetypes.UsedFor_UpdateMobile:
		case basetypes.UsedFor_UpdatePassword:
		case basetypes.UsedFor_UpdateGoogleAuth:
		case basetypes.UsedFor_NewLogin:
		case basetypes.UsedFor_OrderCompleted:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = _usedFor
		return nil
	}
}

func WithSender(sender *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if sender == nil {
			if must {
				return fmt.Errorf("invalid sender")
			}
			return nil
		}
		if *sender == "" {
			return fmt.Errorf("invalid sender")
		}
		h.Sender = sender
		return nil
	}
}

func WithReplyTos(replyTos *[]string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if replyTos == nil {
			if must {
				return fmt.Errorf("invalid replyTos")
			}
			return nil
		}
		if len(*replyTos) == 0 {
			return nil
		}
		h.ReplyTos = replyTos
		return nil
	}
}

func WithCcTos(ccTos *[]string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ccTos == nil {
			if must {
				return fmt.Errorf("invalid cctos")
			}
			return nil
		}
		if len(*ccTos) == 0 {
			return nil
		}
		h.CcTos = ccTos
		return nil
	}
}

func WithSubject(subject *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if subject == nil {
			if must {
				return fmt.Errorf("invalid subject")
			}
			return nil
		}
		if *subject == "" {
			return fmt.Errorf("invalid subject")
		}
		h.Subject = subject
		return nil
	}
}

func WithBody(body *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if body == nil {
			if must {
				return fmt.Errorf("invalid body")
			}
			return nil
		}
		if *body == "" {
			return fmt.Errorf("invalid body")
		}
		h.Body = body
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithVars(vars *templatemwpb.TemplateVars, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if vars == nil {
			if must {
				return fmt.Errorf("invalid vars")
			}
			return nil
		}
		h.Vars = vars
		return nil
	}
}

//nolint
func WithReqs(reqs []*npool.EmailTemplateReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*emailtemplatecrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.LangID == nil {
					return fmt.Errorf("invalid langid")
				}
				if req.UsedFor == nil {
					return fmt.Errorf("invalid usedfor")
				}
			}
			_req := &emailtemplatecrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
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
				case basetypes.UsedFor_GoodBenefit1:
				case basetypes.UsedFor_UpdateEmail:
				case basetypes.UsedFor_UpdateMobile:
				case basetypes.UsedFor_UpdatePassword:
				case basetypes.UsedFor_UpdateGoogleAuth:
				case basetypes.UsedFor_NewLogin:
				case basetypes.UsedFor_OrderCompleted:
				default:
					return fmt.Errorf("invalid usedfor")
				}
				_req.UsedFor = req.UsedFor
			}
			if req.Sender != nil {
				_req.Sender = req.Sender
			}
			if req.ReplyTos != nil {
				_req.ReplyTos = &req.ReplyTos
			}
			if req.CCTos != nil {
				_req.CcTos = &req.CCTos
			}
			if req.Subject != nil {
				_req.Subject = req.Subject
			}
			if req.Body != nil {
				_req.Body = req.Body
			}
			if req.DefaultToUsername != nil {
				_req.DefaultToUsername = req.DefaultToUsername
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:funlen,dupl,gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &emailtemplatecrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
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
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{
				Op:  conds.GetLangID().GetOp(),
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
			case int32(basetypes.UsedFor_GoodBenefit1):
			case int32(basetypes.UsedFor_UpdateEmail):
			case int32(basetypes.UsedFor_UpdateMobile):
			case int32(basetypes.UsedFor_UpdatePassword):
			case int32(basetypes.UsedFor_UpdateGoogleAuth):
			case int32(basetypes.UsedFor_NewLogin):
			case int32(basetypes.UsedFor_OrderCompleted):
			default:
				return fmt.Errorf("invalid usedfor")
			}
			usedFor := conds.GetUsedFor().GetValue()
			h.Conds.UsedFor = &cruder.Cond{
				Op:  conds.GetUsedFor().GetOp(),
				Val: basetypes.UsedFor(usedFor),
			}
		}
		if conds.AppIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppIDs = &cruder.Cond{Op: conds.GetAppIDs().GetOp(), Val: ids}
		}
		if conds.LangIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetLangIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.LangIDs = &cruder.Cond{Op: conds.GetLangIDs().GetOp(), Val: ids}
		}
		if conds.UsedFors != nil {
			usedFors := []string{}
			for _, usedFor := range conds.GetUsedFors().GetValue() {
				switch usedFor {
				case uint32(basetypes.UsedFor_Signup):
				case uint32(basetypes.UsedFor_Signin):
				case uint32(basetypes.UsedFor_Update):
				case uint32(basetypes.UsedFor_Contact):
				case uint32(basetypes.UsedFor_SetWithdrawAddress):
				case uint32(basetypes.UsedFor_Withdraw):
				case uint32(basetypes.UsedFor_CreateInvitationCode):
				case uint32(basetypes.UsedFor_SetCommission):
				case uint32(basetypes.UsedFor_SetTransferTargetUser):
				case uint32(basetypes.UsedFor_Transfer):
				case uint32(basetypes.UsedFor_WithdrawalRequest):
				case uint32(basetypes.UsedFor_WithdrawalCompleted):
				case uint32(basetypes.UsedFor_DepositReceived):
				case uint32(basetypes.UsedFor_KYCApproved):
				case uint32(basetypes.UsedFor_KYCRejected):
				case uint32(basetypes.UsedFor_Announcement):
				case uint32(basetypes.UsedFor_GoodBenefit1):
				case uint32(basetypes.UsedFor_UpdateEmail):
				case uint32(basetypes.UsedFor_UpdateMobile):
				case uint32(basetypes.UsedFor_UpdatePassword):
				case uint32(basetypes.UsedFor_UpdateGoogleAuth):
				case uint32(basetypes.UsedFor_NewLogin):
				case uint32(basetypes.UsedFor_OrderCompleted):
				default:
					return fmt.Errorf("invalid usedfor")
				}
				_usedFor := basetypes.UsedFor(usedFor).String()
				usedFors = append(usedFors, _usedFor)
			}
			h.Conds.UsedFors = &cruder.Cond{Op: conds.GetUsedFors().GetOp(), Val: usedFors}
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
