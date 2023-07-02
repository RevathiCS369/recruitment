// Code generated by ent, DO NOT EDIT.

package reversal_application_ip

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the reversal_application_ip type in the database.
	Label = "reversal_application_ip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ApplicationID"
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
	// FieldExamCodeIP holds the string denoting the examcodeip field in the database.
	FieldExamCodeIP = "exam_code_ip"
	// FieldExamYear holds the string denoting the examyear field in the database.
	FieldExamYear = "exam_year"
	// FieldCentrePreference holds the string denoting the centrepreference field in the database.
	FieldCentrePreference = "centre_preference"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// FieldApplicationStatus holds the string denoting the applicationstatus field in the database.
	FieldApplicationStatus = "application_status"
	// FieldReversalApplnSubmittedDate holds the string denoting the reversalapplnsubmitteddate field in the database.
	FieldReversalApplnSubmittedDate = "reversal_appln_submitted_date"
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
	// FieldAppliactionRemarks holds the string denoting the appliactionremarks field in the database.
	FieldAppliactionRemarks = "appliaction_remarks"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldRoleUserCode holds the string denoting the roleusercode field in the database.
	FieldRoleUserCode = "role_user_code"
	// Table holds the table name of the reversal_application_ip in the database.
	Table = "Reversal_Application_IP"
)

// Columns holds all SQL columns for reversal_application_ip fields.
var Columns = []string{
	FieldID,
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
	FieldExamCodeIP,
	FieldExamYear,
	FieldCentrePreference,
	FieldSignature,
	FieldPhoto,
	FieldApplicationStatus,
	FieldReversalApplnSubmittedDate,
	FieldVARemarks,
	FieldVAUserName,
	FieldVADate,
	FieldCARemarks,
	FieldCAUserName,
	FieldCADate,
	FieldAppliactionRemarks,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldRoleUserCode,
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
		return fmt.Errorf("reversal_application_ip: invalid enum value for Gender field: %q", _gender)
	}
}

// OrderOption defines the ordering options for the Reversal_Application_IP queries.
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

// ByExamCodeIP orders the results by the ExamCodeIP field.
func ByExamCodeIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamCodeIP, opts...).ToFunc()
}

// ByExamYear orders the results by the ExamYear field.
func ByExamYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamYear, opts...).ToFunc()
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

// ByReversalApplnSubmittedDate orders the results by the ReversalApplnSubmittedDate field.
func ByReversalApplnSubmittedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReversalApplnSubmittedDate, opts...).ToFunc()
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
