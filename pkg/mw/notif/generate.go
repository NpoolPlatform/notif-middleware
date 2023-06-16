package notif

import (
	"context"

	mwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/notif"
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

	notifs, err := mwcli.CreateNotifs(ctx, reqs)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}
