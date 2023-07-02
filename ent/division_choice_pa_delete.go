// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/division_choice_pa"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionChoicePADelete is the builder for deleting a Division_Choice_PA entity.
type DivisionChoicePADelete struct {
	config
	hooks    []Hook
	mutation *DivisionChoicePAMutation
}

// Where appends a list predicates to the DivisionChoicePADelete builder.
func (dcpd *DivisionChoicePADelete) Where(ps ...predicate.Division_Choice_PA) *DivisionChoicePADelete {
	dcpd.mutation.Where(ps...)
	return dcpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcpd *DivisionChoicePADelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dcpd.sqlExec, dcpd.mutation, dcpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpd *DivisionChoicePADelete) ExecX(ctx context.Context) int {
	n, err := dcpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcpd *DivisionChoicePADelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(division_choice_pa.Table, sqlgraph.NewFieldSpec(division_choice_pa.FieldID, field.TypeInt32))
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

// DivisionChoicePADeleteOne is the builder for deleting a single Division_Choice_PA entity.
type DivisionChoicePADeleteOne struct {
	dcpd *DivisionChoicePADelete
}

// Where appends a list predicates to the DivisionChoicePADelete builder.
func (dcpdo *DivisionChoicePADeleteOne) Where(ps ...predicate.Division_Choice_PA) *DivisionChoicePADeleteOne {
	dcpdo.dcpd.mutation.Where(ps...)
	return dcpdo
}

// Exec executes the deletion query.
func (dcpdo *DivisionChoicePADeleteOne) Exec(ctx context.Context) error {
	n, err := dcpdo.dcpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{division_choice_pa.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcpdo *DivisionChoicePADeleteOne) ExecX(ctx context.Context) {
	if err := dcpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
