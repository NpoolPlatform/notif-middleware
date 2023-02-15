package template

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	chanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	notifmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	notifchanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	notifchanmgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/channel"

	email "github.com/NpoolPlatform/notif-middleware/pkg/template/email"
)

func FillTemplate(
	ctx context.Context,
	appID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) (
	[]*notifmgrpb.NotifReq, error,
) {
	const maxChannels = int32(100)

	chans, err := notifchanmgrcli.GetChannels(ctx, &notifchanmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		EventType: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(usedFor),
		},
	}, 0, maxChannels)
	if err != nil {
		return nil, err
	}
	if len(chans) == 0 {
		return nil, err
	}

	reqs := []*notifmgrpb.NotifReq{}

	for _, ch := range chans {
		switch ch.Channel {
		case chanmgrpb.NotifChannel_Email:
			_reqs, err := email.FillTemplate(ctx, appID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs)
		case chanmgrpb.NotifChannel_SMS:
			_reqs, err := sms.FillTemplate(ctx, appID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs)
		case chanmgrpb.NotifChannel_Frontend:
			_reqs, err := frontend.FillTemplate(ctx, appID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs)
		}
	}

	return reqs, nil
}
