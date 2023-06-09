// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/division_choice_pm"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionChoicePMDelete is the builder for deleting a Division_Choice_PM entity.
type DivisionChoicePMDelete struct {
	config
	hooks    []Hook
	mutation *DivisionChoicePMMutation
}

// Where appends a list predicates to the DivisionChoicePMDelete builder.
func (dcpd *DivisionChoicePMDelete) Where(ps ...predicate.Division_Choice_PM) *DivisionChoicePMDelete {
	dcpd.mutation.Where(ps...)
	return dcpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcpd *DivisionChoicePMDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dcpd.sqlExec, dcpd.mutation, dcpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpd *DivisionChoicePMDelete) ExecX(ctx context.Context) int {
	n, err := dcpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcpd *DivisionChoicePMDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(division_choice_pm.Table, sqlgraph.NewFieldSpec(division_choice_pm.FieldID, field.TypeInt32))
	if ps := dcpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dcpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dcpd.mutation.done = true
	return affected, err
}

// DivisionChoicePMDeleteOne is the builder for deleting a single Division_Choice_PM entity.
type DivisionChoicePMDeleteOne struct {
	dcpd *DivisionChoicePMDelete
}

// Where appends a list predicates to the DivisionChoicePMDelete builder.
func (dcpdo *DivisionChoicePMDeleteOne) Where(ps ...predicate.Division_Choice_PM) *DivisionChoicePMDeleteOne {
	dcpdo.dcpd.mutation.Where(ps...)
	return dcpdo
}

// Exec executes the deletion query.
func (dcpdo *DivisionChoicePMDeleteOne) Exec(ctx context.Context) error {
	n, err := dcpdo.dcpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{division_choice_pm.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpdo *DivisionChoicePMDeleteOne) ExecX(ctx context.Context) {
	if err := dcpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
