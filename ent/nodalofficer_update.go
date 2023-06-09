// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/center"
	"recruit/ent/exam"
	"recruit/ent/nodalofficer"
	"recruit/ent/notification"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NodalOfficerUpdate is the builder for updating NodalOfficer entities.
type NodalOfficerUpdate struct {
	config
	hooks    []Hook
	mutation *NodalOfficerMutation
}

// Where appends a list predicates to the NodalOfficerUpdate builder.
func (nou *NodalOfficerUpdate) Where(ps ...predicate.NodalOfficer) *NodalOfficerUpdate {
	nou.mutation.Where(ps...)
	return nou
}

// SetNodalOfficerName sets the "NodalOfficerName" field.
func (nou *NodalOfficerUpdate) SetNodalOfficerName(s string) *NodalOfficerUpdate {
	nou.mutation.SetNodalOfficerName(s)
	return nou
}

// SetDesignationID sets the "DesignationID" field.
func (nou *NodalOfficerUpdate) SetDesignationID(i int32) *NodalOfficerUpdate {
	nou.mutation.ResetDesignationID()
	nou.mutation.SetDesignationID(i)
	return nou
}

// AddDesignationID adds i to the "DesignationID" field.
func (nou *NodalOfficerUpdate) AddDesignationID(i int32) *NodalOfficerUpdate {
	nou.mutation.AddDesignationID(i)
	return nou
}

// SetEmailID sets the "EmailID" field.
func (nou *NodalOfficerUpdate) SetEmailID(s string) *NodalOfficerUpdate {
	nou.mutation.SetEmailID(s)
	return nou
}

// SetMobileNumber sets the "MobileNumber" field.
func (nou *NodalOfficerUpdate) SetMobileNumber(s string) *NodalOfficerUpdate {
	nou.mutation.SetMobileNumber(s)
	return nou
}

