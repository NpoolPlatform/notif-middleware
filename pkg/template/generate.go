package template

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	chanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	notifmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	notifchanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	notifchanmgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/notif/channel"

	email "github.com/NpoolPlatform/notif-middleware/pkg/template/email"
	frontend "github.com/NpoolPlatform/notif-middleware/pkg/template/frontend"
	sms "github.com/NpoolPlatform/notif-middleware/pkg/template/sms"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) ([]*notifmgrpb.NotifReq, error) {
	const maxChannels = int32(100)

	chans, _, err := notifchanmgrcli.GetChannels(ctx, &notifchanmgrpb.Conds{
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
		case chanmgrpb.NotifChannel_ChannelEmail:
			_reqs, err := email.GenerateNotifs(ctx, appID, userID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case chanmgrpb.NotifChannel_ChannelSMS:
			_reqs, err := sms.GenerateNotifs(ctx, appID, userID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case chanmgrpb.NotifChannel_ChannelFrontend:
			_reqs, err := frontend.GenerateNotifs(ctx, appID, userID, usedFor, vars)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		}
	}

	return reqs, nil
}

func GenerateText(
	ctx context.Context,
	appID, langID string,
	usedFor basetypes.UsedFor,
	channel chanmgrpb.NotifChannel,
	vars *npool.TemplateVars,
) (*npool.TextInfo, error) {
	switch channel {
	case chanmgrpb.NotifChannel_ChannelEmail:
		return email.GenerateText(ctx, appID, langID, usedFor, vars)
	case chanmgrpb.NotifChannel_ChannelSMS:
		return sms.GenerateText(ctx, appID, langID, usedFor, vars)
	case chanmgrpb.NotifChannel_ChannelFrontend:
		return frontend.GenerateText(ctx, appID, langID, usedFor, vars)
	}

	return nil, fmt.Errorf("unknown channel")
}
