// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/rolemaster"
	"recruit/ent/usermaster"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserMaster is the model entity for the UserMaster schema.
type UserMaster struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int64 `json:"EmployeeID,omitempty"`
	// EmployeeName holds the value of the "EmployeeName" field.
	EmployeeName string `json:"EmployeeName,omitempty"`
	// FacilityID holds the value of the "FacilityID" field.
	FacilityID string `json:"FacilityID,omitempty"`
	// Cadre holds the value of the "Cadre" field.
	Cadre string `json:"Cadre,omitempty"`
	// RoleUserCode holds the value of the "RoleUserCode" field.
	RoleUserCode int32 `json:"RoleUserCode,omitempty"`
	// Mobile holds the value of the "Mobile" field.
	Mobile string `json:"Mobile,omitempty"`
	// EmailID holds the value of the "EmailID" field.
	EmailID string `json:"EmailID,omitempty"`
	// UserName holds the value of the "UserName" field.
	UserName string `json:"UserName,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// OTP holds the value of the "OTP" field.
	OTP int32 `json:"OTP,omitempty"`
	// ExamCode holds the value of the "ExamCode" field.
	ExamCode int32 `json:"ExamCode,omitempty"`
	// ExamCodePS holds the value of the "ExamCodePS" field.
	ExamCodePS int32 `json:"ExamCodePS,omitempty"`
	// OTPRemarks holds the value of the "OTPRemarks" field.
	OTPRemarks string `json:"OTPRemarks,omitempty"`
	// Status holds the value of the "Status" field.
	Status bool `json:"Status,omitempty"`
	// NewPasswordRequest holds the value of the "NewPasswordRequest" field.
	NewPasswordRequest bool `json:"NewPasswordRequest,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// OTPTriggeredTime holds the value of the "OTPTriggeredTime" field.
	OTPTriggeredTime time.Time `json:"OTPTriggeredTime,omitempty"`
	// CreatedBy holds the value of the "CreatedBy" field.
	CreatedBy string `json:"CreatedBy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserMasterQuery when eager-loading is set.
	Edges                             UserMasterEdges `json:"edges"`
	employee_master_usermaster_ref    *int64
	exam_applications_ip_users_ip_ref *int64
	exam_applications_ps_users_ps_ref *int64
	exam_ip_users_ip_type             *int32
	exam_pa_users_ps_type             *int32
	exam_ps_users_ps_type             *int32
	selectValues                      sql.SelectValues
}

// UserMasterEdges holds the relations/edges for other nodes in the graph.
type UserMasterEdges struct {
	// Roles holds the value of the roles edge.
	Roles *RoleMaster `json:"roles,omitempty"`
	// UsermasterRef holds the value of the UsermasterRef edge.
	UsermasterRef []*EmployeeMaster `json:"UsermasterRef,omitempty"`
	// UsersPSRef holds the value of the UsersPSRef edge.
	UsersPSRef []*Exam_Applications_PS `json:"UsersPSRef,omitempty"`
	// UsersIPRef holds the value of the UsersIPRef edge.
	UsersIPRef []*Exam_Applications_IP `json:"UsersIPRef,omitempty"`
	// UsersPsType holds the value of the users_ps_type edge.
	UsersPsType []*Exam_PS `json:"users_ps_type,omitempty"`
	// UsersIPType holds the value of the users_ip_type edge.
	UsersIPType []*Exam_IP `json:"users_ip_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserMasterEdges) RolesOrErr() (*RoleMaster, error) {
	if e.loadedTypes[0] {
		if e.Roles == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: rolemaster.Label}
		}
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// UsermasterRefOrErr returns the UsermasterRef value or an error if the edge
// was not loaded in eager-loading.
func (e UserMasterEdges) UsermasterRefOrErr() ([]*EmployeeMaster, error) {
	if e.loadedTypes[1] {
		return e.UsermasterRef, nil
	}
	return nil, &NotLoadedError{edge: "UsermasterRef"}
}

// UsersPSRefOrErr returns the UsersPSRef value or an error if the edge
// was not loaded in eager-loading.
func (e UserMasterEdges) UsersPSRefOrErr() ([]*Exam_Applications_PS, error) {
	if e.loadedTypes[2] {
		return e.UsersPSRef, nil
	}
	return nil, &NotLoadedError{edge: "UsersPSRef"}
}

// UsersIPRefOrErr returns the UsersIPRef value or an error if the edge
// was not loaded in eager-loading.
func (e UserMasterEdges) UsersIPRefOrErr() ([]*Exam_Applications_IP, error) {
	if e.loadedTypes[3] {
		return e.UsersIPRef, nil
	}
	return nil, &NotLoadedError{edge: "UsersIPRef"}
}

// UsersPsTypeOrErr returns the UsersPsType value or an error if the edge
// was not loaded in eager-loading.
func (e UserMasterEdges) UsersPsTypeOrErr() ([]*Exam_PS, error) {
	if e.loadedTypes[4] {
		return e.UsersPsType, nil
	}
	return nil, &NotLoadedError{edge: "users_ps_type"}
}

