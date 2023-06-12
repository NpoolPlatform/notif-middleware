package handler

import (
	"context"
	"fmt"

	appcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
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
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}

		exist, err := appcli.ExistApp(ctx, *appID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}

		h.AppID = &_appID
		return nil
	}
}

func WithUserID(appID, userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_userID, err := uuid.Parse(*userID)
		if err != nil {
			return err
		}

		exist, err := appusercli.ExistUser(ctx, *appID, *userID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid user")
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

		handler, err := amt1.NewHandler(ctx,
			amt1.WithID(amtID),
			amt1.WithAppID(appID),
		)
		if err != nil {
			return err
		}

		exist, err := handler.ExistAnnouncement(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid announcement id")
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
