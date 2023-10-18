// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendannouncement"
	"github.com/google/uuid"
)

// SendAnnouncementCreate is the builder for creating a SendAnnouncement entity.
type SendAnnouncementCreate struct {
	config
	mutation *SendAnnouncementMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sac *SendAnnouncementCreate) SetCreatedAt(u uint32) *SendAnnouncementCreate {
	sac.mutation.SetCreatedAt(u)
	return sac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableCreatedAt(u *uint32) *SendAnnouncementCreate {
	if u != nil {
		sac.SetCreatedAt(*u)
	}
	return sac
}

// SetUpdatedAt sets the "updated_at" field.
func (sac *SendAnnouncementCreate) SetUpdatedAt(u uint32) *SendAnnouncementCreate {
	sac.mutation.SetUpdatedAt(u)
	return sac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableUpdatedAt(u *uint32) *SendAnnouncementCreate {
	if u != nil {
		sac.SetUpdatedAt(*u)
	}
	return sac
}

// SetDeletedAt sets the "deleted_at" field.
func (sac *SendAnnouncementCreate) SetDeletedAt(u uint32) *SendAnnouncementCreate {
	sac.mutation.SetDeletedAt(u)
	return sac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableDeletedAt(u *uint32) *SendAnnouncementCreate {
	if u != nil {
		sac.SetDeletedAt(*u)
	}
	return sac
}

// SetEntID sets the "ent_id" field.
func (sac *SendAnnouncementCreate) SetEntID(u uuid.UUID) *SendAnnouncementCreate {
	sac.mutation.SetEntID(u)
	return sac
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableEntID(u *uuid.UUID) *SendAnnouncementCreate {
	if u != nil {
		sac.SetEntID(*u)
	}
	return sac
}

// SetAppID sets the "app_id" field.
func (sac *SendAnnouncementCreate) SetAppID(u uuid.UUID) *SendAnnouncementCreate {
	sac.mutation.SetAppID(u)
	return sac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableAppID(u *uuid.UUID) *SendAnnouncementCreate {
	if u != nil {
		sac.SetAppID(*u)
	}
	return sac
}

// SetUserID sets the "user_id" field.
func (sac *SendAnnouncementCreate) SetUserID(u uuid.UUID) *SendAnnouncementCreate {
	sac.mutation.SetUserID(u)
	return sac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableUserID(u *uuid.UUID) *SendAnnouncementCreate {
	if u != nil {
		sac.SetUserID(*u)
	}
	return sac
}

// SetAnnouncementID sets the "announcement_id" field.
func (sac *SendAnnouncementCreate) SetAnnouncementID(u uuid.UUID) *SendAnnouncementCreate {
	sac.mutation.SetAnnouncementID(u)
	return sac
}

// SetNillableAnnouncementID sets the "announcement_id" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableAnnouncementID(u *uuid.UUID) *SendAnnouncementCreate {
	if u != nil {
		sac.SetAnnouncementID(*u)
	}
	return sac
}

// SetChannel sets the "channel" field.
func (sac *SendAnnouncementCreate) SetChannel(s string) *SendAnnouncementCreate {
	sac.mutation.SetChannel(s)
	return sac
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (sac *SendAnnouncementCreate) SetNillableChannel(s *string) *SendAnnouncementCreate {
	if s != nil {
		sac.SetChannel(*s)
	}
	return sac
}

// SetID sets the "id" field.
func (sac *SendAnnouncementCreate) SetID(u uint32) *SendAnnouncementCreate {
	sac.mutation.SetID(u)
	return sac
}

// Mutation returns the SendAnnouncementMutation object of the builder.
func (sac *SendAnnouncementCreate) Mutation() *SendAnnouncementMutation {
	return sac.mutation
}

// Save creates the SendAnnouncement in the database.
func (sac *SendAnnouncementCreate) Save(ctx context.Context) (*SendAnnouncement, error) {
	var (
		err  error
		node *SendAnnouncement
	)
	if err := sac.defaults(); err != nil {
		return nil, err
	}
	if len(sac.hooks) == 0 {
		if err = sac.check(); err != nil {
			return nil, err
		}
		node, err = sac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SendAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sac.check(); err != nil {
				return nil, err
			}
			sac.mutation = mutation
			if node, err = sac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sac.hooks) - 1; i >= 0; i-- {
			if sac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (sac *SendAnnouncementCreate) SaveX(ctx context.Context) *SendAnnouncement {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *SendAnnouncementCreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *SendAnnouncementCreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sac *SendAnnouncementCreate) defaults() error {
	if _, ok := sac.mutation.CreatedAt(); !ok {
		if sendannouncement.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultCreatedAt()
		sac.mutation.SetCreatedAt(v)
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		if sendannouncement.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultUpdatedAt()
		sac.mutation.SetUpdatedAt(v)
	}
	if _, ok := sac.mutation.DeletedAt(); !ok {
		if sendannouncement.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultDeletedAt()
		sac.mutation.SetDeletedAt(v)
	}
	if _, ok := sac.mutation.EntID(); !ok {
		if sendannouncement.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultEntID()
		sac.mutation.SetEntID(v)
	}
	if _, ok := sac.mutation.AppID(); !ok {
		if sendannouncement.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultAppID()
		sac.mutation.SetAppID(v)
	}
	if _, ok := sac.mutation.UserID(); !ok {
		if sendannouncement.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultUserID()
		sac.mutation.SetUserID(v)
	}
	if _, ok := sac.mutation.AnnouncementID(); !ok {
		if sendannouncement.DefaultAnnouncementID == nil {
			return fmt.Errorf("ent: uninitialized sendannouncement.DefaultAnnouncementID (forgotten import ent/runtime?)")
		}
		v := sendannouncement.DefaultAnnouncementID()
		sac.mutation.SetAnnouncementID(v)
	}
	if _, ok := sac.mutation.Channel(); !ok {
		v := sendannouncement.DefaultChannel
		sac.mutation.SetChannel(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sac *SendAnnouncementCreate) check() error {
	if _, ok := sac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SendAnnouncement.created_at"`)}
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SendAnnouncement.updated_at"`)}
	}
	if _, ok := sac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "SendAnnouncement.deleted_at"`)}
	}
	if _, ok := sac.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "SendAnnouncement.ent_id"`)}
	}
	return nil
}

func (sac *SendAnnouncementCreate) sqlSave(ctx context.Context) (*SendAnnouncement, error) {
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (sac *SendAnnouncementCreate) createSpec() (*SendAnnouncement, *sqlgraph.CreateSpec) {
	var (
		_node = &SendAnnouncement{config: sac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sendannouncement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: sendannouncement.FieldID,
			},
		}
	)
	_spec.OnConflict = sac.conflict
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: sendannouncement.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := sac.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := sac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := sac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := sac.mutation.AnnouncementID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: sendannouncement.FieldAnnouncementID,
		})
		_node.AnnouncementID = value
	}
	if value, ok := sac.mutation.Channel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sendannouncement.FieldChannel,
		})
		_node.Channel = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SendAnnouncement.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SendAnnouncementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sac *SendAnnouncementCreate) OnConflict(opts ...sql.ConflictOption) *SendAnnouncementUpsertOne {
	sac.conflict = opts
	return &SendAnnouncementUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SendAnnouncement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sac *SendAnnouncementCreate) OnConflictColumns(columns ...string) *SendAnnouncementUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &SendAnnouncementUpsertOne{
		create: sac,
	}
}

