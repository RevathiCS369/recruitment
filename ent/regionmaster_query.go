// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/circlemaster"
	"recruit/ent/divisionmaster"
	"recruit/ent/facility"
	"recruit/ent/predicate"
	"recruit/ent/regionmaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RegionMasterQuery is the builder for querying RegionMaster entities.
type RegionMasterQuery struct {
	config
	ctx              *QueryContext
	order            []regionmaster.OrderOption
	inters           []Interceptor
	predicates       []predicate.RegionMaster
	withCircleRef    *CircleMasterQuery
	withRegions      *DivisionMasterQuery
	withRegionRefRef *FacilityQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RegionMasterQuery builder.
func (rmq *RegionMasterQuery) Where(ps ...predicate.RegionMaster) *RegionMasterQuery {
	rmq.predicates = append(rmq.predicates, ps...)
	return rmq
}

// Limit the number of records to be returned by this query.
func (rmq *RegionMasterQuery) Limit(limit int) *RegionMasterQuery {
	rmq.ctx.Limit = &limit
	return rmq
}

// Offset to start from.
func (rmq *RegionMasterQuery) Offset(offset int) *RegionMasterQuery {
	rmq.ctx.Offset = &offset
	return rmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rmq *RegionMasterQuery) Unique(unique bool) *RegionMasterQuery {
	rmq.ctx.Unique = &unique
	return rmq
}

// Order specifies how the records should be ordered.
func (rmq *RegionMasterQuery) Order(o ...regionmaster.OrderOption) *RegionMasterQuery {
	rmq.order = append(rmq.order, o...)
	return rmq
}

// QueryCircleRef chains the current query on the "circle_ref" edge.
func (rmq *RegionMasterQuery) QueryCircleRef() *CircleMasterQuery {
	query := (&CircleMasterClient{config: rmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(regionmaster.Table, regionmaster.FieldID, selector),
			sqlgraph.To(circlemaster.Table, circlemaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, regionmaster.CircleRefTable, regionmaster.CircleRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(rmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegions chains the current query on the "regions" edge.
func (rmq *RegionMasterQuery) QueryRegions() *DivisionMasterQuery {
	query := (&DivisionMasterClient{config: rmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(regionmaster.Table, regionmaster.FieldID, selector),
			sqlgraph.To(divisionmaster.Table, divisionmaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, regionmaster.RegionsTable, regionmaster.RegionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegionRefRef chains the current query on the "region_ref_ref" edge.
func (rmq *RegionMasterQuery) QueryRegionRefRef() *FacilityQuery {
	query := (&FacilityClient{config: rmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(regionmaster.Table, regionmaster.FieldID, selector),
			sqlgraph.To(facility.Table, facility.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, regionmaster.RegionRefRefTable, regionmaster.RegionRefRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(rmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RegionMaster entity from the query.
// Returns a *NotFoundError when no RegionMaster was found.
func (rmq *RegionMasterQuery) First(ctx context.Context) (*RegionMaster, error) {
	nodes, err := rmq.Limit(1).All(setContextOp(ctx, rmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{regionmaster.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rmq *RegionMasterQuery) FirstX(ctx context.Context) *RegionMaster {
	node, err := rmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RegionMaster ID from the query.
// Returns a *NotFoundError when no RegionMaster ID was found.
func (rmq *RegionMasterQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = rmq.Limit(1).IDs(setContextOp(ctx, rmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{regionmaster.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rmq *RegionMasterQuery) FirstIDX(ctx context.Context) int32 {
	id, err := rmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RegionMaster entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RegionMaster entity is found.
// Returns a *NotFoundError when no RegionMaster entities are found.
func (rmq *RegionMasterQuery) Only(ctx context.Context) (*RegionMaster, error) {
	nodes, err := rmq.Limit(2).All(setContextOp(ctx, rmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{regionmaster.Label}
	default:
		return nil, &NotSingularError{regionmaster.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rmq *RegionMasterQuery) OnlyX(ctx context.Context) *RegionMaster {
	node, err := rmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RegionMaster ID in the query.
// Returns a *NotSingularError when more than one RegionMaster ID is found.
// Returns a *NotFoundError when no entities are found.
func (rmq *RegionMasterQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = rmq.Limit(2).IDs(setContextOp(ctx, rmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{regionmaster.Label}
	default:
		err = &NotSingularError{regionmaster.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rmq *RegionMasterQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := rmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RegionMasters.
func (rmq *RegionMasterQuery) All(ctx context.Context) ([]*RegionMaster, error) {
	ctx = setContextOp(ctx, rmq.ctx, "All")
	if err := rmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RegionMaster, *RegionMasterQuery]()
	return withInterceptors[[]*RegionMaster](ctx, rmq, qr, rmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rmq *RegionMasterQuery) AllX(ctx context.Context) []*RegionMaster {
	nodes, err := rmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RegionMaster IDs.
func (rmq *RegionMasterQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if rmq.ctx.Unique == nil && rmq.path != nil {
		rmq.Unique(true)
	}
	ctx = setContextOp(ctx, rmq.ctx, "IDs")
	if err = rmq.Select(regionmaster.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rmq *RegionMasterQuery) IDsX(ctx context.Context) []int32 {
	ids, err := rmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rmq *RegionMasterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rmq.ctx, "Count")
	if err := rmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rmq, querierCount[*RegionMasterQuery](), rmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rmq *RegionMasterQuery) CountX(ctx context.Context) int {
	count, err := rmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rmq *RegionMasterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rmq.ctx, "Exist")
	switch _, err := rmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rmq *RegionMasterQuery) ExistX(ctx context.Context) bool {
	exist, err := rmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RegionMasterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rmq *RegionMasterQuery) Clone() *RegionMasterQuery {
	if rmq == nil {
		return nil
	}
	return &RegionMasterQuery{
		config:           rmq.config,
		ctx:              rmq.ctx.Clone(),
		order:            append([]regionmaster.OrderOption{}, rmq.order...),
		inters:           append([]Interceptor{}, rmq.inters...),
		predicates:       append([]predicate.RegionMaster{}, rmq.predicates...),
		withCircleRef:    rmq.withCircleRef.Clone(),
		withRegions:      rmq.withRegions.Clone(),
		withRegionRefRef: rmq.withRegionRefRef.Clone(),
		// clone intermediate query.
		sql:  rmq.sql.Clone(),
		path: rmq.path,
	}
}

// WithCircleRef tells the query-builder to eager-load the nodes that are connected to
// the "circle_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (rmq *RegionMasterQuery) WithCircleRef(opts ...func(*CircleMasterQuery)) *RegionMasterQuery {
	query := (&CircleMasterClient{config: rmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rmq.withCircleRef = query
	return rmq
}

// WithRegions tells the query-builder to eager-load the nodes that are connected to
// the "regions" edge. The optional arguments are used to configure the query builder of the edge.
func (rmq *RegionMasterQuery) WithRegions(opts ...func(*DivisionMasterQuery)) *RegionMasterQuery {
	query := (&DivisionMasterClient{config: rmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rmq.withRegions = query
	return rmq
}

// WithRegionRefRef tells the query-builder to eager-load the nodes that are connected to
// the "region_ref_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (rmq *RegionMasterQuery) WithRegionRefRef(opts ...func(*FacilityQuery)) *RegionMasterQuery {
	query := (&FacilityClient{config: rmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rmq.withRegionRefRef = query
	return rmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RegionCode int32 `json:"RegionCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RegionMaster.Query().
//		GroupBy(regionmaster.FieldRegionCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rmq *RegionMasterQuery) GroupBy(field string, fields ...string) *RegionMasterGroupBy {
	rmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RegionMasterGroupBy{build: rmq}
	grbuild.flds = &rmq.ctx.Fields
	grbuild.label = regionmaster.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RegionCode int32 `json:"RegionCode,omitempty"`
//	}
//
//	client.RegionMaster.Query().
//		Select(regionmaster.FieldRegionCode).
//		Scan(ctx, &v)
func (rmq *RegionMasterQuery) Select(fields ...string) *RegionMasterSelect {
	rmq.ctx.Fields = append(rmq.ctx.Fields, fields...)
	sbuild := &RegionMasterSelect{RegionMasterQuery: rmq}
	sbuild.label = regionmaster.Label
	sbuild.flds, sbuild.scan = &rmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RegionMasterSelect configured with the given aggregations.
func (rmq *RegionMasterQuery) Aggregate(fns ...AggregateFunc) *RegionMasterSelect {
	return rmq.Select().Aggregate(fns...)
}

func (rmq *RegionMasterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rmq); err != nil {
				return err
			}
		}
	}
	for _, f := range rmq.ctx.Fields {
		if !regionmaster.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rmq.path != nil {
		prev, err := rmq.path(ctx)
		if err != nil {
			return err
		}
		rmq.sql = prev
	}
	return nil
}

func (rmq *RegionMasterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RegionMaster, error) {
	var (
		nodes       = []*RegionMaster{}
		withFKs     = rmq.withFKs
		_spec       = rmq.querySpec()
		loadedTypes = [3]bool{
			rmq.withCircleRef != nil,
			rmq.withRegions != nil,
			rmq.withRegionRefRef != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, regionmaster.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RegionMaster).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RegionMaster{config: rmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rmq.withCircleRef; query != nil {
		if err := rmq.loadCircleRef(ctx, query, nodes,
			func(n *RegionMaster) { n.Edges.CircleRef = []*CircleMaster{} },
			func(n *RegionMaster, e *CircleMaster) { n.Edges.CircleRef = append(n.Edges.CircleRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := rmq.withRegions; query != nil {
		if err := rmq.loadRegions(ctx, query, nodes,
			func(n *RegionMaster) { n.Edges.Regions = []*DivisionMaster{} },
			func(n *RegionMaster, e *DivisionMaster) { n.Edges.Regions = append(n.Edges.Regions, e) }); err != nil {
			return nil, err
		}
	}
	if query := rmq.withRegionRefRef; query != nil {
		if err := rmq.loadRegionRefRef(ctx, query, nodes,
			func(n *RegionMaster) { n.Edges.RegionRefRef = []*Facility{} },
			func(n *RegionMaster, e *Facility) { n.Edges.RegionRefRef = append(n.Edges.RegionRefRef, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rmq *RegionMasterQuery) loadCircleRef(ctx context.Context, query *CircleMasterQuery, nodes []*RegionMaster, init func(*RegionMaster), assign func(*RegionMaster, *CircleMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*RegionMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CircleMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(regionmaster.CircleRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.region_master_circle_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "region_master_circle_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "region_master_circle_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rmq *RegionMasterQuery) loadRegions(ctx context.Context, query *DivisionMasterQuery, nodes []*RegionMaster, init func(*RegionMaster), assign func(*RegionMaster, *DivisionMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*RegionMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DivisionMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(regionmaster.RegionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.region_master_regions
		if fk == nil {
			return fmt.Errorf(`foreign-key "region_master_regions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "region_master_regions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rmq *RegionMasterQuery) loadRegionRefRef(ctx context.Context, query *FacilityQuery, nodes []*RegionMaster, init func(*RegionMaster), assign func(*RegionMaster, *Facility)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*RegionMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(facility.FieldRegionID)
	}
	query.Where(predicate.Facility(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(regionmaster.RegionRefRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RegionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "RegionID" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rmq *RegionMasterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rmq.querySpec()
	_spec.Node.Columns = rmq.ctx.Fields
	if len(rmq.ctx.Fields) > 0 {
		_spec.Unique = rmq.ctx.Unique != nil && *rmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rmq.driver, _spec)
}

func (rmq *RegionMasterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(regionmaster.Table, regionmaster.Columns, sqlgraph.NewFieldSpec(regionmaster.FieldID, field.TypeInt32))
	_spec.From = rmq.sql
	if unique := rmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rmq.path != nil {
		_spec.Unique = true
	}
	if fields := rmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, regionmaster.FieldID)
		for i := range fields {
			if fields[i] != regionmaster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rmq *RegionMasterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rmq.driver.Dialect())
	t1 := builder.Table(regionmaster.Table)
	columns := rmq.ctx.Fields
	if len(columns) == 0 {
		columns = regionmaster.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rmq.sql != nil {
		selector = rmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rmq.ctx.Unique != nil && *rmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rmq.predicates {
		p(selector)
	}
	for _, p := range rmq.order {
		p(selector)
	}
	if offset := rmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RegionMasterGroupBy is the group-by builder for RegionMaster entities.
type RegionMasterGroupBy struct {
	selector
	build *RegionMasterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rmgb *RegionMasterGroupBy) Aggregate(fns ...AggregateFunc) *RegionMasterGroupBy {
	rmgb.fns = append(rmgb.fns, fns...)
	return rmgb
}

// Scan applies the selector query and scans the result into the given value.
func (rmgb *RegionMasterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rmgb.build.ctx, "GroupBy")
	if err := rmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegionMasterQuery, *RegionMasterGroupBy](ctx, rmgb.build, rmgb, rmgb.build.inters, v)
}

func (rmgb *RegionMasterGroupBy) sqlScan(ctx context.Context, root *RegionMasterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rmgb.fns))
	for _, fn := range rmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rmgb.flds)+len(rmgb.fns))
		for _, f := range *rmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RegionMasterSelect is the builder for selecting fields of RegionMaster entities.
type RegionMasterSelect struct {
	*RegionMasterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rms *RegionMasterSelect) Aggregate(fns ...AggregateFunc) *RegionMasterSelect {
	rms.fns = append(rms.fns, fns...)
	return rms
}

// Scan applies the selector query and scans the result into the given value.
func (rms *RegionMasterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rms.ctx, "Select")
	if err := rms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegionMasterQuery, *RegionMasterSelect](ctx, rms.RegionMasterQuery, rms, rms.inters, v)
}

func (rms *RegionMasterSelect) sqlScan(ctx context.Context, root *RegionMasterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rms.fns))
	for _, fn := range rms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
