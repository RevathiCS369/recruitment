// Code generated by ent, DO NOT EDIT.

package employeemaster

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the employeemaster type in the database.
	Label = "employee_master"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "EmpID"
	// FieldEmployeeID holds the string denoting the employeeid field in the database.
	FieldEmployeeID = "employee_id"
	// FieldEmployeeName holds the string denoting the employeename field in the database.
	FieldEmployeeName = "employee_name"
	// FieldDOB holds the string denoting the dob field in the database.
	FieldDOB = "dob"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldMobileNumber holds the string denoting the mobilenumber field in the database.
	FieldMobileNumber = "mobile_number"
	// FieldEmailID holds the string denoting the emailid field in the database.
	FieldEmailID = "email_id"
	// FieldEmployeeCategoryCode holds the string denoting the employeecategorycode field in the database.
	FieldEmployeeCategoryCode = "employee_category_code"
	// FieldEmployeeCategory holds the string denoting the employeecategory field in the database.
	FieldEmployeeCategory = "employee_category"
	// FieldPostCode holds the string denoting the postcode field in the database.
	FieldPostCode = "post_code"
	// FieldEmployeePost holds the string denoting the employeepost field in the database.
	FieldEmployeePost = "employee_post"
	// FieldFacilityID holds the string denoting the facilityid field in the database.
	FieldFacilityID = "facility_id"
	// FieldDCCS holds the string denoting the dccs field in the database.
	FieldDCCS = "dccs"
	// FieldDCInPresentCadre holds the string denoting the dcinpresentcadre field in the database.
	FieldDCInPresentCadre = "dc_in_present_cadre"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCadre holds the string denoting the cadre field in the database.
	FieldCadre = "cadre"
	// EdgeUsermasterRef holds the string denoting the usermasterref edge name in mutations.
	EdgeUsermasterRef = "UsermasterRef"
	// EdgeEmpRef holds the string denoting the emp_ref edge name in mutations.
	EdgeEmpRef = "Emp_Ref"
	// UserMasterFieldID holds the string denoting the ID field of the UserMaster.
	UserMasterFieldID = "UserID"
	// Exam_Applications_PSFieldID holds the string denoting the ID field of the Exam_Applications_PS.
	Exam_Applications_PSFieldID = "ApplicationID"
	// Table holds the table name of the employeemaster in the database.
	Table = "EmployeeMaster"
	// UsermasterRefTable is the table that holds the UsermasterRef relation/edge.
	UsermasterRefTable = "UserMaster"
	// UsermasterRefInverseTable is the table name for the UserMaster entity.
	// It exists in this package in order to avoid circular dependency with the "usermaster" package.
	UsermasterRefInverseTable = "UserMaster"
	// UsermasterRefColumn is the table column denoting the UsermasterRef relation/edge.
	UsermasterRefColumn = "employee_master_usermaster_ref"
	// EmpRefTable is the table that holds the Emp_Ref relation/edge.
	EmpRefTable = "Exam_Applications_PS"
	// EmpRefInverseTable is the table name for the Exam_Applications_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ps" package.
	EmpRefInverseTable = "Exam_Applications_PS"
	// EmpRefColumn is the table column denoting the Emp_Ref relation/edge.
	EmpRefColumn = "employee_master_emp_ref"
)

// Columns holds all SQL columns for employeemaster fields.
var Columns = []string{
	FieldID,
	FieldEmployeeID,
	FieldEmployeeName,
	FieldDOB,
	FieldGender,
	FieldMobileNumber,
	FieldEmailID,
	FieldEmployeeCategoryCode,
	FieldEmployeeCategory,
	FieldPostCode,
	FieldEmployeePost,
	FieldFacilityID,
	FieldDCCS,
	FieldDCInPresentCadre,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldCadre,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "EmployeeMaster"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_master_usermaster_ref",
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

var (
	// DefaultUpdatedBy holds the default value on creation for the "UpdatedBy" field.
	DefaultUpdatedBy string
)

// Gender defines the type for the "Gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

func (_gender Gender) String() string {
	return string(_gender)
}

// GenderValidator is a validator for the "Gender" field enum values. It is called by the builders before save.
func GenderValidator(_gender Gender) error {
	switch _gender {
	case GenderMale, GenderFemale:
		return nil
	default:
		return fmt.Errorf("employeemaster: invalid enum value for Gender field: %q", _gender)
	}
}

// OrderOption defines the ordering options for the EmployeeMaster queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmployeeID orders the results by the EmployeeID field.
func ByEmployeeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeID, opts...).ToFunc()
}

// ByEmployeeName orders the results by the EmployeeName field.
func ByEmployeeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeName, opts...).ToFunc()
}

// ByDOB orders the results by the DOB field.
func ByDOB(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDOB, opts...).ToFunc()
}

// ByGender orders the results by the Gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByMobileNumber orders the results by the MobileNumber field.
func ByMobileNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobileNumber, opts...).ToFunc()
}

// ByEmailID orders the results by the EmailID field.
func ByEmailID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailID, opts...).ToFunc()
}

// ByEmployeeCategoryCode orders the results by the EmployeeCategoryCode field.
func ByEmployeeCategoryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeCategoryCode, opts...).ToFunc()
}

// ByEmployeeCategory orders the results by the EmployeeCategory field.
func ByEmployeeCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeCategory, opts...).ToFunc()
}

// ByPostCode orders the results by the PostCode field.
func ByPostCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostCode, opts...).ToFunc()
}

// ByEmployeePost orders the results by the EmployeePost field.
func ByEmployeePost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeePost, opts...).ToFunc()
}

// ByFacilityID orders the results by the FacilityID field.
func ByFacilityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFacilityID, opts...).ToFunc()
}

// ByDCCS orders the results by the DCCS field.
func ByDCCS(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDCCS, opts...).ToFunc()
}

// ByDCInPresentCadre orders the results by the DCInPresentCadre field.
func ByDCInPresentCadre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDCInPresentCadre, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCadre orders the results by the Cadre field.
func ByCadre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCadre, opts...).ToFunc()
}

// ByUsermasterRefCount orders the results by UsermasterRef count.
func ByUsermasterRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsermasterRefStep(), opts...)
	}
}

// ByUsermasterRef orders the results by UsermasterRef terms.
func ByUsermasterRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsermasterRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEmpRefCount orders the results by Emp_Ref count.
func ByEmpRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEmpRefStep(), opts...)
	}
}

// ByEmpRef orders the results by Emp_Ref terms.
func ByEmpRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEmpRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsermasterRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsermasterRefInverseTable, UserMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsermasterRefTable, UsermasterRefColumn),
	)
}
func newEmpRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EmpRefInverseTable, Exam_Applications_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EmpRefTable, EmpRefColumn),
	)
}
