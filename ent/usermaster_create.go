// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/employeemaster"
	"recruit/ent/exam_applications_ip"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/exam_ip"
	"recruit/ent/exam_ps"
	"recruit/ent/rolemaster"
	"recruit/ent/usermaster"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserMasterCreate is the builder for creating a UserMaster entity.
type UserMasterCreate struct {
	config
	mutation *UserMasterMutation
	hooks    []Hook
}

// SetEmployeeID sets the "EmployeeID" field.
func (umc *UserMasterCreate) SetEmployeeID(i int64) *UserMasterCreate {
	umc.mutation.SetEmployeeID(i)
	return umc
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableEmployeeID(i *int64) *UserMasterCreate {
	if i != nil {
		umc.SetEmployeeID(*i)
	}
	return umc
}

// SetEmployeeName sets the "EmployeeName" field.
func (umc *UserMasterCreate) SetEmployeeName(s string) *UserMasterCreate {
	umc.mutation.SetEmployeeName(s)
	return umc
}

// SetNillableEmployeeName sets the "EmployeeName" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableEmployeeName(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetEmployeeName(*s)
	}
	return umc
}

// SetFacilityID sets the "FacilityID" field.
func (umc *UserMasterCreate) SetFacilityID(s string) *UserMasterCreate {
	umc.mutation.SetFacilityID(s)
	return umc
}

// SetNillableFacilityID sets the "FacilityID" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableFacilityID(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetFacilityID(*s)
	}
	return umc
}

// SetCadre sets the "Cadre" field.
func (umc *UserMasterCreate) SetCadre(s string) *UserMasterCreate {
	umc.mutation.SetCadre(s)
	return umc
}

// SetNillableCadre sets the "Cadre" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableCadre(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetCadre(*s)
	}
	return umc
}

// SetRoleUserCode sets the "RoleUserCode" field.
func (umc *UserMasterCreate) SetRoleUserCode(i int32) *UserMasterCreate {
	umc.mutation.SetRoleUserCode(i)
	return umc
}

// SetNillableRoleUserCode sets the "RoleUserCode" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableRoleUserCode(i *int32) *UserMasterCreate {
	if i != nil {
		umc.SetRoleUserCode(*i)
	}
	return umc
}

// SetMobile sets the "Mobile" field.
func (umc *UserMasterCreate) SetMobile(s string) *UserMasterCreate {
	umc.mutation.SetMobile(s)
	return umc
}

// SetNillableMobile sets the "Mobile" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableMobile(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetMobile(*s)
	}
	return umc
}

// SetEmailID sets the "EmailID" field.
func (umc *UserMasterCreate) SetEmailID(s string) *UserMasterCreate {
	umc.mutation.SetEmailID(s)
	return umc
}

// SetNillableEmailID sets the "EmailID" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableEmailID(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetEmailID(*s)
	}
	return umc
}

// SetUserName sets the "UserName" field.
func (umc *UserMasterCreate) SetUserName(s string) *UserMasterCreate {
	umc.mutation.SetUserName(s)
	return umc
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableUserName(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetUserName(*s)
	}
	return umc
}

// SetPassword sets the "Password" field.
func (umc *UserMasterCreate) SetPassword(s string) *UserMasterCreate {
	umc.mutation.SetPassword(s)
	return umc
}

// SetNillablePassword sets the "Password" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillablePassword(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetPassword(*s)
	}
	return umc
}

// SetOTP sets the "OTP" field.
func (umc *UserMasterCreate) SetOTP(i int32) *UserMasterCreate {
	umc.mutation.SetOTP(i)
	return umc
}

// SetNillableOTP sets the "OTP" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableOTP(i *int32) *UserMasterCreate {
	if i != nil {
		umc.SetOTP(*i)
	}
	return umc
}

// SetExamCode sets the "ExamCode" field.
func (umc *UserMasterCreate) SetExamCode(i int32) *UserMasterCreate {
	umc.mutation.SetExamCode(i)
	return umc
}

// SetNillableExamCode sets the "ExamCode" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableExamCode(i *int32) *UserMasterCreate {
	if i != nil {
		umc.SetExamCode(*i)
	}
	return umc
}

