package notif

import (
	"context"

	"github.com/NpoolPlatform/notif-middleware/pkg/mw/template"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
)

func (h *Handler) GenerateNotifs(
	ctx context.Context,
) (
	[]*npool.Notif, error,
) {
	reqs, err := template.GenerateNotifs(ctx, h.AppID.String(), h.UserID.String(), *h.EventType, h.Vars)
	if err != nil {
		return nil, err
	}
	for _, req := range reqs {
		req.Extra = h.Extra
	}

	notifGenerateHandler, err := NewHandler(
		ctx,
		WithReqs(reqs),
	)
	if err != nil {
		return nil, err
	}

	notifs, err := notifGenerateHandler.CreateNotifs(ctx)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}
