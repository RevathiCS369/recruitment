// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exampapers"
	"recruit/ent/papertypes"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaperTypesCreate is the builder for creating a PaperTypes entity.
type PaperTypesCreate struct {
	config
	mutation *PaperTypesMutation
	hooks    []Hook
}

// SetPaperCode sets the "PaperCode" field.
func (ptc *PaperTypesCreate) SetPaperCode(i int32) *PaperTypesCreate {
	ptc.mutation.SetPaperCode(i)
	return ptc
}

// SetNillablePaperCode sets the "PaperCode" field if the given value is not nil.
func (ptc *PaperTypesCreate) SetNillablePaperCode(i *int32) *PaperTypesCreate {
	if i != nil {
		ptc.SetPaperCode(*i)
	}
	return ptc
}

// SetPaperTypeDescription sets the "PaperTypeDescription" field.
func (ptc *PaperTypesCreate) SetPaperTypeDescription(s string) *PaperTypesCreate {
	ptc.mutation.SetPaperTypeDescription(s)
	return ptc
}

// SetOrderNumber sets the "OrderNumber" field.
func (ptc *PaperTypesCreate) SetOrderNumber(s string) *PaperTypesCreate {
	ptc.mutation.SetOrderNumber(s)
	return ptc
}

// SetSequenceNumber sets the "SequenceNumber" field.
func (ptc *PaperTypesCreate) SetSequenceNumber(i int32) *PaperTypesCreate {
	ptc.mutation.SetSequenceNumber(i)
	return ptc
}

// SetCreatedDate sets the "CreatedDate" field.
func (ptc *PaperTypesCreate) SetCreatedDate(t time.Time) *PaperTypesCreate {
	ptc.mutation.SetCreatedDate(t)
	return ptc
}

// SetID sets the "id" field.
func (ptc *PaperTypesCreate) SetID(i int32) *PaperTypesCreate {
	ptc.mutation.SetID(i)
	return ptc
}

// SetPapercodeID sets the "papercode" edge to the ExamPapers entity by ID.
func (ptc *PaperTypesCreate) SetPapercodeID(id int32) *PaperTypesCreate {
	ptc.mutation.SetPapercodeID(id)
	return ptc
}

// SetNillablePapercodeID sets the "papercode" edge to the ExamPapers entity by ID if the given value is not nil.
func (ptc *PaperTypesCreate) SetNillablePapercodeID(id *int32) *PaperTypesCreate {
	if id != nil {
		ptc = ptc.SetPapercodeID(*id)
	}
	return ptc
}

// SetPapercode sets the "papercode" edge to the ExamPapers entity.
func (ptc *PaperTypesCreate) SetPapercode(e *ExamPapers) *PaperTypesCreate {
	return ptc.SetPapercodeID(e.ID)
}

// Mutation returns the PaperTypesMutation object of the builder.
func (ptc *PaperTypesCreate) Mutation() *PaperTypesMutation {
	return ptc.mutation
}

// Save creates the PaperTypes in the database.
func (ptc *PaperTypesCreate) Save(ctx context.Context) (*PaperTypes, error) {
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PaperTypesCreate) SaveX(ctx context.Context) *PaperTypes {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *PaperTypesCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *PaperTypesCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *PaperTypesCreate) check() error {
	if _, ok := ptc.mutation.PaperTypeDescription(); !ok {
		return &ValidationError{Name: "PaperTypeDescription", err: errors.New(`ent: missing required field "PaperTypes.PaperTypeDescription"`)}
	}
	if v, ok := ptc.mutation.PaperTypeDescription(); ok {
		if err := papertypes.PaperTypeDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "PaperTypeDescription", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.PaperTypeDescription": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.OrderNumber(); !ok {
		return &ValidationError{Name: "OrderNumber", err: errors.New(`ent: missing required field "PaperTypes.OrderNumber"`)}
	}
	if v, ok := ptc.mutation.OrderNumber(); ok {
		if err := papertypes.OrderNumberValidator(v); err != nil {
			return &ValidationError{Name: "OrderNumber", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.OrderNumber": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.SequenceNumber(); !ok {
		return &ValidationError{Name: "SequenceNumber", err: errors.New(`ent: missing required field "PaperTypes.SequenceNumber"`)}
	}
	if _, ok := ptc.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "CreatedDate", err: errors.New(`ent: missing required field "PaperTypes.CreatedDate"`)}
	}
	return nil
}

func (ptc *PaperTypesCreate) sqlSave(ctx context.Context) (*PaperTypes, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ptc.mutation.id = &_node.ID
	ptc.mutation.done = true
	return _node, nil
}

func (ptc *PaperTypesCreate) createSpec() (*PaperTypes, *sqlgraph.CreateSpec) {
	var (
		_node = &PaperTypes{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(papertypes.Table, sqlgraph.NewFieldSpec(papertypes.FieldID, field.TypeInt32))
	)
	if id, ok := ptc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ptc.mutation.PaperTypeDescription(); ok {
		_spec.SetField(papertypes.FieldPaperTypeDescription, field.TypeString, value)
		_node.PaperTypeDescription = value
	}
	if value, ok := ptc.mutation.OrderNumber(); ok {
		_spec.SetField(papertypes.FieldOrderNumber, field.TypeString, value)
		_node.OrderNumber = value
	}
	if value, ok := ptc.mutation.SequenceNumber(); ok {
		_spec.SetField(papertypes.FieldSequenceNumber, field.TypeInt32, value)
		_node.SequenceNumber = value
	}
	if value, ok := ptc.mutation.CreatedDate(); ok {
		_spec.SetField(papertypes.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = value
	}
	if nodes := ptc.mutation.PapercodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   papertypes.PapercodeTable,
			Columns: []string{papertypes.PapercodeColumn},
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
	return _node, _spec
}

// PaperTypesCreateBulk is the builder for creating many PaperTypes entities in bulk.
type PaperTypesCreateBulk struct {
	config
	builders []*PaperTypesCreate
}

// Save creates the PaperTypes entities in the database.
func (ptcb *PaperTypesCreateBulk) Save(ctx context.Context) ([]*PaperTypes, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*PaperTypes, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaperTypesMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *PaperTypesCreateBulk) SaveX(ctx context.Context) []*PaperTypes {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *PaperTypesCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *PaperTypesCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}