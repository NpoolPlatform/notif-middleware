// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/usernotif"
	"github.com/google/uuid"
)

// UserNotifCreate is the builder for creating a UserNotif entity.
type UserNotifCreate struct {
	config
	mutation *UserNotifMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (unc *UserNotifCreate) SetCreatedAt(u uint32) *UserNotifCreate {
	unc.mutation.SetCreatedAt(u)
	return unc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableCreatedAt(u *uint32) *UserNotifCreate {
	if u != nil {
		unc.SetCreatedAt(*u)
	}
	return unc
}

// SetUpdatedAt sets the "updated_at" field.
func (unc *UserNotifCreate) SetUpdatedAt(u uint32) *UserNotifCreate {
	unc.mutation.SetUpdatedAt(u)
	return unc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableUpdatedAt(u *uint32) *UserNotifCreate {
	if u != nil {
		unc.SetUpdatedAt(*u)
	}
	return unc
}

// SetDeletedAt sets the "deleted_at" field.
func (unc *UserNotifCreate) SetDeletedAt(u uint32) *UserNotifCreate {
	unc.mutation.SetDeletedAt(u)
	return unc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableDeletedAt(u *uint32) *UserNotifCreate {
	if u != nil {
		unc.SetDeletedAt(*u)
	}
	return unc
}

// SetAppID sets the "app_id" field.
func (unc *UserNotifCreate) SetAppID(u uuid.UUID) *UserNotifCreate {
	unc.mutation.SetAppID(u)
	return unc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableAppID(u *uuid.UUID) *UserNotifCreate {
	if u != nil {
		unc.SetAppID(*u)
	}
	return unc
}

// SetUserID sets the "user_id" field.
func (unc *UserNotifCreate) SetUserID(u uuid.UUID) *UserNotifCreate {
	unc.mutation.SetUserID(u)
	return unc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableUserID(u *uuid.UUID) *UserNotifCreate {
	if u != nil {
		unc.SetUserID(*u)
	}
	return unc
}

// SetNotifID sets the "notif_id" field.
func (unc *UserNotifCreate) SetNotifID(u uuid.UUID) *UserNotifCreate {
	unc.mutation.SetNotifID(u)
	return unc
}

// SetNillableNotifID sets the "notif_id" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableNotifID(u *uuid.UUID) *UserNotifCreate {
	if u != nil {
		unc.SetNotifID(*u)
	}
	return unc
}

// SetID sets the "id" field.
func (unc *UserNotifCreate) SetID(u uuid.UUID) *UserNotifCreate {
	unc.mutation.SetID(u)
	return unc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (unc *UserNotifCreate) SetNillableID(u *uuid.UUID) *UserNotifCreate {
	if u != nil {
		unc.SetID(*u)
	}
	return unc
}

// Mutation returns the UserNotifMutation object of the builder.
func (unc *UserNotifCreate) Mutation() *UserNotifMutation {
	return unc.mutation
}

// Save creates the UserNotif in the database.
func (unc *UserNotifCreate) Save(ctx context.Context) (*UserNotif, error) {
	var (
		err  error
		node *UserNotif
	)
	if err := unc.defaults(); err != nil {
		return nil, err
	}
	if len(unc.hooks) == 0 {
		if err = unc.check(); err != nil {
			return nil, err
		}
		node, err = unc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserNotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = unc.check(); err != nil {
				return nil, err
			}
			unc.mutation = mutation
			if node, err = unc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(unc.hooks) - 1; i >= 0; i-- {
			if unc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = unc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, unc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserNotif)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserNotifMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (unc *UserNotifCreate) SaveX(ctx context.Context) *UserNotif {
	v, err := unc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (unc *UserNotifCreate) Exec(ctx context.Context) error {
	_, err := unc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (unc *UserNotifCreate) ExecX(ctx context.Context) {
	if err := unc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (unc *UserNotifCreate) defaults() error {
	if _, ok := unc.mutation.CreatedAt(); !ok {
		if usernotif.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultCreatedAt()
		unc.mutation.SetCreatedAt(v)
	}
	if _, ok := unc.mutation.UpdatedAt(); !ok {
		if usernotif.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultUpdatedAt()
		unc.mutation.SetUpdatedAt(v)
	}
	if _, ok := unc.mutation.DeletedAt(); !ok {
		if usernotif.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultDeletedAt()
		unc.mutation.SetDeletedAt(v)
	}
	if _, ok := unc.mutation.AppID(); !ok {
		if usernotif.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultAppID()
		unc.mutation.SetAppID(v)
	}
	if _, ok := unc.mutation.UserID(); !ok {
		if usernotif.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultUserID()
		unc.mutation.SetUserID(v)
	}
	if _, ok := unc.mutation.NotifID(); !ok {
		if usernotif.DefaultNotifID == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultNotifID (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultNotifID()
		unc.mutation.SetNotifID(v)
	}
	if _, ok := unc.mutation.ID(); !ok {
		if usernotif.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized usernotif.DefaultID (forgotten import ent/runtime?)")
		}
		v := usernotif.DefaultID()
		unc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (unc *UserNotifCreate) check() error {
	if _, ok := unc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserNotif.created_at"`)}
	}
	if _, ok := unc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserNotif.updated_at"`)}
	}
	if _, ok := unc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "UserNotif.deleted_at"`)}
	}
	return nil
}

func (unc *UserNotifCreate) sqlSave(ctx context.Context) (*UserNotif, error) {
	_node, _spec := unc.createSpec()
	if err := sqlgraph.CreateNode(ctx, unc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (unc *UserNotifCreate) createSpec() (*UserNotif, *sqlgraph.CreateSpec) {
	var (
		_node = &UserNotif{config: unc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usernotif.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: usernotif.FieldID,
			},
		}
	)
	_spec.OnConflict = unc.conflict
	if id, ok := unc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := unc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usernotif.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := unc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usernotif.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := unc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: usernotif.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := unc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usernotif.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := unc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usernotif.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := unc.mutation.NotifID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: usernotif.FieldNotifID,
		})
		_node.NotifID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserNotif.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserNotifUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (unc *UserNotifCreate) OnConflict(opts ...sql.ConflictOption) *UserNotifUpsertOne {
	unc.conflict = opts
	return &UserNotifUpsertOne{
		create: unc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserNotif.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (unc *UserNotifCreate) OnConflictColumns(columns ...string) *UserNotifUpsertOne {
	unc.conflict = append(unc.conflict, sql.ConflictColumns(columns...))
	return &UserNotifUpsertOne{
		create: unc,
	}
}

type (
	// UserNotifUpsertOne is the builder for "upsert"-ing
	//  one UserNotif node.
	UserNotifUpsertOne struct {
		create *UserNotifCreate
	}

	// UserNotifUpsert is the "OnConflict" setter.
	UserNotifUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserNotifUpsert) SetCreatedAt(v uint32) *UserNotifUpsert {
	u.Set(usernotif.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateCreatedAt() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserNotifUpsert) AddCreatedAt(v uint32) *UserNotifUpsert {
	u.Add(usernotif.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserNotifUpsert) SetUpdatedAt(v uint32) *UserNotifUpsert {
	u.Set(usernotif.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateUpdatedAt() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserNotifUpsert) AddUpdatedAt(v uint32) *UserNotifUpsert {
	u.Add(usernotif.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserNotifUpsert) SetDeletedAt(v uint32) *UserNotifUpsert {
	u.Set(usernotif.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateDeletedAt() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserNotifUpsert) AddDeletedAt(v uint32) *UserNotifUpsert {
	u.Add(usernotif.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserNotifUpsert) SetAppID(v uuid.UUID) *UserNotifUpsert {
	u.Set(usernotif.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateAppID() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserNotifUpsert) ClearAppID() *UserNotifUpsert {
	u.SetNull(usernotif.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserNotifUpsert) SetUserID(v uuid.UUID) *UserNotifUpsert {
	u.Set(usernotif.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateUserID() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserNotifUpsert) ClearUserID() *UserNotifUpsert {
	u.SetNull(usernotif.FieldUserID)
	return u
}

// SetNotifID sets the "notif_id" field.
func (u *UserNotifUpsert) SetNotifID(v uuid.UUID) *UserNotifUpsert {
	u.Set(usernotif.FieldNotifID, v)
	return u
}

// UpdateNotifID sets the "notif_id" field to the value that was provided on create.
func (u *UserNotifUpsert) UpdateNotifID() *UserNotifUpsert {
	u.SetExcluded(usernotif.FieldNotifID)
	return u
}

// ClearNotifID clears the value of the "notif_id" field.
func (u *UserNotifUpsert) ClearNotifID() *UserNotifUpsert {
	u.SetNull(usernotif.FieldNotifID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserNotif.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usernotif.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserNotifUpsertOne) UpdateNewValues() *UserNotifUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(usernotif.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserNotif.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserNotifUpsertOne) Ignore() *UserNotifUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserNotifUpsertOne) DoNothing() *UserNotifUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserNotifCreate.OnConflict
// documentation for more info.
func (u *UserNotifUpsertOne) Update(set func(*UserNotifUpsert)) *UserNotifUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserNotifUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserNotifUpsertOne) SetCreatedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserNotifUpsertOne) AddCreatedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateCreatedAt() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserNotifUpsertOne) SetUpdatedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserNotifUpsertOne) AddUpdatedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateUpdatedAt() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserNotifUpsertOne) SetDeletedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserNotifUpsertOne) AddDeletedAt(v uint32) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateDeletedAt() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserNotifUpsertOne) SetAppID(v uuid.UUID) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateAppID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserNotifUpsertOne) ClearAppID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserNotifUpsertOne) SetUserID(v uuid.UUID) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateUserID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserNotifUpsertOne) ClearUserID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearUserID()
	})
}

// SetNotifID sets the "notif_id" field.
func (u *UserNotifUpsertOne) SetNotifID(v uuid.UUID) *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetNotifID(v)
	})
}

// UpdateNotifID sets the "notif_id" field to the value that was provided on create.
func (u *UserNotifUpsertOne) UpdateNotifID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateNotifID()
	})
}

// ClearNotifID clears the value of the "notif_id" field.
func (u *UserNotifUpsertOne) ClearNotifID() *UserNotifUpsertOne {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearNotifID()
	})
}

// Exec executes the query.
func (u *UserNotifUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserNotifCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserNotifUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserNotifUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserNotifUpsertOne.ID is not supported by MySQL driver. Use UserNotifUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserNotifUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserNotifCreateBulk is the builder for creating many UserNotif entities in bulk.
type UserNotifCreateBulk struct {
	config
	builders []*UserNotifCreate
	conflict []sql.ConflictOption
}

// Save creates the UserNotif entities in the database.
func (uncb *UserNotifCreateBulk) Save(ctx context.Context) ([]*UserNotif, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uncb.builders))
	nodes := make([]*UserNotif, len(uncb.builders))
	mutators := make([]Mutator, len(uncb.builders))
	for i := range uncb.builders {
		func(i int, root context.Context) {
			builder := uncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserNotifMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uncb *UserNotifCreateBulk) SaveX(ctx context.Context) []*UserNotif {
	v, err := uncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uncb *UserNotifCreateBulk) Exec(ctx context.Context) error {
	_, err := uncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uncb *UserNotifCreateBulk) ExecX(ctx context.Context) {
	if err := uncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserNotif.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserNotifUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uncb *UserNotifCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserNotifUpsertBulk {
	uncb.conflict = opts
	return &UserNotifUpsertBulk{
		create: uncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserNotif.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uncb *UserNotifCreateBulk) OnConflictColumns(columns ...string) *UserNotifUpsertBulk {
	uncb.conflict = append(uncb.conflict, sql.ConflictColumns(columns...))
	return &UserNotifUpsertBulk{
		create: uncb,
	}
}

// UserNotifUpsertBulk is the builder for "upsert"-ing
// a bulk of UserNotif nodes.
type UserNotifUpsertBulk struct {
	create *UserNotifCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserNotif.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usernotif.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserNotifUpsertBulk) UpdateNewValues() *UserNotifUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(usernotif.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserNotif.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserNotifUpsertBulk) Ignore() *UserNotifUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserNotifUpsertBulk) DoNothing() *UserNotifUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserNotifCreateBulk.OnConflict
// documentation for more info.
func (u *UserNotifUpsertBulk) Update(set func(*UserNotifUpsert)) *UserNotifUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserNotifUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserNotifUpsertBulk) SetCreatedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserNotifUpsertBulk) AddCreatedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateCreatedAt() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserNotifUpsertBulk) SetUpdatedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserNotifUpsertBulk) AddUpdatedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateUpdatedAt() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserNotifUpsertBulk) SetDeletedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserNotifUpsertBulk) AddDeletedAt(v uint32) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateDeletedAt() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserNotifUpsertBulk) SetAppID(v uuid.UUID) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateAppID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserNotifUpsertBulk) ClearAppID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserNotifUpsertBulk) SetUserID(v uuid.UUID) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateUserID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserNotifUpsertBulk) ClearUserID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearUserID()
	})
}

// SetNotifID sets the "notif_id" field.
func (u *UserNotifUpsertBulk) SetNotifID(v uuid.UUID) *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.SetNotifID(v)
	})
}

// UpdateNotifID sets the "notif_id" field to the value that was provided on create.
func (u *UserNotifUpsertBulk) UpdateNotifID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.UpdateNotifID()
	})
}

// ClearNotifID clears the value of the "notif_id" field.
func (u *UserNotifUpsertBulk) ClearNotifID() *UserNotifUpsertBulk {
	return u.Update(func(s *UserNotifUpsert) {
		s.ClearNotifID()
	})
}

// Exec executes the query.
func (u *UserNotifUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserNotifCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserNotifCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserNotifUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}