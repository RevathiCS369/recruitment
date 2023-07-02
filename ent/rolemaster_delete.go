// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/predicate"
	"recruit/ent/rolemaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMasterDelete is the builder for deleting a RoleMaster entity.
type RoleMasterDelete struct {
	config
	hooks    []Hook
	mutation *RoleMasterMutation
}

// Where appends a list predicates to the RoleMasterDelete builder.
func (rmd *RoleMasterDelete) Where(ps ...predicate.RoleMaster) *RoleMasterDelete {
	rmd.mutation.Where(ps...)
	return rmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rmd *RoleMasterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rmd.sqlExec, rmd.mutation, rmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rmd *RoleMasterDelete) ExecX(ctx context.Context) int {
	n, err := rmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rmd *RoleMasterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rolemaster.Table, sqlgraph.NewFieldSpec(rolemaster.FieldID, field.TypeInt32))
	if ps := rmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rmd.mutation.done = true
	return affected, err
}

// RoleMasterDeleteOne is the builder for deleting a single RoleMaster entity.
type RoleMasterDeleteOne struct {
	rmd *RoleMasterDelete
}

// Where appends a list predicates to the RoleMasterDelete builder.
func (rmdo *RoleMasterDeleteOne) Where(ps ...predicate.RoleMaster) *RoleMasterDeleteOne {
	rmdo.rmd.mutation.Where(ps...)
	return rmdo
}

// Exec executes the deletion query.
func (rmdo *RoleMasterDeleteOne) Exec(ctx context.Context) error {
	n, err := rmdo.rmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rolemaster.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rmdo *RoleMasterDeleteOne) ExecX(ctx context.Context) {
	if err := rmdo.Exec(ctx); err != nil {
		panic(err)
	}
}