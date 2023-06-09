// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/txnotifstate"
)

// TxNotifStateDelete is the builder for deleting a TxNotifState entity.
type TxNotifStateDelete struct {
	config
	hooks    []Hook
	mutation *TxNotifStateMutation
}

// Where appends a list predicates to the TxNotifStateDelete builder.
func (tnsd *TxNotifStateDelete) Where(ps ...predicate.TxNotifState) *TxNotifStateDelete {
	tnsd.mutation.Where(ps...)
	return tnsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tnsd *TxNotifStateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tnsd.hooks) == 0 {
		affected, err = tnsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TxNotifStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tnsd.mutation = mutation
			affected, err = tnsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tnsd.hooks) - 1; i >= 0; i-- {
			if tnsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tnsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tnsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tnsd *TxNotifStateDelete) ExecX(ctx context.Context) int {
	n, err := tnsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tnsd *TxNotifStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: txnotifstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: txnotifstate.FieldID,
			},
		},
	}
	if ps := tnsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tnsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TxNotifStateDeleteOne is the builder for deleting a single TxNotifState entity.
type TxNotifStateDeleteOne struct {
	tnsd *TxNotifStateDelete
}

// Exec executes the deletion query.
func (tnsdo *TxNotifStateDeleteOne) Exec(ctx context.Context) error {
	n, err := tnsdo.tnsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{txnotifstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tnsdo *TxNotifStateDeleteOne) ExecX(ctx context.Context) {
	tnsdo.tnsd.ExecX(ctx)
}