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
type RoleMaster struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (RoleMaster) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("RoleUserCode"),
		field.String("RoleName"),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
		field.Bool("Status").Default(true),
	}
}

// Edges of the ExamPapers.

func (RoleMaster) Edges() []ent.Edge {

	return []ent.Edge{edge.To("roles", AdminLogin.Type),
		edge.To("Roles_Ref", UserMaster.Type),
		edge.To("Roles_PS_Ref", Exam_Applications_PS.Type),
		edge.To("Roles_IP_Ref", Exam_Applications_IP.Type)}

}

func (RoleMaster) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "RoleMaster"}}
}
