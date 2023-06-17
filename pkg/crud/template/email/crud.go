package email

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entemailtemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/emailtemplate"

	"github.com/google/uuid"
)

type Req struct {
	ID                *uuid.UUID
	AppID             *uuid.UUID
	LangID            *uuid.UUID
	DefaultToUsername *string
	UsedFor           *basetypes.UsedFor
	Sender            *string
	ReplyTos          *[]string
	CcTos             *[]string
	Subject           *string
	Body              *string
	DeletedAt         *uint32
}

func CreateSet(c *ent.EmailTemplateCreate, req *Req) *ent.EmailTemplateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.DefaultToUsername != nil {
		c.SetDefaultToUsername(*req.DefaultToUsername)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Sender != nil {
		c.SetSender(*req.Sender)
	}
	if req.ReplyTos != nil {
		c.SetReplyTos(*req.ReplyTos)
	}
	if req.CcTos != nil {
		c.SetCcTos(*req.CcTos)
	}
	if req.Subject != nil {
		c.SetSubject(*req.Subject)
	}
	if req.Body != nil {
		c.SetBody(*req.Body)
	}
	return c
}

func UpdateSet(u *ent.EmailTemplateUpdateOne, req *Req) *ent.EmailTemplateUpdateOne {
	if req.DefaultToUsername != nil {
		u = u.SetDefaultToUsername(*req.DefaultToUsername)
	}
	if req.Sender != nil {
		u = u.SetSender(*req.Sender)
	}
	if req.ReplyTos != nil {
		u = u.SetReplyTos(*req.ReplyTos)
	}
	if req.CcTos != nil {
		u = u.SetCcTos(*req.CcTos)
	}
	if req.Subject != nil {
		u = u.SetSubject(*req.Subject)
	}
	if req.Body != nil {
		u = u.SetBody(*req.Body)
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
	Sender   *cruder.Cond
	IDs      *cruder.Cond
	AppIDs   *cruder.Cond
	LangIDs  *cruder.Cond
	UsedFors *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.EmailTemplateQuery, conds *Conds) (*ent.EmailTemplateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entemailtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entemailtemplate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(entemailtemplate.LangID(id))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid used for")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entemailtemplate.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid used for op field")
		}
	}
	if conds.Sender != nil {
		sender, ok := conds.Sender.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid sender")
		}
		switch conds.Sender.Op {
		case cruder.EQ:
			q.Where(entemailtemplate.Sender(sender))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entemailtemplate.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.AppIDs != nil {
		appids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entemailtemplate.AppIDIn(appids...))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.LangIDs != nil {
		langids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(entemailtemplate.LangIDIn(langids...))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.IDs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFors")
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entemailtemplate.UsedForIn(usedFors...))
		default:
			return nil, fmt.Errorf("invalid email field")
		}
	}
	q.Where(entemailtemplate.DeletedAt(0))
	return q, nil
}
