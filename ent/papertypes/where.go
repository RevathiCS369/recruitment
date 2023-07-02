// Code generated by ent, DO NOT EDIT.

package papertypes

import (
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLTE(FieldID, id))
}

// PaperCode applies equality check predicate on the "PaperCode" field. It's identical to PaperCodeEQ.
func PaperCode(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldPaperCode, v))
}

// PaperTypeDescription applies equality check predicate on the "PaperTypeDescription" field. It's identical to PaperTypeDescriptionEQ.
func PaperTypeDescription(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldPaperTypeDescription, v))
}

// OrderNumber applies equality check predicate on the "OrderNumber" field. It's identical to OrderNumberEQ.
func OrderNumber(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldOrderNumber, v))
}

// SequenceNumber applies equality check predicate on the "SequenceNumber" field. It's identical to SequenceNumberEQ.
func SequenceNumber(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldSequenceNumber, v))
}

// CreatedDate applies equality check predicate on the "CreatedDate" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldCreatedDate, v))
}

// PaperCodeEQ applies the EQ predicate on the "PaperCode" field.
func PaperCodeEQ(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldPaperCode, v))
}

// PaperCodeNEQ applies the NEQ predicate on the "PaperCode" field.
func PaperCodeNEQ(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldPaperCode, v))
}

// PaperCodeIn applies the In predicate on the "PaperCode" field.
func PaperCodeIn(vs ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldPaperCode, vs...))
}

// PaperCodeNotIn applies the NotIn predicate on the "PaperCode" field.
func PaperCodeNotIn(vs ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldPaperCode, vs...))
}

// PaperCodeIsNil applies the IsNil predicate on the "PaperCode" field.
func PaperCodeIsNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIsNull(FieldPaperCode))
}

// PaperCodeNotNil applies the NotNil predicate on the "PaperCode" field.
func PaperCodeNotNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotNull(FieldPaperCode))
}

// PaperTypeDescriptionEQ applies the EQ predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionEQ(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionNEQ applies the NEQ predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionNEQ(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionIn applies the In predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionIn(vs ...string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldPaperTypeDescription, vs...))
}

// PaperTypeDescriptionNotIn applies the NotIn predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionNotIn(vs ...string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldPaperTypeDescription, vs...))
}

// PaperTypeDescriptionGT applies the GT predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionGT(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGT(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionGTE applies the GTE predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionGTE(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGTE(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionLT applies the LT predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionLT(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLT(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionLTE applies the LTE predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionLTE(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLTE(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionContains applies the Contains predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionContains(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldContains(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionHasPrefix applies the HasPrefix predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionHasPrefix(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldHasPrefix(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionHasSuffix applies the HasSuffix predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionHasSuffix(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldHasSuffix(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionEqualFold applies the EqualFold predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionEqualFold(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEqualFold(FieldPaperTypeDescription, v))
}

// PaperTypeDescriptionContainsFold applies the ContainsFold predicate on the "PaperTypeDescription" field.
func PaperTypeDescriptionContainsFold(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldContainsFold(FieldPaperTypeDescription, v))
}

// OrderNumberEQ applies the EQ predicate on the "OrderNumber" field.
func OrderNumberEQ(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldOrderNumber, v))
}

// OrderNumberNEQ applies the NEQ predicate on the "OrderNumber" field.
func OrderNumberNEQ(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldOrderNumber, v))
}

// OrderNumberIn applies the In predicate on the "OrderNumber" field.
func OrderNumberIn(vs ...string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldOrderNumber, vs...))
}

// OrderNumberNotIn applies the NotIn predicate on the "OrderNumber" field.
func OrderNumberNotIn(vs ...string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldOrderNumber, vs...))
}

// OrderNumberGT applies the GT predicate on the "OrderNumber" field.
func OrderNumberGT(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGT(FieldOrderNumber, v))
}

// OrderNumberGTE applies the GTE predicate on the "OrderNumber" field.
func OrderNumberGTE(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGTE(FieldOrderNumber, v))
}

// OrderNumberLT applies the LT predicate on the "OrderNumber" field.
func OrderNumberLT(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLT(FieldOrderNumber, v))
}

// OrderNumberLTE applies the LTE predicate on the "OrderNumber" field.
func OrderNumberLTE(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLTE(FieldOrderNumber, v))
}

