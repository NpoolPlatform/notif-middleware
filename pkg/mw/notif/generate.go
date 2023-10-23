package notif

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	usernotifcrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/user"
	usernotifmw "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/user"
	tmplmw "github.com/NpoolPlatform/notif-middleware/pkg/mw/template"
	"github.com/google/uuid"
)

type generateHandler struct {
	*Handler
}

func (h *generateHandler) getNotifUsers(ctx context.Context) ([]string, error) {
	userIDs := []string{}
	const maxLimit = int32(100)
	switch *h.NotifType {
	case basetypes.NotifType_NotifMulticast:
		usernotifHandler, err := usernotifmw.NewHandler(
			ctx,
			usernotifmw.WithOffset(0),
			usernotifmw.WithLimit(maxLimit),
		)
		if err != nil {
			return nil, err
		}
		usernotifHandler.Conds = &usernotifcrud.Conds{
			AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
		}
		userNotifs, _, err := usernotifHandler.GetNotifUsers(ctx)
		if err != nil {
			return nil, err
		}
		if len(userNotifs) == 0 {
			return nil, fmt.Errorf("invalid user notif")
		}
		for _, row := range userNotifs {
			userIDs = append(userIDs, row.UserID)
		}
	case basetypes.NotifType_NotifUnicast:
		if h.UserID == nil {
			return nil, fmt.Errorf("invalid userid")
		}
		userID := h.UserID.String()
		userIDs = append(userIDs, userID)
	default:
		return nil, fmt.Errorf("invalid notiftype")
	}
	return userIDs, nil
}

func (h *generateHandler) createUserNotifs(ctx context.Context, appID, eventID, userID string) ([]*npool.NotifReq, error) {
	reqs := []*npool.NotifReq{}
	templateHandler, err := tmplmw.NewHandler(
		ctx,
		tmplmw.WithAppID(&appID, true),
		tmplmw.WithUserID(&userID, true),
		tmplmw.WithUsedFor(h.EventType, true),
		tmplmw.WithVars(h.Vars, true),
	)
	if err != nil {
		return nil, err
	}
	_reqs, err := templateHandler.GenerateNotifs(ctx)
	if err != nil {
		return nil, err
	}
	for _, req := range _reqs {
		req.Extra = h.Extra
		req.NotifType = h.NotifType
		req.EventID = &eventID
	}
	reqs = append(reqs, _reqs...)
	return reqs, nil
}

func (h *Handler) GenerateNotifs(
	ctx context.Context,
) (
	[]*npool.Notif, error,
) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.EventType == nil {
		return nil, fmt.Errorf("invalid eventtype")
	}
	if h.NotifType == nil {
		return nil, fmt.Errorf("invalid notiftype")
	}
	appID := h.AppID.String()
	eventID := uuid.NewString()

	handler := &generateHandler{
		Handler: h,
	}
	userIDs, err := handler.getNotifUsers(ctx)
	if err != nil {
		return nil, err
	}

	reqs := []*npool.NotifReq{}
	for _, _userID := range userIDs {
		_reqs, err := handler.createUserNotifs(ctx, appID, eventID, _userID)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, _reqs...)
	}

	notifGenerateHandler, err := NewHandler(
		ctx,
		WithReqs(reqs, true),
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
