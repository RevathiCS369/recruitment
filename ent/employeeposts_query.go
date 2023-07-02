// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/employeeposts"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeePostsQuery is the builder for querying EmployeePosts entities.
type EmployeePostsQuery struct {
	config
	ctx        *QueryContext
	order      []employeeposts.OrderOption
	inters     []Interceptor
	predicates []predicate.EmployeePosts
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeePostsQuery builder.
func (epq *EmployeePostsQuery) Where(ps ...predicate.EmployeePosts) *EmployeePostsQuery {
	epq.predicates = append(epq.predicates, ps...)
	return epq
}

// Limit the number of records to be returned by this query.
func (epq *EmployeePostsQuery) Limit(limit int) *EmployeePostsQuery {
	epq.ctx.Limit = &limit
	return epq
}

// Offset to start from.
func (epq *EmployeePostsQuery) Offset(offset int) *EmployeePostsQuery {
	epq.ctx.Offset = &offset
	return epq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (epq *EmployeePostsQuery) Unique(unique bool) *EmployeePostsQuery {
	epq.ctx.Unique = &unique
	return epq
}

// Order specifies how the records should be ordered.
func (epq *EmployeePostsQuery) Order(o ...employeeposts.OrderOption) *EmployeePostsQuery {
	epq.order = append(epq.order, o...)
	return epq
}

// First returns the first EmployeePosts entity from the query.
// Returns a *NotFoundError when no EmployeePosts was found.
func (epq *EmployeePostsQuery) First(ctx context.Context) (*EmployeePosts, error) {
	nodes, err := epq.Limit(1).All(setContextOp(ctx, epq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employeeposts.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epq *EmployeePostsQuery) FirstX(ctx context.Context) *EmployeePosts {
	node, err := epq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmployeePosts ID from the query.
// Returns a *NotFoundError when no EmployeePosts ID was found.
func (epq *EmployeePostsQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = epq.Limit(1).IDs(setContextOp(ctx, epq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employeeposts.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (epq *EmployeePostsQuery) FirstIDX(ctx context.Context) int32 {
	id, err := epq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmployeePosts entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmployeePosts entity is found.
// Returns a *NotFoundError when no EmployeePosts entities are found.
func (epq *EmployeePostsQuery) Only(ctx context.Context) (*EmployeePosts, error) {
	nodes, err := epq.Limit(2).All(setContextOp(ctx, epq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employeeposts.Label}
	default:
		return nil, &NotSingularError{employeeposts.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epq *EmployeePostsQuery) OnlyX(ctx context.Context) *EmployeePosts {
	node, err := epq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmployeePosts ID in the query.
// Returns a *NotSingularError when more than one EmployeePosts ID is found.
// Returns a *NotFoundError when no entities are found.
func (epq *EmployeePostsQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = epq.Limit(2).IDs(setContextOp(ctx, epq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employeeposts.Label}
	default:
		err = &NotSingularError{employeeposts.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (epq *EmployeePostsQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := epq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmployeePostsSlice.
func (epq *EmployeePostsQuery) All(ctx context.Context) ([]*EmployeePosts, error) {
	ctx = setContextOp(ctx, epq.ctx, "All")
	if err := epq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmployeePosts, *EmployeePostsQuery]()
	return withInterceptors[[]*EmployeePosts](ctx, epq, qr, epq.inters)
}

// AllX is like All, but panics if an error occurs.
func (epq *EmployeePostsQuery) AllX(ctx context.Context) []*EmployeePosts {
	nodes, err := epq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmployeePosts IDs.
func (epq *EmployeePostsQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if epq.ctx.Unique == nil && epq.path != nil {
		epq.Unique(true)
	}
	ctx = setContextOp(ctx, epq.ctx, "IDs")
	if err = epq.Select(employeeposts.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epq *EmployeePostsQuery) IDsX(ctx context.Context) []int32 {
	ids, err := epq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epq *EmployeePostsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, epq.ctx, "Count")
	if err := epq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, epq, querierCount[*EmployeePostsQuery](), epq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (epq *EmployeePostsQuery) CountX(ctx context.Context) int {
	count, err := epq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epq *EmployeePostsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, epq.ctx, "Exist")
	switch _, err := epq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (epq *EmployeePostsQuery) ExistX(ctx context.Context) bool {
	exist, err := epq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeePostsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epq *EmployeePostsQuery) Clone() *EmployeePostsQuery {
	if epq == nil {
		return nil
	}
	return &EmployeePostsQuery{
		config:     epq.config,
		ctx:        epq.ctx.Clone(),
		order:      append([]employeeposts.OrderOption{}, epq.order...),
		inters:     append([]Interceptor{}, epq.inters...),
		predicates: append([]predicate.EmployeePosts{}, epq.predicates...),
		// clone intermediate query.
		sql:  epq.sql.Clone(),
		path: epq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PostCode string `json:"PostCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmployeePosts.Query().
//		GroupBy(employeeposts.FieldPostCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (epq *EmployeePostsQuery) GroupBy(field string, fields ...string) *EmployeePostsGroupBy {
	epq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeePostsGroupBy{build: epq}
	grbuild.flds = &epq.ctx.Fields
	grbuild.label = employeeposts.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PostCode string `json:"PostCode,omitempty"`
//	}
//
//	client.EmployeePosts.Query().
//		Select(employeeposts.FieldPostCode).
//		Scan(ctx, &v)
func (epq *EmployeePostsQuery) Select(fields ...string) *EmployeePostsSelect {
	epq.ctx.Fields = append(epq.ctx.Fields, fields...)
	sbuild := &EmployeePostsSelect{EmployeePostsQuery: epq}
	sbuild.label = employeeposts.Label
	sbuild.flds, sbuild.scan = &epq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeePostsSelect configured with the given aggregations.
func (epq *EmployeePostsQuery) Aggregate(fns ...AggregateFunc) *EmployeePostsSelect {
	return epq.Select().Aggregate(fns...)
}

func (epq *EmployeePostsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range epq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, epq); err != nil {
				return err
			}
		}
	}
	for _, f := range epq.ctx.Fields {
		if !employeeposts.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if epq.path != nil {
		prev, err := epq.path(ctx)
		if err != nil {
			return err
		}
		epq.sql = prev
	}
	return nil
}

func (epq *EmployeePostsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmployeePosts, error) {
	var (
		nodes = []*EmployeePosts{}
		_spec = epq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmployeePosts).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmployeePosts{config: epq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, epq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (epq *EmployeePostsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := epq.querySpec()
	_spec.Node.Columns = epq.ctx.Fields
	if len(epq.ctx.Fields) > 0 {
		_spec.Unique = epq.ctx.Unique != nil && *epq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, epq.driver, _spec)
}

func (epq *EmployeePostsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employeeposts.Table, employeeposts.Columns, sqlgraph.NewFieldSpec(employeeposts.FieldID, field.TypeInt32))
	_spec.From = epq.sql
	if unique := epq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if epq.path != nil {
		_spec.Unique = true
	}
	if fields := epq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeeposts.FieldID)
		for i := range fields {
			if fields[i] != employeeposts.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := epq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := epq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := epq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := epq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (epq *EmployeePostsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(employeeposts.Table)
	columns := epq.ctx.Fields
	if len(columns) == 0 {
		columns = employeeposts.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if epq.sql != nil {
		selector = epq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if epq.ctx.Unique != nil && *epq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range epq.predicates {
		p(selector)
	}
	for _, p := range epq.order {
		p(selector)
	}
	if offset := epq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmployeePostsGroupBy is the group-by builder for EmployeePosts entities.
type EmployeePostsGroupBy struct {
	selector
	build *EmployeePostsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epgb *EmployeePostsGroupBy) Aggregate(fns ...AggregateFunc) *EmployeePostsGroupBy {
	epgb.fns = append(epgb.fns, fns...)
	return epgb
}

// Scan applies the selector query and scans the result into the given value.
func (epgb *EmployeePostsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epgb.build.ctx, "GroupBy")
	if err := epgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeePostsQuery, *EmployeePostsGroupBy](ctx, epgb.build, epgb, epgb.build.inters, v)
}

func (epgb *EmployeePostsGroupBy) sqlScan(ctx context.Context, root *EmployeePostsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(epgb.fns))
	for _, fn := range epgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*epgb.flds)+len(epgb.fns))
		for _, f := range *epgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*epgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmployeePostsSelect is the builder for selecting fields of EmployeePosts entities.
type EmployeePostsSelect struct {
	*EmployeePostsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eps *EmployeePostsSelect) Aggregate(fns ...AggregateFunc) *EmployeePostsSelect {
	eps.fns = append(eps.fns, fns...)
	return eps
}

// Scan applies the selector query and scans the result into the given value.
func (eps *EmployeePostsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eps.ctx, "Select")
	if err := eps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeePostsQuery, *EmployeePostsSelect](ctx, eps.EmployeePostsQuery, eps, eps.inters, v)
}

func (eps *EmployeePostsSelect) sqlScan(ctx context.Context, root *EmployeePostsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eps.fns))
	for _, fn := range eps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
