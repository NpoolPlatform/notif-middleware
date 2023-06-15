package email

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/template/email"
	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateEmailTemplate(ctx context.Context, cli *ent.Client) error {
	if _, err := emailtemplatecrud.UpdateSet(
		cli.EmailTemplate.UpdateOneID(*h.ID),
		&emailtemplatecrud.Req{
			DefaultToUsername: h.DefaultToUsername,
			Sender:            h.Sender,
			ReplyTos:          h.ReplyTos,
			CcTos:             h.CcTos,
			Subject:           h.Subject,
			Body:              h.Body,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.LangID == nil {
		return nil, fmt.Errorf("invalid langid")
	}

	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixSetFiat,
		*h.ID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &emailtemplatecrud.Conds{
		ID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
	}
	h.Offset = 0
	h.Limit = 2

	email, err := h.GetEmailTemplateOnly(ctx)
	if err != nil {
		return nil, err
	}
	if email != nil {
		return nil, fmt.Errorf("emailtemplate exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateEmailTemplate(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetEmailTemplate(ctx)
}