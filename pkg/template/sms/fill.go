package sms

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	notifmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
)

func FillTemplate(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) (
	[]*notifmgrpb.NotifReq, error,
) {
	return nil, nil
}
