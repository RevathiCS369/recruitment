// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exampapers"
	"recruit/ent/papertypes"
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaperTypesUpdate is the builder for updating PaperTypes entities.
type PaperTypesUpdate struct {
	config
	hooks    []Hook
	mutation *PaperTypesMutation
}

// Where appends a list predicates to the PaperTypesUpdate builder.
func (ptu *PaperTypesUpdate) Where(ps ...predicate.PaperTypes) *PaperTypesUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetPaperCode sets the "PaperCode" field.
func (ptu *PaperTypesUpdate) SetPaperCode(i int32) *PaperTypesUpdate {
	ptu.mutation.SetPaperCode(i)
	return ptu
}

// SetNillablePaperCode sets the "PaperCode" field if the given value is not nil.
func (ptu *PaperTypesUpdate) SetNillablePaperCode(i *int32) *PaperTypesUpdate {
	if i != nil {
		ptu.SetPaperCode(*i)
	}
	return ptu
}

// ClearPaperCode clears the value of the "PaperCode" field.
func (ptu *PaperTypesUpdate) ClearPaperCode() *PaperTypesUpdate {
	ptu.mutation.ClearPaperCode()
	return ptu
}

// SetPaperTypeDescription sets the "PaperTypeDescription" field.
func (ptu *PaperTypesUpdate) SetPaperTypeDescription(s string) *PaperTypesUpdate {
	ptu.mutation.SetPaperTypeDescription(s)
	return ptu
}

// SetOrderNumber sets the "OrderNumber" field.
func (ptu *PaperTypesUpdate) SetOrderNumber(s string) *PaperTypesUpdate {
	ptu.mutation.SetOrderNumber(s)
	return ptu
}

// SetSequenceNumber sets the "SequenceNumber" field.
func (ptu *PaperTypesUpdate) SetSequenceNumber(i int32) *PaperTypesUpdate {
	ptu.mutation.ResetSequenceNumber()
	ptu.mutation.SetSequenceNumber(i)
	return ptu
}

// SetNillableSequenceNumber sets the "SequenceNumber" field if the given value is not nil.
func (ptu *PaperTypesUpdate) SetNillableSequenceNumber(i *int32) *PaperTypesUpdate {
	if i != nil {
		ptu.SetSequenceNumber(*i)
	}
	return ptu
}

// AddSequenceNumber adds i to the "SequenceNumber" field.
func (ptu *PaperTypesUpdate) AddSequenceNumber(i int32) *PaperTypesUpdate {
	ptu.mutation.AddSequenceNumber(i)
	return ptu
}

// ClearSequenceNumber clears the value of the "SequenceNumber" field.
func (ptu *PaperTypesUpdate) ClearSequenceNumber() *PaperTypesUpdate {
	ptu.mutation.ClearSequenceNumber()
	return ptu
}

