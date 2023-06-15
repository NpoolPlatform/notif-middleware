package frontend

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	notifmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	frontendtmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/frontend"
	frontendtmplmwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/template/frontend"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	tmplreplace "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/replace"

	"github.com/google/uuid"
)

func GenerateNotifs(
	ctx context.Context,
	appID, userID string,
	usedFor basetypes.UsedFor,
	vars *npool.TemplateVars,
) ([]*notifmwpb.NotifReq, error) {
	const maxTemplates = uint32(100)
	eventID := uuid.NewString()

	tmpls, _, err := frontendtmplmwcli.GetFrontendTemplates(ctx, &frontendtmplmwpb.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UsedFor: &basetypes.Uint32Val{
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

	reqs := []*notifmwpb.NotifReq{}
	for _, tmpl := range tmpls {
		title := tmplreplace.ReplaceAll(tmpl.Title, vars)
		content := tmplreplace.ReplaceAll(tmpl.Content, vars)
		useTemplate := true
		channel1 := basetypes.NotifChannel_ChannelFrontend

		reqs = append(reqs, &notifmwpb.NotifReq{
			AppID:       &appID,
			UserID:      &userID,
			LangID:      &tmpl.LangID,
			EventType:   &usedFor,
			UseTemplate: &useTemplate,
			Title:       &title,
			Content:     &content,
			Channel:     &channel1,
			EventID:     &eventID,
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
	tmpl, err := frontendtmplmwcli.GetFrontendTemplateOnly(ctx, &frontendtmplmwpb.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		LangID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: langID,
		},
		UsedFor: &basetypes.Uint32Val{
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
