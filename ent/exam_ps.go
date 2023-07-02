// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"recruit/ent/exam_ps"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Exam_PS is the model entity for the Exam_PS schema.
type Exam_PS struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ExamNameCode holds the value of the "ExamNameCode" field.
	ExamNameCode string `json:"ExamNameCode,omitempty"`
	// ExamName holds the value of the "ExamName" field.
	ExamName string `json:"ExamName,omitempty"`
	// ExamType holds the value of the "ExamType" field.
	ExamType string `json:"ExamType,omitempty"`
	// NotificationCode holds the value of the "NotificationCode" field.
	NotificationCode int32 `json:"NotificationCode,omitempty"`
	// ConductedBy holds the value of the "ConductedBy" field.
	ConductedBy string `json:"ConductedBy,omitempty"`
	// NodalOffice holds the value of the "NodalOffice" field.
	NodalOffice string `json:"NodalOffice,omitempty"`
	// CalendarCode holds the value of the "CalendarCode" field.
	CalendarCode int32 `json:"CalendarCode,omitempty"`
	// PaperCode holds the value of the "PaperCode" field.
	PaperCode int32 `json:"PaperCode,omitempty"`
	// EligibleCadre holds the value of the "EligibleCadre" field.
	EligibleCadre string `json:"EligibleCadre,omitempty"`
	// EligiblePost1 holds the value of the "EligiblePost1" field.
	EligiblePost1 string `json:"EligiblePost1,omitempty"`
	// EligiblePost2 holds the value of the "EligiblePost2" field.
	EligiblePost2 string `json:"EligiblePost2,omitempty"`
	// EligiblePost3 holds the value of the "EligiblePost3" field.
	EligiblePost3 string `json:"EligiblePost3,omitempty"`
	// EligiblePost4 holds the value of the "EligiblePost4" field.
	EligiblePost4 string `json:"EligiblePost4,omitempty"`
	// EligiblePost5 holds the value of the "EligiblePost5" field.
	EligiblePost5 string `json:"EligiblePost5,omitempty"`
	// ExamPost1 holds the value of the "ExamPost1" field.
	ExamPost1 string `json:"ExamPost1,omitempty"`
	// ExamPost2 holds the value of the "ExamPost2" field.
	ExamPost2 string `json:"ExamPost2,omitempty"`
	// ExamPost3 holds the value of the "ExamPost3" field.
	ExamPost3 string `json:"ExamPost3,omitempty"`
	// ExamPost4 holds the value of the "ExamPost4" field.
	ExamPost4 string `json:"ExamPost4,omitempty"`
	// ExamPost5 holds the value of the "ExamPost5" field.
	ExamPost5 string `json:"ExamPost5,omitempty"`
	// EducationCriteria holds the value of the "EducationCriteria" field.
	EducationCriteria string `json:"EducationCriteria,omitempty"`
	// CategoryAgeLimitGEN holds the value of the "CategoryAgeLimitGEN" field.
	CategoryAgeLimitGEN string `json:"CategoryAgeLimitGEN,omitempty"`
	// CategoryAgeLimitSC holds the value of the "CategoryAgeLimitSC" field.
	CategoryAgeLimitSC string `json:"CategoryAgeLimitSC,omitempty"`
	// CategoryAgeLimitST holds the value of the "CategoryAgeLimitST" field.
	CategoryAgeLimitST string `json:"CategoryAgeLimitST,omitempty"`
	// ServiceYears holds the value of the "ServiceYears" field.
	ServiceYears string `json:"ServiceYears,omitempty"`
	// DrivingLicenseRequired holds the value of the "DrivingLicenseRequired" field.
	DrivingLicenseRequired string `json:"DrivingLicenseRequired,omitempty"`
	// ExamPaperCode holds the value of the "ExamPaperCode" field.
	ExamPaperCode string `json:"ExamPaperCode,omitempty"`
	// ExamPaper1 holds the value of the "ExamPaper1" field.
	ExamPaper1 string `json:"ExamPaper1,omitempty"`
	// ExamPaper2 holds the value of the "ExamPaper2" field.
	ExamPaper2 string `json:"ExamPaper2,omitempty"`
	// ExamPaper3 holds the value of the "ExamPaper3" field.
	ExamPaper3 string `json:"ExamPaper3,omitempty"`
	// ExamPaper4 holds the value of the "ExamPaper4" field.
	ExamPaper4 string `json:"ExamPaper4,omitempty"`
	// ExamPaper5 holds the value of the "ExamPaper5" field.
	ExamPaper5 string `json:"ExamPaper5,omitempty"`
	// ExamPaper6 holds the value of the "ExamPaper6" field.
	ExamPaper6 string `json:"ExamPaper6,omitempty"`
	// PayLevelEligibilty holds the value of the "PayLevelEligibilty" field.
	PayLevelEligibilty string `json:"PayLevelEligibilty,omitempty"`
	// CategoryMinMarksSCSTPH holds the value of the "CategoryMinMarksSCSTPH" field.
	CategoryMinMarksSCSTPH string `json:"CategoryMinMarksSCSTPH,omitempty"`
	// CategoryMinMarksGENOBC holds the value of the "CategoryMinMarksGENOBC" field.
	CategoryMinMarksGENOBC string `json:"CategoryMinMarksGENOBC,omitempty"`
	// LocalLanguageAllowed holds the value of the "LocalLanguageAllowed" field.
	LocalLanguageAllowed string `json:"LocalLanguageAllowed,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy string `json:"UpdatedBy,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Exam_PSQuery when eager-loading is set.
	Edges                                  Exam_PSEdges `json:"edges"`
	exam_calendar_examcal_ps_ref           *int32
	exam_papers_papers_ps_ref              *int32
	exam_applications_ps_exam_appln_ps_ref *int64
	notification_notifications_ps          *int32
	user_master_users_ps_type              *int64
	selectValues                           sql.SelectValues
}

