// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/divisionmaster"
	"recruit/ent/predicate"
	"recruit/ent/regionmaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionMasterQuery is the builder for querying DivisionMaster entities.
type DivisionMasterQuery struct {
	config
	ctx         *QueryContext
	order       []divisionmaster.OrderOption
	inters      []Interceptor
	predicates  []predicate.DivisionMaster
	withRegions *RegionMasterQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DivisionMasterQuery builder.
func (dmq *DivisionMasterQuery) Where(ps ...predicate.DivisionMaster) *DivisionMasterQuery {
	dmq.predicates = append(dmq.predicates, ps...)
	return dmq
}

// Limit the number of records to be returned by this query.
func (dmq *DivisionMasterQuery) Limit(limit int) *DivisionMasterQuery {
	dmq.ctx.Limit = &limit
	return dmq
}

// Offset to start from.
func (dmq *DivisionMasterQuery) Offset(offset int) *DivisionMasterQuery {
	dmq.ctx.Offset = &offset
	return dmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmq *DivisionMasterQuery) Unique(unique bool) *DivisionMasterQuery {
	dmq.ctx.Unique = &unique
	return dmq
}

// Order specifies how the records should be ordered.
func (dmq *DivisionMasterQuery) Order(o ...divisionmaster.OrderOption) *DivisionMasterQuery {
	dmq.order = append(dmq.order, o...)
	return dmq
}

// QueryRegions chains the current query on the "regions" edge.
func (dmq *DivisionMasterQuery) QueryRegions() *RegionMasterQuery {
	query := (&RegionMasterClient{config: dmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(divisionmaster.Table, divisionmaster.FieldID, selector),
			sqlgraph.To(regionmaster.Table, regionmaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, divisionmaster.RegionsTable, divisionmaster.RegionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DivisionMaster entity from the query.
// Returns a *NotFoundError when no DivisionMaster was found.
func (dmq *DivisionMasterQuery) First(ctx context.Context) (*DivisionMaster, error) {
	nodes, err := dmq.Limit(1).All(setContextOp(ctx, dmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{divisionmaster.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmq *DivisionMasterQuery) FirstX(ctx context.Context) *DivisionMaster {
	node, err := dmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DivisionMaster ID from the query.
// Returns a *NotFoundError when no DivisionMaster ID was found.
func (dmq *DivisionMasterQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dmq.Limit(1).IDs(setContextOp(ctx, dmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{divisionmaster.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmq *DivisionMasterQuery) FirstIDX(ctx context.Context) int32 {
	id, err := dmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DivisionMaster entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DivisionMaster entity is found.
// Returns a *NotFoundError when no DivisionMaster entities are found.
func (dmq *DivisionMasterQuery) Only(ctx context.Context) (*DivisionMaster, error) {
	nodes, err := dmq.Limit(2).All(setContextOp(ctx, dmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{divisionmaster.Label}
	default:
		return nil, &NotSingularError{divisionmaster.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmq *DivisionMasterQuery) OnlyX(ctx context.Context) *DivisionMaster {
	node, err := dmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DivisionMaster ID in the query.
// Returns a *NotSingularError when more than one DivisionMaster ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmq *DivisionMasterQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dmq.Limit(2).IDs(setContextOp(ctx, dmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{divisionmaster.Label}
	default:
		err = &NotSingularError{divisionmaster.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmq *DivisionMasterQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := dmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DivisionMasters.
func (dmq *DivisionMasterQuery) All(ctx context.Context) ([]*DivisionMaster, error) {
	ctx = setContextOp(ctx, dmq.ctx, "All")
	if err := dmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DivisionMaster, *DivisionMasterQuery]()
	return withInterceptors[[]*DivisionMaster](ctx, dmq, qr, dmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dmq *DivisionMasterQuery) AllX(ctx context.Context) []*DivisionMaster {
	nodes, err := dmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DivisionMaster IDs.
func (dmq *DivisionMasterQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if dmq.ctx.Unique == nil && dmq.path != nil {
		dmq.Unique(true)
	}
	ctx = setContextOp(ctx, dmq.ctx, "IDs")
	if err = dmq.Select(divisionmaster.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmq *DivisionMasterQuery) IDsX(ctx context.Context) []int32 {
	ids, err := dmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmq *DivisionMasterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Count")
	if err := dmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dmq, querierCount[*DivisionMasterQuery](), dmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dmq *DivisionMasterQuery) CountX(ctx context.Context) int {
	count, err := dmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmq *DivisionMasterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Exist")
	switch _, err := dmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dmq *DivisionMasterQuery) ExistX(ctx context.Context) bool {
	exist, err := dmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DivisionMasterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmq *DivisionMasterQuery) Clone() *DivisionMasterQuery {
	if dmq == nil {
		return nil
	}
	return &DivisionMasterQuery{
		config:      dmq.config,
		ctx:         dmq.ctx.Clone(),
		order:       append([]divisionmaster.OrderOption{}, dmq.order...),
		inters:      append([]Interceptor{}, dmq.inters...),
		predicates:  append([]predicate.DivisionMaster{}, dmq.predicates...),
		withRegions: dmq.withRegions.Clone(),
		// clone intermediate query.
		sql:  dmq.sql.Clone(),
		path: dmq.path,
	}
}

// WithRegions tells the query-builder to eager-load the nodes that are connected to
// the "regions" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DivisionMasterQuery) WithRegions(opts ...func(*RegionMasterQuery)) *DivisionMasterQuery {
	query := (&RegionMasterClient{config: dmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dmq.withRegions = query
	return dmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DivisionCode int32 `json:"DivisionCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DivisionMaster.Query().
//		GroupBy(divisionmaster.FieldDivisionCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dmq *DivisionMasterQuery) GroupBy(field string, fields ...string) *DivisionMasterGroupBy {
	dmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DivisionMasterGroupBy{build: dmq}
	grbuild.flds = &dmq.ctx.Fields
	grbuild.label = divisionmaster.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DivisionCode int32 `json:"DivisionCode,omitempty"`
//	}
//
//	client.DivisionMaster.Query().
//		Select(divisionmaster.FieldDivisionCode).
//		Scan(ctx, &v)
func (dmq *DivisionMasterQuery) Select(fields ...string) *DivisionMasterSelect {
	dmq.ctx.Fields = append(dmq.ctx.Fields, fields...)
	sbuild := &DivisionMasterSelect{DivisionMasterQuery: dmq}
	sbuild.label = divisionmaster.Label
	sbuild.flds, sbuild.scan = &dmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DivisionMasterSelect configured with the given aggregations.
func (dmq *DivisionMasterQuery) Aggregate(fns ...AggregateFunc) *DivisionMasterSelect {
	return dmq.Select().Aggregate(fns...)
}

func (dmq *DivisionMasterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dmq); err != nil {
				return err
			}
		}
	}
	for _, f := range dmq.ctx.Fields {
		if !divisionmaster.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dmq.path != nil {
		prev, err := dmq.path(ctx)
		if err != nil {
			return err
		}
		dmq.sql = prev
	}
	return nil
}

func (dmq *DivisionMasterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DivisionMaster, error) {
	var (
		nodes       = []*DivisionMaster{}
		withFKs     = dmq.withFKs
		_spec       = dmq.querySpec()
		loadedTypes = [1]bool{
			dmq.withRegions != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, divisionmaster.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DivisionMaster).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DivisionMaster{config: dmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dmq.withRegions; query != nil {
		if err := dmq.loadRegions(ctx, query, nodes,
			func(n *DivisionMaster) { n.Edges.Regions = []*RegionMaster{} },
			func(n *DivisionMaster, e *RegionMaster) { n.Edges.Regions = append(n.Edges.Regions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dmq *DivisionMasterQuery) loadRegions(ctx context.Context, query *RegionMasterQuery, nodes []*DivisionMaster, init func(*DivisionMaster), assign func(*DivisionMaster, *RegionMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*DivisionMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RegionMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(divisionmaster.RegionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.division_master_regions
		if fk == nil {
			return fmt.Errorf(`foreign-key "division_master_regions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "division_master_regions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dmq *DivisionMasterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmq.querySpec()
	_spec.Node.Columns = dmq.ctx.Fields
	if len(dmq.ctx.Fields) > 0 {
		_spec.Unique = dmq.ctx.Unique != nil && *dmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dmq.driver, _spec)
}

func (dmq *DivisionMasterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(divisionmaster.Table, divisionmaster.Columns, sqlgraph.NewFieldSpec(divisionmaster.FieldID, field.TypeInt32))
	_spec.From = dmq.sql
	if unique := dmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dmq.path != nil {
		_spec.Unique = true
	}
	if fields := dmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, divisionmaster.FieldID)
		for i := range fields {
			if fields[i] != divisionmaster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmq *DivisionMasterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmq.driver.Dialect())
	t1 := builder.Table(divisionmaster.Table)
	columns := dmq.ctx.Fields
	if len(columns) == 0 {
		columns = divisionmaster.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmq.sql != nil {
		selector = dmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmq.ctx.Unique != nil && *dmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dmq.predicates {
		p(selector)
	}
	for _, p := range dmq.order {
		p(selector)
	}
	if offset := dmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DivisionMasterGroupBy is the group-by builder for DivisionMaster entities.
type DivisionMasterGroupBy struct {
	selector
	build *DivisionMasterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmgb *DivisionMasterGroupBy) Aggregate(fns ...AggregateFunc) *DivisionMasterGroupBy {
	dmgb.fns = append(dmgb.fns, fns...)
	return dmgb
}

// Scan applies the selector query and scans the result into the given value.
func (dmgb *DivisionMasterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dmgb.build.ctx, "GroupBy")
	if err := dmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionMasterQuery, *DivisionMasterGroupBy](ctx, dmgb.build, dmgb, dmgb.build.inters, v)
}

func (dmgb *DivisionMasterGroupBy) sqlScan(ctx context.Context, root *DivisionMasterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dmgb.fns))
	for _, fn := range dmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dmgb.flds)+len(dmgb.fns))
		for _, f := range *dmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DivisionMasterSelect is the builder for selecting fields of DivisionMaster entities.
type DivisionMasterSelect struct {
	*DivisionMasterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dms *DivisionMasterSelect) Aggregate(fns ...AggregateFunc) *DivisionMasterSelect {
	dms.fns = append(dms.fns, fns...)
	return dms
}

// Scan applies the selector query and scans the result into the given value.
func (dms *DivisionMasterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dms.ctx, "Select")
	if err := dms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DivisionMasterQuery, *DivisionMasterSelect](ctx, dms.DivisionMasterQuery, dms, dms.inters, v)
}

func (dms *DivisionMasterSelect) sqlScan(ctx context.Context, root *DivisionMasterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dms.fns))
	for _, fn := range dms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}