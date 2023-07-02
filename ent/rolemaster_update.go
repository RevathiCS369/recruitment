// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/adminlogin"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/predicate"
	"recruit/ent/rolemaster"
	"recruit/ent/usermaster"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMasterUpdate is the builder for updating RoleMaster entities.
type RoleMasterUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMasterMutation
}

// Where appends a list predicates to the RoleMasterUpdate builder.
func (rmu *RoleMasterUpdate) Where(ps ...predicate.RoleMaster) *RoleMasterUpdate {
	rmu.mutation.Where(ps...)
	return rmu
}

// SetRoleName sets the "RoleName" field.
func (rmu *RoleMasterUpdate) SetRoleName(s string) *RoleMasterUpdate {
	rmu.mutation.SetRoleName(s)
	return rmu
}

// SetCreatedDate sets the "CreatedDate" field.
func (rmu *RoleMasterUpdate) SetCreatedDate(t time.Time) *RoleMasterUpdate {
	rmu.mutation.SetCreatedDate(t)
	return rmu
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (rmu *RoleMasterUpdate) SetNillableCreatedDate(t *time.Time) *RoleMasterUpdate {
	if t != nil {
		rmu.SetCreatedDate(*t)
	}
	return rmu
}

// ClearCreatedDate clears the value of the "CreatedDate" field.
func (rmu *RoleMasterUpdate) ClearCreatedDate() *RoleMasterUpdate {
	rmu.mutation.ClearCreatedDate()
	return rmu
}

// SetStatus sets the "Status" field.
func (rmu *RoleMasterUpdate) SetStatus(b bool) *RoleMasterUpdate {
	rmu.mutation.SetStatus(b)
	return rmu
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (rmu *RoleMasterUpdate) SetNillableStatus(b *bool) *RoleMasterUpdate {
	if b != nil {
		rmu.SetStatus(*b)
	}
	return rmu
}

// AddRoleIDs adds the "roles" edge to the AdminLogin entity by IDs.
func (rmu *RoleMasterUpdate) AddRoleIDs(ids ...int32) *RoleMasterUpdate {
	rmu.mutation.AddRoleIDs(ids...)
	return rmu
}

// AddRoles adds the "roles" edges to the AdminLogin entity.
func (rmu *RoleMasterUpdate) AddRoles(a ...*AdminLogin) *RoleMasterUpdate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rmu.AddRoleIDs(ids...)
}

// AddRolesRefIDs adds the "Roles_Ref" edge to the UserMaster entity by IDs.
func (rmu *RoleMasterUpdate) AddRolesRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.AddRolesRefIDs(ids...)
	return rmu
}

// AddRolesRef adds the "Roles_Ref" edges to the UserMaster entity.
func (rmu *RoleMasterUpdate) AddRolesRef(u ...*UserMaster) *RoleMasterUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rmu.AddRolesRefIDs(ids...)
}

// AddRolesPSRefIDs adds the "Roles_PS_Ref" edge to the Exam_Applications_PS entity by IDs.
func (rmu *RoleMasterUpdate) AddRolesPSRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.AddRolesPSRefIDs(ids...)
	return rmu
}

// AddRolesPSRef adds the "Roles_PS_Ref" edges to the Exam_Applications_PS entity.
func (rmu *RoleMasterUpdate) AddRolesPSRef(e ...*Exam_Applications_PS) *RoleMasterUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmu.AddRolesPSRefIDs(ids...)
}

// AddRolesIPRefIDs adds the "Roles_IP_Ref" edge to the Exam_Applications_IP entity by IDs.
func (rmu *RoleMasterUpdate) AddRolesIPRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.AddRolesIPRefIDs(ids...)
	return rmu
}

// AddRolesIPRef adds the "Roles_IP_Ref" edges to the Exam_Applications_IP entity.
func (rmu *RoleMasterUpdate) AddRolesIPRef(e ...*Exam_Applications_IP) *RoleMasterUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmu.AddRolesIPRefIDs(ids...)
}

