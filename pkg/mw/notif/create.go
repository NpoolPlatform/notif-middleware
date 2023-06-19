//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createNotif(ctx context.Context, cli *ent.Client) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid lang")
	}
	if h.LangID == nil {
		return fmt.Errorf("invalid logo")
	}
	if h.EventID == nil {
		return fmt.Errorf("invalid eventid")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*h.AppID,
		*h.LangID,
		*h.EventID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := notifcrud.CreateSet(
		cli.Notif.Create(),
		&notifcrud.Req{
			ID:          h.ID,
			AppID:       h.AppID,
			LangID:      h.LangID,
			UserID:      h.UserID,
			EventID:     h.EventID,
			Notified:    h.Notified,
			EventType:   h.EventType,
			UseTemplate: h.UseTemplate,
			Title:       h.Title,
			Content:     h.Content,
			Channel:     h.Channel,
			Extra:       h.Extra,
			NotifType:   h.NotifType,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateNotif(ctx context.Context) (*npool.Notif, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createNotif(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotif(ctx)
}

func (h *Handler) CreateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.AppID = req.AppID
			handler.LangID = req.LangID
			handler.UserID = req.UserID
			handler.EventID = req.EventID
			handler.Notified = req.Notified
			handler.EventType = req.EventType
			handler.UseTemplate = req.UseTemplate
			handler.Title = req.Title
			handler.Content = req.Content
			handler.Channel = req.Channel
			handler.Extra = req.Extra
			handler.NotifType = req.NotifType
			if err := handler.createNotif(ctx, cli); err != nil {
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
