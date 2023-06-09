// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/cadre_choice_ip"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CadreChoiceIPDelete is the builder for deleting a Cadre_Choice_IP entity.
type CadreChoiceIPDelete struct {
	config
	hooks    []Hook
	mutation *CadreChoiceIPMutation
}

// Where appends a list predicates to the CadreChoiceIPDelete builder.
func (ccid *CadreChoiceIPDelete) Where(ps ...predicate.Cadre_Choice_IP) *CadreChoiceIPDelete {
	ccid.mutation.Where(ps...)
	return ccid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccid *CadreChoiceIPDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ccid.sqlExec, ccid.mutation, ccid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ccid *CadreChoiceIPDelete) ExecX(ctx context.Context) int {
	n, err := ccid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccid *CadreChoiceIPDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cadre_choice_ip.Table, sqlgraph.NewFieldSpec(cadre_choice_ip.FieldID, field.TypeInt32))
	if ps := ccid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ccid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ccid.mutation.done = true
	return affected, err
}

// CadreChoiceIPDeleteOne is the builder for deleting a single Cadre_Choice_IP entity.
type CadreChoiceIPDeleteOne struct {
	ccid *CadreChoiceIPDelete
}

// Where appends a list predicates to the CadreChoiceIPDelete builder.
func (ccido *CadreChoiceIPDeleteOne) Where(ps ...predicate.Cadre_Choice_IP) *CadreChoiceIPDeleteOne {
	ccido.ccid.mutation.Where(ps...)
	return ccido
}

// Exec executes the deletion query.
func (ccido *CadreChoiceIPDeleteOne) Exec(ctx context.Context) error {
	n, err := ccido.ccid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cadre_choice_ip.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccido *CadreChoiceIPDeleteOne) ExecX(ctx context.Context) {
	if err := ccido.Exec(ctx); err != nil {
		panic(err)
	}
}
