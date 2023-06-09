// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/ageeligibility"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgeEligibilityDelete is the builder for deleting a AgeEligibility entity.
type AgeEligibilityDelete struct {
	config
	hooks    []Hook
	mutation *AgeEligibilityMutation
}

// Where appends a list predicates to the AgeEligibilityDelete builder.
func (aed *AgeEligibilityDelete) Where(ps ...predicate.AgeEligibility) *AgeEligibilityDelete {
	aed.mutation.Where(ps...)
	return aed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aed *AgeEligibilityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, aed.sqlExec, aed.mutation, aed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aed *AgeEligibilityDelete) ExecX(ctx context.Context) int {
	n, err := aed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aed *AgeEligibilityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ageeligibility.Table, sqlgraph.NewFieldSpec(ageeligibility.FieldID, field.TypeInt32))
	if ps := aed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aed.mutation.done = true
	return affected, err
}

// AgeEligibilityDeleteOne is the builder for deleting a single AgeEligibility entity.
type AgeEligibilityDeleteOne struct {
	aed *AgeEligibilityDelete
}

// Where appends a list predicates to the AgeEligibilityDelete builder.
func (aedo *AgeEligibilityDeleteOne) Where(ps ...predicate.AgeEligibility) *AgeEligibilityDeleteOne {
	aedo.aed.mutation.Where(ps...)
	return aedo
}

// Exec executes the deletion query.
func (aedo *AgeEligibilityDeleteOne) Exec(ctx context.Context) error {
	n, err := aedo.aed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ageeligibility.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aedo *AgeEligibilityDeleteOne) ExecX(ctx context.Context) {
	if err := aedo.Exec(ctx); err != nil {
		panic(err)
	}
}