// Mutation returns the RoleMasterMutation object of the builder.
func (rmu *RoleMasterUpdate) Mutation() *RoleMasterMutation {
	return rmu.mutation
}

// ClearRoles clears all "roles" edges to the AdminLogin entity.
func (rmu *RoleMasterUpdate) ClearRoles() *RoleMasterUpdate {
	rmu.mutation.ClearRoles()
	return rmu
}

// RemoveRoleIDs removes the "roles" edge to AdminLogin entities by IDs.
func (rmu *RoleMasterUpdate) RemoveRoleIDs(ids ...int32) *RoleMasterUpdate {
	rmu.mutation.RemoveRoleIDs(ids...)
	return rmu
}

// RemoveRoles removes "roles" edges to AdminLogin entities.
func (rmu *RoleMasterUpdate) RemoveRoles(a ...*AdminLogin) *RoleMasterUpdate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rmu.RemoveRoleIDs(ids...)
}

// ClearRolesRef clears all "Roles_Ref" edges to the UserMaster entity.
func (rmu *RoleMasterUpdate) ClearRolesRef() *RoleMasterUpdate {
	rmu.mutation.ClearRolesRef()
	return rmu
}

// RemoveRolesRefIDs removes the "Roles_Ref" edge to UserMaster entities by IDs.
func (rmu *RoleMasterUpdate) RemoveRolesRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.RemoveRolesRefIDs(ids...)
	return rmu
}

// RemoveRolesRef removes "Roles_Ref" edges to UserMaster entities.
func (rmu *RoleMasterUpdate) RemoveRolesRef(u ...*UserMaster) *RoleMasterUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rmu.RemoveRolesRefIDs(ids...)
}

// ClearRolesPSRef clears all "Roles_PS_Ref" edges to the Exam_Applications_PS entity.
func (rmu *RoleMasterUpdate) ClearRolesPSRef() *RoleMasterUpdate {
	rmu.mutation.ClearRolesPSRef()
	return rmu
}

// RemoveRolesPSRefIDs removes the "Roles_PS_Ref" edge to Exam_Applications_PS entities by IDs.
func (rmu *RoleMasterUpdate) RemoveRolesPSRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.RemoveRolesPSRefIDs(ids...)
	return rmu
}

// RemoveRolesPSRef removes "Roles_PS_Ref" edges to Exam_Applications_PS entities.
func (rmu *RoleMasterUpdate) RemoveRolesPSRef(e ...*Exam_Applications_PS) *RoleMasterUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmu.RemoveRolesPSRefIDs(ids...)
}

// ClearRolesIPRef clears all "Roles_IP_Ref" edges to the Exam_Applications_IP entity.
func (rmu *RoleMasterUpdate) ClearRolesIPRef() *RoleMasterUpdate {
	rmu.mutation.ClearRolesIPRef()
	return rmu
}

// RemoveRolesIPRefIDs removes the "Roles_IP_Ref" edge to Exam_Applications_IP entities by IDs.
func (rmu *RoleMasterUpdate) RemoveRolesIPRefIDs(ids ...int64) *RoleMasterUpdate {
	rmu.mutation.RemoveRolesIPRefIDs(ids...)
	return rmu
}

