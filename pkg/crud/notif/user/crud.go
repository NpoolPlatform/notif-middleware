package user

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entusernotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	NotifID   *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.UserNotifCreate, req *Req) *ent.UserNotifCreate {
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

func UpdateSet(u *ent.UserNotifUpdateOne, req *Req) *ent.UserNotifUpdateOne {
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	NotifID *cruder.Cond
	IDs     *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.UserNotifQuery, conds *Conds) (*ent.UserNotifQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entusernotif.ID(id))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entusernotif.AppID(appid))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entusernotif.UserID(userid))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.NotifID != nil {
		notifid, ok := conds.NotifID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid notifid")
		}
		switch conds.NotifID.Op {
		case cruder.EQ:
			q.Where(entusernotif.NotifID(notifid))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entusernotif.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid user field")
		}
	}
	q.Where(entusernotif.DeletedAt(0))
	return q, nil
}
