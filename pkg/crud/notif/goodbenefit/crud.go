package goodbenefit

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entgoodbenefit "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/goodbenefit"
	"github.com/google/uuid"
)

type Req struct {
	EntID       *uuid.UUID
	GoodID      *uuid.UUID
	GoodName    *string
	Amount      *string
	State       *basetypes.Result
	Message     *string
	BenefitDate *uint32
	TxID        *uuid.UUID
	Generated   *bool
	DeletedAt   *uint32
}

func CreateSet(c *ent.GoodBenefitCreate, req *Req) *ent.GoodBenefitCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.GoodName != nil {
		c.SetGoodName(*req.GoodName)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.BenefitDate != nil {
		c.SetBenefitDate(*req.BenefitDate)
	}
	if req.TxID != nil {
		c.SetTxID(*req.TxID)
	}
	c.SetGenerated(false)
	return c
}

func UpdateSet(u *ent.GoodBenefitUpdateOne, req *Req) *ent.GoodBenefitUpdateOne {
	if req.Generated != nil {
		u.SetGenerated(*req.Generated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	GoodID           *cruder.Cond
	Generated        *cruder.Cond
	BenefitDateStart *cruder.Cond
	BenefitDateEnd   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.GoodBenefitQuery, conds *Conds) (*ent.GoodBenefitQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.ID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit id op field %s", conds.ID.Op)
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit op field %s", conds.EntID.Op)
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit good id")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid good benefit good id op field %s", conds.GoodID.Op)
		}
	}
	if conds.Generated != nil {
		notified, ok := conds.Generated.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit notified")
		}
		switch conds.Generated.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.Generated(notified))
		default:
			return nil, fmt.Errorf("invalid good benefit notified op field %s", conds.Generated.Op)
		}
	}
	if conds.BenefitDateStart != nil {
		_date, ok := conds.BenefitDateStart.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit benefit date %s", conds.BenefitDateStart.Op)
		}
		switch conds.BenefitDateStart.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		case cruder.LTE:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		case cruder.LT:
			q.Where(entgoodbenefit.BenefitDateLT(_date))
		case cruder.GTE:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		default:
			return nil, fmt.Errorf("invalid good benefit benefit date op field %s", conds.BenefitDateStart.Op)
		}
	}
	if conds.BenefitDateEnd != nil {
		_date, ok := conds.BenefitDateEnd.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid good benefit benefit date %s", conds.BenefitDateEnd.Op)
		}
		switch conds.BenefitDateEnd.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		case cruder.GTE:
			q.Where(entgoodbenefit.BenefitDateGTE(_date))
		case cruder.GT:
			q.Where(entgoodbenefit.BenefitDateGT(_date))
		case cruder.LTE:
			q.Where(entgoodbenefit.BenefitDateLTE(_date))
		default:
			return nil, fmt.Errorf("invalid good benefit benefit date op field %s", conds.BenefitDateEnd.Op)
		}
	}
	return q, nil
}
