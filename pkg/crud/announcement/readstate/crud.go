package readstate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entreadamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readannouncement"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	UserID         *uuid.UUID
	AnnouncementID *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.ReadAnnouncementCreate, req *Req) *ent.ReadAnnouncementCreate {
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

func UpdateSet(u *ent.ReadAnnouncementUpdateOne, req *Req) *ent.ReadAnnouncementUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID              *cruder.Cond
	AppID           *cruder.Cond
	UserID          *cruder.Cond
	AnnouncementID  *cruder.Cond
	AnnouncementIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.ReadAnnouncementQuery, conds *Conds) (*ent.ReadAnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid read announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entreadamt.ID(id))
		default:
			return nil, fmt.Errorf("invalid read announcement id op field %s", conds.ID.Op)
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entreadamt.AppID(id))
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
			q.Where(entreadamt.UserID(id))
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
			q.Where(entreadamt.AnnouncementID(id))
		default:
			return nil, fmt.Errorf("invalid announcement id op field %s", conds.AnnouncementID.Op)
		}
	}
	if conds.AnnouncementIDs != nil {
		ids, ok := conds.AnnouncementIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement ids")
		}
		switch conds.AnnouncementIDs.Op {
		case cruder.IN:
			q.Where(entreadamt.AnnouncementIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid announcement ids op field %s", conds.AnnouncementIDs.Op)
		}
	}
	return q, nil
}
