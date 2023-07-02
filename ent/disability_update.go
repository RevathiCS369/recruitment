// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/disability"
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DisabilityUpdate is the builder for updating Disability entities.
type DisabilityUpdate struct {
	config
	hooks    []Hook
	mutation *DisabilityMutation
}

// Where appends a list predicates to the DisabilityUpdate builder.
func (du *DisabilityUpdate) Where(ps ...predicate.Disability) *DisabilityUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetDisabilityTypeCode sets the "DisabilityTypeCode" field.
func (du *DisabilityUpdate) SetDisabilityTypeCode(s string) *DisabilityUpdate {
	du.mutation.SetDisabilityTypeCode(s)
	return du
}

// SetDisabilityTypeDescription sets the "DisabilityTypeDescription" field.
func (du *DisabilityUpdate) SetDisabilityTypeDescription(s string) *DisabilityUpdate {
	du.mutation.SetDisabilityTypeDescription(s)
	return du
}

// SetDisabilityPercentage sets the "DisabilityPercentage" field.
func (du *DisabilityUpdate) SetDisabilityPercentage(i int32) *DisabilityUpdate {
	du.mutation.ResetDisabilityPercentage()
	du.mutation.SetDisabilityPercentage(i)
	return du
}

// AddDisabilityPercentage adds i to the "DisabilityPercentage" field.
func (du *DisabilityUpdate) AddDisabilityPercentage(i int32) *DisabilityUpdate {
	du.mutation.AddDisabilityPercentage(i)
	return du
}

// SetDisabilityFlag sets the "DisabilityFlag" field.
func (du *DisabilityUpdate) SetDisabilityFlag(df disability.DisabilityFlag) *DisabilityUpdate {
	du.mutation.SetDisabilityFlag(df)
	return du
}

// Mutation returns the DisabilityMutation object of the builder.
func (du *DisabilityUpdate) Mutation() *DisabilityMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DisabilityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DisabilityUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DisabilityUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DisabilityUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DisabilityUpdate) check() error {
	if v, ok := du.mutation.DisabilityFlag(); ok {
		if err := disability.DisabilityFlagValidator(v); err != nil {
			return &ValidationError{Name: "DisabilityFlag", err: fmt.Errorf(`ent: validator failed for field "Disability.DisabilityFlag": %w`, err)}
		}
	}
	return nil
}

func (du *DisabilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(disability.Table, disability.Columns, sqlgraph.NewFieldSpec(disability.FieldID, field.TypeInt32))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DisabilityTypeCode(); ok {
		_spec.SetField(disability.FieldDisabilityTypeCode, field.TypeString, value)
	}
	if value, ok := du.mutation.DisabilityTypeDescription(); ok {
		_spec.SetField(disability.FieldDisabilityTypeDescription, field.TypeString, value)
	}
	if value, ok := du.mutation.DisabilityPercentage(); ok {
		_spec.SetField(disability.FieldDisabilityPercentage, field.TypeInt32, value)
	}
	if value, ok := du.mutation.AddedDisabilityPercentage(); ok {
		_spec.AddField(disability.FieldDisabilityPercentage, field.TypeInt32, value)
	}
	if value, ok := du.mutation.DisabilityFlag(); ok {
		_spec.SetField(disability.FieldDisabilityFlag, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disability.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DisabilityUpdateOne is the builder for updating a single Disability entity.
type DisabilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DisabilityMutation
}

// SetDisabilityTypeCode sets the "DisabilityTypeCode" field.
func (duo *DisabilityUpdateOne) SetDisabilityTypeCode(s string) *DisabilityUpdateOne {
	duo.mutation.SetDisabilityTypeCode(s)
	return duo
}

// SetDisabilityTypeDescription sets the "DisabilityTypeDescription" field.
func (duo *DisabilityUpdateOne) SetDisabilityTypeDescription(s string) *DisabilityUpdateOne {
	duo.mutation.SetDisabilityTypeDescription(s)
	return duo
}

// SetDisabilityPercentage sets the "DisabilityPercentage" field.
func (duo *DisabilityUpdateOne) SetDisabilityPercentage(i int32) *DisabilityUpdateOne {
	duo.mutation.ResetDisabilityPercentage()
	duo.mutation.SetDisabilityPercentage(i)
	return duo
}

// AddDisabilityPercentage adds i to the "DisabilityPercentage" field.
func (duo *DisabilityUpdateOne) AddDisabilityPercentage(i int32) *DisabilityUpdateOne {
	duo.mutation.AddDisabilityPercentage(i)
	return duo
}

// SetDisabilityFlag sets the "DisabilityFlag" field.
func (duo *DisabilityUpdateOne) SetDisabilityFlag(df disability.DisabilityFlag) *DisabilityUpdateOne {
	duo.mutation.SetDisabilityFlag(df)
	return duo
}

// Mutation returns the DisabilityMutation object of the builder.
func (duo *DisabilityUpdateOne) Mutation() *DisabilityMutation {
	return duo.mutation
}

// Where appends a list predicates to the DisabilityUpdate builder.
func (duo *DisabilityUpdateOne) Where(ps ...predicate.Disability) *DisabilityUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DisabilityUpdateOne) Select(field string, fields ...string) *DisabilityUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Disability entity.
func (duo *DisabilityUpdateOne) Save(ctx context.Context) (*Disability, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DisabilityUpdateOne) SaveX(ctx context.Context) *Disability {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DisabilityUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DisabilityUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DisabilityUpdateOne) check() error {
	if v, ok := duo.mutation.DisabilityFlag(); ok {
		if err := disability.DisabilityFlagValidator(v); err != nil {
			return &ValidationError{Name: "DisabilityFlag", err: fmt.Errorf(`ent: validator failed for field "Disability.DisabilityFlag": %w`, err)}
		}
	}
	return nil
}

func (duo *DisabilityUpdateOne) sqlSave(ctx context.Context) (_node *Disability, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(disability.Table, disability.Columns, sqlgraph.NewFieldSpec(disability.FieldID, field.TypeInt32))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Disability.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, disability.FieldID)
		for _, f := range fields {
			if !disability.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != disability.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.DisabilityTypeCode(); ok {
		_spec.SetField(disability.FieldDisabilityTypeCode, field.TypeString, value)
	}
	if value, ok := duo.mutation.DisabilityTypeDescription(); ok {
		_spec.SetField(disability.FieldDisabilityTypeDescription, field.TypeString, value)
	}
	if value, ok := duo.mutation.DisabilityPercentage(); ok {
		_spec.SetField(disability.FieldDisabilityPercentage, field.TypeInt32, value)
	}
	if value, ok := duo.mutation.AddedDisabilityPercentage(); ok {
		_spec.AddField(disability.FieldDisabilityPercentage, field.TypeInt32, value)
	}
	if value, ok := duo.mutation.DisabilityFlag(); ok {
		_spec.SetField(disability.FieldDisabilityFlag, field.TypeEnum, value)
	}
	_node = &Disability{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{disability.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}