// UsersIPTypeOrErr returns the UsersIPType value or an error if the edge
// was not loaded in eager-loading.
func (e UserMasterEdges) UsersIPTypeOrErr() ([]*Exam_IP, error) {
	if e.loadedTypes[5] {
		return e.UsersIPType, nil
	}
	return nil, &NotLoadedError{edge: "users_ip_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserMaster) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usermaster.FieldStatus, usermaster.FieldNewPasswordRequest:
			values[i] = new(sql.NullBool)
		case usermaster.FieldID, usermaster.FieldEmployeeID, usermaster.FieldRoleUserCode, usermaster.FieldOTP, usermaster.FieldExamCode, usermaster.FieldExamCodePS:
			values[i] = new(sql.NullInt64)
		case usermaster.FieldEmployeeName, usermaster.FieldFacilityID, usermaster.FieldCadre, usermaster.FieldMobile, usermaster.FieldEmailID, usermaster.FieldUserName, usermaster.FieldPassword, usermaster.FieldOTPRemarks, usermaster.FieldCreatedBy:
			values[i] = new(sql.NullString)
		case usermaster.FieldCreatedAt, usermaster.FieldOTPTriggeredTime:
			values[i] = new(sql.NullTime)
		case usermaster.ForeignKeys[0]: // employee_master_usermaster_ref
			values[i] = new(sql.NullInt64)
		case usermaster.ForeignKeys[1]: // exam_applications_ip_users_ip_ref
			values[i] = new(sql.NullInt64)
		case usermaster.ForeignKeys[2]: // exam_applications_ps_users_ps_ref
			values[i] = new(sql.NullInt64)
		case usermaster.ForeignKeys[3]: // exam_ip_users_ip_type
			values[i] = new(sql.NullInt64)
		case usermaster.ForeignKeys[4]: // exam_pa_users_ps_type
			values[i] = new(sql.NullInt64)
		case usermaster.ForeignKeys[5]: // exam_ps_users_ps_type
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserMaster fields.
func (um *UserMaster) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usermaster.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			um.ID = int64(value.Int64)
		case usermaster.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				um.EmployeeID = value.Int64
			}
		case usermaster.FieldEmployeeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeName", values[i])
			} else if value.Valid {
				um.EmployeeName = value.String
			}
		case usermaster.FieldFacilityID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FacilityID", values[i])
			} else if value.Valid {
				um.FacilityID = value.String
			}
		case usermaster.FieldCadre:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Cadre", values[i])
			} else if value.Valid {
				um.Cadre = value.String
			}
		case usermaster.FieldRoleUserCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RoleUserCode", values[i])
			} else if value.Valid {
				um.RoleUserCode = int32(value.Int64)
			}
		case usermaster.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Mobile", values[i])
			} else if value.Valid {
				um.Mobile = value.String
			}
		case usermaster.FieldEmailID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmailID", values[i])
			} else if value.Valid {
				um.EmailID = value.String
			}
		case usermaster.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserName", values[i])
			} else if value.Valid {
				um.UserName = value.String
			}
		case usermaster.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Password", values[i])
			} else if value.Valid {
				um.Password = value.String
			}
		case usermaster.FieldOTP:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field OTP", values[i])
			} else if value.Valid {
				um.OTP = int32(value.Int64)
			}
		case usermaster.FieldExamCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamCode", values[i])
			} else if value.Valid {
				um.ExamCode = int32(value.Int64)
			}
		case usermaster.FieldExamCodePS:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ExamCodePS", values[i])
			} else if value.Valid {
				um.ExamCodePS = int32(value.Int64)
			}
		case usermaster.FieldOTPRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OTPRemarks", values[i])
			} else if value.Valid {
				um.OTPRemarks = value.String
			}
		case usermaster.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				um.Status = value.Bool
			}
		case usermaster.FieldNewPasswordRequest:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field NewPasswordRequest", values[i])
			} else if value.Valid {
				um.NewPasswordRequest = value.Bool
			}
		case usermaster.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				um.CreatedAt = value.Time
			}
		case usermaster.FieldOTPTriggeredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field OTPTriggeredTime", values[i])
			} else if value.Valid {
				um.OTPTriggeredTime = value.Time
			}
		case usermaster.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedBy", values[i])
			} else if value.Valid {
				um.CreatedBy = value.String
			}
		case usermaster.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employee_master_usermaster_ref", value)
			} else if value.Valid {
				um.employee_master_usermaster_ref = new(int64)
				*um.employee_master_usermaster_ref = int64(value.Int64)
			}
		case usermaster.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_ip_users_ip_ref", value)
			} else if value.Valid {
				um.exam_applications_ip_users_ip_ref = new(int64)
				*um.exam_applications_ip_users_ip_ref = int64(value.Int64)
			}
		case usermaster.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_ps_users_ps_ref", value)
			} else if value.Valid {
				um.exam_applications_ps_users_ps_ref = new(int64)
				*um.exam_applications_ps_users_ps_ref = int64(value.Int64)
			}
		case usermaster.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_ip_users_ip_type", value)
			} else if value.Valid {
				um.exam_ip_users_ip_type = new(int32)
				*um.exam_ip_users_ip_type = int32(value.Int64)
			}
		case usermaster.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_pa_users_ps_type", value)
			} else if value.Valid {
				um.exam_pa_users_ps_type = new(int32)
				*um.exam_pa_users_ps_type = int32(value.Int64)
			}
		case usermaster.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_ps_users_ps_type", value)
			} else if value.Valid {
				um.exam_ps_users_ps_type = new(int32)
				*um.exam_ps_users_ps_type = int32(value.Int64)
			}
		default:
			um.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserMaster.
