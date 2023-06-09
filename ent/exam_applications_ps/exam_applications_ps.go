// Code generated by ent, DO NOT EDIT.

package exam_applications_ps

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the exam_applications_ps type in the database.
	Label = "exam_applications_ps"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ApplicationID"
	// FieldApplicationNumber holds the string denoting the applicationnumber field in the database.
	FieldApplicationNumber = "application_number"
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
	// FieldEmployeeCategory holds the string denoting the employeecategory field in the database.
	FieldEmployeeCategory = "employee_category"
	// FieldCadre holds the string denoting the cadre field in the database.
	FieldCadre = "cadre"
	// FieldEmployeePost holds the string denoting the employeepost field in the database.
	FieldEmployeePost = "employee_post"
	// FieldFacilityID holds the string denoting the facilityid field in the database.
	FieldFacilityID = "facility_id"
	// FieldDCCS holds the string denoting the dccs field in the database.
	FieldDCCS = "dccs"
	// FieldDCInPresentCadre holds the string denoting the dcinpresentcadre field in the database.
	FieldDCInPresentCadre = "dc_in_present_cadre"
	// FieldDeputationOfficeId holds the string denoting the deputationofficeid field in the database.
	FieldDeputationOfficeId = "deputation_office_id"
	// FieldDisabilityType holds the string denoting the disabilitytype field in the database.
	FieldDisabilityType = "disability_type"
	// FieldDisabilityPercentage holds the string denoting the disabilitypercentage field in the database.
	FieldDisabilityPercentage = "disability_percentage"
	// FieldEducation holds the string denoting the education field in the database.
	FieldEducation = "education"
	// FieldExamNameCode holds the string denoting the examnamecode field in the database.
	FieldExamNameCode = "exam_name_code"
	// FieldExamYear holds the string denoting the examyear field in the database.
	FieldExamYear = "exam_year"
	// FieldExamName holds the string denoting the examname field in the database.
	FieldExamName = "exam_name"
	// FieldCentrePreference holds the string denoting the centrepreference field in the database.
	FieldCentrePreference = "centre_preference"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// FieldApplicationStatus holds the string denoting the applicationstatus field in the database.
	FieldApplicationStatus = "application_status"
	// FieldApplnSubmittedDate holds the string denoting the applnsubmitteddate field in the database.
	FieldApplnSubmittedDate = "appln_submitted_date"
	// FieldVARemarks holds the string denoting the va_remarks field in the database.
	FieldVARemarks = "va_remarks"
	// FieldVAUserName holds the string denoting the va_username field in the database.
	FieldVAUserName = "va_user_name"
	// FieldVADate holds the string denoting the va_date field in the database.
	FieldVADate = "va_date"
	// FieldCARemarks holds the string denoting the ca_remarks field in the database.
	FieldCARemarks = "ca_remarks"
	// FieldCAUserName holds the string denoting the ca_username field in the database.
	FieldCAUserName = "ca_user_name"
	// FieldCADate holds the string denoting the ca_date field in the database.
	FieldCADate = "ca_date"
	// FieldApplicationWdlDate holds the string denoting the applicationwdldate field in the database.
	FieldApplicationWdlDate = "application_wdl_date"
	// FieldNARemarks holds the string denoting the na_remarks field in the database.
	FieldNARemarks = "na_remarks"
	// FieldNAUserName holds the string denoting the na_username field in the database.
	FieldNAUserName = "na_user_name"
	// FieldNADate holds the string denoting the na_date field in the database.
	FieldNADate = "na_date"
	// FieldAppliactionRemarks holds the string denoting the appliactionremarks field in the database.
	FieldAppliactionRemarks = "appliaction_remarks"
	// FieldCadrePreferences holds the string denoting the cadrepreferences field in the database.
	FieldCadrePreferences = "cadre_preferences"
	// FieldDivisionPreferences holds the string denoting the divisionpreferences field in the database.
	FieldDivisionPreferences = "division_preferences"
	// FieldCirclePreferences holds the string denoting the circlepreferences field in the database.
	FieldCirclePreferences = "circle_preferences"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldRoleUserCode holds the string denoting the roleusercode field in the database.
	FieldRoleUserCode = "role_user_code"
	// EdgeUsersPSRef holds the string denoting the userspsref edge name in mutations.
	EdgeUsersPSRef = "UsersPSRef"
	// EdgeExamApplnPSRef holds the string denoting the examappln_ps_ref edge name in mutations.
	EdgeExamApplnPSRef = "ExamAppln_PS_Ref"
	// EdgeOfficePSRef holds the string denoting the office_ps_ref edge name in mutations.
	EdgeOfficePSRef = "Office_PS_Ref"
	// EdgeRoleusers holds the string denoting the roleusers edge name in mutations.
	EdgeRoleusers = "roleusers"
	// UserMasterFieldID holds the string denoting the ID field of the UserMaster.
	UserMasterFieldID = "UserID"
	// Exam_PSFieldID holds the string denoting the ID field of the Exam_PS.
	Exam_PSFieldID = "ExamCodePS"
	// FacilityFieldID holds the string denoting the ID field of the Facility.
	FacilityFieldID = "FacilityID"
	// RoleMasterFieldID holds the string denoting the ID field of the RoleMaster.
	RoleMasterFieldID = "RoleUserCode"
	// Table holds the table name of the exam_applications_ps in the database.
	Table = "Exam_Applications_PS"
	// UsersPSRefTable is the table that holds the UsersPSRef relation/edge.
	UsersPSRefTable = "UserMaster"
	// UsersPSRefInverseTable is the table name for the UserMaster entity.
	// It exists in this package in order to avoid circular dependency with the "usermaster" package.
	UsersPSRefInverseTable = "UserMaster"
	// UsersPSRefColumn is the table column denoting the UsersPSRef relation/edge.
	UsersPSRefColumn = "exam_applications_ps_users_ps_ref"
	// ExamApplnPSRefTable is the table that holds the ExamAppln_PS_Ref relation/edge.
	ExamApplnPSRefTable = "Exam_PS"
	// ExamApplnPSRefInverseTable is the table name for the Exam_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_ps" package.
	ExamApplnPSRefInverseTable = "Exam_PS"
	// ExamApplnPSRefColumn is the table column denoting the ExamAppln_PS_Ref relation/edge.
	ExamApplnPSRefColumn = "exam_applications_ps_exam_appln_ps_ref"
	// OfficePSRefTable is the table that holds the Office_PS_Ref relation/edge.
	OfficePSRefTable = "Facility"
	// OfficePSRefInverseTable is the table name for the Facility entity.
	// It exists in this package in order to avoid circular dependency with the "facility" package.
	OfficePSRefInverseTable = "Facility"
	// OfficePSRefColumn is the table column denoting the Office_PS_Ref relation/edge.
	OfficePSRefColumn = "exam_applications_ps_office_ps_ref"
	// RoleusersTable is the table that holds the roleusers relation/edge.
	RoleusersTable = "Exam_Applications_PS"
	// RoleusersInverseTable is the table name for the RoleMaster entity.
	// It exists in this package in order to avoid circular dependency with the "rolemaster" package.
	RoleusersInverseTable = "RoleMaster"
	// RoleusersColumn is the table column denoting the roleusers relation/edge.
	RoleusersColumn = "role_user_code"
)

