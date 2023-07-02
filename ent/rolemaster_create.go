// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/adminlogin"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/rolemaster"
	"recruit/ent/usermaster"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMasterCreate is the builder for creating a RoleMaster entity.
type RoleMasterCreate struct {
	config
	mutation *RoleMasterMutation
	hooks    []Hook
}

// SetRoleName sets the "RoleName" field.
func (rmc *RoleMasterCreate) SetRoleName(s string) *RoleMasterCreate {
	rmc.mutation.SetRoleName(s)
	return rmc
}

// SetCreatedDate sets the "CreatedDate" field.
func (rmc *RoleMasterCreate) SetCreatedDate(t time.Time) *RoleMasterCreate {
	rmc.mutation.SetCreatedDate(t)
	return rmc
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (rmc *RoleMasterCreate) SetNillableCreatedDate(t *time.Time) *RoleMasterCreate {
	if t != nil {
		rmc.SetCreatedDate(*t)
	}
	return rmc
}

// SetStatus sets the "Status" field.
func (rmc *RoleMasterCreate) SetStatus(b bool) *RoleMasterCreate {
	rmc.mutation.SetStatus(b)
	return rmc
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (rmc *RoleMasterCreate) SetNillableStatus(b *bool) *RoleMasterCreate {
	if b != nil {
		rmc.SetStatus(*b)
	}
	return rmc
}

// SetID sets the "id" field.
func (rmc *RoleMasterCreate) SetID(i int32) *RoleMasterCreate {
	rmc.mutation.SetID(i)
	return rmc
}

// AddRoleIDs adds the "roles" edge to the AdminLogin entity by IDs.
func (rmc *RoleMasterCreate) AddRoleIDs(ids ...int32) *RoleMasterCreate {
	rmc.mutation.AddRoleIDs(ids...)
	return rmc
}

// AddRoles adds the "roles" edges to the AdminLogin entity.
func (rmc *RoleMasterCreate) AddRoles(a ...*AdminLogin) *RoleMasterCreate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rmc.AddRoleIDs(ids...)
}

// AddRolesRefIDs adds the "Roles_Ref" edge to the UserMaster entity by IDs.
func (rmc *RoleMasterCreate) AddRolesRefIDs(ids ...int64) *RoleMasterCreate {
	rmc.mutation.AddRolesRefIDs(ids...)
	return rmc
}

// AddRolesRef adds the "Roles_Ref" edges to the UserMaster entity.
func (rmc *RoleMasterCreate) AddRolesRef(u ...*UserMaster) *RoleMasterCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rmc.AddRolesRefIDs(ids...)
}

// AddRolesPSRefIDs adds the "Roles_PS_Ref" edge to the Exam_Applications_PS entity by IDs.
func (rmc *RoleMasterCreate) AddRolesPSRefIDs(ids ...int64) *RoleMasterCreate {
	rmc.mutation.AddRolesPSRefIDs(ids...)
	return rmc
}

// AddRolesPSRef adds the "Roles_PS_Ref" edges to the Exam_Applications_PS entity.
func (rmc *RoleMasterCreate) AddRolesPSRef(e ...*Exam_Applications_PS) *RoleMasterCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmc.AddRolesPSRefIDs(ids...)
}

// AddRolesIPRefIDs adds the "Roles_IP_Ref" edge to the Exam_Applications_IP entity by IDs.
func (rmc *RoleMasterCreate) AddRolesIPRefIDs(ids ...int64) *RoleMasterCreate {
	rmc.mutation.AddRolesIPRefIDs(ids...)
	return rmc
}

// AddRolesIPRef adds the "Roles_IP_Ref" edges to the Exam_Applications_IP entity.
func (rmc *RoleMasterCreate) AddRolesIPRef(e ...*Exam_Applications_IP) *RoleMasterCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmc.AddRolesIPRefIDs(ids...)
}

// Mutation returns the RoleMasterMutation object of the builder.
func (rmc *RoleMasterCreate) Mutation() *RoleMasterMutation {
	return rmc.mutation
}

// Save creates the RoleMaster in the database.
func (rmc *RoleMasterCreate) Save(ctx context.Context) (*RoleMaster, error) {
	rmc.defaults()
	return withHooks(ctx, rmc.sqlSave, rmc.mutation, rmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rmc *RoleMasterCreate) SaveX(ctx context.Context) *RoleMaster {
	v, err := rmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmc *RoleMasterCreate) Exec(ctx context.Context) error {
	_, err := rmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmc *RoleMasterCreate) ExecX(ctx context.Context) {
	if err := rmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmc *RoleMasterCreate) defaults() {
	if _, ok := rmc.mutation.Status(); !ok {
		v := rolemaster.DefaultStatus
		rmc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmc *RoleMasterCreate) check() error {
	if _, ok := rmc.mutation.RoleName(); !ok {
		return &ValidationError{Name: "RoleName", err: errors.New(`ent: missing required field "RoleMaster.RoleName"`)}
	}
	if _, ok := rmc.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "RoleMaster.Status"`)}
	}
	return nil
}

func (rmc *RoleMasterCreate) sqlSave(ctx context.Context) (*RoleMaster, error) {
	if err := rmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	rmc.mutation.id = &_node.ID
	rmc.mutation.done = true
	return _node, nil
}

func (rmc *RoleMasterCreate) createSpec() (*RoleMaster, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleMaster{config: rmc.config}
		_spec = sqlgraph.NewCreateSpec(rolemaster.Table, sqlgraph.NewFieldSpec(rolemaster.FieldID, field.TypeInt32))
	)
	if id, ok := rmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rmc.mutation.RoleName(); ok {
		_spec.SetField(rolemaster.FieldRoleName, field.TypeString, value)
		_node.RoleName = value
	}
	if value, ok := rmc.mutation.CreatedDate(); ok {
		_spec.SetField(rolemaster.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if value, ok := rmc.mutation.Status(); ok {
		_spec.SetField(rolemaster.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if nodes := rmc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rmc.mutation.RolesRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rmc.mutation.RolesPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rmc.mutation.RolesIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleMasterCreateBulk is the builder for creating many RoleMaster entities in bulk.
type RoleMasterCreateBulk struct {
	config
	builders []*RoleMasterCreate
}

// Save creates the RoleMaster entities in the database.
func (rmcb *RoleMasterCreateBulk) Save(ctx context.Context) ([]*RoleMaster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rmcb.builders))
	nodes := make([]*RoleMaster, len(rmcb.builders))
	mutators := make([]Mutator, len(rmcb.builders))
	for i := range rmcb.builders {
		func(i int, root context.Context) {
			builder := rmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMasterMutation)
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
					_, err = mutators[i+1].Mutate(root, rmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rmcb *RoleMasterCreateBulk) SaveX(ctx context.Context) []*RoleMaster {
	v, err := rmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmcb *RoleMasterCreateBulk) Exec(ctx context.Context) error {
	_, err := rmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmcb *RoleMasterCreateBulk) ExecX(ctx context.Context) {
	if err := rmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