// This includes values selected through modifiers, order, etc.
func (um *UserMaster) Value(name string) (ent.Value, error) {
	return um.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the UserMaster entity.
func (um *UserMaster) QueryRoles() *RoleMasterQuery {
	return NewUserMasterClient(um.config).QueryRoles(um)
}

// QueryUsermasterRef queries the "UsermasterRef" edge of the UserMaster entity.
func (um *UserMaster) QueryUsermasterRef() *EmployeeMasterQuery {
	return NewUserMasterClient(um.config).QueryUsermasterRef(um)
}

// QueryUsersPSRef queries the "UsersPSRef" edge of the UserMaster entity.
func (um *UserMaster) QueryUsersPSRef() *ExamApplicationsPSQuery {
	return NewUserMasterClient(um.config).QueryUsersPSRef(um)
}

// QueryUsersIPRef queries the "UsersIPRef" edge of the UserMaster entity.
func (um *UserMaster) QueryUsersIPRef() *ExamApplicationsIPQuery {
	return NewUserMasterClient(um.config).QueryUsersIPRef(um)
}

// QueryUsersPsType queries the "users_ps_type" edge of the UserMaster entity.
func (um *UserMaster) QueryUsersPsType() *ExamPSQuery {
	return NewUserMasterClient(um.config).QueryUsersPsType(um)
}

// QueryUsersIPType queries the "users_ip_type" edge of the UserMaster entity.
func (um *UserMaster) QueryUsersIPType() *ExamIPQuery {
	return NewUserMasterClient(um.config).QueryUsersIPType(um)
}

// Update returns a builder for updating this UserMaster.
// Note that you need to call UserMaster.Unwrap() before calling this method if this UserMaster
// was returned from a transaction, and the transaction was committed or rolled back.
func (um *UserMaster) Update() *UserMasterUpdateOne {
	return NewUserMasterClient(um.config).UpdateOne(um)
}

// Unwrap unwraps the UserMaster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (um *UserMaster) Unwrap() *UserMaster {
	_tx, ok := um.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserMaster is not a transactional entity")
	}
	um.config.driver = _tx.drv
	return um
}

// String implements the fmt.Stringer.
func (um *UserMaster) String() string {
	var builder strings.Builder
	builder.WriteString("UserMaster(")
	builder.WriteString(fmt.Sprintf("id=%v, ", um.ID))
	builder.WriteString("EmployeeID=")
	builder.WriteString(fmt.Sprintf("%v", um.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("EmployeeName=")
	builder.WriteString(um.EmployeeName)
	builder.WriteString(", ")
	builder.WriteString("FacilityID=")
	builder.WriteString(um.FacilityID)
	builder.WriteString(", ")
	builder.WriteString("Cadre=")
	builder.WriteString(um.Cadre)
	builder.WriteString(", ")
	builder.WriteString("RoleUserCode=")
	builder.WriteString(fmt.Sprintf("%v", um.RoleUserCode))
	builder.WriteString(", ")
	builder.WriteString("Mobile=")
	builder.WriteString(um.Mobile)
	builder.WriteString(", ")
	builder.WriteString("EmailID=")
	builder.WriteString(um.EmailID)
	builder.WriteString(", ")
	builder.WriteString("UserName=")
	builder.WriteString(um.UserName)
	builder.WriteString(", ")
	builder.WriteString("Password=")
	builder.WriteString(um.Password)
	builder.WriteString(", ")
	builder.WriteString("OTP=")
	builder.WriteString(fmt.Sprintf("%v", um.OTP))
	builder.WriteString(", ")
	builder.WriteString("ExamCode=")
	builder.WriteString(fmt.Sprintf("%v", um.ExamCode))
	builder.WriteString(", ")
	builder.WriteString("ExamCodePS=")
	builder.WriteString(fmt.Sprintf("%v", um.ExamCodePS))
	builder.WriteString(", ")
	builder.WriteString("OTPRemarks=")
	builder.WriteString(um.OTPRemarks)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(fmt.Sprintf("%v", um.Status))
	builder.WriteString(", ")
	builder.WriteString("NewPasswordRequest=")
	builder.WriteString(fmt.Sprintf("%v", um.NewPasswordRequest))
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(um.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("OTPTriggeredTime=")
	builder.WriteString(um.OTPTriggeredTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("CreatedBy=")
	builder.WriteString(um.CreatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// UserMasters is a parsable slice of UserMaster.
type UserMasters []*UserMaster