// SetExamCodePS sets the "ExamCodePS" field.
func (umc *UserMasterCreate) SetExamCodePS(i int32) *UserMasterCreate {
	umc.mutation.SetExamCodePS(i)
	return umc
}

// SetNillableExamCodePS sets the "ExamCodePS" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableExamCodePS(i *int32) *UserMasterCreate {
	if i != nil {
		umc.SetExamCodePS(*i)
	}
	return umc
}

// SetOTPRemarks sets the "OTPRemarks" field.
func (umc *UserMasterCreate) SetOTPRemarks(s string) *UserMasterCreate {
	umc.mutation.SetOTPRemarks(s)
	return umc
}

// SetNillableOTPRemarks sets the "OTPRemarks" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableOTPRemarks(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetOTPRemarks(*s)
	}
	return umc
}

// SetStatus sets the "Status" field.
func (umc *UserMasterCreate) SetStatus(b bool) *UserMasterCreate {
	umc.mutation.SetStatus(b)
	return umc
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableStatus(b *bool) *UserMasterCreate {
	if b != nil {
		umc.SetStatus(*b)
	}
	return umc
}

// SetNewPasswordRequest sets the "NewPasswordRequest" field.
func (umc *UserMasterCreate) SetNewPasswordRequest(b bool) *UserMasterCreate {
	umc.mutation.SetNewPasswordRequest(b)
	return umc
}

// SetNillableNewPasswordRequest sets the "NewPasswordRequest" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableNewPasswordRequest(b *bool) *UserMasterCreate {
	if b != nil {
		umc.SetNewPasswordRequest(*b)
	}
	return umc
}

// SetCreatedAt sets the "CreatedAt" field.
func (umc *UserMasterCreate) SetCreatedAt(t time.Time) *UserMasterCreate {
	umc.mutation.SetCreatedAt(t)
	return umc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableCreatedAt(t *time.Time) *UserMasterCreate {
	if t != nil {
		umc.SetCreatedAt(*t)
	}
	return umc
}

// SetOTPTriggeredTime sets the "OTPTriggeredTime" field.
func (umc *UserMasterCreate) SetOTPTriggeredTime(t time.Time) *UserMasterCreate {
	umc.mutation.SetOTPTriggeredTime(t)
	return umc
}

// SetNillableOTPTriggeredTime sets the "OTPTriggeredTime" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableOTPTriggeredTime(t *time.Time) *UserMasterCreate {
	if t != nil {
		umc.SetOTPTriggeredTime(*t)
	}
	return umc
}

// SetCreatedBy sets the "CreatedBy" field.
func (umc *UserMasterCreate) SetCreatedBy(s string) *UserMasterCreate {
	umc.mutation.SetCreatedBy(s)
	return umc
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (umc *UserMasterCreate) SetNillableCreatedBy(s *string) *UserMasterCreate {
	if s != nil {
		umc.SetCreatedBy(*s)
	}
	return umc
}

// SetID sets the "id" field.
func (umc *UserMasterCreate) SetID(i int64) *UserMasterCreate {
	umc.mutation.SetID(i)
	return umc
}

// SetRolesID sets the "roles" edge to the RoleMaster entity by ID.
func (umc *UserMasterCreate) SetRolesID(id int32) *UserMasterCreate {
	umc.mutation.SetRolesID(id)
	return umc
}

// SetNillableRolesID sets the "roles" edge to the RoleMaster entity by ID if the given value is not nil.
func (umc *UserMasterCreate) SetNillableRolesID(id *int32) *UserMasterCreate {
	if id != nil {
		umc = umc.SetRolesID(*id)
	}
	return umc
}

// SetRoles sets the "roles" edge to the RoleMaster entity.
func (umc *UserMasterCreate) SetRoles(r *RoleMaster) *UserMasterCreate {
	return umc.SetRolesID(r.ID)
}

// AddUsermasterRefIDs adds the "UsermasterRef" edge to the EmployeeMaster entity by IDs.
func (umc *UserMasterCreate) AddUsermasterRefIDs(ids ...int64) *UserMasterCreate {
	umc.mutation.AddUsermasterRefIDs(ids...)
	return umc
}

// AddUsermasterRef adds the "UsermasterRef" edges to the EmployeeMaster entity.
func (umc *UserMasterCreate) AddUsermasterRef(e ...*EmployeeMaster) *UserMasterCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return umc.AddUsermasterRefIDs(ids...)
}

