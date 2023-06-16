package sendstate

import (
	"context"
	"fmt"

	sendstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/sendstate"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entsendnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendnotif"
)

func (h *Handler) ExistSendState(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			SendNotif.
			Query().
			Where(
				entsendnotif.ID(*h.ID),
				entsendnotif.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistSendStateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := sendstatecrud.SetQueryConds(cli.SendNotif.Query(), h.Conds)
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
