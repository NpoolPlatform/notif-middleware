package handler

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	AnnouncementID *uuid.UUID
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			return fmt.Errorf("invalid app id")
		}
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		h.AppID = &_appID
		return nil
	}
}

func WithUserID(userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			return fmt.Errorf("invalid user id")
		}
		_userID, err := uuid.Parse(*userID)
		if err != nil {
			return err
		}

		h.UserID = &_userID
		return nil
	}
}

func WithAnnouncementID(appID, amtID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amtID, err := uuid.Parse(*amtID)
		if err != nil {
			return err
		}

		handler, err := announcement1.NewHandler(ctx,
			announcement1.WithID(amtID),
		)
		if err != nil {
			return err
		}

		amt, err := handler.GetAnnouncement(ctx)
		if err != nil {
			return err
		}
		if amt == nil {
			return fmt.Errorf("announcement id not exist")
		}
		if amt.AppID != *appID {
			return fmt.Errorf("wrong app id or announcement id")
		}
		if amt.AnnouncementType != basetypes.NotifType_NotifMulticast {
			return fmt.Errorf("wrong announcement type %v", amt.AnnouncementType.String())
		}

		h.AnnouncementID = &_amtID
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
