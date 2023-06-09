// Code generated by ent, DO NOT EDIT.

package exam_pa

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the exam_pa type in the database.
	Label = "exam_pa"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ExamCodePA"
	// FieldExamNameCode holds the string denoting the examnamecode field in the database.
	FieldExamNameCode = "exam_name_code"
	// FieldExamName holds the string denoting the examname field in the database.
	FieldExamName = "exam_name"
	// FieldExamType holds the string denoting the examtype field in the database.
	FieldExamType = "exam_type"
	// FieldNotificationCode holds the string denoting the notificationcode field in the database.
	FieldNotificationCode = "notification_code"
	// FieldConductedBy holds the string denoting the conductedby field in the database.
	FieldConductedBy = "conducted_by"
	// FieldNodalOffice holds the string denoting the nodaloffice field in the database.
	FieldNodalOffice = "nodal_office"
	// FieldCalendarCode holds the string denoting the calendarcode field in the database.
	FieldCalendarCode = "calendar_code"
	// FieldPaperCode holds the string denoting the papercode field in the database.
	FieldPaperCode = "paper_code"
	// FieldEligibleCadre holds the string denoting the eligiblecadre field in the database.
	FieldEligibleCadre = "eligible_cadre"
	// FieldEligiblePost1 holds the string denoting the eligiblepost1 field in the database.
	FieldEligiblePost1 = "eligible_post1"
	// FieldEligiblePost2 holds the string denoting the eligiblepost2 field in the database.
	FieldEligiblePost2 = "eligible_post2"
	// FieldEligiblePost3 holds the string denoting the eligiblepost3 field in the database.
	FieldEligiblePost3 = "eligible_post3"
	// FieldEligiblePost4 holds the string denoting the eligiblepost4 field in the database.
	FieldEligiblePost4 = "eligible_post4"
	// FieldEligiblePost5 holds the string denoting the eligiblepost5 field in the database.
	FieldEligiblePost5 = "eligible_post5"
	// FieldExamPost1 holds the string denoting the exampost1 field in the database.
	FieldExamPost1 = "exam_post1"
	// FieldExamPost2 holds the string denoting the exampost2 field in the database.
	FieldExamPost2 = "exam_post2"
	// FieldExamPost3 holds the string denoting the exampost3 field in the database.
	FieldExamPost3 = "exam_post3"
	// FieldExamPost4 holds the string denoting the exampost4 field in the database.
	FieldExamPost4 = "exam_post4"
	// FieldExamPost5 holds the string denoting the exampost5 field in the database.
	FieldExamPost5 = "exam_post5"
	// FieldEducationCriteria holds the string denoting the educationcriteria field in the database.
	FieldEducationCriteria = "education_criteria"
	// FieldCategoryAgeLimitGEN holds the string denoting the categoryagelimitgen field in the database.
	FieldCategoryAgeLimitGEN = "category_age_limit_gen"
	// FieldCategoryAgeLimitSC holds the string denoting the categoryagelimitsc field in the database.
	FieldCategoryAgeLimitSC = "category_age_limit_sc"
	// FieldCategoryAgeLimitST holds the string denoting the categoryagelimitst field in the database.
	FieldCategoryAgeLimitST = "category_age_limit_st"
	// FieldServiceYears holds the string denoting the serviceyears field in the database.
	FieldServiceYears = "service_years"
	// FieldDrivingLicenseRequired holds the string denoting the drivinglicenserequired field in the database.
	FieldDrivingLicenseRequired = "driving_license_required"
	// FieldExamPaperCode holds the string denoting the exampapercode field in the database.
	FieldExamPaperCode = "exam_paper_code"
	// FieldExamPaper1 holds the string denoting the exampaper1 field in the database.
	FieldExamPaper1 = "exam_paper1"
	// FieldExamPaper2 holds the string denoting the exampaper2 field in the database.
	FieldExamPaper2 = "exam_paper2"
	// FieldExamPaper3 holds the string denoting the exampaper3 field in the database.
	FieldExamPaper3 = "exam_paper3"
	// FieldExamPaper4 holds the string denoting the exampaper4 field in the database.
	FieldExamPaper4 = "exam_paper4"
	// FieldExamPaper5 holds the string denoting the exampaper5 field in the database.
	FieldExamPaper5 = "exam_paper5"
	// FieldExamPaper6 holds the string denoting the exampaper6 field in the database.
	FieldExamPaper6 = "exam_paper6"
	// FieldPayLevelEligibilty holds the string denoting the payleveleligibilty field in the database.
	FieldPayLevelEligibilty = "pay_level_eligibilty"
	// FieldCategoryMinMarksSCSTPH holds the string denoting the categoryminmarksscstph field in the database.
	FieldCategoryMinMarksSCSTPH = "category_min_marks_scstph"
	// FieldCategoryMinMarksGENOBC holds the string denoting the categoryminmarksgenobc field in the database.
	FieldCategoryMinMarksGENOBC = "category_min_marks_genobc"
	// FieldLocalLanguageAllowed holds the string denoting the locallanguageallowed field in the database.
	FieldLocalLanguageAllowed = "local_language_allowed"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "updated_by"
	// EdgeExamcalPsRef holds the string denoting the examcal_ps_ref edge name in mutations.
	EdgeExamcalPsRef = "examcal_ps_ref"
	// EdgePapersPsRef holds the string denoting the papers_ps_ref edge name in mutations.
	EdgePapersPsRef = "papers_ps_ref"
	// EdgeUsersPsType holds the string denoting the users_ps_type edge name in mutations.
	EdgeUsersPsType = "users_ps_type"
	// EdgeExamApplnPSRef holds the string denoting the examappln_ps_ref edge name in mutations.
	EdgeExamApplnPSRef = "ExamAppln_PS_Ref"
	// EdgeNotificationsPs holds the string denoting the notifications_ps edge name in mutations.
	EdgeNotificationsPs = "notifications_ps"
	// ExamCalendarFieldID holds the string denoting the ID field of the ExamCalendar.
	ExamCalendarFieldID = "CalendarCode"
	// ExamPapersFieldID holds the string denoting the ID field of the ExamPapers.
	ExamPapersFieldID = "PaperCode"
	// UserMasterFieldID holds the string denoting the ID field of the UserMaster.
	UserMasterFieldID = "UserID"
	// Exam_Applications_PSFieldID holds the string denoting the ID field of the Exam_Applications_PS.
	Exam_Applications_PSFieldID = "ApplicationID"
	// NotificationFieldID holds the string denoting the ID field of the Notification.
	NotificationFieldID = "NotifyCode"
	// Table holds the table name of the exam_pa in the database.
	Table = "Exam_PA"
	// ExamcalPsRefTable is the table that holds the examcal_ps_ref relation/edge.
	ExamcalPsRefTable = "ExamCalendar"
	// ExamcalPsRefInverseTable is the table name for the ExamCalendar entity.
	// It exists in this package in order to avoid circular dependency with the "examcalendar" package.
	ExamcalPsRefInverseTable = "ExamCalendar"
	// ExamcalPsRefColumn is the table column denoting the examcal_ps_ref relation/edge.
	ExamcalPsRefColumn = "exam_pa_examcal_ps_ref"
	// PapersPsRefTable is the table that holds the papers_ps_ref relation/edge.
	PapersPsRefTable = "ExamPapers"
	// PapersPsRefInverseTable is the table name for the ExamPapers entity.
	// It exists in this package in order to avoid circular dependency with the "exampapers" package.
	PapersPsRefInverseTable = "ExamPapers"
	// PapersPsRefColumn is the table column denoting the papers_ps_ref relation/edge.
	PapersPsRefColumn = "exam_pa_papers_ps_ref"
	// UsersPsTypeTable is the table that holds the users_ps_type relation/edge.
	UsersPsTypeTable = "UserMaster"
	// UsersPsTypeInverseTable is the table name for the UserMaster entity.
	// It exists in this package in order to avoid circular dependency with the "usermaster" package.
	UsersPsTypeInverseTable = "UserMaster"
	// UsersPsTypeColumn is the table column denoting the users_ps_type relation/edge.
	UsersPsTypeColumn = "exam_pa_users_ps_type"
	// ExamApplnPSRefTable is the table that holds the ExamAppln_PS_Ref relation/edge.
	ExamApplnPSRefTable = "Exam_Applications_PS"
	// ExamApplnPSRefInverseTable is the table name for the Exam_Applications_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ps" package.
	ExamApplnPSRefInverseTable = "Exam_Applications_PS"
	// ExamApplnPSRefColumn is the table column denoting the ExamAppln_PS_Ref relation/edge.
	ExamApplnPSRefColumn = "exam_pa_exam_appln_ps_ref"
	// NotificationsPsTable is the table that holds the notifications_ps relation/edge.
	NotificationsPsTable = "Notification"
	// NotificationsPsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsPsInverseTable = "Notification"
	// NotificationsPsColumn is the table column denoting the notifications_ps relation/edge.
	NotificationsPsColumn = "exam_pa_notifications_ps"
)

