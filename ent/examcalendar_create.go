// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exam"
	"recruit/ent/exam_ip"
	"recruit/ent/exam_ps"
	"recruit/ent/examcalendar"
	"recruit/ent/exampapers"
	"recruit/ent/notification"
	"recruit/ent/vacancyyear"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCalendarCreate is the builder for creating a ExamCalendar entity.
type ExamCalendarCreate struct {
	config
	mutation *ExamCalendarMutation
	hooks    []Hook
}

// SetExamYear sets the "ExamYear" field.
func (ecc *ExamCalendarCreate) SetExamYear(i int32) *ExamCalendarCreate {
	ecc.mutation.SetExamYear(i)
	return ecc
}

// SetExamName sets the "ExamName" field.
func (ecc *ExamCalendarCreate) SetExamName(s string) *ExamCalendarCreate {
	ecc.mutation.SetExamName(s)
	return ecc
}

// SetExamCode sets the "ExamCode" field.
func (ecc *ExamCalendarCreate) SetExamCode(i int32) *ExamCalendarCreate {
	ecc.mutation.SetExamCode(i)
	return ecc
}

// SetNillableExamCode sets the "ExamCode" field if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableExamCode(i *int32) *ExamCalendarCreate {
	if i != nil {
		ecc.SetExamCode(*i)
	}
	return ecc
}

// SetNotificationDate sets the "NotificationDate" field.
func (ecc *ExamCalendarCreate) SetNotificationDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetNotificationDate(t)
	return ecc
}

// SetModelNotificationDate sets the "ModelNotificationDate" field.
func (ecc *ExamCalendarCreate) SetModelNotificationDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetModelNotificationDate(t)
	return ecc
}

// SetApplicationEndDate sets the "ApplicationEndDate" field.
func (ecc *ExamCalendarCreate) SetApplicationEndDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetApplicationEndDate(t)
	return ecc
}

// SetApprovedOrderDate sets the "ApprovedOrderDate" field.
func (ecc *ExamCalendarCreate) SetApprovedOrderDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetApprovedOrderDate(t)
	return ecc
}

// SetTentativeResultDate sets the "TentativeResultDate" field.
func (ecc *ExamCalendarCreate) SetTentativeResultDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetTentativeResultDate(t)
	return ecc
}

// SetNillableTentativeResultDate sets the "TentativeResultDate" field if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableTentativeResultDate(t *time.Time) *ExamCalendarCreate {
	if t != nil {
		ecc.SetTentativeResultDate(*t)
	}
	return ecc
}

// SetCreatedDate sets the "CreatedDate" field.
func (ecc *ExamCalendarCreate) SetCreatedDate(t time.Time) *ExamCalendarCreate {
	ecc.mutation.SetCreatedDate(t)
	return ecc
}

// SetApprovedOrderNumber sets the "ApprovedOrderNumber" field.
func (ecc *ExamCalendarCreate) SetApprovedOrderNumber(s string) *ExamCalendarCreate {
	ecc.mutation.SetApprovedOrderNumber(s)
	return ecc
}

// SetVacancyYears sets the "VacancyYears" field.
func (ecc *ExamCalendarCreate) SetVacancyYears(i []interface{}) *ExamCalendarCreate {
	ecc.mutation.SetVacancyYears(i)
	return ecc
}

// SetExamPapers sets the "ExamPapers" field.
func (ecc *ExamCalendarCreate) SetExamPapers(i []interface{}) *ExamCalendarCreate {
	ecc.mutation.SetExamPapers(i)
	return ecc
}

// SetVacancyYearCode sets the "VacancyYearCode" field.
func (ecc *ExamCalendarCreate) SetVacancyYearCode(i int32) *ExamCalendarCreate {
	ecc.mutation.SetVacancyYearCode(i)
	return ecc
}

// SetNillableVacancyYearCode sets the "VacancyYearCode" field if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableVacancyYearCode(i *int32) *ExamCalendarCreate {
	if i != nil {
		ecc.SetVacancyYearCode(*i)
	}
	return ecc
}

