package frontend

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entfrontendtemplate "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/frontendtemplate"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	UsedFor   *string
	Title     *string
	Content   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.FrontendTemplateCreate, req *Req) *ent.FrontendTemplateCreate {
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
		c.SetUsedFor(*req.UsedFor)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Content != nil {
		c.SetContent(*req.Content)
	}
	return c
}

func UpdateSet(u *ent.FrontendTemplateUpdateOne, req *Req) *ent.FrontendTemplateUpdateOne {
	if req.LangID != nil {
		u = u.SetLangID(*req.LangID)
	}
	if req.UsedFor != nil {
		u = u.SetUsedFor(*req.UsedFor)
	}
	if req.Title != nil {
		u = u.SetTitle(*req.Title)
	}
	if req.Content != nil {
		u = u.SetContent(*req.Content)
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
func SetQueryConds(q *ent.FrontendTemplateQuery, conds *Conds) (*ent.FrontendTemplateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entfrontendtemplate.UsedFor(usedFor))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.AppIDs != nil {
		appids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.AppIDIn(appids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.LangIDs != nil {
		langids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.LangIDs.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.LangIDIn(langids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.IDs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid usedFors")
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entfrontendtemplate.UsedForIn(usedFors...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	q.Where(entfrontendtemplate.DeletedAt(0))
	return q, nil
}