// OrderNumberContains applies the Contains predicate on the "OrderNumber" field.
func OrderNumberContains(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldContains(FieldOrderNumber, v))
}

// OrderNumberHasPrefix applies the HasPrefix predicate on the "OrderNumber" field.
func OrderNumberHasPrefix(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldHasPrefix(FieldOrderNumber, v))
}

// OrderNumberHasSuffix applies the HasSuffix predicate on the "OrderNumber" field.
func OrderNumberHasSuffix(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldHasSuffix(FieldOrderNumber, v))
}

// OrderNumberEqualFold applies the EqualFold predicate on the "OrderNumber" field.
func OrderNumberEqualFold(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEqualFold(FieldOrderNumber, v))
}

// OrderNumberContainsFold applies the ContainsFold predicate on the "OrderNumber" field.
func OrderNumberContainsFold(v string) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldContainsFold(FieldOrderNumber, v))
}

// SequenceNumberEQ applies the EQ predicate on the "SequenceNumber" field.
func SequenceNumberEQ(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldSequenceNumber, v))
}

// SequenceNumberNEQ applies the NEQ predicate on the "SequenceNumber" field.
func SequenceNumberNEQ(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldSequenceNumber, v))
}

// SequenceNumberIn applies the In predicate on the "SequenceNumber" field.
func SequenceNumberIn(vs ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldSequenceNumber, vs...))
}

// SequenceNumberNotIn applies the NotIn predicate on the "SequenceNumber" field.
func SequenceNumberNotIn(vs ...int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldSequenceNumber, vs...))
}

// SequenceNumberGT applies the GT predicate on the "SequenceNumber" field.
func SequenceNumberGT(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGT(FieldSequenceNumber, v))
}

// SequenceNumberGTE applies the GTE predicate on the "SequenceNumber" field.
func SequenceNumberGTE(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGTE(FieldSequenceNumber, v))
}

// SequenceNumberLT applies the LT predicate on the "SequenceNumber" field.
func SequenceNumberLT(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLT(FieldSequenceNumber, v))
}

// SequenceNumberLTE applies the LTE predicate on the "SequenceNumber" field.
func SequenceNumberLTE(v int32) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLTE(FieldSequenceNumber, v))
}

// SequenceNumberIsNil applies the IsNil predicate on the "SequenceNumber" field.
func SequenceNumberIsNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIsNull(FieldSequenceNumber))
}

// SequenceNumberNotNil applies the NotNil predicate on the "SequenceNumber" field.
func SequenceNumberNotNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotNull(FieldSequenceNumber))
}

// CreatedDateEQ applies the EQ predicate on the "CreatedDate" field.
func CreatedDateEQ(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "CreatedDate" field.
func CreatedDateNEQ(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "CreatedDate" field.
func CreatedDateIn(vs ...time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "CreatedDate" field.
func CreatedDateNotIn(vs ...time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "CreatedDate" field.
func CreatedDateGT(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "CreatedDate" field.
func CreatedDateGTE(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "CreatedDate" field.
func CreatedDateLT(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "CreatedDate" field.
func CreatedDateLTE(v time.Time) predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldLTE(FieldCreatedDate, v))
}

// CreatedDateIsNil applies the IsNil predicate on the "CreatedDate" field.
func CreatedDateIsNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldIsNull(FieldCreatedDate))
}

// CreatedDateNotNil applies the NotNil predicate on the "CreatedDate" field.
func CreatedDateNotNil() predicate.PaperTypes {
	return predicate.PaperTypes(sql.FieldNotNull(FieldCreatedDate))
}

// HasPapercode applies the HasEdge predicate on the "papercode" edge.
func HasPapercode() predicate.PaperTypes {
	return predicate.PaperTypes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PapercodeTable, PapercodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPapercodeWith applies the HasEdge predicate on the "papercode" edge with a given conditions (other predicates).
func HasPapercodeWith(preds ...predicate.ExamPapers) predicate.PaperTypes {
	return predicate.PaperTypes(func(s *sql.Selector) {
		step := newPapercodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PaperTypes) predicate.PaperTypes {
	return predicate.PaperTypes(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PaperTypes) predicate.PaperTypes {
	return predicate.PaperTypes(func(s *sql.Selector) {
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
func Not(p predicate.PaperTypes) predicate.PaperTypes {
	return predicate.PaperTypes(func(s *sql.Selector) {
		p(s.Not())
	})
}
