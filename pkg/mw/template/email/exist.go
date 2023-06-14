package email

import (
	"context"

	emailtemplatecrud "github.com/NpoolPlatform/notif-middleware/pkg/crud/template/email"

	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
)

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