// Columns holds all SQL columns for exam_pa fields.
var Columns = []string{
	FieldID,
	FieldExamNameCode,
	FieldExamName,
	FieldExamType,
	FieldNotificationCode,
	FieldConductedBy,
	FieldNodalOffice,
	FieldCalendarCode,
	FieldPaperCode,
	FieldEligibleCadre,
	FieldEligiblePost1,
	FieldEligiblePost2,
	FieldEligiblePost3,
	FieldEligiblePost4,
	FieldEligiblePost5,
	FieldExamPost1,
	FieldExamPost2,
	FieldExamPost3,
	FieldExamPost4,
	FieldExamPost5,
	FieldEducationCriteria,
	FieldCategoryAgeLimitGEN,
	FieldCategoryAgeLimitSC,
	FieldCategoryAgeLimitST,
	FieldServiceYears,
	FieldDrivingLicenseRequired,
	FieldExamPaperCode,
	FieldExamPaper1,
	FieldExamPaper2,
	FieldExamPaper3,
	FieldExamPaper4,
	FieldExamPaper5,
	FieldExamPaper6,
	FieldPayLevelEligibilty,
	FieldCategoryMinMarksSCSTPH,
	FieldCategoryMinMarksGENOBC,
	FieldLocalLanguageAllowed,
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

// OrderOption defines the ordering options for the Exam_PA queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByExamNameCode orders the results by the ExamNameCode field.
func ByExamNameCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamNameCode, opts...).ToFunc()
}

