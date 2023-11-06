package sendstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/sendstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	amt "github.com/NpoolPlatform/notif-middleware/pkg/mw/announcement"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateSendState(ctx context.Context) (info *npool.SendState, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	// get announcement first to get channel attr
	amtID := handler.AnnouncementID.String()
	amtHandler, err := amt.NewHandler(ctx, amt.WithEntID(&amtID, true))
	if err != nil {
		return nil, err
	}

	announcement, err := amtHandler.GetAnnouncement(ctx)
	if err != nil {
		return nil, err
	}
	if announcement == nil {
		return nil, fmt.Errorf("invalid announcement id")
	}

	h.Conds = &crud.Conds{
		AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
	}

	exist, err := h.ExistSendStateConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("send state exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.SendAnnouncement.Create(),
			&crud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				AnnouncementID: h.AnnouncementID,
				Channel:        &announcement.Channel,
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

	return h.GetSendState(ctx)
}

func (h *Handler) CreateSendStates(ctx context.Context) (infos []*npool.SendState, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			h.AppID = req.AppID
			h.UserID = req.UserID
			h.AnnouncementID = req.AnnouncementID
			h.Channel = req.Channel

			h.Conds = &crud.Conds{
				AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID:         &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
				AnnouncementID: &cruder.Cond{Op: cruder.EQ, Val: *h.AnnouncementID},
			}

			exist, err := h.ExistSendStateConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				continue
			}

			info, err := h.CreateSendState(ctx)
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
