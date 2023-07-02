// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/exam_applications_ps"
	"recruit/ent/exam_pa"
	"recruit/ent/examcalendar"
	"recruit/ent/exampapers"
	"recruit/ent/notification"
	"recruit/ent/usermaster"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamPACreate is the builder for creating a Exam_PA entity.
type ExamPACreate struct {
	config
	mutation *ExamPAMutation
	hooks    []Hook
}

// SetExamNameCode sets the "ExamNameCode" field.
func (epc *ExamPACreate) SetExamNameCode(s string) *ExamPACreate {
	epc.mutation.SetExamNameCode(s)
	return epc
}

// SetNillableExamNameCode sets the "ExamNameCode" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamNameCode(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamNameCode(*s)
	}
	return epc
}

// SetExamName sets the "ExamName" field.
func (epc *ExamPACreate) SetExamName(s string) *ExamPACreate {
	epc.mutation.SetExamName(s)
	return epc
}

// SetExamType sets the "ExamType" field.
func (epc *ExamPACreate) SetExamType(s string) *ExamPACreate {
	epc.mutation.SetExamType(s)
	return epc
}

// SetNotificationCode sets the "NotificationCode" field.
func (epc *ExamPACreate) SetNotificationCode(i int32) *ExamPACreate {
	epc.mutation.SetNotificationCode(i)
	return epc
}

// SetNillableNotificationCode sets the "NotificationCode" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableNotificationCode(i *int32) *ExamPACreate {
	if i != nil {
		epc.SetNotificationCode(*i)
	}
	return epc
}

// SetConductedBy sets the "ConductedBy" field.
func (epc *ExamPACreate) SetConductedBy(s string) *ExamPACreate {
	epc.mutation.SetConductedBy(s)
	return epc
}

// SetNodalOffice sets the "NodalOffice" field.
func (epc *ExamPACreate) SetNodalOffice(s string) *ExamPACreate {
	epc.mutation.SetNodalOffice(s)
	return epc
}

// SetNillableNodalOffice sets the "NodalOffice" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableNodalOffice(s *string) *ExamPACreate {
	if s != nil {
		epc.SetNodalOffice(*s)
	}
	return epc
}

// SetCalendarCode sets the "CalendarCode" field.
func (epc *ExamPACreate) SetCalendarCode(i int32) *ExamPACreate {
	epc.mutation.SetCalendarCode(i)
	return epc
}

// SetNillableCalendarCode sets the "CalendarCode" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCalendarCode(i *int32) *ExamPACreate {
	if i != nil {
		epc.SetCalendarCode(*i)
	}
	return epc
}

// SetPaperCode sets the "PaperCode" field.
func (epc *ExamPACreate) SetPaperCode(i int32) *ExamPACreate {
	epc.mutation.SetPaperCode(i)
	return epc
}

// SetNillablePaperCode sets the "PaperCode" field if the given value is not nil.
func (epc *ExamPACreate) SetNillablePaperCode(i *int32) *ExamPACreate {
	if i != nil {
		epc.SetPaperCode(*i)
	}
	return epc
}

// SetEligibleCadre sets the "EligibleCadre" field.
func (epc *ExamPACreate) SetEligibleCadre(s string) *ExamPACreate {
	epc.mutation.SetEligibleCadre(s)
	return epc
}

// SetNillableEligibleCadre sets the "EligibleCadre" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligibleCadre(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligibleCadre(*s)
	}
	return epc
}

// SetEligiblePost1 sets the "EligiblePost1" field.
func (epc *ExamPACreate) SetEligiblePost1(s string) *ExamPACreate {
	epc.mutation.SetEligiblePost1(s)
	return epc
}

// SetNillableEligiblePost1 sets the "EligiblePost1" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligiblePost1(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligiblePost1(*s)
	}
	return epc
}

// SetEligiblePost2 sets the "EligiblePost2" field.
func (epc *ExamPACreate) SetEligiblePost2(s string) *ExamPACreate {
	epc.mutation.SetEligiblePost2(s)
	return epc
}

// SetNillableEligiblePost2 sets the "EligiblePost2" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligiblePost2(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligiblePost2(*s)
	}
	return epc
}

