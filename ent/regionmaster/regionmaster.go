// Code generated by ent, DO NOT EDIT.

package regionmaster

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the regionmaster type in the database.
	Label = "region_master"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "RegionID"
	// FieldRegionCode holds the string denoting the regioncode field in the database.
	FieldRegionCode = "region_code"
	// FieldRegionOfficeId holds the string denoting the regionofficeid field in the database.
	FieldRegionOfficeId = "region_office_id"
	// FieldOfficeType holds the string denoting the officetype field in the database.
	FieldOfficeType = "office_type"
	// FieldRegionOfficeName holds the string denoting the regionofficename field in the database.
	FieldRegionOfficeName = "region_office_name"
	// FieldReportingOfficeType holds the string denoting the reportingofficetype field in the database.
	FieldReportingOfficeType = "reporting_office_type"
	// FieldReportingOfficeCode holds the string denoting the reportingofficecode field in the database.
	FieldReportingOfficeCode = "reporting_office_code"
	// FieldEmailID holds the string denoting the emailid field in the database.
	FieldEmailID = "email_id"
	// FieldMobileNumber holds the string denoting the mobilenumber field in the database.
	FieldMobileNumber = "mobile_number"
	// FieldCircleCode holds the string denoting the circlecode field in the database.
	FieldCircleCode = "circle_code"
	// EdgeCircleRef holds the string denoting the circle_ref edge name in mutations.
	EdgeCircleRef = "circle_ref"
	// EdgeRegions holds the string denoting the regions edge name in mutations.
	EdgeRegions = "regions"
	// EdgeRegionRefRef holds the string denoting the region_ref_ref edge name in mutations.
	EdgeRegionRefRef = "region_ref_ref"
	// CircleMasterFieldID holds the string denoting the ID field of the CircleMaster.
	CircleMasterFieldID = "CircleID"
	// DivisionMasterFieldID holds the string denoting the ID field of the DivisionMaster.
	DivisionMasterFieldID = "DivisionID"
	// FacilityFieldID holds the string denoting the ID field of the Facility.
	FacilityFieldID = "FacilityID"
	// Table holds the table name of the regionmaster in the database.
	Table = "RegionMaster"
	// CircleRefTable is the table that holds the circle_ref relation/edge.
	CircleRefTable = "CircleMaster"
	// CircleRefInverseTable is the table name for the CircleMaster entity.
	// It exists in this package in order to avoid circular dependency with the "circlemaster" package.
	CircleRefInverseTable = "CircleMaster"
	// CircleRefColumn is the table column denoting the circle_ref relation/edge.
	CircleRefColumn = "region_master_circle_ref"
	// RegionsTable is the table that holds the regions relation/edge.
	RegionsTable = "DivisionMaster"
	// RegionsInverseTable is the table name for the DivisionMaster entity.
	// It exists in this package in order to avoid circular dependency with the "divisionmaster" package.
	RegionsInverseTable = "DivisionMaster"
	// RegionsColumn is the table column denoting the regions relation/edge.
	RegionsColumn = "region_master_regions"
	// RegionRefRefTable is the table that holds the region_ref_ref relation/edge.
	RegionRefRefTable = "Facility"
	// RegionRefRefInverseTable is the table name for the Facility entity.
	// It exists in this package in order to avoid circular dependency with the "facility" package.
	RegionRefRefInverseTable = "Facility"
	// RegionRefRefColumn is the table column denoting the region_ref_ref relation/edge.
	RegionRefRefColumn = "region_master_region_ref_ref"
)

// Columns holds all SQL columns for regionmaster fields.
var Columns = []string{
	FieldID,
	FieldRegionCode,
	FieldRegionOfficeId,
	FieldOfficeType,
	FieldRegionOfficeName,
	FieldReportingOfficeType,
	FieldReportingOfficeCode,
	FieldEmailID,
	FieldMobileNumber,
	FieldCircleCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "RegionMaster"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"circle_master_region_ref",
	"division_master_regions",
	"facility_region_ref",
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

// OrderOption defines the ordering options for the RegionMaster queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRegionCode orders the results by the RegionCode field.
func ByRegionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionCode, opts...).ToFunc()
}

// ByRegionOfficeId orders the results by the RegionOfficeId field.
func ByRegionOfficeId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionOfficeId, opts...).ToFunc()
}

// ByOfficeType orders the results by the OfficeType field.
func ByOfficeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOfficeType, opts...).ToFunc()
}

// ByRegionOfficeName orders the results by the RegionOfficeName field.
func ByRegionOfficeName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegionOfficeName, opts...).ToFunc()
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

// ByCircleCode orders the results by the CircleCode field.
func ByCircleCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCircleCode, opts...).ToFunc()
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

// ByRegionsCount orders the results by regions count.
func ByRegionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRegionsStep(), opts...)
	}
}

// ByRegions orders the results by regions terms.
func ByRegions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRegionRefRefCount orders the results by region_ref_ref count.
func ByRegionRefRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRegionRefRefStep(), opts...)
	}
}

// ByRegionRefRef orders the results by region_ref_ref terms.
func ByRegionRefRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRegionRefRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCircleRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CircleRefInverseTable, CircleMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CircleRefTable, CircleRefColumn),
	)
}
func newRegionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegionsInverseTable, DivisionMasterFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RegionsTable, RegionsColumn),
	)
}
func newRegionRefRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RegionRefRefInverseTable, FacilityFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RegionRefRefTable, RegionRefRefColumn),
	)
}