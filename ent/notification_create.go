// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/application"
	"recruit/ent/center"
	"recruit/ent/exam"
	"recruit/ent/nodalofficer"
	"recruit/ent/notification"
	"recruit/ent/vacancyyear"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
}

// SetExamCode sets the "ExamCode" field.
func (nc *NotificationCreate) SetExamCode(i int32) *NotificationCreate {
	nc.mutation.SetExamCode(i)
	return nc
}

// SetNillableExamCode sets the "ExamCode" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableExamCode(i *int32) *NotificationCreate {
	if i != nil {
		nc.SetExamCode(*i)
	}
	return nc
}

// SetExamYear sets the "ExamYear" field.
func (nc *NotificationCreate) SetExamYear(i int32) *NotificationCreate {
	nc.mutation.SetExamYear(i)
	return nc
}

// SetApplicationStartDate sets the "ApplicationStartDate" field.
func (nc *NotificationCreate) SetApplicationStartDate(t time.Time) *NotificationCreate {
	nc.mutation.SetApplicationStartDate(t)
	return nc
}

// SetApplicationEndDate sets the "ApplicationEndDate" field.
func (nc *NotificationCreate) SetApplicationEndDate(t time.Time) *NotificationCreate {
	nc.mutation.SetApplicationEndDate(t)
	return nc
}

// SetVerificationDateByController sets the "VerificationDateByController" field.
func (nc *NotificationCreate) SetVerificationDateByController(t time.Time) *NotificationCreate {
	nc.mutation.SetVerificationDateByController(t)
	return nc
}

// SetCorrectionDateByCandidate sets the "CorrectionDateByCandidate" field.
func (nc *NotificationCreate) SetCorrectionDateByCandidate(t time.Time) *NotificationCreate {
	nc.mutation.SetCorrectionDateByCandidate(t)
	return nc
}

// SetCorrectionVeriyDateByController sets the "CorrectionVeriyDateByController" field.
func (nc *NotificationCreate) SetCorrectionVeriyDateByController(t time.Time) *NotificationCreate {
	nc.mutation.SetCorrectionVeriyDateByController(t)
	return nc
}

// SetHallTicketAllotmentDateByNodalOfficer sets the "HallTicketAllotmentDateByNodalOfficer" field.
func (nc *NotificationCreate) SetHallTicketAllotmentDateByNodalOfficer(t time.Time) *NotificationCreate {
	nc.mutation.SetHallTicketAllotmentDateByNodalOfficer(t)
	return nc
}

// SetHallTicketDownloadDate sets the "HallTicketDownloadDate" field.
func (nc *NotificationCreate) SetHallTicketDownloadDate(t time.Time) *NotificationCreate {
	nc.mutation.SetHallTicketDownloadDate(t)
	return nc
}

// SetNillableHallTicketDownloadDate sets the "HallTicketDownloadDate" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableHallTicketDownloadDate(t *time.Time) *NotificationCreate {
	if t != nil {
		nc.SetHallTicketDownloadDate(*t)
	}
	return nc
}

// SetNotifyFile sets the "NotifyFile" field.
func (nc *NotificationCreate) SetNotifyFile(s string) *NotificationCreate {
	nc.mutation.SetNotifyFile(s)
	return nc
}

// SetNillableNotifyFile sets the "NotifyFile" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableNotifyFile(s *string) *NotificationCreate {
	if s != nil {
		nc.SetNotifyFile(*s)
	}
	return nc
}

// SetSyllabusFile sets the "SyllabusFile" field.
func (nc *NotificationCreate) SetSyllabusFile(s string) *NotificationCreate {
	nc.mutation.SetSyllabusFile(s)
	return nc
}

// SetNillableSyllabusFile sets the "SyllabusFile" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableSyllabusFile(s *string) *NotificationCreate {
	if s != nil {
		nc.SetSyllabusFile(*s)
	}
	return nc
}

// SetVacanciesFile sets the "VacanciesFile" field.
func (nc *NotificationCreate) SetVacanciesFile(s string) *NotificationCreate {
	nc.mutation.SetVacanciesFile(s)
	return nc
}

