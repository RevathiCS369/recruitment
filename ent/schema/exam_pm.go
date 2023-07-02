// Code generated by entimport, DO NOT EDIT.

package schema

import (
	
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent"
	"time"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exam holds the schema definition for the Exam entity.
type Exam_PM struct {
	ent.Schema
}

/// Fields of the Exam.
func (Exam_PM) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("ExamCodePM"), 
	field.String("ExamNameCode").Optional().Unique(),
	field.String("ExamName"), 
	//field.Int32("NumOfPapers"),
	field.String("ExamType"), 	
	field.Int32("NotificationCode").Optional(),		
	field.String("ConductedBy"), field.String("NodalOffice").Optional(),
	field.Int32("CalendarCode").Optional(),
	field.Int32("PaperCode").Optional(),
	field.String("EligibleCadre").Optional(),
	field.String("EligiblePost1").Optional(),
	field.String("EligiblePost2").Optional(),
	field.String("EligiblePost3").Optional(),
	field.String("EligiblePost4").Optional(),
	field.String("EligiblePost5").Optional(),
	field.String("ExamPost1").Optional(),
	field.String("ExamPost2").Optional(),
	field.String("ExamPost3").Optional(),
	field.String("ExamPost4").Optional(),
	field.String("ExamPost5").Optional(),
	field.String("EducationCriteria").Optional(),
	field.String("CategoryAgeLimitGEN").Optional(),
	field.String("CategoryAgeLimitSC").Optional(),
	field.String("CategoryAgeLimitST").Optional(),
	field.String("ServiceYears").Optional(),
	field.String("DrivingLicenseRequired").Optional(),
	field.String("ExamPaperCode").Optional(),
	field.String("ExamPaper1").Optional(),
	field.String("ExamPaper2").Optional(),
	field.String("ExamPaper3").Optional(),
	field.String("ExamPaper4").Optional(),
	field.String("ExamPaper5").Optional(),
	field.String("ExamPaper6").Optional(),
	field.Bool("PayLevelEligibilty").Optional(),
	field.String("CategoryMinMarksSCSTPH").Optional(),
	field.String("CategoryMinMarksGENOBC").Optional(),
	field.Bool("LocalLanguageAllowed").Optional(),
	field.Time("UpdatedAt").SchemaType(map[string]string{
		dialect.Postgres: "date",
	}).Default(time.Now).Optional(),	
	field.String("UpdatedBy").Default("API").Optional(),}
}
// Edges of the Exam.
func (Exam_PM) Edges() []ent.Edge {
	return nil
	/*[]ent.Edge{	
	
	//edge.To("exam", Exam.Type).Ref("papers").Unique().Field("ExamCode"),
	edge.To("examcal_ps_ref", ExamCalendar.Type),
	edge.To("papers_ps_ref", ExamPapers.Type),
	edge.To("users_ps_type", UserMaster.Type),
	edge.To("ExamAppln_PS_Ref" , Exam_Applications_PS.Type),
	edge.To("notifications_ps", Notification.Type),

	}*/

}
func (Exam_PM) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Exam_PM"}}
}