// SetPaperCode sets the "PaperCode" field.
func (ecc *ExamCalendarCreate) SetPaperCode(i int32) *ExamCalendarCreate {
	ecc.mutation.SetPaperCode(i)
	return ecc
}

// SetNillablePaperCode sets the "PaperCode" field if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillablePaperCode(i *int32) *ExamCalendarCreate {
	if i != nil {
		ecc.SetPaperCode(*i)
	}
	return ecc
}

// SetExamCodePS sets the "ExamCodePS" field.
func (ecc *ExamCalendarCreate) SetExamCodePS(i int32) *ExamCalendarCreate {
	ecc.mutation.SetExamCodePS(i)
	return ecc
}

// SetNillableExamCodePS sets the "ExamCodePS" field if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableExamCodePS(i *int32) *ExamCalendarCreate {
	if i != nil {
		ecc.SetExamCodePS(*i)
	}
	return ecc
}

// SetID sets the "id" field.
func (ecc *ExamCalendarCreate) SetID(i int32) *ExamCalendarCreate {
	ecc.mutation.SetID(i)
	return ecc
}

// SetVcyYearsID sets the "vcy_years" edge to the VacancyYear entity by ID.
func (ecc *ExamCalendarCreate) SetVcyYearsID(id int32) *ExamCalendarCreate {
	ecc.mutation.SetVcyYearsID(id)
	return ecc
}

// SetNillableVcyYearsID sets the "vcy_years" edge to the VacancyYear entity by ID if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableVcyYearsID(id *int32) *ExamCalendarCreate {
	if id != nil {
		ecc = ecc.SetVcyYearsID(*id)
	}
	return ecc
}

// SetVcyYears sets the "vcy_years" edge to the VacancyYear entity.
func (ecc *ExamCalendarCreate) SetVcyYears(v *VacancyYear) *ExamCalendarCreate {
	return ecc.SetVcyYearsID(v.ID)
}

// SetExamsID sets the "exams" edge to the Exam entity by ID.
func (ecc *ExamCalendarCreate) SetExamsID(id int32) *ExamCalendarCreate {
	ecc.mutation.SetExamsID(id)
	return ecc
}

// SetNillableExamsID sets the "exams" edge to the Exam entity by ID if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillableExamsID(id *int32) *ExamCalendarCreate {
	if id != nil {
		ecc = ecc.SetExamsID(*id)
	}
	return ecc
}

// SetExams sets the "exams" edge to the Exam entity.
func (ecc *ExamCalendarCreate) SetExams(e *Exam) *ExamCalendarCreate {
	return ecc.SetExamsID(e.ID)
}

// SetPapersID sets the "papers" edge to the ExamPapers entity by ID.
func (ecc *ExamCalendarCreate) SetPapersID(id int32) *ExamCalendarCreate {
	ecc.mutation.SetPapersID(id)
	return ecc
}

// SetNillablePapersID sets the "papers" edge to the ExamPapers entity by ID if the given value is not nil.
func (ecc *ExamCalendarCreate) SetNillablePapersID(id *int32) *ExamCalendarCreate {
	if id != nil {
		ecc = ecc.SetPapersID(*id)
	}
	return ecc
}

// SetPapers sets the "papers" edge to the ExamPapers entity.
func (ecc *ExamCalendarCreate) SetPapers(e *ExamPapers) *ExamCalendarCreate {
	return ecc.SetPapersID(e.ID)
}

// AddNotifyRefIDs adds the "Notify_ref" edge to the Notification entity by IDs.
func (ecc *ExamCalendarCreate) AddNotifyRefIDs(ids ...int32) *ExamCalendarCreate {
	ecc.mutation.AddNotifyRefIDs(ids...)
	return ecc
}