// ByExamName orders the results by the ExamName field.
func ByExamName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamName, opts...).ToFunc()
}

// ByExamType orders the results by the ExamType field.
func ByExamType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamType, opts...).ToFunc()
}

// ByNotificationCode orders the results by the NotificationCode field.
func ByNotificationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotificationCode, opts...).ToFunc()
}

// ByConductedBy orders the results by the ConductedBy field.
func ByConductedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConductedBy, opts...).ToFunc()
}

// ByNodalOffice orders the results by the NodalOffice field.
func ByNodalOffice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodalOffice, opts...).ToFunc()
}

// ByCalendarCode orders the results by the CalendarCode field.
func ByCalendarCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCalendarCode, opts...).ToFunc()
}

// ByPaperCode orders the results by the PaperCode field.
func ByPaperCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaperCode, opts...).ToFunc()
}

// ByEligibleCadre orders the results by the EligibleCadre field.
func ByEligibleCadre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligibleCadre, opts...).ToFunc()
}

// ByEligiblePost1 orders the results by the EligiblePost1 field.
func ByEligiblePost1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligiblePost1, opts...).ToFunc()
}

// ByEligiblePost2 orders the results by the EligiblePost2 field.
func ByEligiblePost2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligiblePost2, opts...).ToFunc()
}

// ByEligiblePost3 orders the results by the EligiblePost3 field.
func ByEligiblePost3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligiblePost3, opts...).ToFunc()
}

// ByEligiblePost4 orders the results by the EligiblePost4 field.
func ByEligiblePost4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligiblePost4, opts...).ToFunc()
}

// ByEligiblePost5 orders the results by the EligiblePost5 field.
func ByEligiblePost5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEligiblePost5, opts...).ToFunc()
}

// ByExamPost1 orders the results by the ExamPost1 field.
func ByExamPost1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPost1, opts...).ToFunc()
}

// ByExamPost2 orders the results by the ExamPost2 field.
func ByExamPost2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPost2, opts...).ToFunc()
}

// ByExamPost3 orders the results by the ExamPost3 field.
func ByExamPost3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPost3, opts...).ToFunc()
}

// ByExamPost4 orders the results by the ExamPost4 field.
func ByExamPost4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPost4, opts...).ToFunc()
}

// ByExamPost5 orders the results by the ExamPost5 field.
func ByExamPost5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPost5, opts...).ToFunc()
}

// ByEducationCriteria orders the results by the EducationCriteria field.
func ByEducationCriteria(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEducationCriteria, opts...).ToFunc()
}

