//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notif"
	"github.com/google/uuid"
)

func (h *Handler) UpdateNotif(ctx context.Context) (*npool.Notif, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.Notified == nil {
		return nil, fmt.Errorf("invalid notified")
	}

	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixCreateNotif,
		*h.ID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	info, err := h.GetNotif(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("notif not exist")
	}
	if info.Notified {
		if *h.Notified != info.Notified {
			return nil, fmt.Errorf("invalid notified")
		}
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := notifcrud.UpdateSet(
			tx.Notif.UpdateOneID(*h.ID),
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
	})
	if err != nil {
		return nil, err
	}

	return h.GetNotif(ctx)
}

// nolint:gocyclo
func (h *Handler) UpdateNotifs(ctx context.Context) ([]*npool.Notif, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				return fmt.Errorf("invalid id")
			}
			if req.Notified == nil {
				return fmt.Errorf("invalid notified")
			}

			h.ID = req.ID
			info, err := h.GetNotif(ctx)
			if err != nil {
				return err
			}
			if info == nil {
				return fmt.Errorf("notif not exist")
			}
			if info.Notified {
				if *req.Notified != info.Notified {
					return fmt.Errorf("invalid notified")
				}
			}

			eventID, err := uuid.Parse(info.EventID)
			if err != nil {
				return fmt.Errorf("invalid event id")
			}

			switch info.Channel {
			case basetypes.NotifChannel_ChannelFrontend:
				if _, err = tx.Notif.
					Update().
					Where(
						entnotif.EventID(eventID),
						entnotif.Channel(basetypes.NotifChannel_ChannelFrontend.String()),
					).
					SetNotified(true).
					Save(_ctx); err != nil {
					return err
				}
			default:
				if _, err := notifcrud.UpdateSet(
					tx.Notif.UpdateOneID(*req.ID),
					&notifcrud.Req{
						Notified: req.Notified,
					},
				).Save(ctx); err != nil {
					return err
				}
			}
			ids = append(ids, *req.ID)
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
