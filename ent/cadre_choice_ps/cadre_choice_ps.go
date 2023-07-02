// Code generated by ent, DO NOT EDIT.

package cadre_choice_ps

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cadre_choice_ps type in the database.
	Label = "cadre_choice_ps"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "CadrePrefId"
	// FieldApplication holds the string denoting the application field in the database.
	FieldApplication = "application"
	// FieldCadrePrefNo holds the string denoting the cadreprefno field in the database.
	FieldCadrePrefNo = "cadre_pref_no"
	// FieldCadrePrefValue holds the string denoting the cadreprefvalue field in the database.
	FieldCadrePrefValue = "cadre_pref_value"
	// FieldEmployeeID holds the string denoting the employeeid field in the database.
	FieldEmployeeID = "employee_id"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// Table holds the table name of the cadre_choice_ps in the database.
	Table = "Cadre_Choice_PS"
)

// Columns holds all SQL columns for cadre_choice_ps fields.
var Columns = []string{
	FieldID,
	FieldApplication,
	FieldCadrePrefNo,
	FieldCadrePrefValue,
	FieldEmployeeID,
	FieldUpdatedAt,
	FieldUpdatedBy,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// DefaultUpdatedBy holds the default value on creation for the "UpdatedBy" field.
	DefaultUpdatedBy string
)

// OrderOption defines the ordering options for the Cadre_Choice_PS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApplication orders the results by the Application field.
func ByApplication(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplication, opts...).ToFunc()
}

// ByCadrePrefNo orders the results by the CadrePrefNo field.
func ByCadrePrefNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCadrePrefNo, opts...).ToFunc()
}

// ByCadrePrefValue orders the results by the CadrePrefValue field.
func ByCadrePrefValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCadrePrefValue, opts...).ToFunc()
}

// ByEmployeeID orders the results by the EmployeeID field.
func ByEmployeeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeID, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}