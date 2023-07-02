// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"recruit/ent/exam_ps"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamPSDelete is the builder for deleting a Exam_PS entity.
type ExamPSDelete struct {
	config
	hooks    []Hook
	mutation *ExamPSMutation
}

// Where appends a list predicates to the ExamPSDelete builder.
func (epd *ExamPSDelete) Where(ps ...predicate.Exam_PS) *ExamPSDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *ExamPSDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epd.sqlExec, epd.mutation, epd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *ExamPSDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *ExamPSDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(exam_ps.Table, sqlgraph.NewFieldSpec(exam_ps.FieldID, field.TypeInt32))
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

// ExamPSDeleteOne is the builder for deleting a single Exam_PS entity.
type ExamPSDeleteOne struct {
	epd *ExamPSDelete
}

// Where appends a list predicates to the ExamPSDelete builder.
func (epdo *ExamPSDeleteOne) Where(ps ...predicate.Exam_PS) *ExamPSDeleteOne {
	epdo.epd.mutation.Where(ps...)
	return epdo
}

// Exec executes the deletion query.
func (epdo *ExamPSDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exam_ps.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *ExamPSDeleteOne) ExecX(ctx context.Context) {
	if err := epdo.Exec(ctx); err != nil {
		panic(err)
	}
}