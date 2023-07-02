// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/employeecategory"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeCategoryQuery is the builder for querying EmployeeCategory entities.
type EmployeeCategoryQuery struct {
	config
	ctx        *QueryContext
	order      []employeecategory.OrderOption
	inters     []Interceptor
	predicates []predicate.EmployeeCategory
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeCategoryQuery builder.
func (ecq *EmployeeCategoryQuery) Where(ps ...predicate.EmployeeCategory) *EmployeeCategoryQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EmployeeCategoryQuery) Limit(limit int) *EmployeeCategoryQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EmployeeCategoryQuery) Offset(offset int) *EmployeeCategoryQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EmployeeCategoryQuery) Unique(unique bool) *EmployeeCategoryQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EmployeeCategoryQuery) Order(o ...employeecategory.OrderOption) *EmployeeCategoryQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// First returns the first EmployeeCategory entity from the query.
// Returns a *NotFoundError when no EmployeeCategory was found.
func (ecq *EmployeeCategoryQuery) First(ctx context.Context) (*EmployeeCategory, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employeecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) FirstX(ctx context.Context) *EmployeeCategory {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmployeeCategory ID from the query.
// Returns a *NotFoundError when no EmployeeCategory ID was found.
func (ecq *EmployeeCategoryQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employeecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmployeeCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmployeeCategory entity is found.
// Returns a *NotFoundError when no EmployeeCategory entities are found.
func (ecq *EmployeeCategoryQuery) Only(ctx context.Context) (*EmployeeCategory, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employeecategory.Label}
	default:
		return nil, &NotSingularError{employeecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) OnlyX(ctx context.Context) *EmployeeCategory {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmployeeCategory ID in the query.
// Returns a *NotSingularError when more than one EmployeeCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EmployeeCategoryQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employeecategory.Label}
	default:
		err = &NotSingularError{employeecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmployeeCategories.
func (ecq *EmployeeCategoryQuery) All(ctx context.Context) ([]*EmployeeCategory, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmployeeCategory, *EmployeeCategoryQuery]()
	return withInterceptors[[]*EmployeeCategory](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) AllX(ctx context.Context) []*EmployeeCategory {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmployeeCategory IDs.
func (ecq *EmployeeCategoryQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(employeecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EmployeeCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EmployeeCategoryQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EmployeeCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Exist")
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EmployeeCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EmployeeCategoryQuery) Clone() *EmployeeCategoryQuery {
	if ecq == nil {
		return nil
	}
	return &EmployeeCategoryQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]employeecategory.OrderOption{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.EmployeeCategory{}, ecq.predicates...),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Categrycode string `json:"Categrycode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmployeeCategory.Query().
//		GroupBy(employeecategory.FieldCategrycode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EmployeeCategoryQuery) GroupBy(field string, fields ...string) *EmployeeCategoryGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeeCategoryGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = employeecategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Categrycode string `json:"Categrycode,omitempty"`
//	}
//
//	client.EmployeeCategory.Query().
//		Select(employeecategory.FieldCategrycode).
//		Scan(ctx, &v)
func (ecq *EmployeeCategoryQuery) Select(fields ...string) *EmployeeCategorySelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EmployeeCategorySelect{EmployeeCategoryQuery: ecq}
	sbuild.label = employeecategory.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeeCategorySelect configured with the given aggregations.
func (ecq *EmployeeCategoryQuery) Aggregate(fns ...AggregateFunc) *EmployeeCategorySelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EmployeeCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !employeecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *EmployeeCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmployeeCategory, error) {
	var (
		nodes   = []*EmployeeCategory{}
		withFKs = ecq.withFKs
		_spec   = ecq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, employeecategory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmployeeCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmployeeCategory{config: ecq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ecq *EmployeeCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EmployeeCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employeecategory.Table, employeecategory.Columns, sqlgraph.NewFieldSpec(employeecategory.FieldID, field.TypeInt32))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeecategory.FieldID)
		for i := range fields {
			if fields[i] != employeecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EmployeeCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(employeecategory.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = employeecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmployeeCategoryGroupBy is the group-by builder for EmployeeCategory entities.
type EmployeeCategoryGroupBy struct {
	selector
	build *EmployeeCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EmployeeCategoryGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeCategoryGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EmployeeCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeCategoryQuery, *EmployeeCategoryGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EmployeeCategoryGroupBy) sqlScan(ctx context.Context, root *EmployeeCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmployeeCategorySelect is the builder for selecting fields of EmployeeCategory entities.
type EmployeeCategorySelect struct {
	*EmployeeCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EmployeeCategorySelect) Aggregate(fns ...AggregateFunc) *EmployeeCategorySelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EmployeeCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeCategoryQuery, *EmployeeCategorySelect](ctx, ecs.EmployeeCategoryQuery, ecs, ecs.inters, v)
}

func (ecs *EmployeeCategorySelect) sqlScan(ctx context.Context, root *EmployeeCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
