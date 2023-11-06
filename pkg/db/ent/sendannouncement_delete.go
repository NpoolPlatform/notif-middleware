// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendannouncement"
)

// SendAnnouncementDelete is the builder for deleting a SendAnnouncement entity.
type SendAnnouncementDelete struct {
	config
	hooks    []Hook
	mutation *SendAnnouncementMutation
}

// Where appends a list predicates to the SendAnnouncementDelete builder.
func (sad *SendAnnouncementDelete) Where(ps ...predicate.SendAnnouncement) *SendAnnouncementDelete {
	sad.mutation.Where(ps...)
	return sad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sad *SendAnnouncementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sad.hooks) == 0 {
		affected, err = sad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SendAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sad.mutation = mutation
			affected, err = sad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sad.hooks) - 1; i >= 0; i-- {
			if sad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sad *SendAnnouncementDelete) ExecX(ctx context.Context) int {
	n, err := sad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sad *SendAnnouncementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sendannouncement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: sendannouncement.FieldID,
			},
		},
	}
	if ps := sad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SendAnnouncementDeleteOne is the builder for deleting a single SendAnnouncement entity.
type SendAnnouncementDeleteOne struct {
	sad *SendAnnouncementDelete
}

// Exec executes the deletion query.
func (sado *SendAnnouncementDeleteOne) Exec(ctx context.Context) error {
	n, err := sado.sad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sendannouncement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sado *SendAnnouncementDeleteOne) ExecX(ctx context.Context) {
	sado.sad.ExecX(ctx)
}