// ByCategoryAgeLimitGEN orders the results by the CategoryAgeLimitGEN field.
func ByCategoryAgeLimitGEN(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryAgeLimitGEN, opts...).ToFunc()
}

// ByCategoryAgeLimitSC orders the results by the CategoryAgeLimitSC field.
func ByCategoryAgeLimitSC(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryAgeLimitSC, opts...).ToFunc()
}

// ByCategoryAgeLimitST orders the results by the CategoryAgeLimitST field.
func ByCategoryAgeLimitST(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryAgeLimitST, opts...).ToFunc()
}

// ByServiceYears orders the results by the ServiceYears field.
func ByServiceYears(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceYears, opts...).ToFunc()
}

// ByDrivingLicenseRequired orders the results by the DrivingLicenseRequired field.
func ByDrivingLicenseRequired(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDrivingLicenseRequired, opts...).ToFunc()
}

// ByExamPaperCode orders the results by the ExamPaperCode field.
func ByExamPaperCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaperCode, opts...).ToFunc()
}

// ByExamPaper1 orders the results by the ExamPaper1 field.
func ByExamPaper1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper1, opts...).ToFunc()
}

// ByExamPaper2 orders the results by the ExamPaper2 field.
func ByExamPaper2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper2, opts...).ToFunc()
}

// ByExamPaper3 orders the results by the ExamPaper3 field.
func ByExamPaper3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper3, opts...).ToFunc()
}

// ByExamPaper4 orders the results by the ExamPaper4 field.
func ByExamPaper4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper4, opts...).ToFunc()
}

// ByExamPaper5 orders the results by the ExamPaper5 field.
func ByExamPaper5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper5, opts...).ToFunc()
}

// ByExamPaper6 orders the results by the ExamPaper6 field.
func ByExamPaper6(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExamPaper6, opts...).ToFunc()
}

// ByPayLevelEligibilty orders the results by the PayLevelEligibilty field.
func ByPayLevelEligibilty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayLevelEligibilty, opts...).ToFunc()
}

// ByCategoryMinMarksSCSTPH orders the results by the CategoryMinMarksSCSTPH field.
func ByCategoryMinMarksSCSTPH(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryMinMarksSCSTPH, opts...).ToFunc()
}

// ByCategoryMinMarksGENOBC orders the results by the CategoryMinMarksGENOBC field.
func ByCategoryMinMarksGENOBC(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryMinMarksGENOBC, opts...).ToFunc()
}

// ByLocalLanguageAllowed orders the results by the LocalLanguageAllowed field.
func ByLocalLanguageAllowed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocalLanguageAllowed, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByExamcalPsRefCount orders the results by examcal_ps_ref count.
func ByExamcalPsRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamcalPsRefStep(), opts...)
	}
}

// ByExamcalPsRef orders the results by examcal_ps_ref terms.
func ByExamcalPsRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamcalPsRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPapersPsRefCount orders the results by papers_ps_ref count.
func ByPapersPsRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPapersPsRefStep(), opts...)
	}
}

// ByPapersPsRef orders the results by papers_ps_ref terms.
func ByPapersPsRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPapersPsRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersPsTypeCount orders the results by users_ps_type count.
func ByUsersPsTypeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersPsTypeStep(), opts...)
	}
}

// ByUsersPsType orders the results by users_ps_type terms.
func ByUsersPsType(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersPsTypeStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByNotificationsPsCount orders the results by notifications_ps count.
func ByNotificationsPsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsPsStep(), opts...)
	}
}

// ByNotificationsPs orders the results by notifications_ps terms.
func ByNotificationsPs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsPsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExamcalPsRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamcalPsRefInverseTable, ExamCalendarFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamcalPsRefTable, ExamcalPsRefColumn),
	)
}
func newPapersPsRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PapersPsRefInverseTable, ExamPapersFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PapersPsRefTable, PapersPsRefColumn),
	)
}
func newUsersPsTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersPsTypeInverseTable, UserMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersPsTypeTable, UsersPsTypeColumn),
	)
}
func newExamApplnPSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamApplnPSRefInverseTable, Exam_Applications_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamApplnPSRefTable, ExamApplnPSRefColumn),
	)
}
func newNotificationsPsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsPsInverseTable, NotificationFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsPsTable, NotificationsPsColumn),
	)
}
