package tx

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/tx"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/tx"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	enttx "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/txnotifstate"
)

type queryHandler struct {
	*Handler
	stm   *ent.TxNotifStateSelect
	infos []*npool.Tx
	total uint32
}

func (h *queryHandler) selectTx(stm *ent.TxNotifStateQuery) {
	h.stm = stm.Select(
		enttx.FieldID,
		enttx.FieldTxID,
		enttx.FieldNotifState,
		enttx.FieldTxType,
		enttx.FieldCreatedAt,
		enttx.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTx(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid notif tx id")
	}
	h.selectTx(
		cli.TxNotifState.
			Query().
			Where(
				enttx.ID(*h.ID),
				enttx.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryTxsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.TxNotifState.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectTx(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTxs(ctx context.Context) ([]*npool.Tx, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTxsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetTx(ctx context.Context) (info *npool.Tx, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTx(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetTxOnly(ctx context.Context) (*npool.Tx, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTxsByConds(_ctx, cli); err != nil {
			return err
		}

		_, err := handler.stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 0 {
		return nil, fmt.Errorf("to many record")
	}

	return handler.infos[0], nil
}
