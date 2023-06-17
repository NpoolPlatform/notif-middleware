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

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createEmailTemplate(ctx context.Context, cli *ent.Client) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid lang")
	}
	if h.LangID == nil {
		return fmt.Errorf("invalid logo")
	}
	if h.UsedFor == nil {
		return fmt.Errorf("invalid usedFor")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*h.AppID,
		*h.LangID,
		h.UsedFor,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &emailtemplatecrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		LangID:  &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
		UsedFor: &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistEmailTemplateConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("emailtemplate exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	info, err := emailtemplatecrud.CreateSet(
		cli.EmailTemplate.Create(),
		&emailtemplatecrud.Req{
			ID:                h.ID,
			AppID:             h.AppID,
			LangID:            h.LangID,
			UsedFor:           h.UsedFor,
			Sender:            h.Sender,
			ReplyTos:          h.ReplyTos,
			CcTos:             h.CcTos,
			Subject:           h.Subject,
			Body:              h.Body,
			DefaultToUsername: h.DefaultToUsername,
		},
	).Save(ctx)
	if err != nil {
		return err
	}
	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateEmailTemplate(ctx context.Context) (*npool.EmailTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createEmailTemplate(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetEmailTemplate(ctx)
}

func (h *Handler) CreateEmailTemplates(ctx context.Context) ([]*npool.EmailTemplate, error) {
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.AppID = req.AppID
			handler.LangID = req.LangID
			handler.UsedFor = req.UsedFor
			handler.Sender = req.Sender
			handler.ReplyTos = req.ReplyTos
			handler.CcTos = req.CcTos
			handler.Subject = req.Subject
			handler.Body = req.Body
			handler.DefaultToUsername = req.DefaultToUsername
			if err := handler.createEmailTemplate(ctx, cli); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &emailtemplatecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetEmailTemplates(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
