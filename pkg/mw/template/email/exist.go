package email

import (
	"context"
	"fmt"

	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entemailtemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/emailtemplate"
)

func (h *Handler) ExistEmailTemplate(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			EmailTemplate.
			Query().
			Where(
				entemailtemplate.EntID(*h.EntID),
				entemailtemplate.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistEmailTemplateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := emailtemplatecrud.SetQueryConds(cli.EmailTemplate.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
