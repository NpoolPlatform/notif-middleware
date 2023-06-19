package readstate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement/readstate"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/announcement/readstate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateReadState(ctx context.Context) (info *npool.ReadState, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	h.Conds = &crud.Conds{
		AppID: &cruder.Cond{
			Op:  cruder.EQ,
			Val: *h.AppID,
		},
		UserID: &cruder.Cond{
			Op:  cruder.EQ,
			Val: *h.UserID,
		},
		AnnouncementID: &cruder.Cond{
			Op:  cruder.EQ,
			Val: *h.AnnouncementID,
		},
	}
	h.Offset = 0
	h.Limit = 1

	infos, _, err := h.GetReadStates(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) > 0 {
		return nil, fmt.Errorf("read state exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.ReadAnnouncement.Create(),
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

	return h.GetReadState(ctx)
}