// SetEligiblePost3 sets the "EligiblePost3" field.
func (epc *ExamPACreate) SetEligiblePost3(s string) *ExamPACreate {
	epc.mutation.SetEligiblePost3(s)
	return epc
}

// SetNillableEligiblePost3 sets the "EligiblePost3" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligiblePost3(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligiblePost3(*s)
	}
	return epc
}

// SetEligiblePost4 sets the "EligiblePost4" field.
func (epc *ExamPACreate) SetEligiblePost4(s string) *ExamPACreate {
	epc.mutation.SetEligiblePost4(s)
	return epc
}

// SetNillableEligiblePost4 sets the "EligiblePost4" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligiblePost4(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligiblePost4(*s)
	}
	return epc
}

// SetEligiblePost5 sets the "EligiblePost5" field.
func (epc *ExamPACreate) SetEligiblePost5(s string) *ExamPACreate {
	epc.mutation.SetEligiblePost5(s)
	return epc
}

// SetNillableEligiblePost5 sets the "EligiblePost5" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEligiblePost5(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEligiblePost5(*s)
	}
	return epc
}

// SetExamPost1 sets the "ExamPost1" field.
func (epc *ExamPACreate) SetExamPost1(s string) *ExamPACreate {
	epc.mutation.SetExamPost1(s)
	return epc
}

// SetNillableExamPost1 sets the "ExamPost1" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPost1(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPost1(*s)
	}
	return epc
}

// SetExamPost2 sets the "ExamPost2" field.
func (epc *ExamPACreate) SetExamPost2(s string) *ExamPACreate {
	epc.mutation.SetExamPost2(s)
	return epc
}

// SetNillableExamPost2 sets the "ExamPost2" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPost2(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPost2(*s)
	}
	return epc
}

// SetExamPost3 sets the "ExamPost3" field.
func (epc *ExamPACreate) SetExamPost3(s string) *ExamPACreate {
	epc.mutation.SetExamPost3(s)
	return epc
}

// SetNillableExamPost3 sets the "ExamPost3" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPost3(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPost3(*s)
	}
	return epc
}

// SetExamPost4 sets the "ExamPost4" field.
func (epc *ExamPACreate) SetExamPost4(s string) *ExamPACreate {
	epc.mutation.SetExamPost4(s)
	return epc
}

// SetNillableExamPost4 sets the "ExamPost4" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPost4(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPost4(*s)
	}
	return epc
}

// SetExamPost5 sets the "ExamPost5" field.
func (epc *ExamPACreate) SetExamPost5(s string) *ExamPACreate {
	epc.mutation.SetExamPost5(s)
	return epc
}

// SetNillableExamPost5 sets the "ExamPost5" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPost5(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPost5(*s)
	}
	return epc
}

// SetEducationCriteria sets the "EducationCriteria" field.
func (epc *ExamPACreate) SetEducationCriteria(s string) *ExamPACreate {
	epc.mutation.SetEducationCriteria(s)
	return epc
}

// SetNillableEducationCriteria sets the "EducationCriteria" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableEducationCriteria(s *string) *ExamPACreate {
	if s != nil {
		epc.SetEducationCriteria(*s)
	}
	return epc
}

// SetCategoryAgeLimitGEN sets the "CategoryAgeLimitGEN" field.
func (epc *ExamPACreate) SetCategoryAgeLimitGEN(s string) *ExamPACreate {
	epc.mutation.SetCategoryAgeLimitGEN(s)
	return epc
}

// SetNillableCategoryAgeLimitGEN sets the "CategoryAgeLimitGEN" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCategoryAgeLimitGEN(s *string) *ExamPACreate {
	if s != nil {
		epc.SetCategoryAgeLimitGEN(*s)
	}
	return epc
}

// SetCategoryAgeLimitSC sets the "CategoryAgeLimitSC" field.
func (epc *ExamPACreate) SetCategoryAgeLimitSC(s string) *ExamPACreate {
	epc.mutation.SetCategoryAgeLimitSC(s)
	return epc
}

// SetNillableCategoryAgeLimitSC sets the "CategoryAgeLimitSC" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCategoryAgeLimitSC(s *string) *ExamPACreate {
	if s != nil {
		epc.SetCategoryAgeLimitSC(*s)
	}
	return epc
}

