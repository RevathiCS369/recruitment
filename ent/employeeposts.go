// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/employeeposts"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EmployeePosts is the model entity for the EmployeePosts schema.
type EmployeePosts struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// PostCode holds the value of the "PostCode" field.
	PostCode string `json:"PostCode,omitempty"`
	// PostDescription holds the value of the "PostDescription" field.
	PostDescription string `json:"PostDescription,omitempty"`
	// Group holds the value of the "Group" field.
	Group string `json:"Group,omitempty"`
	// PayLevel holds the value of the "PayLevel" field.
	PayLevel string `json:"PayLevel,omitempty"`
	// Scale holds the value of the "Scale" field.
	Scale string `json:"Scale,omitempty"`
	// BaseCadreFlag holds the value of the "BaseCadreFlag" field.
	BaseCadreFlag bool `json:"BaseCadreFlag,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmployeePosts) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employeeposts.FieldBaseCadreFlag:
			values[i] = new(sql.NullBool)
		case employeeposts.FieldID:
			values[i] = new(sql.NullInt64)
		case employeeposts.FieldPostCode, employeeposts.FieldPostDescription, employeeposts.FieldGroup, employeeposts.FieldPayLevel, employeeposts.FieldScale:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmployeePosts fields.
func (ep *EmployeePosts) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employeeposts.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ep.ID = int32(value.Int64)
		case employeeposts.FieldPostCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostCode", values[i])
			} else if value.Valid {
				ep.PostCode = value.String
			}
		case employeeposts.FieldPostDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostDescription", values[i])
			} else if value.Valid {
				ep.PostDescription = value.String
			}
		case employeeposts.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Group", values[i])
			} else if value.Valid {
				ep.Group = value.String
			}
		case employeeposts.FieldPayLevel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PayLevel", values[i])
			} else if value.Valid {
				ep.PayLevel = value.String
			}
		case employeeposts.FieldScale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Scale", values[i])
			} else if value.Valid {
				ep.Scale = value.String
			}
		case employeeposts.FieldBaseCadreFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field BaseCadreFlag", values[i])
			} else if value.Valid {
				ep.BaseCadreFlag = value.Bool
			}
		default:
			ep.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmployeePosts.
// This includes values selected through modifiers, order, etc.
func (ep *EmployeePosts) Value(name string) (ent.Value, error) {
	return ep.selectValues.Get(name)
}

// Update returns a builder for updating this EmployeePosts.
// Note that you need to call EmployeePosts.Unwrap() before calling this method if this EmployeePosts
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EmployeePosts) Update() *EmployeePostsUpdateOne {
	return NewEmployeePostsClient(ep.config).UpdateOne(ep)
}

// Unwrap unwraps the EmployeePosts entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *EmployeePosts) Unwrap() *EmployeePosts {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmployeePosts is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EmployeePosts) String() string {
	var builder strings.Builder
	builder.WriteString("EmployeePosts(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("PostCode=")
	builder.WriteString(ep.PostCode)
	builder.WriteString(", ")
	builder.WriteString("PostDescription=")
	builder.WriteString(ep.PostDescription)
	builder.WriteString(", ")
	builder.WriteString("Group=")
	builder.WriteString(ep.Group)
	builder.WriteString(", ")
	builder.WriteString("PayLevel=")
	builder.WriteString(ep.PayLevel)
	builder.WriteString(", ")
	builder.WriteString("Scale=")
	builder.WriteString(ep.Scale)
	builder.WriteString(", ")
	builder.WriteString("BaseCadreFlag=")
	builder.WriteString(fmt.Sprintf("%v", ep.BaseCadreFlag))
	builder.WriteByte(')')
	return builder.String()
}

// EmployeePostsSlice is a parsable slice of EmployeePosts.
type EmployeePostsSlice []*EmployeePosts