type (
	// SendAnnouncementUpsertOne is the builder for "upsert"-ing
	//  one SendAnnouncement node.
	SendAnnouncementUpsertOne struct {
		create *SendAnnouncementCreate
	}

	// SendAnnouncementUpsert is the "OnConflict" setter.
	SendAnnouncementUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *SendAnnouncementUpsert) SetCreatedAt(v uint32) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateCreatedAt() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SendAnnouncementUpsert) AddCreatedAt(v uint32) *SendAnnouncementUpsert {
	u.Add(sendannouncement.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SendAnnouncementUpsert) SetUpdatedAt(v uint32) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateUpdatedAt() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SendAnnouncementUpsert) AddUpdatedAt(v uint32) *SendAnnouncementUpsert {
	u.Add(sendannouncement.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SendAnnouncementUpsert) SetDeletedAt(v uint32) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateDeletedAt() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SendAnnouncementUpsert) AddDeletedAt(v uint32) *SendAnnouncementUpsert {
	u.Add(sendannouncement.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *SendAnnouncementUpsert) SetEntID(v uuid.UUID) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateEntID() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *SendAnnouncementUpsert) SetAppID(v uuid.UUID) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateAppID() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *SendAnnouncementUpsert) ClearAppID() *SendAnnouncementUpsert {
	u.SetNull(sendannouncement.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *SendAnnouncementUpsert) SetUserID(v uuid.UUID) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateUserID() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *SendAnnouncementUpsert) ClearUserID() *SendAnnouncementUpsert {
	u.SetNull(sendannouncement.FieldUserID)
	return u
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *SendAnnouncementUpsert) SetAnnouncementID(v uuid.UUID) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldAnnouncementID, v)
	return u
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateAnnouncementID() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldAnnouncementID)
	return u
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *SendAnnouncementUpsert) ClearAnnouncementID() *SendAnnouncementUpsert {
	u.SetNull(sendannouncement.FieldAnnouncementID)
	return u
}

