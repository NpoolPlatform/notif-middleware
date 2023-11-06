// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendannouncement"
	"github.com/google/uuid"
)

// SendAnnouncementUpdate is the builder for updating SendAnnouncement entities.
type SendAnnouncementUpdate struct {
	config
	hooks     []Hook
	mutation  *SendAnnouncementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SendAnnouncementUpdate builder.
func (sau *SendAnnouncementUpdate) Where(ps ...predicate.SendAnnouncement) *SendAnnouncementUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetCreatedAt sets the "created_at" field.
func (sau *SendAnnouncementUpdate) SetCreatedAt(u uint32) *SendAnnouncementUpdate {
	sau.mutation.ResetCreatedAt()
	sau.mutation.SetCreatedAt(u)
	return sau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableCreatedAt(u *uint32) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetCreatedAt(*u)
	}
	return sau
}

// AddCreatedAt adds u to the "created_at" field.
func (sau *SendAnnouncementUpdate) AddCreatedAt(u int32) *SendAnnouncementUpdate {
	sau.mutation.AddCreatedAt(u)
	return sau
}

// SetUpdatedAt sets the "updated_at" field.
func (sau *SendAnnouncementUpdate) SetUpdatedAt(u uint32) *SendAnnouncementUpdate {
	sau.mutation.ResetUpdatedAt()
	sau.mutation.SetUpdatedAt(u)
	return sau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (sau *SendAnnouncementUpdate) AddUpdatedAt(u int32) *SendAnnouncementUpdate {
	sau.mutation.AddUpdatedAt(u)
	return sau
}

// SetDeletedAt sets the "deleted_at" field.
func (sau *SendAnnouncementUpdate) SetDeletedAt(u uint32) *SendAnnouncementUpdate {
	sau.mutation.ResetDeletedAt()
	sau.mutation.SetDeletedAt(u)
	return sau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableDeletedAt(u *uint32) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetDeletedAt(*u)
	}
	return sau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (sau *SendAnnouncementUpdate) AddDeletedAt(u int32) *SendAnnouncementUpdate {
	sau.mutation.AddDeletedAt(u)
	return sau
}

// SetEntID sets the "ent_id" field.
func (sau *SendAnnouncementUpdate) SetEntID(u uuid.UUID) *SendAnnouncementUpdate {
	sau.mutation.SetEntID(u)
	return sau
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableEntID(u *uuid.UUID) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetEntID(*u)
	}
	return sau
}

// SetAppID sets the "app_id" field.
func (sau *SendAnnouncementUpdate) SetAppID(u uuid.UUID) *SendAnnouncementUpdate {
	sau.mutation.SetAppID(u)
	return sau
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableAppID(u *uuid.UUID) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetAppID(*u)
	}
	return sau
}

// ClearAppID clears the value of the "app_id" field.
func (sau *SendAnnouncementUpdate) ClearAppID() *SendAnnouncementUpdate {
	sau.mutation.ClearAppID()
	return sau
}

// SetUserID sets the "user_id" field.
func (sau *SendAnnouncementUpdate) SetUserID(u uuid.UUID) *SendAnnouncementUpdate {
	sau.mutation.SetUserID(u)
	return sau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableUserID(u *uuid.UUID) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetUserID(*u)
	}
	return sau
}

// ClearUserID clears the value of the "user_id" field.
func (sau *SendAnnouncementUpdate) ClearUserID() *SendAnnouncementUpdate {
	sau.mutation.ClearUserID()
	return sau
}

// SetAnnouncementID sets the "announcement_id" field.
func (sau *SendAnnouncementUpdate) SetAnnouncementID(u uuid.UUID) *SendAnnouncementUpdate {
	sau.mutation.SetAnnouncementID(u)
	return sau
}

// SetNillableAnnouncementID sets the "announcement_id" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableAnnouncementID(u *uuid.UUID) *SendAnnouncementUpdate {
	if u != nil {
		sau.SetAnnouncementID(*u)
	}
	return sau
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (sau *SendAnnouncementUpdate) ClearAnnouncementID() *SendAnnouncementUpdate {
	sau.mutation.ClearAnnouncementID()
	return sau
}

