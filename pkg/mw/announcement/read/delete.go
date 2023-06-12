package read

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/read"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/read"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) DeleteReadAnnouncement(ctx context.Context) (*npool.ReadAnnouncement, error) {
	info, err := h.GetReadAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.ReadAnnouncement.UpdateOneID(*h.ID),
			&crud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
