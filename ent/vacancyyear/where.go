// Code generated by ent, DO NOT EDIT.

package vacancyyear

import (
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldID, id))
}

// FromDate applies equality check predicate on the "FromDate" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldFromDate, v))
}

// ToDate applies equality check predicate on the "ToDate" field. It's identical to ToDateEQ.
func ToDate(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldToDate, v))
}

// NotifyCode applies equality check predicate on the "NotifyCode" field. It's identical to NotifyCodeEQ.
func NotifyCode(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldNotifyCode, v))
}

// VacancyYear applies equality check predicate on the "VacancyYear" field. It's identical to VacancyYearEQ.
func VacancyYear(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldVacancyYear, v))
}

// CalendarCode applies equality check predicate on the "CalendarCode" field. It's identical to CalendarCodeEQ.
func CalendarCode(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldCalendarCode, v))
}

// FromDateEQ applies the EQ predicate on the "FromDate" field.
func FromDateEQ(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldFromDate, v))
}

// FromDateNEQ applies the NEQ predicate on the "FromDate" field.
func FromDateNEQ(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldFromDate, v))
}

// FromDateIn applies the In predicate on the "FromDate" field.
func FromDateIn(vs ...time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldFromDate, vs...))
}

// FromDateNotIn applies the NotIn predicate on the "FromDate" field.
func FromDateNotIn(vs ...time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldFromDate, vs...))
}

// FromDateGT applies the GT predicate on the "FromDate" field.
func FromDateGT(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldFromDate, v))
}

// FromDateGTE applies the GTE predicate on the "FromDate" field.
func FromDateGTE(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldFromDate, v))
}

// FromDateLT applies the LT predicate on the "FromDate" field.
func FromDateLT(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldFromDate, v))
}

// FromDateLTE applies the LTE predicate on the "FromDate" field.
func FromDateLTE(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldFromDate, v))
}

// ToDateEQ applies the EQ predicate on the "ToDate" field.
func ToDateEQ(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldToDate, v))
}

// ToDateNEQ applies the NEQ predicate on the "ToDate" field.
func ToDateNEQ(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldToDate, v))
}

// ToDateIn applies the In predicate on the "ToDate" field.
func ToDateIn(vs ...time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldToDate, vs...))
}

// ToDateNotIn applies the NotIn predicate on the "ToDate" field.
func ToDateNotIn(vs ...time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldToDate, vs...))
}

// ToDateGT applies the GT predicate on the "ToDate" field.
func ToDateGT(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldToDate, v))
}

// ToDateGTE applies the GTE predicate on the "ToDate" field.
func ToDateGTE(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldToDate, v))
}

// ToDateLT applies the LT predicate on the "ToDate" field.
func ToDateLT(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldToDate, v))
}

// ToDateLTE applies the LTE predicate on the "ToDate" field.
func ToDateLTE(v time.Time) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldToDate, v))
}

// NotifyCodeEQ applies the EQ predicate on the "NotifyCode" field.
func NotifyCodeEQ(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldNotifyCode, v))
}

// NotifyCodeNEQ applies the NEQ predicate on the "NotifyCode" field.
func NotifyCodeNEQ(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldNotifyCode, v))
}

// NotifyCodeIn applies the In predicate on the "NotifyCode" field.
func NotifyCodeIn(vs ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldNotifyCode, vs...))
}

// NotifyCodeNotIn applies the NotIn predicate on the "NotifyCode" field.
func NotifyCodeNotIn(vs ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldNotifyCode, vs...))
}

// NotifyCodeGT applies the GT predicate on the "NotifyCode" field.
func NotifyCodeGT(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldNotifyCode, v))
}

// NotifyCodeGTE applies the GTE predicate on the "NotifyCode" field.
func NotifyCodeGTE(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldNotifyCode, v))
}

// NotifyCodeLT applies the LT predicate on the "NotifyCode" field.
func NotifyCodeLT(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldNotifyCode, v))
}

// NotifyCodeLTE applies the LTE predicate on the "NotifyCode" field.
func NotifyCodeLTE(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldNotifyCode, v))
}

// NotifyCodeIsNil applies the IsNil predicate on the "NotifyCode" field.
func NotifyCodeIsNil() predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIsNull(FieldNotifyCode))
}

// NotifyCodeNotNil applies the NotNil predicate on the "NotifyCode" field.
func NotifyCodeNotNil() predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotNull(FieldNotifyCode))
}

// VacancyYearEQ applies the EQ predicate on the "VacancyYear" field.
func VacancyYearEQ(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldVacancyYear, v))
}

