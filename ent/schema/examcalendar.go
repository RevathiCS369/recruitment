package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"

	//"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ExamCalendar struct {
	ent.Schema
}

func (ExamCalendar) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("CalendarCode"),

		field.Int32("ExamYear"),
		field.String("ExamName").NotEmpty(),
		field.Int32("ExamCode").Optional(),
		field.Time("NotificationDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("ModelNotificationDate").SchemaType(map[string]string{
			//dialect.Postgres: "date",
			dialect.Postgres: "date",
		}),
		field.Time("ApplicationEndDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("ApprovedOrderDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.Time("TentativeResultDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		//field.String("ExamName").NotEmpty(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}),
		field.String("ApprovedOrderNumber"),
		field.JSON("VacancyYears", []interface{}{}).
			Optional(),
		field.JSON("ExamPapers", []interface{}{}).
			Optional(),
		field.Int32("VacancyYearCode").Optional(),
		field.Int32("PaperCode").Optional(),
	}

}
func (ExamCalendar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vcy_years", VacancyYear.Type).Ref("vacancy_ref").Unique().Field("VacancyYearCode"),
		edge.From("exams", Exam.Type).Ref("exams_ref").Unique().Field("ExamCode"),
		edge.From("papers", ExamPapers.Type).Ref("papers_ref").Unique().Field("PaperCode"),
		edge.To("Notify_ref", Notification.Type)}
}
func (ExamCalendar) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamCalendar"}}
}
