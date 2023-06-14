package readstate

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/readstate"
	entreadstate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readnotif"

	readstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/readstate"
)

type queryHandler struct {
	*Handler
	stm   *ent.ReadNotifSelect
	infos []*npool.ReadState
	total uint32
}

func (h *queryHandler) selectReadState(stm *ent.ReadNotifQuery) {
	h.stm = stm.Select(
		entreadstate.FieldID,
		entreadstate.FieldAppID,
		entreadstate.FieldUserID,
		entreadstate.FieldNotifID,
		entreadstate.FieldCreatedAt,
		entreadstate.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryReadState(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid readnotifid")
	}

	h.selectReadState(
		cli.ReadNotif.
			Query().
			Where(
				entreadstate.ID(*h.ID),
				entreadstate.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryReadStates(ctx context.Context, cli *ent.Client) error {
	stm, err := readstatecrud.SetQueryConds(cli.ReadNotif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectReadState(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetReadState(ctx context.Context) (*npool.ReadState, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadState(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetReadStates(ctx context.Context) ([]*npool.ReadState, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadStates(ctx, cli); err != nil {
			return err
		}
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetReadStateOnly(ctx context.Context) (info *npool.ReadState, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryReadStates(_ctx, cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
