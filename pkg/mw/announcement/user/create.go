package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/user"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/user"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
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

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
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

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAnnouncementUser(ctx)
}

func (h *Handler) CreateAnnouncementUsers(ctx context.Context) (infos []*npool.AnnouncementUser, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			h.AppID = req.AppID
			h.UserID = req.UserID
			h.AnnouncementID = req.AnnouncementID

			h.Conds = &crud.Conds{
				AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
				AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
			}

			exist, err := h.ExistAnnouncementUserConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				continue
			}

			info, err := h.CreateAnnouncementUser(ctx)
			if err != nil {
				return err
			}
			infos = append(infos, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
