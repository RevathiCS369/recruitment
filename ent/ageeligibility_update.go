// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/ageeligibility"
	"recruit/ent/exameligibility"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgeEligibilityUpdate is the builder for updating AgeEligibility entities.
type AgeEligibilityUpdate struct {
	config
	hooks    []Hook
	mutation *AgeEligibilityMutation
}

// Where appends a list predicates to the AgeEligibilityUpdate builder.
func (aeu *AgeEligibilityUpdate) Where(ps ...predicate.AgeEligibility) *AgeEligibilityUpdate {
	aeu.mutation.Where(ps...)
	return aeu
}

// SetEligibilityCode sets the "EligibilityCode" field.
func (aeu *AgeEligibilityUpdate) SetEligibilityCode(i int32) *AgeEligibilityUpdate {
	aeu.mutation.SetEligibilityCode(i)
	return aeu
}

// SetNillableEligibilityCode sets the "EligibilityCode" field if the given value is not nil.
func (aeu *AgeEligibilityUpdate) SetNillableEligibilityCode(i *int32) *AgeEligibilityUpdate {
	if i != nil {
		aeu.SetEligibilityCode(*i)
	}
	return aeu
}

// ClearEligibilityCode clears the value of the "EligibilityCode" field.
func (aeu *AgeEligibilityUpdate) ClearEligibilityCode() *AgeEligibilityUpdate {
	aeu.mutation.ClearEligibilityCode()
	return aeu
}

// SetAge sets the "Age" field.
func (aeu *AgeEligibilityUpdate) SetAge(i int32) *AgeEligibilityUpdate {
	aeu.mutation.ResetAge()
	aeu.mutation.SetAge(i)
	return aeu
}

// SetNillableAge sets the "Age" field if the given value is not nil.
func (aeu *AgeEligibilityUpdate) SetNillableAge(i *int32) *AgeEligibilityUpdate {
	if i != nil {
		aeu.SetAge(*i)
	}
	return aeu
}

// AddAge adds i to the "Age" field.
func (aeu *AgeEligibilityUpdate) AddAge(i int32) *AgeEligibilityUpdate {
	aeu.mutation.AddAge(i)
	return aeu
}

// ClearAge clears the value of the "Age" field.
func (aeu *AgeEligibilityUpdate) ClearAge() *AgeEligibilityUpdate {
	aeu.mutation.ClearAge()
	return aeu
}

// SetCategoryID sets the "CategoryID" field.
func (aeu *AgeEligibilityUpdate) SetCategoryID(i int32) *AgeEligibilityUpdate {
	aeu.mutation.ResetCategoryID()
	aeu.mutation.SetCategoryID(i)
	return aeu
}

// SetNillableCategoryID sets the "CategoryID" field if the given value is not nil.
func (aeu *AgeEligibilityUpdate) SetNillableCategoryID(i *int32) *AgeEligibilityUpdate {
	if i != nil {
		aeu.SetCategoryID(*i)
	}
	return aeu
}

// AddCategoryID adds i to the "CategoryID" field.
func (aeu *AgeEligibilityUpdate) AddCategoryID(i int32) *AgeEligibilityUpdate {
	aeu.mutation.AddCategoryID(i)
	return aeu
}

// ClearCategoryID clears the value of the "CategoryID" field.
func (aeu *AgeEligibilityUpdate) ClearCategoryID() *AgeEligibilityUpdate {
	aeu.mutation.ClearCategoryID()
	return aeu
}

// SetExamEligibilityID sets the "exam_eligibility" edge to the ExamEligibility entity by ID.
func (aeu *AgeEligibilityUpdate) SetExamEligibilityID(id int32) *AgeEligibilityUpdate {
	aeu.mutation.SetExamEligibilityID(id)
	return aeu
}

// SetNillableExamEligibilityID sets the "exam_eligibility" edge to the ExamEligibility entity by ID if the given value is not nil.
func (aeu *AgeEligibilityUpdate) SetNillableExamEligibilityID(id *int32) *AgeEligibilityUpdate {
	if id != nil {
		aeu = aeu.SetExamEligibilityID(*id)
	}
	return aeu
}

// SetExamEligibility sets the "exam_eligibility" edge to the ExamEligibility entity.
func (aeu *AgeEligibilityUpdate) SetExamEligibility(e *ExamEligibility) *AgeEligibilityUpdate {
	return aeu.SetExamEligibilityID(e.ID)
}

// Mutation returns the AgeEligibilityMutation object of the builder.
func (aeu *AgeEligibilityUpdate) Mutation() *AgeEligibilityMutation {
	return aeu.mutation
}

