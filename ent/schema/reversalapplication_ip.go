// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Reversal_Application_IP struct {
	ent.Schema
}

func (Reversal_Application_IP) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ApplicationID"), 
		field.Int64("EmployeeID").Optional(),
		field.String("EmployeeName").Optional(),
		//field.Time("DOB").SchemaType(map[string]string{
		//	dialect.Postgres: "date",
		//}).Optional(),
		field.String("DOB").Optional(),
		field.Enum("Gender").Values("Male", "Female"),
		field.String("MobileNumber").Optional(),
		field.String("EmailID").Optional(),		
		field.String("EmployeeCategory").Optional(),
		field.String("Cadre").Optional(),
		field.String("EmployeePost").Optional(),
		
		field.String("FacilityID").Optional(),
		/*field.Time("DCCS").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Time("DCInPresentCadre").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),*/
		field.String("DCCS").Optional(),
		field.String("DCInPresentCadre").Optional(),
		field.String("DeputationOfficeId").Optional(),
		field.String("DisabilityType").Optional(),
		field.String("DisabilityPercentage").Optional(),
		field.String("Education").Optional(),
		field.Int32("ExamCodeIP").Optional(),
		field.String("ExamYear").Optional(),
		field.String("CentrePreference").Optional(),
		field.String("Signature").Optional(),
		field.String("Photo").Optional(),
		field.String("ApplicationStatus").Optional(),
		field.Time("ReversalApplnSubmittedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("VA_Remarks").Optional(),
		field.String("VA_UserName").Optional(),
		field.Time("VA_Date").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("CA_Remarks").Optional(),
		field.String("CA_UserName").Optional(),
		field.Time("CA_Date").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),	
		field.String("AppliactionRemarks").Optional(),
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Default(time.Now).Optional(),	
		field.String("UpdatedBy").Default("API").Optional(),
		field.Int32("RoleUserCode").Optional(),
		}
		
		
}
func (Reversal_Application_IP) Edges() []ent.Edge {
	return nil
	/*[]ent.Edge{
		edge.From("emps", EmployeeMaster.Type).Ref("Emp_Ref").Unique().Field("EmployeeID"),
		//edge.From("users", UserMaster.Type).Ref("Users_PS_Ref").Unique().Field("UserName"),	
		edge.To("UsersIPRef", UserMaster.Type),
		edge.From("examsip", Exam_PS.Type).Ref("ExamAppln_IP_Ref").Unique().Field("ExamCodeIP"),	
		//edge.From("offices", Facility.Type).Ref("Office_PS_Ref").Unique().Field("FacilityOfficeID"),
		edge.To("Office_IP_Ref", Facility.Type),
		edge.From("roleusers", RoleMaster.Type).Ref("Roles_IP_Ref").Unique().Field("RoleUserCode"),}*/

}


func (Reversal_Application_IP) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Reversal_Application_IP"}}
}