// SetNillableVacanciesFile sets the "VacanciesFile" field if the given value is not nil.
func (nc *NotificationCreate) SetNillableVacanciesFile(s *string) *NotificationCreate {
	if s != nil {
		nc.SetVacanciesFile(*s)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotificationCreate) SetID(i int32) *NotificationCreate {
	nc.mutation.SetID(i)
	return nc
}

// AddApplicationIDs adds the "applications" edge to the Application entity by IDs.
func (nc *NotificationCreate) AddApplicationIDs(ids ...int32) *NotificationCreate {
	nc.mutation.AddApplicationIDs(ids...)
	return nc
}

// AddApplications adds the "applications" edges to the Application entity.
func (nc *NotificationCreate) AddApplications(a ...*Application) *NotificationCreate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nc.AddApplicationIDs(ids...)
}

// AddCenterIDs adds the "centers" edge to the Center entity by IDs.
func (nc *NotificationCreate) AddCenterIDs(ids ...int32) *NotificationCreate {
	nc.mutation.AddCenterIDs(ids...)
	return nc
}

// AddCenters adds the "centers" edges to the Center entity.
func (nc *NotificationCreate) AddCenters(c ...*Center) *NotificationCreate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nc.AddCenterIDs(ids...)
}

// AddNodalOfficerIDs adds the "nodal_officers" edge to the NodalOfficer entity by IDs.
func (nc *NotificationCreate) AddNodalOfficerIDs(ids ...int32) *NotificationCreate {
	nc.mutation.AddNodalOfficerIDs(ids...)
	return nc
}

