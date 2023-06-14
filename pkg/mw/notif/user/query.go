package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	entuser "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"

	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"
)

type queryHandler struct {
	*Handler
	stm   *ent.UserNotifSelect
	infos []*npool.UserNotif
	total uint32
}

func (h *queryHandler) selectUser(stm *ent.UserNotifQuery) {
	h.stm = stm.Select(
		entuser.FieldID,
		entuser.FieldAppID,
		entuser.FieldUserID,
		entuser.FieldNotifID,
		entuser.FieldCreatedAt,
		entuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid usernotifid")
	}

	h.selectUser(
		cli.UserNotif.
			Query().
			Where(
				entuser.ID(*h.ID),
				entuser.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.UserNotif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetUser(ctx context.Context) (*npool.UserNotif, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUser(cli); err != nil {
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

func (h *Handler) GetUsers(ctx context.Context) ([]*npool.UserNotif, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUsers(ctx, cli); err != nil {
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

func (h *Handler) GetUserOnly(ctx context.Context) (info *npool.UserNotif, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUsers(_ctx, cli); err != nil {
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
