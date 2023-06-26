package notif

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entnotif "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notif"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	LangID      *uuid.UUID
	UserID      *uuid.UUID
	EventID     *uuid.UUID
	Notified    *bool
	EventType   *basetypes.UsedFor
	UseTemplate *bool
	Title       *string
	Content     *string
	Channel     *basetypes.NotifChannel
	Extra       *string
	NotifType   *basetypes.NotifType
	DeletedAt   *uint32
}

func CreateSet(c *ent.NotifCreate, req *Req) *ent.NotifCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.Notified != nil {
		c.SetNotified(*req.Notified)
	}
	if req.EventType != nil {
		c.SetEventType(req.EventType.String())
	}
	if req.UseTemplate != nil {
		c.SetUseTemplate(*req.UseTemplate)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Content != nil {
		c.SetContent(*req.Content)
	}
	if req.Channel != nil {
		c.SetChannel(req.Channel.String())
	}
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}
	if req.NotifType != nil {
		c.SetType(req.NotifType.String())
	}
	return c
}

func UpdateSet(u *ent.NotifUpdateOne, req *Req) *ent.NotifUpdateOne {
	if req.Title != nil {
		u = u.SetTitle(*req.Title)
	}
	if req.Content != nil {
		u = u.SetContent(*req.Content)
	}
	if req.Channel != nil {
		u = u.SetChannel(req.Channel.String())
	}
	if req.Notified != nil {
		u = u.SetNotified(*req.Notified)
	}
	if req.UseTemplate != nil {
		u = u.SetUseTemplate(*req.UseTemplate)
	}
	if req.Extra != nil {
		u = u.SetExtra(*req.Extra)
	}
	if req.NotifType != nil {
		u = u.SetType(req.NotifType.String())
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	LangID      *cruder.Cond
	Notified    *cruder.Cond
	EventType   *cruder.Cond
	UseTemplate *cruder.Cond
	Channel     *cruder.Cond
	Extra       *cruder.Cond
	Type        *cruder.Cond
	EventID     *cruder.Cond
	IDs         *cruder.Cond
	EventTypes  *cruder.Cond
	Channels    *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.NotifQuery, conds *Conds) (*ent.NotifQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entnotif.ID(id))
		case cruder.NEQ:
			q.Where(entnotif.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.AppID != nil {
		appid, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entnotif.AppID(appid))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.UserID != nil {
		userid, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entnotif.UserID(userid))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.LangID != nil {
		langid, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(entnotif.LangID(langid))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.EventID != nil {
		eventid, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(entnotif.EventID(eventid))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.Notified != nil {
		notified, ok := conds.Notified.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid notified")
		}
		switch conds.Notified.Op {
		case cruder.EQ:
			q.Where(entnotif.Notified(notified))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.EventType != nil {
		eventType, ok := conds.EventType.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid eventType")
		}
		switch conds.EventType.Op {
		case cruder.EQ:
			q.Where(entnotif.EventType(eventType.String()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.Type != nil {
		_type, ok := conds.Type.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid type")
		}
		switch conds.Type.Op {
		case cruder.EQ:
			q.Where(entnotif.Type(_type))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.UseTemplate != nil {
		useTemplate, ok := conds.UseTemplate.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid useTemplate")
		}
		switch conds.UseTemplate.Op {
		case cruder.EQ:
			q.Where(entnotif.UseTemplate(useTemplate))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channelsss")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entnotif.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.Extra != nil {
		extra, ok := conds.Extra.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid extra")
		}
		switch conds.Extra.Op {
		case cruder.LIKE:
			q.Where(entnotif.ExtraContains(extra))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entnotif.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.EventTypes != nil {
		eventTypes, ok := conds.EventTypes.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid eventTypes")
		}
		switch conds.EventTypes.Op {
		case cruder.IN:
			q.Where(entnotif.EventTypeIn(eventTypes...))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.Channels != nil {
		channels, ok := conds.Channels.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid channels")
		}
		switch conds.Channels.Op {
		case cruder.IN:
			q.Where(entnotif.ChannelIn(channels...))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	q.Where(entnotif.DeletedAt(0))
	return q, nil
}
