package user

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entusernotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	EventType *basetypes.UsedFor
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
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	return c
}

func UpdateSet(u *ent.UserNotifUpdateOne, req *Req) *ent.UserNotifUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	EventType *cruder.Cond
	IDs       *cruder.Cond
}

// nolint:gocyclo
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
	if conds.EventType != nil {
		eventtype, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid notifid")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entusernotif.EventType(eventtype.String()))
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
