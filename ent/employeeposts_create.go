// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/employeeposts"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeePostsCreate is the builder for creating a EmployeePosts entity.
type EmployeePostsCreate struct {
	config
	mutation *EmployeePostsMutation
	hooks    []Hook
}

// SetPostCode sets the "PostCode" field.
func (epc *EmployeePostsCreate) SetPostCode(s string) *EmployeePostsCreate {
	epc.mutation.SetPostCode(s)
	return epc
}

// SetPostDescription sets the "PostDescription" field.
func (epc *EmployeePostsCreate) SetPostDescription(s string) *EmployeePostsCreate {
	epc.mutation.SetPostDescription(s)
	return epc
}

// SetGroup sets the "Group" field.
func (epc *EmployeePostsCreate) SetGroup(s string) *EmployeePostsCreate {
	epc.mutation.SetGroup(s)
	return epc
}

// SetPayLevel sets the "PayLevel" field.
func (epc *EmployeePostsCreate) SetPayLevel(s string) *EmployeePostsCreate {
	epc.mutation.SetPayLevel(s)
	return epc
}

// SetScale sets the "Scale" field.
func (epc *EmployeePostsCreate) SetScale(s string) *EmployeePostsCreate {
	epc.mutation.SetScale(s)
	return epc
}

// SetBaseCadreFlag sets the "BaseCadreFlag" field.
func (epc *EmployeePostsCreate) SetBaseCadreFlag(b bool) *EmployeePostsCreate {
	epc.mutation.SetBaseCadreFlag(b)
	return epc
}

// SetID sets the "id" field.
func (epc *EmployeePostsCreate) SetID(i int32) *EmployeePostsCreate {
	epc.mutation.SetID(i)
	return epc
}

// Mutation returns the EmployeePostsMutation object of the builder.
func (epc *EmployeePostsCreate) Mutation() *EmployeePostsMutation {
	return epc.mutation
}

// Save creates the EmployeePosts in the database.
func (epc *EmployeePostsCreate) Save(ctx context.Context) (*EmployeePosts, error) {
	return withHooks(ctx, epc.sqlSave, epc.mutation, epc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (epc *EmployeePostsCreate) SaveX(ctx context.Context) *EmployeePosts {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epc *EmployeePostsCreate) Exec(ctx context.Context) error {
	_, err := epc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epc *EmployeePostsCreate) ExecX(ctx context.Context) {
	if err := epc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epc *EmployeePostsCreate) check() error {
	if _, ok := epc.mutation.PostCode(); !ok {
		return &ValidationError{Name: "PostCode", err: errors.New(`ent: missing required field "EmployeePosts.PostCode"`)}
	}
	if _, ok := epc.mutation.PostDescription(); !ok {
		return &ValidationError{Name: "PostDescription", err: errors.New(`ent: missing required field "EmployeePosts.PostDescription"`)}
	}
	if _, ok := epc.mutation.Group(); !ok {
		return &ValidationError{Name: "Group", err: errors.New(`ent: missing required field "EmployeePosts.Group"`)}
	}
	if _, ok := epc.mutation.PayLevel(); !ok {
		return &ValidationError{Name: "PayLevel", err: errors.New(`ent: missing required field "EmployeePosts.PayLevel"`)}
	}
	if _, ok := epc.mutation.Scale(); !ok {
		return &ValidationError{Name: "Scale", err: errors.New(`ent: missing required field "EmployeePosts.Scale"`)}
	}
	if _, ok := epc.mutation.BaseCadreFlag(); !ok {
		return &ValidationError{Name: "BaseCadreFlag", err: errors.New(`ent: missing required field "EmployeePosts.BaseCadreFlag"`)}
	}
	return nil
}

func (epc *EmployeePostsCreate) sqlSave(ctx context.Context) (*EmployeePosts, error) {
	if err := epc.check(); err != nil {
		return nil, err
	}
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	epc.mutation.id = &_node.ID
	epc.mutation.done = true
	return _node, nil
}

func (epc *EmployeePostsCreate) createSpec() (*EmployeePosts, *sqlgraph.CreateSpec) {
	var (
		_node = &EmployeePosts{config: epc.config}
		_spec = sqlgraph.NewCreateSpec(employeeposts.Table, sqlgraph.NewFieldSpec(employeeposts.FieldID, field.TypeInt32))
	)
	if id, ok := epc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := epc.mutation.PostCode(); ok {
		_spec.SetField(employeeposts.FieldPostCode, field.TypeString, value)
		_node.PostCode = value
	}
	if value, ok := epc.mutation.PostDescription(); ok {
		_spec.SetField(employeeposts.FieldPostDescription, field.TypeString, value)
		_node.PostDescription = value
	}
	if value, ok := epc.mutation.Group(); ok {
		_spec.SetField(employeeposts.FieldGroup, field.TypeString, value)
		_node.Group = value
	}
	if value, ok := epc.mutation.PayLevel(); ok {
		_spec.SetField(employeeposts.FieldPayLevel, field.TypeString, value)
		_node.PayLevel = value
	}
	if value, ok := epc.mutation.Scale(); ok {
		_spec.SetField(employeeposts.FieldScale, field.TypeString, value)
		_node.Scale = value
	}
	if value, ok := epc.mutation.BaseCadreFlag(); ok {
		_spec.SetField(employeeposts.FieldBaseCadreFlag, field.TypeBool, value)
		_node.BaseCadreFlag = value
	}
	return _node, _spec
}

// EmployeePostsCreateBulk is the builder for creating many EmployeePosts entities in bulk.
type EmployeePostsCreateBulk struct {
	config
	builders []*EmployeePostsCreate
}

// Save creates the EmployeePosts entities in the database.
func (epcb *EmployeePostsCreateBulk) Save(ctx context.Context) ([]*EmployeePosts, error) {
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*EmployeePosts, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmployeePostsMutation)
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
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (epcb *EmployeePostsCreateBulk) SaveX(ctx context.Context) []*EmployeePosts {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epcb *EmployeePostsCreateBulk) Exec(ctx context.Context) error {
	_, err := epcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epcb *EmployeePostsCreateBulk) ExecX(ctx context.Context) {
	if err := epcb.Exec(ctx); err != nil {
		panic(err)
	}
}
