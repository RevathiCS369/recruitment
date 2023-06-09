// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/employeedesignation"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EmployeeDesignation is the model entity for the EmployeeDesignation schema.
type EmployeeDesignation struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// DesignationCode holds the value of the "DesignationCode" field.
	DesignationCode string `json:"DesignationCode,omitempty"`
	// DesignationDescription holds the value of the "DesignationDescription" field.
	DesignationDescription string `json:"DesignationDescription,omitempty"`
	selectValues           sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmployeeDesignation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case employeedesignation.FieldID:
			values[i] = new(sql.NullInt64)
		case employeedesignation.FieldDesignationCode, employeedesignation.FieldDesignationDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmployeeDesignation fields.
func (ed *EmployeeDesignation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employeedesignation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ed.ID = int32(value.Int64)
		case employeedesignation.FieldDesignationCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DesignationCode", values[i])
			} else if value.Valid {
				ed.DesignationCode = value.String
			}
		case employeedesignation.FieldDesignationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DesignationDescription", values[i])
			} else if value.Valid {
				ed.DesignationDescription = value.String
			}
		default:
			ed.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmployeeDesignation.
// This includes values selected through modifiers, order, etc.
func (ed *EmployeeDesignation) Value(name string) (ent.Value, error) {
	return ed.selectValues.Get(name)
}

// Update returns a builder for updating this EmployeeDesignation.
// Note that you need to call EmployeeDesignation.Unwrap() before calling this method if this EmployeeDesignation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ed *EmployeeDesignation) Update() *EmployeeDesignationUpdateOne {
	return NewEmployeeDesignationClient(ed.config).UpdateOne(ed)
}

// Unwrap unwraps the EmployeeDesignation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ed *EmployeeDesignation) Unwrap() *EmployeeDesignation {
	_tx, ok := ed.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmployeeDesignation is not a transactional entity")
	}
	ed.config.driver = _tx.drv
	return ed
}

// String implements the fmt.Stringer.
func (ed *EmployeeDesignation) String() string {
	var builder strings.Builder
	builder.WriteString("EmployeeDesignation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ed.ID))
	builder.WriteString("DesignationCode=")
	builder.WriteString(ed.DesignationCode)
	builder.WriteString(", ")
	builder.WriteString("DesignationDescription=")
	builder.WriteString(ed.DesignationDescription)
	builder.WriteByte(')')
	return builder.String()
}

// EmployeeDesignations is a parsable slice of EmployeeDesignation.
type EmployeeDesignations []*EmployeeDesignation
