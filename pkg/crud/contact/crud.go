package contact

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/contact"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type Req struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	UsedFor     *basetypes.UsedFor
	Account     *string
	AccountType *basetypes.SignMethod
	Sender      *string
	DeletedAt   *uint32
}

func CreateSet(c *ent.ContactCreate, in *npool.ContactReq) *ent.ContactCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UsedFor != nil {
		c.SetUsedFor(in.GetUsedFor().String())
	}
	if in.Account != nil {
		c.SetAccount(in.GetAccount())
	}
	if in.AccountType != nil {
		c.SetAccountType(in.GetAccountType().String())
	}
	if in.Sender != nil {
		c.SetSender(in.GetSender())
	}
	return c
}

func Update(u *ent.ContactUpdateOne, req *Req) (*ent.ContactUpdateOne, error) {
	if req.Account != nil {
		u.SetAccount(*req.Account)
	}
	if req.AccountType != nil {
		u.SetAccountType(req.AccountType.String())
	}
	if req.Sender != nil {
		u.SetSender(*req.Sender)
	}

	return u, nil
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.ContactQuery, conds *npool.Conds) (*ent.ContactQuery, error) {
	if conds == nil {
		return q, nil
	}

	if conds.ID != nil {
		id, err := uuid.Parse(conds.GetID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetID().GetOp() {
		case cruder.EQ:
			q.Where(contact.ID(id))
		default:
			return nil, fmt.Errorf("invalid contact field")
		}
	}

	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			q.Where(contact.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field")
		}
	}

	if conds.AccountType != nil {
		switch conds.GetAccountType().GetOp() {
		case cruder.EQ:
			q.Where(contact.AccountType(basetypes.SignMethod(conds.GetAccountType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid account type op field")
		}
	}

	if conds.UsedFor != nil {
		switch conds.GetUsedFor().GetOp() {
		case cruder.EQ:
			q.Where(contact.UsedFor(basetypes.UsedFor(conds.GetUsedFor().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid used for op field")
		}
	}

	return q, nil
}
