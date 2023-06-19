package sendstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/sendstate"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Channel *basetypes.NotifChannel
	Reqs    []*crud.Req
	Conds   *crud.Conds
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

func WithChannel(channel *basetypes.NotifChannel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
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
				Op: conds.GetAppID().GetOp(), Val: id,
			}
		}
		if conds.UserID != nil {
			userID, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op: conds.GetUserID().GetOp(), Val: userID,
			}
		}

		if conds.AnnouncementID != nil {
			amtID, err := uuid.Parse(conds.GetAnnouncementID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AnnouncementID = &cruder.Cond{
				Op: conds.GetAnnouncementID().GetOp(), Val: amtID,
			}
		}
		if conds.Channel != nil {
			channel := conds.GetChannel().GetValue()
			h.Conds.Channel = &cruder.Cond{
				Op: conds.GetChannel().GetOp(), Val: basetypes.NotifChannel(channel),
			}
		}
		if conds.UserIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetUserIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.UserIDs = &cruder.Cond{
				Op: conds.GetUserIDs().GetOp(), Val: ids,
			}
		}
		return nil
	}
}

func WithReqs(reqs []*npool.SendStateReq) func(context.Context, *Handler) error {
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
			_req.AppID = &appID

			// UserID
			userID, err := uuid.Parse(req.UserID.String())
			if err != nil {
				return err
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
			_req.AnnouncementID = &amtID

			// Channel
			info, err := amtHandler.GetAnnouncement(_ctx)
			if err != nil {
				return err
			}
			_req.Channel = &info.Channel

			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
