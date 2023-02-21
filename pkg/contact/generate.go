package contact

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

func GenerateContact(
	ctx context.Context,
	subject, body, appID, sender, senderName string,
	usedFor basetypes.UsedFor,
) (*npool.TextInfo, error) {
	info, err := mgrcli.GetContactOnly(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UsedFor: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(usedFor.Number()),
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

	body = fmt.Sprintf("From: %v<br>Name: %v<br>%v", sender, senderName, body)
	body = strings.ReplaceAll(body, "\n", "<br>")

	return &npool.TextInfo{
		Subject:  subject,
		Content:  body,
		From:     info.Sender,
		To:       info.Account,
		ReplyTos: []string{sender},
	}, nil
}
