package tx

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/tx"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) UpdateTx(ctx context.Context) (info *npool.Tx, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.TxNotifState.UpdateOneID(*h.ID),
			&crud.Req{
				ID:         h.ID,
				NotifState: h.NotifState,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}
