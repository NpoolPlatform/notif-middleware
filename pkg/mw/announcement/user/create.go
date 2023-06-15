package user

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateAnnouncementUser(ctx context.Context) (info *npool.AnnouncementUser, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	// TODO:Check Exist
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.UserAnnouncement.Create(),
			&crud.Req{
				ID:             h.ID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				AnnouncementID: h.AnnouncementID,
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

	return h.GetAnnouncementUser(ctx)
}
