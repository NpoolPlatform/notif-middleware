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
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notifuser"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// NotifUserQuery is the builder for querying NotifUser entities.
type NotifUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NotifUser
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotifUserQuery builder.
func (nuq *NotifUserQuery) Where(ps ...predicate.NotifUser) *NotifUserQuery {
	nuq.predicates = append(nuq.predicates, ps...)
	return nuq
}

// Limit adds a limit step to the query.
func (nuq *NotifUserQuery) Limit(limit int) *NotifUserQuery {
	nuq.limit = &limit
	return nuq
}

// Offset adds an offset step to the query.
func (nuq *NotifUserQuery) Offset(offset int) *NotifUserQuery {
	nuq.offset = &offset
	return nuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nuq *NotifUserQuery) Unique(unique bool) *NotifUserQuery {
	nuq.unique = &unique
	return nuq
}

// Order adds an order step to the query.
func (nuq *NotifUserQuery) Order(o ...OrderFunc) *NotifUserQuery {
	nuq.order = append(nuq.order, o...)
	return nuq
}

// First returns the first NotifUser entity from the query.
// Returns a *NotFoundError when no NotifUser was found.
func (nuq *NotifUserQuery) First(ctx context.Context) (*NotifUser, error) {
	nodes, err := nuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notifuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nuq *NotifUserQuery) FirstX(ctx context.Context) *NotifUser {
	node, err := nuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotifUser ID from the query.
// Returns a *NotFoundError when no NotifUser ID was found.
func (nuq *NotifUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notifuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nuq *NotifUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotifUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotifUser entity is found.
// Returns a *NotFoundError when no NotifUser entities are found.
func (nuq *NotifUserQuery) Only(ctx context.Context) (*NotifUser, error) {
	nodes, err := nuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notifuser.Label}
	default:
		return nil, &NotSingularError{notifuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nuq *NotifUserQuery) OnlyX(ctx context.Context) *NotifUser {
	node, err := nuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotifUser ID in the query.
// Returns a *NotSingularError when more than one NotifUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (nuq *NotifUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notifuser.Label}
	default:
		err = &NotSingularError{notifuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nuq *NotifUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotifUsers.
func (nuq *NotifUserQuery) All(ctx context.Context) ([]*NotifUser, error) {
	if err := nuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nuq *NotifUserQuery) AllX(ctx context.Context) []*NotifUser {
	nodes, err := nuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotifUser IDs.
func (nuq *NotifUserQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := nuq.Select(notifuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nuq *NotifUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nuq *NotifUserQuery) Count(ctx context.Context) (int, error) {
	if err := nuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nuq *NotifUserQuery) CountX(ctx context.Context) int {
	count, err := nuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nuq *NotifUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := nuq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nuq *NotifUserQuery) ExistX(ctx context.Context) bool {
	exist, err := nuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotifUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nuq *NotifUserQuery) Clone() *NotifUserQuery {
	if nuq == nil {
		return nil
	}
	return &NotifUserQuery{
		config:     nuq.config,
		limit:      nuq.limit,
		offset:     nuq.offset,
		order:      append([]OrderFunc{}, nuq.order...),
		predicates: append([]predicate.NotifUser{}, nuq.predicates...),
		// clone intermediate query.
		sql:    nuq.sql.Clone(),
		path:   nuq.path,
		unique: nuq.unique,
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
//	client.NotifUser.Query().
//		GroupBy(notifuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nuq *NotifUserQuery) GroupBy(field string, fields ...string) *NotifUserGroupBy {
	grbuild := &NotifUserGroupBy{config: nuq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nuq.sqlQuery(ctx), nil
	}
	grbuild.label = notifuser.Label
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
//	client.NotifUser.Query().
//		Select(notifuser.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (nuq *NotifUserQuery) Select(fields ...string) *NotifUserSelect {
	nuq.fields = append(nuq.fields, fields...)
	selbuild := &NotifUserSelect{NotifUserQuery: nuq}
	selbuild.label = notifuser.Label
	selbuild.flds, selbuild.scan = &nuq.fields, selbuild.Scan
	return selbuild
}

func (nuq *NotifUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nuq.fields {
		if !notifuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nuq.path != nil {
		prev, err := nuq.path(ctx)
		if err != nil {
			return err
		}
		nuq.sql = prev
	}
	if notifuser.Policy == nil {
		return errors.New("ent: uninitialized notifuser.Policy (forgotten import ent/runtime?)")
	}
	if err := notifuser.Policy.EvalQuery(ctx, nuq); err != nil {
		return err
	}
	return nil
}

func (nuq *NotifUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotifUser, error) {
	var (
		nodes = []*NotifUser{}
		_spec = nuq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*NotifUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &NotifUser{config: nuq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(nuq.modifiers) > 0 {
		_spec.Modifiers = nuq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (nuq *NotifUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nuq.querySpec()
	if len(nuq.modifiers) > 0 {
		_spec.Modifiers = nuq.modifiers
	}
	_spec.Node.Columns = nuq.fields
	if len(nuq.fields) > 0 {
		_spec.Unique = nuq.unique != nil && *nuq.unique
	}
	return sqlgraph.CountNodes(ctx, nuq.driver, _spec)
}

func (nuq *NotifUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nuq *NotifUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notifuser.Table,
			Columns: notifuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: notifuser.FieldID,
			},
		},
		From:   nuq.sql,
		Unique: true,
	}
	if unique := nuq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nuq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifuser.FieldID)
		for i := range fields {
			if fields[i] != notifuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nuq *NotifUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nuq.driver.Dialect())
	t1 := builder.Table(notifuser.Table)
	columns := nuq.fields
	if len(columns) == 0 {
		columns = notifuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nuq.sql != nil {
		selector = nuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nuq.unique != nil && *nuq.unique {
		selector.Distinct()
	}
	for _, m := range nuq.modifiers {
		m(selector)
	}
	for _, p := range nuq.predicates {
		p(selector)
	}
	for _, p := range nuq.order {
		p(selector)
	}
	if offset := nuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (nuq *NotifUserQuery) ForUpdate(opts ...sql.LockOption) *NotifUserQuery {
	if nuq.driver.Dialect() == dialect.Postgres {
		nuq.Unique(false)
	}
	nuq.modifiers = append(nuq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return nuq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (nuq *NotifUserQuery) ForShare(opts ...sql.LockOption) *NotifUserQuery {
	if nuq.driver.Dialect() == dialect.Postgres {
		nuq.Unique(false)
	}
	nuq.modifiers = append(nuq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return nuq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nuq *NotifUserQuery) Modify(modifiers ...func(s *sql.Selector)) *NotifUserSelect {
	nuq.modifiers = append(nuq.modifiers, modifiers...)
	return nuq.Select()
}

// NotifUserGroupBy is the group-by builder for NotifUser entities.
type NotifUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nugb *NotifUserGroupBy) Aggregate(fns ...AggregateFunc) *NotifUserGroupBy {
	nugb.fns = append(nugb.fns, fns...)
	return nugb
}

// Scan applies the group-by query and scans the result into the given value.
func (nugb *NotifUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nugb.path(ctx)
	if err != nil {
		return err
	}
	nugb.sql = query
	return nugb.sqlScan(ctx, v)
}

func (nugb *NotifUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nugb.fields {
		if !notifuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nugb *NotifUserGroupBy) sqlQuery() *sql.Selector {
	selector := nugb.sql.Select()
	aggregation := make([]string, 0, len(nugb.fns))
	for _, fn := range nugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nugb.fields)+len(nugb.fns))
		for _, f := range nugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nugb.fields...)...)
}

// NotifUserSelect is the builder for selecting fields of NotifUser entities.
type NotifUserSelect struct {
	*NotifUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nus *NotifUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nus.prepareQuery(ctx); err != nil {
		return err
	}
	nus.sql = nus.NotifUserQuery.sqlQuery(ctx)
	return nus.sqlScan(ctx, v)
}

func (nus *NotifUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nus.sql.Query()
	if err := nus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nus *NotifUserSelect) Modify(modifiers ...func(s *sql.Selector)) *NotifUserSelect {
	nus.modifiers = append(nus.modifiers, modifiers...)
	return nus
}