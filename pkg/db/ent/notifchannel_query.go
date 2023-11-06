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
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notifchannel"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
)

// NotifChannelQuery is the builder for querying NotifChannel entities.
type NotifChannelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NotifChannel
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotifChannelQuery builder.
func (ncq *NotifChannelQuery) Where(ps ...predicate.NotifChannel) *NotifChannelQuery {
	ncq.predicates = append(ncq.predicates, ps...)
	return ncq
}

// Limit adds a limit step to the query.
func (ncq *NotifChannelQuery) Limit(limit int) *NotifChannelQuery {
	ncq.limit = &limit
	return ncq
}

// Offset adds an offset step to the query.
func (ncq *NotifChannelQuery) Offset(offset int) *NotifChannelQuery {
	ncq.offset = &offset
	return ncq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ncq *NotifChannelQuery) Unique(unique bool) *NotifChannelQuery {
	ncq.unique = &unique
	return ncq
}

// Order adds an order step to the query.
func (ncq *NotifChannelQuery) Order(o ...OrderFunc) *NotifChannelQuery {
	ncq.order = append(ncq.order, o...)
	return ncq
}

// First returns the first NotifChannel entity from the query.
// Returns a *NotFoundError when no NotifChannel was found.
func (ncq *NotifChannelQuery) First(ctx context.Context) (*NotifChannel, error) {
	nodes, err := ncq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notifchannel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ncq *NotifChannelQuery) FirstX(ctx context.Context) *NotifChannel {
	node, err := ncq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotifChannel ID from the query.
// Returns a *NotFoundError when no NotifChannel ID was found.
func (ncq *NotifChannelQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ncq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notifchannel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ncq *NotifChannelQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := ncq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotifChannel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotifChannel entity is found.
// Returns a *NotFoundError when no NotifChannel entities are found.
func (ncq *NotifChannelQuery) Only(ctx context.Context) (*NotifChannel, error) {
	nodes, err := ncq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notifchannel.Label}
	default:
		return nil, &NotSingularError{notifchannel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ncq *NotifChannelQuery) OnlyX(ctx context.Context) *NotifChannel {
	node, err := ncq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotifChannel ID in the query.
// Returns a *NotSingularError when more than one NotifChannel ID is found.
// Returns a *NotFoundError when no entities are found.
func (ncq *NotifChannelQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ncq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notifchannel.Label}
	default:
		err = &NotSingularError{notifchannel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ncq *NotifChannelQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := ncq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotifChannels.
func (ncq *NotifChannelQuery) All(ctx context.Context) ([]*NotifChannel, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ncq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ncq *NotifChannelQuery) AllX(ctx context.Context) []*NotifChannel {
	nodes, err := ncq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotifChannel IDs.
func (ncq *NotifChannelQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := ncq.Select(notifchannel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ncq *NotifChannelQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := ncq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ncq *NotifChannelQuery) Count(ctx context.Context) (int, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ncq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ncq *NotifChannelQuery) CountX(ctx context.Context) int {
	count, err := ncq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ncq *NotifChannelQuery) Exist(ctx context.Context) (bool, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ncq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ncq *NotifChannelQuery) ExistX(ctx context.Context) bool {
	exist, err := ncq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotifChannelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ncq *NotifChannelQuery) Clone() *NotifChannelQuery {
	if ncq == nil {
		return nil
	}
	return &NotifChannelQuery{
		config:     ncq.config,
		limit:      ncq.limit,
		offset:     ncq.offset,
		order:      append([]OrderFunc{}, ncq.order...),
		predicates: append([]predicate.NotifChannel{}, ncq.predicates...),
		// clone intermediate query.
		sql:    ncq.sql.Clone(),
		path:   ncq.path,
		unique: ncq.unique,
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
//	client.NotifChannel.Query().
//		GroupBy(notifchannel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ncq *NotifChannelQuery) GroupBy(field string, fields ...string) *NotifChannelGroupBy {
	grbuild := &NotifChannelGroupBy{config: ncq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ncq.sqlQuery(ctx), nil
	}
	grbuild.label = notifchannel.Label
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
//	client.NotifChannel.Query().
//		Select(notifchannel.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ncq *NotifChannelQuery) Select(fields ...string) *NotifChannelSelect {
	ncq.fields = append(ncq.fields, fields...)
	selbuild := &NotifChannelSelect{NotifChannelQuery: ncq}
	selbuild.label = notifchannel.Label
	selbuild.flds, selbuild.scan = &ncq.fields, selbuild.Scan
	return selbuild
}

func (ncq *NotifChannelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ncq.fields {
		if !notifchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ncq.path != nil {
		prev, err := ncq.path(ctx)
		if err != nil {
			return err
		}
		ncq.sql = prev
	}
	if notifchannel.Policy == nil {
		return errors.New("ent: uninitialized notifchannel.Policy (forgotten import ent/runtime?)")
	}
	if err := notifchannel.Policy.EvalQuery(ctx, ncq); err != nil {
		return err
	}
	return nil
}

func (ncq *NotifChannelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotifChannel, error) {
	var (
		nodes = []*NotifChannel{}
		_spec = ncq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*NotifChannel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &NotifChannel{config: ncq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ncq.modifiers) > 0 {
		_spec.Modifiers = ncq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ncq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ncq *NotifChannelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ncq.querySpec()
	if len(ncq.modifiers) > 0 {
		_spec.Modifiers = ncq.modifiers
	}
	_spec.Node.Columns = ncq.fields
	if len(ncq.fields) > 0 {
		_spec.Unique = ncq.unique != nil && *ncq.unique
	}
	return sqlgraph.CountNodes(ctx, ncq.driver, _spec)
}

func (ncq *NotifChannelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ncq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ncq *NotifChannelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notifchannel.Table,
			Columns: notifchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: notifchannel.FieldID,
			},
		},
		From:   ncq.sql,
		Unique: true,
	}
	if unique := ncq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ncq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifchannel.FieldID)
		for i := range fields {
			if fields[i] != notifchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ncq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ncq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ncq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ncq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ncq *NotifChannelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ncq.driver.Dialect())
	t1 := builder.Table(notifchannel.Table)
	columns := ncq.fields
	if len(columns) == 0 {
		columns = notifchannel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ncq.sql != nil {
		selector = ncq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ncq.unique != nil && *ncq.unique {
		selector.Distinct()
	}
	for _, m := range ncq.modifiers {
		m(selector)
	}
	for _, p := range ncq.predicates {
		p(selector)
	}
	for _, p := range ncq.order {
		p(selector)
	}
	if offset := ncq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ncq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ncq *NotifChannelQuery) ForUpdate(opts ...sql.LockOption) *NotifChannelQuery {
	if ncq.driver.Dialect() == dialect.Postgres {
		ncq.Unique(false)
	}
	ncq.modifiers = append(ncq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ncq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ncq *NotifChannelQuery) ForShare(opts ...sql.LockOption) *NotifChannelQuery {
	if ncq.driver.Dialect() == dialect.Postgres {
		ncq.Unique(false)
	}
	ncq.modifiers = append(ncq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ncq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ncq *NotifChannelQuery) Modify(modifiers ...func(s *sql.Selector)) *NotifChannelSelect {
	ncq.modifiers = append(ncq.modifiers, modifiers...)
	return ncq.Select()
}

// NotifChannelGroupBy is the group-by builder for NotifChannel entities.
type NotifChannelGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ncgb *NotifChannelGroupBy) Aggregate(fns ...AggregateFunc) *NotifChannelGroupBy {
	ncgb.fns = append(ncgb.fns, fns...)
	return ncgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ncgb *NotifChannelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ncgb.path(ctx)
	if err != nil {
		return err
	}
	ncgb.sql = query
	return ncgb.sqlScan(ctx, v)
}

func (ncgb *NotifChannelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ncgb.fields {
		if !notifchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ncgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ncgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ncgb *NotifChannelGroupBy) sqlQuery() *sql.Selector {
	selector := ncgb.sql.Select()
	aggregation := make([]string, 0, len(ncgb.fns))
	for _, fn := range ncgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ncgb.fields)+len(ncgb.fns))
		for _, f := range ncgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ncgb.fields...)...)
}

// NotifChannelSelect is the builder for selecting fields of NotifChannel entities.
type NotifChannelSelect struct {
	*NotifChannelQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ncs *NotifChannelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ncs.prepareQuery(ctx); err != nil {
		return err
	}
	ncs.sql = ncs.NotifChannelQuery.sqlQuery(ctx)
	return ncs.sqlScan(ctx, v)
}

func (ncs *NotifChannelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ncs.sql.Query()
	if err := ncs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ncs *NotifChannelSelect) Modify(modifiers ...func(s *sql.Selector)) *NotifChannelSelect {
	ncs.modifiers = append(ncs.modifiers, modifiers...)
	return ncs
}
