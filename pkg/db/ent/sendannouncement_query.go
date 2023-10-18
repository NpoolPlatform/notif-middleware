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
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/sendannouncement"
)

// SendAnnouncementQuery is the builder for querying SendAnnouncement entities.
type SendAnnouncementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SendAnnouncement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SendAnnouncementQuery builder.
func (saq *SendAnnouncementQuery) Where(ps ...predicate.SendAnnouncement) *SendAnnouncementQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit adds a limit step to the query.
func (saq *SendAnnouncementQuery) Limit(limit int) *SendAnnouncementQuery {
	saq.limit = &limit
	return saq
}

// Offset adds an offset step to the query.
func (saq *SendAnnouncementQuery) Offset(offset int) *SendAnnouncementQuery {
	saq.offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SendAnnouncementQuery) Unique(unique bool) *SendAnnouncementQuery {
	saq.unique = &unique
	return saq
}

// Order adds an order step to the query.
func (saq *SendAnnouncementQuery) Order(o ...OrderFunc) *SendAnnouncementQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// First returns the first SendAnnouncement entity from the query.
// Returns a *NotFoundError when no SendAnnouncement was found.
func (saq *SendAnnouncementQuery) First(ctx context.Context) (*SendAnnouncement, error) {
	nodes, err := saq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sendannouncement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SendAnnouncementQuery) FirstX(ctx context.Context) *SendAnnouncement {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SendAnnouncement ID from the query.
// Returns a *NotFoundError when no SendAnnouncement ID was found.
func (saq *SendAnnouncementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = saq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sendannouncement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SendAnnouncementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SendAnnouncement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SendAnnouncement entity is found.
// Returns a *NotFoundError when no SendAnnouncement entities are found.
func (saq *SendAnnouncementQuery) Only(ctx context.Context) (*SendAnnouncement, error) {
	nodes, err := saq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sendannouncement.Label}
	default:
		return nil, &NotSingularError{sendannouncement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SendAnnouncementQuery) OnlyX(ctx context.Context) *SendAnnouncement {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SendAnnouncement ID in the query.
// Returns a *NotSingularError when more than one SendAnnouncement ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SendAnnouncementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = saq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sendannouncement.Label}
	default:
		err = &NotSingularError{sendannouncement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SendAnnouncementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SendAnnouncements.
func (saq *SendAnnouncementQuery) All(ctx context.Context) ([]*SendAnnouncement, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return saq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (saq *SendAnnouncementQuery) AllX(ctx context.Context) []*SendAnnouncement {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SendAnnouncement IDs.
func (saq *SendAnnouncementQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := saq.Select(sendannouncement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SendAnnouncementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SendAnnouncementQuery) Count(ctx context.Context) (int, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return saq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SendAnnouncementQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SendAnnouncementQuery) Exist(ctx context.Context) (bool, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return saq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *SendAnnouncementQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SendAnnouncementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SendAnnouncementQuery) Clone() *SendAnnouncementQuery {
	if saq == nil {
		return nil
	}
	return &SendAnnouncementQuery{
		config:     saq.config,
		limit:      saq.limit,
		offset:     saq.offset,
		order:      append([]OrderFunc{}, saq.order...),
		predicates: append([]predicate.SendAnnouncement{}, saq.predicates...),
		// clone intermediate query.
		sql:    saq.sql.Clone(),
		path:   saq.path,
		unique: saq.unique,
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
//	client.SendAnnouncement.Query().
//		GroupBy(sendannouncement.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (saq *SendAnnouncementQuery) GroupBy(field string, fields ...string) *SendAnnouncementGroupBy {
	grbuild := &SendAnnouncementGroupBy{config: saq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return saq.sqlQuery(ctx), nil
	}
	grbuild.label = sendannouncement.Label
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
//	client.SendAnnouncement.Query().
//		Select(sendannouncement.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (saq *SendAnnouncementQuery) Select(fields ...string) *SendAnnouncementSelect {
	saq.fields = append(saq.fields, fields...)
	selbuild := &SendAnnouncementSelect{SendAnnouncementQuery: saq}
	selbuild.label = sendannouncement.Label
	selbuild.flds, selbuild.scan = &saq.fields, selbuild.Scan
	return selbuild
}

func (saq *SendAnnouncementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range saq.fields {
		if !sendannouncement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	if sendannouncement.Policy == nil {
		return errors.New("ent: uninitialized sendannouncement.Policy (forgotten import ent/runtime?)")
	}
	if err := sendannouncement.Policy.EvalQuery(ctx, saq); err != nil {
		return err
	}
	return nil
}

func (saq *SendAnnouncementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SendAnnouncement, error) {
	var (
		nodes = []*SendAnnouncement{}
		_spec = saq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SendAnnouncement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SendAnnouncement{config: saq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (saq *SendAnnouncementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	if len(saq.modifiers) > 0 {
		_spec.Modifiers = saq.modifiers
	}
	_spec.Node.Columns = saq.fields
	if len(saq.fields) > 0 {
		_spec.Unique = saq.unique != nil && *saq.unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *SendAnnouncementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := saq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (saq *SendAnnouncementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sendannouncement.Table,
			Columns: sendannouncement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: sendannouncement.FieldID,
			},
		},
		From:   saq.sql,
		Unique: true,
	}
	if unique := saq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := saq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sendannouncement.FieldID)
		for i := range fields {
			if fields[i] != sendannouncement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *SendAnnouncementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(sendannouncement.Table)
	columns := saq.fields
	if len(columns) == 0 {
		columns = sendannouncement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.unique != nil && *saq.unique {
		selector.Distinct()
	}
	for _, m := range saq.modifiers {
		m(selector)
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (saq *SendAnnouncementQuery) ForUpdate(opts ...sql.LockOption) *SendAnnouncementQuery {
	if saq.driver.Dialect() == dialect.Postgres {
		saq.Unique(false)
	}
	saq.modifiers = append(saq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return saq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (saq *SendAnnouncementQuery) ForShare(opts ...sql.LockOption) *SendAnnouncementQuery {
	if saq.driver.Dialect() == dialect.Postgres {
		saq.Unique(false)
	}
	saq.modifiers = append(saq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return saq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (saq *SendAnnouncementQuery) Modify(modifiers ...func(s *sql.Selector)) *SendAnnouncementSelect {
	saq.modifiers = append(saq.modifiers, modifiers...)
	return saq.Select()
}

// SendAnnouncementGroupBy is the group-by builder for SendAnnouncement entities.
type SendAnnouncementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SendAnnouncementGroupBy) Aggregate(fns ...AggregateFunc) *SendAnnouncementGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sagb *SendAnnouncementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sagb.path(ctx)
	if err != nil {
		return err
	}
	sagb.sql = query
	return sagb.sqlScan(ctx, v)
}

func (sagb *SendAnnouncementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sagb.fields {
		if !sendannouncement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sagb *SendAnnouncementGroupBy) sqlQuery() *sql.Selector {
	selector := sagb.sql.Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sagb.fields)+len(sagb.fns))
		for _, f := range sagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sagb.fields...)...)
}

// SendAnnouncementSelect is the builder for selecting fields of SendAnnouncement entities.
type SendAnnouncementSelect struct {
	*SendAnnouncementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SendAnnouncementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	sas.sql = sas.SendAnnouncementQuery.sqlQuery(ctx)
	return sas.sqlScan(ctx, v)
}

func (sas *SendAnnouncementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sas.sql.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sas *SendAnnouncementSelect) Modify(modifiers ...func(s *sql.Selector)) *SendAnnouncementSelect {
	sas.modifiers = append(sas.modifiers, modifiers...)
	return sas
}
