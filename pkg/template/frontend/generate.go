package frontend

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	chanmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	notifmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	frontendtmplmgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"
	frontendtmplmgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/template/frontend"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	tmplreplace "github.com/NpoolPlatform/notif-middleware/pkg/template/replace"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) ([]*notifmgrpb.NotifReq, error) {
	const maxTemplates = int32(100)

	tmpls, _, err := frontendtmplmgrcli.GetFrontendTemplates(ctx, &frontendtmplmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UsedFor: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(usedFor),
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
		title := tmplreplace.ReplaceAll(tmpl.Title, vars)
		content := tmplreplace.ReplaceAll(tmpl.Content, vars)
		useTemplate := true
		channel1 := chanmgrpb.NotifChannel_ChannelFrontend

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

func GenerateText(
	ctx context.Context,
	appID, langID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) (*npool.TextInfo, error) {
	tmpl, err := frontendtmplmgrcli.GetFrontendTemplateOnly(ctx, &frontendtmplmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: langID,
		},
		UsedFor: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(usedFor),
		},
	})
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, nil
	}

	title := tmplreplace.ReplaceAll(tmpl.Content, vars)
	content := tmplreplace.ReplaceAll(tmpl.Content, vars)

	return &npool.TextInfo{
		Subject: title,
		Content: content,
	}, nil
}