// SetCreatedDate sets the "CreatedDate" field.
func (ptu *PaperTypesUpdate) SetCreatedDate(t time.Time) *PaperTypesUpdate {
	ptu.mutation.SetCreatedDate(t)
	return ptu
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (ptu *PaperTypesUpdate) SetNillableCreatedDate(t *time.Time) *PaperTypesUpdate {
	if t != nil {
		ptu.SetCreatedDate(*t)
	}
	return ptu
}

// ClearCreatedDate clears the value of the "CreatedDate" field.
func (ptu *PaperTypesUpdate) ClearCreatedDate() *PaperTypesUpdate {
	ptu.mutation.ClearCreatedDate()
	return ptu
}

// SetPapercodeID sets the "papercode" edge to the ExamPapers entity by ID.
func (ptu *PaperTypesUpdate) SetPapercodeID(id int32) *PaperTypesUpdate {
	ptu.mutation.SetPapercodeID(id)
	return ptu
}

// SetNillablePapercodeID sets the "papercode" edge to the ExamPapers entity by ID if the given value is not nil.
func (ptu *PaperTypesUpdate) SetNillablePapercodeID(id *int32) *PaperTypesUpdate {
	if id != nil {
		ptu = ptu.SetPapercodeID(*id)
	}
	return ptu
}

// SetPapercode sets the "papercode" edge to the ExamPapers entity.
func (ptu *PaperTypesUpdate) SetPapercode(e *ExamPapers) *PaperTypesUpdate {
	return ptu.SetPapercodeID(e.ID)
}

// Mutation returns the PaperTypesMutation object of the builder.
func (ptu *PaperTypesUpdate) Mutation() *PaperTypesMutation {
	return ptu.mutation
}

// ClearPapercode clears the "papercode" edge to the ExamPapers entity.
func (ptu *PaperTypesUpdate) ClearPapercode() *PaperTypesUpdate {
	ptu.mutation.ClearPapercode()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PaperTypesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ptu.sqlSave, ptu.mutation, ptu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PaperTypesUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PaperTypesUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PaperTypesUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *PaperTypesUpdate) check() error {
	if v, ok := ptu.mutation.PaperTypeDescription(); ok {
		if err := papertypes.PaperTypeDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "PaperTypeDescription", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.PaperTypeDescription": %w`, err)}
		}
	}
	if v, ok := ptu.mutation.OrderNumber(); ok {
		if err := papertypes.OrderNumberValidator(v); err != nil {
			return &ValidationError{Name: "OrderNumber", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.OrderNumber": %w`, err)}
		}
	}
	return nil
}

func (ptu *PaperTypesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ptu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(papertypes.Table, papertypes.Columns, sqlgraph.NewFieldSpec(papertypes.FieldID, field.TypeInt32))
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.PaperTypeDescription(); ok {
		_spec.SetField(papertypes.FieldPaperTypeDescription, field.TypeString, value)
	}
	if value, ok := ptu.mutation.OrderNumber(); ok {
		_spec.SetField(papertypes.FieldOrderNumber, field.TypeString, value)
	}
	if value, ok := ptu.mutation.SequenceNumber(); ok {
		_spec.SetField(papertypes.FieldSequenceNumber, field.TypeInt32, value)
	}
	if value, ok := ptu.mutation.AddedSequenceNumber(); ok {
		_spec.AddField(papertypes.FieldSequenceNumber, field.TypeInt32, value)
	}
	if ptu.mutation.SequenceNumberCleared() {
		_spec.ClearField(papertypes.FieldSequenceNumber, field.TypeInt32)
	}
	if value, ok := ptu.mutation.CreatedDate(); ok {
		_spec.SetField(papertypes.FieldCreatedDate, field.TypeTime, value)
	}
	if ptu.mutation.CreatedDateCleared() {
		_spec.ClearField(papertypes.FieldCreatedDate, field.TypeTime)
	}
	if ptu.mutation.PapercodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.PapercodeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{papertypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptu.mutation.done = true
	return n, nil
}

// PaperTypesUpdateOne is the builder for updating a single PaperTypes entity.
type PaperTypesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaperTypesMutation
}

// SetPaperCode sets the "PaperCode" field.
func (ptuo *PaperTypesUpdateOne) SetPaperCode(i int32) *PaperTypesUpdateOne {
	ptuo.mutation.SetPaperCode(i)
	return ptuo
}

// SetNillablePaperCode sets the "PaperCode" field if the given value is not nil.
func (ptuo *PaperTypesUpdateOne) SetNillablePaperCode(i *int32) *PaperTypesUpdateOne {
	if i != nil {
		ptuo.SetPaperCode(*i)
	}
	return ptuo
}

// ClearPaperCode clears the value of the "PaperCode" field.
func (ptuo *PaperTypesUpdateOne) ClearPaperCode() *PaperTypesUpdateOne {
	ptuo.mutation.ClearPaperCode()
	return ptuo
}

// SetPaperTypeDescription sets the "PaperTypeDescription" field.
func (ptuo *PaperTypesUpdateOne) SetPaperTypeDescription(s string) *PaperTypesUpdateOne {
	ptuo.mutation.SetPaperTypeDescription(s)
	return ptuo
}

// SetOrderNumber sets the "OrderNumber" field.
func (ptuo *PaperTypesUpdateOne) SetOrderNumber(s string) *PaperTypesUpdateOne {
	ptuo.mutation.SetOrderNumber(s)
	return ptuo
}

// SetSequenceNumber sets the "SequenceNumber" field.
func (ptuo *PaperTypesUpdateOne) SetSequenceNumber(i int32) *PaperTypesUpdateOne {
	ptuo.mutation.ResetSequenceNumber()
	ptuo.mutation.SetSequenceNumber(i)
	return ptuo
}

// SetNillableSequenceNumber sets the "SequenceNumber" field if the given value is not nil.
func (ptuo *PaperTypesUpdateOne) SetNillableSequenceNumber(i *int32) *PaperTypesUpdateOne {
	if i != nil {
		ptuo.SetSequenceNumber(*i)
	}
	return ptuo
}

// AddSequenceNumber adds i to the "SequenceNumber" field.
func (ptuo *PaperTypesUpdateOne) AddSequenceNumber(i int32) *PaperTypesUpdateOne {
	ptuo.mutation.AddSequenceNumber(i)
	return ptuo
}

// ClearSequenceNumber clears the value of the "SequenceNumber" field.
func (ptuo *PaperTypesUpdateOne) ClearSequenceNumber() *PaperTypesUpdateOne {
	ptuo.mutation.ClearSequenceNumber()
	return ptuo
}

// SetCreatedDate sets the "CreatedDate" field.
func (ptuo *PaperTypesUpdateOne) SetCreatedDate(t time.Time) *PaperTypesUpdateOne {
	ptuo.mutation.SetCreatedDate(t)
	return ptuo
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (ptuo *PaperTypesUpdateOne) SetNillableCreatedDate(t *time.Time) *PaperTypesUpdateOne {
	if t != nil {
		ptuo.SetCreatedDate(*t)
	}
	return ptuo
}

// ClearCreatedDate clears the value of the "CreatedDate" field.
func (ptuo *PaperTypesUpdateOne) ClearCreatedDate() *PaperTypesUpdateOne {
	ptuo.mutation.ClearCreatedDate()
	return ptuo
}

// SetPapercodeID sets the "papercode" edge to the ExamPapers entity by ID.
func (ptuo *PaperTypesUpdateOne) SetPapercodeID(id int32) *PaperTypesUpdateOne {
	ptuo.mutation.SetPapercodeID(id)
	return ptuo
}

// SetNillablePapercodeID sets the "papercode" edge to the ExamPapers entity by ID if the given value is not nil.
func (ptuo *PaperTypesUpdateOne) SetNillablePapercodeID(id *int32) *PaperTypesUpdateOne {
	if id != nil {
		ptuo = ptuo.SetPapercodeID(*id)
	}
	return ptuo
}

// SetPapercode sets the "papercode" edge to the ExamPapers entity.
func (ptuo *PaperTypesUpdateOne) SetPapercode(e *ExamPapers) *PaperTypesUpdateOne {
	return ptuo.SetPapercodeID(e.ID)
}

// Mutation returns the PaperTypesMutation object of the builder.
func (ptuo *PaperTypesUpdateOne) Mutation() *PaperTypesMutation {
	return ptuo.mutation
}

// ClearPapercode clears the "papercode" edge to the ExamPapers entity.
func (ptuo *PaperTypesUpdateOne) ClearPapercode() *PaperTypesUpdateOne {
	ptuo.mutation.ClearPapercode()
	return ptuo
}

// Where appends a list predicates to the PaperTypesUpdate builder.
func (ptuo *PaperTypesUpdateOne) Where(ps ...predicate.PaperTypes) *PaperTypesUpdateOne {
	ptuo.mutation.Where(ps...)
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PaperTypesUpdateOne) Select(field string, fields ...string) *PaperTypesUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PaperTypes entity.
func (ptuo *PaperTypesUpdateOne) Save(ctx context.Context) (*PaperTypes, error) {
	return withHooks(ctx, ptuo.sqlSave, ptuo.mutation, ptuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PaperTypesUpdateOne) SaveX(ctx context.Context) *PaperTypes {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PaperTypesUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PaperTypesUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *PaperTypesUpdateOne) check() error {
	if v, ok := ptuo.mutation.PaperTypeDescription(); ok {
		if err := papertypes.PaperTypeDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "PaperTypeDescription", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.PaperTypeDescription": %w`, err)}
		}
	}
	if v, ok := ptuo.mutation.OrderNumber(); ok {
		if err := papertypes.OrderNumberValidator(v); err != nil {
			return &ValidationError{Name: "OrderNumber", err: fmt.Errorf(`ent: validator failed for field "PaperTypes.OrderNumber": %w`, err)}
		}
	}
	return nil
}

func (ptuo *PaperTypesUpdateOne) sqlSave(ctx context.Context) (_node *PaperTypes, err error) {
	if err := ptuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(papertypes.Table, papertypes.Columns, sqlgraph.NewFieldSpec(papertypes.FieldID, field.TypeInt32))
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaperTypes.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, papertypes.FieldID)
		for _, f := range fields {
			if !papertypes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != papertypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.PaperTypeDescription(); ok {
		_spec.SetField(papertypes.FieldPaperTypeDescription, field.TypeString, value)
	}
	if value, ok := ptuo.mutation.OrderNumber(); ok {
		_spec.SetField(papertypes.FieldOrderNumber, field.TypeString, value)
	}
	if value, ok := ptuo.mutation.SequenceNumber(); ok {
		_spec.SetField(papertypes.FieldSequenceNumber, field.TypeInt32, value)
	}
	if value, ok := ptuo.mutation.AddedSequenceNumber(); ok {
		_spec.AddField(papertypes.FieldSequenceNumber, field.TypeInt32, value)
	}
	if ptuo.mutation.SequenceNumberCleared() {
		_spec.ClearField(papertypes.FieldSequenceNumber, field.TypeInt32)
	}
	if value, ok := ptuo.mutation.CreatedDate(); ok {
		_spec.SetField(papertypes.FieldCreatedDate, field.TypeTime, value)
	}
	if ptuo.mutation.CreatedDateCleared() {
		_spec.ClearField(papertypes.FieldCreatedDate, field.TypeTime)
	}
	if ptuo.mutation.PapercodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.PapercodeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PaperTypes{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{papertypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptuo.mutation.done = true
	return _node, nil
}
