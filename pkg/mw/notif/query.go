package notif

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	entnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notif"

	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"
)

type queryHandler struct {
	*Handler
	stm   *ent.NotifSelect
	infos []*npool.Notif
	total uint32
}

func (h *queryHandler) selectNotif(stm *ent.NotifQuery) {
	h.stm = stm.Select(
		entnotif.FieldID,
		entnotif.FieldAppID,
		entnotif.FieldLangID,
		entnotif.FieldUserID,
		entnotif.FieldEventID,
		entnotif.FieldNotified,
		entnotif.FieldEventType,
		entnotif.FieldUseTemplate,
		entnotif.FieldTitle,
		entnotif.FieldContent,
		entnotif.FieldChannel,
		entnotif.FieldExtra,
		entnotif.FieldCreatedAt,
		entnotif.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryNotif(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid notifid")
	}

	h.selectNotif(
		cli.Notif.
			Query().
			Where(
				entnotif.ID(*h.ID),
				entnotif.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryNotifs(ctx context.Context, cli *ent.Client) error {
	stm, err := notifcrud.SetQueryConds(cli.Notif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectNotif(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetNotif(ctx context.Context) (*npool.Notif, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotif(cli); err != nil {
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

func (h *Handler) GetNotifs(ctx context.Context) ([]*npool.Notif, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(ctx, cli); err != nil {
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

func (h *Handler) GetNotifOnly(ctx context.Context) (info *npool.Notif, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifs(_ctx, cli); err != nil {
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
