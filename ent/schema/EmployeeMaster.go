package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type EmployeeMaster struct {
	ent.Schema
}

// Fields of the User.
func (EmployeeMaster) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("EmpID"),
		field.Int64("EmployeeID").Unique(),
		field.String("EmployeeName").Optional(),
		//field.Time("DOB").SchemaType(map[string]string{
		//	dialect.Postgres: "date",
		//}).Optional(),
		field.String("DOB").Optional(),
		field.Enum("Gender").Values("Male", "Female"),
		field.String("MobileNumber").Optional(),
		field.String("EmailID").Optional(),
		field.String("EmployeeCategoryCode").Optional(),
		field.String("EmployeeCategory").Optional(),
		field.String("PostCode").Optional(),
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
		field.Time("UpdatedAt").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.String("UpdatedBy").Default("API").Optional(),
		field.String("Cadre").Optional(),
	}

	//Employee Number,Correct/Incorrect, Remarks
}

// Edges of the User.
func (EmployeeMaster) Edges() []ent.Edge {
	return []ent.Edge{edge.To("UsermasterRef", UserMaster.Type),
		edge.To("Emp_Ref", Exam_Applications_PS.Type)}
}

func (EmployeeMaster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeeMaster"}}
}
