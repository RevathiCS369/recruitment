// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/application"
	"recruit/ent/center"
	"recruit/ent/nodalofficer"
	"recruit/ent/notification"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CenterCreate is the builder for creating a Center entity.
type CenterCreate struct {
	config
	mutation *CenterMutation
	hooks    []Hook
}

// SetNotifyCode sets the "NotifyCode" field.
func (cc *CenterCreate) SetNotifyCode(i int32) *CenterCreate {
	cc.mutation.SetNotifyCode(i)
	return cc
}

// SetNillableNotifyCode sets the "NotifyCode" field if the given value is not nil.
func (cc *CenterCreate) SetNillableNotifyCode(i *int32) *CenterCreate {
	if i != nil {
		cc.SetNotifyCode(*i)
	}
	return cc
}

// SetNodalOfficerCode sets the "NodalOfficerCode" field.
func (cc *CenterCreate) SetNodalOfficerCode(i int32) *CenterCreate {
	cc.mutation.SetNodalOfficerCode(i)
	return cc
}

// SetNillableNodalOfficerCode sets the "NodalOfficerCode" field if the given value is not nil.
func (cc *CenterCreate) SetNillableNodalOfficerCode(i *int32) *CenterCreate {
	if i != nil {
		cc.SetNodalOfficerCode(*i)
	}
	return cc
}

// SetCenterName sets the "CenterName" field.
func (cc *CenterCreate) SetCenterName(s string) *CenterCreate {
	cc.mutation.SetCenterName(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CenterCreate) SetID(i int32) *CenterCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddApplicationIDs adds the "applications" edge to the Application entity by IDs.
func (cc *CenterCreate) AddApplicationIDs(ids ...int32) *CenterCreate {
	cc.mutation.AddApplicationIDs(ids...)
	return cc
}

// AddApplications adds the "applications" edges to the Application entity.
func (cc *CenterCreate) AddApplications(a ...*Application) *CenterCreate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddApplicationIDs(ids...)
}

// SetNodalOfficerID sets the "nodal_officer" edge to the NodalOfficer entity by ID.
func (cc *CenterCreate) SetNodalOfficerID(id int32) *CenterCreate {
	cc.mutation.SetNodalOfficerID(id)
	return cc
}

// SetNillableNodalOfficerID sets the "nodal_officer" edge to the NodalOfficer entity by ID if the given value is not nil.
func (cc *CenterCreate) SetNillableNodalOfficerID(id *int32) *CenterCreate {
	if id != nil {
		cc = cc.SetNodalOfficerID(*id)
	}
	return cc
}

// SetNodalOfficer sets the "nodal_officer" edge to the NodalOfficer entity.
func (cc *CenterCreate) SetNodalOfficer(n *NodalOfficer) *CenterCreate {
	return cc.SetNodalOfficerID(n.ID)
}

// SetNotificationID sets the "notification" edge to the Notification entity by ID.
func (cc *CenterCreate) SetNotificationID(id int32) *CenterCreate {
	cc.mutation.SetNotificationID(id)
	return cc
}

// SetNillableNotificationID sets the "notification" edge to the Notification entity by ID if the given value is not nil.
func (cc *CenterCreate) SetNillableNotificationID(id *int32) *CenterCreate {
	if id != nil {
		cc = cc.SetNotificationID(*id)
	}
	return cc
}

// SetNotification sets the "notification" edge to the Notification entity.
func (cc *CenterCreate) SetNotification(n *Notification) *CenterCreate {
	return cc.SetNotificationID(n.ID)
}

// Mutation returns the CenterMutation object of the builder.
func (cc *CenterCreate) Mutation() *CenterMutation {
	return cc.mutation
}

// Save creates the Center in the database.
func (cc *CenterCreate) Save(ctx context.Context) (*Center, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CenterCreate) SaveX(ctx context.Context) *Center {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CenterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CenterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CenterCreate) check() error {
	if _, ok := cc.mutation.CenterName(); !ok {
		return &ValidationError{Name: "CenterName", err: errors.New(`ent: missing required field "Center.CenterName"`)}
	}
	return nil
}

func (cc *CenterCreate) sqlSave(ctx context.Context) (*Center, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CenterCreate) createSpec() (*Center, *sqlgraph.CreateSpec) {
	var (
		_node = &Center{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(center.Table, sqlgraph.NewFieldSpec(center.FieldID, field.TypeInt32))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CenterName(); ok {
		_spec.SetField(center.FieldCenterName, field.TypeString, value)
		_node.CenterName = value
	}
	if nodes := cc.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   center.ApplicationsTable,
			Columns: []string{center.ApplicationsColumn},
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
	if nodes := cc.mutation.NodalOfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   center.NodalOfficerTable,
			Columns: []string{center.NodalOfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nodalofficer.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NodalOfficerCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NotificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   center.NotificationTable,
			Columns: []string{center.NotificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NotifyCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CenterCreateBulk is the builder for creating many Center entities in bulk.
type CenterCreateBulk struct {
	config
	builders []*CenterCreate
}

// Save creates the Center entities in the database.
func (ccb *CenterCreateBulk) Save(ctx context.Context) ([]*Center, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Center, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CenterMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CenterCreateBulk) SaveX(ctx context.Context) []*Center {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CenterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CenterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
