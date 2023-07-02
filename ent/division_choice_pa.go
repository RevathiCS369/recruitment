// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/division_choice_pa"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Division_Choice_PA is the model entity for the Division_Choice_PA schema.
type Division_Choice_PA struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Application holds the value of the "Application" field.
	Application string `json:"Application,omitempty"`
	// CadrePrefNo holds the value of the "CadrePrefNo" field.
	CadrePrefNo string `json:"CadrePrefNo,omitempty"`
	// CadrePrefValue holds the value of the "CadrePrefValue" field.
	CadrePrefValue string `json:"CadrePrefValue,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int64 `json:"EmployeeID,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy    string `json:"UpdatedBy,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Division_Choice_PA) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case division_choice_pa.FieldID, division_choice_pa.FieldEmployeeID:
			values[i] = new(sql.NullInt64)
		case division_choice_pa.FieldApplication, division_choice_pa.FieldCadrePrefNo, division_choice_pa.FieldCadrePrefValue, division_choice_pa.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case division_choice_pa.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Division_Choice_PA fields.
func (dcp *Division_Choice_PA) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case division_choice_pa.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dcp.ID = int32(value.Int64)
		case division_choice_pa.FieldApplication:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Application", values[i])
			} else if value.Valid {
				dcp.Application = value.String
			}
		case division_choice_pa.FieldCadrePrefNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CadrePrefNo", values[i])
			} else if value.Valid {
				dcp.CadrePrefNo = value.String
			}
		case division_choice_pa.FieldCadrePrefValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CadrePrefValue", values[i])
			} else if value.Valid {
				dcp.CadrePrefValue = value.String
			}
		case division_choice_pa.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				dcp.EmployeeID = value.Int64
			}
		case division_choice_pa.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				dcp.UpdatedAt = value.Time
			}
		case division_choice_pa.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				dcp.UpdatedBy = value.String
			}
		default:
			dcp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Division_Choice_PA.
// This includes values selected through modifiers, order, etc.
func (dcp *Division_Choice_PA) Value(name string) (ent.Value, error) {
	return dcp.selectValues.Get(name)
}

// Update returns a builder for updating this Division_Choice_PA.
// Note that you need to call Division_Choice_PA.Unwrap() before calling this method if this Division_Choice_PA
// was returned from a transaction, and the transaction was committed or rolled back.
func (dcp *Division_Choice_PA) Update() *DivisionChoicePAUpdateOne {
	return NewDivisionChoicePAClient(dcp.config).UpdateOne(dcp)
}

// Unwrap unwraps the Division_Choice_PA entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dcp *Division_Choice_PA) Unwrap() *Division_Choice_PA {
	_tx, ok := dcp.config.driver.(*txDriver)
	if !ok {
		panic("ent: Division_Choice_PA is not a transactional entity")
	}
	dcp.config.driver = _tx.drv
	return dcp
}

// String implements the fmt.Stringer.
func (dcp *Division_Choice_PA) String() string {
	var builder strings.Builder
	builder.WriteString("Division_Choice_PA(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dcp.ID))
	builder.WriteString("Application=")
	builder.WriteString(dcp.Application)
	builder.WriteString(", ")
	builder.WriteString("CadrePrefNo=")
	builder.WriteString(dcp.CadrePrefNo)
	builder.WriteString(", ")
	builder.WriteString("CadrePrefValue=")
	builder.WriteString(dcp.CadrePrefValue)
	builder.WriteString(", ")
	builder.WriteString("EmployeeID=")
	builder.WriteString(fmt.Sprintf("%v", dcp.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(dcp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedBy=")
	builder.WriteString(dcp.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Division_Choice_PAs is a parsable slice of Division_Choice_PA.
type Division_Choice_PAs []*Division_Choice_PA
