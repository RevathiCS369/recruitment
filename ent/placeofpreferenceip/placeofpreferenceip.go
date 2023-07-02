// Code generated by ent, DO NOT EDIT.

package placeofpreferenceip

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the placeofpreferenceip type in the database.
	Label = "place_of_preference_ip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "PlacePrefId"
	// FieldApplicationID holds the string denoting the applicationid field in the database.
	FieldApplicationID = "application_id"
	// FieldPlacePrefNo holds the string denoting the placeprefno field in the database.
	FieldPlacePrefNo = "place_pref_no"
	// FieldPlacePrefValue holds the string denoting the placeprefvalue field in the database.
	FieldPlacePrefValue = "place_pref_value"
	// FieldEmployeeID holds the string denoting the employeeid field in the database.
	FieldEmployeeID = "employee_id"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// EdgeApplnIPRef holds the string denoting the applnip_ref edge name in mutations.
	EdgeApplnIPRef = "ApplnIP_Ref"
	// Exam_Applications_IPFieldID holds the string denoting the ID field of the Exam_Applications_IP.
	Exam_Applications_IPFieldID = "ApplicationID"
	// Table holds the table name of the placeofpreferenceip in the database.
	Table = "PlaceOfPreferenceIP"
	// ApplnIPRefTable is the table that holds the ApplnIP_Ref relation/edge.
	ApplnIPRefTable = "PlaceOfPreferenceIP"
	// ApplnIPRefInverseTable is the table name for the Exam_Applications_IP entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ip" package.
	ApplnIPRefInverseTable = "Exam_Applications_IP"
	// ApplnIPRefColumn is the table column denoting the ApplnIP_Ref relation/edge.
	ApplnIPRefColumn = "application_id"
)

// Columns holds all SQL columns for placeofpreferenceip fields.
var Columns = []string{
	FieldID,
	FieldApplicationID,
	FieldPlacePrefNo,
	FieldPlacePrefValue,
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
	// DefaultUpdatedBy holds the default value on creation for the "UpdatedBy" field.
	DefaultUpdatedBy string
)

// OrderOption defines the ordering options for the PlaceOfPreferenceIP queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApplicationID orders the results by the ApplicationID field.
func ByApplicationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationID, opts...).ToFunc()
}

// ByPlacePrefNo orders the results by the PlacePrefNo field.
func ByPlacePrefNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlacePrefNo, opts...).ToFunc()
}

// ByPlacePrefValue orders the results by the PlacePrefValue field.
func ByPlacePrefValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlacePrefValue, opts...).ToFunc()
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

// ByApplnIPRefField orders the results by ApplnIP_Ref field.
func ByApplnIPRefField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApplnIPRefStep(), sql.OrderByField(field, opts...))
	}
}
func newApplnIPRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApplnIPRefInverseTable, Exam_Applications_IPFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ApplnIPRefTable, ApplnIPRefColumn),
	)
}