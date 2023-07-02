// Code generated by ent, DO NOT EDIT.

package vacancyyear

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the vacancyyear type in the database.
	Label = "vacancy_year"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "VacancyYearCode"
	// FieldFromDate holds the string denoting the fromdate field in the database.
	FieldFromDate = "from_date"
	// FieldToDate holds the string denoting the todate field in the database.
	FieldToDate = "to_date"
	// FieldNotifyCode holds the string denoting the notifycode field in the database.
	FieldNotifyCode = "notify_code"
	// FieldVacancyYear holds the string denoting the vacancyyear field in the database.
	FieldVacancyYear = "vacancy_year"
	// FieldCalendarCode holds the string denoting the calendarcode field in the database.
	FieldCalendarCode = "calendar_code"
	// EdgeVacancyRef holds the string denoting the vacancy_ref edge name in mutations.
	EdgeVacancyRef = "vacancy_ref"
	// EdgeExams holds the string denoting the exams edge name in mutations.
	EdgeExams = "exams"
	// ExamCalendarFieldID holds the string denoting the ID field of the ExamCalendar.
	ExamCalendarFieldID = "CalendarCode"
	// ExamFieldID holds the string denoting the ID field of the Exam.
	ExamFieldID = "ExamCode"
	// Table holds the table name of the vacancyyear in the database.
	Table = "VacancyYears"
	// VacancyRefTable is the table that holds the vacancy_ref relation/edge.
	VacancyRefTable = "ExamCalendar"
	// VacancyRefInverseTable is the table name for the ExamCalendar entity.
	// It exists in this package in order to avoid circular dependency with the "examcalendar" package.
	VacancyRefInverseTable = "ExamCalendar"
	// VacancyRefColumn is the table column denoting the vacancy_ref relation/edge.
	VacancyRefColumn = "vacancy_year_code"
	// ExamsTable is the table that holds the exams relation/edge.
	ExamsTable = "Exam"
	// ExamsInverseTable is the table name for the Exam entity.
	// It exists in this package in order to avoid circular dependency with the "exam" package.
	ExamsInverseTable = "Exam"
	// ExamsColumn is the table column denoting the exams relation/edge.
	ExamsColumn = "vacancy_year_exams"
)

// Columns holds all SQL columns for vacancyyear fields.
var Columns = []string{
	FieldID,
	FieldFromDate,
	FieldToDate,
	FieldNotifyCode,
	FieldVacancyYear,
	FieldCalendarCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "VacancyYears"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"notification_vacancy_years",
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

// OrderOption defines the ordering options for the VacancyYear queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFromDate orders the results by the FromDate field.
func ByFromDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromDate, opts...).ToFunc()
}

// ByToDate orders the results by the ToDate field.
func ByToDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToDate, opts...).ToFunc()
}

// ByNotifyCode orders the results by the NotifyCode field.
func ByNotifyCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifyCode, opts...).ToFunc()
}

// ByVacancyYear orders the results by the VacancyYear field.
func ByVacancyYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVacancyYear, opts...).ToFunc()
}

// ByCalendarCode orders the results by the CalendarCode field.
func ByCalendarCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCalendarCode, opts...).ToFunc()
}

// ByVacancyRefCount orders the results by vacancy_ref count.
func ByVacancyRefCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVacancyRefStep(), opts...)
	}
}

// ByVacancyRef orders the results by vacancy_ref terms.
func ByVacancyRef(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVacancyRefStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByExamsCount orders the results by exams count.
func ByExamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamsStep(), opts...)
	}
}

// ByExams orders the results by exams terms.
func ByExams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVacancyRefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VacancyRefInverseTable, ExamCalendarFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VacancyRefTable, VacancyRefColumn),
	)
}
func newExamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamsInverseTable, ExamFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
	)
}