// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/smstemplate"
)

// SMSTemplateQuery is the builder for querying SMSTemplate entities.
type SMSTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SMSTemplate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SMSTemplateQuery builder.
func (stq *SMSTemplateQuery) Where(ps ...predicate.SMSTemplate) *SMSTemplateQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *SMSTemplateQuery) Limit(limit int) *SMSTemplateQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *SMSTemplateQuery) Offset(offset int) *SMSTemplateQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SMSTemplateQuery) Unique(unique bool) *SMSTemplateQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *SMSTemplateQuery) Order(o ...OrderFunc) *SMSTemplateQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first SMSTemplate entity from the query.
// Returns a *NotFoundError when no SMSTemplate was found.
func (stq *SMSTemplateQuery) First(ctx context.Context) (*SMSTemplate, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{smstemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SMSTemplateQuery) FirstX(ctx context.Context) *SMSTemplate {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SMSTemplate ID from the query.
// Returns a *NotFoundError when no SMSTemplate ID was found.
func (stq *SMSTemplateQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{smstemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SMSTemplateQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SMSTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SMSTemplate entity is found.
// Returns a *NotFoundError when no SMSTemplate entities are found.
func (stq *SMSTemplateQuery) Only(ctx context.Context) (*SMSTemplate, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{smstemplate.Label}
	default:
		return nil, &NotSingularError{smstemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SMSTemplateQuery) OnlyX(ctx context.Context) *SMSTemplate {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SMSTemplate ID in the query.
// Returns a *NotSingularError when more than one SMSTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SMSTemplateQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{smstemplate.Label}
	default:
		err = &NotSingularError{smstemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SMSTemplateQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SMSTemplates.
func (stq *SMSTemplateQuery) All(ctx context.Context) ([]*SMSTemplate, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *SMSTemplateQuery) AllX(ctx context.Context) []*SMSTemplate {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SMSTemplate IDs.
func (stq *SMSTemplateQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := stq.Select(smstemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SMSTemplateQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SMSTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SMSTemplateQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SMSTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SMSTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SMSTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SMSTemplateQuery) Clone() *SMSTemplateQuery {
	if stq == nil {
		return nil
	}
	return &SMSTemplateQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.SMSTemplate{}, stq.predicates...),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SMSTemplate.Query().
//		GroupBy(smstemplate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (stq *SMSTemplateQuery) GroupBy(field string, fields ...string) *SMSTemplateGroupBy {
	grbuild := &SMSTemplateGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = smstemplate.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.SMSTemplate.Query().
//		Select(smstemplate.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (stq *SMSTemplateQuery) Select(fields ...string) *SMSTemplateSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &SMSTemplateSelect{SMSTemplateQuery: stq}
	selbuild.label = smstemplate.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

func (stq *SMSTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !smstemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	if smstemplate.Policy == nil {
		return errors.New("ent: uninitialized smstemplate.Policy (forgotten import ent/runtime?)")
	}
	if err := smstemplate.Policy.EvalQuery(ctx, stq); err != nil {
		return err
	}
	return nil
}

func (stq *SMSTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SMSTemplate, error) {
	var (
		nodes = []*SMSTemplate{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SMSTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SMSTemplate{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (stq *SMSTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SMSTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := stq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (stq *SMSTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smstemplate.Table,
			Columns: smstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: smstemplate.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smstemplate.FieldID)
		for i := range fields {
			if fields[i] != smstemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SMSTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(smstemplate.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = smstemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, m := range stq.modifiers {
		m(selector)
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (stq *SMSTemplateQuery) ForUpdate(opts ...sql.LockOption) *SMSTemplateQuery {
	if stq.driver.Dialect() == dialect.Postgres {
		stq.Unique(false)
	}
	stq.modifiers = append(stq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return stq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (stq *SMSTemplateQuery) ForShare(opts ...sql.LockOption) *SMSTemplateQuery {
	if stq.driver.Dialect() == dialect.Postgres {
		stq.Unique(false)
	}
	stq.modifiers = append(stq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return stq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (stq *SMSTemplateQuery) Modify(modifiers ...func(s *sql.Selector)) *SMSTemplateSelect {
	stq.modifiers = append(stq.modifiers, modifiers...)
	return stq.Select()
}

// SMSTemplateGroupBy is the group-by builder for SMSTemplate entities.
type SMSTemplateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SMSTemplateGroupBy) Aggregate(fns ...AggregateFunc) *SMSTemplateGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *SMSTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *SMSTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range stgb.fields {
		if !smstemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *SMSTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// SMSTemplateSelect is the builder for selecting fields of SMSTemplate entities.
type SMSTemplateSelect struct {
	*SMSTemplateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SMSTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.SMSTemplateQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *SMSTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sts *SMSTemplateSelect) Modify(modifiers ...func(s *sql.Selector)) *SMSTemplateSelect {
	sts.modifiers = append(sts.modifiers, modifiers...)
	return sts
}
