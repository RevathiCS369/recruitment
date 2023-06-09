// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/application"
	"recruit/ent/center"
	"recruit/ent/notification"
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApplicationUpdate is the builder for updating Application entities.
type ApplicationUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (au *ApplicationUpdate) Where(ps ...predicate.Application) *ApplicationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetEmployeeID sets the "EmployeeID" field.
func (au *ApplicationUpdate) SetEmployeeID(i int32) *ApplicationUpdate {
	au.mutation.ResetEmployeeID()
	au.mutation.SetEmployeeID(i)
	return au
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (au *ApplicationUpdate) AddEmployeeID(i int32) *ApplicationUpdate {
	au.mutation.AddEmployeeID(i)
	return au
}

// SetNotifyCode sets the "NotifyCode" field.
func (au *ApplicationUpdate) SetNotifyCode(i int32) *ApplicationUpdate {
	au.mutation.SetNotifyCode(i)
	return au
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableNotifyCode(i *int32) *ApplicationUpdate {
	if i != nil {
		au.SetNotifyCode(*i)
	}
	return au
}

// ClearNotifyCode clears the value of the "NotifyCode" field.
func (au *ApplicationUpdate) ClearNotifyCode() *ApplicationUpdate {
	au.mutation.ClearNotifyCode()
	return au
}

// SetHallTicketNumber sets the "HallTicketNumber" field.
func (au *ApplicationUpdate) SetHallTicketNumber(s string) *ApplicationUpdate {
	au.mutation.SetHallTicketNumber(s)
	return au
}

// SetNillableHallTicketNumber sets the "HallTicketNumber" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableHallTicketNumber(s *string) *ApplicationUpdate {
	if s != nil {
		au.SetHallTicketNumber(*s)
	}
	return au
}

// ClearHallTicketNumber clears the value of the "HallTicketNumber" field.
func (au *ApplicationUpdate) ClearHallTicketNumber() *ApplicationUpdate {
	au.mutation.ClearHallTicketNumber()
	return au
}

// SetCenterCode sets the "CenterCode" field.
func (au *ApplicationUpdate) SetCenterCode(i int32) *ApplicationUpdate {
	au.mutation.SetCenterCode(i)
	return au
}

// SetNillableCenterCode sets the "CenterCode" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableCenterCode(i *int32) *ApplicationUpdate {
	if i != nil {
		au.SetCenterCode(*i)
	}
	return au
}

// ClearCenterCode clears the value of the "CenterCode" field.
func (au *ApplicationUpdate) ClearCenterCode() *ApplicationUpdate {
	au.mutation.ClearCenterCode()
	return au
}

// SetAppliedStamp sets the "AppliedStamp" field.
func (au *ApplicationUpdate) SetAppliedStamp(t time.Time) *ApplicationUpdate {
	au.mutation.SetAppliedStamp(t)
	return au
}

// SetCenterID sets the "center" edge to the Center entity by ID.
func (au *ApplicationUpdate) SetCenterID(id int32) *ApplicationUpdate {
	au.mutation.SetCenterID(id)
	return au
}

// SetNillableCenterID sets the "center" edge to the Center entity by ID if the given value is not nil.
func (au *ApplicationUpdate) SetNillableCenterID(id *int32) *ApplicationUpdate {
	if id != nil {
		au = au.SetCenterID(*id)
	}
	return au
}

// SetCenter sets the "center" edge to the Center entity.
func (au *ApplicationUpdate) SetCenter(c *Center) *ApplicationUpdate {
	return au.SetCenterID(c.ID)
}

// SetNotificationID sets the "notification" edge to the Notification entity by ID.
func (au *ApplicationUpdate) SetNotificationID(id int32) *ApplicationUpdate {
	au.mutation.SetNotificationID(id)
	return au
}

// SetNillableNotificationID sets the "notification" edge to the Notification entity by ID if the given value is not nil.
func (au *ApplicationUpdate) SetNillableNotificationID(id *int32) *ApplicationUpdate {
	if id != nil {
		au = au.SetNotificationID(*id)
	}
	return au
}

