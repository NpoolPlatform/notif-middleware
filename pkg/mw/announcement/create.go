package announcement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Title == nil {
		return fmt.Errorf("title empty")
	}
	if h.Content == nil {
		return fmt.Errorf("content empty")
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

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.Announcement.Create(),
			&crud.Req{
				ID:      h.ID,
				AppID:   h.AppID,
				Title:   h.Title,
				Content: h.Content,
				LangID:  h.LangID,
				Channel: h.Channel,
				Type:    h.Type,
				EndAt:   &h.EndAt,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncement(ctx)
}
