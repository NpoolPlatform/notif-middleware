package user

import (
	"context"
	"fmt"

	appcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/user"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds *crud.Conds
	Reqs  []*crud.Req
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Handler: _handler,
	}

	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, h); err != nil {
			return nil, err
		}
	}
	return h, nil
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
		if conds.AnnouncementID != nil {
			id, err := uuid.Parse(conds.GetAnnouncementID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AnnouncementID = &cruder.Cond{
				Op: conds.GetAnnouncementID().GetOp(), Val: id,
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.AnnouncementUserReq) func(context.Context, *Handler) error {
	return func(_ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range _reqs {
			if req.AppID == nil || req.UserID == nil || req.AnnouncementID == nil {
				continue
			}

			// AppID
			_req := &crud.Req{}
			appID, err := uuid.Parse(req.AppID.String())
			if err != nil {
				return err
			}
			exist, err := appcli.ExistApp(_ctx, req.AppID.String())
			if err != nil {
				return err
			}
			if !exist {
				return fmt.Errorf("invalid app id %v", req.AppID.String())
			}
			_req.AppID = &appID

			// UserID
			userID, err := uuid.Parse(req.UserID.String())
			if err != nil {
				return err
			}

			exist, err = appusercli.ExistUser(_ctx, req.AppID.String(), req.UserID.String())
			if err != nil {
				return err
			}
			if !exist {
				return fmt.Errorf("invalid user")
			}
			_req.UserID = &userID

			// AnnouncementID
			amtID, err := uuid.Parse(req.AnnouncementID.String())
			if err != nil {
				return err
			}
			_amtID := req.AnnouncementID.String()
			amtHandler, err := amt1.NewHandler(_ctx, amt1.WithID(&_amtID))
			if err != nil {
				return err
			}

			exist, err = amtHandler.ExistAnnouncement(_ctx)
			if err != nil {
				return err
			}
			if !exist {
				return fmt.Errorf("invalid user")
			}
			_req.AnnouncementID = &amtID

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
