// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/exam"
	"recruit/ent/examcalendar"
	"recruit/ent/exampapers"
	"recruit/ent/notification"
	"recruit/ent/predicate"
	"recruit/ent/vacancyyear"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCalendarQuery is the builder for querying ExamCalendar entities.
type ExamCalendarQuery struct {
	config
	ctx           *QueryContext
	order         []examcalendar.OrderOption
	inters        []Interceptor
	predicates    []predicate.ExamCalendar
	withVcyYears  *VacancyYearQuery
	withExams     *ExamQuery
	withPapers    *ExamPapersQuery
	withNotifyRef *NotificationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamCalendarQuery builder.
func (ecq *ExamCalendarQuery) Where(ps ...predicate.ExamCalendar) *ExamCalendarQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *ExamCalendarQuery) Limit(limit int) *ExamCalendarQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *ExamCalendarQuery) Offset(offset int) *ExamCalendarQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *ExamCalendarQuery) Unique(unique bool) *ExamCalendarQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *ExamCalendarQuery) Order(o ...examcalendar.OrderOption) *ExamCalendarQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryVcyYears chains the current query on the "vcy_years" edge.
func (ecq *ExamCalendarQuery) QueryVcyYears() *VacancyYearQuery {
	query := (&VacancyYearClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcalendar.Table, examcalendar.FieldID, selector),
			sqlgraph.To(vacancyyear.Table, vacancyyear.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, examcalendar.VcyYearsTable, examcalendar.VcyYearsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExams chains the current query on the "exams" edge.
func (ecq *ExamCalendarQuery) QueryExams() *ExamQuery {
	query := (&ExamClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcalendar.Table, examcalendar.FieldID, selector),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, examcalendar.ExamsTable, examcalendar.ExamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPapers chains the current query on the "papers" edge.
func (ecq *ExamCalendarQuery) QueryPapers() *ExamPapersQuery {
	query := (&ExamPapersClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcalendar.Table, examcalendar.FieldID, selector),
			sqlgraph.To(exampapers.Table, exampapers.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, examcalendar.PapersTable, examcalendar.PapersColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyRef chains the current query on the "Notify_ref" edge.
func (ecq *ExamCalendarQuery) QueryNotifyRef() *NotificationQuery {
	query := (&NotificationClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcalendar.Table, examcalendar.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcalendar.NotifyRefTable, examcalendar.NotifyRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExamCalendar entity from the query.
// Returns a *NotFoundError when no ExamCalendar was found.
func (ecq *ExamCalendarQuery) First(ctx context.Context) (*ExamCalendar, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examcalendar.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *ExamCalendarQuery) FirstX(ctx context.Context) *ExamCalendar {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamCalendar ID from the query.
// Returns a *NotFoundError when no ExamCalendar ID was found.
func (ecq *ExamCalendarQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examcalendar.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *ExamCalendarQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamCalendar entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamCalendar entity is found.
// Returns a *NotFoundError when no ExamCalendar entities are found.
func (ecq *ExamCalendarQuery) Only(ctx context.Context) (*ExamCalendar, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examcalendar.Label}
	default:
		return nil, &NotSingularError{examcalendar.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *ExamCalendarQuery) OnlyX(ctx context.Context) *ExamCalendar {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamCalendar ID in the query.
// Returns a *NotSingularError when more than one ExamCalendar ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *ExamCalendarQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examcalendar.Label}
	default:
		err = &NotSingularError{examcalendar.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *ExamCalendarQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamCalendars.
func (ecq *ExamCalendarQuery) All(ctx context.Context) ([]*ExamCalendar, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamCalendar, *ExamCalendarQuery]()
	return withInterceptors[[]*ExamCalendar](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *ExamCalendarQuery) AllX(ctx context.Context) []*ExamCalendar {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamCalendar IDs.
func (ecq *ExamCalendarQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(examcalendar.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *ExamCalendarQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *ExamCalendarQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*ExamCalendarQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *ExamCalendarQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *ExamCalendarQuery) Exist(ctx context.Context) (bool, error) {
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
func (ecq *ExamCalendarQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamCalendarQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *ExamCalendarQuery) Clone() *ExamCalendarQuery {
	if ecq == nil {
		return nil
	}
	return &ExamCalendarQuery{
		config:        ecq.config,
		ctx:           ecq.ctx.Clone(),
		order:         append([]examcalendar.OrderOption{}, ecq.order...),
		inters:        append([]Interceptor{}, ecq.inters...),
		predicates:    append([]predicate.ExamCalendar{}, ecq.predicates...),
		withVcyYears:  ecq.withVcyYears.Clone(),
		withExams:     ecq.withExams.Clone(),
		withPapers:    ecq.withPapers.Clone(),
		withNotifyRef: ecq.withNotifyRef.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithVcyYears tells the query-builder to eager-load the nodes that are connected to
// the "vcy_years" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExamCalendarQuery) WithVcyYears(opts ...func(*VacancyYearQuery)) *ExamCalendarQuery {
	query := (&VacancyYearClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withVcyYears = query
	return ecq
}

// WithExams tells the query-builder to eager-load the nodes that are connected to
// the "exams" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExamCalendarQuery) WithExams(opts ...func(*ExamQuery)) *ExamCalendarQuery {
	query := (&ExamClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withExams = query
	return ecq
}

// WithPapers tells the query-builder to eager-load the nodes that are connected to
// the "papers" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExamCalendarQuery) WithPapers(opts ...func(*ExamPapersQuery)) *ExamCalendarQuery {
	query := (&ExamPapersClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withPapers = query
	return ecq
}

// WithNotifyRef tells the query-builder to eager-load the nodes that are connected to
// the "Notify_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExamCalendarQuery) WithNotifyRef(opts ...func(*NotificationQuery)) *ExamCalendarQuery {
	query := (&NotificationClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withNotifyRef = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExamYear int32 `json:"ExamYear,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamCalendar.Query().
//		GroupBy(examcalendar.FieldExamYear).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *ExamCalendarQuery) GroupBy(field string, fields ...string) *ExamCalendarGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamCalendarGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = examcalendar.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExamYear int32 `json:"ExamYear,omitempty"`
//	}
//
//	client.ExamCalendar.Query().
//		Select(examcalendar.FieldExamYear).
//		Scan(ctx, &v)
func (ecq *ExamCalendarQuery) Select(fields ...string) *ExamCalendarSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &ExamCalendarSelect{ExamCalendarQuery: ecq}
	sbuild.label = examcalendar.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamCalendarSelect configured with the given aggregations.
func (ecq *ExamCalendarQuery) Aggregate(fns ...AggregateFunc) *ExamCalendarSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *ExamCalendarQuery) prepareQuery(ctx context.Context) error {
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
		if !examcalendar.ValidColumn(f) {
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

func (ecq *ExamCalendarQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamCalendar, error) {
	var (
		nodes       = []*ExamCalendar{}
		_spec       = ecq.querySpec()
		loadedTypes = [4]bool{
			ecq.withVcyYears != nil,
			ecq.withExams != nil,
			ecq.withPapers != nil,
			ecq.withNotifyRef != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamCalendar).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamCalendar{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := ecq.withVcyYears; query != nil {
		if err := ecq.loadVcyYears(ctx, query, nodes, nil,
			func(n *ExamCalendar, e *VacancyYear) { n.Edges.VcyYears = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withExams; query != nil {
		if err := ecq.loadExams(ctx, query, nodes, nil,
			func(n *ExamCalendar, e *Exam) { n.Edges.Exams = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withPapers; query != nil {
		if err := ecq.loadPapers(ctx, query, nodes, nil,
			func(n *ExamCalendar, e *ExamPapers) { n.Edges.Papers = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withNotifyRef; query != nil {
		if err := ecq.loadNotifyRef(ctx, query, nodes,
			func(n *ExamCalendar) { n.Edges.NotifyRef = []*Notification{} },
			func(n *ExamCalendar, e *Notification) { n.Edges.NotifyRef = append(n.Edges.NotifyRef, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *ExamCalendarQuery) loadVcyYears(ctx context.Context, query *VacancyYearQuery, nodes []*ExamCalendar, init func(*ExamCalendar), assign func(*ExamCalendar, *VacancyYear)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*ExamCalendar)
	for i := range nodes {
		fk := nodes[i].VacancyYearCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(vacancyyear.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "VacancyYearCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecq *ExamCalendarQuery) loadExams(ctx context.Context, query *ExamQuery, nodes []*ExamCalendar, init func(*ExamCalendar), assign func(*ExamCalendar, *Exam)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*ExamCalendar)
	for i := range nodes {
		fk := nodes[i].ExamCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exam.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ExamCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecq *ExamCalendarQuery) loadPapers(ctx context.Context, query *ExamPapersQuery, nodes []*ExamCalendar, init func(*ExamCalendar), assign func(*ExamCalendar, *ExamPapers)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*ExamCalendar)
	for i := range nodes {
		fk := nodes[i].PaperCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exampapers.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "PaperCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecq *ExamCalendarQuery) loadNotifyRef(ctx context.Context, query *NotificationQuery, nodes []*ExamCalendar, init func(*ExamCalendar), assign func(*ExamCalendar, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*ExamCalendar)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcalendar.NotifyRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_calendar_notify_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_calendar_notify_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_calendar_notify_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ecq *ExamCalendarQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *ExamCalendarQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examcalendar.Table, examcalendar.Columns, sqlgraph.NewFieldSpec(examcalendar.FieldID, field.TypeInt32))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examcalendar.FieldID)
		for i := range fields {
			if fields[i] != examcalendar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ecq.withVcyYears != nil {
			_spec.Node.AddColumnOnce(examcalendar.FieldVacancyYearCode)
		}
		if ecq.withExams != nil {
			_spec.Node.AddColumnOnce(examcalendar.FieldExamCode)
		}
		if ecq.withPapers != nil {
			_spec.Node.AddColumnOnce(examcalendar.FieldPaperCode)
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

func (ecq *ExamCalendarQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(examcalendar.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = examcalendar.Columns
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

// ExamCalendarGroupBy is the group-by builder for ExamCalendar entities.
type ExamCalendarGroupBy struct {
	selector
	build *ExamCalendarQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *ExamCalendarGroupBy) Aggregate(fns ...AggregateFunc) *ExamCalendarGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *ExamCalendarGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCalendarQuery, *ExamCalendarGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *ExamCalendarGroupBy) sqlScan(ctx context.Context, root *ExamCalendarQuery, v any) error {
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

// ExamCalendarSelect is the builder for selecting fields of ExamCalendar entities.
type ExamCalendarSelect struct {
	*ExamCalendarQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *ExamCalendarSelect) Aggregate(fns ...AggregateFunc) *ExamCalendarSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *ExamCalendarSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCalendarQuery, *ExamCalendarSelect](ctx, ecs.ExamCalendarQuery, ecs, ecs.inters, v)
}

func (ecs *ExamCalendarSelect) sqlScan(ctx context.Context, root *ExamCalendarQuery, v any) error {
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
