package notif

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	tmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
)

func GenerateNotifs(
	ctx context.Context,
	appID string,
	usedFor basetypes.UsedFor,
	vars *tmplmwpb.TemplateVars,
) (
	[]*npool.Notif, error,
) {
	return nil, nil
}
