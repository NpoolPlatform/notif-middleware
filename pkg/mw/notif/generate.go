package notif

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	templatemw "github.com/NpoolPlatform/notif-middleware/pkg/mw/template"
)

func (h *Handler) GenerateNotifs(
	ctx context.Context,
) (
	[]*npool.Notif, error,
) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}
	if h.EventType == nil {
		return nil, fmt.Errorf("invalid eventtype")
	}
	appID := h.AppID.String()
	userID := h.UserID.String()

	templateHandler, err := templatemw.NewHandler(
		ctx,
		templatemw.WithAppID(&appID),
		templatemw.WithUserID(&userID),
		templatemw.WithUsedFor(h.EventType),
		templatemw.WithVars(h.Vars),
	)
	if err != nil {
		return nil, err
	}

	reqs, err := templateHandler.GenerateNotifs(ctx)
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
