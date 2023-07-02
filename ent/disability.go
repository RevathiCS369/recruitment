// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/disability"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Disability is the model entity for the Disability schema.
type Disability struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// DisabilityTypeCode holds the value of the "DisabilityTypeCode" field.
	DisabilityTypeCode string `json:"DisabilityTypeCode,omitempty"`
	// DisabilityTypeDescription holds the value of the "DisabilityTypeDescription" field.
	DisabilityTypeDescription string `json:"DisabilityTypeDescription,omitempty"`
	// DisabilityPercentage holds the value of the "DisabilityPercentage" field.
	DisabilityPercentage int32 `json:"DisabilityPercentage,omitempty"`
	// DisabilityFlag holds the value of the "DisabilityFlag" field.
	DisabilityFlag disability.DisabilityFlag `json:"DisabilityFlag,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DisabilityQuery when eager-loading is set.
	Edges               DisabilityEdges `json:"edges"`
	exam_papers_dis_ref *int32
	selectValues        sql.SelectValues
}

// DisabilityEdges holds the relations/edges for other nodes in the graph.
type DisabilityEdges struct {
	// DisRef holds the value of the dis_ref edge.
	DisRef []*ExamPapers `json:"dis_ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DisRefOrErr returns the DisRef value or an error if the edge
// was not loaded in eager-loading.
func (e DisabilityEdges) DisRefOrErr() ([]*ExamPapers, error) {
	if e.loadedTypes[0] {
		return e.DisRef, nil
	}
	return nil, &NotLoadedError{edge: "dis_ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Disability) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case disability.FieldID, disability.FieldDisabilityPercentage:
			values[i] = new(sql.NullInt64)
		case disability.FieldDisabilityTypeCode, disability.FieldDisabilityTypeDescription, disability.FieldDisabilityFlag:
			values[i] = new(sql.NullString)
		case disability.ForeignKeys[0]: // exam_papers_dis_ref
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Disability fields.
func (d *Disability) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case disability.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int32(value.Int64)
		case disability.FieldDisabilityTypeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DisabilityTypeCode", values[i])
			} else if value.Valid {
				d.DisabilityTypeCode = value.String
			}
		case disability.FieldDisabilityTypeDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DisabilityTypeDescription", values[i])
			} else if value.Valid {
				d.DisabilityTypeDescription = value.String
			}
		case disability.FieldDisabilityPercentage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field DisabilityPercentage", values[i])
			} else if value.Valid {
				d.DisabilityPercentage = int32(value.Int64)
			}
		case disability.FieldDisabilityFlag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DisabilityFlag", values[i])
			} else if value.Valid {
				d.DisabilityFlag = disability.DisabilityFlag(value.String)
			}
		case disability.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_papers_dis_ref", value)
			} else if value.Valid {
				d.exam_papers_dis_ref = new(int32)
				*d.exam_papers_dis_ref = int32(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Disability.
// This includes values selected through modifiers, order, etc.
func (d *Disability) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryDisRef queries the "dis_ref" edge of the Disability entity.
func (d *Disability) QueryDisRef() *ExamPapersQuery {
	return NewDisabilityClient(d.config).QueryDisRef(d)
}

// Update returns a builder for updating this Disability.
// Note that you need to call Disability.Unwrap() before calling this method if this Disability
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Disability) Update() *DisabilityUpdateOne {
	return NewDisabilityClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Disability entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Disability) Unwrap() *Disability {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Disability is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Disability) String() string {
	var builder strings.Builder
	builder.WriteString("Disability(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("DisabilityTypeCode=")
	builder.WriteString(d.DisabilityTypeCode)
	builder.WriteString(", ")
	builder.WriteString("DisabilityTypeDescription=")
	builder.WriteString(d.DisabilityTypeDescription)
	builder.WriteString(", ")
	builder.WriteString("DisabilityPercentage=")
	builder.WriteString(fmt.Sprintf("%v", d.DisabilityPercentage))
	builder.WriteString(", ")
	builder.WriteString("DisabilityFlag=")
	builder.WriteString(fmt.Sprintf("%v", d.DisabilityFlag))
	builder.WriteByte(')')
	return builder.String()
}

// Disabilities is a parsable slice of Disability.
type Disabilities []*Disability
