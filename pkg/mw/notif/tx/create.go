package tx

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/tx"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.NotifState == nil {
		return fmt.Errorf("notif state empty")
	}
	return nil
}

func (h *Handler) CreateTx(ctx context.Context) (info *npool.Tx, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.TxNotifState.Create(),
			&crud.Req{
				ID:         h.ID,
				TxID:       h.TxID,
				NotifState: h.NotifState,
				TxType:     h.TxType,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}
