// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/predicate"
	"recruit/ent/recommendationsipapplications"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecommendationsIPApplicationsQuery is the builder for querying RecommendationsIPApplications entities.
type RecommendationsIPApplicationsQuery struct {
	config
	ctx          *QueryContext
	order        []recommendationsipapplications.OrderOption
	inters       []Interceptor
	predicates   []predicate.RecommendationsIPApplications
	withApplnRef *ExamApplicationsIPQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecommendationsIPApplicationsQuery builder.
func (riaq *RecommendationsIPApplicationsQuery) Where(ps ...predicate.RecommendationsIPApplications) *RecommendationsIPApplicationsQuery {
	riaq.predicates = append(riaq.predicates, ps...)
	return riaq
}

// Limit the number of records to be returned by this query.
func (riaq *RecommendationsIPApplicationsQuery) Limit(limit int) *RecommendationsIPApplicationsQuery {
	riaq.ctx.Limit = &limit
	return riaq
}

// Offset to start from.
func (riaq *RecommendationsIPApplicationsQuery) Offset(offset int) *RecommendationsIPApplicationsQuery {
	riaq.ctx.Offset = &offset
	return riaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (riaq *RecommendationsIPApplicationsQuery) Unique(unique bool) *RecommendationsIPApplicationsQuery {
	riaq.ctx.Unique = &unique
	return riaq
}

// Order specifies how the records should be ordered.
func (riaq *RecommendationsIPApplicationsQuery) Order(o ...recommendationsipapplications.OrderOption) *RecommendationsIPApplicationsQuery {
	riaq.order = append(riaq.order, o...)
	return riaq
}

// QueryApplnRef chains the current query on the "ApplnRef" edge.
func (riaq *RecommendationsIPApplicationsQuery) QueryApplnRef() *ExamApplicationsIPQuery {
	query := (&ExamApplicationsIPClient{config: riaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := riaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := riaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recommendationsipapplications.Table, recommendationsipapplications.FieldID, selector),
			sqlgraph.To(exam_applications_ip.Table, exam_applications_ip.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recommendationsipapplications.ApplnRefTable, recommendationsipapplications.ApplnRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(riaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RecommendationsIPApplications entity from the query.
// Returns a *NotFoundError when no RecommendationsIPApplications was found.
func (riaq *RecommendationsIPApplicationsQuery) First(ctx context.Context) (*RecommendationsIPApplications, error) {
	nodes, err := riaq.Limit(1).All(setContextOp(ctx, riaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recommendationsipapplications.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) FirstX(ctx context.Context) *RecommendationsIPApplications {
	node, err := riaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecommendationsIPApplications ID from the query.
// Returns a *NotFoundError when no RecommendationsIPApplications ID was found.
func (riaq *RecommendationsIPApplicationsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = riaq.Limit(1).IDs(setContextOp(ctx, riaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recommendationsipapplications.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := riaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecommendationsIPApplications entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecommendationsIPApplications entity is found.
// Returns a *NotFoundError when no RecommendationsIPApplications entities are found.
func (riaq *RecommendationsIPApplicationsQuery) Only(ctx context.Context) (*RecommendationsIPApplications, error) {
	nodes, err := riaq.Limit(2).All(setContextOp(ctx, riaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recommendationsipapplications.Label}
	default:
		return nil, &NotSingularError{recommendationsipapplications.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) OnlyX(ctx context.Context) *RecommendationsIPApplications {
	node, err := riaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecommendationsIPApplications ID in the query.
// Returns a *NotSingularError when more than one RecommendationsIPApplications ID is found.
// Returns a *NotFoundError when no entities are found.
func (riaq *RecommendationsIPApplicationsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = riaq.Limit(2).IDs(setContextOp(ctx, riaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recommendationsipapplications.Label}
	default:
		err = &NotSingularError{recommendationsipapplications.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := riaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecommendationsIPApplicationsSlice.
func (riaq *RecommendationsIPApplicationsQuery) All(ctx context.Context) ([]*RecommendationsIPApplications, error) {
	ctx = setContextOp(ctx, riaq.ctx, "All")
	if err := riaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RecommendationsIPApplications, *RecommendationsIPApplicationsQuery]()
	return withInterceptors[[]*RecommendationsIPApplications](ctx, riaq, qr, riaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) AllX(ctx context.Context) []*RecommendationsIPApplications {
	nodes, err := riaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecommendationsIPApplications IDs.
func (riaq *RecommendationsIPApplicationsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if riaq.ctx.Unique == nil && riaq.path != nil {
		riaq.Unique(true)
	}
	ctx = setContextOp(ctx, riaq.ctx, "IDs")
	if err = riaq.Select(recommendationsipapplications.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := riaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (riaq *RecommendationsIPApplicationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, riaq.ctx, "Count")
	if err := riaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, riaq, querierCount[*RecommendationsIPApplicationsQuery](), riaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) CountX(ctx context.Context) int {
	count, err := riaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (riaq *RecommendationsIPApplicationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, riaq.ctx, "Exist")
	switch _, err := riaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (riaq *RecommendationsIPApplicationsQuery) ExistX(ctx context.Context) bool {
	exist, err := riaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecommendationsIPApplicationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (riaq *RecommendationsIPApplicationsQuery) Clone() *RecommendationsIPApplicationsQuery {
	if riaq == nil {
		return nil
	}
	return &RecommendationsIPApplicationsQuery{
		config:       riaq.config,
		ctx:          riaq.ctx.Clone(),
		order:        append([]recommendationsipapplications.OrderOption{}, riaq.order...),
		inters:       append([]Interceptor{}, riaq.inters...),
		predicates:   append([]predicate.RecommendationsIPApplications{}, riaq.predicates...),
		withApplnRef: riaq.withApplnRef.Clone(),
		// clone intermediate query.
		sql:  riaq.sql.Clone(),
		path: riaq.path,
	}
}

// WithApplnRef tells the query-builder to eager-load the nodes that are connected to
// the "ApplnRef" edge. The optional arguments are used to configure the query builder of the edge.
func (riaq *RecommendationsIPApplicationsQuery) WithApplnRef(opts ...func(*ExamApplicationsIPQuery)) *RecommendationsIPApplicationsQuery {
	query := (&ExamApplicationsIPClient{config: riaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	riaq.withApplnRef = query
	return riaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplicationID int64 `json:"ApplicationID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RecommendationsIPApplications.Query().
//		GroupBy(recommendationsipapplications.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (riaq *RecommendationsIPApplicationsQuery) GroupBy(field string, fields ...string) *RecommendationsIPApplicationsGroupBy {
	riaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RecommendationsIPApplicationsGroupBy{build: riaq}
	grbuild.flds = &riaq.ctx.Fields
	grbuild.label = recommendationsipapplications.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ApplicationID int64 `json:"ApplicationID,omitempty"`
//	}
//
//	client.RecommendationsIPApplications.Query().
//		Select(recommendationsipapplications.FieldApplicationID).
//		Scan(ctx, &v)
func (riaq *RecommendationsIPApplicationsQuery) Select(fields ...string) *RecommendationsIPApplicationsSelect {
	riaq.ctx.Fields = append(riaq.ctx.Fields, fields...)
	sbuild := &RecommendationsIPApplicationsSelect{RecommendationsIPApplicationsQuery: riaq}
	sbuild.label = recommendationsipapplications.Label
	sbuild.flds, sbuild.scan = &riaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RecommendationsIPApplicationsSelect configured with the given aggregations.
func (riaq *RecommendationsIPApplicationsQuery) Aggregate(fns ...AggregateFunc) *RecommendationsIPApplicationsSelect {
	return riaq.Select().Aggregate(fns...)
}

func (riaq *RecommendationsIPApplicationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range riaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, riaq); err != nil {
				return err
			}
		}
	}
	for _, f := range riaq.ctx.Fields {
		if !recommendationsipapplications.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if riaq.path != nil {
		prev, err := riaq.path(ctx)
		if err != nil {
			return err
		}
		riaq.sql = prev
	}
	return nil
}

func (riaq *RecommendationsIPApplicationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecommendationsIPApplications, error) {
	var (
		nodes       = []*RecommendationsIPApplications{}
		_spec       = riaq.querySpec()
		loadedTypes = [1]bool{
			riaq.withApplnRef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecommendationsIPApplications).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecommendationsIPApplications{config: riaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, riaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := riaq.withApplnRef; query != nil {
		if err := riaq.loadApplnRef(ctx, query, nodes, nil,
			func(n *RecommendationsIPApplications, e *Exam_Applications_IP) { n.Edges.ApplnRef = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (riaq *RecommendationsIPApplicationsQuery) loadApplnRef(ctx context.Context, query *ExamApplicationsIPQuery, nodes []*RecommendationsIPApplications, init func(*RecommendationsIPApplications), assign func(*RecommendationsIPApplications, *Exam_Applications_IP)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RecommendationsIPApplications)
	for i := range nodes {
		fk := nodes[i].ApplicationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exam_applications_ip.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ApplicationID" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (riaq *RecommendationsIPApplicationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := riaq.querySpec()
	_spec.Node.Columns = riaq.ctx.Fields
	if len(riaq.ctx.Fields) > 0 {
		_spec.Unique = riaq.ctx.Unique != nil && *riaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, riaq.driver, _spec)
}

func (riaq *RecommendationsIPApplicationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(recommendationsipapplications.Table, recommendationsipapplications.Columns, sqlgraph.NewFieldSpec(recommendationsipapplications.FieldID, field.TypeInt64))
	_spec.From = riaq.sql
	if unique := riaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if riaq.path != nil {
		_spec.Unique = true
	}
	if fields := riaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recommendationsipapplications.FieldID)
		for i := range fields {
			if fields[i] != recommendationsipapplications.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if riaq.withApplnRef != nil {
			_spec.Node.AddColumnOnce(recommendationsipapplications.FieldApplicationID)
		}
	}
	if ps := riaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := riaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := riaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := riaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (riaq *RecommendationsIPApplicationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(riaq.driver.Dialect())
	t1 := builder.Table(recommendationsipapplications.Table)
	columns := riaq.ctx.Fields
	if len(columns) == 0 {
		columns = recommendationsipapplications.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if riaq.sql != nil {
		selector = riaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if riaq.ctx.Unique != nil && *riaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range riaq.predicates {
		p(selector)
	}
	for _, p := range riaq.order {
		p(selector)
	}
	if offset := riaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := riaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecommendationsIPApplicationsGroupBy is the group-by builder for RecommendationsIPApplications entities.
type RecommendationsIPApplicationsGroupBy struct {
	selector
	build *RecommendationsIPApplicationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (riagb *RecommendationsIPApplicationsGroupBy) Aggregate(fns ...AggregateFunc) *RecommendationsIPApplicationsGroupBy {
	riagb.fns = append(riagb.fns, fns...)
	return riagb
}

// Scan applies the selector query and scans the result into the given value.
func (riagb *RecommendationsIPApplicationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, riagb.build.ctx, "GroupBy")
	if err := riagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecommendationsIPApplicationsQuery, *RecommendationsIPApplicationsGroupBy](ctx, riagb.build, riagb, riagb.build.inters, v)
}

func (riagb *RecommendationsIPApplicationsGroupBy) sqlScan(ctx context.Context, root *RecommendationsIPApplicationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(riagb.fns))
	for _, fn := range riagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*riagb.flds)+len(riagb.fns))
		for _, f := range *riagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*riagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := riagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RecommendationsIPApplicationsSelect is the builder for selecting fields of RecommendationsIPApplications entities.
type RecommendationsIPApplicationsSelect struct {
	*RecommendationsIPApplicationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rias *RecommendationsIPApplicationsSelect) Aggregate(fns ...AggregateFunc) *RecommendationsIPApplicationsSelect {
	rias.fns = append(rias.fns, fns...)
	return rias
}

// Scan applies the selector query and scans the result into the given value.
func (rias *RecommendationsIPApplicationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rias.ctx, "Select")
	if err := rias.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecommendationsIPApplicationsQuery, *RecommendationsIPApplicationsSelect](ctx, rias.RecommendationsIPApplicationsQuery, rias, rias.inters, v)
}

func (rias *RecommendationsIPApplicationsSelect) sqlScan(ctx context.Context, root *RecommendationsIPApplicationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rias.fns))
	for _, fn := range rias.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rias.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rias.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
