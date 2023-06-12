package user

import (
	"context"
	"fmt"

	appcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusercli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/user"
	amt1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	AnnouncementID *uuid.UUID
	Conds          *crud.Conds
	Offset         int32
	Limit          int32
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

func WithAnnouncementID(amtID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_amtID, err := uuid.Parse(*amtID)
		if err != nil {
			return err
		}

		handler, err := amt1.NewHandler(ctx, amt1.WithID(amtID))
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
		return nil
	}
}