// VacancyYearNEQ applies the NEQ predicate on the "VacancyYear" field.
func VacancyYearNEQ(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldVacancyYear, v))
}

// VacancyYearIn applies the In predicate on the "VacancyYear" field.
func VacancyYearIn(vs ...string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldVacancyYear, vs...))
}

// VacancyYearNotIn applies the NotIn predicate on the "VacancyYear" field.
func VacancyYearNotIn(vs ...string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldVacancyYear, vs...))
}

// VacancyYearGT applies the GT predicate on the "VacancyYear" field.
func VacancyYearGT(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldVacancyYear, v))
}

// VacancyYearGTE applies the GTE predicate on the "VacancyYear" field.
func VacancyYearGTE(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldVacancyYear, v))
}

// VacancyYearLT applies the LT predicate on the "VacancyYear" field.
func VacancyYearLT(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldVacancyYear, v))
}

// VacancyYearLTE applies the LTE predicate on the "VacancyYear" field.
func VacancyYearLTE(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldVacancyYear, v))
}

// VacancyYearContains applies the Contains predicate on the "VacancyYear" field.
func VacancyYearContains(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldContains(FieldVacancyYear, v))
}

// VacancyYearHasPrefix applies the HasPrefix predicate on the "VacancyYear" field.
func VacancyYearHasPrefix(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldHasPrefix(FieldVacancyYear, v))
}

// VacancyYearHasSuffix applies the HasSuffix predicate on the "VacancyYear" field.
func VacancyYearHasSuffix(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldHasSuffix(FieldVacancyYear, v))
}

// VacancyYearEqualFold applies the EqualFold predicate on the "VacancyYear" field.
func VacancyYearEqualFold(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEqualFold(FieldVacancyYear, v))
}

// VacancyYearContainsFold applies the ContainsFold predicate on the "VacancyYear" field.
func VacancyYearContainsFold(v string) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldContainsFold(FieldVacancyYear, v))
}

// CalendarCodeEQ applies the EQ predicate on the "CalendarCode" field.
func CalendarCodeEQ(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldEQ(FieldCalendarCode, v))
}

// CalendarCodeNEQ applies the NEQ predicate on the "CalendarCode" field.
func CalendarCodeNEQ(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNEQ(FieldCalendarCode, v))
}

// CalendarCodeIn applies the In predicate on the "CalendarCode" field.
func CalendarCodeIn(vs ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIn(FieldCalendarCode, vs...))
}

// CalendarCodeNotIn applies the NotIn predicate on the "CalendarCode" field.
func CalendarCodeNotIn(vs ...int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotIn(FieldCalendarCode, vs...))
}

// CalendarCodeGT applies the GT predicate on the "CalendarCode" field.
func CalendarCodeGT(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGT(FieldCalendarCode, v))
}

// CalendarCodeGTE applies the GTE predicate on the "CalendarCode" field.
func CalendarCodeGTE(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldGTE(FieldCalendarCode, v))
}

// CalendarCodeLT applies the LT predicate on the "CalendarCode" field.
func CalendarCodeLT(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLT(FieldCalendarCode, v))
}

// CalendarCodeLTE applies the LTE predicate on the "CalendarCode" field.
func CalendarCodeLTE(v int32) predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldLTE(FieldCalendarCode, v))
}

// CalendarCodeIsNil applies the IsNil predicate on the "CalendarCode" field.
func CalendarCodeIsNil() predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldIsNull(FieldCalendarCode))
}

// CalendarCodeNotNil applies the NotNil predicate on the "CalendarCode" field.
func CalendarCodeNotNil() predicate.VacancyYear {
	return predicate.VacancyYear(sql.FieldNotNull(FieldCalendarCode))
}

// HasVacancyRef applies the HasEdge predicate on the "vacancy_ref" edge.
func HasVacancyRef() predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VacancyRefTable, VacancyRefColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVacancyRefWith applies the HasEdge predicate on the "vacancy_ref" edge with a given conditions (other predicates).
func HasVacancyRefWith(preds ...predicate.ExamCalendar) predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		step := newVacancyRefStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExams applies the HasEdge predicate on the "exams" edge.
func HasExams() predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamsWith applies the HasEdge predicate on the "exams" edge with a given conditions (other predicates).
func HasExamsWith(preds ...predicate.Exam) predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		step := newExamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VacancyYear) predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VacancyYear) predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VacancyYear) predicate.VacancyYear {
	return predicate.VacancyYear(func(s *sql.Selector) {
		p(s.Not())
	})
}
