package notif

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	tmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	template "github.com/NpoolPlatform/notif-middleware/pkg/template"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *tmplmwpb.TemplateVars,
) (
	[]*npool.Notif, error,
) {
	reqs, err := template.FillTemplate(ctx, appID, userID, usedFor, vars)
	if err != nil {
		return nil, err
	}

	notifs, err := mgrcli.CreateNotifs(ctx, reqs)
	if err != nil {
		return nil, err
	}

	return expandMany(notifs), nil
}