// AddNotifyRef adds the "Notify_ref" edges to the Notification entity.
func (ecc *ExamCalendarCreate) AddNotifyRef(n ...*Notification) *ExamCalendarCreate {
	ids := make([]int32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ecc.AddNotifyRefIDs(ids...)
}

// AddExamcalPsRefIDs adds the "examcal_ps_ref" edge to the Exam_PS entity by IDs.
func (ecc *ExamCalendarCreate) AddExamcalPsRefIDs(ids ...int32) *ExamCalendarCreate {
	ecc.mutation.AddExamcalPsRefIDs(ids...)
	return ecc
}

// AddExamcalPsRef adds the "examcal_ps_ref" edges to the Exam_PS entity.
func (ecc *ExamCalendarCreate) AddExamcalPsRef(e ...*Exam_PS) *ExamCalendarCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddExamcalPsRefIDs(ids...)
}

// AddExamcalIPRefIDs adds the "examcal_ip_ref" edge to the Exam_IP entity by IDs.
func (ecc *ExamCalendarCreate) AddExamcalIPRefIDs(ids ...int32) *ExamCalendarCreate {
	ecc.mutation.AddExamcalIPRefIDs(ids...)
	return ecc
}

// AddExamcalIPRef adds the "examcal_ip_ref" edges to the Exam_IP entity.
func (ecc *ExamCalendarCreate) AddExamcalIPRef(e ...*Exam_IP) *ExamCalendarCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddExamcalIPRefIDs(ids...)
}

// Mutation returns the ExamCalendarMutation object of the builder.
func (ecc *ExamCalendarCreate) Mutation() *ExamCalendarMutation {
	return ecc.mutation
}

