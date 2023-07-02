// Code generated by ent, DO NOT EDIT.

package disability

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Disability {
	return predicate.Disability(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Disability {
	return predicate.Disability(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Disability {
	return predicate.Disability(sql.FieldLTE(FieldID, id))
}

// DisabilityTypeCode applies equality check predicate on the "DisabilityTypeCode" field. It's identical to DisabilityTypeCodeEQ.
func DisabilityTypeCode(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityTypeCode, v))
}

// DisabilityTypeDescription applies equality check predicate on the "DisabilityTypeDescription" field. It's identical to DisabilityTypeDescriptionEQ.
func DisabilityTypeDescription(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityTypeDescription, v))
}

// DisabilityPercentage applies equality check predicate on the "DisabilityPercentage" field. It's identical to DisabilityPercentageEQ.
func DisabilityPercentage(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityPercentage, v))
}

// DisabilityTypeCodeEQ applies the EQ predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeEQ(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeNEQ applies the NEQ predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeNEQ(v string) predicate.Disability {
	return predicate.Disability(sql.FieldNEQ(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeIn applies the In predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeIn(vs ...string) predicate.Disability {
	return predicate.Disability(sql.FieldIn(FieldDisabilityTypeCode, vs...))
}

// DisabilityTypeCodeNotIn applies the NotIn predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeNotIn(vs ...string) predicate.Disability {
	return predicate.Disability(sql.FieldNotIn(FieldDisabilityTypeCode, vs...))
}

// DisabilityTypeCodeGT applies the GT predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeGT(v string) predicate.Disability {
	return predicate.Disability(sql.FieldGT(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeGTE applies the GTE predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeGTE(v string) predicate.Disability {
	return predicate.Disability(sql.FieldGTE(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeLT applies the LT predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeLT(v string) predicate.Disability {
	return predicate.Disability(sql.FieldLT(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeLTE applies the LTE predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeLTE(v string) predicate.Disability {
	return predicate.Disability(sql.FieldLTE(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeContains applies the Contains predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeContains(v string) predicate.Disability {
	return predicate.Disability(sql.FieldContains(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeHasPrefix applies the HasPrefix predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeHasPrefix(v string) predicate.Disability {
	return predicate.Disability(sql.FieldHasPrefix(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeHasSuffix applies the HasSuffix predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeHasSuffix(v string) predicate.Disability {
	return predicate.Disability(sql.FieldHasSuffix(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeEqualFold applies the EqualFold predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeEqualFold(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEqualFold(FieldDisabilityTypeCode, v))
}

// DisabilityTypeCodeContainsFold applies the ContainsFold predicate on the "DisabilityTypeCode" field.
func DisabilityTypeCodeContainsFold(v string) predicate.Disability {
	return predicate.Disability(sql.FieldContainsFold(FieldDisabilityTypeCode, v))
}

// DisabilityTypeDescriptionEQ applies the EQ predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionEQ(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionNEQ applies the NEQ predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionNEQ(v string) predicate.Disability {
	return predicate.Disability(sql.FieldNEQ(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionIn applies the In predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionIn(vs ...string) predicate.Disability {
	return predicate.Disability(sql.FieldIn(FieldDisabilityTypeDescription, vs...))
}

// DisabilityTypeDescriptionNotIn applies the NotIn predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionNotIn(vs ...string) predicate.Disability {
	return predicate.Disability(sql.FieldNotIn(FieldDisabilityTypeDescription, vs...))
}

// DisabilityTypeDescriptionGT applies the GT predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionGT(v string) predicate.Disability {
	return predicate.Disability(sql.FieldGT(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionGTE applies the GTE predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionGTE(v string) predicate.Disability {
	return predicate.Disability(sql.FieldGTE(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionLT applies the LT predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionLT(v string) predicate.Disability {
	return predicate.Disability(sql.FieldLT(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionLTE applies the LTE predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionLTE(v string) predicate.Disability {
	return predicate.Disability(sql.FieldLTE(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionContains applies the Contains predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionContains(v string) predicate.Disability {
	return predicate.Disability(sql.FieldContains(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionHasPrefix applies the HasPrefix predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionHasPrefix(v string) predicate.Disability {
	return predicate.Disability(sql.FieldHasPrefix(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionHasSuffix applies the HasSuffix predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionHasSuffix(v string) predicate.Disability {
	return predicate.Disability(sql.FieldHasSuffix(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionEqualFold applies the EqualFold predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionEqualFold(v string) predicate.Disability {
	return predicate.Disability(sql.FieldEqualFold(FieldDisabilityTypeDescription, v))
}

// DisabilityTypeDescriptionContainsFold applies the ContainsFold predicate on the "DisabilityTypeDescription" field.
func DisabilityTypeDescriptionContainsFold(v string) predicate.Disability {
	return predicate.Disability(sql.FieldContainsFold(FieldDisabilityTypeDescription, v))
}

// DisabilityPercentageEQ applies the EQ predicate on the "DisabilityPercentage" field.
func DisabilityPercentageEQ(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityPercentage, v))
}

// DisabilityPercentageNEQ applies the NEQ predicate on the "DisabilityPercentage" field.
func DisabilityPercentageNEQ(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldNEQ(FieldDisabilityPercentage, v))
}

// DisabilityPercentageIn applies the In predicate on the "DisabilityPercentage" field.
func DisabilityPercentageIn(vs ...int32) predicate.Disability {
	return predicate.Disability(sql.FieldIn(FieldDisabilityPercentage, vs...))
}

// DisabilityPercentageNotIn applies the NotIn predicate on the "DisabilityPercentage" field.
func DisabilityPercentageNotIn(vs ...int32) predicate.Disability {
	return predicate.Disability(sql.FieldNotIn(FieldDisabilityPercentage, vs...))
}

// DisabilityPercentageGT applies the GT predicate on the "DisabilityPercentage" field.
func DisabilityPercentageGT(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldGT(FieldDisabilityPercentage, v))
}

// DisabilityPercentageGTE applies the GTE predicate on the "DisabilityPercentage" field.
func DisabilityPercentageGTE(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldGTE(FieldDisabilityPercentage, v))
}

// DisabilityPercentageLT applies the LT predicate on the "DisabilityPercentage" field.
func DisabilityPercentageLT(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldLT(FieldDisabilityPercentage, v))
}

// DisabilityPercentageLTE applies the LTE predicate on the "DisabilityPercentage" field.
func DisabilityPercentageLTE(v int32) predicate.Disability {
	return predicate.Disability(sql.FieldLTE(FieldDisabilityPercentage, v))
}

// DisabilityFlagEQ applies the EQ predicate on the "DisabilityFlag" field.
func DisabilityFlagEQ(v DisabilityFlag) predicate.Disability {
	return predicate.Disability(sql.FieldEQ(FieldDisabilityFlag, v))
}

// DisabilityFlagNEQ applies the NEQ predicate on the "DisabilityFlag" field.
func DisabilityFlagNEQ(v DisabilityFlag) predicate.Disability {
	return predicate.Disability(sql.FieldNEQ(FieldDisabilityFlag, v))
}

// DisabilityFlagIn applies the In predicate on the "DisabilityFlag" field.
func DisabilityFlagIn(vs ...DisabilityFlag) predicate.Disability {
	return predicate.Disability(sql.FieldIn(FieldDisabilityFlag, vs...))
}

// DisabilityFlagNotIn applies the NotIn predicate on the "DisabilityFlag" field.
func DisabilityFlagNotIn(vs ...DisabilityFlag) predicate.Disability {
	return predicate.Disability(sql.FieldNotIn(FieldDisabilityFlag, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Disability) predicate.Disability {
	return predicate.Disability(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Disability) predicate.Disability {
	return predicate.Disability(func(s *sql.Selector) {
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
func Not(p predicate.Disability) predicate.Disability {
	return predicate.Disability(func(s *sql.Selector) {
		p(s.Not())
	})
}
