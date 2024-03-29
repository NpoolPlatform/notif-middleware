// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/announcement"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AnnouncementUpdate is the builder for updating Announcement entities.
type AnnouncementUpdate struct {
	config
	hooks     []Hook
	mutation  *AnnouncementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AnnouncementUpdate builder.
func (au *AnnouncementUpdate) Where(ps ...predicate.Announcement) *AnnouncementUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AnnouncementUpdate) SetCreatedAt(u uint32) *AnnouncementUpdate {
	au.mutation.ResetCreatedAt()
	au.mutation.SetCreatedAt(u)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableCreatedAt(u *uint32) *AnnouncementUpdate {
	if u != nil {
		au.SetCreatedAt(*u)
	}
	return au
}

// AddCreatedAt adds u to the "created_at" field.
func (au *AnnouncementUpdate) AddCreatedAt(u int32) *AnnouncementUpdate {
	au.mutation.AddCreatedAt(u)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AnnouncementUpdate) SetUpdatedAt(u uint32) *AnnouncementUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(u)
	return au
}

// AddUpdatedAt adds u to the "updated_at" field.
func (au *AnnouncementUpdate) AddUpdatedAt(u int32) *AnnouncementUpdate {
	au.mutation.AddUpdatedAt(u)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AnnouncementUpdate) SetDeletedAt(u uint32) *AnnouncementUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(u)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableDeletedAt(u *uint32) *AnnouncementUpdate {
	if u != nil {
		au.SetDeletedAt(*u)
	}
	return au
}

// AddDeletedAt adds u to the "deleted_at" field.
func (au *AnnouncementUpdate) AddDeletedAt(u int32) *AnnouncementUpdate {
	au.mutation.AddDeletedAt(u)
	return au
}

