package notif

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	tmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	mwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/notif"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	template "github.com/NpoolPlatform/notif-middleware/pkg/mw/template"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *tmplmwpb.TemplateVars,
	extra *string,
) (
	[]*npool.Notif, error,
) {
	reqs, err := template.GenerateNotifs(ctx, appID, userID, usedFor, vars)
	if err != nil {
		return nil, err
	}

	for _, req := range reqs {
		req.Extra = extra
	}

	notifs, err := mwcli.CreateNotifs(ctx, reqs)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}
