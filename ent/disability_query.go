// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/disability"
	"recruit/ent/exampapers"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DisabilityQuery is the builder for querying Disability entities.
type DisabilityQuery struct {
	config
	ctx        *QueryContext
	order      []disability.OrderOption
	inters     []Interceptor
	predicates []predicate.Disability
	withDisRef *ExamPapersQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DisabilityQuery builder.
func (dq *DisabilityQuery) Where(ps ...predicate.Disability) *DisabilityQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DisabilityQuery) Limit(limit int) *DisabilityQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DisabilityQuery) Offset(offset int) *DisabilityQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DisabilityQuery) Unique(unique bool) *DisabilityQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DisabilityQuery) Order(o ...disability.OrderOption) *DisabilityQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDisRef chains the current query on the "dis_ref" edge.
func (dq *DisabilityQuery) QueryDisRef() *ExamPapersQuery {
	query := (&ExamPapersClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(disability.Table, disability.FieldID, selector),
			sqlgraph.To(exampapers.Table, exampapers.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, disability.DisRefTable, disability.DisRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Disability entity from the query.
// Returns a *NotFoundError when no Disability was found.
func (dq *DisabilityQuery) First(ctx context.Context) (*Disability, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{disability.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DisabilityQuery) FirstX(ctx context.Context) *Disability {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Disability ID from the query.
// Returns a *NotFoundError when no Disability ID was found.
func (dq *DisabilityQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{disability.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DisabilityQuery) FirstIDX(ctx context.Context) int32 {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Disability entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Disability entity is found.
// Returns a *NotFoundError when no Disability entities are found.
func (dq *DisabilityQuery) Only(ctx context.Context) (*Disability, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{disability.Label}
	default:
		return nil, &NotSingularError{disability.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DisabilityQuery) OnlyX(ctx context.Context) *Disability {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Disability ID in the query.
// Returns a *NotSingularError when more than one Disability ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DisabilityQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{disability.Label}
	default:
		err = &NotSingularError{disability.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DisabilityQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Disabilities.
func (dq *DisabilityQuery) All(ctx context.Context) ([]*Disability, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Disability, *DisabilityQuery]()
	return withInterceptors[[]*Disability](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DisabilityQuery) AllX(ctx context.Context) []*Disability {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Disability IDs.
func (dq *DisabilityQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(disability.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DisabilityQuery) IDsX(ctx context.Context) []int32 {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DisabilityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DisabilityQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DisabilityQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DisabilityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DisabilityQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DisabilityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DisabilityQuery) Clone() *DisabilityQuery {
	if dq == nil {
		return nil
	}
	return &DisabilityQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]disability.OrderOption{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Disability{}, dq.predicates...),
		withDisRef: dq.withDisRef.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDisRef tells the query-builder to eager-load the nodes that are connected to
// the "dis_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DisabilityQuery) WithDisRef(opts ...func(*ExamPapersQuery)) *DisabilityQuery {
	query := (&ExamPapersClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDisRef = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DisabilityTypeCode string `json:"DisabilityTypeCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Disability.Query().
//		GroupBy(disability.FieldDisabilityTypeCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DisabilityQuery) GroupBy(field string, fields ...string) *DisabilityGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DisabilityGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = disability.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DisabilityTypeCode string `json:"DisabilityTypeCode,omitempty"`
//	}
//
//	client.Disability.Query().
//		Select(disability.FieldDisabilityTypeCode).
//		Scan(ctx, &v)
func (dq *DisabilityQuery) Select(fields ...string) *DisabilitySelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DisabilitySelect{DisabilityQuery: dq}
	sbuild.label = disability.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DisabilitySelect configured with the given aggregations.
func (dq *DisabilityQuery) Aggregate(fns ...AggregateFunc) *DisabilitySelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DisabilityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !disability.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DisabilityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Disability, error) {
	var (
		nodes       = []*Disability{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDisRef != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, disability.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Disability).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Disability{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDisRef; query != nil {
		if err := dq.loadDisRef(ctx, query, nodes,
			func(n *Disability) { n.Edges.DisRef = []*ExamPapers{} },
			func(n *Disability, e *ExamPapers) { n.Edges.DisRef = append(n.Edges.DisRef, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DisabilityQuery) loadDisRef(ctx context.Context, query *ExamPapersQuery, nodes []*Disability, init func(*Disability), assign func(*Disability, *ExamPapers)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Disability)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ExamPapers(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(disability.DisRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.disability_dis_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "disability_dis_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "disability_dis_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DisabilityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DisabilityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(disability.Table, disability.Columns, sqlgraph.NewFieldSpec(disability.FieldID, field.TypeInt32))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, disability.FieldID)
		for i := range fields {
			if fields[i] != disability.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DisabilityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(disability.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = disability.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DisabilityGroupBy is the group-by builder for Disability entities.
type DisabilityGroupBy struct {
	selector
	build *DisabilityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DisabilityGroupBy) Aggregate(fns ...AggregateFunc) *DisabilityGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DisabilityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DisabilityQuery, *DisabilityGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DisabilityGroupBy) sqlScan(ctx context.Context, root *DisabilityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DisabilitySelect is the builder for selecting fields of Disability entities.
type DisabilitySelect struct {
	*DisabilityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DisabilitySelect) Aggregate(fns ...AggregateFunc) *DisabilitySelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DisabilitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DisabilityQuery, *DisabilitySelect](ctx, ds.DisabilityQuery, ds, ds.inters, v)
}

func (ds *DisabilitySelect) sqlScan(ctx context.Context, root *DisabilityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
