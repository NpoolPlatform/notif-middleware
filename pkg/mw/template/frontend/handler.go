package frontend

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/frontend"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uuid.UUID
	AppID   *uuid.UUID
	LangID  *uuid.UUID
	UsedFor *basetypes.UsedFor
	Title   *string
	Content *string
	Reqs    []*frontendtemplatecrud.Req
	Conds   *frontendtemplatecrud.Conds
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

func WithUsedFor(_usedFor *basetypes.UsedFor) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _usedFor == nil {
			return nil
		}
		switch *_usedFor {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		default:
			return fmt.Errorf("Invalid UsedFor")
		}
		h.UsedFor = _usedFor
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

// nolint:gocyclo
func WithReqs(reqs []*npool.FrontendTemplateReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*frontendtemplatecrud.Req{}
		for _, req := range reqs {
			_req := &frontendtemplatecrud.Req{}
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
				case basetypes.UsedFor_WithdrawalRequest:
				case basetypes.UsedFor_WithdrawalCompleted:
				case basetypes.UsedFor_DepositReceived:
				case basetypes.UsedFor_KYCApproved:
				case basetypes.UsedFor_KYCRejected:
				case basetypes.UsedFor_Announcement:
				default:
					return fmt.Errorf("Invalid UsedFor")
				}
				_req.UsedFor = req.UsedFor
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &frontendtemplatecrud.Conds{}
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
			case uint32(basetypes.UsedFor_WithdrawalRequest):
			case uint32(basetypes.UsedFor_WithdrawalCompleted):
			case uint32(basetypes.UsedFor_DepositReceived):
			case uint32(basetypes.UsedFor_KYCApproved):
			case uint32(basetypes.UsedFor_KYCRejected):
			case uint32(basetypes.UsedFor_Announcement):
			default:
				return fmt.Errorf("Invalid UsedFor")
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
