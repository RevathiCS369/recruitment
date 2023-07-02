// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/regionmaster"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RegionMaster is the model entity for the RegionMaster schema.
type RegionMaster struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// RegionCode holds the value of the "RegionCode" field.
	RegionCode int32 `json:"RegionCode,omitempty"`
	// RegionOfficeId holds the value of the "RegionOfficeId" field.
	RegionOfficeId string `json:"RegionOfficeId,omitempty"`
	// OfficeType holds the value of the "OfficeType" field.
	OfficeType string `json:"OfficeType,omitempty"`
	// RegionOfficeName holds the value of the "RegionOfficeName" field.
	RegionOfficeName string `json:"RegionOfficeName,omitempty"`
	// ReportingOfficeType holds the value of the "ReportingOfficeType" field.
	ReportingOfficeType string `json:"ReportingOfficeType,omitempty"`
	// ReportingOfficeCode holds the value of the "ReportingOfficeCode" field.
	ReportingOfficeCode string `json:"ReportingOfficeCode,omitempty"`
	// EmailID holds the value of the "EmailID" field.
	EmailID string `json:"EmailID,omitempty"`
	// MobileNumber holds the value of the "MobileNumber" field.
	MobileNumber int32 `json:"MobileNumber,omitempty"`
	// CircleCode holds the value of the "CircleCode" field.
	CircleCode int32 `json:"CircleCode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RegionMasterQuery when eager-loading is set.
	Edges                    RegionMasterEdges `json:"edges"`
	circle_master_region_ref *int32
	division_master_regions  *int32
	facility_region_ref      *int32
	selectValues             sql.SelectValues
}

// RegionMasterEdges holds the relations/edges for other nodes in the graph.
type RegionMasterEdges struct {
	// CircleRef holds the value of the circle_ref edge.
	CircleRef []*CircleMaster `json:"circle_ref,omitempty"`
	// Regions holds the value of the regions edge.
	Regions []*DivisionMaster `json:"regions,omitempty"`
	// RegionRefRef holds the value of the region_ref_ref edge.
	RegionRefRef []*Facility `json:"region_ref_ref,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CircleRefOrErr returns the CircleRef value or an error if the edge
// was not loaded in eager-loading.
func (e RegionMasterEdges) CircleRefOrErr() ([]*CircleMaster, error) {
	if e.loadedTypes[0] {
		return e.CircleRef, nil
	}
	return nil, &NotLoadedError{edge: "circle_ref"}
}

// RegionsOrErr returns the Regions value or an error if the edge
// was not loaded in eager-loading.
func (e RegionMasterEdges) RegionsOrErr() ([]*DivisionMaster, error) {
	if e.loadedTypes[1] {
		return e.Regions, nil
	}
	return nil, &NotLoadedError{edge: "regions"}
}

