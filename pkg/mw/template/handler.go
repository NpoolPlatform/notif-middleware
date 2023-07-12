package template

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	templatemwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	"github.com/google/uuid"
)

type Handler struct {
	AppID   *uuid.UUID
	UserID  *uuid.UUID
	UsedFor *basetypes.UsedFor
	LangID  *uuid.UUID
	Vars    *templatemwpb.TemplateVars
	Channel *basetypes.NotifChannel
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
		case basetypes.UsedFor_GoodBenefit1:
		default:
			return fmt.Errorf("invalid %v usedfor", *_usedFor)
		}
		h.UsedFor = _usedFor
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

func WithVars(vars *templatemwpb.TemplateVars) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Vars = vars
		return nil
	}
}
