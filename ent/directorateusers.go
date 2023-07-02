// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/directorateusers"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// DirectorateUsers is the model entity for the DirectorateUsers schema.
type DirectorateUsers struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Role holds the value of the "Role" field.
	Role string `json:"Role,omitempty"`
	// EmployeedID holds the value of the "EmployeedID" field.
	EmployeedID int32 `json:"EmployeedID,omitempty"`
	// EmployeeName holds the value of the "EmployeeName" field.
	EmployeeName string `json:"EmployeeName,omitempty"`
	// EmailId holds the value of the "EmailId" field.
	EmailId string `json:"EmailId,omitempty"`
	// MobileNumber holds the value of the "MobileNumber" field.
	MobileNumber int64 `json:"MobileNumber,omitempty"`
	// SequenceNumber holds the value of the "SequenceNumber" field.
	SequenceNumber int32 `json:"SequenceNumber,omitempty"`
	// Status holds the value of the "Status" field.
	Status string `json:"Status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DirectorateUsersQuery when eager-loading is set.
	Edges        DirectorateUsersEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DirectorateUsersEdges holds the relations/edges for other nodes in the graph.
type DirectorateUsersEdges struct {
	// EmployeeUser holds the value of the employee_user edge.
	EmployeeUser []*Employees `json:"employee_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeeUserOrErr returns the EmployeeUser value or an error if the edge
// was not loaded in eager-loading.
func (e DirectorateUsersEdges) EmployeeUserOrErr() ([]*Employees, error) {
	if e.loadedTypes[0] {
		return e.EmployeeUser, nil
	}
	return nil, &NotLoadedError{edge: "employee_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DirectorateUsers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case directorateusers.FieldID, directorateusers.FieldEmployeedID, directorateusers.FieldMobileNumber, directorateusers.FieldSequenceNumber:
			values[i] = new(sql.NullInt64)
		case directorateusers.FieldRole, directorateusers.FieldEmployeeName, directorateusers.FieldEmailId, directorateusers.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DirectorateUsers fields.
func (du *DirectorateUsers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case directorateusers.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			du.ID = int32(value.Int64)
		case directorateusers.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Role", values[i])
			} else if value.Valid {
				du.Role = value.String
			}
		case directorateusers.FieldEmployeedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeedID", values[i])
			} else if value.Valid {
				du.EmployeedID = int32(value.Int64)
			}
		case directorateusers.FieldEmployeeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeName", values[i])
			} else if value.Valid {
				du.EmployeeName = value.String
			}
		case directorateusers.FieldEmailId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmailId", values[i])
			} else if value.Valid {
				du.EmailId = value.String
			}
		case directorateusers.FieldMobileNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field MobileNumber", values[i])
			} else if value.Valid {
				du.MobileNumber = value.Int64
			}
		case directorateusers.FieldSequenceNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field SequenceNumber", values[i])
			} else if value.Valid {
				du.SequenceNumber = int32(value.Int64)
			}
		case directorateusers.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				du.Status = value.String
			}
		default:
			du.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DirectorateUsers.
// This includes values selected through modifiers, order, etc.
func (du *DirectorateUsers) Value(name string) (ent.Value, error) {
	return du.selectValues.Get(name)
}

// QueryEmployeeUser queries the "employee_user" edge of the DirectorateUsers entity.
func (du *DirectorateUsers) QueryEmployeeUser() *EmployeesQuery {
	return NewDirectorateUsersClient(du.config).QueryEmployeeUser(du)
}

// Update returns a builder for updating this DirectorateUsers.
// Note that you need to call DirectorateUsers.Unwrap() before calling this method if this DirectorateUsers
// was returned from a transaction, and the transaction was committed or rolled back.
func (du *DirectorateUsers) Update() *DirectorateUsersUpdateOne {
	return NewDirectorateUsersClient(du.config).UpdateOne(du)
}

// Unwrap unwraps the DirectorateUsers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (du *DirectorateUsers) Unwrap() *DirectorateUsers {
	_tx, ok := du.config.driver.(*txDriver)
	if !ok {
		panic("ent: DirectorateUsers is not a transactional entity")
	}
	du.config.driver = _tx.drv
	return du
}

// String implements the fmt.Stringer.
func (du *DirectorateUsers) String() string {
	var builder strings.Builder
	builder.WriteString("DirectorateUsers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", du.ID))
	builder.WriteString("Role=")
	builder.WriteString(du.Role)
	builder.WriteString(", ")
	builder.WriteString("EmployeedID=")
	builder.WriteString(fmt.Sprintf("%v", du.EmployeedID))
	builder.WriteString(", ")
	builder.WriteString("EmployeeName=")
	builder.WriteString(du.EmployeeName)
	builder.WriteString(", ")
	builder.WriteString("EmailId=")
	builder.WriteString(du.EmailId)
	builder.WriteString(", ")
	builder.WriteString("MobileNumber=")
	builder.WriteString(fmt.Sprintf("%v", du.MobileNumber))
	builder.WriteString(", ")
	builder.WriteString("SequenceNumber=")
	builder.WriteString(fmt.Sprintf("%v", du.SequenceNumber))
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(du.Status)
	builder.WriteByte(')')
	return builder.String()
}

// DirectorateUsersSlice is a parsable slice of DirectorateUsers.
type DirectorateUsersSlice []*DirectorateUsers