// SetEntID sets the "ent_id" field.
func (au *AnnouncementUpdate) SetEntID(u uuid.UUID) *AnnouncementUpdate {
	au.mutation.SetEntID(u)
	return au
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableEntID(u *uuid.UUID) *AnnouncementUpdate {
	if u != nil {
		au.SetEntID(*u)
	}
	return au
}

// SetAppID sets the "app_id" field.
func (au *AnnouncementUpdate) SetAppID(u uuid.UUID) *AnnouncementUpdate {
	au.mutation.SetAppID(u)
	return au
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableAppID(u *uuid.UUID) *AnnouncementUpdate {
	if u != nil {
		au.SetAppID(*u)
	}
	return au
}

// ClearAppID clears the value of the "app_id" field.
func (au *AnnouncementUpdate) ClearAppID() *AnnouncementUpdate {
	au.mutation.ClearAppID()
	return au
}

// SetLangID sets the "lang_id" field.
func (au *AnnouncementUpdate) SetLangID(u uuid.UUID) *AnnouncementUpdate {
	au.mutation.SetLangID(u)
	return au
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableLangID(u *uuid.UUID) *AnnouncementUpdate {
	if u != nil {
		au.SetLangID(*u)
	}
	return au
}

// ClearLangID clears the value of the "lang_id" field.
func (au *AnnouncementUpdate) ClearLangID() *AnnouncementUpdate {
	au.mutation.ClearLangID()
	return au
}

// SetTitle sets the "title" field.
func (au *AnnouncementUpdate) SetTitle(s string) *AnnouncementUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableTitle(s *string) *AnnouncementUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// ClearTitle clears the value of the "title" field.
func (au *AnnouncementUpdate) ClearTitle() *AnnouncementUpdate {
	au.mutation.ClearTitle()
	return au
}

// SetContent sets the "content" field.
func (au *AnnouncementUpdate) SetContent(s string) *AnnouncementUpdate {
	au.mutation.SetContent(s)
	return au
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableContent(s *string) *AnnouncementUpdate {
	if s != nil {
		au.SetContent(*s)
	}
	return au
}

// ClearContent clears the value of the "content" field.
func (au *AnnouncementUpdate) ClearContent() *AnnouncementUpdate {
	au.mutation.ClearContent()
	return au
}

// SetChannel sets the "channel" field.
func (au *AnnouncementUpdate) SetChannel(s string) *AnnouncementUpdate {
	au.mutation.SetChannel(s)
	return au
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableChannel(s *string) *AnnouncementUpdate {
	if s != nil {
		au.SetChannel(*s)
	}
	return au
}

// ClearChannel clears the value of the "channel" field.
func (au *AnnouncementUpdate) ClearChannel() *AnnouncementUpdate {
	au.mutation.ClearChannel()
	return au
}

// SetStartAt sets the "start_at" field.
func (au *AnnouncementUpdate) SetStartAt(u uint32) *AnnouncementUpdate {
	au.mutation.ResetStartAt()
	au.mutation.SetStartAt(u)
	return au
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableStartAt(u *uint32) *AnnouncementUpdate {
	if u != nil {
		au.SetStartAt(*u)
	}
	return au
}

// AddStartAt adds u to the "start_at" field.
func (au *AnnouncementUpdate) AddStartAt(u int32) *AnnouncementUpdate {
	au.mutation.AddStartAt(u)
	return au
}

// ClearStartAt clears the value of the "start_at" field.
func (au *AnnouncementUpdate) ClearStartAt() *AnnouncementUpdate {
	au.mutation.ClearStartAt()
	return au
}

// SetEndAt sets the "end_at" field.
func (au *AnnouncementUpdate) SetEndAt(u uint32) *AnnouncementUpdate {
	au.mutation.ResetEndAt()
	au.mutation.SetEndAt(u)
	return au
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableEndAt(u *uint32) *AnnouncementUpdate {
	if u != nil {
		au.SetEndAt(*u)
	}
	return au
}

// AddEndAt adds u to the "end_at" field.
func (au *AnnouncementUpdate) AddEndAt(u int32) *AnnouncementUpdate {
	au.mutation.AddEndAt(u)
	return au
}

// ClearEndAt clears the value of the "end_at" field.
func (au *AnnouncementUpdate) ClearEndAt() *AnnouncementUpdate {
	au.mutation.ClearEndAt()
	return au
}

// SetType sets the "type" field.
func (au *AnnouncementUpdate) SetType(s string) *AnnouncementUpdate {
	au.mutation.SetType(s)
	return au
}

// SetNillableType sets the "type" field if the given value is not nil.
func (au *AnnouncementUpdate) SetNillableType(s *string) *AnnouncementUpdate {
	if s != nil {
		au.SetType(*s)
	}
	return au
}

// ClearType clears the value of the "type" field.
func (au *AnnouncementUpdate) ClearType() *AnnouncementUpdate {
	au.mutation.ClearType()
	return au
}

// Mutation returns the AnnouncementMutation object of the builder.
func (au *AnnouncementUpdate) Mutation() *AnnouncementMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnnouncementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := au.defaults(); err != nil {
		return 0, err
	}
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnnouncementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnnouncementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnnouncementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AnnouncementUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if announcement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized announcement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := announcement.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AnnouncementUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AnnouncementUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AnnouncementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: announcement.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldEntID,
		})
	}
	if value, ok := au.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldAppID,
		})
	}
	if au.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: announcement.FieldAppID,
		})
	}
	if value, ok := au.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldLangID,
		})
	}
	if au.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: announcement.FieldLangID,
		})
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTitle,
		})
	}
	if au.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldTitle,
		})
	}
	if value, ok := au.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldContent,
		})
	}
	if au.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldContent,
		})
	}
	if value, ok := au.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldChannel,
		})
	}
	if au.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldChannel,
		})
	}
	if value, ok := au.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldStartAt,
		})
	}
	if value, ok := au.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldStartAt,
		})
	}
	if au.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: announcement.FieldStartAt,
		})
	}
	if value, ok := au.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldEndAt,
		})
	}
	if value, ok := au.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldEndAt,
		})
	}
	if au.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: announcement.FieldEndAt,
		})
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldType,
		})
	}
	if au.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldType,
		})
	}
	_spec.Modifiers = au.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{announcement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AnnouncementUpdateOne is the builder for updating a single Announcement entity.
type AnnouncementUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AnnouncementMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (auo *AnnouncementUpdateOne) SetCreatedAt(u uint32) *AnnouncementUpdateOne {
	auo.mutation.ResetCreatedAt()
	auo.mutation.SetCreatedAt(u)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableCreatedAt(u *uint32) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetCreatedAt(*u)
	}
	return auo
}

