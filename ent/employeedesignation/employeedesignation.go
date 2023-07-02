// Code generated by ent, DO NOT EDIT.

package employeedesignation

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the employeedesignation type in the database.
	Label = "employee_designation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "DesignationID"
	// FieldDesignationCode holds the string denoting the designationcode field in the database.
	FieldDesignationCode = "designation_code"
	// FieldDesignationDescription holds the string denoting the designationdescription field in the database.
	FieldDesignationDescription = "designation_description"
	// Table holds the table name of the employeedesignation in the database.
	Table = "EmployeeDesignation"
)

// Columns holds all SQL columns for employeedesignation fields.
var Columns = []string{
	FieldID,
	FieldDesignationCode,
	FieldDesignationDescription,
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

// OrderOption defines the ordering options for the EmployeeDesignation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDesignationCode orders the results by the DesignationCode field.
func ByDesignationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesignationCode, opts...).ToFunc()
}

// ByDesignationDescription orders the results by the DesignationDescription field.
func ByDesignationDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesignationDescription, opts...).ToFunc()
}