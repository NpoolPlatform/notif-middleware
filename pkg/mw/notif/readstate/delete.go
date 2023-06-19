package readstate

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	readstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/readstate"
)

func (h *Handler) DeleteReadState(ctx context.Context) (*npool.ReadState, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetReadState(ctx)
	if err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := readstatecrud.UpdateSet(
			cli.ReadNotif.UpdateOneID(*h.ID),
			&readstatecrud.Req{
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
