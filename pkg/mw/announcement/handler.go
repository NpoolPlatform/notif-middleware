package announcement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	constant "github.com/NpoolPlatform/notif-middleware/pkg/const"
	amtcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/google/uuid"
)

type Handler struct {
	ID      *uuid.UUID
	AppID   *uuid.UUID
	LangID  *uuid.UUID
	Title   *string
	Content *string
	Channel *basetypes.NotifChannel
	Type    *npool.AnnouncementType
	EndAt   int32
	Conds   *amtcrud.Conds
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

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_appID, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}
		// TODO: judge app id
		// exist, err := appcli.ExistApp(ctx, *appID)
		// if err != nil {
		// 	return err
		// }
		// if !exist {
		// 	return fmt.Errorf("invalid app")
		// }

		h.AppID = &_appID
		return nil
	}
}

func WithLangID(langID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_langID, err := uuid.Parse(*langID)
		if err != nil {
			return err
		}
		// TODO: judge lang id
		// exist, err := appcli.ExistApp(ctx, *appID)
		// if err != nil {
		// 	return err
		// }
		// if !exist {
		// 	return fmt.Errorf("invalid app")
		// }

		h.LangID = &_langID
		return nil
	}
}

func WithTitle(title *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if title == nil {
			return nil
		}
		const leastTitleLen = 4
		if len(*title) < leastTitleLen {
			return fmt.Errorf("name %v too short", *title)
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
		const leastContentLen = 4
		if len(*content) < leastContentLen {
			return fmt.Errorf("content %v too short", *content)
		}
		h.Content = content
		return nil
	}
}

func WithChannel(channel *basetypes.NotifChannel) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if channel == nil {
			return nil
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

func WithAnnouncementType(_type *npool.AnnouncementType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case npool.AnnouncementType_Broadcast:
		case npool.AnnouncementType_Multicast:
		default:
			return fmt.Errorf("type %v invalid", *_type)
		}
		h.Type = _type
		return nil
	}
}

func WithEndAt(endAt int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if endAt < int32(time.Now().Unix()) {
			return fmt.Errorf("invalid end at")
		}
		h.EndAt = endAt
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
		h.Conds = &amtcrud.Conds{}
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
