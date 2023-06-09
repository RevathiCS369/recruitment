// Code generated by ent, DO NOT EDIT.

package employeeposts

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the employeeposts type in the database.
	Label = "employee_posts"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "PostID"
	// FieldPostCode holds the string denoting the postcode field in the database.
	FieldPostCode = "post_code"
	// FieldPostDescription holds the string denoting the postdescription field in the database.
	FieldPostDescription = "post_description"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldPayLevel holds the string denoting the paylevel field in the database.
	FieldPayLevel = "pay_level"
	// FieldScale holds the string denoting the scale field in the database.
	FieldScale = "scale"
	// FieldBaseCadreFlag holds the string denoting the basecadreflag field in the database.
	FieldBaseCadreFlag = "base_cadre_flag"
	// EdgeEmpPosts holds the string denoting the emp_posts edge name in mutations.
	EdgeEmpPosts = "emp_posts"
	// EdgePostEligibility holds the string denoting the posteligibility edge name in mutations.
	EdgePostEligibility = "PostEligibility"
	// EmployeesFieldID holds the string denoting the ID field of the Employees.
	EmployeesFieldID = "RegistrationID"
	// EligibilityMasterFieldID holds the string denoting the ID field of the EligibilityMaster.
	EligibilityMasterFieldID = "EligibilityCode"
	// Table holds the table name of the employeeposts in the database.
	Table = "EmployeePosts"
	// EmpPostsTable is the table that holds the emp_posts relation/edge.
	EmpPostsTable = "Employees"
	// EmpPostsInverseTable is the table name for the Employees entity.
	// It exists in this package in order to avoid circular dependency with the "employees" package.
	EmpPostsInverseTable = "Employees"
	// EmpPostsColumn is the table column denoting the emp_posts relation/edge.
	EmpPostsColumn = "employee_posts_emp_posts"
	// PostEligibilityTable is the table that holds the PostEligibility relation/edge.
	PostEligibilityTable = "EligibilityMaster"
	// PostEligibilityInverseTable is the table name for the EligibilityMaster entity.
	// It exists in this package in order to avoid circular dependency with the "eligibilitymaster" package.
	PostEligibilityInverseTable = "EligibilityMaster"
	// PostEligibilityColumn is the table column denoting the PostEligibility relation/edge.
	PostEligibilityColumn = "employee_posts_post_eligibility"
)

// Columns holds all SQL columns for employeeposts fields.
var Columns = []string{
	FieldID,
	FieldPostCode,
	FieldPostDescription,
	FieldGroup,
	FieldPayLevel,
	FieldScale,
	FieldBaseCadreFlag,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "EmployeePosts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"eligibility_master_post_eligibility",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the EmployeePosts queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPostCode orders the results by the PostCode field.
func ByPostCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostCode, opts...).ToFunc()
}

// ByPostDescription orders the results by the PostDescription field.
func ByPostDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostDescription, opts...).ToFunc()
}

// ByGroup orders the results by the Group field.
func ByGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroup, opts...).ToFunc()
}

// ByPayLevel orders the results by the PayLevel field.
func ByPayLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayLevel, opts...).ToFunc()
}

// ByScale orders the results by the Scale field.
func ByScale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScale, opts...).ToFunc()
}

// ByBaseCadreFlag orders the results by the BaseCadreFlag field.
func ByBaseCadreFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBaseCadreFlag, opts...).ToFunc()
}

// ByEmpPostsCount orders the results by emp_posts count.
func ByEmpPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmpPostsStep(), opts...)
	}
}

// ByEmpPosts orders the results by emp_posts terms.
func ByEmpPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmpPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostEligibilityCount orders the results by PostEligibility count.
func ByPostEligibilityCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostEligibilityStep(), opts...)
	}
}

// ByPostEligibility orders the results by PostEligibility terms.
func ByPostEligibility(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostEligibilityStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEmpPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmpPostsInverseTable, EmployeesFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmpPostsTable, EmpPostsColumn),
	)
}
func newPostEligibilityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostEligibilityInverseTable, EligibilityMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostEligibilityTable, PostEligibilityColumn),
	)
}
