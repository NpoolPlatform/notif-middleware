// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/emailtemplate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// EmailTemplateUpdate is the builder for updating EmailTemplate entities.
type EmailTemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *EmailTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EmailTemplateUpdate builder.
func (etu *EmailTemplateUpdate) Where(ps ...predicate.EmailTemplate) *EmailTemplateUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetCreatedAt sets the "created_at" field.
func (etu *EmailTemplateUpdate) SetCreatedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetCreatedAt()
	etu.mutation.SetCreatedAt(u)
	return etu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableCreatedAt(u *uint32) *EmailTemplateUpdate {
	if u != nil {
		etu.SetCreatedAt(*u)
	}
	return etu
}

// AddCreatedAt adds u to the "created_at" field.
func (etu *EmailTemplateUpdate) AddCreatedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddCreatedAt(u)
	return etu
}

// SetUpdatedAt sets the "updated_at" field.
func (etu *EmailTemplateUpdate) SetUpdatedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetUpdatedAt()
	etu.mutation.SetUpdatedAt(u)
	return etu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (etu *EmailTemplateUpdate) AddUpdatedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddUpdatedAt(u)
	return etu
}

// SetDeletedAt sets the "deleted_at" field.
func (etu *EmailTemplateUpdate) SetDeletedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetDeletedAt()
	etu.mutation.SetDeletedAt(u)
	return etu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableDeletedAt(u *uint32) *EmailTemplateUpdate {
	if u != nil {
		etu.SetDeletedAt(*u)
	}
	return etu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (etu *EmailTemplateUpdate) AddDeletedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddDeletedAt(u)
	return etu
}

// SetAppID sets the "app_id" field.
func (etu *EmailTemplateUpdate) SetAppID(u uuid.UUID) *EmailTemplateUpdate {
	etu.mutation.SetAppID(u)
	return etu
}

// SetLangID sets the "lang_id" field.
func (etu *EmailTemplateUpdate) SetLangID(u uuid.UUID) *EmailTemplateUpdate {
	etu.mutation.SetLangID(u)
	return etu
}

// SetDefaultToUsername sets the "default_to_username" field.
func (etu *EmailTemplateUpdate) SetDefaultToUsername(s string) *EmailTemplateUpdate {
	etu.mutation.SetDefaultToUsername(s)
	return etu
}

// SetUsedFor sets the "used_for" field.
func (etu *EmailTemplateUpdate) SetUsedFor(s string) *EmailTemplateUpdate {
	etu.mutation.SetUsedFor(s)
	return etu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableUsedFor(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetUsedFor(*s)
	}
	return etu
}

// ClearUsedFor clears the value of the "used_for" field.
func (etu *EmailTemplateUpdate) ClearUsedFor() *EmailTemplateUpdate {
	etu.mutation.ClearUsedFor()
	return etu
}

// SetSender sets the "sender" field.
func (etu *EmailTemplateUpdate) SetSender(s string) *EmailTemplateUpdate {
	etu.mutation.SetSender(s)
	return etu
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableSender(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetSender(*s)
	}
	return etu
}

// ClearSender clears the value of the "sender" field.
func (etu *EmailTemplateUpdate) ClearSender() *EmailTemplateUpdate {
	etu.mutation.ClearSender()
	return etu
}

// SetReplyTos sets the "reply_tos" field.
func (etu *EmailTemplateUpdate) SetReplyTos(s []string) *EmailTemplateUpdate {
	etu.mutation.SetReplyTos(s)
	return etu
}

// ClearReplyTos clears the value of the "reply_tos" field.
func (etu *EmailTemplateUpdate) ClearReplyTos() *EmailTemplateUpdate {
	etu.mutation.ClearReplyTos()
	return etu
}

// SetCcTos sets the "cc_tos" field.
func (etu *EmailTemplateUpdate) SetCcTos(s []string) *EmailTemplateUpdate {
	etu.mutation.SetCcTos(s)
	return etu
}

// ClearCcTos clears the value of the "cc_tos" field.
func (etu *EmailTemplateUpdate) ClearCcTos() *EmailTemplateUpdate {
	etu.mutation.ClearCcTos()
	return etu
}

// SetSubject sets the "subject" field.
func (etu *EmailTemplateUpdate) SetSubject(s string) *EmailTemplateUpdate {
	etu.mutation.SetSubject(s)
	return etu
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableSubject(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetSubject(*s)
	}
	return etu
}

// ClearSubject clears the value of the "subject" field.
func (etu *EmailTemplateUpdate) ClearSubject() *EmailTemplateUpdate {
	etu.mutation.ClearSubject()
	return etu
}

// SetBody sets the "body" field.
func (etu *EmailTemplateUpdate) SetBody(s string) *EmailTemplateUpdate {
	etu.mutation.SetBody(s)
	return etu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableBody(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetBody(*s)
	}
	return etu
}

