package generate

import (
	"context"
	"fmt"
	"strings"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	mgrpb "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"

	mgrcli "github.com/NpoolPlatform/notif-manager/pkg/client/contact"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func (h *Handler) GenerateContact(ctx context.Context) (*npool.TextInfo, error) {
	info, err := mgrcli.GetContactOnly(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: h.AppID.String(),
		},
		UsedFor: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(h.UsedFor.Number()),
		},
		AccountType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(basetypes.SignMethod_Email),
		},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("contact not exist")
	}

	_body := fmt.Sprintf("From: %v<br>Name: %v<br>%v", h.Sender, h.SenderName, h.Body)
	body := strings.ReplaceAll(_body, "\n", "<br>")

	return &npool.TextInfo{
		Subject:  *h.Subject,
		Content:  body,
		From:     info.Sender,
		To:       info.Account,
		ReplyTos: []string{*h.Sender},
	}, nil
}
