package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamPapers holds the schema definition for the ExamPapers entity.
type PaperTypes struct {
	ent.Schema
}

// Fields of the ExamPapers.
func (PaperTypes) Fields() []ent.Field {
	return []ent.Field{field.Int32("id").StorageKey("PaperTypeCode"),
		field.Int32("PaperCode").Optional(),
		field.String("PaperTypeDescription").MaxLen(100).NotEmpty(),
		field.String("OrderNumber").MaxLen(100),
		field.Int32("SequenceNumber").Optional(),
		field.Time("CreatedDate").SchemaType(map[string]string{
			dialect.Postgres: "date",
		}).Optional(),
	}
}

// Edges of the ExamPapers.
func (PaperTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("papercode", ExamPapers.Type).Ref("exampapers_types").Unique().Field("PaperCode"),
	}
}