// SetNotification sets the "notification" edge to the Notification entity.
func (au *ApplicationUpdate) SetNotification(n *Notification) *ApplicationUpdate {
	return au.SetNotificationID(n.ID)
}

// Mutation returns the ApplicationMutation object of the builder.
func (au *ApplicationUpdate) Mutation() *ApplicationMutation {
	return au.mutation
}

// ClearCenter clears the "center" edge to the Center entity.
func (au *ApplicationUpdate) ClearCenter() *ApplicationUpdate {
	au.mutation.ClearCenter()
	return au
}

// ClearNotification clears the "notification" edge to the Notification entity.
func (au *ApplicationUpdate) ClearNotification() *ApplicationUpdate {
	au.mutation.ClearNotification()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApplicationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApplicationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApplicationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeInt32))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.EmployeeID(); ok {
		_spec.SetField(application.FieldEmployeeID, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedEmployeeID(); ok {
		_spec.AddField(application.FieldEmployeeID, field.TypeInt32, value)
	}
	if value, ok := au.mutation.HallTicketNumber(); ok {
		_spec.SetField(application.FieldHallTicketNumber, field.TypeString, value)
	}
	if au.mutation.HallTicketNumberCleared() {
		_spec.ClearField(application.FieldHallTicketNumber, field.TypeString)
	}
	if value, ok := au.mutation.AppliedStamp(); ok {
		_spec.SetField(application.FieldAppliedStamp, field.TypeTime, value)
	}
	if au.mutation.CenterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.CenterTable,
			Columns: []string{application.CenterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CenterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.CenterTable,
			Columns: []string{application.CenterColumn},
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
	if au.mutation.NotificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.NotificationTable,
			Columns: []string{application.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.NotificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.NotificationTable,
			Columns: []string{application.NotificationColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApplicationUpdateOne is the builder for updating a single Application entity.
type ApplicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationMutation
}

// SetEmployeeID sets the "EmployeeID" field.
func (auo *ApplicationUpdateOne) SetEmployeeID(i int32) *ApplicationUpdateOne {
	auo.mutation.ResetEmployeeID()
	auo.mutation.SetEmployeeID(i)
	return auo
}

// AddEmployeeID adds i to the "EmployeeID" field.
func (auo *ApplicationUpdateOne) AddEmployeeID(i int32) *ApplicationUpdateOne {
	auo.mutation.AddEmployeeID(i)
	return auo
}

// SetNotifyCode sets the "NotifyCode" field.
func (auo *ApplicationUpdateOne) SetNotifyCode(i int32) *ApplicationUpdateOne {
	auo.mutation.SetNotifyCode(i)
	return auo
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableNotifyCode(i *int32) *ApplicationUpdateOne {
	if i != nil {
		auo.SetNotifyCode(*i)
	}
	return auo
}

// ClearNotifyCode clears the value of the "NotifyCode" field.
func (auo *ApplicationUpdateOne) ClearNotifyCode() *ApplicationUpdateOne {
	auo.mutation.ClearNotifyCode()
	return auo
}

// SetHallTicketNumber sets the "HallTicketNumber" field.
func (auo *ApplicationUpdateOne) SetHallTicketNumber(s string) *ApplicationUpdateOne {
	auo.mutation.SetHallTicketNumber(s)
	return auo
}

// SetNillableHallTicketNumber sets the "HallTicketNumber" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableHallTicketNumber(s *string) *ApplicationUpdateOne {
	if s != nil {
		auo.SetHallTicketNumber(*s)
	}
	return auo
}

// ClearHallTicketNumber clears the value of the "HallTicketNumber" field.
func (auo *ApplicationUpdateOne) ClearHallTicketNumber() *ApplicationUpdateOne {
	auo.mutation.ClearHallTicketNumber()
	return auo
}

// SetCenterCode sets the "CenterCode" field.
func (auo *ApplicationUpdateOne) SetCenterCode(i int32) *ApplicationUpdateOne {
	auo.mutation.SetCenterCode(i)
	return auo
}

// SetNillableCenterCode sets the "CenterCode" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableCenterCode(i *int32) *ApplicationUpdateOne {
	if i != nil {
		auo.SetCenterCode(*i)
	}
	return auo
}

// ClearCenterCode clears the value of the "CenterCode" field.
func (auo *ApplicationUpdateOne) ClearCenterCode() *ApplicationUpdateOne {
	auo.mutation.ClearCenterCode()
	return auo
}

// SetAppliedStamp sets the "AppliedStamp" field.
func (auo *ApplicationUpdateOne) SetAppliedStamp(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetAppliedStamp(t)
	return auo
}

// SetCenterID sets the "center" edge to the Center entity by ID.
func (auo *ApplicationUpdateOne) SetCenterID(id int32) *ApplicationUpdateOne {
	auo.mutation.SetCenterID(id)
	return auo
}

// SetNillableCenterID sets the "center" edge to the Center entity by ID if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableCenterID(id *int32) *ApplicationUpdateOne {
	if id != nil {
		auo = auo.SetCenterID(*id)
	}
	return auo
}

// SetCenter sets the "center" edge to the Center entity.
func (auo *ApplicationUpdateOne) SetCenter(c *Center) *ApplicationUpdateOne {
	return auo.SetCenterID(c.ID)
}

// SetNotificationID sets the "notification" edge to the Notification entity by ID.
func (auo *ApplicationUpdateOne) SetNotificationID(id int32) *ApplicationUpdateOne {
	auo.mutation.SetNotificationID(id)
	return auo
}

// SetNillableNotificationID sets the "notification" edge to the Notification entity by ID if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableNotificationID(id *int32) *ApplicationUpdateOne {
	if id != nil {
		auo = auo.SetNotificationID(*id)
	}
	return auo
}

// SetNotification sets the "notification" edge to the Notification entity.
func (auo *ApplicationUpdateOne) SetNotification(n *Notification) *ApplicationUpdateOne {
	return auo.SetNotificationID(n.ID)
}

// Mutation returns the ApplicationMutation object of the builder.
func (auo *ApplicationUpdateOne) Mutation() *ApplicationMutation {
	return auo.mutation
}

// ClearCenter clears the "center" edge to the Center entity.
func (auo *ApplicationUpdateOne) ClearCenter() *ApplicationUpdateOne {
	auo.mutation.ClearCenter()
	return auo
}

// ClearNotification clears the "notification" edge to the Notification entity.
func (auo *ApplicationUpdateOne) ClearNotification() *ApplicationUpdateOne {
	auo.mutation.ClearNotification()
	return auo
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (auo *ApplicationUpdateOne) Where(ps ...predicate.Application) *ApplicationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApplicationUpdateOne) Select(field string, fields ...string) *ApplicationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Application entity.
func (auo *ApplicationUpdateOne) Save(ctx context.Context) (*Application, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApplicationUpdateOne) SaveX(ctx context.Context) *Application {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ApplicationUpdateOne) sqlSave(ctx context.Context) (_node *Application, err error) {
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeInt32))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Application.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, application.FieldID)
		for _, f := range fields {
			if !application.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != application.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.EmployeeID(); ok {
		_spec.SetField(application.FieldEmployeeID, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedEmployeeID(); ok {
		_spec.AddField(application.FieldEmployeeID, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.HallTicketNumber(); ok {
		_spec.SetField(application.FieldHallTicketNumber, field.TypeString, value)
	}
	if auo.mutation.HallTicketNumberCleared() {
		_spec.ClearField(application.FieldHallTicketNumber, field.TypeString)
	}
	if value, ok := auo.mutation.AppliedStamp(); ok {
		_spec.SetField(application.FieldAppliedStamp, field.TypeTime, value)
	}
	if auo.mutation.CenterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.CenterTable,
			Columns: []string{application.CenterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CenterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.CenterTable,
			Columns: []string{application.CenterColumn},
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
	if auo.mutation.NotificationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.NotificationTable,
			Columns: []string{application.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.NotificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.NotificationTable,
			Columns: []string{application.NotificationColumn},
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
	_node = &Application{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
