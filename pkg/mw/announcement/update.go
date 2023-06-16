package announcement

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) UpdateAnnouncement(ctx context.Context) (info *npool.Announcement, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Announcement.UpdateOneID(*h.ID),
			&crud.Req{
				Title:   h.Title,
				Content: h.Content,
				EndAt:   h.EndAt,
				Type:    h.Type,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncement(ctx)
}