// Columns holds all SQL columns for exam_applications_ps fields.
var Columns = []string{
	FieldID,
	FieldApplicationNumber,
	FieldEmployeeID,
	FieldEmployeeName,
	FieldDOB,
	FieldGender,
	FieldMobileNumber,
	FieldEmailID,
	FieldEmployeeCategory,
	FieldCadre,
	FieldEmployeePost,
	FieldFacilityID,
	FieldDCCS,
	FieldDCInPresentCadre,
	FieldDeputationOfficeId,
	FieldDisabilityType,
	FieldDisabilityPercentage,
	FieldEducation,
	FieldExamNameCode,
	FieldExamYear,
	FieldExamName,
	FieldCentrePreference,
	FieldSignature,
	FieldPhoto,
	FieldApplicationStatus,
	FieldApplnSubmittedDate,
	FieldVARemarks,
	FieldVAUserName,
	FieldVADate,
	FieldCARemarks,
	FieldCAUserName,
	FieldCADate,
	FieldApplicationWdlDate,
	FieldNARemarks,
	FieldNAUserName,
	FieldNADate,
	FieldAppliactionRemarks,
	FieldCadrePreferences,
	FieldDivisionPreferences,
	FieldCirclePreferences,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldRoleUserCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "Exam_Applications_PS"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"employee_master_emp_ref",
	"exam_pa_exam_appln_ps_ref",
	"exam_ps_exam_appln_ps_ref",
	"facility_office_ps_ref",
	"user_master_users_ps_ref",
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
	// DefaultApplnSubmittedDate holds the default value on creation for the "ApplnSubmittedDate" field.
	DefaultApplnSubmittedDate func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
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
		return fmt.Errorf("exam_applications_ps: invalid enum value for Gender field: %q", _gender)
	}
}

// OrderOption defines the ordering options for the Exam_Applications_PS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApplicationNumber orders the results by the ApplicationNumber field.
func ByApplicationNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationNumber, opts...).ToFunc()
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

// ByEmployeeCategory orders the results by the EmployeeCategory field.
func ByEmployeeCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmployeeCategory, opts...).ToFunc()
}

