package tx

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"
	enttx "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/txnotifstate"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	TxID       *uuid.UUID
	NotifState *basetypes.TxState
	TxType     *basetypes.UsedFor
	DeletedAt  *uint32
}

func CreateSet(c *ent.TxNotifStateCreate, req *Req) *ent.TxNotifStateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.TxID != nil {
		c.SetTxID(*req.TxID)
	}
	if req.NotifState != nil {
		c.SetNotifState(req.NotifState.String())
	}
	if req.TxType != nil {
		c.SetTxType(req.TxType.String())
	}
	return c
}

func UpdateSet(u *ent.TxNotifStateUpdateOne, req *Req) *ent.TxNotifStateUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	TxID       *cruder.Cond
	NotifState *cruder.Cond
	TxType     *cruder.Cond
}

func SetQueryConds(q *ent.TxNotifStateQuery, conds *Conds) (*ent.TxNotifStateQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid tx state id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enttx.ID(id))
		default:
			return nil, fmt.Errorf("invalid tx state id op field")
		}
	}
	if conds.TxID != nil {
		id, ok := conds.TxID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid tx id")
		}
		switch conds.TxID.Op {
		case cruder.EQ:
			q.Where(enttx.TxID(id))
		default:
			return nil, fmt.Errorf("invalid tx id op field")
		}
	}
	if conds.NotifState != nil {
		txState, ok := conds.NotifState.Val.(basetypes.TxState)
		if !ok {
			return nil, fmt.Errorf("invalid tx state")
		}
		switch conds.NotifState.Op {
		case cruder.EQ:
			q.Where(enttx.NotifState(txState.String()))
		default:
			return nil, fmt.Errorf("invalid tx state op field")
		}
	}
	if conds.TxType != nil {
		_type, ok := conds.TxType.Val.(basetypes.TxType)
		if !ok {
			return nil, fmt.Errorf("invalid tx type")
		}
		switch conds.TxType.Op {
		case cruder.EQ:
			q.Where(enttx.TxType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid tx type op field")
		}
	}
	return q, nil
}
