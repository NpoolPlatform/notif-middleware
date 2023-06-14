package sendstate

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	amt "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateSendState(ctx context.Context) (info *npool.SendState, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	// get announcement first to get channel
	amtID := handler.AnnouncementID.String()
	amtHandler, err := amt.NewHandler(ctx, amt.WithID(&amtID))
	if err != nil {
		return nil, err
	}

	announcement, err := amtHandler.GetAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if announcement == nil {
		return nil, fmt.Errorf("invalid announcement id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.SendAnnouncement.Create(),
			&crud.Req{
				ID:             h.ID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				AnnouncementID: h.AnnouncementID,
				Channel:        &announcement.Channel,
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

	return h.GetSendState(ctx)
}
