package email

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	notifmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	emailtmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtmplmwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/template/email"

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
	const maxTemplates = int32(100)
	eventID := uuid.NewString()

	tmpls, _, err := emailtmplmwcli.GetEmailTemplates(ctx, &emailtmplmwpb.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UsedFor: &basetypes.Int32Val{
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

	reqs := []*notifmwpb.NotifReq{}
	for _, tmpl := range tmpls {
		title := tmplreplace.ReplaceAll(tmpl.Subject, vars)
		content := tmplreplace.ReplaceAll(tmpl.Body, vars)
		useTemplate := true
		channel1 := basetypes.NotifChannel_ChannelEmail

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
	tmpl, err := emailtmplmwcli.GetEmailTemplateOnly(ctx, &emailtmplmwpb.Conds{
		AppID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		LangID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: langID,
		},
		UsedFor: &basetypes.Int32Val{
			Op:    cruder.EQ,
			Value: int32(usedFor),
		},
	})
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, nil
	}

	title := tmplreplace.ReplaceAll(tmpl.Subject, vars)
	content := tmplreplace.ReplaceAll(tmpl.Body, vars)

	return &npool.TextInfo{
		Subject:  title,
		Content:  content,
		From:     tmpl.Sender,
		ToCCs:    tmpl.CCTos,
		ReplyTos: tmpl.ReplyTos,
	}, nil
}
