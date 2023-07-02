// Code generated by ent, DO NOT EDIT.

package facility

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the facility type in the database.
	Label = "facility"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "FacilityID"
	// FieldFacilityCode holds the string denoting the facilitycode field in the database.
	FieldFacilityCode = "facility_code"
	// FieldOfficeType holds the string denoting the officetype field in the database.
	FieldOfficeType = "office_type"
	// FieldFacilityOfficeID holds the string denoting the facilityofficeid field in the database.
	FieldFacilityOfficeID = "facility_office_id"
	// FieldFacilityName holds the string denoting the facilityname field in the database.
	FieldFacilityName = "facility_name"
	// FieldReportingOfficeType holds the string denoting the reportingofficetype field in the database.
	FieldReportingOfficeType = "reporting_office_type"
	// FieldReportingOfficeCode holds the string denoting the reportingofficecode field in the database.
	FieldReportingOfficeCode = "reporting_office_code"
	// FieldEmailID holds the string denoting the emailid field in the database.
	FieldEmailID = "email_id"
	// FieldMobileNumber holds the string denoting the mobilenumber field in the database.
	FieldMobileNumber = "mobile_number"
	// FieldDivisionCode holds the string denoting the divisioncode field in the database.
	FieldDivisionCode = "division_code"
	// FieldDivisionName holds the string denoting the divisionname field in the database.
	FieldDivisionName = "division_name"
	// FieldDivisionID holds the string denoting the divisionid field in the database.
	FieldDivisionID = "division_id"
	// FieldRegionCode holds the string denoting the regioncode field in the database.
	FieldRegionCode = "region_code"
	// FieldRegionID holds the string denoting the regionid field in the database.
	FieldRegionID = "region_id"
	// FieldRegionName holds the string denoting the regionname field in the database.
	FieldRegionName = "region_name"
	// FieldCircleCode holds the string denoting the circlecode field in the database.
	FieldCircleCode = "circle_code"
	// FieldCircleID holds the string denoting the circleid field in the database.
	FieldCircleID = "circle_id"
	// FieldCircleName holds the string denoting the circlename field in the database.
	FieldCircleName = "circle_name"
	// FieldReportingOfficeID holds the string denoting the reportingofficeid field in the database.
	FieldReportingOfficeID = "reporting_office_id"
	// FieldReportingOfficeName holds the string denoting the reportingofficename field in the database.
	FieldReportingOfficeName = "reporting_office_name"
	// EdgeDivisions holds the string denoting the divisions edge name in mutations.
	EdgeDivisions = "divisions"
	// EdgeRegions holds the string denoting the regions edge name in mutations.
	EdgeRegions = "regions"
	// EdgeCircles holds the string denoting the circles edge name in mutations.
	EdgeCircles = "circles"
	// EdgeCircleRef holds the string denoting the circle_ref edge name in mutations.
	EdgeCircleRef = "circle_ref"
	// EdgeOfficePSRef holds the string denoting the office_ps_ref edge name in mutations.
	EdgeOfficePSRef = "Office_PS_Ref"
	// EdgeOfficeIPRef holds the string denoting the office_ip_ref edge name in mutations.
	EdgeOfficeIPRef = "Office_IP_Ref"
	// DivisionMasterFieldID holds the string denoting the ID field of the DivisionMaster.
	DivisionMasterFieldID = "DivisionID"
	// RegionMasterFieldID holds the string denoting the ID field of the RegionMaster.
	RegionMasterFieldID = "RegionID"
	// CircleMasterFieldID holds the string denoting the ID field of the CircleMaster.
	CircleMasterFieldID = "CircleID"
	// Exam_Applications_PSFieldID holds the string denoting the ID field of the Exam_Applications_PS.
	Exam_Applications_PSFieldID = "ApplicationID"
	// Exam_Applications_IPFieldID holds the string denoting the ID field of the Exam_Applications_IP.
	Exam_Applications_IPFieldID = "ApplicationID"
	// Table holds the table name of the facility in the database.
	Table = "Facility"
	// DivisionsTable is the table that holds the divisions relation/edge.
	DivisionsTable = "Facility"
	// DivisionsInverseTable is the table name for the DivisionMaster entity.
	// It exists in this package in order to avoid circular dependency with the "divisionmaster" package.
	DivisionsInverseTable = "DivisionMaster"
	// DivisionsColumn is the table column denoting the divisions relation/edge.
	DivisionsColumn = "division_id"
	// RegionsTable is the table that holds the regions relation/edge.
	RegionsTable = "Facility"
	// RegionsInverseTable is the table name for the RegionMaster entity.
	// It exists in this package in order to avoid circular dependency with the "regionmaster" package.
	RegionsInverseTable = "RegionMaster"
	// RegionsColumn is the table column denoting the regions relation/edge.
	RegionsColumn = "region_id"
	// CirclesTable is the table that holds the circles relation/edge.
	CirclesTable = "Facility"
	// CirclesInverseTable is the table name for the CircleMaster entity.
	// It exists in this package in order to avoid circular dependency with the "circlemaster" package.
	CirclesInverseTable = "CircleMaster"
	// CirclesColumn is the table column denoting the circles relation/edge.
	CirclesColumn = "circle_id"
	// CircleRefTable is the table that holds the circle_ref relation/edge.
	CircleRefTable = "CircleMaster"
	// CircleRefInverseTable is the table name for the CircleMaster entity.
	// It exists in this package in order to avoid circular dependency with the "circlemaster" package.
	CircleRefInverseTable = "CircleMaster"
	// CircleRefColumn is the table column denoting the circle_ref relation/edge.
	CircleRefColumn = "facility_circle_ref"
	// OfficePSRefTable is the table that holds the Office_PS_Ref relation/edge.
	OfficePSRefTable = "Exam_Applications_PS"
	// OfficePSRefInverseTable is the table name for the Exam_Applications_PS entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ps" package.
	OfficePSRefInverseTable = "Exam_Applications_PS"
	// OfficePSRefColumn is the table column denoting the Office_PS_Ref relation/edge.
	OfficePSRefColumn = "facility_office_ps_ref"
	// OfficeIPRefTable is the table that holds the Office_IP_Ref relation/edge.
	OfficeIPRefTable = "Exam_Applications_IP"
	// OfficeIPRefInverseTable is the table name for the Exam_Applications_IP entity.
	// It exists in this package in order to avoid circular dependency with the "exam_applications_ip" package.
	OfficeIPRefInverseTable = "Exam_Applications_IP"
	// OfficeIPRefColumn is the table column denoting the Office_IP_Ref relation/edge.
	OfficeIPRefColumn = "facility_office_ip_ref"
)

