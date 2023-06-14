package contact

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

func CreateSet(c *ent.ContactCreate, in *Req) *ent.ContactCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.AppID != nil {
		c.SetAppID(*in.AppID)
	}
	if in.UsedFor != nil {
		c.SetUsedFor(in.UsedFor.String())
	}
	if in.Account != nil {
		c.SetAccount(*in.Account)
	}
	if in.AccountType != nil {
		c.SetAccountType(in.AccountType.String())
	}
	if in.Sender != nil {
		c.SetSender(*in.Sender)
	}
	return c
}

func UpdateSet(u *ent.ContactUpdateOne, req *Req) *ent.ContactUpdateOne {
	if req.Account != nil {
		u.SetAccount(*req.Account)
	}
	if req.AccountType != nil {
		u.SetAccountType(req.AccountType.String())
	}
	if req.Sender != nil {
		u.SetSender(*req.Sender)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	AppID       *cruder.Cond
	AccountType *cruder.Cond
	UsedFor     *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.ContactQuery, conds *Conds) (*ent.ContactQuery, error) {
	if conds == nil {
		return q, nil
	}

	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contact id")
		}

		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(contact.ID(id))
		default:
			return nil, fmt.Errorf("invalid contact field")
		}
	}

	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}

		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(contact.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field")
		}
	}

	if conds.AccountType != nil {
		accountType, ok := conds.AccountType.Val.(basetypes.SignMethod)
		if !ok {
			return nil, fmt.Errorf("invalid account type")
		}

		switch conds.AccountType.Op {
		case cruder.EQ:
			q.Where(contact.AccountType(accountType.String()))
		default:
			return nil, fmt.Errorf("invalid account type op field")
		}
	}

	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid used for")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(contact.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid used for op field")
		}
	}

	return q, nil
}