// RemoveRolesIPRef removes "Roles_IP_Ref" edges to Exam_Applications_IP entities.
func (rmu *RoleMasterUpdate) RemoveRolesIPRef(e ...*Exam_Applications_IP) *RoleMasterUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmu.RemoveRolesIPRefIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rmu *RoleMasterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rmu.sqlSave, rmu.mutation, rmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmu *RoleMasterUpdate) SaveX(ctx context.Context) int {
	affected, err := rmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rmu *RoleMasterUpdate) Exec(ctx context.Context) error {
	_, err := rmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmu *RoleMasterUpdate) ExecX(ctx context.Context) {
	if err := rmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rmu *RoleMasterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rolemaster.Table, rolemaster.Columns, sqlgraph.NewFieldSpec(rolemaster.FieldID, field.TypeInt32))
	if ps := rmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmu.mutation.RoleName(); ok {
		_spec.SetField(rolemaster.FieldRoleName, field.TypeString, value)
	}
	if value, ok := rmu.mutation.CreatedDate(); ok {
		_spec.SetField(rolemaster.FieldCreatedDate, field.TypeTime, value)
	}
	if rmu.mutation.CreatedDateCleared() {
		_spec.ClearField(rolemaster.FieldCreatedDate, field.TypeTime)
	}
	if value, ok := rmu.mutation.Status(); ok {
		_spec.SetField(rolemaster.FieldStatus, field.TypeBool, value)
	}
	if rmu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !rmu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmu.mutation.RolesRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RemovedRolesRefIDs(); len(nodes) > 0 && !rmu.mutation.RolesRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RolesRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmu.mutation.RolesPSRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RemovedRolesPSRefIDs(); len(nodes) > 0 && !rmu.mutation.RolesPSRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RolesPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmu.mutation.RolesIPRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RemovedRolesIPRefIDs(); len(nodes) > 0 && !rmu.mutation.RolesIPRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RolesIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolemaster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rmu.mutation.done = true
	return n, nil
}

// RoleMasterUpdateOne is the builder for updating a single RoleMaster entity.
type RoleMasterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMasterMutation
}

// SetRoleName sets the "RoleName" field.
func (rmuo *RoleMasterUpdateOne) SetRoleName(s string) *RoleMasterUpdateOne {
	rmuo.mutation.SetRoleName(s)
	return rmuo
}

// SetCreatedDate sets the "CreatedDate" field.
func (rmuo *RoleMasterUpdateOne) SetCreatedDate(t time.Time) *RoleMasterUpdateOne {
	rmuo.mutation.SetCreatedDate(t)
	return rmuo
}

// SetNillableCreatedDate sets the "CreatedDate" field if the given value is not nil.
func (rmuo *RoleMasterUpdateOne) SetNillableCreatedDate(t *time.Time) *RoleMasterUpdateOne {
	if t != nil {
		rmuo.SetCreatedDate(*t)
	}
	return rmuo
}

// ClearCreatedDate clears the value of the "CreatedDate" field.
func (rmuo *RoleMasterUpdateOne) ClearCreatedDate() *RoleMasterUpdateOne {
	rmuo.mutation.ClearCreatedDate()
	return rmuo
}

// SetStatus sets the "Status" field.
func (rmuo *RoleMasterUpdateOne) SetStatus(b bool) *RoleMasterUpdateOne {
	rmuo.mutation.SetStatus(b)
	return rmuo
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (rmuo *RoleMasterUpdateOne) SetNillableStatus(b *bool) *RoleMasterUpdateOne {
	if b != nil {
		rmuo.SetStatus(*b)
	}
	return rmuo
}

// AddRoleIDs adds the "roles" edge to the AdminLogin entity by IDs.
func (rmuo *RoleMasterUpdateOne) AddRoleIDs(ids ...int32) *RoleMasterUpdateOne {
	rmuo.mutation.AddRoleIDs(ids...)
	return rmuo
}

// AddRoles adds the "roles" edges to the AdminLogin entity.
func (rmuo *RoleMasterUpdateOne) AddRoles(a ...*AdminLogin) *RoleMasterUpdateOne {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rmuo.AddRoleIDs(ids...)
}

// AddRolesRefIDs adds the "Roles_Ref" edge to the UserMaster entity by IDs.
func (rmuo *RoleMasterUpdateOne) AddRolesRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.AddRolesRefIDs(ids...)
	return rmuo
}

// AddRolesRef adds the "Roles_Ref" edges to the UserMaster entity.
func (rmuo *RoleMasterUpdateOne) AddRolesRef(u ...*UserMaster) *RoleMasterUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rmuo.AddRolesRefIDs(ids...)
}

