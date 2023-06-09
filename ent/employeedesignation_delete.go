// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/employeedesignation"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmployeeDesignationDelete is the builder for deleting a EmployeeDesignation entity.
type EmployeeDesignationDelete struct {
	config
	hooks    []Hook
	mutation *EmployeeDesignationMutation
}

// Where appends a list predicates to the EmployeeDesignationDelete builder.
func (edd *EmployeeDesignationDelete) Where(ps ...predicate.EmployeeDesignation) *EmployeeDesignationDelete {
	edd.mutation.Where(ps...)
	return edd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (edd *EmployeeDesignationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, edd.sqlExec, edd.mutation, edd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (edd *EmployeeDesignationDelete) ExecX(ctx context.Context) int {
	n, err := edd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (edd *EmployeeDesignationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(employeedesignation.Table, sqlgraph.NewFieldSpec(employeedesignation.FieldID, field.TypeInt32))
	if ps := edd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, edd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	edd.mutation.done = true
	return affected, err
}

// EmployeeDesignationDeleteOne is the builder for deleting a single EmployeeDesignation entity.
type EmployeeDesignationDeleteOne struct {
	edd *EmployeeDesignationDelete
}

// Where appends a list predicates to the EmployeeDesignationDelete builder.
func (eddo *EmployeeDesignationDeleteOne) Where(ps ...predicate.EmployeeDesignation) *EmployeeDesignationDeleteOne {
	eddo.edd.mutation.Where(ps...)
	return eddo
}

// Exec executes the deletion query.
func (eddo *EmployeeDesignationDeleteOne) Exec(ctx context.Context) error {
	n, err := eddo.edd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{employeedesignation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eddo *EmployeeDesignationDeleteOne) ExecX(ctx context.Context) {
	if err := eddo.Exec(ctx); err != nil {
		panic(err)
	}
}
