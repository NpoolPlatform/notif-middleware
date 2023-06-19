package sendstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/sendstate"
	entsendstate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendnotif"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	sendstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/sendstate"
)

type queryHandler struct {
	*Handler
	stm   *ent.SendNotifSelect
	infos []*npool.SendState
	total uint32
}

func (h *queryHandler) selectSendState(stm *ent.SendNotifQuery) {
	h.stm = stm.Select(
		entsendstate.FieldID,
		entsendstate.FieldAppID,
		entsendstate.FieldUserID,
		entsendstate.FieldNotifID,
		entsendstate.FieldChannel,
		entsendstate.FieldCreatedAt,
		entsendstate.FieldUpdatedAt,
	)
}

func (h *queryHandler) querySendState(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid sendnotifid")
	}

	h.selectSendState(
		cli.SendNotif.
			Query().
			Where(
				entsendstate.ID(*h.ID),
				entsendstate.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Channel = basetypes.NotifChannel(basetypes.NotifChannel_value[info.ChannelStr])
	}
}

func (h *queryHandler) querySendStates(ctx context.Context, cli *ent.Client) error {
	stm, err := sendstatecrud.SetQueryConds(cli.SendNotif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectSendState(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSendState(ctx context.Context) (*npool.SendState, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendState(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		if err := handler.scan(_ctx); err != nil {
			return err
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
	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetSendStates(ctx context.Context) ([]*npool.SendState, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendStates(ctx, cli); err != nil {
			return err
		}
		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetSendStateOnly(ctx context.Context) (info *npool.SendState, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySendStates(_ctx, cli); err != nil {
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
