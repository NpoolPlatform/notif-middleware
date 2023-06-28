package goodbenefit

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.GoodID == nil {
		return fmt.Errorf("good id is empty")
	}
	if h.GoodName == nil {
		return fmt.Errorf("good name id is empty")
	}
	if h.Amount == nil {
		return fmt.Errorf("amount is empty")
	}
	if h.State == nil {
		return fmt.Errorf("state is empty")
	}
	if h.BenefitDate == nil {
		return fmt.Errorf("benefit date is empty")
	}
	if h.Generated == nil {
		return fmt.Errorf("generated is empty")
	}
	if *h.State == basetypes.Result_Success {
		if h.Amount == nil || h.TxID == nil {
			return fmt.Errorf("amount or tx id can not be empty")
		}
	}

	return nil
}

func (h *Handler) CreateGoodBenefit(ctx context.Context) (*npool.GoodBenefit, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.GoodBenefit.Create(),
			&crud.Req{
				ID:          h.ID,
				GoodID:      h.GoodID,
				GoodName:    h.GoodName,
				Amount:      h.Amount,
				State:       h.State,
				Message:     h.Message,
				BenefitDate: h.BenefitDate,
				TxID:        h.TxID,
				Generated:   h.Generated,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGoodBenefit(ctx)
}
