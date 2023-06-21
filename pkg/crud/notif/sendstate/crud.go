package sendstate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entsendnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendnotif"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	EventID   *uuid.UUID
	Channel   *basetypes.NotifChannel
	DeletedAt *uint32
}

func CreateSet(c *ent.SendNotifCreate, req *Req) *ent.SendNotifCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.Channel != nil {
		c.SetChannel(req.Channel.String())
	}
	return c
}

func UpdateSet(u *ent.SendNotifUpdateOne, req *Req) *ent.SendNotifUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	EventID *cruder.Cond
	Channel *cruder.Cond
	EndAt   *cruder.Cond
	IDs     *cruder.Cond
	UserIDs *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.SendNotifQuery, conds *Conds) (*ent.SendNotifQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsendnotif.ID(id))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entsendnotif.AppID(appid))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entsendnotif.UserID(userid))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.EventID != nil {
		eventid, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(entsendnotif.EventID(eventid))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channel")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entsendnotif.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entsendnotif.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	if conds.UserIDs != nil {
		userids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entsendnotif.UserIDIn(userids...))
		default:
			return nil, fmt.Errorf("invalid sendstate field")
		}
	}
	q.Where(entsendnotif.DeletedAt(0))
	return q, nil
}
