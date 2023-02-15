package sms

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	chanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	notifmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	smstmplmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/sms"
	smstmplmgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/template/sms"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	tmplreplace "github.com/NpoolPlatform/notif-middleware/pkg/template/replace"
)

func FillTemplate(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) (
	[]*notifmgrpb.NotifReq, error,
) {
	const maxTemplates = int32(100)

	tmpls, _, err := smstmplmgrcli.GetSMSTemplates(ctx, &smstmplmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UsedFor: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(usedFor),
		},
	}, 0, maxTemplates)
	if err != nil {
		return nil, err
	}
	if len(tmpls) == 0 {
		return nil, nil
	}

	reqs := []*notifmgrpb.NotifReq{}
	for _, tmpl := range tmpls {
		title := tmplreplace.ReplaceAll(tmpl.Subject, vars)
		content := tmplreplace.ReplaceAll(tmpl.Message, vars)
		useTemplate := true
		channel1 := chanmgrpb.NotifChannel_ChannelSMS

		reqs = append(reqs, &notifmgrpb.NotifReq{
			AppID:       &appID,
			UserID:      &userID,
			LangID:      &tmpl.LangID,
			EventType:   &usedFor,
			UseTemplate: &useTemplate,
			Title:       &title,
			Content:     &content,
			Channel:     &channel1,
		})
	}

	return reqs, nil
}