// Save creates the ExamCalendar in the database.
func (ecc *ExamCalendarCreate) Save(ctx context.Context) (*ExamCalendar, error) {
	return withHooks(ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *ExamCalendarCreate) SaveX(ctx context.Context) *ExamCalendar {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *ExamCalendarCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *ExamCalendarCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *ExamCalendarCreate) check() error {
	if _, ok := ecc.mutation.ExamYear(); !ok {
		return &ValidationError{Name: "ExamYear", err: errors.New(`ent: missing required field "ExamCalendar.ExamYear"`)}
	}
	if _, ok := ecc.mutation.ExamName(); !ok {
		return &ValidationError{Name: "ExamName", err: errors.New(`ent: missing required field "ExamCalendar.ExamName"`)}
	}
	if v, ok := ecc.mutation.ExamName(); ok {
		if err := examcalendar.ExamNameValidator(v); err != nil {
			return &ValidationError{Name: "ExamName", err: fmt.Errorf(`ent: validator failed for field "ExamCalendar.ExamName": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.NotificationDate(); !ok {
		return &ValidationError{Name: "NotificationDate", err: errors.New(`ent: missing required field "ExamCalendar.NotificationDate"`)}
	}
	if _, ok := ecc.mutation.ModelNotificationDate(); !ok {
		return &ValidationError{Name: "ModelNotificationDate", err: errors.New(`ent: missing required field "ExamCalendar.ModelNotificationDate"`)}
	}
	if _, ok := ecc.mutation.ApplicationEndDate(); !ok {
		return &ValidationError{Name: "ApplicationEndDate", err: errors.New(`ent: missing required field "ExamCalendar.ApplicationEndDate"`)}
	}
	if _, ok := ecc.mutation.ApprovedOrderDate(); !ok {
		return &ValidationError{Name: "ApprovedOrderDate", err: errors.New(`ent: missing required field "ExamCalendar.ApprovedOrderDate"`)}
	}
	if _, ok := ecc.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "CreatedDate", err: errors.New(`ent: missing required field "ExamCalendar.CreatedDate"`)}
	}
	if _, ok := ecc.mutation.ApprovedOrderNumber(); !ok {
		return &ValidationError{Name: "ApprovedOrderNumber", err: errors.New(`ent: missing required field "ExamCalendar.ApprovedOrderNumber"`)}
	}
	return nil
}

func (ecc *ExamCalendarCreate) sqlSave(ctx context.Context) (*ExamCalendar, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *ExamCalendarCreate) createSpec() (*ExamCalendar, *sqlgraph.CreateSpec) {
	var (
		_node = &ExamCalendar{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(examcalendar.Table, sqlgraph.NewFieldSpec(examcalendar.FieldID, field.TypeInt32))
	)
	if id, ok := ecc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ecc.mutation.ExamYear(); ok {
		_spec.SetField(examcalendar.FieldExamYear, field.TypeInt32, value)
		_node.ExamYear = value
	}
	if value, ok := ecc.mutation.ExamName(); ok {
		_spec.SetField(examcalendar.FieldExamName, field.TypeString, value)
		_node.ExamName = value
	}
	if value, ok := ecc.mutation.NotificationDate(); ok {
		_spec.SetField(examcalendar.FieldNotificationDate, field.TypeTime, value)
		_node.NotificationDate = value
	}
	if value, ok := ecc.mutation.ModelNotificationDate(); ok {
		_spec.SetField(examcalendar.FieldModelNotificationDate, field.TypeTime, value)
		_node.ModelNotificationDate = value
	}
	if value, ok := ecc.mutation.ApplicationEndDate(); ok {
		_spec.SetField(examcalendar.FieldApplicationEndDate, field.TypeTime, value)
		_node.ApplicationEndDate = value
	}
	if value, ok := ecc.mutation.ApprovedOrderDate(); ok {
		_spec.SetField(examcalendar.FieldApprovedOrderDate, field.TypeTime, value)
		_node.ApprovedOrderDate = value
	}
	if value, ok := ecc.mutation.TentativeResultDate(); ok {
		_spec.SetField(examcalendar.FieldTentativeResultDate, field.TypeTime, value)
		_node.TentativeResultDate = value
	}
	if value, ok := ecc.mutation.CreatedDate(); ok {
		_spec.SetField(examcalendar.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if value, ok := ecc.mutation.ApprovedOrderNumber(); ok {
		_spec.SetField(examcalendar.FieldApprovedOrderNumber, field.TypeString, value)
		_node.ApprovedOrderNumber = value
	}
	if value, ok := ecc.mutation.VacancyYears(); ok {
		_spec.SetField(examcalendar.FieldVacancyYears, field.TypeJSON, value)
		_node.VacancyYears = value
	}
	if value, ok := ecc.mutation.ExamPapers(); ok {
		_spec.SetField(examcalendar.FieldExamPapers, field.TypeJSON, value)
		_node.ExamPapers = value
	}
	if value, ok := ecc.mutation.ExamCodePS(); ok {
		_spec.SetField(examcalendar.FieldExamCodePS, field.TypeInt32, value)
		_node.ExamCodePS = value
	}
	if nodes := ecc.mutation.VcyYearsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examcalendar.VcyYearsTable,
			Columns: []string{examcalendar.VcyYearsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vacancyyear.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VacancyYearCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examcalendar.ExamsTable,
			Columns: []string{examcalendar.ExamsColumn},
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
	if nodes := ecc.mutation.PapersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examcalendar.PapersTable,
			Columns: []string{examcalendar.PapersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exampapers.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PaperCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.NotifyRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcalendar.NotifyRefTable,
			Columns: []string{examcalendar.NotifyRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.ExamcalPsRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcalendar.ExamcalPsRefTable,
			Columns: []string{examcalendar.ExamcalPsRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_ps.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.ExamcalIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcalendar.ExamcalIPRefTable,
			Columns: []string{examcalendar.ExamcalIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_ip.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExamCalendarCreateBulk is the builder for creating many ExamCalendar entities in bulk.
type ExamCalendarCreateBulk struct {
	config
	builders []*ExamCalendarCreate
}

// Save creates the ExamCalendar entities in the database.
func (eccb *ExamCalendarCreateBulk) Save(ctx context.Context) ([]*ExamCalendar, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*ExamCalendar, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExamCalendarMutation)
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
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *ExamCalendarCreateBulk) SaveX(ctx context.Context) []*ExamCalendar {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *ExamCalendarCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *ExamCalendarCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}