// AddCreatedAt adds u to the "created_at" field.
func (auo *AnnouncementUpdateOne) AddCreatedAt(u int32) *AnnouncementUpdateOne {
	auo.mutation.AddCreatedAt(u)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AnnouncementUpdateOne) SetUpdatedAt(u uint32) *AnnouncementUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(u)
	return auo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auo *AnnouncementUpdateOne) AddUpdatedAt(u int32) *AnnouncementUpdateOne {
	auo.mutation.AddUpdatedAt(u)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AnnouncementUpdateOne) SetDeletedAt(u uint32) *AnnouncementUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(u)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableDeletedAt(u *uint32) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetDeletedAt(*u)
	}
	return auo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auo *AnnouncementUpdateOne) AddDeletedAt(u int32) *AnnouncementUpdateOne {
	auo.mutation.AddDeletedAt(u)
	return auo
}

// SetEntID sets the "ent_id" field.
func (auo *AnnouncementUpdateOne) SetEntID(u uuid.UUID) *AnnouncementUpdateOne {
	auo.mutation.SetEntID(u)
	return auo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableEntID(u *uuid.UUID) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetEntID(*u)
	}
	return auo
}

// SetAppID sets the "app_id" field.
func (auo *AnnouncementUpdateOne) SetAppID(u uuid.UUID) *AnnouncementUpdateOne {
	auo.mutation.SetAppID(u)
	return auo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableAppID(u *uuid.UUID) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetAppID(*u)
	}
	return auo
}

// ClearAppID clears the value of the "app_id" field.
func (auo *AnnouncementUpdateOne) ClearAppID() *AnnouncementUpdateOne {
	auo.mutation.ClearAppID()
	return auo
}

// SetLangID sets the "lang_id" field.
func (auo *AnnouncementUpdateOne) SetLangID(u uuid.UUID) *AnnouncementUpdateOne {
	auo.mutation.SetLangID(u)
	return auo
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableLangID(u *uuid.UUID) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetLangID(*u)
	}
	return auo
}

// ClearLangID clears the value of the "lang_id" field.
func (auo *AnnouncementUpdateOne) ClearLangID() *AnnouncementUpdateOne {
	auo.mutation.ClearLangID()
	return auo
}

// SetTitle sets the "title" field.
func (auo *AnnouncementUpdateOne) SetTitle(s string) *AnnouncementUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableTitle(s *string) *AnnouncementUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// ClearTitle clears the value of the "title" field.
func (auo *AnnouncementUpdateOne) ClearTitle() *AnnouncementUpdateOne {
	auo.mutation.ClearTitle()
	return auo
}

