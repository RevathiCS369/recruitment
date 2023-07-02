// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EmployeedID holds the value of the "EmployeedID" field.
	EmployeedID string `json:"EmployeedID,omitempty"`
	// IDVerified holds the value of the "IDVerified" field.
	IDVerified bool `json:"IDVerified,omitempty"`
	// IDRemStatus holds the value of the "IDRemStatus" field.
	IDRemStatus bool `json:"IDRemStatus,omitempty"`
	// IDRemarks holds the value of the "IDRemarks" field.
	IDRemarks string `json:"IDRemarks,omitempty"`
	// EmployeedName holds the value of the "EmployeedName" field.
	EmployeedName string `json:"EmployeedName,omitempty"`
	// NameVerified holds the value of the "nameVerified" field.
	NameVerified bool `json:"nameVerified,omitempty"`
	// NameRemStatus holds the value of the "nameRemStatus" field.
	NameRemStatus bool `json:"nameRemStatus,omitempty"`
	// NameRemarks holds the value of the "nameRemarks" field.
	NameRemarks string `json:"nameRemarks,omitempty"`
	// DOB holds the value of the "DOB" field.
	DOB time.Time `json:"DOB,omitempty"`
	// DOBVerified holds the value of the "DOBVerified" field.
	DOBVerified bool `json:"DOBVerified,omitempty"`
	// DOBRemStatus holds the value of the "DOBRemStatus" field.
	DOBRemStatus bool `json:"DOBRemStatus,omitempty"`
	// DOBRemarks holds the value of the "DOBRemarks" field.
	DOBRemarks string `json:"DOBRemarks,omitempty"`
	// Gender holds the value of the "Gender" field.
	Gender user.Gender `json:"Gender,omitempty"`
	// GenderVerified holds the value of the "genderVerified" field.
	GenderVerified bool `json:"genderVerified,omitempty"`
	// GenderRemStatus holds the value of the "genderRemStatus" field.
	GenderRemStatus bool `json:"genderRemStatus,omitempty"`
	// GenderRemarks holds the value of the "genderRemarks" field.
	GenderRemarks string `json:"genderRemarks,omitempty"`
	// Cadreid holds the value of the "Cadreid" field.
	Cadreid int32 `json:"Cadreid,omitempty"`
	// CadreidVerified holds the value of the "cadreidVerified" field.
	CadreidVerified bool `json:"cadreidVerified,omitempty"`
	// CadreidRemStatus holds the value of the "cadreidRemStatus" field.
	CadreidRemStatus bool `json:"cadreidRemStatus,omitempty"`
	// CadreidRemarks holds the value of the "cadreidRemarks" field.
	CadreidRemarks string `json:"cadreidRemarks,omitempty"`
	// OfficeID holds the value of the "OfficeID" field.
	OfficeID int32 `json:"OfficeID,omitempty"`
	// OfficeIDVerified holds the value of the "officeIDVerified" field.
	OfficeIDVerified bool `json:"officeIDVerified,omitempty"`
	// OfficeIDRemStatus holds the value of the "officeIDRemStatus" field.
	OfficeIDRemStatus bool `json:"officeIDRemStatus,omitempty"`
	// OfficeIDRemarks holds the value of the "officeIDRemarks" field.
	OfficeIDRemarks string `json:"officeIDRemarks,omitempty"`
	// PH holds the value of the "PH" field.
	PH bool `json:"PH,omitempty"`
	// PHVerified holds the value of the "PHVerified" field.
	PHVerified bool `json:"PHVerified,omitempty"`
	// PHRemStatus holds the value of the "PHRemStatus" field.
	PHRemStatus bool `json:"PHRemStatus,omitempty"`
	// PHRemarks holds the value of the "PHRemarks" field.
	PHRemarks string `json:"PHRemarks,omitempty"`
	// PHDetails holds the value of the "PHDetails" field.
	PHDetails string `json:"PHDetails,omitempty"`
	// PHDetailsVerified holds the value of the "PHDetailsVerified" field.
	PHDetailsVerified bool `json:"PHDetailsVerified,omitempty"`
	// PHDetailsRemStatus holds the value of the "PHDetailsRemStatus" field.
	PHDetailsRemStatus bool `json:"PHDetailsRemStatus,omitempty"`
	// PHDetailsRemarks holds the value of the "PHDetailsRemarks" field.
	PHDetailsRemarks string `json:"PHDetailsRemarks,omitempty"`
	// APSWorking holds the value of the "APSWorking" field.
	APSWorking bool `json:"APSWorking,omitempty"`
	// APSWorkingVerified holds the value of the "APSWorkingVerified" field.
	APSWorkingVerified bool `json:"APSWorkingVerified,omitempty"`
	// APSWorkingRemStatus holds the value of the "APSWorkingRemStatus" field.
	APSWorkingRemStatus bool `json:"APSWorkingRemStatus,omitempty"`
	// APSWorkingRemarks holds the value of the "APSWorkingRemarks" field.
	APSWorkingRemarks string `json:"APSWorkingRemarks,omitempty"`
	// Profilestatus holds the value of the "profilestatus" field.
	Profilestatus bool `json:"profilestatus,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIDVerified, user.FieldIDRemStatus, user.FieldNameVerified, user.FieldNameRemStatus, user.FieldDOBVerified, user.FieldDOBRemStatus, user.FieldGenderVerified, user.FieldGenderRemStatus, user.FieldCadreidVerified, user.FieldCadreidRemStatus, user.FieldOfficeIDVerified, user.FieldOfficeIDRemStatus, user.FieldPH, user.FieldPHVerified, user.FieldPHRemStatus, user.FieldPHDetailsVerified, user.FieldPHDetailsRemStatus, user.FieldAPSWorking, user.FieldAPSWorkingVerified, user.FieldAPSWorkingRemStatus, user.FieldProfilestatus:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldCadreid, user.FieldOfficeID:
			values[i] = new(sql.NullInt64)
		case user.FieldEmployeedID, user.FieldIDRemarks, user.FieldEmployeedName, user.FieldNameRemarks, user.FieldDOBRemarks, user.FieldGender, user.FieldGenderRemarks, user.FieldCadreidRemarks, user.FieldOfficeIDRemarks, user.FieldPHRemarks, user.FieldPHDetails, user.FieldPHDetailsRemarks, user.FieldAPSWorkingRemarks:
			values[i] = new(sql.NullString)
		case user.FieldDOB:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldEmployeedID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeedID", values[i])
			} else if value.Valid {
				u.EmployeedID = value.String
			}
		case user.FieldIDVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IDVerified", values[i])
			} else if value.Valid {
				u.IDVerified = value.Bool
			}
		case user.FieldIDRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IDRemStatus", values[i])
			} else if value.Valid {
				u.IDRemStatus = value.Bool
			}
		case user.FieldIDRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field IDRemarks", values[i])
			} else if value.Valid {
				u.IDRemarks = value.String
			}
		case user.FieldEmployeedName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EmployeedName", values[i])
			} else if value.Valid {
				u.EmployeedName = value.String
			}
		case user.FieldNameVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field nameVerified", values[i])
			} else if value.Valid {
				u.NameVerified = value.Bool
			}
		case user.FieldNameRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field nameRemStatus", values[i])
			} else if value.Valid {
				u.NameRemStatus = value.Bool
			}
		case user.FieldNameRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nameRemarks", values[i])
			} else if value.Valid {
				u.NameRemarks = value.String
			}
		case user.FieldDOB:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DOB", values[i])
			} else if value.Valid {
				u.DOB = value.Time
			}
		case user.FieldDOBVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field DOBVerified", values[i])
			} else if value.Valid {
				u.DOBVerified = value.Bool
			}
		case user.FieldDOBRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field DOBRemStatus", values[i])
			} else if value.Valid {
				u.DOBRemStatus = value.Bool
			}
		case user.FieldDOBRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DOBRemarks", values[i])
			} else if value.Valid {
				u.DOBRemarks = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Gender", values[i])
			} else if value.Valid {
				u.Gender = user.Gender(value.String)
			}
		case user.FieldGenderVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field genderVerified", values[i])
			} else if value.Valid {
				u.GenderVerified = value.Bool
			}
		case user.FieldGenderRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field genderRemStatus", values[i])
			} else if value.Valid {
				u.GenderRemStatus = value.Bool
			}
		case user.FieldGenderRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field genderRemarks", values[i])
			} else if value.Valid {
				u.GenderRemarks = value.String
			}
		case user.FieldCadreid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Cadreid", values[i])
			} else if value.Valid {
				u.Cadreid = int32(value.Int64)
			}
		case user.FieldCadreidVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cadreidVerified", values[i])
			} else if value.Valid {
				u.CadreidVerified = value.Bool
			}
		case user.FieldCadreidRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cadreidRemStatus", values[i])
			} else if value.Valid {
				u.CadreidRemStatus = value.Bool
			}
		case user.FieldCadreidRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cadreidRemarks", values[i])
			} else if value.Valid {
				u.CadreidRemarks = value.String
			}
		case user.FieldOfficeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field OfficeID", values[i])
			} else if value.Valid {
				u.OfficeID = int32(value.Int64)
			}
		case user.FieldOfficeIDVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field officeIDVerified", values[i])
			} else if value.Valid {
				u.OfficeIDVerified = value.Bool
			}
		case user.FieldOfficeIDRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field officeIDRemStatus", values[i])
			} else if value.Valid {
				u.OfficeIDRemStatus = value.Bool
			}
		case user.FieldOfficeIDRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field officeIDRemarks", values[i])
			} else if value.Valid {
				u.OfficeIDRemarks = value.String
			}
		case user.FieldPH:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field PH", values[i])
			} else if value.Valid {
				u.PH = value.Bool
			}
		case user.FieldPHVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field PHVerified", values[i])
			} else if value.Valid {
				u.PHVerified = value.Bool
			}
		case user.FieldPHRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field PHRemStatus", values[i])
			} else if value.Valid {
				u.PHRemStatus = value.Bool
			}
		case user.FieldPHRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PHRemarks", values[i])
			} else if value.Valid {
				u.PHRemarks = value.String
			}
		case user.FieldPHDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PHDetails", values[i])
			} else if value.Valid {
				u.PHDetails = value.String
			}
		case user.FieldPHDetailsVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field PHDetailsVerified", values[i])
			} else if value.Valid {
				u.PHDetailsVerified = value.Bool
			}
		case user.FieldPHDetailsRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field PHDetailsRemStatus", values[i])
			} else if value.Valid {
				u.PHDetailsRemStatus = value.Bool
			}
		case user.FieldPHDetailsRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PHDetailsRemarks", values[i])
			} else if value.Valid {
				u.PHDetailsRemarks = value.String
			}
		case user.FieldAPSWorking:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field APSWorking", values[i])
			} else if value.Valid {
				u.APSWorking = value.Bool
			}
		case user.FieldAPSWorkingVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field APSWorkingVerified", values[i])
			} else if value.Valid {
				u.APSWorkingVerified = value.Bool
			}
		case user.FieldAPSWorkingRemStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field APSWorkingRemStatus", values[i])
			} else if value.Valid {
				u.APSWorkingRemStatus = value.Bool
			}
		case user.FieldAPSWorkingRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field APSWorkingRemarks", values[i])
			} else if value.Valid {
				u.APSWorkingRemarks = value.String
			}
		case user.FieldProfilestatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field profilestatus", values[i])
			} else if value.Valid {
				u.Profilestatus = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("EmployeedID=")
	builder.WriteString(u.EmployeedID)
	builder.WriteString(", ")
	builder.WriteString("IDVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.IDVerified))
	builder.WriteString(", ")
	builder.WriteString("IDRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.IDRemStatus))
	builder.WriteString(", ")
	builder.WriteString("IDRemarks=")
	builder.WriteString(u.IDRemarks)
	builder.WriteString(", ")
	builder.WriteString("EmployeedName=")
	builder.WriteString(u.EmployeedName)
	builder.WriteString(", ")
	builder.WriteString("nameVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.NameVerified))
	builder.WriteString(", ")
	builder.WriteString("nameRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.NameRemStatus))
	builder.WriteString(", ")
	builder.WriteString("nameRemarks=")
	builder.WriteString(u.NameRemarks)
	builder.WriteString(", ")
	builder.WriteString("DOB=")
	builder.WriteString(u.DOB.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("DOBVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.DOBVerified))
	builder.WriteString(", ")
	builder.WriteString("DOBRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.DOBRemStatus))
	builder.WriteString(", ")
	builder.WriteString("DOBRemarks=")
	builder.WriteString(u.DOBRemarks)
	builder.WriteString(", ")
	builder.WriteString("Gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	builder.WriteString("genderVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.GenderVerified))
	builder.WriteString(", ")
	builder.WriteString("genderRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.GenderRemStatus))
	builder.WriteString(", ")
	builder.WriteString("genderRemarks=")
	builder.WriteString(u.GenderRemarks)
	builder.WriteString(", ")
	builder.WriteString("Cadreid=")
	builder.WriteString(fmt.Sprintf("%v", u.Cadreid))
	builder.WriteString(", ")
	builder.WriteString("cadreidVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.CadreidVerified))
	builder.WriteString(", ")
	builder.WriteString("cadreidRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.CadreidRemStatus))
	builder.WriteString(", ")
	builder.WriteString("cadreidRemarks=")
	builder.WriteString(u.CadreidRemarks)
	builder.WriteString(", ")
	builder.WriteString("OfficeID=")
	builder.WriteString(fmt.Sprintf("%v", u.OfficeID))
	builder.WriteString(", ")
	builder.WriteString("officeIDVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.OfficeIDVerified))
	builder.WriteString(", ")
	builder.WriteString("officeIDRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.OfficeIDRemStatus))
	builder.WriteString(", ")
	builder.WriteString("officeIDRemarks=")
	builder.WriteString(u.OfficeIDRemarks)
	builder.WriteString(", ")
	builder.WriteString("PH=")
	builder.WriteString(fmt.Sprintf("%v", u.PH))
	builder.WriteString(", ")
	builder.WriteString("PHVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.PHVerified))
	builder.WriteString(", ")
	builder.WriteString("PHRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.PHRemStatus))
	builder.WriteString(", ")
	builder.WriteString("PHRemarks=")
	builder.WriteString(u.PHRemarks)
	builder.WriteString(", ")
	builder.WriteString("PHDetails=")
	builder.WriteString(u.PHDetails)
	builder.WriteString(", ")
	builder.WriteString("PHDetailsVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.PHDetailsVerified))
	builder.WriteString(", ")
	builder.WriteString("PHDetailsRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.PHDetailsRemStatus))
	builder.WriteString(", ")
	builder.WriteString("PHDetailsRemarks=")
	builder.WriteString(u.PHDetailsRemarks)
	builder.WriteString(", ")
	builder.WriteString("APSWorking=")
	builder.WriteString(fmt.Sprintf("%v", u.APSWorking))
	builder.WriteString(", ")
	builder.WriteString("APSWorkingVerified=")
	builder.WriteString(fmt.Sprintf("%v", u.APSWorkingVerified))
	builder.WriteString(", ")
	builder.WriteString("APSWorkingRemStatus=")
	builder.WriteString(fmt.Sprintf("%v", u.APSWorkingRemStatus))
	builder.WriteString(", ")
	builder.WriteString("APSWorkingRemarks=")
	builder.WriteString(u.APSWorkingRemarks)
	builder.WriteString(", ")
	builder.WriteString("profilestatus=")
	builder.WriteString(fmt.Sprintf("%v", u.Profilestatus))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User