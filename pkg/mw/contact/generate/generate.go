package generate

import (
	"context"
	"fmt"
	"strings"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/contact"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) GenerateContact(ctx context.Context) (*npool.TextInfo, error) {
	h.Conds.AppID = &cruder.Cond{
		Op:  cruder.EQ, Val: *h.AppID,
	}
	h.Conds.UsedFor = &cruder.Cond{
		Op:  cruder.EQ, Val: int32(h.UsedFor.Number()),
	}
	h.Conds.AccountType = &cruder.Cond{
		Op:  cruder.EQ, Val: int32(basetypes.SignMethod_Email),
	}

	info, err := h.GetContactOnly(ctx)
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