// Columns holds all SQL columns for facility fields.
var Columns = []string{
	FieldID,
	FieldFacilityCode,
	FieldOfficeType,
	FieldFacilityOfficeID,
	FieldFacilityName,
	FieldReportingOfficeType,
	FieldReportingOfficeCode,
	FieldEmailID,
	FieldMobileNumber,
	FieldDivisionCode,
	FieldDivisionName,
	FieldDivisionID,
	FieldRegionCode,
	FieldRegionID,
	FieldRegionName,
	FieldCircleCode,
	FieldCircleID,
	FieldCircleName,
	FieldReportingOfficeID,
	FieldReportingOfficeName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "Facility"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_applications_ip_office_ip_ref",
	"exam_applications_ps_office_ps_ref",
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

// OrderOption defines the ordering options for the Facility queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFacilityCode orders the results by the FacilityCode field.
func ByFacilityCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFacilityCode, opts...).ToFunc()
}

// ByOfficeType orders the results by the OfficeType field.
func ByOfficeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOfficeType, opts...).ToFunc()
}

// ByFacilityOfficeID orders the results by the FacilityOfficeID field.
func ByFacilityOfficeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFacilityOfficeID, opts...).ToFunc()
}

// ByFacilityName orders the results by the FacilityName field.
func ByFacilityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFacilityName, opts...).ToFunc()
}

