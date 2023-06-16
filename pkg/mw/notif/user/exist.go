package user

import (
	"context"
	"fmt"

	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entusernotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"
)

func (h *Handler) ExistUser(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			UserNotif.
			Query().
			Where(
				entusernotif.ID(*h.ID),
				entusernotif.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistUserConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(cli.UserNotif.Query(), h.Conds)
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
