// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/predicate"
	"recruit/ent/usermaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserMasterDelete is the builder for deleting a UserMaster entity.
type UserMasterDelete struct {
	config
	hooks    []Hook
	mutation *UserMasterMutation
}

// Where appends a list predicates to the UserMasterDelete builder.
func (umd *UserMasterDelete) Where(ps ...predicate.UserMaster) *UserMasterDelete {
	umd.mutation.Where(ps...)
	return umd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (umd *UserMasterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, umd.sqlExec, umd.mutation, umd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (umd *UserMasterDelete) ExecX(ctx context.Context) int {
	n, err := umd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (umd *UserMasterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usermaster.Table, sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64))
	if ps := umd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, umd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	umd.mutation.done = true
	return affected, err
}

// UserMasterDeleteOne is the builder for deleting a single UserMaster entity.
type UserMasterDeleteOne struct {
	umd *UserMasterDelete
}

// Where appends a list predicates to the UserMasterDelete builder.
func (umdo *UserMasterDeleteOne) Where(ps ...predicate.UserMaster) *UserMasterDeleteOne {
	umdo.umd.mutation.Where(ps...)
	return umdo
}

// Exec executes the deletion query.
func (umdo *UserMasterDeleteOne) Exec(ctx context.Context) error {
	n, err := umdo.umd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usermaster.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (umdo *UserMasterDeleteOne) ExecX(ctx context.Context) {
	if err := umdo.Exec(ctx); err != nil {
		panic(err)
	}
}
