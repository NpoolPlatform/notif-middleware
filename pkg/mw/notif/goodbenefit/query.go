package goodbenefit

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/notif-middleware/pkg/crud/notif/goodbenefit"
	"github.com/NpoolPlatform/notif-middleware/pkg/db"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entgoodbenefit "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/goodbenefit"
)

type queryHandler struct {
	*Handler
	stm   *ent.GoodBenefitSelect
	infos []*npool.GoodBenefit
	total uint32
}

func (h *queryHandler) selectGoodBenefit(stm *ent.GoodBenefitQuery) {
	h.stm = stm.Select(
		entgoodbenefit.FieldID,
		entgoodbenefit.FieldGoodID,
		entgoodbenefit.FieldGoodName,
		entgoodbenefit.FieldAmount,
		entgoodbenefit.FieldState,
		entgoodbenefit.FieldMessage,
		entgoodbenefit.FieldBenefitDate,
		entgoodbenefit.FieldTxID,
		entgoodbenefit.FieldCreatedAt,
		entgoodbenefit.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.State = basetypes.Result(basetypes.Result_value[info.StateStr])
	}
}

func (h *queryHandler) queryGoodBenefit(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("good benefit id is empty")
	}
	h.selectGoodBenefit(
		cli.GoodBenefit.
			Query().
			Where(
				entgoodbenefit.ID(*h.ID),
				entgoodbenefit.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryGoodBenefitsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.GoodBenefit.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectGoodBenefit(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetGoodBenefits(ctx context.Context) ([]*npool.GoodBenefit, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBenefitsByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetGoodBenefit(ctx context.Context) (info *npool.GoodBenefit, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBenefit(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