// SetCategoryAgeLimitST sets the "CategoryAgeLimitST" field.
func (epc *ExamPACreate) SetCategoryAgeLimitST(s string) *ExamPACreate {
	epc.mutation.SetCategoryAgeLimitST(s)
	return epc
}

// SetNillableCategoryAgeLimitST sets the "CategoryAgeLimitST" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCategoryAgeLimitST(s *string) *ExamPACreate {
	if s != nil {
		epc.SetCategoryAgeLimitST(*s)
	}
	return epc
}

// SetServiceYears sets the "ServiceYears" field.
func (epc *ExamPACreate) SetServiceYears(s string) *ExamPACreate {
	epc.mutation.SetServiceYears(s)
	return epc
}

// SetNillableServiceYears sets the "ServiceYears" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableServiceYears(s *string) *ExamPACreate {
	if s != nil {
		epc.SetServiceYears(*s)
	}
	return epc
}

// SetDrivingLicenseRequired sets the "DrivingLicenseRequired" field.
func (epc *ExamPACreate) SetDrivingLicenseRequired(s string) *ExamPACreate {
	epc.mutation.SetDrivingLicenseRequired(s)
	return epc
}

// SetNillableDrivingLicenseRequired sets the "DrivingLicenseRequired" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableDrivingLicenseRequired(s *string) *ExamPACreate {
	if s != nil {
		epc.SetDrivingLicenseRequired(*s)
	}
	return epc
}

// SetExamPaperCode sets the "ExamPaperCode" field.
func (epc *ExamPACreate) SetExamPaperCode(s string) *ExamPACreate {
	epc.mutation.SetExamPaperCode(s)
	return epc
}

// SetNillableExamPaperCode sets the "ExamPaperCode" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaperCode(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaperCode(*s)
	}
	return epc
}

// SetExamPaper1 sets the "ExamPaper1" field.
func (epc *ExamPACreate) SetExamPaper1(s string) *ExamPACreate {
	epc.mutation.SetExamPaper1(s)
	return epc
}

// SetNillableExamPaper1 sets the "ExamPaper1" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper1(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper1(*s)
	}
	return epc
}

// SetExamPaper2 sets the "ExamPaper2" field.
func (epc *ExamPACreate) SetExamPaper2(s string) *ExamPACreate {
	epc.mutation.SetExamPaper2(s)
	return epc
}

// SetNillableExamPaper2 sets the "ExamPaper2" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper2(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper2(*s)
	}
	return epc
}

// SetExamPaper3 sets the "ExamPaper3" field.
func (epc *ExamPACreate) SetExamPaper3(s string) *ExamPACreate {
	epc.mutation.SetExamPaper3(s)
	return epc
}

// SetNillableExamPaper3 sets the "ExamPaper3" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper3(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper3(*s)
	}
	return epc
}

// SetExamPaper4 sets the "ExamPaper4" field.
func (epc *ExamPACreate) SetExamPaper4(s string) *ExamPACreate {
	epc.mutation.SetExamPaper4(s)
	return epc
}

// SetNillableExamPaper4 sets the "ExamPaper4" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper4(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper4(*s)
	}
	return epc
}

// SetExamPaper5 sets the "ExamPaper5" field.
func (epc *ExamPACreate) SetExamPaper5(s string) *ExamPACreate {
	epc.mutation.SetExamPaper5(s)
	return epc
}

// SetNillableExamPaper5 sets the "ExamPaper5" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper5(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper5(*s)
	}
	return epc
}

// SetExamPaper6 sets the "ExamPaper6" field.
func (epc *ExamPACreate) SetExamPaper6(s string) *ExamPACreate {
	epc.mutation.SetExamPaper6(s)
	return epc
}

// SetNillableExamPaper6 sets the "ExamPaper6" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableExamPaper6(s *string) *ExamPACreate {
	if s != nil {
		epc.SetExamPaper6(*s)
	}
	return epc
}

// SetPayLevelEligibilty sets the "PayLevelEligibilty" field.
func (epc *ExamPACreate) SetPayLevelEligibilty(b bool) *ExamPACreate {
	epc.mutation.SetPayLevelEligibilty(b)
	return epc
}