// AddUsersPSRefIDs adds the "UsersPSRef" edge to the Exam_Applications_PS entity by IDs.
func (umc *UserMasterCreate) AddUsersPSRefIDs(ids ...int64) *UserMasterCreate {
	umc.mutation.AddUsersPSRefIDs(ids...)
	return umc
}

// AddUsersPSRef adds the "UsersPSRef" edges to the Exam_Applications_PS entity.
func (umc *UserMasterCreate) AddUsersPSRef(e ...*Exam_Applications_PS) *UserMasterCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return umc.AddUsersPSRefIDs(ids...)
}

// AddUsersIPRefIDs adds the "UsersIPRef" edge to the Exam_Applications_IP entity by IDs.
func (umc *UserMasterCreate) AddUsersIPRefIDs(ids ...int64) *UserMasterCreate {
	umc.mutation.AddUsersIPRefIDs(ids...)
	return umc
}

// AddUsersIPRef adds the "UsersIPRef" edges to the Exam_Applications_IP entity.
func (umc *UserMasterCreate) AddUsersIPRef(e ...*Exam_Applications_IP) *UserMasterCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return umc.AddUsersIPRefIDs(ids...)
}

// AddUsersPsTypeIDs adds the "users_ps_type" edge to the Exam_PS entity by IDs.
func (umc *UserMasterCreate) AddUsersPsTypeIDs(ids ...int32) *UserMasterCreate {
	umc.mutation.AddUsersPsTypeIDs(ids...)
	return umc
}

// AddUsersPsType adds the "users_ps_type" edges to the Exam_PS entity.
func (umc *UserMasterCreate) AddUsersPsType(e ...*Exam_PS) *UserMasterCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return umc.AddUsersPsTypeIDs(ids...)
}

// AddUsersIPTypeIDs adds the "users_ip_type" edge to the Exam_IP entity by IDs.
func (umc *UserMasterCreate) AddUsersIPTypeIDs(ids ...int32) *UserMasterCreate {
	umc.mutation.AddUsersIPTypeIDs(ids...)
	return umc
}

// AddUsersIPType adds the "users_ip_type" edges to the Exam_IP entity.
func (umc *UserMasterCreate) AddUsersIPType(e ...*Exam_IP) *UserMasterCreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return umc.AddUsersIPTypeIDs(ids...)
}

// Mutation returns the UserMasterMutation object of the builder.
func (umc *UserMasterCreate) Mutation() *UserMasterMutation {
	return umc.mutation
}