// SetChannel sets the "channel" field.
func (sau *SendAnnouncementUpdate) SetChannel(s string) *SendAnnouncementUpdate {
	sau.mutation.SetChannel(s)
	return sau
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (sau *SendAnnouncementUpdate) SetNillableChannel(s *string) *SendAnnouncementUpdate {
	if s != nil {
		sau.SetChannel(*s)
	}
	return sau
}

// ClearChannel clears the value of the "channel" field.
func (sau *SendAnnouncementUpdate) ClearChannel() *SendAnnouncementUpdate {
	sau.mutation.ClearChannel()
	return sau
}

// Mutation returns the SendAnnouncementMutation object of the builder.
func (sau *SendAnnouncementUpdate) Mutation() *SendAnnouncementMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SendAnnouncementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := sau.defaults(); err != nil {
		return 0, err
	}
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SendAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SendAnnouncementUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SendAnnouncementUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SendAnnouncementUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *SendAnnouncementUpdate) defaults() error {
	if _, ok := sau.mutation.UpdatedAt(); !ok {
		if sendannouncement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := sendannouncement.UpdateDefaultUpdatedAt()
		sau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sau *SendAnnouncementUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SendAnnouncementUpdate {
	sau.modifiers = append(sau.modifiers, modifiers...)
	return sau
}

func (sau *SendAnnouncementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sendannouncement.Table,
			Columns: sendannouncement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: sendannouncement.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldCreatedAt,
		})
	}
	if value, ok := sau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldCreatedAt,
		})
	}
	if value, ok := sau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldUpdatedAt,
		})
	}
	if value, ok := sau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldUpdatedAt,
		})
	}
	if value, ok := sau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldDeletedAt,
		})
	}
	if value, ok := sau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldDeletedAt,
		})
	}
	if value, ok := sau.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldEntID,
		})
	}
	if value, ok := sau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAppID,
		})
	}
	if sau.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldAppID,
		})
	}
	if value, ok := sau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldUserID,
		})
	}
	if sau.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldUserID,
		})
	}
	if value, ok := sau.mutation.AnnouncementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAnnouncementID,
		})
	}
	if sau.mutation.AnnouncementIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldAnnouncementID,
		})
	}
	if value, ok := sau.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sendannouncement.FieldChannel,
		})
	}
	if sau.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sendannouncement.FieldChannel,
		})
	}
	_spec.Modifiers = sau.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sendannouncement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SendAnnouncementUpdateOne is the builder for updating a single SendAnnouncement entity.
type SendAnnouncementUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SendAnnouncementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (sauo *SendAnnouncementUpdateOne) SetCreatedAt(u uint32) *SendAnnouncementUpdateOne {
	sauo.mutation.ResetCreatedAt()
	sauo.mutation.SetCreatedAt(u)
	return sauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableCreatedAt(u *uint32) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetCreatedAt(*u)
	}
	return sauo
}

// AddCreatedAt adds u to the "created_at" field.
func (sauo *SendAnnouncementUpdateOne) AddCreatedAt(u int32) *SendAnnouncementUpdateOne {
	sauo.mutation.AddCreatedAt(u)
	return sauo
}

// SetUpdatedAt sets the "updated_at" field.
func (sauo *SendAnnouncementUpdateOne) SetUpdatedAt(u uint32) *SendAnnouncementUpdateOne {
	sauo.mutation.ResetUpdatedAt()
	sauo.mutation.SetUpdatedAt(u)
	return sauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (sauo *SendAnnouncementUpdateOne) AddUpdatedAt(u int32) *SendAnnouncementUpdateOne {
	sauo.mutation.AddUpdatedAt(u)
	return sauo
}

// SetDeletedAt sets the "deleted_at" field.
func (sauo *SendAnnouncementUpdateOne) SetDeletedAt(u uint32) *SendAnnouncementUpdateOne {
	sauo.mutation.ResetDeletedAt()
	sauo.mutation.SetDeletedAt(u)
	return sauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableDeletedAt(u *uint32) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetDeletedAt(*u)
	}
	return sauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (sauo *SendAnnouncementUpdateOne) AddDeletedAt(u int32) *SendAnnouncementUpdateOne {
	sauo.mutation.AddDeletedAt(u)
	return sauo
}

// SetEntID sets the "ent_id" field.
func (sauo *SendAnnouncementUpdateOne) SetEntID(u uuid.UUID) *SendAnnouncementUpdateOne {
	sauo.mutation.SetEntID(u)
	return sauo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableEntID(u *uuid.UUID) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetEntID(*u)
	}
	return sauo
}

// SetAppID sets the "app_id" field.
func (sauo *SendAnnouncementUpdateOne) SetAppID(u uuid.UUID) *SendAnnouncementUpdateOne {
	sauo.mutation.SetAppID(u)
	return sauo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableAppID(u *uuid.UUID) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetAppID(*u)
	}
	return sauo
}

// ClearAppID clears the value of the "app_id" field.
func (sauo *SendAnnouncementUpdateOne) ClearAppID() *SendAnnouncementUpdateOne {
	sauo.mutation.ClearAppID()
	return sauo
}

