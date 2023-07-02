// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/exameligibility"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamEligibilityDelete is the builder for deleting a ExamEligibility entity.
type ExamEligibilityDelete struct {
	config
	hooks    []Hook
	mutation *ExamEligibilityMutation
}

// Where appends a list predicates to the ExamEligibilityDelete builder.
func (eed *ExamEligibilityDelete) Where(ps ...predicate.ExamEligibility) *ExamEligibilityDelete {
	eed.mutation.Where(ps...)
	return eed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eed *ExamEligibilityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eed.sqlExec, eed.mutation, eed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eed *ExamEligibilityDelete) ExecX(ctx context.Context) int {
	n, err := eed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eed *ExamEligibilityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(exameligibility.Table, sqlgraph.NewFieldSpec(exameligibility.FieldID, field.TypeInt32))
	if ps := eed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eed.mutation.done = true
	return affected, err
}

// ExamEligibilityDeleteOne is the builder for deleting a single ExamEligibility entity.
type ExamEligibilityDeleteOne struct {
	eed *ExamEligibilityDelete
}

// Where appends a list predicates to the ExamEligibilityDelete builder.
func (eedo *ExamEligibilityDeleteOne) Where(ps ...predicate.ExamEligibility) *ExamEligibilityDeleteOne {
	eedo.eed.mutation.Where(ps...)
	return eedo
}

// Exec executes the deletion query.
func (eedo *ExamEligibilityDeleteOne) Exec(ctx context.Context) error {
	n, err := eedo.eed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exameligibility.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eedo *ExamEligibilityDeleteOne) ExecX(ctx context.Context) {
	if err := eedo.Exec(ctx); err != nil {
		panic(err)
	}
}
