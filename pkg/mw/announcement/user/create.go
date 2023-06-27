package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	announcement1 "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/google/uuid"
)

func (h *Handler) CreateAnnouncementUser(ctx context.Context) (info *npool.AnnouncementUser, err error) {
	h.Conds = &crud.Conds{
		AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
	}
	exist, err := h.ExistAnnouncementUserConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("announcement user exist")
	}

	// only announcement type is Multicast can create user
	announcementID := h.AnnouncementID.String()
	announcementHandler, err := announcement1.NewHandler(ctx, announcement1.WithID(&announcementID))
	if err != nil {
		return nil, err
	}
	announcement, err := announcementHandler.GetAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if announcement.AnnouncementType != basetypes.NotifType_NotifMulticast {
		return nil, fmt.Errorf("wrong announcement type %v", announcement.AnnouncementType.String())
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := crud.CreateSet(
			cli.UserAnnouncement.Create(),
			&crud.Req{
				ID:             h.ID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				AnnouncementID: h.AnnouncementID,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncementUser(ctx)
}