// Exam_PSEdges holds the relations/edges for other nodes in the graph.
type Exam_PSEdges struct {
	// ExamcalPsRef holds the value of the examcal_ps_ref edge.
	ExamcalPsRef []*ExamCalendar `json:"examcal_ps_ref,omitempty"`
	// PapersPsRef holds the value of the papers_ps_ref edge.
	PapersPsRef []*ExamPapers `json:"papers_ps_ref,omitempty"`
	// UsersPsType holds the value of the users_ps_type edge.
	UsersPsType []*UserMaster `json:"users_ps_type,omitempty"`
	// ExamApplnPSRef holds the value of the ExamAppln_PS_Ref edge.
	ExamApplnPSRef []*Exam_Applications_PS `json:"ExamAppln_PS_Ref,omitempty"`
	// NotificationsPs holds the value of the notifications_ps edge.
	NotificationsPs []*Notification `json:"notifications_ps,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ExamcalPsRefOrErr returns the ExamcalPsRef value or an error if the edge
// was not loaded in eager-loading.
func (e Exam_PSEdges) ExamcalPsRefOrErr() ([]*ExamCalendar, error) {
	if e.loadedTypes[0] {
		return e.ExamcalPsRef, nil
	}
	return nil, &NotLoadedError{edge: "examcal_ps_ref"}
}

// PapersPsRefOrErr returns the PapersPsRef value or an error if the edge
// was not loaded in eager-loading.
func (e Exam_PSEdges) PapersPsRefOrErr() ([]*ExamPapers, error) {
	if e.loadedTypes[1] {
		return e.PapersPsRef, nil
	}
	return nil, &NotLoadedError{edge: "papers_ps_ref"}
}

// UsersPsTypeOrErr returns the UsersPsType value or an error if the edge
// was not loaded in eager-loading.
func (e Exam_PSEdges) UsersPsTypeOrErr() ([]*UserMaster, error) {
	if e.loadedTypes[2] {
		return e.UsersPsType, nil
	}
	return nil, &NotLoadedError{edge: "users_ps_type"}
}

// ExamApplnPSRefOrErr returns the ExamApplnPSRef value or an error if the edge
// was not loaded in eager-loading.
func (e Exam_PSEdges) ExamApplnPSRefOrErr() ([]*Exam_Applications_PS, error) {
	if e.loadedTypes[3] {
		return e.ExamApplnPSRef, nil
	}
	return nil, &NotLoadedError{edge: "ExamAppln_PS_Ref"}
}

// NotificationsPsOrErr returns the NotificationsPs value or an error if the edge
// was not loaded in eager-loading.
func (e Exam_PSEdges) NotificationsPsOrErr() ([]*Notification, error) {
	if e.loadedTypes[4] {
		return e.NotificationsPs, nil
	}
	return nil, &NotLoadedError{edge: "notifications_ps"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exam_PS) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exam_ps.FieldID, exam_ps.FieldNotificationCode, exam_ps.FieldCalendarCode, exam_ps.FieldPaperCode:
			values[i] = new(sql.NullInt64)
		case exam_ps.FieldExamNameCode, exam_ps.FieldExamName, exam_ps.FieldExamType, exam_ps.FieldConductedBy, exam_ps.FieldNodalOffice, exam_ps.FieldEligibleCadre, exam_ps.FieldEligiblePost1, exam_ps.FieldEligiblePost2, exam_ps.FieldEligiblePost3, exam_ps.FieldEligiblePost4, exam_ps.FieldEligiblePost5, exam_ps.FieldExamPost1, exam_ps.FieldExamPost2, exam_ps.FieldExamPost3, exam_ps.FieldExamPost4, exam_ps.FieldExamPost5, exam_ps.FieldEducationCriteria, exam_ps.FieldCategoryAgeLimitGEN, exam_ps.FieldCategoryAgeLimitSC, exam_ps.FieldCategoryAgeLimitST, exam_ps.FieldServiceYears, exam_ps.FieldDrivingLicenseRequired, exam_ps.FieldExamPaperCode, exam_ps.FieldExamPaper1, exam_ps.FieldExamPaper2, exam_ps.FieldExamPaper3, exam_ps.FieldExamPaper4, exam_ps.FieldExamPaper5, exam_ps.FieldExamPaper6, exam_ps.FieldPayLevelEligibilty, exam_ps.FieldCategoryMinMarksSCSTPH, exam_ps.FieldCategoryMinMarksGENOBC, exam_ps.FieldLocalLanguageAllowed, exam_ps.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case exam_ps.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case exam_ps.ForeignKeys[0]: // exam_calendar_examcal_ps_ref
			values[i] = new(sql.NullInt64)
		case exam_ps.ForeignKeys[1]: // exam_papers_papers_ps_ref
			values[i] = new(sql.NullInt64)
		case exam_ps.ForeignKeys[2]: // exam_applications_ps_exam_appln_ps_ref
			values[i] = new(sql.NullInt64)
		case exam_ps.ForeignKeys[3]: // notification_notifications_ps
			values[i] = new(sql.NullInt64)
		case exam_ps.ForeignKeys[4]: // user_master_users_ps_type
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exam_PS fields.
func (ep *Exam_PS) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exam_ps.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ep.ID = int32(value.Int64)
		case exam_ps.FieldExamNameCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamNameCode", values[i])
			} else if value.Valid {
				ep.ExamNameCode = value.String
			}
		case exam_ps.FieldExamName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamName", values[i])
			} else if value.Valid {
				ep.ExamName = value.String
			}
		case exam_ps.FieldExamType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamType", values[i])
			} else if value.Valid {
				ep.ExamType = value.String
			}
		case exam_ps.FieldNotificationCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field NotificationCode", values[i])
			} else if value.Valid {
				ep.NotificationCode = int32(value.Int64)
			}
		case exam_ps.FieldConductedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ConductedBy", values[i])
			} else if value.Valid {
				ep.ConductedBy = value.String
			}
		case exam_ps.FieldNodalOffice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field NodalOffice", values[i])
			} else if value.Valid {
				ep.NodalOffice = value.String
			}
		case exam_ps.FieldCalendarCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CalendarCode", values[i])
			} else if value.Valid {
				ep.CalendarCode = int32(value.Int64)
			}
		case exam_ps.FieldPaperCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PaperCode", values[i])
			} else if value.Valid {
				ep.PaperCode = int32(value.Int64)
			}
		case exam_ps.FieldEligibleCadre:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligibleCadre", values[i])
			} else if value.Valid {
				ep.EligibleCadre = value.String
			}
		case exam_ps.FieldEligiblePost1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligiblePost1", values[i])
			} else if value.Valid {
				ep.EligiblePost1 = value.String
			}
		case exam_ps.FieldEligiblePost2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligiblePost2", values[i])
			} else if value.Valid {
				ep.EligiblePost2 = value.String
			}
		case exam_ps.FieldEligiblePost3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligiblePost3", values[i])
			} else if value.Valid {
				ep.EligiblePost3 = value.String
			}
		case exam_ps.FieldEligiblePost4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligiblePost4", values[i])
			} else if value.Valid {
				ep.EligiblePost4 = value.String
			}
		case exam_ps.FieldEligiblePost5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EligiblePost5", values[i])
			} else if value.Valid {
				ep.EligiblePost5 = value.String
			}
		case exam_ps.FieldExamPost1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPost1", values[i])
			} else if value.Valid {
				ep.ExamPost1 = value.String
			}
		case exam_ps.FieldExamPost2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPost2", values[i])
			} else if value.Valid {
				ep.ExamPost2 = value.String
			}
		case exam_ps.FieldExamPost3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPost3", values[i])
			} else if value.Valid {
				ep.ExamPost3 = value.String
			}
		case exam_ps.FieldExamPost4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPost4", values[i])
			} else if value.Valid {
				ep.ExamPost4 = value.String
			}
		case exam_ps.FieldExamPost5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPost5", values[i])
			} else if value.Valid {
				ep.ExamPost5 = value.String
			}
		case exam_ps.FieldEducationCriteria:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field EducationCriteria", values[i])
			} else if value.Valid {
				ep.EducationCriteria = value.String
			}
		case exam_ps.FieldCategoryAgeLimitGEN:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryAgeLimitGEN", values[i])
			} else if value.Valid {
				ep.CategoryAgeLimitGEN = value.String
			}
		case exam_ps.FieldCategoryAgeLimitSC:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryAgeLimitSC", values[i])
			} else if value.Valid {
				ep.CategoryAgeLimitSC = value.String
			}
		case exam_ps.FieldCategoryAgeLimitST:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryAgeLimitST", values[i])
			} else if value.Valid {
				ep.CategoryAgeLimitST = value.String
			}
		case exam_ps.FieldServiceYears:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ServiceYears", values[i])
			} else if value.Valid {
				ep.ServiceYears = value.String
			}
		case exam_ps.FieldDrivingLicenseRequired:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DrivingLicenseRequired", values[i])
			} else if value.Valid {
				ep.DrivingLicenseRequired = value.String
			}
		case exam_ps.FieldExamPaperCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaperCode", values[i])
			} else if value.Valid {
				ep.ExamPaperCode = value.String
			}
		case exam_ps.FieldExamPaper1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper1", values[i])
			} else if value.Valid {
				ep.ExamPaper1 = value.String
			}
		case exam_ps.FieldExamPaper2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper2", values[i])
			} else if value.Valid {
				ep.ExamPaper2 = value.String
			}
		case exam_ps.FieldExamPaper3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper3", values[i])
			} else if value.Valid {
				ep.ExamPaper3 = value.String
			}
		case exam_ps.FieldExamPaper4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper4", values[i])
			} else if value.Valid {
				ep.ExamPaper4 = value.String
			}
		case exam_ps.FieldExamPaper5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper5", values[i])
			} else if value.Valid {
				ep.ExamPaper5 = value.String
			}
		case exam_ps.FieldExamPaper6:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ExamPaper6", values[i])
			} else if value.Valid {
				ep.ExamPaper6 = value.String
			}
		case exam_ps.FieldPayLevelEligibilty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PayLevelEligibilty", values[i])
			} else if value.Valid {
				ep.PayLevelEligibilty = value.String
			}
		case exam_ps.FieldCategoryMinMarksSCSTPH:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryMinMarksSCSTPH", values[i])
			} else if value.Valid {
				ep.CategoryMinMarksSCSTPH = value.String
			}
		case exam_ps.FieldCategoryMinMarksGENOBC:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CategoryMinMarksGENOBC", values[i])
			} else if value.Valid {
				ep.CategoryMinMarksGENOBC = value.String
			}
		case exam_ps.FieldLocalLanguageAllowed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LocalLanguageAllowed", values[i])
			} else if value.Valid {
				ep.LocalLanguageAllowed = value.String
			}
		case exam_ps.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				ep.UpdatedAt = value.Time
			}
		case exam_ps.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				ep.UpdatedBy = value.String
			}
		case exam_ps.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_calendar_examcal_ps_ref", value)
			} else if value.Valid {
				ep.exam_calendar_examcal_ps_ref = new(int32)
				*ep.exam_calendar_examcal_ps_ref = int32(value.Int64)
			}
		case exam_ps.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_papers_papers_ps_ref", value)
			} else if value.Valid {
				ep.exam_papers_papers_ps_ref = new(int32)
				*ep.exam_papers_papers_ps_ref = int32(value.Int64)
			}
		case exam_ps.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_applications_ps_exam_appln_ps_ref", value)
			} else if value.Valid {
				ep.exam_applications_ps_exam_appln_ps_ref = new(int64)
				*ep.exam_applications_ps_exam_appln_ps_ref = int64(value.Int64)
			}
		case exam_ps.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field notification_notifications_ps", value)
			} else if value.Valid {
				ep.notification_notifications_ps = new(int32)
				*ep.notification_notifications_ps = int32(value.Int64)
			}
		case exam_ps.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_master_users_ps_type", value)
			} else if value.Valid {
				ep.user_master_users_ps_type = new(int64)
				*ep.user_master_users_ps_type = int64(value.Int64)
			}
		default:
			ep.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Exam_PS.
// This includes values selected through modifiers, order, etc.
func (ep *Exam_PS) Value(name string) (ent.Value, error) {
	return ep.selectValues.Get(name)
}

// QueryExamcalPsRef queries the "examcal_ps_ref" edge of the Exam_PS entity.
func (ep *Exam_PS) QueryExamcalPsRef() *ExamCalendarQuery {
	return NewExamPSClient(ep.config).QueryExamcalPsRef(ep)
}

// QueryPapersPsRef queries the "papers_ps_ref" edge of the Exam_PS entity.
func (ep *Exam_PS) QueryPapersPsRef() *ExamPapersQuery {
	return NewExamPSClient(ep.config).QueryPapersPsRef(ep)
}

// QueryUsersPsType queries the "users_ps_type" edge of the Exam_PS entity.
func (ep *Exam_PS) QueryUsersPsType() *UserMasterQuery {
	return NewExamPSClient(ep.config).QueryUsersPsType(ep)
}

// QueryExamApplnPSRef queries the "ExamAppln_PS_Ref" edge of the Exam_PS entity.
func (ep *Exam_PS) QueryExamApplnPSRef() *ExamApplicationsPSQuery {
	return NewExamPSClient(ep.config).QueryExamApplnPSRef(ep)
}

// QueryNotificationsPs queries the "notifications_ps" edge of the Exam_PS entity.
func (ep *Exam_PS) QueryNotificationsPs() *NotificationQuery {
	return NewExamPSClient(ep.config).QueryNotificationsPs(ep)
}

// Update returns a builder for updating this Exam_PS.
// Note that you need to call Exam_PS.Unwrap() before calling this method if this Exam_PS
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *Exam_PS) Update() *ExamPSUpdateOne {
	return NewExamPSClient(ep.config).UpdateOne(ep)
}

// Unwrap unwraps the Exam_PS entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *Exam_PS) Unwrap() *Exam_PS {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exam_PS is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *Exam_PS) String() string {
	var builder strings.Builder
	builder.WriteString("Exam_PS(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("ExamNameCode=")
	builder.WriteString(ep.ExamNameCode)
	builder.WriteString(", ")
	builder.WriteString("ExamName=")
	builder.WriteString(ep.ExamName)
	builder.WriteString(", ")
	builder.WriteString("ExamType=")
	builder.WriteString(ep.ExamType)
	builder.WriteString(", ")
	builder.WriteString("NotificationCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.NotificationCode))
	builder.WriteString(", ")
	builder.WriteString("ConductedBy=")
	builder.WriteString(ep.ConductedBy)
	builder.WriteString(", ")
	builder.WriteString("NodalOffice=")
	builder.WriteString(ep.NodalOffice)
	builder.WriteString(", ")
	builder.WriteString("CalendarCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.CalendarCode))
	builder.WriteString(", ")
	builder.WriteString("PaperCode=")
	builder.WriteString(fmt.Sprintf("%v", ep.PaperCode))
	builder.WriteString(", ")
	builder.WriteString("EligibleCadre=")
	builder.WriteString(ep.EligibleCadre)
	builder.WriteString(", ")
	builder.WriteString("EligiblePost1=")
	builder.WriteString(ep.EligiblePost1)
	builder.WriteString(", ")
	builder.WriteString("EligiblePost2=")
	builder.WriteString(ep.EligiblePost2)
	builder.WriteString(", ")
	builder.WriteString("EligiblePost3=")
	builder.WriteString(ep.EligiblePost3)
	builder.WriteString(", ")
	builder.WriteString("EligiblePost4=")
	builder.WriteString(ep.EligiblePost4)
	builder.WriteString(", ")
	builder.WriteString("EligiblePost5=")
	builder.WriteString(ep.EligiblePost5)
	builder.WriteString(", ")
	builder.WriteString("ExamPost1=")
	builder.WriteString(ep.ExamPost1)
	builder.WriteString(", ")
	builder.WriteString("ExamPost2=")
	builder.WriteString(ep.ExamPost2)
	builder.WriteString(", ")
	builder.WriteString("ExamPost3=")
	builder.WriteString(ep.ExamPost3)
	builder.WriteString(", ")
	builder.WriteString("ExamPost4=")
	builder.WriteString(ep.ExamPost4)
	builder.WriteString(", ")
	builder.WriteString("ExamPost5=")
	builder.WriteString(ep.ExamPost5)
	builder.WriteString(", ")
	builder.WriteString("EducationCriteria=")
	builder.WriteString(ep.EducationCriteria)
	builder.WriteString(", ")
	builder.WriteString("CategoryAgeLimitGEN=")
	builder.WriteString(ep.CategoryAgeLimitGEN)
	builder.WriteString(", ")
	builder.WriteString("CategoryAgeLimitSC=")
	builder.WriteString(ep.CategoryAgeLimitSC)
	builder.WriteString(", ")
	builder.WriteString("CategoryAgeLimitST=")
	builder.WriteString(ep.CategoryAgeLimitST)
	builder.WriteString(", ")
	builder.WriteString("ServiceYears=")
	builder.WriteString(ep.ServiceYears)
	builder.WriteString(", ")
	builder.WriteString("DrivingLicenseRequired=")
	builder.WriteString(ep.DrivingLicenseRequired)
	builder.WriteString(", ")
	builder.WriteString("ExamPaperCode=")
	builder.WriteString(ep.ExamPaperCode)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper1=")
	builder.WriteString(ep.ExamPaper1)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper2=")
	builder.WriteString(ep.ExamPaper2)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper3=")
	builder.WriteString(ep.ExamPaper3)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper4=")
	builder.WriteString(ep.ExamPaper4)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper5=")
	builder.WriteString(ep.ExamPaper5)
	builder.WriteString(", ")
	builder.WriteString("ExamPaper6=")
	builder.WriteString(ep.ExamPaper6)
	builder.WriteString(", ")
	builder.WriteString("PayLevelEligibilty=")
	builder.WriteString(ep.PayLevelEligibilty)
	builder.WriteString(", ")
	builder.WriteString("CategoryMinMarksSCSTPH=")
	builder.WriteString(ep.CategoryMinMarksSCSTPH)
	builder.WriteString(", ")
	builder.WriteString("CategoryMinMarksGENOBC=")
	builder.WriteString(ep.CategoryMinMarksGENOBC)
	builder.WriteString(", ")
	builder.WriteString("LocalLanguageAllowed=")
	builder.WriteString(ep.LocalLanguageAllowed)
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(ep.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedBy=")
	builder.WriteString(ep.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Exam_PSs is a parsable slice of Exam_PS.
type Exam_PSs []*Exam_PS