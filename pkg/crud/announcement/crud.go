package announcement

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mw/v1/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	entamt "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	Title     *string
	Content   *string
	Channel   *basetypes.NotifChannel
	Type      *npool.AnnouncementType
	EndAt     *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.AnnouncementCreate, req *Req) *ent.AnnouncementCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
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
	if req.Type != nil {
		c.SetType(req.Type.String())
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	return c
}

func UpdateSet(u *ent.AnnouncementUpdateOne, req *Req) *ent.AnnouncementUpdateOne {
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Content != nil {
		u.SetContent(*req.Content)
	}
	if req.Channel != nil {
		u.SetChannel(req.Channel.String())
	}
	if req.Type != nil {
		u.SetType(req.Type.String())
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	EndAt   *cruder.Cond
	Channel *cruder.Cond
}

func SetQueryConds(q *ent.AnnouncementQuery, conds *Conds) (*ent.AnnouncementQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid announcement id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entamt.ID(id))
		default:
			return nil, fmt.Errorf("invalid announcement op field")
		}
	}

	if conds.EndAt != nil {
		endAt, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end at")
		}
		switch conds.EndAt.Op {
		case cruder.GTE:
			q.Where(entamt.EndAt(endAt))
		case cruder.EQ:
			q.Where(entamt.EndAt(endAt))
		default:
			return nil, fmt.Errorf("invalid end at op field")
		}
	}

	if conds.Channel != nil {
		channel, ok := conds.Channel.Val.(basetypes.NotifChannel)
		if !ok {
			return nil, fmt.Errorf("invalid channel")
		}
		switch conds.Channel.Op {
		case cruder.EQ:
			q.Where(entamt.Channel(channel.String()))
		default:
			return nil, fmt.Errorf("invalid channel op field")
		}
	}

	return q, nil
}
