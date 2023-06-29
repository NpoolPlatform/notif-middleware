// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notif"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// NotifUpdate is the builder for updating Notif entities.
type NotifUpdate struct {
	config
	hooks     []Hook
	mutation  *NotifMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NotifUpdate builder.
func (nu *NotifUpdate) Where(ps ...predicate.Notif) *NotifUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NotifUpdate) SetCreatedAt(u uint32) *NotifUpdate {
	nu.mutation.ResetCreatedAt()
	nu.mutation.SetCreatedAt(u)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableCreatedAt(u *uint32) *NotifUpdate {
	if u != nil {
		nu.SetCreatedAt(*u)
	}
	return nu
}

// AddCreatedAt adds u to the "created_at" field.
func (nu *NotifUpdate) AddCreatedAt(u int32) *NotifUpdate {
	nu.mutation.AddCreatedAt(u)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotifUpdate) SetUpdatedAt(u uint32) *NotifUpdate {
	nu.mutation.ResetUpdatedAt()
	nu.mutation.SetUpdatedAt(u)
	return nu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (nu *NotifUpdate) AddUpdatedAt(u int32) *NotifUpdate {
	nu.mutation.AddUpdatedAt(u)
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NotifUpdate) SetDeletedAt(u uint32) *NotifUpdate {
	nu.mutation.ResetDeletedAt()
	nu.mutation.SetDeletedAt(u)
	return nu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableDeletedAt(u *uint32) *NotifUpdate {
	if u != nil {
		nu.SetDeletedAt(*u)
	}
	return nu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (nu *NotifUpdate) AddDeletedAt(u int32) *NotifUpdate {
	nu.mutation.AddDeletedAt(u)
	return nu
}

// SetAppID sets the "app_id" field.
func (nu *NotifUpdate) SetAppID(u uuid.UUID) *NotifUpdate {
	nu.mutation.SetAppID(u)
	return nu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableAppID(u *uuid.UUID) *NotifUpdate {
	if u != nil {
		nu.SetAppID(*u)
	}
	return nu
}

// ClearAppID clears the value of the "app_id" field.
func (nu *NotifUpdate) ClearAppID() *NotifUpdate {
	nu.mutation.ClearAppID()
	return nu
}

// SetUserID sets the "user_id" field.
func (nu *NotifUpdate) SetUserID(u uuid.UUID) *NotifUpdate {
	nu.mutation.SetUserID(u)
	return nu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableUserID(u *uuid.UUID) *NotifUpdate {
	if u != nil {
		nu.SetUserID(*u)
	}
	return nu
}

// ClearUserID clears the value of the "user_id" field.
func (nu *NotifUpdate) ClearUserID() *NotifUpdate {
	nu.mutation.ClearUserID()
	return nu
}

// SetNotified sets the "notified" field.
func (nu *NotifUpdate) SetNotified(b bool) *NotifUpdate {
	nu.mutation.SetNotified(b)
	return nu
}

// SetNillableNotified sets the "notified" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableNotified(b *bool) *NotifUpdate {
	if b != nil {
		nu.SetNotified(*b)
	}
	return nu
}

// ClearNotified clears the value of the "notified" field.
func (nu *NotifUpdate) ClearNotified() *NotifUpdate {
	nu.mutation.ClearNotified()
	return nu
}

// SetLangID sets the "lang_id" field.
func (nu *NotifUpdate) SetLangID(u uuid.UUID) *NotifUpdate {
	nu.mutation.SetLangID(u)
	return nu
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableLangID(u *uuid.UUID) *NotifUpdate {
	if u != nil {
		nu.SetLangID(*u)
	}
	return nu
}

// ClearLangID clears the value of the "lang_id" field.
func (nu *NotifUpdate) ClearLangID() *NotifUpdate {
	nu.mutation.ClearLangID()
	return nu
}

// SetEventID sets the "event_id" field.
func (nu *NotifUpdate) SetEventID(u uuid.UUID) *NotifUpdate {
	nu.mutation.SetEventID(u)
	return nu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableEventID(u *uuid.UUID) *NotifUpdate {
	if u != nil {
		nu.SetEventID(*u)
	}
	return nu
}

// ClearEventID clears the value of the "event_id" field.
func (nu *NotifUpdate) ClearEventID() *NotifUpdate {
	nu.mutation.ClearEventID()
	return nu
}

// SetEventType sets the "event_type" field.
func (nu *NotifUpdate) SetEventType(s string) *NotifUpdate {
	nu.mutation.SetEventType(s)
	return nu
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableEventType(s *string) *NotifUpdate {
	if s != nil {
		nu.SetEventType(*s)
	}
	return nu
}

// ClearEventType clears the value of the "event_type" field.
func (nu *NotifUpdate) ClearEventType() *NotifUpdate {
	nu.mutation.ClearEventType()
	return nu
}

// SetUseTemplate sets the "use_template" field.
func (nu *NotifUpdate) SetUseTemplate(b bool) *NotifUpdate {
	nu.mutation.SetUseTemplate(b)
	return nu
}

// SetNillableUseTemplate sets the "use_template" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableUseTemplate(b *bool) *NotifUpdate {
	if b != nil {
		nu.SetUseTemplate(*b)
	}
	return nu
}

// ClearUseTemplate clears the value of the "use_template" field.
func (nu *NotifUpdate) ClearUseTemplate() *NotifUpdate {
	nu.mutation.ClearUseTemplate()
	return nu
}

// SetTitle sets the "title" field.
func (nu *NotifUpdate) SetTitle(s string) *NotifUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableTitle(s *string) *NotifUpdate {
	if s != nil {
		nu.SetTitle(*s)
	}
	return nu
}

// ClearTitle clears the value of the "title" field.
func (nu *NotifUpdate) ClearTitle() *NotifUpdate {
	nu.mutation.ClearTitle()
	return nu
}

// SetContent sets the "content" field.
func (nu *NotifUpdate) SetContent(s string) *NotifUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableContent(s *string) *NotifUpdate {
	if s != nil {
		nu.SetContent(*s)
	}
	return nu
}

// ClearContent clears the value of the "content" field.
func (nu *NotifUpdate) ClearContent() *NotifUpdate {
	nu.mutation.ClearContent()
	return nu
}

// SetChannel sets the "channel" field.
func (nu *NotifUpdate) SetChannel(s string) *NotifUpdate {
	nu.mutation.SetChannel(s)
	return nu
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableChannel(s *string) *NotifUpdate {
	if s != nil {
		nu.SetChannel(*s)
	}
	return nu
}

// ClearChannel clears the value of the "channel" field.
func (nu *NotifUpdate) ClearChannel() *NotifUpdate {
	nu.mutation.ClearChannel()
	return nu
}

// SetExtra sets the "extra" field.
func (nu *NotifUpdate) SetExtra(s string) *NotifUpdate {
	nu.mutation.SetExtra(s)
	return nu
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableExtra(s *string) *NotifUpdate {
	if s != nil {
		nu.SetExtra(*s)
	}
	return nu
}

// ClearExtra clears the value of the "extra" field.
func (nu *NotifUpdate) ClearExtra() *NotifUpdate {
	nu.mutation.ClearExtra()
	return nu
}

// SetType sets the "type" field.
func (nu *NotifUpdate) SetType(s string) *NotifUpdate {
	nu.mutation.SetType(s)
	return nu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nu *NotifUpdate) SetNillableType(s *string) *NotifUpdate {
	if s != nil {
		nu.SetType(*s)
	}
	return nu
}

// ClearType clears the value of the "type" field.
func (nu *NotifUpdate) ClearType() *NotifUpdate {
	nu.mutation.ClearType()
	return nu
}

// Mutation returns the NotifMutation object of the builder.
func (nu *NotifUpdate) Mutation() *NotifMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotifUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := nu.defaults(); err != nil {
		return 0, err
	}
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotifUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotifUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotifUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NotifUpdate) defaults() error {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		if notif.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized notif.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := notif.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nu *NotifUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotifUpdate {
	nu.modifiers = append(nu.modifiers, modifiers...)
	return nu
}

func (nu *NotifUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notif.Table,
			Columns: notif.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notif.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldCreatedAt,
		})
	}
	if value, ok := nu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldCreatedAt,
		})
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldUpdatedAt,
		})
	}
	if value, ok := nu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldUpdatedAt,
		})
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldDeletedAt,
		})
	}
	if value, ok := nu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldDeletedAt,
		})
	}
	if value, ok := nu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldAppID,
		})
	}
	if nu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldAppID,
		})
	}
	if value, ok := nu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldUserID,
		})
	}
	if nu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldUserID,
		})
	}
	if value, ok := nu.mutation.Notified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notif.FieldNotified,
		})
	}
	if nu.mutation.NotifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: notif.FieldNotified,
		})
	}
	if value, ok := nu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldLangID,
		})
	}
	if nu.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldLangID,
		})
	}
	if value, ok := nu.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldEventID,
		})
	}
	if nu.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldEventID,
		})
	}
	if value, ok := nu.mutation.EventType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldEventType,
		})
	}
	if nu.mutation.EventTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldEventType,
		})
	}
	if value, ok := nu.mutation.UseTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notif.FieldUseTemplate,
		})
	}
	if nu.mutation.UseTemplateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: notif.FieldUseTemplate,
		})
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldTitle,
		})
	}
	if nu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldTitle,
		})
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldContent,
		})
	}
	if nu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldContent,
		})
	}
	if value, ok := nu.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldChannel,
		})
	}
	if nu.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldChannel,
		})
	}
	if value, ok := nu.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldExtra,
		})
	}
	if nu.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldExtra,
		})
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldType,
		})
	}
	if nu.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldType,
		})
	}
	_spec.Modifiers = nu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notif.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// NotifUpdateOne is the builder for updating a single Notif entity.
type NotifUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NotifMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NotifUpdateOne) SetCreatedAt(u uint32) *NotifUpdateOne {
	nuo.mutation.ResetCreatedAt()
	nuo.mutation.SetCreatedAt(u)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableCreatedAt(u *uint32) *NotifUpdateOne {
	if u != nil {
		nuo.SetCreatedAt(*u)
	}
	return nuo
}

// AddCreatedAt adds u to the "created_at" field.
func (nuo *NotifUpdateOne) AddCreatedAt(u int32) *NotifUpdateOne {
	nuo.mutation.AddCreatedAt(u)
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotifUpdateOne) SetUpdatedAt(u uint32) *NotifUpdateOne {
	nuo.mutation.ResetUpdatedAt()
	nuo.mutation.SetUpdatedAt(u)
	return nuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (nuo *NotifUpdateOne) AddUpdatedAt(u int32) *NotifUpdateOne {
	nuo.mutation.AddUpdatedAt(u)
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NotifUpdateOne) SetDeletedAt(u uint32) *NotifUpdateOne {
	nuo.mutation.ResetDeletedAt()
	nuo.mutation.SetDeletedAt(u)
	return nuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableDeletedAt(u *uint32) *NotifUpdateOne {
	if u != nil {
		nuo.SetDeletedAt(*u)
	}
	return nuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (nuo *NotifUpdateOne) AddDeletedAt(u int32) *NotifUpdateOne {
	nuo.mutation.AddDeletedAt(u)
	return nuo
}

// SetAppID sets the "app_id" field.
func (nuo *NotifUpdateOne) SetAppID(u uuid.UUID) *NotifUpdateOne {
	nuo.mutation.SetAppID(u)
	return nuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableAppID(u *uuid.UUID) *NotifUpdateOne {
	if u != nil {
		nuo.SetAppID(*u)
	}
	return nuo
}

// ClearAppID clears the value of the "app_id" field.
func (nuo *NotifUpdateOne) ClearAppID() *NotifUpdateOne {
	nuo.mutation.ClearAppID()
	return nuo
}

// SetUserID sets the "user_id" field.
func (nuo *NotifUpdateOne) SetUserID(u uuid.UUID) *NotifUpdateOne {
	nuo.mutation.SetUserID(u)
	return nuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableUserID(u *uuid.UUID) *NotifUpdateOne {
	if u != nil {
		nuo.SetUserID(*u)
	}
	return nuo
}

// ClearUserID clears the value of the "user_id" field.
func (nuo *NotifUpdateOne) ClearUserID() *NotifUpdateOne {
	nuo.mutation.ClearUserID()
	return nuo
}

// SetNotified sets the "notified" field.
func (nuo *NotifUpdateOne) SetNotified(b bool) *NotifUpdateOne {
	nuo.mutation.SetNotified(b)
	return nuo
}

// SetNillableNotified sets the "notified" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableNotified(b *bool) *NotifUpdateOne {
	if b != nil {
		nuo.SetNotified(*b)
	}
	return nuo
}

// ClearNotified clears the value of the "notified" field.
func (nuo *NotifUpdateOne) ClearNotified() *NotifUpdateOne {
	nuo.mutation.ClearNotified()
	return nuo
}

// SetLangID sets the "lang_id" field.
func (nuo *NotifUpdateOne) SetLangID(u uuid.UUID) *NotifUpdateOne {
	nuo.mutation.SetLangID(u)
	return nuo
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableLangID(u *uuid.UUID) *NotifUpdateOne {
	if u != nil {
		nuo.SetLangID(*u)
	}
	return nuo
}

// ClearLangID clears the value of the "lang_id" field.
func (nuo *NotifUpdateOne) ClearLangID() *NotifUpdateOne {
	nuo.mutation.ClearLangID()
	return nuo
}

// SetEventID sets the "event_id" field.
func (nuo *NotifUpdateOne) SetEventID(u uuid.UUID) *NotifUpdateOne {
	nuo.mutation.SetEventID(u)
	return nuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableEventID(u *uuid.UUID) *NotifUpdateOne {
	if u != nil {
		nuo.SetEventID(*u)
	}
	return nuo
}

// ClearEventID clears the value of the "event_id" field.
func (nuo *NotifUpdateOne) ClearEventID() *NotifUpdateOne {
	nuo.mutation.ClearEventID()
	return nuo
}

// SetEventType sets the "event_type" field.
func (nuo *NotifUpdateOne) SetEventType(s string) *NotifUpdateOne {
	nuo.mutation.SetEventType(s)
	return nuo
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableEventType(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetEventType(*s)
	}
	return nuo
}

// ClearEventType clears the value of the "event_type" field.
func (nuo *NotifUpdateOne) ClearEventType() *NotifUpdateOne {
	nuo.mutation.ClearEventType()
	return nuo
}

// SetUseTemplate sets the "use_template" field.
func (nuo *NotifUpdateOne) SetUseTemplate(b bool) *NotifUpdateOne {
	nuo.mutation.SetUseTemplate(b)
	return nuo
}

// SetNillableUseTemplate sets the "use_template" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableUseTemplate(b *bool) *NotifUpdateOne {
	if b != nil {
		nuo.SetUseTemplate(*b)
	}
	return nuo
}

// ClearUseTemplate clears the value of the "use_template" field.
func (nuo *NotifUpdateOne) ClearUseTemplate() *NotifUpdateOne {
	nuo.mutation.ClearUseTemplate()
	return nuo
}

// SetTitle sets the "title" field.
func (nuo *NotifUpdateOne) SetTitle(s string) *NotifUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableTitle(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetTitle(*s)
	}
	return nuo
}

// ClearTitle clears the value of the "title" field.
func (nuo *NotifUpdateOne) ClearTitle() *NotifUpdateOne {
	nuo.mutation.ClearTitle()
	return nuo
}

// SetContent sets the "content" field.
func (nuo *NotifUpdateOne) SetContent(s string) *NotifUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableContent(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetContent(*s)
	}
	return nuo
}

// ClearContent clears the value of the "content" field.
func (nuo *NotifUpdateOne) ClearContent() *NotifUpdateOne {
	nuo.mutation.ClearContent()
	return nuo
}

// SetChannel sets the "channel" field.
func (nuo *NotifUpdateOne) SetChannel(s string) *NotifUpdateOne {
	nuo.mutation.SetChannel(s)
	return nuo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableChannel(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetChannel(*s)
	}
	return nuo
}

// ClearChannel clears the value of the "channel" field.
func (nuo *NotifUpdateOne) ClearChannel() *NotifUpdateOne {
	nuo.mutation.ClearChannel()
	return nuo
}

// SetExtra sets the "extra" field.
func (nuo *NotifUpdateOne) SetExtra(s string) *NotifUpdateOne {
	nuo.mutation.SetExtra(s)
	return nuo
}

// SetNillableExtra sets the "extra" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableExtra(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetExtra(*s)
	}
	return nuo
}

// ClearExtra clears the value of the "extra" field.
func (nuo *NotifUpdateOne) ClearExtra() *NotifUpdateOne {
	nuo.mutation.ClearExtra()
	return nuo
}

// SetType sets the "type" field.
func (nuo *NotifUpdateOne) SetType(s string) *NotifUpdateOne {
	nuo.mutation.SetType(s)
	return nuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nuo *NotifUpdateOne) SetNillableType(s *string) *NotifUpdateOne {
	if s != nil {
		nuo.SetType(*s)
	}
	return nuo
}

// ClearType clears the value of the "type" field.
func (nuo *NotifUpdateOne) ClearType() *NotifUpdateOne {
	nuo.mutation.ClearType()
	return nuo
}

// Mutation returns the NotifMutation object of the builder.
func (nuo *NotifUpdateOne) Mutation() *NotifMutation {
	return nuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotifUpdateOne) Select(field string, fields ...string) *NotifUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notif entity.
func (nuo *NotifUpdateOne) Save(ctx context.Context) (*Notif, error) {
	var (
		err  error
		node *Notif
	)
	if err := nuo.defaults(); err != nil {
		return nil, err
	}
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotifMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Notif)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NotifMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotifUpdateOne) SaveX(ctx context.Context) *Notif {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotifUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotifUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NotifUpdateOne) defaults() error {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		if notif.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized notif.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := notif.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nuo *NotifUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NotifUpdateOne {
	nuo.modifiers = append(nuo.modifiers, modifiers...)
	return nuo
}

func (nuo *NotifUpdateOne) sqlSave(ctx context.Context) (_node *Notif, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notif.Table,
			Columns: notif.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notif.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notif.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notif.FieldID)
		for _, f := range fields {
			if !notif.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notif.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldCreatedAt,
		})
	}
	if value, ok := nuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldCreatedAt,
		})
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldUpdatedAt,
		})
	}
	if value, ok := nuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldUpdatedAt,
		})
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldDeletedAt,
		})
	}
	if value, ok := nuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: notif.FieldDeletedAt,
		})
	}
	if value, ok := nuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldAppID,
		})
	}
	if nuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldAppID,
		})
	}
	if value, ok := nuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldUserID,
		})
	}
	if nuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldUserID,
		})
	}
	if value, ok := nuo.mutation.Notified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notif.FieldNotified,
		})
	}
	if nuo.mutation.NotifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: notif.FieldNotified,
		})
	}
	if value, ok := nuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldLangID,
		})
	}
	if nuo.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldLangID,
		})
	}
	if value, ok := nuo.mutation.EventID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: notif.FieldEventID,
		})
	}
	if nuo.mutation.EventIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: notif.FieldEventID,
		})
	}
	if value, ok := nuo.mutation.EventType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldEventType,
		})
	}
	if nuo.mutation.EventTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldEventType,
		})
	}
	if value, ok := nuo.mutation.UseTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notif.FieldUseTemplate,
		})
	}
	if nuo.mutation.UseTemplateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: notif.FieldUseTemplate,
		})
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldTitle,
		})
	}
	if nuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldTitle,
		})
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldContent,
		})
	}
	if nuo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldContent,
		})
	}
	if value, ok := nuo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldChannel,
		})
	}
	if nuo.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldChannel,
		})
	}
	if value, ok := nuo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldExtra,
		})
	}
	if nuo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldExtra,
		})
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notif.FieldType,
		})
	}
	if nuo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: notif.FieldType,
		})
	}
	_spec.Modifiers = nuo.modifiers
	_node = &Notif{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notif.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