// ByCadre orders the results by the Cadre field.
func ByCadre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCadre, opts...).ToFunc()
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

// ByDeputationOfficeId orders the results by the DeputationOfficeId field.
func ByDeputationOfficeId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeputationOfficeId, opts...).ToFunc()
}

// ByDisabilityType orders the results by the DisabilityType field.
func ByDisabilityType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabilityType, opts...).ToFunc()
}

// ByDisabilityPercentage orders the results by the DisabilityPercentage field.
func ByDisabilityPercentage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabilityPercentage, opts...).ToFunc()
}

// ByEducation orders the results by the Education field.
func ByEducation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEducation, opts...).ToFunc()
}

// ByExamNameCode orders the results by the ExamNameCode field.
func ByExamNameCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamNameCode, opts...).ToFunc()
}

// ByExamYear orders the results by the ExamYear field.
func ByExamYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamYear, opts...).ToFunc()
}

// ByExamName orders the results by the ExamName field.
func ByExamName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamName, opts...).ToFunc()
}

// ByCentrePreference orders the results by the CentrePreference field.
func ByCentrePreference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCentrePreference, opts...).ToFunc()
}

// BySignature orders the results by the Signature field.
func BySignature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignature, opts...).ToFunc()
}

// ByPhoto orders the results by the Photo field.
func ByPhoto(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoto, opts...).ToFunc()
}

// ByApplicationStatus orders the results by the ApplicationStatus field.
func ByApplicationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationStatus, opts...).ToFunc()
}

// ByApplnSubmittedDate orders the results by the ApplnSubmittedDate field.
func ByApplnSubmittedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplnSubmittedDate, opts...).ToFunc()
}

// ByVARemarks orders the results by the VA_Remarks field.
func ByVARemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVARemarks, opts...).ToFunc()
}

// ByVAUserName orders the results by the VA_UserName field.
func ByVAUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVAUserName, opts...).ToFunc()
}

// ByVADate orders the results by the VA_Date field.
func ByVADate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVADate, opts...).ToFunc()
}

// ByCARemarks orders the results by the CA_Remarks field.
func ByCARemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCARemarks, opts...).ToFunc()
}

// ByCAUserName orders the results by the CA_UserName field.
func ByCAUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCAUserName, opts...).ToFunc()
}

// ByCADate orders the results by the CA_Date field.
func ByCADate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCADate, opts...).ToFunc()
}

// ByApplicationWdlDate orders the results by the ApplicationWdlDate field.
func ByApplicationWdlDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationWdlDate, opts...).ToFunc()
}

// ByNARemarks orders the results by the NA_Remarks field.
func ByNARemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNARemarks, opts...).ToFunc()
}

// ByNAUserName orders the results by the NA_UserName field.
func ByNAUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNAUserName, opts...).ToFunc()
}

// ByNADate orders the results by the NA_Date field.
func ByNADate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNADate, opts...).ToFunc()
}

// ByAppliactionRemarks orders the results by the AppliactionRemarks field.
func ByAppliactionRemarks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppliactionRemarks, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByRoleUserCode orders the results by the RoleUserCode field.
func ByRoleUserCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleUserCode, opts...).ToFunc()
}

// ByUsersPSRefCount orders the results by UsersPSRef count.
func ByUsersPSRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersPSRefStep(), opts...)
	}
}

// ByUsersPSRef orders the results by UsersPSRef terms.
func ByUsersPSRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersPSRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamApplnPSRefCount orders the results by ExamAppln_PS_Ref count.
func ByExamApplnPSRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamApplnPSRefStep(), opts...)
	}
}

// ByExamApplnPSRef orders the results by ExamAppln_PS_Ref terms.
func ByExamApplnPSRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamApplnPSRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOfficePSRefCount orders the results by Office_PS_Ref count.
func ByOfficePSRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOfficePSRefStep(), opts...)
	}
}

// ByOfficePSRef orders the results by Office_PS_Ref terms.
func ByOfficePSRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOfficePSRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoleusersField orders the results by roleusers field.
func ByRoleusersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleusersStep(), sql.OrderByField(field, opts...))
	}
}
func newUsersPSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersPSRefInverseTable, UserMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersPSRefTable, UsersPSRefColumn),
	)
}
func newExamApplnPSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamApplnPSRefInverseTable, Exam_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamApplnPSRefTable, ExamApplnPSRefColumn),
	)
}
func newOfficePSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OfficePSRefInverseTable, FacilityFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OfficePSRefTable, OfficePSRefColumn),
	)
}
func newRoleusersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleusersInverseTable, RoleMasterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoleusersTable, RoleusersColumn),
	)
}
