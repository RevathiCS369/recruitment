// Code generated by ent, DO NOT EDIT.

package rolemaster

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the rolemaster type in the database.
	Label = "role_master"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "RoleUserCode"
	// FieldRoleName holds the string denoting the rolename field in the database.
	FieldRoleName = "role_name"
	// FieldCreatedDate holds the string denoting the createddate field in the database.
	FieldCreatedDate = "created_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeRolesRef holds the string denoting the roles_ref edge name in mutations.
	EdgeRolesRef = "Roles_Ref"
	// EdgeRolesPSRef holds the string denoting the roles_ps_ref edge name in mutations.
	EdgeRolesPSRef = "Roles_PS_Ref"
	// EdgeRolesIPRef holds the string denoting the roles_ip_ref edge name in mutations.
	EdgeRolesIPRef = "Roles_IP_Ref"
	// AdminLoginFieldID holds the string denoting the ID field of the AdminLogin.
	AdminLoginFieldID = "LoginId"
	// UserMasterFieldID holds the string denoting the ID field of the UserMaster.
	UserMasterFieldID = "UserID"
	// Exam_Applications_PSFieldID holds the string denoting the ID field of the Exam_Applications_PS.
	Exam_Applications_PSFieldID = "ApplicationID"
	// Exam_Applications_IPFieldID holds the string denoting the ID field of the Exam_Applications_IP.
	Exam_Applications_IPFieldID = "ApplicationID"
	// Table holds the table name of the rolemaster in the database.
	Table = "RoleMaster"
	// RolesTable is the table that holds the roles relation/edge.
	RolesTable = "AdminLogin"
	// RolesInverseTable is the table name for the AdminLogin entity.
	// It exists in this package in order to avoid circular dependency with the "adminlogin" package.
	RolesInverseTable = "AdminLogin"
	// RolesColumn is the table column denoting the roles relation/edge.
	RolesColumn = "role_user_code"
	// RolesRefTable is the table that holds the Roles_Ref relation/edge.
	RolesRefTable = "UserMaster"
	// RolesRefInverseTable is the table name for the UserMaster entity.
	// It exists in this package in order to avoid circular dependency with the "usermaster" package.
	RolesRefInverseTable = "UserMaster"
	// RolesRefColumn is the table column denoting the Roles_Ref relation/edge.
	RolesRefColumn = "role_user_code"
	// RolesPSRefTable is the table that holds the Roles_PS_Ref relation/edge.
	RolesPSRefTable = "Exam_Applications_PS"
	// RolesPSRefInverseTable is the table name for the Exam_Applications_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ps" package.
	RolesPSRefInverseTable = "Exam_Applications_PS"
	// RolesPSRefColumn is the table column denoting the Roles_PS_Ref relation/edge.
	RolesPSRefColumn = "role_user_code"
	// RolesIPRefTable is the table that holds the Roles_IP_Ref relation/edge.
	RolesIPRefTable = "Exam_Applications_IP"
	// RolesIPRefInverseTable is the table name for the Exam_Applications_IP entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ip" package.
	RolesIPRefInverseTable = "Exam_Applications_IP"
	// RolesIPRefColumn is the table column denoting the Roles_IP_Ref relation/edge.
	RolesIPRefColumn = "role_user_code"
)

// Columns holds all SQL columns for rolemaster fields.
var Columns = []string{
	FieldID,
	FieldRoleName,
	FieldCreatedDate,
	FieldStatus,
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
	// DefaultStatus holds the default value on creation for the "Status" field.
	DefaultStatus bool
)

// OrderOption defines the ordering options for the RoleMaster queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoleName orders the results by the RoleName field.
func ByRoleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleName, opts...).ToFunc()
}

// ByCreatedDate orders the results by the CreatedDate field.
func ByCreatedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedDate, opts...).ToFunc()
}

// ByStatus orders the results by the Status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRolesRefCount orders the results by Roles_Ref count.
func ByRolesRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesRefStep(), opts...)
	}
}

// ByRolesRef orders the results by Roles_Ref terms.
func ByRolesRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRolesPSRefCount orders the results by Roles_PS_Ref count.
func ByRolesPSRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesPSRefStep(), opts...)
	}
}

// ByRolesPSRef orders the results by Roles_PS_Ref terms.
func ByRolesPSRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesPSRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRolesIPRefCount orders the results by Roles_IP_Ref count.
func ByRolesIPRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesIPRefStep(), opts...)
	}
}

// ByRolesIPRef orders the results by Roles_IP_Ref terms.
func ByRolesIPRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesIPRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, AdminLoginFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
	)
}
func newRolesRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesRefInverseTable, UserMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RolesRefTable, RolesRefColumn),
	)
}
func newRolesPSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesPSRefInverseTable, Exam_Applications_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RolesPSRefTable, RolesPSRefColumn),
	)
}
func newRolesIPRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesIPRefInverseTable, Exam_Applications_IPFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RolesIPRefTable, RolesIPRefColumn),
	)
}
