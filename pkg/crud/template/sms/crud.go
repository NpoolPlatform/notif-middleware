package sms

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entsmstemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/smstemplate"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	UsedFor   *basetypes.UsedFor
	Subject   *string
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.SMSTemplateCreate, req *Req) *ent.SMSTemplateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Subject != nil {
		c.SetSubject(*req.Subject)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.SMSTemplateUpdateOne, req *Req) *ent.SMSTemplateUpdateOne {
	if req.LangID != nil {
		u = u.SetLangID(*req.LangID)
	}
	if req.Subject != nil {
		u = u.SetSubject(*req.Subject)
	}
	if req.Message != nil {
		u = u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	AppID    *cruder.Cond
	LangID   *cruder.Cond
	UsedFor  *cruder.Cond
	IDs      *cruder.Cond
	AppIDs   *cruder.Cond
	LangIDs  *cruder.Cond
	UsedFors *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.SMSTemplateQuery, conds *Conds) (*ent.SMSTemplateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entsmstemplate.UsedFor(usedFor))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.AppIDs != nil {
		appids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.AppIDIn(appids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.LangIDs != nil {
		langids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(entsmstemplate.LangIDIn(langids...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.IDs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFors")
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entsmstemplate.UsedForIn(usedFors...))
		default:
			return nil, fmt.Errorf("invalid sms field")
		}
	}
	q.Where(entsmstemplate.DeletedAt(0))
	return q, nil
}