// SetNillablePayLevelEligibilty sets the "PayLevelEligibilty" field if the given value is not nil.
func (epc *ExamPACreate) SetNillablePayLevelEligibilty(b *bool) *ExamPACreate {
	if b != nil {
		epc.SetPayLevelEligibilty(*b)
	}
	return epc
}

// SetCategoryMinMarksSCSTPH sets the "CategoryMinMarksSCSTPH" field.
func (epc *ExamPACreate) SetCategoryMinMarksSCSTPH(s string) *ExamPACreate {
	epc.mutation.SetCategoryMinMarksSCSTPH(s)
	return epc
}

// SetNillableCategoryMinMarksSCSTPH sets the "CategoryMinMarksSCSTPH" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCategoryMinMarksSCSTPH(s *string) *ExamPACreate {
	if s != nil {
		epc.SetCategoryMinMarksSCSTPH(*s)
	}
	return epc
}

// SetCategoryMinMarksGENOBC sets the "CategoryMinMarksGENOBC" field.
func (epc *ExamPACreate) SetCategoryMinMarksGENOBC(s string) *ExamPACreate {
	epc.mutation.SetCategoryMinMarksGENOBC(s)
	return epc
}

// SetNillableCategoryMinMarksGENOBC sets the "CategoryMinMarksGENOBC" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableCategoryMinMarksGENOBC(s *string) *ExamPACreate {
	if s != nil {
		epc.SetCategoryMinMarksGENOBC(*s)
	}
	return epc
}

// SetLocalLanguageAllowed sets the "LocalLanguageAllowed" field.
func (epc *ExamPACreate) SetLocalLanguageAllowed(b bool) *ExamPACreate {
	epc.mutation.SetLocalLanguageAllowed(b)
	return epc
}

