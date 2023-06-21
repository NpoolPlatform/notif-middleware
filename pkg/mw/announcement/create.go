package announcement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("app is empty")
	}
	if h.Title == nil {
		return fmt.Errorf("title is empty")
	}
	if h.Content == nil {
		return fmt.Errorf("content is empty")
	}
	if h.Type == nil {
		return fmt.Errorf("type is empty")
	}
	if h.Channel == nil {
		return fmt.Errorf("channel is empty")
	}
	if h.StartAt == nil {
		return fmt.Errorf("start at is empty")
	}
	if h.EndAt == nil {
		return fmt.Errorf("end at is empty")
	}
	return nil
}

func (h *Handler) CreateAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.Announcement.Create(),
			&crud.Req{
				ID:      h.ID,
				AppID:   h.AppID,
				Title:   h.Title,
				Content: h.Content,
				LangID:  h.LangID,
				Channel: h.Channel,
				Type:    h.Type,
				StartAt: h.StartAt,
				EndAt:   h.EndAt,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncement(ctx)
}