// SetChannel sets the "channel" field.
func (u *SendAnnouncementUpsert) SetChannel(v string) *SendAnnouncementUpsert {
	u.Set(sendannouncement.FieldChannel, v)
	return u
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *SendAnnouncementUpsert) UpdateChannel() *SendAnnouncementUpsert {
	u.SetExcluded(sendannouncement.FieldChannel)
	return u
}

// ClearChannel clears the value of the "channel" field.
func (u *SendAnnouncementUpsert) ClearChannel() *SendAnnouncementUpsert {
	u.SetNull(sendannouncement.FieldChannel)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SendAnnouncement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sendannouncement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SendAnnouncementUpsertOne) UpdateNewValues() *SendAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sendannouncement.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.SendAnnouncement.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SendAnnouncementUpsertOne) Ignore() *SendAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SendAnnouncementUpsertOne) DoNothing() *SendAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SendAnnouncementCreate.OnConflict
// documentation for more info.
func (u *SendAnnouncementUpsertOne) Update(set func(*SendAnnouncementUpsert)) *SendAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SendAnnouncementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SendAnnouncementUpsertOne) SetCreatedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SendAnnouncementUpsertOne) AddCreatedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateCreatedAt() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SendAnnouncementUpsertOne) SetUpdatedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SendAnnouncementUpsertOne) AddUpdatedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateUpdatedAt() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SendAnnouncementUpsertOne) SetDeletedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SendAnnouncementUpsertOne) AddDeletedAt(v uint32) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateDeletedAt() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *SendAnnouncementUpsertOne) SetEntID(v uuid.UUID) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateEntID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *SendAnnouncementUpsertOne) SetAppID(v uuid.UUID) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateAppID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *SendAnnouncementUpsertOne) ClearAppID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *SendAnnouncementUpsertOne) SetUserID(v uuid.UUID) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateUserID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *SendAnnouncementUpsertOne) ClearUserID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *SendAnnouncementUpsertOne) SetAnnouncementID(v uuid.UUID) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateAnnouncementID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateAnnouncementID()
	})
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *SendAnnouncementUpsertOne) ClearAnnouncementID() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearAnnouncementID()
	})
}

// SetChannel sets the "channel" field.
func (u *SendAnnouncementUpsertOne) SetChannel(v string) *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetChannel(v)
	})
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *SendAnnouncementUpsertOne) UpdateChannel() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateChannel()
	})
}

