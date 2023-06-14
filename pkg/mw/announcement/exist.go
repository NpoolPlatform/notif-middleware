package announcement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
)

func (h *Handler) ExistAnnouncement(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Announcement.
			Query().
			Where(
				entamt.ID(*h.ID),
				entamt.AppID(*h.AppID),
				entamt.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