// RegionRefRefOrErr returns the RegionRefRef value or an error if the edge
// was not loaded in eager-loading.
func (e RegionMasterEdges) RegionRefRefOrErr() ([]*Facility, error) {
	if e.loadedTypes[2] {
		return e.RegionRefRef, nil
	}
	return nil, &NotLoadedError{edge: "region_ref_ref"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RegionMaster) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case regionmaster.FieldID, regionmaster.FieldRegionCode, regionmaster.FieldMobileNumber, regionmaster.FieldCircleCode:
			values[i] = new(sql.NullInt64)
		case regionmaster.FieldRegionOfficeId, regionmaster.FieldOfficeType, regionmaster.FieldRegionOfficeName, regionmaster.FieldReportingOfficeType, regionmaster.FieldReportingOfficeCode, regionmaster.FieldEmailID:
			values[i] = new(sql.NullString)
		case regionmaster.ForeignKeys[0]: // circle_master_region_ref
			values[i] = new(sql.NullInt64)
		case regionmaster.ForeignKeys[1]: // division_master_regions
			values[i] = new(sql.NullInt64)
		case regionmaster.ForeignKeys[2]: // facility_region_ref
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RegionMaster fields.
func (rm *RegionMaster) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case regionmaster.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rm.ID = int32(value.Int64)
		case regionmaster.FieldRegionCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RegionCode", values[i])
			} else if value.Valid {
				rm.RegionCode = int32(value.Int64)
			}
		case regionmaster.FieldRegionOfficeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field RegionOfficeId", values[i])
			} else if value.Valid {
				rm.RegionOfficeId = value.String
			}
		case regionmaster.FieldOfficeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OfficeType", values[i])
			} else if value.Valid {
				rm.OfficeType = value.String
			}
		case regionmaster.FieldRegionOfficeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field RegionOfficeName", values[i])
			} else if value.Valid {
				rm.RegionOfficeName = value.String
			}
		case regionmaster.FieldReportingOfficeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ReportingOfficeType", values[i])
			} else if value.Valid {
				rm.ReportingOfficeType = value.String
			}
		case regionmaster.FieldReportingOfficeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ReportingOfficeCode", values[i])
			} else if value.Valid {
				rm.ReportingOfficeCode = value.String
			}
		case regionmaster.FieldEmailID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmailID", values[i])
			} else if value.Valid {
				rm.EmailID = value.String
			}
		case regionmaster.FieldMobileNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field MobileNumber", values[i])
			} else if value.Valid {
				rm.MobileNumber = int32(value.Int64)
			}
		case regionmaster.FieldCircleCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CircleCode", values[i])
			} else if value.Valid {
				rm.CircleCode = int32(value.Int64)
			}
		case regionmaster.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field circle_master_region_ref", value)
			} else if value.Valid {
				rm.circle_master_region_ref = new(int32)
				*rm.circle_master_region_ref = int32(value.Int64)
			}
		case regionmaster.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field division_master_regions", value)
			} else if value.Valid {
				rm.division_master_regions = new(int32)
				*rm.division_master_regions = int32(value.Int64)
			}
		case regionmaster.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field facility_region_ref", value)
			} else if value.Valid {
				rm.facility_region_ref = new(int32)
				*rm.facility_region_ref = int32(value.Int64)
			}
		default:
			rm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RegionMaster.
// This includes values selected through modifiers, order, etc.
func (rm *RegionMaster) Value(name string) (ent.Value, error) {
	return rm.selectValues.Get(name)
}

// QueryCircleRef queries the "circle_ref" edge of the RegionMaster entity.
func (rm *RegionMaster) QueryCircleRef() *CircleMasterQuery {
	return NewRegionMasterClient(rm.config).QueryCircleRef(rm)
}

// QueryRegions queries the "regions" edge of the RegionMaster entity.
func (rm *RegionMaster) QueryRegions() *DivisionMasterQuery {
	return NewRegionMasterClient(rm.config).QueryRegions(rm)
}

// QueryRegionRefRef queries the "region_ref_ref" edge of the RegionMaster entity.
func (rm *RegionMaster) QueryRegionRefRef() *FacilityQuery {
	return NewRegionMasterClient(rm.config).QueryRegionRefRef(rm)
}

// Update returns a builder for updating this RegionMaster.
// Note that you need to call RegionMaster.Unwrap() before calling this method if this RegionMaster
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RegionMaster) Update() *RegionMasterUpdateOne {
	return NewRegionMasterClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RegionMaster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RegionMaster) Unwrap() *RegionMaster {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RegionMaster is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RegionMaster) String() string {
	var builder strings.Builder
	builder.WriteString("RegionMaster(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("RegionCode=")
	builder.WriteString(fmt.Sprintf("%v", rm.RegionCode))
	builder.WriteString(", ")
	builder.WriteString("RegionOfficeId=")
	builder.WriteString(rm.RegionOfficeId)
	builder.WriteString(", ")
	builder.WriteString("OfficeType=")
	builder.WriteString(rm.OfficeType)
	builder.WriteString(", ")
	builder.WriteString("RegionOfficeName=")
	builder.WriteString(rm.RegionOfficeName)
	builder.WriteString(", ")
	builder.WriteString("ReportingOfficeType=")
	builder.WriteString(rm.ReportingOfficeType)
	builder.WriteString(", ")
	builder.WriteString("ReportingOfficeCode=")
	builder.WriteString(rm.ReportingOfficeCode)
	builder.WriteString(", ")
	builder.WriteString("EmailID=")
	builder.WriteString(rm.EmailID)
	builder.WriteString(", ")
	builder.WriteString("MobileNumber=")
	builder.WriteString(fmt.Sprintf("%v", rm.MobileNumber))
	builder.WriteString(", ")
	builder.WriteString("CircleCode=")
	builder.WriteString(fmt.Sprintf("%v", rm.CircleCode))
	builder.WriteByte(')')
	return builder.String()
}

// RegionMasters is a parsable slice of RegionMaster.
type RegionMasters []*RegionMaster