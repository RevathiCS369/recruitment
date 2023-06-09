package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type ExamPapers struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (ExamPapers) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("PaperCode"),
		field.String("PaperDescription").MaxLen(100).NotEmpty(),
		field.Int32("ExamCode").Optional(),
		field.Bool("CompetitiveQualifying").Default(false),
		field.Bool("ExceptionForDisability").Default(false),
		field.Int("MaximumMarks").Positive(),
		field.Int("Duration").Positive(),
		field.String("localLanguageAllowedQuestionPaper").MaxLen(10),
		field.String("localLanguageAllowedAnswerPaper").MaxLen(10),
		field.String("OrderNumber").MaxLen(20),
		field.String("PaperStatus").MaxLen(10).NotEmpty(),
		field.Int32("CalendarCode").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Int32("PaperTypeCode").Optional(),
		field.String("PaperTypeName").Optional(),
		field.Int32("DisabilityTypeID").Optional(),
		field.Int32("ExamCodePS").Optional(),
	}
}

// Edges of the ExamPapers.

func (ExamPapers) Edges() []ent.Edge {

	return []ent.Edge{edge.To("centers", Center.Type),
		edge.From("exam", Exam.Type).Ref("papers").Unique().Field("ExamCode"),
		edge.To("exampapers_types", PaperTypes.Type),
		edge.To("papers_ref", ExamCalendar.Type),
		edge.To("ExamPaperEligibility", EligibilityMaster.Type),
		//edge.From("disabilitiesreference", Disability.Type).Ref("dis_ref").Unique().Field("DisabilityTypeID"),
		edge.To("dis_ref", Disability.Type),
		//edge.From("exam_papers_ps", Exam_PS.Type).Ref("papers_ps_ref").Unique().Field("ExamCodePS"),
		edge.To("papers_ps_ref", Exam_PS.Type),
		edge.To("papers_ip_ref", Exam_IP.Type),
	}

}

func (ExamPapers) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ExamPapers"}}
}
