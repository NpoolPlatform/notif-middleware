//nolint:nolintlint,dupl
package notif

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif"
	readstatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/readstate"
	readstatemw "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/readstate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

func (h *Handler) UpdateNotif(ctx context.Context) (*npool.Notif, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
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

	info, err := h.GetNotif(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("notif not exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		switch info.NotifType {
		case basetypes.NotifType_NotifMulticast:
			break
		case basetypes.NotifType_NotifUnicast:
			break
		}
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
			h.ID = req.ID
			info, err := h.GetNotif(ctx)
			if err != nil {
				return err
			}
			if info == nil {
				return fmt.Errorf("notif not exist")
			}

			switch info.NotifType {
			case basetypes.NotifType_NotifMulticast:
				if req.UserID == nil {
					return fmt.Errorf("invalid userid")
				}

				readstateHandler, err := readstatemw.NewHandler(
					ctx,
					readstatemw.WithAppID(&info.AppID),
				)
				if err != nil {
					return err
				}
				readstateHandler.Conds = &readstatecrud.Conds{
					AppID:   &cruder.Cond{Op: cruder.EQ, Val: *readstateHandler.AppID},
					UserID:  &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
					NotifID: &cruder.Cond{Op: cruder.EQ, Val: *req.ID},
				}
				exist, err := readstateHandler.ExistReadStateConds(ctx)
				if err != nil {
					return err
				}
				if !exist {
					readstateHandler.UserID = req.UserID
					readstateHandler.NotifID = req.ID
					_, err := readstateHandler.CreateReadState(ctx)
					if err != nil {
						return err
					}
				}
			case basetypes.NotifType_NotifUnicast:
				if info.Notified {
					if *h.Notified != info.Notified {
						return fmt.Errorf("invalid notified")
					}
				}
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
