package readstate

import (
	"context"
	"fmt"

	readstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/readstate"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entreadnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readnotif"
)

func (h *Handler) ExistReadState(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			ReadNotif.
			Query().
			Where(
				entreadnotif.ID(*h.ID),
				entreadnotif.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistReadStateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := readstatecrud.SetQueryConds(cli.ReadNotif.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