// ClearBody clears the value of the "body" field.
func (etu *EmailTemplateUpdate) ClearBody() *EmailTemplateUpdate {
	etu.mutation.ClearBody()
	return etu
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etu *EmailTemplateUpdate) Mutation() *EmailTemplateMutation {
	return etu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EmailTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := etu.defaults(); err != nil {
		return 0, err
	}
	if len(etu.hooks) == 0 {
		affected, err = etu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			etu.mutation = mutation
			affected, err = etu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(etu.hooks) - 1; i >= 0; i-- {
			if etu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EmailTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EmailTemplateUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EmailTemplateUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etu *EmailTemplateUpdate) defaults() error {
	if _, ok := etu.mutation.UpdatedAt(); !ok {
		if emailtemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized emailtemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := emailtemplate.UpdateDefaultUpdatedAt()
		etu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etu *EmailTemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmailTemplateUpdate {
	etu.modifiers = append(etu.modifiers, modifiers...)
	return etu
}

func (etu *EmailTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailtemplate.Table,
			Columns: emailtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailtemplate.FieldID,
			},
		},
	}
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldCreatedAt,
		})
	}
	if value, ok := etu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldCreatedAt,
		})
	}
	if value, ok := etu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldUpdatedAt,
		})
	}
	if value, ok := etu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldUpdatedAt,
		})
	}
	if value, ok := etu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldDeletedAt,
		})
	}
	if value, ok := etu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldDeletedAt,
		})
	}
	if value, ok := etu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: emailtemplate.FieldAppID,
		})
	}
	if value, ok := etu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: emailtemplate.FieldLangID,
		})
	}
	if value, ok := etu.mutation.DefaultToUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldDefaultToUsername,
		})
	}
	if value, ok := etu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldUsedFor,
		})
	}
	if etu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldUsedFor,
		})
	}
	if value, ok := etu.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldSender,
		})
	}
	if etu.mutation.SenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldSender,
		})
	}
	if value, ok := etu.mutation.ReplyTos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: emailtemplate.FieldReplyTos,
		})
	}
	if etu.mutation.ReplyTosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: emailtemplate.FieldReplyTos,
		})
	}
	if value, ok := etu.mutation.CcTos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: emailtemplate.FieldCcTos,
		})
	}
	if etu.mutation.CcTosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: emailtemplate.FieldCcTos,
		})
	}
	if value, ok := etu.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldSubject,
		})
	}
	if etu.mutation.SubjectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldSubject,
		})
	}
	if value, ok := etu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldBody,
		})
	}
	if etu.mutation.BodyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldBody,
		})
	}
	_spec.Modifiers = etu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EmailTemplateUpdateOne is the builder for updating a single EmailTemplate entity.
type EmailTemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EmailTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (etuo *EmailTemplateUpdateOne) SetCreatedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetCreatedAt()
	etuo.mutation.SetCreatedAt(u)
	return etuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableCreatedAt(u *uint32) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetCreatedAt(*u)
	}
	return etuo
}

// AddCreatedAt adds u to the "created_at" field.
func (etuo *EmailTemplateUpdateOne) AddCreatedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddCreatedAt(u)
	return etuo
}

// SetUpdatedAt sets the "updated_at" field.
func (etuo *EmailTemplateUpdateOne) SetUpdatedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetUpdatedAt()
	etuo.mutation.SetUpdatedAt(u)
	return etuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (etuo *EmailTemplateUpdateOne) AddUpdatedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddUpdatedAt(u)
	return etuo
}

// SetDeletedAt sets the "deleted_at" field.
func (etuo *EmailTemplateUpdateOne) SetDeletedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetDeletedAt()
	etuo.mutation.SetDeletedAt(u)
	return etuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableDeletedAt(u *uint32) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetDeletedAt(*u)
	}
	return etuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (etuo *EmailTemplateUpdateOne) AddDeletedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddDeletedAt(u)
	return etuo
}

// SetAppID sets the "app_id" field.
func (etuo *EmailTemplateUpdateOne) SetAppID(u uuid.UUID) *EmailTemplateUpdateOne {
	etuo.mutation.SetAppID(u)
	return etuo
}

// SetLangID sets the "lang_id" field.
func (etuo *EmailTemplateUpdateOne) SetLangID(u uuid.UUID) *EmailTemplateUpdateOne {
	etuo.mutation.SetLangID(u)
	return etuo
}

// SetDefaultToUsername sets the "default_to_username" field.
func (etuo *EmailTemplateUpdateOne) SetDefaultToUsername(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetDefaultToUsername(s)
	return etuo
}

// SetUsedFor sets the "used_for" field.
func (etuo *EmailTemplateUpdateOne) SetUsedFor(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetUsedFor(s)
	return etuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableUsedFor(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetUsedFor(*s)
	}
	return etuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (etuo *EmailTemplateUpdateOne) ClearUsedFor() *EmailTemplateUpdateOne {
	etuo.mutation.ClearUsedFor()
	return etuo
}

// SetSender sets the "sender" field.
func (etuo *EmailTemplateUpdateOne) SetSender(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetSender(s)
	return etuo
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableSender(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetSender(*s)
	}
	return etuo
}

// ClearSender clears the value of the "sender" field.
func (etuo *EmailTemplateUpdateOne) ClearSender() *EmailTemplateUpdateOne {
	etuo.mutation.ClearSender()
	return etuo
}

// SetReplyTos sets the "reply_tos" field.
func (etuo *EmailTemplateUpdateOne) SetReplyTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.SetReplyTos(s)
	return etuo
}

