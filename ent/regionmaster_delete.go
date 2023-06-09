// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/predicate"
	"recruit/ent/regionmaster"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RegionMasterDelete is the builder for deleting a RegionMaster entity.
type RegionMasterDelete struct {
	config
	hooks    []Hook
	mutation *RegionMasterMutation
}

// Where appends a list predicates to the RegionMasterDelete builder.
func (rmd *RegionMasterDelete) Where(ps ...predicate.RegionMaster) *RegionMasterDelete {
	rmd.mutation.Where(ps...)
	return rmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rmd *RegionMasterDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rmd.sqlExec, rmd.mutation, rmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rmd *RegionMasterDelete) ExecX(ctx context.Context) int {
	n, err := rmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rmd *RegionMasterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(regionmaster.Table, sqlgraph.NewFieldSpec(regionmaster.FieldID, field.TypeInt32))
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

// RegionMasterDeleteOne is the builder for deleting a single RegionMaster entity.
type RegionMasterDeleteOne struct {
	rmd *RegionMasterDelete
}

// Where appends a list predicates to the RegionMasterDelete builder.
func (rmdo *RegionMasterDeleteOne) Where(ps ...predicate.RegionMaster) *RegionMasterDeleteOne {
	rmdo.rmd.mutation.Where(ps...)
	return rmdo
}

// Exec executes the deletion query.
func (rmdo *RegionMasterDeleteOne) Exec(ctx context.Context) error {
	n, err := rmdo.rmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{regionmaster.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rmdo *RegionMasterDeleteOne) ExecX(ctx context.Context) {
	if err := rmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
