// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/exampapers"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamPapersDelete is the builder for deleting a ExamPapers entity.
type ExamPapersDelete struct {
	config
	hooks    []Hook
	mutation *ExamPapersMutation
}

// Where appends a list predicates to the ExamPapersDelete builder.
func (epd *ExamPapersDelete) Where(ps ...predicate.ExamPapers) *ExamPapersDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *ExamPapersDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epd.sqlExec, epd.mutation, epd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *ExamPapersDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *ExamPapersDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(exampapers.Table, sqlgraph.NewFieldSpec(exampapers.FieldID, field.TypeInt32))
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	epd.mutation.done = true
	return affected, err
}

// ExamPapersDeleteOne is the builder for deleting a single ExamPapers entity.
type ExamPapersDeleteOne struct {
	epd *ExamPapersDelete
}

// Where appends a list predicates to the ExamPapersDelete builder.
func (epdo *ExamPapersDeleteOne) Where(ps ...predicate.ExamPapers) *ExamPapersDeleteOne {
	epdo.epd.mutation.Where(ps...)
	return epdo
}

// Exec executes the deletion query.
func (epdo *ExamPapersDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exampapers.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *ExamPapersDeleteOne) ExecX(ctx context.Context) {
	if err := epdo.Exec(ctx); err != nil {
		panic(err)
	}
}
