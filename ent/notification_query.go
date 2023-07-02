// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/application"
	"recruit/ent/center"
	"recruit/ent/exam"
	"recruit/ent/exam_ip"
	"recruit/ent/exam_ps"
	"recruit/ent/nodalofficer"
	"recruit/ent/notification"
	"recruit/ent/predicate"
	"recruit/ent/vacancyyear"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationQuery is the builder for querying Notification entities.
type NotificationQuery struct {
	config
	ctx                 *QueryContext
	order               []notification.OrderOption
	inters              []Interceptor
	predicates          []predicate.Notification
	withApplications    *ApplicationQuery
	withCenters         *CenterQuery
	withNodalOfficers   *NodalOfficerQuery
	withExam            *ExamQuery
	withVacancyYears    *VacancyYearQuery
	withNotifyRef       *NotificationQuery
	withNotificationsPs *ExamPSQuery
	withNotificationsIP *ExamIPQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationQuery builder.
func (nq *NotificationQuery) Where(ps ...predicate.Notification) *NotificationQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NotificationQuery) Limit(limit int) *NotificationQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NotificationQuery) Offset(offset int) *NotificationQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NotificationQuery) Unique(unique bool) *NotificationQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NotificationQuery) Order(o ...notification.OrderOption) *NotificationQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryApplications chains the current query on the "applications" edge.
func (nq *NotificationQuery) QueryApplications() *ApplicationQuery {
	query := (&ApplicationClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.ApplicationsTable, notification.ApplicationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCenters chains the current query on the "centers" edge.
func (nq *NotificationQuery) QueryCenters() *CenterQuery {
	query := (&CenterClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(center.Table, center.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.CentersTable, notification.CentersColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodalOfficers chains the current query on the "nodal_officers" edge.
func (nq *NotificationQuery) QueryNodalOfficers() *NodalOfficerQuery {
	query := (&NodalOfficerClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(nodalofficer.Table, nodalofficer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.NodalOfficersTable, notification.NodalOfficersColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExam chains the current query on the "exam" edge.
func (nq *NotificationQuery) QueryExam() *ExamQuery {
	query := (&ExamClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notification.ExamTable, notification.ExamColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVacancyYears chains the current query on the "vacancy_years" edge.
func (nq *NotificationQuery) QueryVacancyYears() *VacancyYearQuery {
	query := (&VacancyYearClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(vacancyyear.Table, vacancyyear.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.VacancyYearsTable, notification.VacancyYearsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyRef chains the current query on the "notify_ref" edge.
func (nq *NotificationQuery) QueryNotifyRef() *NotificationQuery {
	query := (&NotificationClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, notification.NotifyRefTable, notification.NotifyRefPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotificationsPs chains the current query on the "notifications_ps" edge.
func (nq *NotificationQuery) QueryNotificationsPs() *ExamPSQuery {
	query := (&ExamPSClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(exam_ps.Table, exam_ps.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.NotificationsPsTable, notification.NotificationsPsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotificationsIP chains the current query on the "notifications_ip" edge.
func (nq *NotificationQuery) QueryNotificationsIP() *ExamIPQuery {
	query := (&ExamIPClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(exam_ip.Table, exam_ip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notification.NotificationsIPTable, notification.NotificationsIPColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Notification entity from the query.
// Returns a *NotFoundError when no Notification was found.
func (nq *NotificationQuery) First(ctx context.Context) (*Notification, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NotificationQuery) FirstX(ctx context.Context) *Notification {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Notification ID from the query.
// Returns a *NotFoundError when no Notification ID was found.
func (nq *NotificationQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NotificationQuery) FirstIDX(ctx context.Context) int32 {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Notification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Notification entity is found.
// Returns a *NotFoundError when no Notification entities are found.
func (nq *NotificationQuery) Only(ctx context.Context) (*Notification, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notification.Label}
	default:
		return nil, &NotSingularError{notification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NotificationQuery) OnlyX(ctx context.Context) *Notification {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Notification ID in the query.
// Returns a *NotSingularError when more than one Notification ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NotificationQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notification.Label}
	default:
		err = &NotSingularError{notification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NotificationQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Notifications.
func (nq *NotificationQuery) All(ctx context.Context) ([]*Notification, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Notification, *NotificationQuery]()
	return withInterceptors[[]*Notification](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NotificationQuery) AllX(ctx context.Context) []*Notification {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Notification IDs.
func (nq *NotificationQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(notification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NotificationQuery) IDsX(ctx context.Context) []int32 {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NotificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NotificationQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NotificationQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NotificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NotificationQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NotificationQuery) Clone() *NotificationQuery {
	if nq == nil {
		return nil
	}
	return &NotificationQuery{
		config:              nq.config,
		ctx:                 nq.ctx.Clone(),
		order:               append([]notification.OrderOption{}, nq.order...),
		inters:              append([]Interceptor{}, nq.inters...),
		predicates:          append([]predicate.Notification{}, nq.predicates...),
		withApplications:    nq.withApplications.Clone(),
		withCenters:         nq.withCenters.Clone(),
		withNodalOfficers:   nq.withNodalOfficers.Clone(),
		withExam:            nq.withExam.Clone(),
		withVacancyYears:    nq.withVacancyYears.Clone(),
		withNotifyRef:       nq.withNotifyRef.Clone(),
		withNotificationsPs: nq.withNotificationsPs.Clone(),
		withNotificationsIP: nq.withNotificationsIP.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithApplications tells the query-builder to eager-load the nodes that are connected to
// the "applications" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithApplications(opts ...func(*ApplicationQuery)) *NotificationQuery {
	query := (&ApplicationClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withApplications = query
	return nq
}

// WithCenters tells the query-builder to eager-load the nodes that are connected to
// the "centers" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithCenters(opts ...func(*CenterQuery)) *NotificationQuery {
	query := (&CenterClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withCenters = query
	return nq
}

// WithNodalOfficers tells the query-builder to eager-load the nodes that are connected to
// the "nodal_officers" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithNodalOfficers(opts ...func(*NodalOfficerQuery)) *NotificationQuery {
	query := (&NodalOfficerClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNodalOfficers = query
	return nq
}

// WithExam tells the query-builder to eager-load the nodes that are connected to
// the "exam" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithExam(opts ...func(*ExamQuery)) *NotificationQuery {
	query := (&ExamClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withExam = query
	return nq
}

// WithVacancyYears tells the query-builder to eager-load the nodes that are connected to
// the "vacancy_years" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithVacancyYears(opts ...func(*VacancyYearQuery)) *NotificationQuery {
	query := (&VacancyYearClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withVacancyYears = query
	return nq
}

// WithNotifyRef tells the query-builder to eager-load the nodes that are connected to
// the "notify_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithNotifyRef(opts ...func(*NotificationQuery)) *NotificationQuery {
	query := (&NotificationClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNotifyRef = query
	return nq
}

// WithNotificationsPs tells the query-builder to eager-load the nodes that are connected to
// the "notifications_ps" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithNotificationsPs(opts ...func(*ExamPSQuery)) *NotificationQuery {
	query := (&ExamPSClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNotificationsPs = query
	return nq
}

// WithNotificationsIP tells the query-builder to eager-load the nodes that are connected to
// the "notifications_ip" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithNotificationsIP(opts ...func(*ExamIPQuery)) *NotificationQuery {
	query := (&ExamIPClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNotificationsIP = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExamCode int32 `json:"ExamCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Notification.Query().
//		GroupBy(notification.FieldExamCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NotificationQuery) GroupBy(field string, fields ...string) *NotificationGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotificationGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = notification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExamCode int32 `json:"ExamCode,omitempty"`
//	}
//
//	client.Notification.Query().
//		Select(notification.FieldExamCode).
//		Scan(ctx, &v)
func (nq *NotificationQuery) Select(fields ...string) *NotificationSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NotificationSelect{NotificationQuery: nq}
	sbuild.label = notification.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotificationSelect configured with the given aggregations.
func (nq *NotificationQuery) Aggregate(fns ...AggregateFunc) *NotificationSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NotificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !notification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NotificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Notification, error) {
	var (
		nodes       = []*Notification{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [8]bool{
			nq.withApplications != nil,
			nq.withCenters != nil,
			nq.withNodalOfficers != nil,
			nq.withExam != nil,
			nq.withVacancyYears != nil,
			nq.withNotifyRef != nil,
			nq.withNotificationsPs != nil,
			nq.withNotificationsIP != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, notification.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Notification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Notification{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withApplications; query != nil {
		if err := nq.loadApplications(ctx, query, nodes,
			func(n *Notification) { n.Edges.Applications = []*Application{} },
			func(n *Notification, e *Application) { n.Edges.Applications = append(n.Edges.Applications, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withCenters; query != nil {
		if err := nq.loadCenters(ctx, query, nodes,
			func(n *Notification) { n.Edges.Centers = []*Center{} },
			func(n *Notification, e *Center) { n.Edges.Centers = append(n.Edges.Centers, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNodalOfficers; query != nil {
		if err := nq.loadNodalOfficers(ctx, query, nodes,
			func(n *Notification) { n.Edges.NodalOfficers = []*NodalOfficer{} },
			func(n *Notification, e *NodalOfficer) { n.Edges.NodalOfficers = append(n.Edges.NodalOfficers, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withExam; query != nil {
		if err := nq.loadExam(ctx, query, nodes, nil,
			func(n *Notification, e *Exam) { n.Edges.Exam = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withVacancyYears; query != nil {
		if err := nq.loadVacancyYears(ctx, query, nodes,
			func(n *Notification) { n.Edges.VacancyYears = []*VacancyYear{} },
			func(n *Notification, e *VacancyYear) { n.Edges.VacancyYears = append(n.Edges.VacancyYears, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNotifyRef; query != nil {
		if err := nq.loadNotifyRef(ctx, query, nodes,
			func(n *Notification) { n.Edges.NotifyRef = []*Notification{} },
			func(n *Notification, e *Notification) { n.Edges.NotifyRef = append(n.Edges.NotifyRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNotificationsPs; query != nil {
		if err := nq.loadNotificationsPs(ctx, query, nodes,
			func(n *Notification) { n.Edges.NotificationsPs = []*Exam_PS{} },
			func(n *Notification, e *Exam_PS) { n.Edges.NotificationsPs = append(n.Edges.NotificationsPs, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNotificationsIP; query != nil {
		if err := nq.loadNotificationsIP(ctx, query, nodes,
			func(n *Notification) { n.Edges.NotificationsIP = []*Exam_IP{} },
			func(n *Notification, e *Exam_IP) { n.Edges.NotificationsIP = append(n.Edges.NotificationsIP, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NotificationQuery) loadApplications(ctx context.Context, query *ApplicationQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Application)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(application.FieldNotifyCode)
	}
	query.Where(predicate.Application(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.ApplicationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifyCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "NotifyCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NotificationQuery) loadCenters(ctx context.Context, query *CenterQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Center)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(center.FieldNotifyCode)
	}
	query.Where(predicate.Center(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.CentersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifyCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "NotifyCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NotificationQuery) loadNodalOfficers(ctx context.Context, query *NodalOfficerQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *NodalOfficer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(nodalofficer.FieldNotifyCode)
	}
	query.Where(predicate.NodalOfficer(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.NodalOfficersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifyCode
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "NotifyCode" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NotificationQuery) loadExam(ctx context.Context, query *ExamQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Exam)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*Notification)
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
func (nq *NotificationQuery) loadVacancyYears(ctx context.Context, query *VacancyYearQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *VacancyYear)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.VacancyYear(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.VacancyYearsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.notification_vacancy_years
		if fk == nil {
			return fmt.Errorf(`foreign-key "notification_vacancy_years" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notification_vacancy_years" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NotificationQuery) loadNotifyRef(ctx context.Context, query *NotificationQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Notification)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int32]*Notification)
	nids := make(map[int32]map[*Notification]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notification.NotifyRefTable)
		s.Join(joinT).On(s.C(notification.FieldID), joinT.C(notification.NotifyRefPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(notification.NotifyRefPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notification.NotifyRefPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int32(values[0].(*sql.NullInt64).Int64)
				inValue := int32(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Notification]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Notification](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "notify_ref" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NotificationQuery) loadNotificationsPs(ctx context.Context, query *ExamPSQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Exam_PS)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam_PS(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.NotificationsPsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.notification_notifications_ps
		if fk == nil {
			return fmt.Errorf(`foreign-key "notification_notifications_ps" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notification_notifications_ps" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NotificationQuery) loadNotificationsIP(ctx context.Context, query *ExamIPQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Exam_IP)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Notification)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam_IP(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notification.NotificationsIPColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.notification_notifications_ip
		if fk == nil {
			return fmt.Errorf(`foreign-key "notification_notifications_ip" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notification_notifications_ip" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NotificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NotificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for i := range fields {
			if fields[i] != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if nq.withExam != nil {
			_spec.Node.AddColumnOnce(notification.FieldExamCode)
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NotificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(notification.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = notification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotificationGroupBy is the group-by builder for Notification entities.
type NotificationGroupBy struct {
	selector
	build *NotificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NotificationGroupBy) Aggregate(fns ...AggregateFunc) *NotificationGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NotificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationQuery, *NotificationGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NotificationGroupBy) sqlScan(ctx context.Context, root *NotificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotificationSelect is the builder for selecting fields of Notification entities.
type NotificationSelect struct {
	*NotificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NotificationSelect) Aggregate(fns ...AggregateFunc) *NotificationSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NotificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationQuery, *NotificationSelect](ctx, ns.NotificationQuery, ns, ns.inters, v)
}

func (ns *NotificationSelect) sqlScan(ctx context.Context, root *NotificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