// AddRolesPSRefIDs adds the "Roles_PS_Ref" edge to the Exam_Applications_PS entity by IDs.
func (rmuo *RoleMasterUpdateOne) AddRolesPSRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.AddRolesPSRefIDs(ids...)
	return rmuo
}

// AddRolesPSRef adds the "Roles_PS_Ref" edges to the Exam_Applications_PS entity.
func (rmuo *RoleMasterUpdateOne) AddRolesPSRef(e ...*Exam_Applications_PS) *RoleMasterUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmuo.AddRolesPSRefIDs(ids...)
}

// AddRolesIPRefIDs adds the "Roles_IP_Ref" edge to the Exam_Applications_IP entity by IDs.
func (rmuo *RoleMasterUpdateOne) AddRolesIPRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.AddRolesIPRefIDs(ids...)
	return rmuo
}

// AddRolesIPRef adds the "Roles_IP_Ref" edges to the Exam_Applications_IP entity.
func (rmuo *RoleMasterUpdateOne) AddRolesIPRef(e ...*Exam_Applications_IP) *RoleMasterUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmuo.AddRolesIPRefIDs(ids...)
}

// Mutation returns the RoleMasterMutation object of the builder.
func (rmuo *RoleMasterUpdateOne) Mutation() *RoleMasterMutation {
	return rmuo.mutation
}

// ClearRoles clears all "roles" edges to the AdminLogin entity.
func (rmuo *RoleMasterUpdateOne) ClearRoles() *RoleMasterUpdateOne {
	rmuo.mutation.ClearRoles()
	return rmuo
}

// RemoveRoleIDs removes the "roles" edge to AdminLogin entities by IDs.
func (rmuo *RoleMasterUpdateOne) RemoveRoleIDs(ids ...int32) *RoleMasterUpdateOne {
	rmuo.mutation.RemoveRoleIDs(ids...)
	return rmuo
}

// RemoveRoles removes "roles" edges to AdminLogin entities.
func (rmuo *RoleMasterUpdateOne) RemoveRoles(a ...*AdminLogin) *RoleMasterUpdateOne {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rmuo.RemoveRoleIDs(ids...)
}

// ClearRolesRef clears all "Roles_Ref" edges to the UserMaster entity.
func (rmuo *RoleMasterUpdateOne) ClearRolesRef() *RoleMasterUpdateOne {
	rmuo.mutation.ClearRolesRef()
	return rmuo
}

// RemoveRolesRefIDs removes the "Roles_Ref" edge to UserMaster entities by IDs.
func (rmuo *RoleMasterUpdateOne) RemoveRolesRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.RemoveRolesRefIDs(ids...)
	return rmuo
}

// RemoveRolesRef removes "Roles_Ref" edges to UserMaster entities.
func (rmuo *RoleMasterUpdateOne) RemoveRolesRef(u ...*UserMaster) *RoleMasterUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rmuo.RemoveRolesRefIDs(ids...)
}

// ClearRolesPSRef clears all "Roles_PS_Ref" edges to the Exam_Applications_PS entity.
func (rmuo *RoleMasterUpdateOne) ClearRolesPSRef() *RoleMasterUpdateOne {
	rmuo.mutation.ClearRolesPSRef()
	return rmuo
}

// RemoveRolesPSRefIDs removes the "Roles_PS_Ref" edge to Exam_Applications_PS entities by IDs.
func (rmuo *RoleMasterUpdateOne) RemoveRolesPSRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.RemoveRolesPSRefIDs(ids...)
	return rmuo
}

// RemoveRolesPSRef removes "Roles_PS_Ref" edges to Exam_Applications_PS entities.
func (rmuo *RoleMasterUpdateOne) RemoveRolesPSRef(e ...*Exam_Applications_PS) *RoleMasterUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmuo.RemoveRolesPSRefIDs(ids...)
}

