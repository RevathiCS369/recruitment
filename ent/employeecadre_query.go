// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/employeecadre"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeCadreQuery is the builder for querying EmployeeCadre entities.
type EmployeeCadreQuery struct {
	config
	ctx        *QueryContext
	order      []employeecadre.OrderOption
	inters     []Interceptor
	predicates []predicate.EmployeeCadre
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeCadreQuery builder.
func (ecq *EmployeeCadreQuery) Where(ps ...predicate.EmployeeCadre) *EmployeeCadreQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EmployeeCadreQuery) Limit(limit int) *EmployeeCadreQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EmployeeCadreQuery) Offset(offset int) *EmployeeCadreQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EmployeeCadreQuery) Unique(unique bool) *EmployeeCadreQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EmployeeCadreQuery) Order(o ...employeecadre.OrderOption) *EmployeeCadreQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// First returns the first EmployeeCadre entity from the query.
// Returns a *NotFoundError when no EmployeeCadre was found.
func (ecq *EmployeeCadreQuery) First(ctx context.Context) (*EmployeeCadre, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employeecadre.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) FirstX(ctx context.Context) *EmployeeCadre {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmployeeCadre ID from the query.
// Returns a *NotFoundError when no EmployeeCadre ID was found.
func (ecq *EmployeeCadreQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employeecadre.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmployeeCadre entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmployeeCadre entity is found.
// Returns a *NotFoundError when no EmployeeCadre entities are found.
func (ecq *EmployeeCadreQuery) Only(ctx context.Context) (*EmployeeCadre, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employeecadre.Label}
	default:
		return nil, &NotSingularError{employeecadre.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) OnlyX(ctx context.Context) *EmployeeCadre {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmployeeCadre ID in the query.
// Returns a *NotSingularError when more than one EmployeeCadre ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EmployeeCadreQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employeecadre.Label}
	default:
		err = &NotSingularError{employeecadre.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmployeeCadres.
func (ecq *EmployeeCadreQuery) All(ctx context.Context) ([]*EmployeeCadre, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmployeeCadre, *EmployeeCadreQuery]()
	return withInterceptors[[]*EmployeeCadre](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) AllX(ctx context.Context) []*EmployeeCadre {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmployeeCadre IDs.
func (ecq *EmployeeCadreQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(employeecadre.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EmployeeCadreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EmployeeCadreQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EmployeeCadreQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EmployeeCadreQuery) Exist(ctx context.Context) (bool, error) {
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
func (ecq *EmployeeCadreQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeCadreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EmployeeCadreQuery) Clone() *EmployeeCadreQuery {
	if ecq == nil {
		return nil
	}
	return &EmployeeCadreQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]employeecadre.OrderOption{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.EmployeeCadre{}, ecq.predicates...),
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
//		Cadrecode string `json:"cadrecode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmployeeCadre.Query().
//		GroupBy(employeecadre.FieldCadrecode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EmployeeCadreQuery) GroupBy(field string, fields ...string) *EmployeeCadreGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeeCadreGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = employeecadre.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Cadrecode string `json:"cadrecode,omitempty"`
//	}
//
//	client.EmployeeCadre.Query().
//		Select(employeecadre.FieldCadrecode).
//		Scan(ctx, &v)
func (ecq *EmployeeCadreQuery) Select(fields ...string) *EmployeeCadreSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EmployeeCadreSelect{EmployeeCadreQuery: ecq}
	sbuild.label = employeecadre.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeeCadreSelect configured with the given aggregations.
func (ecq *EmployeeCadreQuery) Aggregate(fns ...AggregateFunc) *EmployeeCadreSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EmployeeCadreQuery) prepareQuery(ctx context.Context) error {
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
		if !employeecadre.ValidColumn(f) {
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

func (ecq *EmployeeCadreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmployeeCadre, error) {
	var (
		nodes = []*EmployeeCadre{}
		_spec = ecq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmployeeCadre).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmployeeCadre{config: ecq.config}
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

func (ecq *EmployeeCadreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EmployeeCadreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employeecadre.Table, employeecadre.Columns, sqlgraph.NewFieldSpec(employeecadre.FieldID, field.TypeInt32))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeecadre.FieldID)
		for i := range fields {
			if fields[i] != employeecadre.FieldID {
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

func (ecq *EmployeeCadreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(employeecadre.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = employeecadre.Columns
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

// EmployeeCadreGroupBy is the group-by builder for EmployeeCadre entities.
type EmployeeCadreGroupBy struct {
	selector
	build *EmployeeCadreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EmployeeCadreGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeCadreGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EmployeeCadreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeCadreQuery, *EmployeeCadreGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EmployeeCadreGroupBy) sqlScan(ctx context.Context, root *EmployeeCadreQuery, v any) error {
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

// EmployeeCadreSelect is the builder for selecting fields of EmployeeCadre entities.
type EmployeeCadreSelect struct {
	*EmployeeCadreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EmployeeCadreSelect) Aggregate(fns ...AggregateFunc) *EmployeeCadreSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EmployeeCadreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeCadreQuery, *EmployeeCadreSelect](ctx, ecs.EmployeeCadreQuery, ecs, ecs.inters, v)
}

func (ecs *EmployeeCadreSelect) sqlScan(ctx context.Context, root *EmployeeCadreQuery, v any) error {
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