// SetContent sets the "content" field.
func (auo *AnnouncementUpdateOne) SetContent(s string) *AnnouncementUpdateOne {
	auo.mutation.SetContent(s)
	return auo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableContent(s *string) *AnnouncementUpdateOne {
	if s != nil {
		auo.SetContent(*s)
	}
	return auo
}

// ClearContent clears the value of the "content" field.
func (auo *AnnouncementUpdateOne) ClearContent() *AnnouncementUpdateOne {
	auo.mutation.ClearContent()
	return auo
}

// SetChannel sets the "channel" field.
func (auo *AnnouncementUpdateOne) SetChannel(s string) *AnnouncementUpdateOne {
	auo.mutation.SetChannel(s)
	return auo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableChannel(s *string) *AnnouncementUpdateOne {
	if s != nil {
		auo.SetChannel(*s)
	}
	return auo
}

// ClearChannel clears the value of the "channel" field.
func (auo *AnnouncementUpdateOne) ClearChannel() *AnnouncementUpdateOne {
	auo.mutation.ClearChannel()
	return auo
}

// SetStartAt sets the "start_at" field.
func (auo *AnnouncementUpdateOne) SetStartAt(u uint32) *AnnouncementUpdateOne {
	auo.mutation.ResetStartAt()
	auo.mutation.SetStartAt(u)
	return auo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableStartAt(u *uint32) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetStartAt(*u)
	}
	return auo
}

// AddStartAt adds u to the "start_at" field.
func (auo *AnnouncementUpdateOne) AddStartAt(u int32) *AnnouncementUpdateOne {
	auo.mutation.AddStartAt(u)
	return auo
}

// ClearStartAt clears the value of the "start_at" field.
func (auo *AnnouncementUpdateOne) ClearStartAt() *AnnouncementUpdateOne {
	auo.mutation.ClearStartAt()
	return auo
}

// SetEndAt sets the "end_at" field.
func (auo *AnnouncementUpdateOne) SetEndAt(u uint32) *AnnouncementUpdateOne {
	auo.mutation.ResetEndAt()
	auo.mutation.SetEndAt(u)
	return auo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableEndAt(u *uint32) *AnnouncementUpdateOne {
	if u != nil {
		auo.SetEndAt(*u)
	}
	return auo
}

// AddEndAt adds u to the "end_at" field.
func (auo *AnnouncementUpdateOne) AddEndAt(u int32) *AnnouncementUpdateOne {
	auo.mutation.AddEndAt(u)
	return auo
}

// ClearEndAt clears the value of the "end_at" field.
func (auo *AnnouncementUpdateOne) ClearEndAt() *AnnouncementUpdateOne {
	auo.mutation.ClearEndAt()
	return auo
}

// SetType sets the "type" field.
func (auo *AnnouncementUpdateOne) SetType(s string) *AnnouncementUpdateOne {
	auo.mutation.SetType(s)
	return auo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (auo *AnnouncementUpdateOne) SetNillableType(s *string) *AnnouncementUpdateOne {
	if s != nil {
		auo.SetType(*s)
	}
	return auo
}

// ClearType clears the value of the "type" field.
func (auo *AnnouncementUpdateOne) ClearType() *AnnouncementUpdateOne {
	auo.mutation.ClearType()
	return auo
}

// Mutation returns the AnnouncementMutation object of the builder.
func (auo *AnnouncementUpdateOne) Mutation() *AnnouncementMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnnouncementUpdateOne) Select(field string, fields ...string) *AnnouncementUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Announcement entity.
func (auo *AnnouncementUpdateOne) Save(ctx context.Context) (*Announcement, error) {
	var (
		err  error
		node *Announcement
	)
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Announcement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AnnouncementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnnouncementUpdateOne) SaveX(ctx context.Context) *Announcement {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnnouncementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnnouncementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AnnouncementUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if announcement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized announcement.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := announcement.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AnnouncementUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AnnouncementUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AnnouncementUpdateOne) sqlSave(ctx context.Context) (_node *Announcement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: announcement.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Announcement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, announcement.FieldID)
		for _, f := range fields {
			if !announcement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != announcement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldEntID,
		})
	}
	if value, ok := auo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldAppID,
		})
	}
	if auo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: announcement.FieldAppID,
		})
	}
	if value, ok := auo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: announcement.FieldLangID,
		})
	}
	if auo.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: announcement.FieldLangID,
		})
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldTitle,
		})
	}
	if auo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldTitle,
		})
	}
	if value, ok := auo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldContent,
		})
	}
	if auo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldContent,
		})
	}
	if value, ok := auo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldChannel,
		})
	}
	if auo.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldChannel,
		})
	}
	if value, ok := auo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldStartAt,
		})
	}
	if value, ok := auo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldStartAt,
		})
	}
	if auo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: announcement.FieldStartAt,
		})
	}
	if value, ok := auo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldEndAt,
		})
	}
	if value, ok := auo.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: announcement.FieldEndAt,
		})
	}
	if auo.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: announcement.FieldEndAt,
		})
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: announcement.FieldType,
		})
	}
	if auo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: announcement.FieldType,
		})
	}
	_spec.Modifiers = auo.modifiers
	_node = &Announcement{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{announcement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