// ClearRolesIPRef clears all "Roles_IP_Ref" edges to the Exam_Applications_IP entity.
func (rmuo *RoleMasterUpdateOne) ClearRolesIPRef() *RoleMasterUpdateOne {
	rmuo.mutation.ClearRolesIPRef()
	return rmuo
}

// RemoveRolesIPRefIDs removes the "Roles_IP_Ref" edge to Exam_Applications_IP entities by IDs.
func (rmuo *RoleMasterUpdateOne) RemoveRolesIPRefIDs(ids ...int64) *RoleMasterUpdateOne {
	rmuo.mutation.RemoveRolesIPRefIDs(ids...)
	return rmuo
}

// RemoveRolesIPRef removes "Roles_IP_Ref" edges to Exam_Applications_IP entities.
func (rmuo *RoleMasterUpdateOne) RemoveRolesIPRef(e ...*Exam_Applications_IP) *RoleMasterUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rmuo.RemoveRolesIPRefIDs(ids...)
}

// Where appends a list predicates to the RoleMasterUpdate builder.
func (rmuo *RoleMasterUpdateOne) Where(ps ...predicate.RoleMaster) *RoleMasterUpdateOne {
	rmuo.mutation.Where(ps...)
	return rmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rmuo *RoleMasterUpdateOne) Select(field string, fields ...string) *RoleMasterUpdateOne {
	rmuo.fields = append([]string{field}, fields...)
	return rmuo
}

// Save executes the query and returns the updated RoleMaster entity.
func (rmuo *RoleMasterUpdateOne) Save(ctx context.Context) (*RoleMaster, error) {
	return withHooks(ctx, rmuo.sqlSave, rmuo.mutation, rmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmuo *RoleMasterUpdateOne) SaveX(ctx context.Context) *RoleMaster {
	node, err := rmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rmuo *RoleMasterUpdateOne) Exec(ctx context.Context) error {
	_, err := rmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmuo *RoleMasterUpdateOne) ExecX(ctx context.Context) {
	if err := rmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rmuo *RoleMasterUpdateOne) sqlSave(ctx context.Context) (_node *RoleMaster, err error) {
	_spec := sqlgraph.NewUpdateSpec(rolemaster.Table, rolemaster.Columns, sqlgraph.NewFieldSpec(rolemaster.FieldID, field.TypeInt32))
	id, ok := rmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoleMaster.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolemaster.FieldID)
		for _, f := range fields {
			if !rolemaster.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rolemaster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmuo.mutation.RoleName(); ok {
		_spec.SetField(rolemaster.FieldRoleName, field.TypeString, value)
	}
	if value, ok := rmuo.mutation.CreatedDate(); ok {
		_spec.SetField(rolemaster.FieldCreatedDate, field.TypeTime, value)
	}
	if rmuo.mutation.CreatedDateCleared() {
		_spec.ClearField(rolemaster.FieldCreatedDate, field.TypeTime)
	}
	if value, ok := rmuo.mutation.Status(); ok {
		_spec.SetField(rolemaster.FieldStatus, field.TypeBool, value)
	}
	if rmuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !rmuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesTable,
			Columns: []string{rolemaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlogin.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmuo.mutation.RolesRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RemovedRolesRefIDs(); len(nodes) > 0 && !rmuo.mutation.RolesRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RolesRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesRefTable,
			Columns: []string{rolemaster.RolesRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmuo.mutation.RolesPSRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RemovedRolesPSRefIDs(); len(nodes) > 0 && !rmuo.mutation.RolesPSRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RolesPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesPSRefTable,
			Columns: []string{rolemaster.RolesPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmuo.mutation.RolesIPRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RemovedRolesIPRefIDs(); len(nodes) > 0 && !rmuo.mutation.RolesIPRefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RolesIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rolemaster.RolesIPRefTable,
			Columns: []string{rolemaster.RolesIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RoleMaster{config: rmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolemaster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rmuo.mutation.done = true
	return _node, nil
}