// ClearReplyTos clears the value of the "reply_tos" field.
func (etuo *EmailTemplateUpdateOne) ClearReplyTos() *EmailTemplateUpdateOne {
	etuo.mutation.ClearReplyTos()
	return etuo
}

// SetCcTos sets the "cc_tos" field.
func (etuo *EmailTemplateUpdateOne) SetCcTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.SetCcTos(s)
	return etuo
}

// ClearCcTos clears the value of the "cc_tos" field.
func (etuo *EmailTemplateUpdateOne) ClearCcTos() *EmailTemplateUpdateOne {
	etuo.mutation.ClearCcTos()
	return etuo
}

// SetSubject sets the "subject" field.
func (etuo *EmailTemplateUpdateOne) SetSubject(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetSubject(s)
	return etuo
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableSubject(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetSubject(*s)
	}
	return etuo
}

// ClearSubject clears the value of the "subject" field.
func (etuo *EmailTemplateUpdateOne) ClearSubject() *EmailTemplateUpdateOne {
	etuo.mutation.ClearSubject()
	return etuo
}

// SetBody sets the "body" field.
func (etuo *EmailTemplateUpdateOne) SetBody(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetBody(s)
	return etuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableBody(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetBody(*s)
	}
	return etuo
}

// ClearBody clears the value of the "body" field.
func (etuo *EmailTemplateUpdateOne) ClearBody() *EmailTemplateUpdateOne {
	etuo.mutation.ClearBody()
	return etuo
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etuo *EmailTemplateUpdateOne) Mutation() *EmailTemplateMutation {
	return etuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EmailTemplateUpdateOne) Select(field string, fields ...string) *EmailTemplateUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EmailTemplate entity.
func (etuo *EmailTemplateUpdateOne) Save(ctx context.Context) (*EmailTemplate, error) {
	var (
		err  error
		node *EmailTemplate
	)
	if err := etuo.defaults(); err != nil {
		return nil, err
	}
	if len(etuo.hooks) == 0 {
		node, err = etuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			etuo.mutation = mutation
			node, err = etuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(etuo.hooks) - 1; i >= 0; i-- {
			if etuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, etuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EmailTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EmailTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EmailTemplateUpdateOne) SaveX(ctx context.Context) *EmailTemplate {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EmailTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EmailTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etuo *EmailTemplateUpdateOne) defaults() error {
	if _, ok := etuo.mutation.UpdatedAt(); !ok {
		if emailtemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized emailtemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := emailtemplate.UpdateDefaultUpdatedAt()
		etuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etuo *EmailTemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmailTemplateUpdateOne {
	etuo.modifiers = append(etuo.modifiers, modifiers...)
	return etuo
}

func (etuo *EmailTemplateUpdateOne) sqlSave(ctx context.Context) (_node *EmailTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailtemplate.Table,
			Columns: emailtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailtemplate.FieldID,
			},
		},
	}
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailtemplate.FieldID)
		for _, f := range fields {
			if !emailtemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emailtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldCreatedAt,
		})
	}
	if value, ok := etuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldCreatedAt,
		})
	}
	if value, ok := etuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldUpdatedAt,
		})
	}
	if value, ok := etuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldUpdatedAt,
		})
	}
	if value, ok := etuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldDeletedAt,
		})
	}
	if value, ok := etuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: emailtemplate.FieldDeletedAt,
		})
	}
	if value, ok := etuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: emailtemplate.FieldAppID,
		})
	}
	if value, ok := etuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: emailtemplate.FieldLangID,
		})
	}
	if value, ok := etuo.mutation.DefaultToUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldDefaultToUsername,
		})
	}
	if value, ok := etuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldUsedFor,
		})
	}
	if etuo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldUsedFor,
		})
	}
	if value, ok := etuo.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldSender,
		})
	}
	if etuo.mutation.SenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldSender,
		})
	}
	if value, ok := etuo.mutation.ReplyTos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: emailtemplate.FieldReplyTos,
		})
	}
	if etuo.mutation.ReplyTosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: emailtemplate.FieldReplyTos,
		})
	}
	if value, ok := etuo.mutation.CcTos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: emailtemplate.FieldCcTos,
		})
	}
	if etuo.mutation.CcTosCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: emailtemplate.FieldCcTos,
		})
	}
	if value, ok := etuo.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldSubject,
		})
	}
	if etuo.mutation.SubjectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldSubject,
		})
	}
	if value, ok := etuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emailtemplate.FieldBody,
		})
	}
	if etuo.mutation.BodyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: emailtemplate.FieldBody,
		})
	}
	_spec.Modifiers = etuo.modifiers
	_node = &EmailTemplate{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}