package sendstate

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement/handler"
	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds *crud.Conds
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
		if conds.AnnouncementID != nil {
			id, err := uuid.Parse(conds.GetAnnouncementID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AnnouncementID = &cruder.Cond{
				Op: conds.GetAnnouncementID().GetOp(), Val: id,
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
