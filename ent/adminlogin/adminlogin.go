// Code generated by ent, DO NOT EDIT.

package adminlogin

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the adminlogin type in the database.
	Label = "admin_login"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "LoginId"
	// FieldRoleUserCode holds the string denoting the roleusercode field in the database.
	FieldRoleUserCode = "role_user_code"
	// FieldRoleName holds the string denoting the rolename field in the database.
	FieldRoleName = "role_name"
	// FieldCreatedDate holds the string denoting the createddate field in the database.
	FieldCreatedDate = "created_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldEmployeedID holds the string denoting the employeedid field in the database.
	FieldEmployeedID = "employeed_id"
	// FieldEmployeeName holds the string denoting the employeename field in the database.
	FieldEmployeeName = "employee_name"
	// FieldEmailid holds the string denoting the emailid field in the database.
	FieldEmailid = "emailid"
	// FieldMobileNumber holds the string denoting the mobilenumber field in the database.
	FieldMobileNumber = "mobile_number"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldOTP holds the string denoting the otp field in the database.
	FieldOTP = "otp"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldVerifyRemarks holds the string denoting the verifyremarks field in the database.
	FieldVerifyRemarks = "verify_remarks"
	// EdgeRoleMaster holds the string denoting the role_master edge name in mutations.
	EdgeRoleMaster = "role_master"
	// RoleMasterFieldID holds the string denoting the ID field of the RoleMaster.
	RoleMasterFieldID = "RoleUserCode"
	// Table holds the table name of the adminlogin in the database.
	Table = "AdminLogin"
	// RoleMasterTable is the table that holds the role_master relation/edge.
	RoleMasterTable = "AdminLogin"
	// RoleMasterInverseTable is the table name for the RoleMaster entity.
	// It exists in this package in order to avoid circular dependency with the "rolemaster" package.
	RoleMasterInverseTable = "RoleMaster"
	// RoleMasterColumn is the table column denoting the role_master relation/edge.
	RoleMasterColumn = "role_user_code"
)

// Columns holds all SQL columns for adminlogin fields.
var Columns = []string{
	FieldID,
	FieldRoleUserCode,
	FieldRoleName,
	FieldCreatedDate,
	FieldStatus,
	FieldEmployeedID,
	FieldEmployeeName,
	FieldEmailid,
	FieldMobileNumber,
	FieldUsername,
	FieldOTP,
	FieldPassword,
	FieldVerifyRemarks,
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

// OrderOption defines the ordering options for the AdminLogin queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoleUserCode orders the results by the RoleUserCode field.
func ByRoleUserCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleUserCode, opts...).ToFunc()
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

// ByEmployeedID orders the results by the EmployeedID field.
func ByEmployeedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeedID, opts...).ToFunc()
}

// ByEmployeeName orders the results by the EmployeeName field.
func ByEmployeeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeName, opts...).ToFunc()
}

// ByEmailid orders the results by the Emailid field.
func ByEmailid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailid, opts...).ToFunc()
}

// ByMobileNumber orders the results by the MobileNumber field.
func ByMobileNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobileNumber, opts...).ToFunc()
}

// ByUsername orders the results by the Username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByOTP orders the results by the OTP field.
func ByOTP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOTP, opts...).ToFunc()
}

// ByPassword orders the results by the Password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByVerifyRemarks orders the results by the VerifyRemarks field.
func ByVerifyRemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerifyRemarks, opts...).ToFunc()
}

// ByRoleMasterField orders the results by role_master field.
func ByRoleMasterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleMasterStep(), sql.OrderByField(field, opts...))
	}
}
func newRoleMasterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleMasterInverseTable, RoleMasterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoleMasterTable, RoleMasterColumn),
	)
}
