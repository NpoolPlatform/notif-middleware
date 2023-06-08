// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendnotif"
)

// SendNotifDelete is the builder for deleting a SendNotif entity.
type SendNotifDelete struct {
	config
	hooks    []Hook
	mutation *SendNotifMutation
}

// Where appends a list predicates to the SendNotifDelete builder.
func (snd *SendNotifDelete) Where(ps ...predicate.SendNotif) *SendNotifDelete {
	snd.mutation.Where(ps...)
	return snd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (snd *SendNotifDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(snd.hooks) == 0 {
		affected, err = snd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SendNotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			snd.mutation = mutation
			affected, err = snd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(snd.hooks) - 1; i >= 0; i-- {
			if snd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = snd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, snd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (snd *SendNotifDelete) ExecX(ctx context.Context) int {
	n, err := snd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (snd *SendNotifDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sendnotif.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: sendnotif.FieldID,
			},
		},
	}
	if ps := snd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, snd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SendNotifDeleteOne is the builder for deleting a single SendNotif entity.
type SendNotifDeleteOne struct {
	snd *SendNotifDelete
}

// Exec executes the deletion query.
func (sndo *SendNotifDeleteOne) Exec(ctx context.Context) error {
	n, err := sndo.snd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sendnotif.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sndo *SendNotifDeleteOne) ExecX(ctx context.Context) {
	sndo.snd.ExecX(ctx)
}
