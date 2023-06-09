// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_ip"
	"recruit/ent/examcalendar"
	"recruit/ent/exampapers"
	"recruit/ent/notification"
	"recruit/ent/predicate"
	"recruit/ent/usermaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamIPQuery is the builder for querying Exam_IP entities.
type ExamIPQuery struct {
	config
	ctx                 *QueryContext
	order               []exam_ip.OrderOption
	inters              []Interceptor
	predicates          []predicate.Exam_IP
	withExamcalIPRef    *ExamCalendarQuery
	withPapersIPRef     *ExamPapersQuery
	withUsersIPType     *UserMasterQuery
	withExamApplnIPRef  *ExamApplicationsIPQuery
	withNotificationsIP *NotificationQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamIPQuery builder.
func (eiq *ExamIPQuery) Where(ps ...predicate.Exam_IP) *ExamIPQuery {
	eiq.predicates = append(eiq.predicates, ps...)
	return eiq
}

// Limit the number of records to be returned by this query.
func (eiq *ExamIPQuery) Limit(limit int) *ExamIPQuery {
	eiq.ctx.Limit = &limit
	return eiq
}

// Offset to start from.
func (eiq *ExamIPQuery) Offset(offset int) *ExamIPQuery {
	eiq.ctx.Offset = &offset
	return eiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eiq *ExamIPQuery) Unique(unique bool) *ExamIPQuery {
	eiq.ctx.Unique = &unique
	return eiq
}

// Order specifies how the records should be ordered.
func (eiq *ExamIPQuery) Order(o ...exam_ip.OrderOption) *ExamIPQuery {
	eiq.order = append(eiq.order, o...)
	return eiq
}

// QueryExamcalIPRef chains the current query on the "examcal_ip_ref" edge.
func (eiq *ExamIPQuery) QueryExamcalIPRef() *ExamCalendarQuery {
	query := (&ExamCalendarClient{config: eiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_ip.Table, exam_ip.FieldID, selector),
			sqlgraph.To(examcalendar.Table, examcalendar.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_ip.ExamcalIPRefTable, exam_ip.ExamcalIPRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPapersIPRef chains the current query on the "papers_ip_ref" edge.
func (eiq *ExamIPQuery) QueryPapersIPRef() *ExamPapersQuery {
	query := (&ExamPapersClient{config: eiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_ip.Table, exam_ip.FieldID, selector),
			sqlgraph.To(exampapers.Table, exampapers.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_ip.PapersIPRefTable, exam_ip.PapersIPRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsersIPType chains the current query on the "users_ip_type" edge.
func (eiq *ExamIPQuery) QueryUsersIPType() *UserMasterQuery {
	query := (&UserMasterClient{config: eiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_ip.Table, exam_ip.FieldID, selector),
			sqlgraph.To(usermaster.Table, usermaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_ip.UsersIPTypeTable, exam_ip.UsersIPTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamApplnIPRef chains the current query on the "ExamAppln_IP_Ref" edge.
func (eiq *ExamIPQuery) QueryExamApplnIPRef() *ExamApplicationsIPQuery {
	query := (&ExamApplicationsIPClient{config: eiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_ip.Table, exam_ip.FieldID, selector),
			sqlgraph.To(exam_applications_ip.Table, exam_applications_ip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_ip.ExamApplnIPRefTable, exam_ip.ExamApplnIPRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotificationsIP chains the current query on the "notifications_ip" edge.
func (eiq *ExamIPQuery) QueryNotificationsIP() *NotificationQuery {
	query := (&NotificationClient{config: eiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_ip.Table, exam_ip.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_ip.NotificationsIPTable, exam_ip.NotificationsIPColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Exam_IP entity from the query.
// Returns a *NotFoundError when no Exam_IP was found.
func (eiq *ExamIPQuery) First(ctx context.Context) (*Exam_IP, error) {
	nodes, err := eiq.Limit(1).All(setContextOp(ctx, eiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exam_ip.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eiq *ExamIPQuery) FirstX(ctx context.Context) *Exam_IP {
	node, err := eiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Exam_IP ID from the query.
// Returns a *NotFoundError when no Exam_IP ID was found.
func (eiq *ExamIPQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eiq.Limit(1).IDs(setContextOp(ctx, eiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exam_ip.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eiq *ExamIPQuery) FirstIDX(ctx context.Context) int32 {
	id, err := eiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Exam_IP entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Exam_IP entity is found.
// Returns a *NotFoundError when no Exam_IP entities are found.
func (eiq *ExamIPQuery) Only(ctx context.Context) (*Exam_IP, error) {
	nodes, err := eiq.Limit(2).All(setContextOp(ctx, eiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exam_ip.Label}
	default:
		return nil, &NotSingularError{exam_ip.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eiq *ExamIPQuery) OnlyX(ctx context.Context) *Exam_IP {
	node, err := eiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Exam_IP ID in the query.
// Returns a *NotSingularError when more than one Exam_IP ID is found.
// Returns a *NotFoundError when no entities are found.
func (eiq *ExamIPQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eiq.Limit(2).IDs(setContextOp(ctx, eiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exam_ip.Label}
	default:
		err = &NotSingularError{exam_ip.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eiq *ExamIPQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := eiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Exam_IPs.
func (eiq *ExamIPQuery) All(ctx context.Context) ([]*Exam_IP, error) {
	ctx = setContextOp(ctx, eiq.ctx, "All")
	if err := eiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Exam_IP, *ExamIPQuery]()
	return withInterceptors[[]*Exam_IP](ctx, eiq, qr, eiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eiq *ExamIPQuery) AllX(ctx context.Context) []*Exam_IP {
	nodes, err := eiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Exam_IP IDs.
func (eiq *ExamIPQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if eiq.ctx.Unique == nil && eiq.path != nil {
		eiq.Unique(true)
	}
	ctx = setContextOp(ctx, eiq.ctx, "IDs")
	if err = eiq.Select(exam_ip.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eiq *ExamIPQuery) IDsX(ctx context.Context) []int32 {
	ids, err := eiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eiq *ExamIPQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eiq.ctx, "Count")
	if err := eiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eiq, querierCount[*ExamIPQuery](), eiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eiq *ExamIPQuery) CountX(ctx context.Context) int {
	count, err := eiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eiq *ExamIPQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eiq.ctx, "Exist")
	switch _, err := eiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eiq *ExamIPQuery) ExistX(ctx context.Context) bool {
	exist, err := eiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamIPQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eiq *ExamIPQuery) Clone() *ExamIPQuery {
	if eiq == nil {
		return nil
	}
	return &ExamIPQuery{
		config:              eiq.config,
		ctx:                 eiq.ctx.Clone(),
		order:               append([]exam_ip.OrderOption{}, eiq.order...),
		inters:              append([]Interceptor{}, eiq.inters...),
		predicates:          append([]predicate.Exam_IP{}, eiq.predicates...),
		withExamcalIPRef:    eiq.withExamcalIPRef.Clone(),
		withPapersIPRef:     eiq.withPapersIPRef.Clone(),
		withUsersIPType:     eiq.withUsersIPType.Clone(),
		withExamApplnIPRef:  eiq.withExamApplnIPRef.Clone(),
		withNotificationsIP: eiq.withNotificationsIP.Clone(),
		// clone intermediate query.
		sql:  eiq.sql.Clone(),
		path: eiq.path,
	}
}

// WithExamcalIPRef tells the query-builder to eager-load the nodes that are connected to
// the "examcal_ip_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *ExamIPQuery) WithExamcalIPRef(opts ...func(*ExamCalendarQuery)) *ExamIPQuery {
	query := (&ExamCalendarClient{config: eiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eiq.withExamcalIPRef = query
	return eiq
}

// WithPapersIPRef tells the query-builder to eager-load the nodes that are connected to
// the "papers_ip_ref" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *ExamIPQuery) WithPapersIPRef(opts ...func(*ExamPapersQuery)) *ExamIPQuery {
	query := (&ExamPapersClient{config: eiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eiq.withPapersIPRef = query
	return eiq
}

// WithUsersIPType tells the query-builder to eager-load the nodes that are connected to
// the "users_ip_type" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *ExamIPQuery) WithUsersIPType(opts ...func(*UserMasterQuery)) *ExamIPQuery {
	query := (&UserMasterClient{config: eiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eiq.withUsersIPType = query
	return eiq
}

// WithExamApplnIPRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamAppln_IP_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *ExamIPQuery) WithExamApplnIPRef(opts ...func(*ExamApplicationsIPQuery)) *ExamIPQuery {
	query := (&ExamApplicationsIPClient{config: eiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eiq.withExamApplnIPRef = query
	return eiq
}

// WithNotificationsIP tells the query-builder to eager-load the nodes that are connected to
// the "notifications_ip" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *ExamIPQuery) WithNotificationsIP(opts ...func(*NotificationQuery)) *ExamIPQuery {
	query := (&NotificationClient{config: eiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eiq.withNotificationsIP = query
	return eiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ExamNameCode string `json:"ExamNameCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamIP.Query().
//		GroupBy(exam_ip.FieldExamNameCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eiq *ExamIPQuery) GroupBy(field string, fields ...string) *ExamIPGroupBy {
	eiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamIPGroupBy{build: eiq}
	grbuild.flds = &eiq.ctx.Fields
	grbuild.label = exam_ip.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ExamNameCode string `json:"ExamNameCode,omitempty"`
//	}
//
//	client.ExamIP.Query().
//		Select(exam_ip.FieldExamNameCode).
//		Scan(ctx, &v)
func (eiq *ExamIPQuery) Select(fields ...string) *ExamIPSelect {
	eiq.ctx.Fields = append(eiq.ctx.Fields, fields...)
	sbuild := &ExamIPSelect{ExamIPQuery: eiq}
	sbuild.label = exam_ip.Label
	sbuild.flds, sbuild.scan = &eiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamIPSelect configured with the given aggregations.
func (eiq *ExamIPQuery) Aggregate(fns ...AggregateFunc) *ExamIPSelect {
	return eiq.Select().Aggregate(fns...)
}

func (eiq *ExamIPQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eiq); err != nil {
				return err
			}
		}
	}
	for _, f := range eiq.ctx.Fields {
		if !exam_ip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eiq.path != nil {
		prev, err := eiq.path(ctx)
		if err != nil {
			return err
		}
		eiq.sql = prev
	}
	return nil
}

func (eiq *ExamIPQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Exam_IP, error) {
	var (
		nodes       = []*Exam_IP{}
		withFKs     = eiq.withFKs
		_spec       = eiq.querySpec()
		loadedTypes = [5]bool{
			eiq.withExamcalIPRef != nil,
			eiq.withPapersIPRef != nil,
			eiq.withUsersIPType != nil,
			eiq.withExamApplnIPRef != nil,
			eiq.withNotificationsIP != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, exam_ip.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Exam_IP).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Exam_IP{config: eiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eiq.withExamcalIPRef; query != nil {
		if err := eiq.loadExamcalIPRef(ctx, query, nodes,
			func(n *Exam_IP) { n.Edges.ExamcalIPRef = []*ExamCalendar{} },
			func(n *Exam_IP, e *ExamCalendar) { n.Edges.ExamcalIPRef = append(n.Edges.ExamcalIPRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eiq.withPapersIPRef; query != nil {
		if err := eiq.loadPapersIPRef(ctx, query, nodes,
			func(n *Exam_IP) { n.Edges.PapersIPRef = []*ExamPapers{} },
			func(n *Exam_IP, e *ExamPapers) { n.Edges.PapersIPRef = append(n.Edges.PapersIPRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eiq.withUsersIPType; query != nil {
		if err := eiq.loadUsersIPType(ctx, query, nodes,
			func(n *Exam_IP) { n.Edges.UsersIPType = []*UserMaster{} },
			func(n *Exam_IP, e *UserMaster) { n.Edges.UsersIPType = append(n.Edges.UsersIPType, e) }); err != nil {
			return nil, err
		}
	}
	if query := eiq.withExamApplnIPRef; query != nil {
		if err := eiq.loadExamApplnIPRef(ctx, query, nodes,
			func(n *Exam_IP) { n.Edges.ExamApplnIPRef = []*Exam_Applications_IP{} },
			func(n *Exam_IP, e *Exam_Applications_IP) { n.Edges.ExamApplnIPRef = append(n.Edges.ExamApplnIPRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eiq.withNotificationsIP; query != nil {
		if err := eiq.loadNotificationsIP(ctx, query, nodes,
			func(n *Exam_IP) { n.Edges.NotificationsIP = []*Notification{} },
			func(n *Exam_IP, e *Notification) { n.Edges.NotificationsIP = append(n.Edges.NotificationsIP, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eiq *ExamIPQuery) loadExamcalIPRef(ctx context.Context, query *ExamCalendarQuery, nodes []*Exam_IP, init func(*Exam_IP), assign func(*Exam_IP, *ExamCalendar)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Exam_IP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ExamCalendar(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_ip.ExamcalIPRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_ip_examcal_ip_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_ip_examcal_ip_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_ip_examcal_ip_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eiq *ExamIPQuery) loadPapersIPRef(ctx context.Context, query *ExamPapersQuery, nodes []*Exam_IP, init func(*Exam_IP), assign func(*Exam_IP, *ExamPapers)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Exam_IP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ExamPapers(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_ip.PapersIPRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_ip_papers_ip_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_ip_papers_ip_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_ip_papers_ip_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eiq *ExamIPQuery) loadUsersIPType(ctx context.Context, query *UserMasterQuery, nodes []*Exam_IP, init func(*Exam_IP), assign func(*Exam_IP, *UserMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Exam_IP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_ip.UsersIPTypeColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_ip_users_ip_type
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_ip_users_ip_type" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_ip_users_ip_type" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eiq *ExamIPQuery) loadExamApplnIPRef(ctx context.Context, query *ExamApplicationsIPQuery, nodes []*Exam_IP, init func(*Exam_IP), assign func(*Exam_IP, *Exam_Applications_IP)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Exam_IP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam_Applications_IP(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_ip.ExamApplnIPRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_ip_exam_appln_ip_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_ip_exam_appln_ip_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_ip_exam_appln_ip_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eiq *ExamIPQuery) loadNotificationsIP(ctx context.Context, query *NotificationQuery, nodes []*Exam_IP, init func(*Exam_IP), assign func(*Exam_IP, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Exam_IP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_ip.NotificationsIPColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_ip_notifications_ip
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_ip_notifications_ip" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_ip_notifications_ip" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eiq *ExamIPQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eiq.querySpec()
	_spec.Node.Columns = eiq.ctx.Fields
	if len(eiq.ctx.Fields) > 0 {
		_spec.Unique = eiq.ctx.Unique != nil && *eiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eiq.driver, _spec)
}

func (eiq *ExamIPQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(exam_ip.Table, exam_ip.Columns, sqlgraph.NewFieldSpec(exam_ip.FieldID, field.TypeInt32))
	_spec.From = eiq.sql
	if unique := eiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eiq.path != nil {
		_spec.Unique = true
	}
	if fields := eiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exam_ip.FieldID)
		for i := range fields {
			if fields[i] != exam_ip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eiq *ExamIPQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eiq.driver.Dialect())
	t1 := builder.Table(exam_ip.Table)
	columns := eiq.ctx.Fields
	if len(columns) == 0 {
		columns = exam_ip.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eiq.sql != nil {
		selector = eiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eiq.ctx.Unique != nil && *eiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eiq.predicates {
		p(selector)
	}
	for _, p := range eiq.order {
		p(selector)
	}
	if offset := eiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamIPGroupBy is the group-by builder for Exam_IP entities.
type ExamIPGroupBy struct {
	selector
	build *ExamIPQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eigb *ExamIPGroupBy) Aggregate(fns ...AggregateFunc) *ExamIPGroupBy {
	eigb.fns = append(eigb.fns, fns...)
	return eigb
}

// Scan applies the selector query and scans the result into the given value.
func (eigb *ExamIPGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eigb.build.ctx, "GroupBy")
	if err := eigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamIPQuery, *ExamIPGroupBy](ctx, eigb.build, eigb, eigb.build.inters, v)
}

func (eigb *ExamIPGroupBy) sqlScan(ctx context.Context, root *ExamIPQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eigb.fns))
	for _, fn := range eigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eigb.flds)+len(eigb.fns))
		for _, f := range *eigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamIPSelect is the builder for selecting fields of ExamIP entities.
type ExamIPSelect struct {
	*ExamIPQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eis *ExamIPSelect) Aggregate(fns ...AggregateFunc) *ExamIPSelect {
	eis.fns = append(eis.fns, fns...)
	return eis
}

// Scan applies the selector query and scans the result into the given value.
func (eis *ExamIPSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eis.ctx, "Select")
	if err := eis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamIPQuery, *ExamIPSelect](ctx, eis.ExamIPQuery, eis, eis.inters, v)
}

func (eis *ExamIPSelect) sqlScan(ctx context.Context, root *ExamIPQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eis.fns))
	for _, fn := range eis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
