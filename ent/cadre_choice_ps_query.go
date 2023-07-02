// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/cadre_choice_ps"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreChoicePSQuery is the builder for querying Cadre_Choice_PS entities.
type CadreChoicePSQuery struct {
	config
	ctx        *QueryContext
	order      []cadre_choice_ps.OrderOption
	inters     []Interceptor
	predicates []predicate.Cadre_Choice_PS
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CadreChoicePSQuery builder.
func (ccpq *CadreChoicePSQuery) Where(ps ...predicate.Cadre_Choice_PS) *CadreChoicePSQuery {
	ccpq.predicates = append(ccpq.predicates, ps...)
	return ccpq
}

// Limit the number of records to be returned by this query.
func (ccpq *CadreChoicePSQuery) Limit(limit int) *CadreChoicePSQuery {
	ccpq.ctx.Limit = &limit
	return ccpq
}

// Offset to start from.
func (ccpq *CadreChoicePSQuery) Offset(offset int) *CadreChoicePSQuery {
	ccpq.ctx.Offset = &offset
	return ccpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccpq *CadreChoicePSQuery) Unique(unique bool) *CadreChoicePSQuery {
	ccpq.ctx.Unique = &unique
	return ccpq
}

// Order specifies how the records should be ordered.
func (ccpq *CadreChoicePSQuery) Order(o ...cadre_choice_ps.OrderOption) *CadreChoicePSQuery {
	ccpq.order = append(ccpq.order, o...)
	return ccpq
}

// First returns the first Cadre_Choice_PS entity from the query.
// Returns a *NotFoundError when no Cadre_Choice_PS was found.
func (ccpq *CadreChoicePSQuery) First(ctx context.Context) (*Cadre_Choice_PS, error) {
	nodes, err := ccpq.Limit(1).All(setContextOp(ctx, ccpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cadre_choice_ps.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) FirstX(ctx context.Context) *Cadre_Choice_PS {
	node, err := ccpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cadre_Choice_PS ID from the query.
// Returns a *NotFoundError when no Cadre_Choice_PS ID was found.
func (ccpq *CadreChoicePSQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ccpq.Limit(1).IDs(setContextOp(ctx, ccpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cadre_choice_ps.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ccpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cadre_Choice_PS entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cadre_Choice_PS entity is found.
// Returns a *NotFoundError when no Cadre_Choice_PS entities are found.
func (ccpq *CadreChoicePSQuery) Only(ctx context.Context) (*Cadre_Choice_PS, error) {
	nodes, err := ccpq.Limit(2).All(setContextOp(ctx, ccpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cadre_choice_ps.Label}
	default:
		return nil, &NotSingularError{cadre_choice_ps.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) OnlyX(ctx context.Context) *Cadre_Choice_PS {
	node, err := ccpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cadre_Choice_PS ID in the query.
// Returns a *NotSingularError when more than one Cadre_Choice_PS ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccpq *CadreChoicePSQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ccpq.Limit(2).IDs(setContextOp(ctx, ccpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cadre_choice_ps.Label}
	default:
		err = &NotSingularError{cadre_choice_ps.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ccpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cadre_Choice_PSs.
func (ccpq *CadreChoicePSQuery) All(ctx context.Context) ([]*Cadre_Choice_PS, error) {
	ctx = setContextOp(ctx, ccpq.ctx, "All")
	if err := ccpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cadre_Choice_PS, *CadreChoicePSQuery]()
	return withInterceptors[[]*Cadre_Choice_PS](ctx, ccpq, qr, ccpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) AllX(ctx context.Context) []*Cadre_Choice_PS {
	nodes, err := ccpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cadre_Choice_PS IDs.
func (ccpq *CadreChoicePSQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ccpq.ctx.Unique == nil && ccpq.path != nil {
		ccpq.Unique(true)
	}
	ctx = setContextOp(ctx, ccpq.ctx, "IDs")
	if err = ccpq.Select(cadre_choice_ps.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ccpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccpq *CadreChoicePSQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ccpq.ctx, "Count")
	if err := ccpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ccpq, querierCount[*CadreChoicePSQuery](), ccpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) CountX(ctx context.Context) int {
	count, err := ccpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccpq *CadreChoicePSQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ccpq.ctx, "Exist")
	switch _, err := ccpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ccpq *CadreChoicePSQuery) ExistX(ctx context.Context) bool {
	exist, err := ccpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CadreChoicePSQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccpq *CadreChoicePSQuery) Clone() *CadreChoicePSQuery {
	if ccpq == nil {
		return nil
	}
	return &CadreChoicePSQuery{
		config:     ccpq.config,
		ctx:        ccpq.ctx.Clone(),
		order:      append([]cadre_choice_ps.OrderOption{}, ccpq.order...),
		inters:     append([]Interceptor{}, ccpq.inters...),
		predicates: append([]predicate.Cadre_Choice_PS{}, ccpq.predicates...),
		// clone intermediate query.
		sql:  ccpq.sql.Clone(),
		path: ccpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Application string `json:"Application,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CadreChoicePS.Query().
//		GroupBy(cadre_choice_ps.FieldApplication).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ccpq *CadreChoicePSQuery) GroupBy(field string, fields ...string) *CadreChoicePSGroupBy {
	ccpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CadreChoicePSGroupBy{build: ccpq}
	grbuild.flds = &ccpq.ctx.Fields
	grbuild.label = cadre_choice_ps.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Application string `json:"Application,omitempty"`
//	}
//
//	client.CadreChoicePS.Query().
//		Select(cadre_choice_ps.FieldApplication).
//		Scan(ctx, &v)
func (ccpq *CadreChoicePSQuery) Select(fields ...string) *CadreChoicePSSelect {
	ccpq.ctx.Fields = append(ccpq.ctx.Fields, fields...)
	sbuild := &CadreChoicePSSelect{CadreChoicePSQuery: ccpq}
	sbuild.label = cadre_choice_ps.Label
	sbuild.flds, sbuild.scan = &ccpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CadreChoicePSSelect configured with the given aggregations.
func (ccpq *CadreChoicePSQuery) Aggregate(fns ...AggregateFunc) *CadreChoicePSSelect {
	return ccpq.Select().Aggregate(fns...)
}

func (ccpq *CadreChoicePSQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ccpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ccpq); err != nil {
				return err
			}
		}
	}
	for _, f := range ccpq.ctx.Fields {
		if !cadre_choice_ps.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccpq.path != nil {
		prev, err := ccpq.path(ctx)
		if err != nil {
			return err
		}
		ccpq.sql = prev
	}
	return nil
}

func (ccpq *CadreChoicePSQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cadre_Choice_PS, error) {
	var (
		nodes = []*Cadre_Choice_PS{}
		_spec = ccpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cadre_Choice_PS).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cadre_Choice_PS{config: ccpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccpq *CadreChoicePSQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccpq.querySpec()
	_spec.Node.Columns = ccpq.ctx.Fields
	if len(ccpq.ctx.Fields) > 0 {
		_spec.Unique = ccpq.ctx.Unique != nil && *ccpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ccpq.driver, _spec)
}

func (ccpq *CadreChoicePSQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cadre_choice_ps.Table, cadre_choice_ps.Columns, sqlgraph.NewFieldSpec(cadre_choice_ps.FieldID, field.TypeInt32))
	_spec.From = ccpq.sql
	if unique := ccpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ccpq.path != nil {
		_spec.Unique = true
	}
	if fields := ccpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cadre_choice_ps.FieldID)
		for i := range fields {
			if fields[i] != cadre_choice_ps.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccpq *CadreChoicePSQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccpq.driver.Dialect())
	t1 := builder.Table(cadre_choice_ps.Table)
	columns := ccpq.ctx.Fields
	if len(columns) == 0 {
		columns = cadre_choice_ps.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccpq.sql != nil {
		selector = ccpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccpq.ctx.Unique != nil && *ccpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ccpq.predicates {
		p(selector)
	}
	for _, p := range ccpq.order {
		p(selector)
	}
	if offset := ccpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CadreChoicePSGroupBy is the group-by builder for Cadre_Choice_PS entities.
type CadreChoicePSGroupBy struct {
	selector
	build *CadreChoicePSQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccpgb *CadreChoicePSGroupBy) Aggregate(fns ...AggregateFunc) *CadreChoicePSGroupBy {
	ccpgb.fns = append(ccpgb.fns, fns...)
	return ccpgb
}

// Scan applies the selector query and scans the result into the given value.
func (ccpgb *CadreChoicePSGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccpgb.build.ctx, "GroupBy")
	if err := ccpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CadreChoicePSQuery, *CadreChoicePSGroupBy](ctx, ccpgb.build, ccpgb, ccpgb.build.inters, v)
}

func (ccpgb *CadreChoicePSGroupBy) sqlScan(ctx context.Context, root *CadreChoicePSQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ccpgb.fns))
	for _, fn := range ccpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ccpgb.flds)+len(ccpgb.fns))
		for _, f := range *ccpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ccpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CadreChoicePSSelect is the builder for selecting fields of CadreChoicePS entities.
type CadreChoicePSSelect struct {
	*CadreChoicePSQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ccps *CadreChoicePSSelect) Aggregate(fns ...AggregateFunc) *CadreChoicePSSelect {
	ccps.fns = append(ccps.fns, fns...)
	return ccps
}

// Scan applies the selector query and scans the result into the given value.
func (ccps *CadreChoicePSSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ccps.ctx, "Select")
	if err := ccps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CadreChoicePSQuery, *CadreChoicePSSelect](ctx, ccps.CadreChoicePSQuery, ccps, ccps.inters, v)
}

func (ccps *CadreChoicePSSelect) sqlScan(ctx context.Context, root *CadreChoicePSQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ccps.fns))
	for _, fn := range ccps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ccps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}