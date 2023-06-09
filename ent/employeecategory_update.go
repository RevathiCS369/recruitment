// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/employeecategory"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeCategoryUpdate is the builder for updating EmployeeCategory entities.
type EmployeeCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *EmployeeCategoryMutation
}

// Where appends a list predicates to the EmployeeCategoryUpdate builder.
func (ecu *EmployeeCategoryUpdate) Where(ps ...predicate.EmployeeCategory) *EmployeeCategoryUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetCategrycode sets the "Categrycode" field.
func (ecu *EmployeeCategoryUpdate) SetCategrycode(s string) *EmployeeCategoryUpdate {
	ecu.mutation.SetCategrycode(s)
	return ecu
}

// SetCategoryDescription sets the "CategoryDescription" field.
func (ecu *EmployeeCategoryUpdate) SetCategoryDescription(s string) *EmployeeCategoryUpdate {
	ecu.mutation.SetCategoryDescription(s)
	return ecu
}

// SetMinimumMarks sets the "MinimumMarks" field.
func (ecu *EmployeeCategoryUpdate) SetMinimumMarks(i int32) *EmployeeCategoryUpdate {
	ecu.mutation.ResetMinimumMarks()
	ecu.mutation.SetMinimumMarks(i)
	return ecu
}

// SetNillableMinimumMarks sets the "MinimumMarks" field if the given value is not nil.
func (ecu *EmployeeCategoryUpdate) SetNillableMinimumMarks(i *int32) *EmployeeCategoryUpdate {
	if i != nil {
		ecu.SetMinimumMarks(*i)
	}
	return ecu
}

// AddMinimumMarks adds i to the "MinimumMarks" field.
func (ecu *EmployeeCategoryUpdate) AddMinimumMarks(i int32) *EmployeeCategoryUpdate {
	ecu.mutation.AddMinimumMarks(i)
	return ecu
}

// ClearMinimumMarks clears the value of the "MinimumMarks" field.
func (ecu *EmployeeCategoryUpdate) ClearMinimumMarks() *EmployeeCategoryUpdate {
	ecu.mutation.ClearMinimumMarks()
	return ecu
}

// Mutation returns the EmployeeCategoryMutation object of the builder.
func (ecu *EmployeeCategoryUpdate) Mutation() *EmployeeCategoryMutation {
	return ecu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EmployeeCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EmployeeCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EmployeeCategoryUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EmployeeCategoryUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecu *EmployeeCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(employeecategory.Table, employeecategory.Columns, sqlgraph.NewFieldSpec(employeecategory.FieldID, field.TypeInt32))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Categrycode(); ok {
		_spec.SetField(employeecategory.FieldCategrycode, field.TypeString, value)
	}
	if value, ok := ecu.mutation.CategoryDescription(); ok {
		_spec.SetField(employeecategory.FieldCategoryDescription, field.TypeString, value)
	}
	if value, ok := ecu.mutation.MinimumMarks(); ok {
		_spec.SetField(employeecategory.FieldMinimumMarks, field.TypeInt32, value)
	}
	if value, ok := ecu.mutation.AddedMinimumMarks(); ok {
		_spec.AddField(employeecategory.FieldMinimumMarks, field.TypeInt32, value)
	}
	if ecu.mutation.MinimumMarksCleared() {
		_spec.ClearField(employeecategory.FieldMinimumMarks, field.TypeInt32)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EmployeeCategoryUpdateOne is the builder for updating a single EmployeeCategory entity.
type EmployeeCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmployeeCategoryMutation
}

// SetCategrycode sets the "Categrycode" field.
func (ecuo *EmployeeCategoryUpdateOne) SetCategrycode(s string) *EmployeeCategoryUpdateOne {
	ecuo.mutation.SetCategrycode(s)
	return ecuo
}

// SetCategoryDescription sets the "CategoryDescription" field.
func (ecuo *EmployeeCategoryUpdateOne) SetCategoryDescription(s string) *EmployeeCategoryUpdateOne {
	ecuo.mutation.SetCategoryDescription(s)
	return ecuo
}

// SetMinimumMarks sets the "MinimumMarks" field.
func (ecuo *EmployeeCategoryUpdateOne) SetMinimumMarks(i int32) *EmployeeCategoryUpdateOne {
	ecuo.mutation.ResetMinimumMarks()
	ecuo.mutation.SetMinimumMarks(i)
	return ecuo
}

// SetNillableMinimumMarks sets the "MinimumMarks" field if the given value is not nil.
func (ecuo *EmployeeCategoryUpdateOne) SetNillableMinimumMarks(i *int32) *EmployeeCategoryUpdateOne {
	if i != nil {
		ecuo.SetMinimumMarks(*i)
	}
	return ecuo
}

// AddMinimumMarks adds i to the "MinimumMarks" field.
func (ecuo *EmployeeCategoryUpdateOne) AddMinimumMarks(i int32) *EmployeeCategoryUpdateOne {
	ecuo.mutation.AddMinimumMarks(i)
	return ecuo
}

// ClearMinimumMarks clears the value of the "MinimumMarks" field.
func (ecuo *EmployeeCategoryUpdateOne) ClearMinimumMarks() *EmployeeCategoryUpdateOne {
	ecuo.mutation.ClearMinimumMarks()
	return ecuo
}

// Mutation returns the EmployeeCategoryMutation object of the builder.
func (ecuo *EmployeeCategoryUpdateOne) Mutation() *EmployeeCategoryMutation {
	return ecuo.mutation
}

// Where appends a list predicates to the EmployeeCategoryUpdate builder.
func (ecuo *EmployeeCategoryUpdateOne) Where(ps ...predicate.EmployeeCategory) *EmployeeCategoryUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EmployeeCategoryUpdateOne) Select(field string, fields ...string) *EmployeeCategoryUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EmployeeCategory entity.
func (ecuo *EmployeeCategoryUpdateOne) Save(ctx context.Context) (*EmployeeCategory, error) {
	return withHooks(ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EmployeeCategoryUpdateOne) SaveX(ctx context.Context) *EmployeeCategory {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EmployeeCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EmployeeCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecuo *EmployeeCategoryUpdateOne) sqlSave(ctx context.Context) (_node *EmployeeCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(employeecategory.Table, employeecategory.Columns, sqlgraph.NewFieldSpec(employeecategory.FieldID, field.TypeInt32))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmployeeCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employeecategory.FieldID)
		for _, f := range fields {
			if !employeecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != employeecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Categrycode(); ok {
		_spec.SetField(employeecategory.FieldCategrycode, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.CategoryDescription(); ok {
		_spec.SetField(employeecategory.FieldCategoryDescription, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.MinimumMarks(); ok {
		_spec.SetField(employeecategory.FieldMinimumMarks, field.TypeInt32, value)
	}
	if value, ok := ecuo.mutation.AddedMinimumMarks(); ok {
		_spec.AddField(employeecategory.FieldMinimumMarks, field.TypeInt32, value)
	}
	if ecuo.mutation.MinimumMarksCleared() {
		_spec.ClearField(employeecategory.FieldMinimumMarks, field.TypeInt32)
	}
	_node = &EmployeeCategory{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employeecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
