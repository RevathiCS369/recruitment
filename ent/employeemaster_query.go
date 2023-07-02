// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/employeemaster"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/predicate"
	"recruit/ent/usermaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeMasterQuery is the builder for querying EmployeeMaster entities.
type EmployeeMasterQuery struct {
	config
	ctx               *QueryContext
	order             []employeemaster.OrderOption
	inters            []Interceptor
	predicates        []predicate.EmployeeMaster
	withUsermasterRef *UserMasterQuery
	withEmpRef        *ExamApplicationsPSQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeMasterQuery builder.
func (emq *EmployeeMasterQuery) Where(ps ...predicate.EmployeeMaster) *EmployeeMasterQuery {
	emq.predicates = append(emq.predicates, ps...)
	return emq
}

// Limit the number of records to be returned by this query.
func (emq *EmployeeMasterQuery) Limit(limit int) *EmployeeMasterQuery {
	emq.ctx.Limit = &limit
	return emq
}

// Offset to start from.
func (emq *EmployeeMasterQuery) Offset(offset int) *EmployeeMasterQuery {
	emq.ctx.Offset = &offset
	return emq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (emq *EmployeeMasterQuery) Unique(unique bool) *EmployeeMasterQuery {
	emq.ctx.Unique = &unique
	return emq
}

// Order specifies how the records should be ordered.
func (emq *EmployeeMasterQuery) Order(o ...employeemaster.OrderOption) *EmployeeMasterQuery {
	emq.order = append(emq.order, o...)
	return emq
}

// QueryUsermasterRef chains the current query on the "UsermasterRef" edge.
func (emq *EmployeeMasterQuery) QueryUsermasterRef() *UserMasterQuery {
	query := (&UserMasterClient{config: emq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := emq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := emq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employeemaster.Table, employeemaster.FieldID, selector),
			sqlgraph.To(usermaster.Table, usermaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employeemaster.UsermasterRefTable, employeemaster.UsermasterRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(emq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmpRef chains the current query on the "Emp_Ref" edge.
func (emq *EmployeeMasterQuery) QueryEmpRef() *ExamApplicationsPSQuery {
	query := (&ExamApplicationsPSClient{config: emq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := emq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := emq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(employeemaster.Table, employeemaster.FieldID, selector),
			sqlgraph.To(exam_applications_ps.Table, exam_applications_ps.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employeemaster.EmpRefTable, employeemaster.EmpRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(emq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmployeeMaster entity from the query.
// Returns a *NotFoundError when no EmployeeMaster was found.
func (emq *EmployeeMasterQuery) First(ctx context.Context) (*EmployeeMaster, error) {
	nodes, err := emq.Limit(1).All(setContextOp(ctx, emq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employeemaster.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (emq *EmployeeMasterQuery) FirstX(ctx context.Context) *EmployeeMaster {
	node, err := emq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmployeeMaster ID from the query.
// Returns a *NotFoundError when no EmployeeMaster ID was found.
func (emq *EmployeeMasterQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = emq.Limit(1).IDs(setContextOp(ctx, emq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employeemaster.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (emq *EmployeeMasterQuery) FirstIDX(ctx context.Context) int64 {
	id, err := emq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmployeeMaster entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmployeeMaster entity is found.
// Returns a *NotFoundError when no EmployeeMaster entities are found.
func (emq *EmployeeMasterQuery) Only(ctx context.Context) (*EmployeeMaster, error) {
	nodes, err := emq.Limit(2).All(setContextOp(ctx, emq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employeemaster.Label}
	default:
		return nil, &NotSingularError{employeemaster.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (emq *EmployeeMasterQuery) OnlyX(ctx context.Context) *EmployeeMaster {
	node, err := emq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmployeeMaster ID in the query.
// Returns a *NotSingularError when more than one EmployeeMaster ID is found.
// Returns a *NotFoundError when no entities are found.
func (emq *EmployeeMasterQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = emq.Limit(2).IDs(setContextOp(ctx, emq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employeemaster.Label}
	default:
		err = &NotSingularError{employeemaster.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (emq *EmployeeMasterQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := emq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmployeeMasters.
func (emq *EmployeeMasterQuery) All(ctx context.Context) ([]*EmployeeMaster, error) {
	ctx = setContextOp(ctx, emq.ctx, "All")
	if err := emq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmployeeMaster, *EmployeeMasterQuery]()
	return withInterceptors[[]*EmployeeMaster](ctx, emq, qr, emq.inters)
}

// AllX is like All, but panics if an error occurs.
func (emq *EmployeeMasterQuery) AllX(ctx context.Context) []*EmployeeMaster {
	nodes, err := emq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmployeeMaster IDs.
func (emq *EmployeeMasterQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if emq.ctx.Unique == nil && emq.path != nil {
		emq.Unique(true)
	}
	ctx = setContextOp(ctx, emq.ctx, "IDs")
	if err = emq.Select(employeemaster.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (emq *EmployeeMasterQuery) IDsX(ctx context.Context) []int64 {
	ids, err := emq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (emq *EmployeeMasterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, emq.ctx, "Count")
	if err := emq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, emq, querierCount[*EmployeeMasterQuery](), emq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (emq *EmployeeMasterQuery) CountX(ctx context.Context) int {
	count, err := emq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (emq *EmployeeMasterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, emq.ctx, "Exist")
	switch _, err := emq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (emq *EmployeeMasterQuery) ExistX(ctx context.Context) bool {
	exist, err := emq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeMasterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (emq *EmployeeMasterQuery) Clone() *EmployeeMasterQuery {
	if emq == nil {
		return nil
	}
	return &EmployeeMasterQuery{
		config:            emq.config,
		ctx:               emq.ctx.Clone(),
		order:             append([]employeemaster.OrderOption{}, emq.order...),
		inters:            append([]Interceptor{}, emq.inters...),
		predicates:        append([]predicate.EmployeeMaster{}, emq.predicates...),
		withUsermasterRef: emq.withUsermasterRef.Clone(),
		withEmpRef:        emq.withEmpRef.Clone(),
		// clone intermediate query.
		sql:  emq.sql.Clone(),
		path: emq.path,
	}
}

// WithUsermasterRef tells the query-builder to eager-load the nodes that are connected to
// the "UsermasterRef" edge. The optional arguments are used to configure the query builder of the edge.
func (emq *EmployeeMasterQuery) WithUsermasterRef(opts ...func(*UserMasterQuery)) *EmployeeMasterQuery {
	query := (&UserMasterClient{config: emq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	emq.withUsermasterRef = query
	return emq
}

// WithEmpRef tells the query-builder to eager-load the nodes that are connected to
// the "Emp_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (emq *EmployeeMasterQuery) WithEmpRef(opts ...func(*ExamApplicationsPSQuery)) *EmployeeMasterQuery {
	query := (&ExamApplicationsPSClient{config: emq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	emq.withEmpRef = query
	return emq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EmployeeID int64 `json:"EmployeeID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmployeeMaster.Query().
//		GroupBy(employeemaster.FieldEmployeeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (emq *EmployeeMasterQuery) GroupBy(field string, fields ...string) *EmployeeMasterGroupBy {
	emq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmployeeMasterGroupBy{build: emq}
	grbuild.flds = &emq.ctx.Fields
	grbuild.label = employeemaster.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EmployeeID int64 `json:"EmployeeID,omitempty"`
//	}
//
//	client.EmployeeMaster.Query().
//		Select(employeemaster.FieldEmployeeID).
//		Scan(ctx, &v)
func (emq *EmployeeMasterQuery) Select(fields ...string) *EmployeeMasterSelect {
	emq.ctx.Fields = append(emq.ctx.Fields, fields...)
	sbuild := &EmployeeMasterSelect{EmployeeMasterQuery: emq}
	sbuild.label = employeemaster.Label
	sbuild.flds, sbuild.scan = &emq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmployeeMasterSelect configured with the given aggregations.
func (emq *EmployeeMasterQuery) Aggregate(fns ...AggregateFunc) *EmployeeMasterSelect {
	return emq.Select().Aggregate(fns...)
}

func (emq *EmployeeMasterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range emq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, emq); err != nil {
				return err
			}
		}
	}
	for _, f := range emq.ctx.Fields {
		if !employeemaster.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if emq.path != nil {
		prev, err := emq.path(ctx)
		if err != nil {
			return err
		}
		emq.sql = prev
	}
	return nil
}

func (emq *EmployeeMasterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmployeeMaster, error) {
	var (
		nodes       = []*EmployeeMaster{}
		withFKs     = emq.withFKs
		_spec       = emq.querySpec()
		loadedTypes = [2]bool{
			emq.withUsermasterRef != nil,
			emq.withEmpRef != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, employeemaster.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmployeeMaster).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmployeeMaster{config: emq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, emq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := emq.withUsermasterRef; query != nil {
		if err := emq.loadUsermasterRef(ctx, query, nodes,
			func(n *EmployeeMaster) { n.Edges.UsermasterRef = []*UserMaster{} },
			func(n *EmployeeMaster, e *UserMaster) { n.Edges.UsermasterRef = append(n.Edges.UsermasterRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := emq.withEmpRef; query != nil {
		if err := emq.loadEmpRef(ctx, query, nodes,
			func(n *EmployeeMaster) { n.Edges.EmpRef = []*Exam_Applications_PS{} },
			func(n *EmployeeMaster, e *Exam_Applications_PS) { n.Edges.EmpRef = append(n.Edges.EmpRef, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (emq *EmployeeMasterQuery) loadUsermasterRef(ctx context.Context, query *UserMasterQuery, nodes []*EmployeeMaster, init func(*EmployeeMaster), assign func(*EmployeeMaster, *UserMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*EmployeeMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(employeemaster.UsermasterRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.employee_master_usermaster_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "employee_master_usermaster_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "employee_master_usermaster_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (emq *EmployeeMasterQuery) loadEmpRef(ctx context.Context, query *ExamApplicationsPSQuery, nodes []*EmployeeMaster, init func(*EmployeeMaster), assign func(*EmployeeMaster, *Exam_Applications_PS)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*EmployeeMaster)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam_Applications_PS(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(employeemaster.EmpRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.employee_master_emp_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "employee_master_emp_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "employee_master_emp_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (emq *EmployeeMasterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := emq.querySpec()
	_spec.Node.Columns = emq.ctx.Fields
	if len(emq.ctx.Fields) > 0 {
		_spec.Unique = emq.ctx.Unique != nil && *emq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, emq.driver, _spec)
}

func (emq *EmployeeMasterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(employeemaster.Table, employeemaster.Columns, sqlgraph.NewFieldSpec(employeemaster.FieldID, field.TypeInt64))
	_spec.From = emq.sql
	if unique := emq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if emq.path != nil {
		_spec.Unique = true
	}
	if fields := emq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeemaster.FieldID)
		for i := range fields {
			if fields[i] != employeemaster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := emq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := emq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := emq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := emq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (emq *EmployeeMasterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(emq.driver.Dialect())
	t1 := builder.Table(employeemaster.Table)
	columns := emq.ctx.Fields
	if len(columns) == 0 {
		columns = employeemaster.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if emq.sql != nil {
		selector = emq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if emq.ctx.Unique != nil && *emq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range emq.predicates {
		p(selector)
	}
	for _, p := range emq.order {
		p(selector)
	}
	if offset := emq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := emq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmployeeMasterGroupBy is the group-by builder for EmployeeMaster entities.
type EmployeeMasterGroupBy struct {
	selector
	build *EmployeeMasterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (emgb *EmployeeMasterGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeMasterGroupBy {
	emgb.fns = append(emgb.fns, fns...)
	return emgb
}

// Scan applies the selector query and scans the result into the given value.
func (emgb *EmployeeMasterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, emgb.build.ctx, "GroupBy")
	if err := emgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeMasterQuery, *EmployeeMasterGroupBy](ctx, emgb.build, emgb, emgb.build.inters, v)
}

func (emgb *EmployeeMasterGroupBy) sqlScan(ctx context.Context, root *EmployeeMasterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(emgb.fns))
	for _, fn := range emgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*emgb.flds)+len(emgb.fns))
		for _, f := range *emgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*emgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := emgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmployeeMasterSelect is the builder for selecting fields of EmployeeMaster entities.
type EmployeeMasterSelect struct {
	*EmployeeMasterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ems *EmployeeMasterSelect) Aggregate(fns ...AggregateFunc) *EmployeeMasterSelect {
	ems.fns = append(ems.fns, fns...)
	return ems
}

// Scan applies the selector query and scans the result into the given value.
func (ems *EmployeeMasterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ems.ctx, "Select")
	if err := ems.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmployeeMasterQuery, *EmployeeMasterSelect](ctx, ems.EmployeeMasterQuery, ems, ems.inters, v)
}

func (ems *EmployeeMasterSelect) sqlScan(ctx context.Context, root *EmployeeMasterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ems.fns))
	for _, fn := range ems.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ems.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ems.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