// SetNillableLocalLanguageAllowed sets the "LocalLanguageAllowed" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableLocalLanguageAllowed(b *bool) *ExamPACreate {
	if b != nil {
		epc.SetLocalLanguageAllowed(*b)
	}
	return epc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (epc *ExamPACreate) SetUpdatedAt(t time.Time) *ExamPACreate {
	epc.mutation.SetUpdatedAt(t)
	return epc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableUpdatedAt(t *time.Time) *ExamPACreate {
	if t != nil {
		epc.SetUpdatedAt(*t)
	}
	return epc
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (epc *ExamPACreate) SetUpdatedBy(s string) *ExamPACreate {
	epc.mutation.SetUpdatedBy(s)
	return epc
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (epc *ExamPACreate) SetNillableUpdatedBy(s *string) *ExamPACreate {
	if s != nil {
		epc.SetUpdatedBy(*s)
	}
	return epc
}

// SetID sets the "id" field.
func (epc *ExamPACreate) SetID(i int32) *ExamPACreate {
	epc.mutation.SetID(i)
	return epc
}

// AddExamcalPsRefIDs adds the "examcal_ps_ref" edge to the ExamCalendar entity by IDs.
func (epc *ExamPACreate) AddExamcalPsRefIDs(ids ...int32) *ExamPACreate {
	epc.mutation.AddExamcalPsRefIDs(ids...)
	return epc
}

// AddExamcalPsRef adds the "examcal_ps_ref" edges to the ExamCalendar entity.
func (epc *ExamPACreate) AddExamcalPsRef(e ...*ExamCalendar) *ExamPACreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddExamcalPsRefIDs(ids...)
}

// AddPapersPsRefIDs adds the "papers_ps_ref" edge to the ExamPapers entity by IDs.
func (epc *ExamPACreate) AddPapersPsRefIDs(ids ...int32) *ExamPACreate {
	epc.mutation.AddPapersPsRefIDs(ids...)
	return epc
}

// AddPapersPsRef adds the "papers_ps_ref" edges to the ExamPapers entity.
func (epc *ExamPACreate) AddPapersPsRef(e ...*ExamPapers) *ExamPACreate {
	ids := make([]int32, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddPapersPsRefIDs(ids...)
}

// AddUsersPsTypeIDs adds the "users_ps_type" edge to the UserMaster entity by IDs.
func (epc *ExamPACreate) AddUsersPsTypeIDs(ids ...int64) *ExamPACreate {
	epc.mutation.AddUsersPsTypeIDs(ids...)
	return epc
}

// AddUsersPsType adds the "users_ps_type" edges to the UserMaster entity.
func (epc *ExamPACreate) AddUsersPsType(u ...*UserMaster) *ExamPACreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return epc.AddUsersPsTypeIDs(ids...)
}

// AddExamApplnPSRefIDs adds the "ExamAppln_PS_Ref" edge to the Exam_Applications_PS entity by IDs.
func (epc *ExamPACreate) AddExamApplnPSRefIDs(ids ...int64) *ExamPACreate {
	epc.mutation.AddExamApplnPSRefIDs(ids...)
	return epc
}

// AddExamApplnPSRef adds the "ExamAppln_PS_Ref" edges to the Exam_Applications_PS entity.
func (epc *ExamPACreate) AddExamApplnPSRef(e ...*Exam_Applications_PS) *ExamPACreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return epc.AddExamApplnPSRefIDs(ids...)
}

// AddNotificationsPIDs adds the "notifications_ps" edge to the Notification entity by IDs.
func (epc *ExamPACreate) AddNotificationsPIDs(ids ...int32) *ExamPACreate {
	epc.mutation.AddNotificationsPIDs(ids...)
	return epc
}

// AddNotificationsPs adds the "notifications_ps" edges to the Notification entity.
func (epc *ExamPACreate) AddNotificationsPs(n ...*Notification) *ExamPACreate {
	ids := make([]int32, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return epc.AddNotificationsPIDs(ids...)
}

// Mutation returns the ExamPAMutation object of the builder.
func (epc *ExamPACreate) Mutation() *ExamPAMutation {
	return epc.mutation
}

// Save creates the Exam_PA in the database.
func (epc *ExamPACreate) Save(ctx context.Context) (*Exam_PA, error) {
	epc.defaults()
	return withHooks(ctx, epc.sqlSave, epc.mutation, epc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (epc *ExamPACreate) SaveX(ctx context.Context) *Exam_PA {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epc *ExamPACreate) Exec(ctx context.Context) error {
	_, err := epc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epc *ExamPACreate) ExecX(ctx context.Context) {
	if err := epc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (epc *ExamPACreate) defaults() {
	if _, ok := epc.mutation.UpdatedAt(); !ok {
		v := exam_pa.DefaultUpdatedAt()
		epc.mutation.SetUpdatedAt(v)
	}
	if _, ok := epc.mutation.UpdatedBy(); !ok {
		v := exam_pa.DefaultUpdatedBy
		epc.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epc *ExamPACreate) check() error {
	if _, ok := epc.mutation.ExamName(); !ok {
		return &ValidationError{Name: "ExamName", err: errors.New(`ent: missing required field "Exam_PA.ExamName"`)}
	}
	if _, ok := epc.mutation.ExamType(); !ok {
		return &ValidationError{Name: "ExamType", err: errors.New(`ent: missing required field "Exam_PA.ExamType"`)}
	}
	if _, ok := epc.mutation.ConductedBy(); !ok {
		return &ValidationError{Name: "ConductedBy", err: errors.New(`ent: missing required field "Exam_PA.ConductedBy"`)}
	}
	return nil
}

func (epc *ExamPACreate) sqlSave(ctx context.Context) (*Exam_PA, error) {
	if err := epc.check(); err != nil {
		return nil, err
	}
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	epc.mutation.id = &_node.ID
	epc.mutation.done = true
	return _node, nil
}

func (epc *ExamPACreate) createSpec() (*Exam_PA, *sqlgraph.CreateSpec) {
	var (
		_node = &Exam_PA{config: epc.config}
		_spec = sqlgraph.NewCreateSpec(exam_pa.Table, sqlgraph.NewFieldSpec(exam_pa.FieldID, field.TypeInt32))
	)
	if id, ok := epc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := epc.mutation.ExamNameCode(); ok {
		_spec.SetField(exam_pa.FieldExamNameCode, field.TypeString, value)
		_node.ExamNameCode = value
	}
	if value, ok := epc.mutation.ExamName(); ok {
		_spec.SetField(exam_pa.FieldExamName, field.TypeString, value)
		_node.ExamName = value
	}
	if value, ok := epc.mutation.ExamType(); ok {
		_spec.SetField(exam_pa.FieldExamType, field.TypeString, value)
		_node.ExamType = value
	}
	if value, ok := epc.mutation.NotificationCode(); ok {
		_spec.SetField(exam_pa.FieldNotificationCode, field.TypeInt32, value)
		_node.NotificationCode = value
	}
	if value, ok := epc.mutation.ConductedBy(); ok {
		_spec.SetField(exam_pa.FieldConductedBy, field.TypeString, value)
		_node.ConductedBy = value
	}
	if value, ok := epc.mutation.NodalOffice(); ok {
		_spec.SetField(exam_pa.FieldNodalOffice, field.TypeString, value)
		_node.NodalOffice = value
	}
	if value, ok := epc.mutation.CalendarCode(); ok {
		_spec.SetField(exam_pa.FieldCalendarCode, field.TypeInt32, value)
		_node.CalendarCode = value
	}
	if value, ok := epc.mutation.PaperCode(); ok {
		_spec.SetField(exam_pa.FieldPaperCode, field.TypeInt32, value)
		_node.PaperCode = value
	}
	if value, ok := epc.mutation.EligibleCadre(); ok {
		_spec.SetField(exam_pa.FieldEligibleCadre, field.TypeString, value)
		_node.EligibleCadre = value
	}
	if value, ok := epc.mutation.EligiblePost1(); ok {
		_spec.SetField(exam_pa.FieldEligiblePost1, field.TypeString, value)
		_node.EligiblePost1 = value
	}
	if value, ok := epc.mutation.EligiblePost2(); ok {
		_spec.SetField(exam_pa.FieldEligiblePost2, field.TypeString, value)
		_node.EligiblePost2 = value
	}
	if value, ok := epc.mutation.EligiblePost3(); ok {
		_spec.SetField(exam_pa.FieldEligiblePost3, field.TypeString, value)
		_node.EligiblePost3 = value
	}
	if value, ok := epc.mutation.EligiblePost4(); ok {
		_spec.SetField(exam_pa.FieldEligiblePost4, field.TypeString, value)
		_node.EligiblePost4 = value
	}
	if value, ok := epc.mutation.EligiblePost5(); ok {
		_spec.SetField(exam_pa.FieldEligiblePost5, field.TypeString, value)
		_node.EligiblePost5 = value
	}
	if value, ok := epc.mutation.ExamPost1(); ok {
		_spec.SetField(exam_pa.FieldExamPost1, field.TypeString, value)
		_node.ExamPost1 = value
	}
	if value, ok := epc.mutation.ExamPost2(); ok {
		_spec.SetField(exam_pa.FieldExamPost2, field.TypeString, value)
		_node.ExamPost2 = value
	}
	if value, ok := epc.mutation.ExamPost3(); ok {
		_spec.SetField(exam_pa.FieldExamPost3, field.TypeString, value)
		_node.ExamPost3 = value
	}
	if value, ok := epc.mutation.ExamPost4(); ok {
		_spec.SetField(exam_pa.FieldExamPost4, field.TypeString, value)
		_node.ExamPost4 = value
	}
	if value, ok := epc.mutation.ExamPost5(); ok {
		_spec.SetField(exam_pa.FieldExamPost5, field.TypeString, value)
		_node.ExamPost5 = value
	}
	if value, ok := epc.mutation.EducationCriteria(); ok {
		_spec.SetField(exam_pa.FieldEducationCriteria, field.TypeString, value)
		_node.EducationCriteria = value
	}
	if value, ok := epc.mutation.CategoryAgeLimitGEN(); ok {
		_spec.SetField(exam_pa.FieldCategoryAgeLimitGEN, field.TypeString, value)
		_node.CategoryAgeLimitGEN = value
	}
	if value, ok := epc.mutation.CategoryAgeLimitSC(); ok {
		_spec.SetField(exam_pa.FieldCategoryAgeLimitSC, field.TypeString, value)
		_node.CategoryAgeLimitSC = value
	}
	if value, ok := epc.mutation.CategoryAgeLimitST(); ok {
		_spec.SetField(exam_pa.FieldCategoryAgeLimitST, field.TypeString, value)
		_node.CategoryAgeLimitST = value
	}
	if value, ok := epc.mutation.ServiceYears(); ok {
		_spec.SetField(exam_pa.FieldServiceYears, field.TypeString, value)
		_node.ServiceYears = value
	}
	if value, ok := epc.mutation.DrivingLicenseRequired(); ok {
		_spec.SetField(exam_pa.FieldDrivingLicenseRequired, field.TypeString, value)
		_node.DrivingLicenseRequired = value
	}
	if value, ok := epc.mutation.ExamPaperCode(); ok {
		_spec.SetField(exam_pa.FieldExamPaperCode, field.TypeString, value)
		_node.ExamPaperCode = value
	}
	if value, ok := epc.mutation.ExamPaper1(); ok {
		_spec.SetField(exam_pa.FieldExamPaper1, field.TypeString, value)
		_node.ExamPaper1 = value
	}
	if value, ok := epc.mutation.ExamPaper2(); ok {
		_spec.SetField(exam_pa.FieldExamPaper2, field.TypeString, value)
		_node.ExamPaper2 = value
	}
	if value, ok := epc.mutation.ExamPaper3(); ok {
		_spec.SetField(exam_pa.FieldExamPaper3, field.TypeString, value)
		_node.ExamPaper3 = value
	}
	if value, ok := epc.mutation.ExamPaper4(); ok {
		_spec.SetField(exam_pa.FieldExamPaper4, field.TypeString, value)
		_node.ExamPaper4 = value
	}
	if value, ok := epc.mutation.ExamPaper5(); ok {
		_spec.SetField(exam_pa.FieldExamPaper5, field.TypeString, value)
		_node.ExamPaper5 = value
	}
	if value, ok := epc.mutation.ExamPaper6(); ok {
		_spec.SetField(exam_pa.FieldExamPaper6, field.TypeString, value)
		_node.ExamPaper6 = value
	}
	if value, ok := epc.mutation.PayLevelEligibilty(); ok {
		_spec.SetField(exam_pa.FieldPayLevelEligibilty, field.TypeBool, value)
		_node.PayLevelEligibilty = value
	}
	if value, ok := epc.mutation.CategoryMinMarksSCSTPH(); ok {
		_spec.SetField(exam_pa.FieldCategoryMinMarksSCSTPH, field.TypeString, value)
		_node.CategoryMinMarksSCSTPH = value
	}
	if value, ok := epc.mutation.CategoryMinMarksGENOBC(); ok {
		_spec.SetField(exam_pa.FieldCategoryMinMarksGENOBC, field.TypeString, value)
		_node.CategoryMinMarksGENOBC = value
	}
	if value, ok := epc.mutation.LocalLanguageAllowed(); ok {
		_spec.SetField(exam_pa.FieldLocalLanguageAllowed, field.TypeBool, value)
		_node.LocalLanguageAllowed = value
	}
	if value, ok := epc.mutation.UpdatedAt(); ok {
		_spec.SetField(exam_pa.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := epc.mutation.UpdatedBy(); ok {
		_spec.SetField(exam_pa.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if nodes := epc.mutation.ExamcalPsRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam_pa.ExamcalPsRefTable,
			Columns: []string{exam_pa.ExamcalPsRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcalendar.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.PapersPsRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam_pa.PapersPsRefTable,
			Columns: []string{exam_pa.PapersPsRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exampapers.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.UsersPsTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam_pa.UsersPsTypeTable,
			Columns: []string{exam_pa.UsersPsTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermaster.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.ExamApplnPSRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam_pa.ExamApplnPSRefTable,
			Columns: []string{exam_pa.ExamApplnPSRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_ps.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := epc.mutation.NotificationsPsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam_pa.NotificationsPsTable,
			Columns: []string{exam_pa.NotificationsPsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExamPACreateBulk is the builder for creating many Exam_PA entities in bulk.
type ExamPACreateBulk struct {
	config
	builders []*ExamPACreate
}

// Save creates the Exam_PA entities in the database.
func (epcb *ExamPACreateBulk) Save(ctx context.Context) ([]*Exam_PA, error) {
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*Exam_PA, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExamPAMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (epcb *ExamPACreateBulk) SaveX(ctx context.Context) []*Exam_PA {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epcb *ExamPACreateBulk) Exec(ctx context.Context) error {
	_, err := epcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epcb *ExamPACreateBulk) ExecX(ctx context.Context) {
	if err := epcb.Exec(ctx); err != nil {
		panic(err)
	}
}