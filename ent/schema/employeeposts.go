// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EmployeePosts struct {
	ent.Schema
}

func (EmployeePosts) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").StorageKey("PostID"), 
		field.String("PostCode").Unique(),
		field.String("PostDescription"),
		field.String("Group"),
		field.String("PayLevel"),
		field.String("Scale"),
		field.Bool("BaseCadreFlag"),}
}
func (EmployeePosts) Edges() []ent.Edge {
	return []ent.Edge{edge.To("emp_posts",Employees.Type),
	edge.To("PostEligibility", EligibilityMaster.Type),
}
}
func (EmployeePosts) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "EmployeePosts"}}
}
