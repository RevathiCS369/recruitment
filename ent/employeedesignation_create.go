// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/employeedesignation"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeDesignationCreate is the builder for creating a EmployeeDesignation entity.
type EmployeeDesignationCreate struct {
	config
	mutation *EmployeeDesignationMutation
	hooks    []Hook
}

// SetDesignationCode sets the "DesignationCode" field.
func (edc *EmployeeDesignationCreate) SetDesignationCode(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDesignationCode(s)
	return edc
}

// SetDesignationDescription sets the "DesignationDescription" field.
func (edc *EmployeeDesignationCreate) SetDesignationDescription(s string) *EmployeeDesignationCreate {
	edc.mutation.SetDesignationDescription(s)
	return edc
}

// SetID sets the "id" field.
func (edc *EmployeeDesignationCreate) SetID(i int32) *EmployeeDesignationCreate {
	edc.mutation.SetID(i)
	return edc
}

// Mutation returns the EmployeeDesignationMutation object of the builder.
func (edc *EmployeeDesignationCreate) Mutation() *EmployeeDesignationMutation {
	return edc.mutation
}

// Save creates the EmployeeDesignation in the database.
func (edc *EmployeeDesignationCreate) Save(ctx context.Context) (*EmployeeDesignation, error) {
	return withHooks(ctx, edc.sqlSave, edc.mutation, edc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (edc *EmployeeDesignationCreate) SaveX(ctx context.Context) *EmployeeDesignation {
	v, err := edc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edc *EmployeeDesignationCreate) Exec(ctx context.Context) error {
	_, err := edc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edc *EmployeeDesignationCreate) ExecX(ctx context.Context) {
	if err := edc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (edc *EmployeeDesignationCreate) check() error {
	if _, ok := edc.mutation.DesignationCode(); !ok {
		return &ValidationError{Name: "DesignationCode", err: errors.New(`ent: missing required field "EmployeeDesignation.DesignationCode"`)}
	}
	if _, ok := edc.mutation.DesignationDescription(); !ok {
		return &ValidationError{Name: "DesignationDescription", err: errors.New(`ent: missing required field "EmployeeDesignation.DesignationDescription"`)}
	}
	return nil
}

func (edc *EmployeeDesignationCreate) sqlSave(ctx context.Context) (*EmployeeDesignation, error) {
	if err := edc.check(); err != nil {
		return nil, err
	}
	_node, _spec := edc.createSpec()
	if err := sqlgraph.CreateNode(ctx, edc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	edc.mutation.id = &_node.ID
	edc.mutation.done = true
	return _node, nil
}

func (edc *EmployeeDesignationCreate) createSpec() (*EmployeeDesignation, *sqlgraph.CreateSpec) {
	var (
		_node = &EmployeeDesignation{config: edc.config}
		_spec = sqlgraph.NewCreateSpec(employeedesignation.Table, sqlgraph.NewFieldSpec(employeedesignation.FieldID, field.TypeInt32))
	)
	if id, ok := edc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := edc.mutation.DesignationCode(); ok {
		_spec.SetField(employeedesignation.FieldDesignationCode, field.TypeString, value)
		_node.DesignationCode = value
	}
	if value, ok := edc.mutation.DesignationDescription(); ok {
		_spec.SetField(employeedesignation.FieldDesignationDescription, field.TypeString, value)
		_node.DesignationDescription = value
	}
	return _node, _spec
}

// EmployeeDesignationCreateBulk is the builder for creating many EmployeeDesignation entities in bulk.
type EmployeeDesignationCreateBulk struct {
	config
	builders []*EmployeeDesignationCreate
}

// Save creates the EmployeeDesignation entities in the database.
func (edcb *EmployeeDesignationCreateBulk) Save(ctx context.Context) ([]*EmployeeDesignation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(edcb.builders))
	nodes := make([]*EmployeeDesignation, len(edcb.builders))
	mutators := make([]Mutator, len(edcb.builders))
	for i := range edcb.builders {
		func(i int, root context.Context) {
			builder := edcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeeDesignationMutation)
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
					_, err = mutators[i+1].Mutate(root, edcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, edcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, edcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (edcb *EmployeeDesignationCreateBulk) SaveX(ctx context.Context) []*EmployeeDesignation {
	v, err := edcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (edcb *EmployeeDesignationCreateBulk) Exec(ctx context.Context) error {
	_, err := edcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (edcb *EmployeeDesignationCreateBulk) ExecX(ctx context.Context) {
	if err := edcb.Exec(ctx); err != nil {
		panic(err)
	}
}