// AddNodalOfficers adds the "nodal_officers" edges to the NodalOfficer entity.
func (nc *NotificationCreate) AddNodalOfficers(n ...*NodalOfficer) *NotificationCreate {
	ids := make([]int32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddNodalOfficerIDs(ids...)
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (nc *NotificationCreate) SetExamID(id int32) *NotificationCreate {
	nc.mutation.SetExamID(id)
	return nc
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (nc *NotificationCreate) SetNillableExamID(id *int32) *NotificationCreate {
	if id != nil {
		nc = nc.SetExamID(*id)
	}
	return nc
}

// SetExam sets the "exam" edge to the Exam entity.
func (nc *NotificationCreate) SetExam(e *Exam) *NotificationCreate {
	return nc.SetExamID(e.ID)
}

// AddVacancyYearIDs adds the "vacancy_years" edge to the VacancyYear entity by IDs.
func (nc *NotificationCreate) AddVacancyYearIDs(ids ...int32) *NotificationCreate {
	nc.mutation.AddVacancyYearIDs(ids...)
	return nc
}

// AddVacancyYears adds the "vacancy_years" edges to the VacancyYear entity.
func (nc *NotificationCreate) AddVacancyYears(v ...*VacancyYear) *NotificationCreate {
	ids := make([]int32, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return nc.AddVacancyYearIDs(ids...)
}

// AddNotifyRefIDs adds the "notify_ref" edge to the Notification entity by IDs.
func (nc *NotificationCreate) AddNotifyRefIDs(ids ...int32) *NotificationCreate {
	nc.mutation.AddNotifyRefIDs(ids...)
	return nc
}

// AddNotifyRef adds the "notify_ref" edges to the Notification entity.
func (nc *NotificationCreate) AddNotifyRef(n ...*Notification) *NotificationCreate {
	ids := make([]int32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddNotifyRefIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.ExamYear(); !ok {
		return &ValidationError{Name: "ExamYear", err: errors.New(`ent: missing required field "Notification.ExamYear"`)}
	}
	if _, ok := nc.mutation.ApplicationStartDate(); !ok {
		return &ValidationError{Name: "ApplicationStartDate", err: errors.New(`ent: missing required field "Notification.ApplicationStartDate"`)}
	}
	if _, ok := nc.mutation.ApplicationEndDate(); !ok {
		return &ValidationError{Name: "ApplicationEndDate", err: errors.New(`ent: missing required field "Notification.ApplicationEndDate"`)}
	}
	if _, ok := nc.mutation.VerificationDateByController(); !ok {
		return &ValidationError{Name: "VerificationDateByController", err: errors.New(`ent: missing required field "Notification.VerificationDateByController"`)}
	}
	if _, ok := nc.mutation.CorrectionDateByCandidate(); !ok {
		return &ValidationError{Name: "CorrectionDateByCandidate", err: errors.New(`ent: missing required field "Notification.CorrectionDateByCandidate"`)}
	}
	if _, ok := nc.mutation.CorrectionVeriyDateByController(); !ok {
		return &ValidationError{Name: "CorrectionVeriyDateByController", err: errors.New(`ent: missing required field "Notification.CorrectionVeriyDateByController"`)}
	}
	if _, ok := nc.mutation.HallTicketAllotmentDateByNodalOfficer(); !ok {
		return &ValidationError{Name: "HallTicketAllotmentDateByNodalOfficer", err: errors.New(`ent: missing required field "Notification.HallTicketAllotmentDateByNodalOfficer"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notification.Table, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.ExamYear(); ok {
		_spec.SetField(notification.FieldExamYear, field.TypeInt32, value)
		_node.ExamYear = value
	}
	if value, ok := nc.mutation.ApplicationStartDate(); ok {
		_spec.SetField(notification.FieldApplicationStartDate, field.TypeTime, value)
		_node.ApplicationStartDate = value
	}
	if value, ok := nc.mutation.ApplicationEndDate(); ok {
		_spec.SetField(notification.FieldApplicationEndDate, field.TypeTime, value)
		_node.ApplicationEndDate = value
	}
	if value, ok := nc.mutation.VerificationDateByController(); ok {
		_spec.SetField(notification.FieldVerificationDateByController, field.TypeTime, value)
		_node.VerificationDateByController = value
	}
	if value, ok := nc.mutation.CorrectionDateByCandidate(); ok {
		_spec.SetField(notification.FieldCorrectionDateByCandidate, field.TypeTime, value)
		_node.CorrectionDateByCandidate = value
	}
	if value, ok := nc.mutation.CorrectionVeriyDateByController(); ok {
		_spec.SetField(notification.FieldCorrectionVeriyDateByController, field.TypeTime, value)
		_node.CorrectionVeriyDateByController = value
	}
	if value, ok := nc.mutation.HallTicketAllotmentDateByNodalOfficer(); ok {
		_spec.SetField(notification.FieldHallTicketAllotmentDateByNodalOfficer, field.TypeTime, value)
		_node.HallTicketAllotmentDateByNodalOfficer = value
	}
	if value, ok := nc.mutation.HallTicketDownloadDate(); ok {
		_spec.SetField(notification.FieldHallTicketDownloadDate, field.TypeTime, value)
		_node.HallTicketDownloadDate = value
	}
	if value, ok := nc.mutation.NotifyFile(); ok {
		_spec.SetField(notification.FieldNotifyFile, field.TypeString, value)
		_node.NotifyFile = value
	}
	if value, ok := nc.mutation.SyllabusFile(); ok {
		_spec.SetField(notification.FieldSyllabusFile, field.TypeString, value)
		_node.SyllabusFile = value
	}
	if value, ok := nc.mutation.VacanciesFile(); ok {
		_spec.SetField(notification.FieldVacanciesFile, field.TypeString, value)
		_node.VacanciesFile = value
	}
	if nodes := nc.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notification.ApplicationsTable,
			Columns: []string{notification.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.CentersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notification.CentersTable,
			Columns: []string{notification.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NodalOfficersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notification.NodalOfficersTable,
			Columns: []string{notification.NodalOfficersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodalofficer.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.ExamTable,
			Columns: []string{notification.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExamCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.VacancyYearsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   notification.VacancyYearsTable,
			Columns: []string{notification.VacancyYearsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vacancyyear.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NotifyRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notification.NotifyRefTable,
			Columns: notification.NotifyRefPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	builders []*NotificationCreate
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
