package readstate

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/readstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) DeleteReadState(ctx context.Context) (*npool.ReadState, error) {
	info, err := h.GetReadState(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil { // dtm required
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.ReadState.UpdateOneID(*h.ID),
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
