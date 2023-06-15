package notif

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateNotif(ctx context.Context, cli *ent.Client) error {
	if _, err := notifcrud.UpdateSet(
		cli.Notif.UpdateOneID(*h.ID),
		&notifcrud.Req{
			Notified:    h.Notified,
			UseTemplate: h.UseTemplate,
			Title:       h.Title,
			Content:     h.Content,
			Channel:     h.Channel,
			Extra:       h.Extra,
			NotifType:   h.NotifType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateNotif(ctx context.Context) (*npool.Notif, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.LangID == nil {
		return nil, fmt.Errorf("invalid langid")
	}

	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixSetFiat,
		*h.ID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &notifcrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
	}
	h.Offset = 0
	h.Limit = 2

	notif, err := h.GetNotifOnly(ctx)
	if err != nil {
		return nil, err
	}
	if notif != nil {
		return nil, fmt.Errorf("notif exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateNotif(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotif(ctx)
}

func (h *Handler) UpdateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	handler := &updateHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.AppID = req.AppID
			handler.LangID = req.LangID
			if err := handler.updateNotif(_ctx, cli); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &notifcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetNotifs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