// ByReportingOfficeType orders the results by the ReportingOfficeType field.
func ByReportingOfficeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportingOfficeType, opts...).ToFunc()
}

// ByReportingOfficeCode orders the results by the ReportingOfficeCode field.
func ByReportingOfficeCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportingOfficeCode, opts...).ToFunc()
}

// ByEmailID orders the results by the EmailID field.
func ByEmailID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailID, opts...).ToFunc()
}

// ByMobileNumber orders the results by the MobileNumber field.
func ByMobileNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobileNumber, opts...).ToFunc()
}

// ByDivisionCode orders the results by the DivisionCode field.
func ByDivisionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDivisionCode, opts...).ToFunc()
}

// ByDivisionName orders the results by the DivisionName field.
func ByDivisionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDivisionName, opts...).ToFunc()
}

// ByDivisionID orders the results by the DivisionID field.
func ByDivisionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDivisionID, opts...).ToFunc()
}

// ByRegionCode orders the results by the RegionCode field.
func ByRegionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionCode, opts...).ToFunc()
}

// ByRegionID orders the results by the RegionID field.
func ByRegionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionID, opts...).ToFunc()
}

// ByRegionName orders the results by the RegionName field.
func ByRegionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionName, opts...).ToFunc()
}

// ByCircleCode orders the results by the CircleCode field.
func ByCircleCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleCode, opts...).ToFunc()
}

// ByCircleID orders the results by the CircleID field.
func ByCircleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleID, opts...).ToFunc()
}

// ByCircleName orders the results by the CircleName field.
func ByCircleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleName, opts...).ToFunc()
}

// ByReportingOfficeID orders the results by the ReportingOfficeID field.
func ByReportingOfficeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportingOfficeID, opts...).ToFunc()
}

// ByReportingOfficeName orders the results by the ReportingOfficeName field.
func ByReportingOfficeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportingOfficeName, opts...).ToFunc()
}

// ByDivisionsField orders the results by divisions field.
func ByDivisionsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDivisionsStep(), sql.OrderByField(field, opts...))
	}
}

// ByRegionsField orders the results by regions field.
func ByRegionsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegionsStep(), sql.OrderByField(field, opts...))
	}
}

// ByCirclesField orders the results by circles field.
func ByCirclesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCirclesStep(), sql.OrderByField(field, opts...))
	}
}

// ByCircleRefCount orders the results by circle_ref count.
func ByCircleRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCircleRefStep(), opts...)
	}
}

// ByCircleRef orders the results by circle_ref terms.
func ByCircleRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCircleRefStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByOfficeIPRefCount orders the results by Office_IP_Ref count.
func ByOfficeIPRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOfficeIPRefStep(), opts...)
	}
}

// ByOfficeIPRef orders the results by Office_IP_Ref terms.
func ByOfficeIPRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOfficeIPRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDivisionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DivisionsInverseTable, DivisionMasterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DivisionsTable, DivisionsColumn),
	)
}
func newRegionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegionsInverseTable, RegionMasterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RegionsTable, RegionsColumn),
	)
}
func newCirclesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CirclesInverseTable, CircleMasterFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CirclesTable, CirclesColumn),
	)
}
func newCircleRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CircleRefInverseTable, CircleMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CircleRefTable, CircleRefColumn),
	)
}
func newOfficePSRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OfficePSRefInverseTable, Exam_Applications_PSFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OfficePSRefTable, OfficePSRefColumn),
	)
}
func newOfficeIPRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OfficeIPRefInverseTable, Exam_Applications_IPFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OfficeIPRefTable, OfficeIPRefColumn),
	)
}