// Save creates the UserMaster in the database.
func (umc *UserMasterCreate) Save(ctx context.Context) (*UserMaster, error) {
	umc.defaults()
	return withHooks(ctx, umc.sqlSave, umc.mutation, umc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (umc *UserMasterCreate) SaveX(ctx context.Context) *UserMaster {
	v, err := umc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umc *UserMasterCreate) Exec(ctx context.Context) error {
	_, err := umc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umc *UserMasterCreate) ExecX(ctx context.Context) {
	if err := umc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (umc *UserMasterCreate) defaults() {
	if _, ok := umc.mutation.Status(); !ok {
		v := usermaster.DefaultStatus
		umc.mutation.SetStatus(v)
	}
	if _, ok := umc.mutation.OTPTriggeredTime(); !ok {
		v := usermaster.DefaultOTPTriggeredTime()
		umc.mutation.SetOTPTriggeredTime(v)
	}
	if _, ok := umc.mutation.CreatedBy(); !ok {
		v := usermaster.DefaultCreatedBy
		umc.mutation.SetCreatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umc *UserMasterCreate) check() error {
	return nil
}

func (umc *UserMasterCreate) sqlSave(ctx context.Context) (*UserMaster, error) {
	if err := umc.check(); err != nil {
		return nil, err
	}
	_node, _spec := umc.createSpec()
	if err := sqlgraph.CreateNode(ctx, umc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	umc.mutation.id = &_node.ID
	umc.mutation.done = true
	return _node, nil
}

func (umc *UserMasterCreate) createSpec() (*UserMaster, *sqlgraph.CreateSpec) {
	var (
		_node = &UserMaster{config: umc.config}
		_spec = sqlgraph.NewCreateSpec(usermaster.Table, sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64))
	)
	if id, ok := umc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := umc.mutation.EmployeeID(); ok {
		_spec.SetField(usermaster.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := umc.mutation.EmployeeName(); ok {
		_spec.SetField(usermaster.FieldEmployeeName, field.TypeString, value)
		_node.EmployeeName = value
	}
	if value, ok := umc.mutation.FacilityID(); ok {
		_spec.SetField(usermaster.FieldFacilityID, field.TypeString, value)
		_node.FacilityID = value
	}
	if value, ok := umc.mutation.Cadre(); ok {
		_spec.SetField(usermaster.FieldCadre, field.TypeString, value)
		_node.Cadre = value
	}
	if value, ok := umc.mutation.Mobile(); ok {
		_spec.SetField(usermaster.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := umc.mutation.EmailID(); ok {
		_spec.SetField(usermaster.FieldEmailID, field.TypeString, value)
		_node.EmailID = value
	}
	if value, ok := umc.mutation.UserName(); ok {
		_spec.SetField(usermaster.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := umc.mutation.Password(); ok {
		_spec.SetField(usermaster.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := umc.mutation.OTP(); ok {
		_spec.SetField(usermaster.FieldOTP, field.TypeInt32, value)
		_node.OTP = value
	}
	if value, ok := umc.mutation.ExamCode(); ok {
		_spec.SetField(usermaster.FieldExamCode, field.TypeInt32, value)
		_node.ExamCode = value
	}
	if value, ok := umc.mutation.ExamCodePS(); ok {
		_spec.SetField(usermaster.FieldExamCodePS, field.TypeInt32, value)
		_node.ExamCodePS = value
	}
	if value, ok := umc.mutation.OTPRemarks(); ok {
		_spec.SetField(usermaster.FieldOTPRemarks, field.TypeString, value)
		_node.OTPRemarks = value
	}
	if value, ok := umc.mutation.Status(); ok {
		_spec.SetField(usermaster.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := umc.mutation.NewPasswordRequest(); ok {
		_spec.SetField(usermaster.FieldNewPasswordRequest, field.TypeBool, value)
		_node.NewPasswordRequest = value
	}
	if value, ok := umc.mutation.CreatedAt(); ok {
		_spec.SetField(usermaster.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := umc.mutation.OTPTriggeredTime(); ok {
		_spec.SetField(usermaster.FieldOTPTriggeredTime, field.TypeTime, value)
		_node.OTPTriggeredTime = value
	}
	if value, ok := umc.mutation.CreatedBy(); ok {
		_spec.SetField(usermaster.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if nodes := umc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermaster.RolesTable,
			Columns: []string{usermaster.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rolemaster.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleUserCode = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umc.mutation.UsermasterRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usermaster.UsermasterRefTable,
			Columns: []string{usermaster.UsermasterRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employeemaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umc.mutation.UsersPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usermaster.UsersPSRefTable,
			Columns: []string{usermaster.UsersPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umc.mutation.UsersIPRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usermaster.UsersIPRefTable,
			Columns: []string{usermaster.UsersIPRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ip.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umc.mutation.UsersPsTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usermaster.UsersPsTypeTable,
			Columns: []string{usermaster.UsersPsTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_ps.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umc.mutation.UsersIPTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usermaster.UsersIPTypeTable,
			Columns: []string{usermaster.UsersIPTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_ip.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserMasterCreateBulk is the builder for creating many UserMaster entities in bulk.
type UserMasterCreateBulk struct {
	config
	builders []*UserMasterCreate
}

// Save creates the UserMaster entities in the database.
func (umcb *UserMasterCreateBulk) Save(ctx context.Context) ([]*UserMaster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(umcb.builders))
	nodes := make([]*UserMaster, len(umcb.builders))
	mutators := make([]Mutator, len(umcb.builders))
	for i := range umcb.builders {
		func(i int, root context.Context) {
			builder := umcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMasterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, umcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, umcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, umcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (umcb *UserMasterCreateBulk) SaveX(ctx context.Context) []*UserMaster {
	v, err := umcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umcb *UserMasterCreateBulk) Exec(ctx context.Context) error {
	_, err := umcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umcb *UserMasterCreateBulk) ExecX(ctx context.Context) {
	if err := umcb.Exec(ctx); err != nil {
		panic(err)
	}
}
