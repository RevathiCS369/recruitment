// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/cadre_choice_ip"
	"recruit/ent/exam_applications_ip"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Cadre_Choice_IP is the model entity for the Cadre_Choice_IP schema.
type Cadre_Choice_IP struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ApplicationID holds the value of the "ApplicationID" field.
	ApplicationID int64 `json:"ApplicationID,omitempty"`
	// CadrePrefNo holds the value of the "CadrePrefNo" field.
	CadrePrefNo string `json:"CadrePrefNo,omitempty"`
	// CadrePrefValue holds the value of the "CadrePrefValue" field.
	CadrePrefValue string `json:"CadrePrefValue,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int64 `json:"EmployeeID,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy string `json:"UpdatedBy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Cadre_Choice_IPQuery when eager-loading is set.
	Edges        Cadre_Choice_IPEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Cadre_Choice_IPEdges holds the relations/edges for other nodes in the graph.
type Cadre_Choice_IPEdges struct {
	// ApplnIPRef holds the value of the ApplnIP_Ref edge.
	ApplnIPRef *Exam_Applications_IP `json:"ApplnIP_Ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplnIPRefOrErr returns the ApplnIPRef value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Cadre_Choice_IPEdges) ApplnIPRefOrErr() (*Exam_Applications_IP, error) {
	if e.loadedTypes[0] {
		if e.ApplnIPRef == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: exam_applications_ip.Label}
		}
		return e.ApplnIPRef, nil
	}
	return nil, &NotLoadedError{edge: "ApplnIP_Ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cadre_Choice_IP) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cadre_choice_ip.FieldID, cadre_choice_ip.FieldApplicationID, cadre_choice_ip.FieldEmployeeID:
			values[i] = new(sql.NullInt64)
		case cadre_choice_ip.FieldCadrePrefNo, cadre_choice_ip.FieldCadrePrefValue, cadre_choice_ip.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case cadre_choice_ip.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cadre_Choice_IP fields.
func (cci *Cadre_Choice_IP) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cadre_choice_ip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cci.ID = int32(value.Int64)
		case cadre_choice_ip.FieldApplicationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ApplicationID", values[i])
			} else if value.Valid {
				cci.ApplicationID = value.Int64
			}
		case cadre_choice_ip.FieldCadrePrefNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CadrePrefNo", values[i])
			} else if value.Valid {
				cci.CadrePrefNo = value.String
			}
		case cadre_choice_ip.FieldCadrePrefValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CadrePrefValue", values[i])
			} else if value.Valid {
				cci.CadrePrefValue = value.String
			}
		case cadre_choice_ip.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				cci.EmployeeID = value.Int64
			}
		case cadre_choice_ip.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				cci.UpdatedAt = value.Time
			}
		case cadre_choice_ip.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				cci.UpdatedBy = value.String
			}
		default:
			cci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cadre_Choice_IP.
// This includes values selected through modifiers, order, etc.
func (cci *Cadre_Choice_IP) Value(name string) (ent.Value, error) {
	return cci.selectValues.Get(name)
}

// QueryApplnIPRef queries the "ApplnIP_Ref" edge of the Cadre_Choice_IP entity.
func (cci *Cadre_Choice_IP) QueryApplnIPRef() *ExamApplicationsIPQuery {
	return NewCadreChoiceIPClient(cci.config).QueryApplnIPRef(cci)
}

// Update returns a builder for updating this Cadre_Choice_IP.
// Note that you need to call Cadre_Choice_IP.Unwrap() before calling this method if this Cadre_Choice_IP
// was returned from a transaction, and the transaction was committed or rolled back.
func (cci *Cadre_Choice_IP) Update() *CadreChoiceIPUpdateOne {
	return NewCadreChoiceIPClient(cci.config).UpdateOne(cci)
}

// Unwrap unwraps the Cadre_Choice_IP entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cci *Cadre_Choice_IP) Unwrap() *Cadre_Choice_IP {
	_tx, ok := cci.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cadre_Choice_IP is not a transactional entity")
	}
	cci.config.driver = _tx.drv
	return cci
}

// String implements the fmt.Stringer.
func (cci *Cadre_Choice_IP) String() string {
	var builder strings.Builder
	builder.WriteString("Cadre_Choice_IP(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cci.ID))
	builder.WriteString("ApplicationID=")
	builder.WriteString(fmt.Sprintf("%v", cci.ApplicationID))
	builder.WriteString(", ")
	builder.WriteString("CadrePrefNo=")
	builder.WriteString(cci.CadrePrefNo)
	builder.WriteString(", ")
	builder.WriteString("CadrePrefValue=")
	builder.WriteString(cci.CadrePrefValue)
	builder.WriteString(", ")
	builder.WriteString("EmployeeID=")
	builder.WriteString(fmt.Sprintf("%v", cci.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(cci.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedBy=")
	builder.WriteString(cci.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Cadre_Choice_IPs is a parsable slice of Cadre_Choice_IP.
type Cadre_Choice_IPs []*Cadre_Choice_IP
