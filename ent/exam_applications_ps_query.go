// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/exam_ps"
	"recruit/ent/facility"
	"recruit/ent/predicate"
	"recruit/ent/rolemaster"
	"recruit/ent/usermaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamApplicationsPSQuery is the builder for querying Exam_Applications_PS entities.
type ExamApplicationsPSQuery struct {
	config
	ctx                *QueryContext
	order              []exam_applications_ps.OrderOption
	inters             []Interceptor
	predicates         []predicate.Exam_Applications_PS
	withUsersPSRef     *UserMasterQuery
	withExamApplnPSRef *ExamPSQuery
	withOfficePSRef    *FacilityQuery
	withRoleusers      *RoleMasterQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamApplicationsPSQuery builder.
func (eapq *ExamApplicationsPSQuery) Where(ps ...predicate.Exam_Applications_PS) *ExamApplicationsPSQuery {
	eapq.predicates = append(eapq.predicates, ps...)
	return eapq
}

// Limit the number of records to be returned by this query.
func (eapq *ExamApplicationsPSQuery) Limit(limit int) *ExamApplicationsPSQuery {
	eapq.ctx.Limit = &limit
	return eapq
}

// Offset to start from.
func (eapq *ExamApplicationsPSQuery) Offset(offset int) *ExamApplicationsPSQuery {
	eapq.ctx.Offset = &offset
	return eapq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eapq *ExamApplicationsPSQuery) Unique(unique bool) *ExamApplicationsPSQuery {
	eapq.ctx.Unique = &unique
	return eapq
}

// Order specifies how the records should be ordered.
func (eapq *ExamApplicationsPSQuery) Order(o ...exam_applications_ps.OrderOption) *ExamApplicationsPSQuery {
	eapq.order = append(eapq.order, o...)
	return eapq
}

// QueryUsersPSRef chains the current query on the "UsersPSRef" edge.
func (eapq *ExamApplicationsPSQuery) QueryUsersPSRef() *UserMasterQuery {
	query := (&UserMasterClient{config: eapq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_applications_ps.Table, exam_applications_ps.FieldID, selector),
			sqlgraph.To(usermaster.Table, usermaster.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_applications_ps.UsersPSRefTable, exam_applications_ps.UsersPSRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExamApplnPSRef chains the current query on the "ExamAppln_PS_Ref" edge.
func (eapq *ExamApplicationsPSQuery) QueryExamApplnPSRef() *ExamPSQuery {
	query := (&ExamPSClient{config: eapq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_applications_ps.Table, exam_applications_ps.FieldID, selector),
			sqlgraph.To(exam_ps.Table, exam_ps.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_applications_ps.ExamApplnPSRefTable, exam_applications_ps.ExamApplnPSRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOfficePSRef chains the current query on the "Office_PS_Ref" edge.
func (eapq *ExamApplicationsPSQuery) QueryOfficePSRef() *FacilityQuery {
	query := (&FacilityClient{config: eapq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_applications_ps.Table, exam_applications_ps.FieldID, selector),
			sqlgraph.To(facility.Table, facility.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, exam_applications_ps.OfficePSRefTable, exam_applications_ps.OfficePSRefColumn),
		)
		fromU = sqlgraph.SetNeighbors(eapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoleusers chains the current query on the "roleusers" edge.
func (eapq *ExamApplicationsPSQuery) QueryRoleusers() *RoleMasterQuery {
	query := (&RoleMasterClient{config: eapq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eapq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eapq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(exam_applications_ps.Table, exam_applications_ps.FieldID, selector),
			sqlgraph.To(rolemaster.Table, rolemaster.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, exam_applications_ps.RoleusersTable, exam_applications_ps.RoleusersColumn),
		)
		fromU = sqlgraph.SetNeighbors(eapq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Exam_Applications_PS entity from the query.
// Returns a *NotFoundError when no Exam_Applications_PS was found.
func (eapq *ExamApplicationsPSQuery) First(ctx context.Context) (*Exam_Applications_PS, error) {
	nodes, err := eapq.Limit(1).All(setContextOp(ctx, eapq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exam_applications_ps.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) FirstX(ctx context.Context) *Exam_Applications_PS {
	node, err := eapq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Exam_Applications_PS ID from the query.
// Returns a *NotFoundError when no Exam_Applications_PS ID was found.
func (eapq *ExamApplicationsPSQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = eapq.Limit(1).IDs(setContextOp(ctx, eapq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exam_applications_ps.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) FirstIDX(ctx context.Context) int64 {
	id, err := eapq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Exam_Applications_PS entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Exam_Applications_PS entity is found.
// Returns a *NotFoundError when no Exam_Applications_PS entities are found.
func (eapq *ExamApplicationsPSQuery) Only(ctx context.Context) (*Exam_Applications_PS, error) {
	nodes, err := eapq.Limit(2).All(setContextOp(ctx, eapq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exam_applications_ps.Label}
	default:
		return nil, &NotSingularError{exam_applications_ps.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) OnlyX(ctx context.Context) *Exam_Applications_PS {
	node, err := eapq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Exam_Applications_PS ID in the query.
// Returns a *NotSingularError when more than one Exam_Applications_PS ID is found.
// Returns a *NotFoundError when no entities are found.
func (eapq *ExamApplicationsPSQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = eapq.Limit(2).IDs(setContextOp(ctx, eapq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exam_applications_ps.Label}
	default:
		err = &NotSingularError{exam_applications_ps.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := eapq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Exam_Applications_PSs.
func (eapq *ExamApplicationsPSQuery) All(ctx context.Context) ([]*Exam_Applications_PS, error) {
	ctx = setContextOp(ctx, eapq.ctx, "All")
	if err := eapq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Exam_Applications_PS, *ExamApplicationsPSQuery]()
	return withInterceptors[[]*Exam_Applications_PS](ctx, eapq, qr, eapq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) AllX(ctx context.Context) []*Exam_Applications_PS {
	nodes, err := eapq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Exam_Applications_PS IDs.
func (eapq *ExamApplicationsPSQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if eapq.ctx.Unique == nil && eapq.path != nil {
		eapq.Unique(true)
	}
	ctx = setContextOp(ctx, eapq.ctx, "IDs")
	if err = eapq.Select(exam_applications_ps.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) IDsX(ctx context.Context) []int64 {
	ids, err := eapq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eapq *ExamApplicationsPSQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eapq.ctx, "Count")
	if err := eapq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eapq, querierCount[*ExamApplicationsPSQuery](), eapq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) CountX(ctx context.Context) int {
	count, err := eapq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eapq *ExamApplicationsPSQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eapq.ctx, "Exist")
	switch _, err := eapq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eapq *ExamApplicationsPSQuery) ExistX(ctx context.Context) bool {
	exist, err := eapq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamApplicationsPSQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eapq *ExamApplicationsPSQuery) Clone() *ExamApplicationsPSQuery {
	if eapq == nil {
		return nil
	}
	return &ExamApplicationsPSQuery{
		config:             eapq.config,
		ctx:                eapq.ctx.Clone(),
		order:              append([]exam_applications_ps.OrderOption{}, eapq.order...),
		inters:             append([]Interceptor{}, eapq.inters...),
		predicates:         append([]predicate.Exam_Applications_PS{}, eapq.predicates...),
		withUsersPSRef:     eapq.withUsersPSRef.Clone(),
		withExamApplnPSRef: eapq.withExamApplnPSRef.Clone(),
		withOfficePSRef:    eapq.withOfficePSRef.Clone(),
		withRoleusers:      eapq.withRoleusers.Clone(),
		// clone intermediate query.
		sql:  eapq.sql.Clone(),
		path: eapq.path,
	}
}

// WithUsersPSRef tells the query-builder to eager-load the nodes that are connected to
// the "UsersPSRef" edge. The optional arguments are used to configure the query builder of the edge.
func (eapq *ExamApplicationsPSQuery) WithUsersPSRef(opts ...func(*UserMasterQuery)) *ExamApplicationsPSQuery {
	query := (&UserMasterClient{config: eapq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eapq.withUsersPSRef = query
	return eapq
}

// WithExamApplnPSRef tells the query-builder to eager-load the nodes that are connected to
// the "ExamAppln_PS_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (eapq *ExamApplicationsPSQuery) WithExamApplnPSRef(opts ...func(*ExamPSQuery)) *ExamApplicationsPSQuery {
	query := (&ExamPSClient{config: eapq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eapq.withExamApplnPSRef = query
	return eapq
}

// WithOfficePSRef tells the query-builder to eager-load the nodes that are connected to
// the "Office_PS_Ref" edge. The optional arguments are used to configure the query builder of the edge.
func (eapq *ExamApplicationsPSQuery) WithOfficePSRef(opts ...func(*FacilityQuery)) *ExamApplicationsPSQuery {
	query := (&FacilityClient{config: eapq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eapq.withOfficePSRef = query
	return eapq
}

// WithRoleusers tells the query-builder to eager-load the nodes that are connected to
// the "roleusers" edge. The optional arguments are used to configure the query builder of the edge.
func (eapq *ExamApplicationsPSQuery) WithRoleusers(opts ...func(*RoleMasterQuery)) *ExamApplicationsPSQuery {
	query := (&RoleMasterClient{config: eapq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eapq.withRoleusers = query
	return eapq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplicationNumber string `json:"ApplicationNumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamApplicationsPS.Query().
//		GroupBy(exam_applications_ps.FieldApplicationNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eapq *ExamApplicationsPSQuery) GroupBy(field string, fields ...string) *ExamApplicationsPSGroupBy {
	eapq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamApplicationsPSGroupBy{build: eapq}
	grbuild.flds = &eapq.ctx.Fields
	grbuild.label = exam_applications_ps.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ApplicationNumber string `json:"ApplicationNumber,omitempty"`
//	}
//
//	client.ExamApplicationsPS.Query().
//		Select(exam_applications_ps.FieldApplicationNumber).
//		Scan(ctx, &v)
func (eapq *ExamApplicationsPSQuery) Select(fields ...string) *ExamApplicationsPSSelect {
	eapq.ctx.Fields = append(eapq.ctx.Fields, fields...)
	sbuild := &ExamApplicationsPSSelect{ExamApplicationsPSQuery: eapq}
	sbuild.label = exam_applications_ps.Label
	sbuild.flds, sbuild.scan = &eapq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamApplicationsPSSelect configured with the given aggregations.
func (eapq *ExamApplicationsPSQuery) Aggregate(fns ...AggregateFunc) *ExamApplicationsPSSelect {
	return eapq.Select().Aggregate(fns...)
}

func (eapq *ExamApplicationsPSQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eapq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eapq); err != nil {
				return err
			}
		}
	}
	for _, f := range eapq.ctx.Fields {
		if !exam_applications_ps.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eapq.path != nil {
		prev, err := eapq.path(ctx)
		if err != nil {
			return err
		}
		eapq.sql = prev
	}
	return nil
}

func (eapq *ExamApplicationsPSQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Exam_Applications_PS, error) {
	var (
		nodes       = []*Exam_Applications_PS{}
		withFKs     = eapq.withFKs
		_spec       = eapq.querySpec()
		loadedTypes = [4]bool{
			eapq.withUsersPSRef != nil,
			eapq.withExamApplnPSRef != nil,
			eapq.withOfficePSRef != nil,
			eapq.withRoleusers != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, exam_applications_ps.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Exam_Applications_PS).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Exam_Applications_PS{config: eapq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eapq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eapq.withUsersPSRef; query != nil {
		if err := eapq.loadUsersPSRef(ctx, query, nodes,
			func(n *Exam_Applications_PS) { n.Edges.UsersPSRef = []*UserMaster{} },
			func(n *Exam_Applications_PS, e *UserMaster) { n.Edges.UsersPSRef = append(n.Edges.UsersPSRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eapq.withExamApplnPSRef; query != nil {
		if err := eapq.loadExamApplnPSRef(ctx, query, nodes,
			func(n *Exam_Applications_PS) { n.Edges.ExamApplnPSRef = []*Exam_PS{} },
			func(n *Exam_Applications_PS, e *Exam_PS) { n.Edges.ExamApplnPSRef = append(n.Edges.ExamApplnPSRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eapq.withOfficePSRef; query != nil {
		if err := eapq.loadOfficePSRef(ctx, query, nodes,
			func(n *Exam_Applications_PS) { n.Edges.OfficePSRef = []*Facility{} },
			func(n *Exam_Applications_PS, e *Facility) { n.Edges.OfficePSRef = append(n.Edges.OfficePSRef, e) }); err != nil {
			return nil, err
		}
	}
	if query := eapq.withRoleusers; query != nil {
		if err := eapq.loadRoleusers(ctx, query, nodes, nil,
			func(n *Exam_Applications_PS, e *RoleMaster) { n.Edges.Roleusers = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eapq *ExamApplicationsPSQuery) loadUsersPSRef(ctx context.Context, query *UserMasterQuery, nodes []*Exam_Applications_PS, init func(*Exam_Applications_PS), assign func(*Exam_Applications_PS, *UserMaster)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Exam_Applications_PS)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserMaster(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_applications_ps.UsersPSRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_applications_ps_users_ps_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_applications_ps_users_ps_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_applications_ps_users_ps_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eapq *ExamApplicationsPSQuery) loadExamApplnPSRef(ctx context.Context, query *ExamPSQuery, nodes []*Exam_Applications_PS, init func(*Exam_Applications_PS), assign func(*Exam_Applications_PS, *Exam_PS)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Exam_Applications_PS)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam_PS(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_applications_ps.ExamApplnPSRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_applications_ps_exam_appln_ps_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_applications_ps_exam_appln_ps_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_applications_ps_exam_appln_ps_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eapq *ExamApplicationsPSQuery) loadOfficePSRef(ctx context.Context, query *FacilityQuery, nodes []*Exam_Applications_PS, init func(*Exam_Applications_PS), assign func(*Exam_Applications_PS, *Facility)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Exam_Applications_PS)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Facility(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(exam_applications_ps.OfficePSRefColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_applications_ps_office_ps_ref
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_applications_ps_office_ps_ref" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_applications_ps_office_ps_ref" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eapq *ExamApplicationsPSQuery) loadRoleusers(ctx context.Context, query *RoleMasterQuery, nodes []*Exam_Applications_PS, init func(*Exam_Applications_PS), assign func(*Exam_Applications_PS, *RoleMaster)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*Exam_Applications_PS)
	for i := range nodes {
		fk := nodes[i].RoleUserCode
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(rolemaster.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "RoleUserCode" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eapq *ExamApplicationsPSQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eapq.querySpec()
	_spec.Node.Columns = eapq.ctx.Fields
	if len(eapq.ctx.Fields) > 0 {
		_spec.Unique = eapq.ctx.Unique != nil && *eapq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eapq.driver, _spec)
}

func (eapq *ExamApplicationsPSQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(exam_applications_ps.Table, exam_applications_ps.Columns, sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64))
	_spec.From = eapq.sql
	if unique := eapq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eapq.path != nil {
		_spec.Unique = true
	}
	if fields := eapq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exam_applications_ps.FieldID)
		for i := range fields {
			if fields[i] != exam_applications_ps.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eapq.withRoleusers != nil {
			_spec.Node.AddColumnOnce(exam_applications_ps.FieldRoleUserCode)
		}
	}
	if ps := eapq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eapq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eapq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eapq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eapq *ExamApplicationsPSQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eapq.driver.Dialect())
	t1 := builder.Table(exam_applications_ps.Table)
	columns := eapq.ctx.Fields
	if len(columns) == 0 {
		columns = exam_applications_ps.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eapq.sql != nil {
		selector = eapq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eapq.ctx.Unique != nil && *eapq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eapq.predicates {
		p(selector)
	}
	for _, p := range eapq.order {
		p(selector)
	}
	if offset := eapq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eapq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamApplicationsPSGroupBy is the group-by builder for Exam_Applications_PS entities.
type ExamApplicationsPSGroupBy struct {
	selector
	build *ExamApplicationsPSQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eapgb *ExamApplicationsPSGroupBy) Aggregate(fns ...AggregateFunc) *ExamApplicationsPSGroupBy {
	eapgb.fns = append(eapgb.fns, fns...)
	return eapgb
}

// Scan applies the selector query and scans the result into the given value.
func (eapgb *ExamApplicationsPSGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eapgb.build.ctx, "GroupBy")
	if err := eapgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamApplicationsPSQuery, *ExamApplicationsPSGroupBy](ctx, eapgb.build, eapgb, eapgb.build.inters, v)
}

func (eapgb *ExamApplicationsPSGroupBy) sqlScan(ctx context.Context, root *ExamApplicationsPSQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eapgb.fns))
	for _, fn := range eapgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eapgb.flds)+len(eapgb.fns))
		for _, f := range *eapgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eapgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eapgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamApplicationsPSSelect is the builder for selecting fields of ExamApplicationsPS entities.
type ExamApplicationsPSSelect struct {
	*ExamApplicationsPSQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eaps *ExamApplicationsPSSelect) Aggregate(fns ...AggregateFunc) *ExamApplicationsPSSelect {
	eaps.fns = append(eaps.fns, fns...)
	return eaps
}

// Scan applies the selector query and scans the result into the given value.
func (eaps *ExamApplicationsPSSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eaps.ctx, "Select")
	if err := eaps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamApplicationsPSQuery, *ExamApplicationsPSSelect](ctx, eaps.ExamApplicationsPSQuery, eaps, eaps.inters, v)
}

func (eaps *ExamApplicationsPSSelect) sqlScan(ctx context.Context, root *ExamApplicationsPSQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eaps.fns))
	for _, fn := range eaps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eaps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eaps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
