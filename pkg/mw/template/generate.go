package template

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	notifmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"
	notifchanmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/channel"

	notifchanmw "github.com/NpoolPlatform/notif-middleware/pkg/mw/notif/channel"

	email "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/email"
	frontend "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/frontend"
	sms "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/sms"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) ([]*notifmwpb.NotifReq, error) {
	const maxChannels = int32(100)
	chanHandler, err := notifchanmw.NewHandler(
		ctx,
		notifchanmw.WithConds(&notifchanmwpb.Conds{
			AppID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			EventType: &basetypes.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(usedFor),
			},
		}),
		notifchanmw.WithOffset(0),
		notifchanmw.WithLimit(maxChannels),
	)
	if err != nil {
		return nil, err
	}

	chans, _, err := chanHandler.GetChannels(ctx)
	if err != nil {
		return nil, err
	}
	if len(chans) == 0 {
		return nil, err
	}

	reqs := []*notifmwpb.NotifReq{}

	for _, ch := range chans {
		switch ch.Channel {
		case basetypes.NotifChannel_ChannelEmail:
			emailHandler, err := email.NewHandler(
				ctx,
				email.WithAppID(&appID),
				email.WithUserID(&userID),
				email.WithUsedFor(&usedFor),
				email.WithVars(vars),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := emailHandler.GenerateNotifs(ctx)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case basetypes.NotifChannel_ChannelSMS:
			smsHandler, err := sms.NewHandler(
				ctx,
				sms.WithAppID(&appID),
				sms.WithUserID(&userID),
				sms.WithUsedFor(&usedFor),
				sms.WithVars(vars),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := smsHandler.GenerateNotifs(ctx)
			if err != nil {
				return nil, err
			}
			reqs = append(reqs, _reqs...)
		case basetypes.NotifChannel_ChannelFrontend:
			frontendHandler, err := frontend.NewHandler(
				ctx,
				frontend.WithAppID(&appID),
				frontend.WithUserID(&userID),
				frontend.WithUsedFor(&usedFor),
				frontend.WithVars(vars),
			)
			if err != nil {
				return nil, err
			}
			_reqs, err := frontendHandler.GenerateNotifs(ctx)
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
	channel basetypes.NotifChannel,
	vars *npool.TemplateVars,
) (*npool.TextInfo, error) {
	switch channel {
	case basetypes.NotifChannel_ChannelEmail:
		emailHandler, err := email.NewHandler(
			ctx,
			email.WithAppID(&appID),
			email.WithLangID(&langID),
			email.WithUsedFor(&usedFor),
			email.WithVars(vars),
		)
		if err != nil {
			return nil, err
		}
		return emailHandler.GenerateText(ctx)
	case basetypes.NotifChannel_ChannelSMS:
		smsHandler, err := sms.NewHandler(
			ctx,
			sms.WithAppID(&appID),
			sms.WithLangID(&langID),
			sms.WithUsedFor(&usedFor),
			sms.WithVars(vars),
		)
		if err != nil {
			return nil, err
		}
		return smsHandler.GenerateText(ctx)
	case basetypes.NotifChannel_ChannelFrontend:
		frontendHandler, err := frontend.NewHandler(
			ctx,
			frontend.WithAppID(&appID),
			frontend.WithLangID(&langID),
			frontend.WithUsedFor(&usedFor),
			frontend.WithVars(vars),
		)
		if err != nil {
			return nil, err
		}
		return frontendHandler.GenerateText(ctx)
	}

	return nil, fmt.Errorf("unknown channel")
}
