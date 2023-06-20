package readstate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entreadnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/readnotif"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	NotifID   *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.ReadNotifCreate, req *Req) *ent.ReadNotifCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.NotifID != nil {
		c.SetNotifID(*req.NotifID)
	}
	return c
}

func UpdateSet(u *ent.ReadNotifUpdateOne, req *Req) *ent.ReadNotifUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	NotifID *cruder.Cond
	IDs     *cruder.Cond
}

// nolint:gocyclo
func SetQueryConds(q *ent.ReadNotifQuery, conds *Conds) (*ent.ReadNotifQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entreadnotif.ID(id))
		default:
			return nil, fmt.Errorf("invalid readstate field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entreadnotif.AppID(appid))
		default:
			return nil, fmt.Errorf("invalid readstate field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entreadnotif.UserID(userid))
		default:
			return nil, fmt.Errorf("invalid readstate field")
		}
	}
	if conds.NotifID != nil {
		notifid, ok := conds.NotifID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid notifid")
		}
		switch conds.NotifID.Op {
		case cruder.EQ:
			q.Where(entreadnotif.NotifID(notifid))
		default:
			return nil, fmt.Errorf("invalid readstate field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entreadnotif.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid readstate field")
		}
	}
	q.Where(entreadnotif.DeletedAt(0))
	return q, nil
}
