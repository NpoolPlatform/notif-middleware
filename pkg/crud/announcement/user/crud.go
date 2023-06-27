package user

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entuseramt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/userannouncement"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	AnnouncementID *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.UserAnnouncementCreate, req *Req) *ent.UserAnnouncementCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AnnouncementID != nil {
		c.SetAnnouncementID(*req.AnnouncementID)
	}
	return c
}

func UpdateSet(u *ent.UserAnnouncementUpdateOne, req *Req) *ent.UserAnnouncementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	AnnouncementID *cruder.Cond
}

func SetQueryConds(q *ent.UserAnnouncementQuery, conds *Conds) (*ent.UserAnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entuseramt.ID(id))
		default:
			return nil, fmt.Errorf("invalid user announcement id op field %s", conds.ID.Op)
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entuseramt.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app id op field %s", conds.AppID.Op)
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entuseramt.UserID(id))
		default:
			return nil, fmt.Errorf("invalid user id op field %s", conds.UserID.Op)
		}
	}
	if conds.AnnouncementID != nil {
		id, ok := conds.AnnouncementID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement id")
		}
		switch conds.AnnouncementID.Op {
		case cruder.EQ:
			q.Where(entuseramt.AnnouncementID(id))
		default:
			return nil, fmt.Errorf("invalid announcement id op field %s", conds.AnnouncementID.Op)
		}
	}
	return q, nil
}
