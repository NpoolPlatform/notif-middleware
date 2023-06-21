package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/user"
	entuser "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"

	usercrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"
)

type queryHandler struct {
	*Handler
	stm   *ent.UserNotifSelect
	infos []*npool.NotifUser
	total uint32
}

func (h *queryHandler) selectNotifUser(stm *ent.UserNotifQuery) {
	h.stm = stm.Select(
		entuser.FieldID,
		entuser.FieldAppID,
		entuser.FieldUserID,
		entuser.FieldEventType,
		entuser.FieldCreatedAt,
		entuser.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryNotifUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid usernotifid")
	}

	h.selectNotifUser(
		cli.UserNotif.
			Query().
			Where(
				entuser.ID(*h.ID),
				entuser.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryNotifUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := usercrud.SetQueryConds(cli.UserNotif.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectNotifUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetNotifUser(ctx context.Context) (*npool.NotifUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUser(cli); err != nil {
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

	return handler.infos[0], nil
}

func (h *Handler) GetNotifUsers(ctx context.Context) ([]*npool.NotifUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUsers(ctx, cli); err != nil {
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

	return handler.infos, handler.total, nil
}

func (h *Handler) GetNotifUserOnly(ctx context.Context) (info *npool.NotifUser, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryNotifUsers(_ctx, cli); err != nil {
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
