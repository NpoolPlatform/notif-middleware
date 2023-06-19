package email

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template"

	notifmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif"

	emailtmplmwpb "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	tmplreplace "github.com/NpoolPlatform/notif-middleware/pkg/mw/template/replace"

	"github.com/google/uuid"
)

func (h *Handler) GenerateNotifs(ctx context.Context) ([]*notifmwpb.NotifReq, error) {
	const maxTemplates = int32(100)
	eventID := uuid.NewString()
	appID := h.AppID.String()
	userID := h.UserID.String()

	emailtmplHandler, err := NewHandler(
		ctx,
		WithConds(&emailtmplmwpb.Conds{
			AppID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			UsedFor: &basetypes.Int32Val{
				Op:    cruder.EQ,
				Value: int32(*h.UsedFor),
			},
		}),
		WithOffset(0),
		WithLimit(maxTemplates),
	)
	if err != nil {
		return nil, err
	}

	tmpls, _, err := emailtmplHandler.GetEmailTemplates(ctx)
	if err != nil {
		return nil, err
	}
	if len(tmpls) == 0 {
		return nil, nil
	}

	reqs := []*notifmwpb.NotifReq{}
	for _, tmpl := range tmpls {
		title := tmplreplace.ReplaceAll(tmpl.Subject, h.Vars)
		content := tmplreplace.ReplaceAll(tmpl.Body, h.Vars)
		useTemplate := true
		channel1 := basetypes.NotifChannel_ChannelEmail

		reqs = append(reqs, &notifmwpb.NotifReq{
			AppID:       &appID,
			UserID:      &userID,
			LangID:      &tmpl.LangID,
			EventType:   h.UsedFor,
			UseTemplate: &useTemplate,
			Title:       &title,
			Content:     &content,
			Channel:     &channel1,
			EventID:     &eventID,
		})
	}

	return reqs, nil
}

func (h *Handler) GenerateText(ctx context.Context) (*npool.TextInfo, error) {
	appID := h.AppID.String()
	langID := h.LangID.String()
	emailtmplHandler, err := NewHandler(
		ctx,
		WithConds(&emailtmplmwpb.Conds{
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
				Value: int32(*h.UsedFor),
			},
		}),
	)
	if err != nil {
		return nil, err
	}
	tmpl, err := emailtmplHandler.GetEmailTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if tmpl == nil {
		return nil, nil
	}

	title := tmplreplace.ReplaceAll(tmpl.Subject, h.Vars)
	content := tmplreplace.ReplaceAll(tmpl.Body, h.Vars)

	return &npool.TextInfo{
		Subject:  title,
		Content:  content,
		From:     tmpl.Sender,
		ToCCs:    tmpl.CCTos,
		ReplyTos: tmpl.ReplyTos,
	}, nil
}