// SetNotifyCode sets the "NotifyCode" field.
func (nou *NodalOfficerUpdate) SetNotifyCode(i int32) *NodalOfficerUpdate {
	nou.mutation.SetNotifyCode(i)
	return nou
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (nou *NodalOfficerUpdate) SetNillableNotifyCode(i *int32) *NodalOfficerUpdate {
	if i != nil {
		nou.SetNotifyCode(*i)
	}
	return nou
}

// ClearNotifyCode clears the value of the "NotifyCode" field.
func (nou *NodalOfficerUpdate) ClearNotifyCode() *NodalOfficerUpdate {
	nou.mutation.ClearNotifyCode()
	return nou
}

// SetExamCode sets the "ExamCode" field.
func (nou *NodalOfficerUpdate) SetExamCode(i int32) *NodalOfficerUpdate {
	nou.mutation.SetExamCode(i)
	return nou
}

// SetNillableExamCode sets the "ExamCode" field if the given value is not nil.
func (nou *NodalOfficerUpdate) SetNillableExamCode(i *int32) *NodalOfficerUpdate {
	if i != nil {
		nou.SetExamCode(*i)
	}
	return nou
}

// ClearExamCode clears the value of the "ExamCode" field.
func (nou *NodalOfficerUpdate) ClearExamCode() *NodalOfficerUpdate {
	nou.mutation.ClearExamCode()
	return nou
}

// SetHallTicketApproved sets the "HallTicketApproved" field.
func (nou *NodalOfficerUpdate) SetHallTicketApproved(s string) *NodalOfficerUpdate {
	nou.mutation.SetHallTicketApproved(s)
	return nou
}

// SetNillableHallTicketApproved sets the "HallTicketApproved" field if the given value is not nil.
func (nou *NodalOfficerUpdate) SetNillableHallTicketApproved(s *string) *NodalOfficerUpdate {
	if s != nil {
		nou.SetHallTicketApproved(*s)
	}
	return nou
}

// ClearHallTicketApproved clears the value of the "HallTicketApproved" field.
func (nou *NodalOfficerUpdate) ClearHallTicketApproved() *NodalOfficerUpdate {
	nou.mutation.ClearHallTicketApproved()
	return nou
}

// AddCenterIDs adds the "centers" edge to the Center entity by IDs.
func (nou *NodalOfficerUpdate) AddCenterIDs(ids ...int32) *NodalOfficerUpdate {
	nou.mutation.AddCenterIDs(ids...)
	return nou
}

// AddCenters adds the "centers" edges to the Center entity.
func (nou *NodalOfficerUpdate) AddCenters(c ...*Center) *NodalOfficerUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nou.AddCenterIDs(ids...)
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (nou *NodalOfficerUpdate) SetExamID(id int32) *NodalOfficerUpdate {
	nou.mutation.SetExamID(id)
	return nou
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (nou *NodalOfficerUpdate) SetNillableExamID(id *int32) *NodalOfficerUpdate {
	if id != nil {
		nou = nou.SetExamID(*id)
	}
	return nou
}

// SetExam sets the "exam" edge to the Exam entity.
func (nou *NodalOfficerUpdate) SetExam(e *Exam) *NodalOfficerUpdate {
	return nou.SetExamID(e.ID)
}

// SetNotificationID sets the "notification" edge to the Notification entity by ID.
func (nou *NodalOfficerUpdate) SetNotificationID(id int32) *NodalOfficerUpdate {
	nou.mutation.SetNotificationID(id)
	return nou
}

// SetNillableNotificationID sets the "notification" edge to the Notification entity by ID if the given value is not nil.
func (nou *NodalOfficerUpdate) SetNillableNotificationID(id *int32) *NodalOfficerUpdate {
	if id != nil {
		nou = nou.SetNotificationID(*id)
	}
	return nou
}

// SetNotification sets the "notification" edge to the Notification entity.
func (nou *NodalOfficerUpdate) SetNotification(n *Notification) *NodalOfficerUpdate {
	return nou.SetNotificationID(n.ID)
}

// Mutation returns the NodalOfficerMutation object of the builder.
func (nou *NodalOfficerUpdate) Mutation() *NodalOfficerMutation {
	return nou.mutation
}

// ClearCenters clears all "centers" edges to the Center entity.
func (nou *NodalOfficerUpdate) ClearCenters() *NodalOfficerUpdate {
	nou.mutation.ClearCenters()
	return nou
}

// RemoveCenterIDs removes the "centers" edge to Center entities by IDs.
func (nou *NodalOfficerUpdate) RemoveCenterIDs(ids ...int32) *NodalOfficerUpdate {
	nou.mutation.RemoveCenterIDs(ids...)
	return nou
}

// RemoveCenters removes "centers" edges to Center entities.
func (nou *NodalOfficerUpdate) RemoveCenters(c ...*Center) *NodalOfficerUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nou.RemoveCenterIDs(ids...)
}

// ClearExam clears the "exam" edge to the Exam entity.
func (nou *NodalOfficerUpdate) ClearExam() *NodalOfficerUpdate {
	nou.mutation.ClearExam()
	return nou
}

// ClearNotification clears the "notification" edge to the Notification entity.
func (nou *NodalOfficerUpdate) ClearNotification() *NodalOfficerUpdate {
	nou.mutation.ClearNotification()
	return nou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nou *NodalOfficerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nou.sqlSave, nou.mutation, nou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nou *NodalOfficerUpdate) SaveX(ctx context.Context) int {
	affected, err := nou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nou *NodalOfficerUpdate) Exec(ctx context.Context) error {
	_, err := nou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nou *NodalOfficerUpdate) ExecX(ctx context.Context) {
	if err := nou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nou *NodalOfficerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(nodalofficer.Table, nodalofficer.Columns, sqlgraph.NewFieldSpec(nodalofficer.FieldID, field.TypeInt32))
	if ps := nou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nou.mutation.NodalOfficerName(); ok {
		_spec.SetField(nodalofficer.FieldNodalOfficerName, field.TypeString, value)
	}
	if value, ok := nou.mutation.DesignationID(); ok {
		_spec.SetField(nodalofficer.FieldDesignationID, field.TypeInt32, value)
	}
	if value, ok := nou.mutation.AddedDesignationID(); ok {
		_spec.AddField(nodalofficer.FieldDesignationID, field.TypeInt32, value)
	}
	if value, ok := nou.mutation.EmailID(); ok {
		_spec.SetField(nodalofficer.FieldEmailID, field.TypeString, value)
	}
	if value, ok := nou.mutation.MobileNumber(); ok {
		_spec.SetField(nodalofficer.FieldMobileNumber, field.TypeString, value)
	}
	if value, ok := nou.mutation.HallTicketApproved(); ok {
		_spec.SetField(nodalofficer.FieldHallTicketApproved, field.TypeString, value)
	}
	if nou.mutation.HallTicketApprovedCleared() {
		_spec.ClearField(nodalofficer.FieldHallTicketApproved, field.TypeString)
	}
	if nou.mutation.CentersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nou.mutation.RemovedCentersIDs(); len(nodes) > 0 && !nou.mutation.CentersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nou.mutation.CentersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nou.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.ExamTable,
			Columns: []string{nodalofficer.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nou.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.ExamTable,
			Columns: []string{nodalofficer.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nou.mutation.NotificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.NotificationTable,
			Columns: []string{nodalofficer.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nou.mutation.NotificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.NotificationTable,
			Columns: []string{nodalofficer.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodalofficer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nou.mutation.done = true
	return n, nil
}

// NodalOfficerUpdateOne is the builder for updating a single NodalOfficer entity.
type NodalOfficerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodalOfficerMutation
}

// SetNodalOfficerName sets the "NodalOfficerName" field.
func (nouo *NodalOfficerUpdateOne) SetNodalOfficerName(s string) *NodalOfficerUpdateOne {
	nouo.mutation.SetNodalOfficerName(s)
	return nouo
}

// SetDesignationID sets the "DesignationID" field.
func (nouo *NodalOfficerUpdateOne) SetDesignationID(i int32) *NodalOfficerUpdateOne {
	nouo.mutation.ResetDesignationID()
	nouo.mutation.SetDesignationID(i)
	return nouo
}

// AddDesignationID adds i to the "DesignationID" field.
func (nouo *NodalOfficerUpdateOne) AddDesignationID(i int32) *NodalOfficerUpdateOne {
	nouo.mutation.AddDesignationID(i)
	return nouo
}

// SetEmailID sets the "EmailID" field.
func (nouo *NodalOfficerUpdateOne) SetEmailID(s string) *NodalOfficerUpdateOne {
	nouo.mutation.SetEmailID(s)
	return nouo
}

// SetMobileNumber sets the "MobileNumber" field.
func (nouo *NodalOfficerUpdateOne) SetMobileNumber(s string) *NodalOfficerUpdateOne {
	nouo.mutation.SetMobileNumber(s)
	return nouo
}

// SetNotifyCode sets the "NotifyCode" field.
func (nouo *NodalOfficerUpdateOne) SetNotifyCode(i int32) *NodalOfficerUpdateOne {
	nouo.mutation.SetNotifyCode(i)
	return nouo
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (nouo *NodalOfficerUpdateOne) SetNillableNotifyCode(i *int32) *NodalOfficerUpdateOne {
	if i != nil {
		nouo.SetNotifyCode(*i)
	}
	return nouo
}

// ClearNotifyCode clears the value of the "NotifyCode" field.
func (nouo *NodalOfficerUpdateOne) ClearNotifyCode() *NodalOfficerUpdateOne {
	nouo.mutation.ClearNotifyCode()
	return nouo
}

// SetExamCode sets the "ExamCode" field.
func (nouo *NodalOfficerUpdateOne) SetExamCode(i int32) *NodalOfficerUpdateOne {
	nouo.mutation.SetExamCode(i)
	return nouo
}

// SetNillableExamCode sets the "ExamCode" field if the given value is not nil.
func (nouo *NodalOfficerUpdateOne) SetNillableExamCode(i *int32) *NodalOfficerUpdateOne {
	if i != nil {
		nouo.SetExamCode(*i)
	}
	return nouo
}

// ClearExamCode clears the value of the "ExamCode" field.
func (nouo *NodalOfficerUpdateOne) ClearExamCode() *NodalOfficerUpdateOne {
	nouo.mutation.ClearExamCode()
	return nouo
}

// SetHallTicketApproved sets the "HallTicketApproved" field.
func (nouo *NodalOfficerUpdateOne) SetHallTicketApproved(s string) *NodalOfficerUpdateOne {
	nouo.mutation.SetHallTicketApproved(s)
	return nouo
}

// SetNillableHallTicketApproved sets the "HallTicketApproved" field if the given value is not nil.
func (nouo *NodalOfficerUpdateOne) SetNillableHallTicketApproved(s *string) *NodalOfficerUpdateOne {
	if s != nil {
		nouo.SetHallTicketApproved(*s)
	}
	return nouo
}

// ClearHallTicketApproved clears the value of the "HallTicketApproved" field.
func (nouo *NodalOfficerUpdateOne) ClearHallTicketApproved() *NodalOfficerUpdateOne {
	nouo.mutation.ClearHallTicketApproved()
	return nouo
}

// AddCenterIDs adds the "centers" edge to the Center entity by IDs.
func (nouo *NodalOfficerUpdateOne) AddCenterIDs(ids ...int32) *NodalOfficerUpdateOne {
	nouo.mutation.AddCenterIDs(ids...)
	return nouo
}

// AddCenters adds the "centers" edges to the Center entity.
func (nouo *NodalOfficerUpdateOne) AddCenters(c ...*Center) *NodalOfficerUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nouo.AddCenterIDs(ids...)
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (nouo *NodalOfficerUpdateOne) SetExamID(id int32) *NodalOfficerUpdateOne {
	nouo.mutation.SetExamID(id)
	return nouo
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (nouo *NodalOfficerUpdateOne) SetNillableExamID(id *int32) *NodalOfficerUpdateOne {
	if id != nil {
		nouo = nouo.SetExamID(*id)
	}
	return nouo
}

// SetExam sets the "exam" edge to the Exam entity.
func (nouo *NodalOfficerUpdateOne) SetExam(e *Exam) *NodalOfficerUpdateOne {
	return nouo.SetExamID(e.ID)
}

// SetNotificationID sets the "notification" edge to the Notification entity by ID.
func (nouo *NodalOfficerUpdateOne) SetNotificationID(id int32) *NodalOfficerUpdateOne {
	nouo.mutation.SetNotificationID(id)
	return nouo
}

// SetNillableNotificationID sets the "notification" edge to the Notification entity by ID if the given value is not nil.
func (nouo *NodalOfficerUpdateOne) SetNillableNotificationID(id *int32) *NodalOfficerUpdateOne {
	if id != nil {
		nouo = nouo.SetNotificationID(*id)
	}
	return nouo
}

// SetNotification sets the "notification" edge to the Notification entity.
func (nouo *NodalOfficerUpdateOne) SetNotification(n *Notification) *NodalOfficerUpdateOne {
	return nouo.SetNotificationID(n.ID)
}

// Mutation returns the NodalOfficerMutation object of the builder.
func (nouo *NodalOfficerUpdateOne) Mutation() *NodalOfficerMutation {
	return nouo.mutation
}

// ClearCenters clears all "centers" edges to the Center entity.
func (nouo *NodalOfficerUpdateOne) ClearCenters() *NodalOfficerUpdateOne {
	nouo.mutation.ClearCenters()
	return nouo
}

// RemoveCenterIDs removes the "centers" edge to Center entities by IDs.
func (nouo *NodalOfficerUpdateOne) RemoveCenterIDs(ids ...int32) *NodalOfficerUpdateOne {
	nouo.mutation.RemoveCenterIDs(ids...)
	return nouo
}

// RemoveCenters removes "centers" edges to Center entities.
func (nouo *NodalOfficerUpdateOne) RemoveCenters(c ...*Center) *NodalOfficerUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return nouo.RemoveCenterIDs(ids...)
}

// ClearExam clears the "exam" edge to the Exam entity.
func (nouo *NodalOfficerUpdateOne) ClearExam() *NodalOfficerUpdateOne {
	nouo.mutation.ClearExam()
	return nouo
}

// ClearNotification clears the "notification" edge to the Notification entity.
func (nouo *NodalOfficerUpdateOne) ClearNotification() *NodalOfficerUpdateOne {
	nouo.mutation.ClearNotification()
	return nouo
}

// Where appends a list predicates to the NodalOfficerUpdate builder.
func (nouo *NodalOfficerUpdateOne) Where(ps ...predicate.NodalOfficer) *NodalOfficerUpdateOne {
	nouo.mutation.Where(ps...)
	return nouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nouo *NodalOfficerUpdateOne) Select(field string, fields ...string) *NodalOfficerUpdateOne {
	nouo.fields = append([]string{field}, fields...)
	return nouo
}

// Save executes the query and returns the updated NodalOfficer entity.
func (nouo *NodalOfficerUpdateOne) Save(ctx context.Context) (*NodalOfficer, error) {
	return withHooks(ctx, nouo.sqlSave, nouo.mutation, nouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nouo *NodalOfficerUpdateOne) SaveX(ctx context.Context) *NodalOfficer {
	node, err := nouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nouo *NodalOfficerUpdateOne) Exec(ctx context.Context) error {
	_, err := nouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nouo *NodalOfficerUpdateOne) ExecX(ctx context.Context) {
	if err := nouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nouo *NodalOfficerUpdateOne) sqlSave(ctx context.Context) (_node *NodalOfficer, err error) {
	_spec := sqlgraph.NewUpdateSpec(nodalofficer.Table, nodalofficer.Columns, sqlgraph.NewFieldSpec(nodalofficer.FieldID, field.TypeInt32))
	id, ok := nouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NodalOfficer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodalofficer.FieldID)
		for _, f := range fields {
			if !nodalofficer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nodalofficer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nouo.mutation.NodalOfficerName(); ok {
		_spec.SetField(nodalofficer.FieldNodalOfficerName, field.TypeString, value)
	}
	if value, ok := nouo.mutation.DesignationID(); ok {
		_spec.SetField(nodalofficer.FieldDesignationID, field.TypeInt32, value)
	}
	if value, ok := nouo.mutation.AddedDesignationID(); ok {
		_spec.AddField(nodalofficer.FieldDesignationID, field.TypeInt32, value)
	}
	if value, ok := nouo.mutation.EmailID(); ok {
		_spec.SetField(nodalofficer.FieldEmailID, field.TypeString, value)
	}
	if value, ok := nouo.mutation.MobileNumber(); ok {
		_spec.SetField(nodalofficer.FieldMobileNumber, field.TypeString, value)
	}
	if value, ok := nouo.mutation.HallTicketApproved(); ok {
		_spec.SetField(nodalofficer.FieldHallTicketApproved, field.TypeString, value)
	}
	if nouo.mutation.HallTicketApprovedCleared() {
		_spec.ClearField(nodalofficer.FieldHallTicketApproved, field.TypeString)
	}
	if nouo.mutation.CentersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nouo.mutation.RemovedCentersIDs(); len(nodes) > 0 && !nouo.mutation.CentersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nouo.mutation.CentersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodalofficer.CentersTable,
			Columns: []string{nodalofficer.CentersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nouo.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.ExamTable,
			Columns: []string{nodalofficer.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nouo.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.ExamTable,
			Columns: []string{nodalofficer.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nouo.mutation.NotificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.NotificationTable,
			Columns: []string{nodalofficer.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nouo.mutation.NotificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nodalofficer.NotificationTable,
			Columns: []string{nodalofficer.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NodalOfficer{config: nouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodalofficer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nouo.mutation.done = true
	return _node, nil
}
