// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/application"
	"recruit/ent/center"
	"recruit/ent/notification"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Application is the model entity for the Application schema.
type Application struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// EmployeeID holds the value of the "EmployeeID" field.
	EmployeeID int32 `json:"EmployeeID,omitempty"`
	// NotifyCode holds the value of the "NotifyCode" field.
	NotifyCode int32 `json:"NotifyCode,omitempty"`
	// HallTicketNumber holds the value of the "HallTicketNumber" field.
	HallTicketNumber string `json:"HallTicketNumber,omitempty"`
	// CenterCode holds the value of the "CenterCode" field.
	CenterCode int32 `json:"CenterCode,omitempty"`
	// AppliedStamp holds the value of the "AppliedStamp" field.
	AppliedStamp time.Time `json:"AppliedStamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationQuery when eager-loading is set.
	Edges        ApplicationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ApplicationEdges holds the relations/edges for other nodes in the graph.
type ApplicationEdges struct {
	// Center holds the value of the center edge.
	Center *Center `json:"center,omitempty"`
	// Notification holds the value of the notification edge.
	Notification *Notification `json:"notification,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CenterOrErr returns the Center value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) CenterOrErr() (*Center, error) {
	if e.loadedTypes[0] {
		if e.Center == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: center.Label}
		}
		return e.Center, nil
	}
	return nil, &NotLoadedError{edge: "center"}
}

// NotificationOrErr returns the Notification value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) NotificationOrErr() (*Notification, error) {
	if e.loadedTypes[1] {
		if e.Notification == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: notification.Label}
		}
		return e.Notification, nil
	}
	return nil, &NotLoadedError{edge: "notification"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldID, application.FieldEmployeeID, application.FieldNotifyCode, application.FieldCenterCode:
			values[i] = new(sql.NullInt64)
		case application.FieldHallTicketNumber:
			values[i] = new(sql.NullString)
		case application.FieldAppliedStamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int32(value.Int64)
		case application.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeeID", values[i])
			} else if value.Valid {
				a.EmployeeID = int32(value.Int64)
			}
		case application.FieldNotifyCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field NotifyCode", values[i])
			} else if value.Valid {
				a.NotifyCode = int32(value.Int64)
			}
		case application.FieldHallTicketNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field HallTicketNumber", values[i])
			} else if value.Valid {
				a.HallTicketNumber = value.String
			}
		case application.FieldCenterCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CenterCode", values[i])
			} else if value.Valid {
				a.CenterCode = int32(value.Int64)
			}
		case application.FieldAppliedStamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field AppliedStamp", values[i])
			} else if value.Valid {
				a.AppliedStamp = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Application.
// This includes values selected through modifiers, order, etc.
func (a *Application) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryCenter queries the "center" edge of the Application entity.
func (a *Application) QueryCenter() *CenterQuery {
	return NewApplicationClient(a.config).QueryCenter(a)
}

// QueryNotification queries the "notification" edge of the Application entity.
func (a *Application) QueryNotification() *NotificationQuery {
	return NewApplicationClient(a.config).QueryNotification(a)
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return NewApplicationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("EmployeeID=")
	builder.WriteString(fmt.Sprintf("%v", a.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("NotifyCode=")
	builder.WriteString(fmt.Sprintf("%v", a.NotifyCode))
	builder.WriteString(", ")
	builder.WriteString("HallTicketNumber=")
	builder.WriteString(a.HallTicketNumber)
	builder.WriteString(", ")
	builder.WriteString("CenterCode=")
	builder.WriteString(fmt.Sprintf("%v", a.CenterCode))
	builder.WriteString(", ")
	builder.WriteString("AppliedStamp=")
	builder.WriteString(a.AppliedStamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application