package channel

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/channel"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateChannel(ctx context.Context) (info *npool.Channel, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.NotifChannel.Create(),
			&crud.Req{
				ID:        h.ID,
				AppID:     h.AppID,
				Channel:   h.Channel,
				EventType: h.EventType,
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

	return h.GetChannel(ctx)
}