// ClearExamEligibility clears the "exam_eligibility" edge to the ExamEligibility entity.
func (aeu *AgeEligibilityUpdate) ClearExamEligibility() *AgeEligibilityUpdate {
	aeu.mutation.ClearExamEligibility()
	return aeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aeu *AgeEligibilityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aeu.sqlSave, aeu.mutation, aeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeu *AgeEligibilityUpdate) SaveX(ctx context.Context) int {
	affected, err := aeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aeu *AgeEligibilityUpdate) Exec(ctx context.Context) error {
	_, err := aeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeu *AgeEligibilityUpdate) ExecX(ctx context.Context) {
	if err := aeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aeu *AgeEligibilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ageeligibility.Table, ageeligibility.Columns, sqlgraph.NewFieldSpec(ageeligibility.FieldID, field.TypeInt32))
	if ps := aeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeu.mutation.Age(); ok {
		_spec.SetField(ageeligibility.FieldAge, field.TypeInt32, value)
	}
	if value, ok := aeu.mutation.AddedAge(); ok {
		_spec.AddField(ageeligibility.FieldAge, field.TypeInt32, value)
	}
	if aeu.mutation.AgeCleared() {
		_spec.ClearField(ageeligibility.FieldAge, field.TypeInt32)
	}
	if value, ok := aeu.mutation.CategoryID(); ok {
		_spec.SetField(ageeligibility.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := aeu.mutation.AddedCategoryID(); ok {
		_spec.AddField(ageeligibility.FieldCategoryID, field.TypeInt32, value)
	}
	if aeu.mutation.CategoryIDCleared() {
		_spec.ClearField(ageeligibility.FieldCategoryID, field.TypeInt32)
	}
	if aeu.mutation.ExamEligibilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ageeligibility.ExamEligibilityTable,
			Columns: []string{ageeligibility.ExamEligibilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exameligibility.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeu.mutation.ExamEligibilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ageeligibility.ExamEligibilityTable,
			Columns: []string{ageeligibility.ExamEligibilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exameligibility.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ageeligibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aeu.mutation.done = true
	return n, nil
}

// AgeEligibilityUpdateOne is the builder for updating a single AgeEligibility entity.
type AgeEligibilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgeEligibilityMutation
}

// SetEligibilityCode sets the "EligibilityCode" field.
func (aeuo *AgeEligibilityUpdateOne) SetEligibilityCode(i int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.SetEligibilityCode(i)
	return aeuo
}

// SetNillableEligibilityCode sets the "EligibilityCode" field if the given value is not nil.
func (aeuo *AgeEligibilityUpdateOne) SetNillableEligibilityCode(i *int32) *AgeEligibilityUpdateOne {
	if i != nil {
		aeuo.SetEligibilityCode(*i)
	}
	return aeuo
}

// ClearEligibilityCode clears the value of the "EligibilityCode" field.
func (aeuo *AgeEligibilityUpdateOne) ClearEligibilityCode() *AgeEligibilityUpdateOne {
	aeuo.mutation.ClearEligibilityCode()
	return aeuo
}

// SetAge sets the "Age" field.
func (aeuo *AgeEligibilityUpdateOne) SetAge(i int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.ResetAge()
	aeuo.mutation.SetAge(i)
	return aeuo
}

// SetNillableAge sets the "Age" field if the given value is not nil.
func (aeuo *AgeEligibilityUpdateOne) SetNillableAge(i *int32) *AgeEligibilityUpdateOne {
	if i != nil {
		aeuo.SetAge(*i)
	}
	return aeuo
}

// AddAge adds i to the "Age" field.
func (aeuo *AgeEligibilityUpdateOne) AddAge(i int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.AddAge(i)
	return aeuo
}

// ClearAge clears the value of the "Age" field.
func (aeuo *AgeEligibilityUpdateOne) ClearAge() *AgeEligibilityUpdateOne {
	aeuo.mutation.ClearAge()
	return aeuo
}

// SetCategoryID sets the "CategoryID" field.
func (aeuo *AgeEligibilityUpdateOne) SetCategoryID(i int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.ResetCategoryID()
	aeuo.mutation.SetCategoryID(i)
	return aeuo
}

// SetNillableCategoryID sets the "CategoryID" field if the given value is not nil.
func (aeuo *AgeEligibilityUpdateOne) SetNillableCategoryID(i *int32) *AgeEligibilityUpdateOne {
	if i != nil {
		aeuo.SetCategoryID(*i)
	}
	return aeuo
}

// AddCategoryID adds i to the "CategoryID" field.
func (aeuo *AgeEligibilityUpdateOne) AddCategoryID(i int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.AddCategoryID(i)
	return aeuo
}

// ClearCategoryID clears the value of the "CategoryID" field.
func (aeuo *AgeEligibilityUpdateOne) ClearCategoryID() *AgeEligibilityUpdateOne {
	aeuo.mutation.ClearCategoryID()
	return aeuo
}

// SetExamEligibilityID sets the "exam_eligibility" edge to the ExamEligibility entity by ID.
func (aeuo *AgeEligibilityUpdateOne) SetExamEligibilityID(id int32) *AgeEligibilityUpdateOne {
	aeuo.mutation.SetExamEligibilityID(id)
	return aeuo
}

// SetNillableExamEligibilityID sets the "exam_eligibility" edge to the ExamEligibility entity by ID if the given value is not nil.
func (aeuo *AgeEligibilityUpdateOne) SetNillableExamEligibilityID(id *int32) *AgeEligibilityUpdateOne {
	if id != nil {
		aeuo = aeuo.SetExamEligibilityID(*id)
	}
	return aeuo
}

// SetExamEligibility sets the "exam_eligibility" edge to the ExamEligibility entity.
func (aeuo *AgeEligibilityUpdateOne) SetExamEligibility(e *ExamEligibility) *AgeEligibilityUpdateOne {
	return aeuo.SetExamEligibilityID(e.ID)
}

// Mutation returns the AgeEligibilityMutation object of the builder.
func (aeuo *AgeEligibilityUpdateOne) Mutation() *AgeEligibilityMutation {
	return aeuo.mutation
}

// ClearExamEligibility clears the "exam_eligibility" edge to the ExamEligibility entity.
func (aeuo *AgeEligibilityUpdateOne) ClearExamEligibility() *AgeEligibilityUpdateOne {
	aeuo.mutation.ClearExamEligibility()
	return aeuo
}

// Where appends a list predicates to the AgeEligibilityUpdate builder.
func (aeuo *AgeEligibilityUpdateOne) Where(ps ...predicate.AgeEligibility) *AgeEligibilityUpdateOne {
	aeuo.mutation.Where(ps...)
	return aeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aeuo *AgeEligibilityUpdateOne) Select(field string, fields ...string) *AgeEligibilityUpdateOne {
	aeuo.fields = append([]string{field}, fields...)
	return aeuo
}

// Save executes the query and returns the updated AgeEligibility entity.
func (aeuo *AgeEligibilityUpdateOne) Save(ctx context.Context) (*AgeEligibility, error) {
	return withHooks(ctx, aeuo.sqlSave, aeuo.mutation, aeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeuo *AgeEligibilityUpdateOne) SaveX(ctx context.Context) *AgeEligibility {
	node, err := aeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aeuo *AgeEligibilityUpdateOne) Exec(ctx context.Context) error {
	_, err := aeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeuo *AgeEligibilityUpdateOne) ExecX(ctx context.Context) {
	if err := aeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aeuo *AgeEligibilityUpdateOne) sqlSave(ctx context.Context) (_node *AgeEligibility, err error) {
	_spec := sqlgraph.NewUpdateSpec(ageeligibility.Table, ageeligibility.Columns, sqlgraph.NewFieldSpec(ageeligibility.FieldID, field.TypeInt32))
	id, ok := aeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AgeEligibility.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ageeligibility.FieldID)
		for _, f := range fields {
			if !ageeligibility.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ageeligibility.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeuo.mutation.Age(); ok {
		_spec.SetField(ageeligibility.FieldAge, field.TypeInt32, value)
	}
	if value, ok := aeuo.mutation.AddedAge(); ok {
		_spec.AddField(ageeligibility.FieldAge, field.TypeInt32, value)
	}
	if aeuo.mutation.AgeCleared() {
		_spec.ClearField(ageeligibility.FieldAge, field.TypeInt32)
	}
	if value, ok := aeuo.mutation.CategoryID(); ok {
		_spec.SetField(ageeligibility.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := aeuo.mutation.AddedCategoryID(); ok {
		_spec.AddField(ageeligibility.FieldCategoryID, field.TypeInt32, value)
	}
	if aeuo.mutation.CategoryIDCleared() {
		_spec.ClearField(ageeligibility.FieldCategoryID, field.TypeInt32)
	}
	if aeuo.mutation.ExamEligibilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ageeligibility.ExamEligibilityTable,
			Columns: []string{ageeligibility.ExamEligibilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exameligibility.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeuo.mutation.ExamEligibilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ageeligibility.ExamEligibilityTable,
			Columns: []string{ageeligibility.ExamEligibilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exameligibility.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AgeEligibility{config: aeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ageeligibility.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aeuo.mutation.done = true
	return _node, nil
}