// ClearChannel clears the value of the "channel" field.
func (u *SendAnnouncementUpsertOne) ClearChannel() *SendAnnouncementUpsertOne {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearChannel()
	})
}

// Exec executes the query.
func (u *SendAnnouncementUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SendAnnouncementCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SendAnnouncementUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SendAnnouncementUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SendAnnouncementUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SendAnnouncementCreateBulk is the builder for creating many SendAnnouncement entities in bulk.
type SendAnnouncementCreateBulk struct {
	config
	builders []*SendAnnouncementCreate
	conflict []sql.ConflictOption
}

// Save creates the SendAnnouncement entities in the database.
func (sacb *SendAnnouncementCreateBulk) Save(ctx context.Context) ([]*SendAnnouncement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SendAnnouncement, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SendAnnouncementMutation)
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
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *SendAnnouncementCreateBulk) SaveX(ctx context.Context) []*SendAnnouncement {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *SendAnnouncementCreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *SendAnnouncementCreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SendAnnouncement.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SendAnnouncementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sacb *SendAnnouncementCreateBulk) OnConflict(opts ...sql.ConflictOption) *SendAnnouncementUpsertBulk {
	sacb.conflict = opts
	return &SendAnnouncementUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SendAnnouncement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sacb *SendAnnouncementCreateBulk) OnConflictColumns(columns ...string) *SendAnnouncementUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &SendAnnouncementUpsertBulk{
		create: sacb,
	}
}

// SendAnnouncementUpsertBulk is the builder for "upsert"-ing
// a bulk of SendAnnouncement nodes.
type SendAnnouncementUpsertBulk struct {
	create *SendAnnouncementCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SendAnnouncement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sendannouncement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SendAnnouncementUpsertBulk) UpdateNewValues() *SendAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sendannouncement.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SendAnnouncement.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SendAnnouncementUpsertBulk) Ignore() *SendAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SendAnnouncementUpsertBulk) DoNothing() *SendAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SendAnnouncementCreateBulk.OnConflict
// documentation for more info.
func (u *SendAnnouncementUpsertBulk) Update(set func(*SendAnnouncementUpsert)) *SendAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SendAnnouncementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SendAnnouncementUpsertBulk) SetCreatedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SendAnnouncementUpsertBulk) AddCreatedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateCreatedAt() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SendAnnouncementUpsertBulk) SetUpdatedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SendAnnouncementUpsertBulk) AddUpdatedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateUpdatedAt() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SendAnnouncementUpsertBulk) SetDeletedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SendAnnouncementUpsertBulk) AddDeletedAt(v uint32) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateDeletedAt() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *SendAnnouncementUpsertBulk) SetEntID(v uuid.UUID) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateEntID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *SendAnnouncementUpsertBulk) SetAppID(v uuid.UUID) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateAppID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *SendAnnouncementUpsertBulk) ClearAppID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *SendAnnouncementUpsertBulk) SetUserID(v uuid.UUID) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateUserID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *SendAnnouncementUpsertBulk) ClearUserID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *SendAnnouncementUpsertBulk) SetAnnouncementID(v uuid.UUID) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateAnnouncementID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateAnnouncementID()
	})
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *SendAnnouncementUpsertBulk) ClearAnnouncementID() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearAnnouncementID()
	})
}

// SetChannel sets the "channel" field.
func (u *SendAnnouncementUpsertBulk) SetChannel(v string) *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.SetChannel(v)
	})
}

// UpdateChannel sets the "channel" field to the value that was provided on create.
func (u *SendAnnouncementUpsertBulk) UpdateChannel() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.UpdateChannel()
	})
}

// ClearChannel clears the value of the "channel" field.
func (u *SendAnnouncementUpsertBulk) ClearChannel() *SendAnnouncementUpsertBulk {
	return u.Update(func(s *SendAnnouncementUpsert) {
		s.ClearChannel()
	})
}

// Exec executes the query.
func (u *SendAnnouncementUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SendAnnouncementCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SendAnnouncementCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SendAnnouncementUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
