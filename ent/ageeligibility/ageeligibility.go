// Code generated by ent, DO NOT EDIT.

package ageeligibility

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the ageeligibility type in the database.
	Label = "age_eligibility"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "AgeElibilityCode"
	// FieldEligibilityCode holds the string denoting the eligibilitycode field in the database.
	FieldEligibilityCode = "eligibility_code"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldCategoryID holds the string denoting the categoryid field in the database.
	FieldCategoryID = "category_id"
	// EdgeExamEligibility holds the string denoting the exam_eligibility edge name in mutations.
	EdgeExamEligibility = "exam_eligibility"
	// ExamEligibilityFieldID holds the string denoting the ID field of the ExamEligibility.
	ExamEligibilityFieldID = "EligibilityCode"
	// Table holds the table name of the ageeligibility in the database.
	Table = "AgeEligibility"
	// ExamEligibilityTable is the table that holds the exam_eligibility relation/edge.
	ExamEligibilityTable = "AgeEligibility"
	// ExamEligibilityInverseTable is the table name for the ExamEligibility entity.
	// It exists in this package in order to avoid circular dependency with the "exameligibility" package.
	ExamEligibilityInverseTable = "ExamEligibility"
	// ExamEligibilityColumn is the table column denoting the exam_eligibility relation/edge.
	ExamEligibilityColumn = "eligibility_code"
)

// Columns holds all SQL columns for ageeligibility fields.
var Columns = []string{
	FieldID,
	FieldEligibilityCode,
	FieldAge,
	FieldCategoryID,
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

// OrderOption defines the ordering options for the AgeEligibility queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEligibilityCode orders the results by the EligibilityCode field.
func ByEligibilityCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligibilityCode, opts...).ToFunc()
}

// ByAge orders the results by the Age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByCategoryID orders the results by the CategoryID field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByExamEligibilityField orders the results by exam_eligibility field.
func ByExamEligibilityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamEligibilityStep(), sql.OrderByField(field, opts...))
	}
}
func newExamEligibilityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamEligibilityInverseTable, ExamEligibilityFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExamEligibilityTable, ExamEligibilityColumn),
	)
}