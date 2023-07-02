// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/division_choice_pa"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionChoicePACreate is the builder for creating a Division_Choice_PA entity.
type DivisionChoicePACreate struct {
	config
	mutation *DivisionChoicePAMutation
	hooks    []Hook
}

// SetApplication sets the "Application" field.
func (dcpc *DivisionChoicePACreate) SetApplication(s string) *DivisionChoicePACreate {
	dcpc.mutation.SetApplication(s)
	return dcpc
}

// SetNillableApplication sets the "Application" field if the given value is not nil.
func (dcpc *DivisionChoicePACreate) SetNillableApplication(s *string) *DivisionChoicePACreate {
	if s != nil {
		dcpc.SetApplication(*s)
	}
	return dcpc
}

// SetCadrePrefNo sets the "CadrePrefNo" field.
func (dcpc *DivisionChoicePACreate) SetCadrePrefNo(s string) *DivisionChoicePACreate {
	dcpc.mutation.SetCadrePrefNo(s)
	return dcpc
}

// SetCadrePrefValue sets the "CadrePrefValue" field.
func (dcpc *DivisionChoicePACreate) SetCadrePrefValue(s string) *DivisionChoicePACreate {
	dcpc.mutation.SetCadrePrefValue(s)
	return dcpc
}

// SetEmployeeID sets the "EmployeeID" field.
func (dcpc *DivisionChoicePACreate) SetEmployeeID(i int64) *DivisionChoicePACreate {
	dcpc.mutation.SetEmployeeID(i)
	return dcpc
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (dcpc *DivisionChoicePACreate) SetNillableEmployeeID(i *int64) *DivisionChoicePACreate {
	if i != nil {
		dcpc.SetEmployeeID(*i)
	}
	return dcpc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (dcpc *DivisionChoicePACreate) SetUpdatedAt(t time.Time) *DivisionChoicePACreate {
	dcpc.mutation.SetUpdatedAt(t)
	return dcpc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (dcpc *DivisionChoicePACreate) SetNillableUpdatedAt(t *time.Time) *DivisionChoicePACreate {
	if t != nil {
		dcpc.SetUpdatedAt(*t)
	}
	return dcpc
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (dcpc *DivisionChoicePACreate) SetUpdatedBy(s string) *DivisionChoicePACreate {
	dcpc.mutation.SetUpdatedBy(s)
	return dcpc
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (dcpc *DivisionChoicePACreate) SetNillableUpdatedBy(s *string) *DivisionChoicePACreate {
	if s != nil {
		dcpc.SetUpdatedBy(*s)
	}
	return dcpc
}

// SetID sets the "id" field.
func (dcpc *DivisionChoicePACreate) SetID(i int32) *DivisionChoicePACreate {
	dcpc.mutation.SetID(i)
	return dcpc
}

// Mutation returns the DivisionChoicePAMutation object of the builder.
func (dcpc *DivisionChoicePACreate) Mutation() *DivisionChoicePAMutation {
	return dcpc.mutation
}

// Save creates the Division_Choice_PA in the database.
func (dcpc *DivisionChoicePACreate) Save(ctx context.Context) (*Division_Choice_PA, error) {
	dcpc.defaults()
	return withHooks(ctx, dcpc.sqlSave, dcpc.mutation, dcpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dcpc *DivisionChoicePACreate) SaveX(ctx context.Context) *Division_Choice_PA {
	v, err := dcpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcpc *DivisionChoicePACreate) Exec(ctx context.Context) error {
	_, err := dcpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpc *DivisionChoicePACreate) ExecX(ctx context.Context) {
	if err := dcpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcpc *DivisionChoicePACreate) defaults() {
	if _, ok := dcpc.mutation.UpdatedAt(); !ok {
		v := division_choice_pa.DefaultUpdatedAt()
		dcpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dcpc.mutation.UpdatedBy(); !ok {
		v := division_choice_pa.DefaultUpdatedBy
		dcpc.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcpc *DivisionChoicePACreate) check() error {
	if _, ok := dcpc.mutation.CadrePrefNo(); !ok {
		return &ValidationError{Name: "CadrePrefNo", err: errors.New(`ent: missing required field "Division_Choice_PA.CadrePrefNo"`)}
	}
	if _, ok := dcpc.mutation.CadrePrefValue(); !ok {
		return &ValidationError{Name: "CadrePrefValue", err: errors.New(`ent: missing required field "Division_Choice_PA.CadrePrefValue"`)}
	}
	return nil
}

func (dcpc *DivisionChoicePACreate) sqlSave(ctx context.Context) (*Division_Choice_PA, error) {
	if err := dcpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dcpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	dcpc.mutation.id = &_node.ID
	dcpc.mutation.done = true
	return _node, nil
}

func (dcpc *DivisionChoicePACreate) createSpec() (*Division_Choice_PA, *sqlgraph.CreateSpec) {
	var (
		_node = &Division_Choice_PA{config: dcpc.config}
		_spec = sqlgraph.NewCreateSpec(division_choice_pa.Table, sqlgraph.NewFieldSpec(division_choice_pa.FieldID, field.TypeInt32))
	)
	if id, ok := dcpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dcpc.mutation.Application(); ok {
		_spec.SetField(division_choice_pa.FieldApplication, field.TypeString, value)
		_node.Application = value
	}
	if value, ok := dcpc.mutation.CadrePrefNo(); ok {
		_spec.SetField(division_choice_pa.FieldCadrePrefNo, field.TypeString, value)
		_node.CadrePrefNo = value
	}
	if value, ok := dcpc.mutation.CadrePrefValue(); ok {
		_spec.SetField(division_choice_pa.FieldCadrePrefValue, field.TypeString, value)
		_node.CadrePrefValue = value
	}
	if value, ok := dcpc.mutation.EmployeeID(); ok {
		_spec.SetField(division_choice_pa.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := dcpc.mutation.UpdatedAt(); ok {
		_spec.SetField(division_choice_pa.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dcpc.mutation.UpdatedBy(); ok {
		_spec.SetField(division_choice_pa.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	return _node, _spec
}

// DivisionChoicePACreateBulk is the builder for creating many Division_Choice_PA entities in bulk.
type DivisionChoicePACreateBulk struct {
	config
	builders []*DivisionChoicePACreate
}

// Save creates the Division_Choice_PA entities in the database.
func (dcpcb *DivisionChoicePACreateBulk) Save(ctx context.Context) ([]*Division_Choice_PA, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcpcb.builders))
	nodes := make([]*Division_Choice_PA, len(dcpcb.builders))
	mutators := make([]Mutator, len(dcpcb.builders))
	for i := range dcpcb.builders {
		func(i int, root context.Context) {
			builder := dcpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DivisionChoicePAMutation)
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
					_, err = mutators[i+1].Mutate(root, dcpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcpcb *DivisionChoicePACreateBulk) SaveX(ctx context.Context) []*Division_Choice_PA {
	v, err := dcpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcpcb *DivisionChoicePACreateBulk) Exec(ctx context.Context) error {
	_, err := dcpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpcb *DivisionChoicePACreateBulk) ExecX(ctx context.Context) {
	if err := dcpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
