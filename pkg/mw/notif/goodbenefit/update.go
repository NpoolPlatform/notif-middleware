package goodbenefit

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) UpdateGoodBenefit(ctx context.Context) (info *npool.GoodBenefit, err error) {
	if h.ID == nil {
		return nil, fmt.Errorf("id is empty")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.GoodBenefit.UpdateOneID(*h.ID),
			&crud.Req{
				ID:       h.ID,
				Notified: h.Notified,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodBenefit(ctx)
}