// SetUserID sets the "user_id" field.
func (sauo *SendAnnouncementUpdateOne) SetUserID(u uuid.UUID) *SendAnnouncementUpdateOne {
	sauo.mutation.SetUserID(u)
	return sauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableUserID(u *uuid.UUID) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetUserID(*u)
	}
	return sauo
}

// ClearUserID clears the value of the "user_id" field.
func (sauo *SendAnnouncementUpdateOne) ClearUserID() *SendAnnouncementUpdateOne {
	sauo.mutation.ClearUserID()
	return sauo
}

// SetAnnouncementID sets the "announcement_id" field.
func (sauo *SendAnnouncementUpdateOne) SetAnnouncementID(u uuid.UUID) *SendAnnouncementUpdateOne {
	sauo.mutation.SetAnnouncementID(u)
	return sauo
}

// SetNillableAnnouncementID sets the "announcement_id" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableAnnouncementID(u *uuid.UUID) *SendAnnouncementUpdateOne {
	if u != nil {
		sauo.SetAnnouncementID(*u)
	}
	return sauo
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (sauo *SendAnnouncementUpdateOne) ClearAnnouncementID() *SendAnnouncementUpdateOne {
	sauo.mutation.ClearAnnouncementID()
	return sauo
}

// SetChannel sets the "channel" field.
func (sauo *SendAnnouncementUpdateOne) SetChannel(s string) *SendAnnouncementUpdateOne {
	sauo.mutation.SetChannel(s)
	return sauo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (sauo *SendAnnouncementUpdateOne) SetNillableChannel(s *string) *SendAnnouncementUpdateOne {
	if s != nil {
		sauo.SetChannel(*s)
	}
	return sauo
}

// ClearChannel clears the value of the "channel" field.
func (sauo *SendAnnouncementUpdateOne) ClearChannel() *SendAnnouncementUpdateOne {
	sauo.mutation.ClearChannel()
	return sauo
}

// Mutation returns the SendAnnouncementMutation object of the builder.
func (sauo *SendAnnouncementUpdateOne) Mutation() *SendAnnouncementMutation {
	return sauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SendAnnouncementUpdateOne) Select(field string, fields ...string) *SendAnnouncementUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SendAnnouncement entity.
func (sauo *SendAnnouncementUpdateOne) Save(ctx context.Context) (*SendAnnouncement, error) {
	var (
		err  error
		node *SendAnnouncement
	)
	if err := sauo.defaults(); err != nil {
		return nil, err
	}
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SendAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SendAnnouncement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SendAnnouncementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SendAnnouncementUpdateOne) SaveX(ctx context.Context) *SendAnnouncement {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SendAnnouncementUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SendAnnouncementUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *SendAnnouncementUpdateOne) defaults() error {
	if _, ok := sauo.mutation.UpdatedAt(); !ok {
		if sendannouncement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := sendannouncement.UpdateDefaultUpdatedAt()
		sauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sauo *SendAnnouncementUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SendAnnouncementUpdateOne {
	sauo.modifiers = append(sauo.modifiers, modifiers...)
	return sauo
}

func (sauo *SendAnnouncementUpdateOne) sqlSave(ctx context.Context) (_node *SendAnnouncement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sendannouncement.Table,
			Columns: sendannouncement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: sendannouncement.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SendAnnouncement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sendannouncement.FieldID)
		for _, f := range fields {
			if !sendannouncement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sendannouncement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldCreatedAt,
		})
	}
	if value, ok := sauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldCreatedAt,
		})
	}
	if value, ok := sauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldUpdatedAt,
		})
	}
	if value, ok := sauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldUpdatedAt,
		})
	}
	if value, ok := sauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldDeletedAt,
		})
	}
	if value, ok := sauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldDeletedAt,
		})
	}
	if value, ok := sauo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldEntID,
		})
	}
	if value, ok := sauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAppID,
		})
	}
	if sauo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldAppID,
		})
	}
	if value, ok := sauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldUserID,
		})
	}
	if sauo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldUserID,
		})
	}
	if value, ok := sauo.mutation.AnnouncementID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAnnouncementID,
		})
	}
	if sauo.mutation.AnnouncementIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: sendannouncement.FieldAnnouncementID,
		})
	}
	if value, ok := sauo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sendannouncement.FieldChannel,
		})
	}
	if sauo.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sendannouncement.FieldChannel,
		})
	}
	_spec.Modifiers = sauo.modifiers
	_node = &SendAnnouncement{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sendannouncement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
