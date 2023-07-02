// Code generated by ent, DO NOT EDIT.

package employeeposts

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldID, id))
}

// PostCode applies equality check predicate on the "PostCode" field. It's identical to PostCodeEQ.
func PostCode(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPostCode, v))
}

// PostDescription applies equality check predicate on the "PostDescription" field. It's identical to PostDescriptionEQ.
func PostDescription(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPostDescription, v))
}

// Group applies equality check predicate on the "Group" field. It's identical to GroupEQ.
func Group(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldGroup, v))
}

// PayLevel applies equality check predicate on the "PayLevel" field. It's identical to PayLevelEQ.
func PayLevel(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPayLevel, v))
}

// Scale applies equality check predicate on the "Scale" field. It's identical to ScaleEQ.
func Scale(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldScale, v))
}

// BaseCadreFlag applies equality check predicate on the "BaseCadreFlag" field. It's identical to BaseCadreFlagEQ.
func BaseCadreFlag(v bool) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldBaseCadreFlag, v))
}

// PostCodeEQ applies the EQ predicate on the "PostCode" field.
func PostCodeEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPostCode, v))
}

// PostCodeNEQ applies the NEQ predicate on the "PostCode" field.
func PostCodeNEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldPostCode, v))
}

// PostCodeIn applies the In predicate on the "PostCode" field.
func PostCodeIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldPostCode, vs...))
}

// PostCodeNotIn applies the NotIn predicate on the "PostCode" field.
func PostCodeNotIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldPostCode, vs...))
}

// PostCodeGT applies the GT predicate on the "PostCode" field.
func PostCodeGT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldPostCode, v))
}

// PostCodeGTE applies the GTE predicate on the "PostCode" field.
func PostCodeGTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldPostCode, v))
}

// PostCodeLT applies the LT predicate on the "PostCode" field.
func PostCodeLT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldPostCode, v))
}

// PostCodeLTE applies the LTE predicate on the "PostCode" field.
func PostCodeLTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldPostCode, v))
}

// PostCodeContains applies the Contains predicate on the "PostCode" field.
func PostCodeContains(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContains(FieldPostCode, v))
}

// PostCodeHasPrefix applies the HasPrefix predicate on the "PostCode" field.
func PostCodeHasPrefix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasPrefix(FieldPostCode, v))
}

// PostCodeHasSuffix applies the HasSuffix predicate on the "PostCode" field.
func PostCodeHasSuffix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasSuffix(FieldPostCode, v))
}

// PostCodeEqualFold applies the EqualFold predicate on the "PostCode" field.
func PostCodeEqualFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEqualFold(FieldPostCode, v))
}

// PostCodeContainsFold applies the ContainsFold predicate on the "PostCode" field.
func PostCodeContainsFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContainsFold(FieldPostCode, v))
}

// PostDescriptionEQ applies the EQ predicate on the "PostDescription" field.
func PostDescriptionEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPostDescription, v))
}

// PostDescriptionNEQ applies the NEQ predicate on the "PostDescription" field.
func PostDescriptionNEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldPostDescription, v))
}

// PostDescriptionIn applies the In predicate on the "PostDescription" field.
func PostDescriptionIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldPostDescription, vs...))
}

// PostDescriptionNotIn applies the NotIn predicate on the "PostDescription" field.
func PostDescriptionNotIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldPostDescription, vs...))
}

// PostDescriptionGT applies the GT predicate on the "PostDescription" field.
func PostDescriptionGT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldPostDescription, v))
}

// PostDescriptionGTE applies the GTE predicate on the "PostDescription" field.
func PostDescriptionGTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldPostDescription, v))
}

// PostDescriptionLT applies the LT predicate on the "PostDescription" field.
func PostDescriptionLT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldPostDescription, v))
}

// PostDescriptionLTE applies the LTE predicate on the "PostDescription" field.
func PostDescriptionLTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldPostDescription, v))
}

// PostDescriptionContains applies the Contains predicate on the "PostDescription" field.
func PostDescriptionContains(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContains(FieldPostDescription, v))
}

// PostDescriptionHasPrefix applies the HasPrefix predicate on the "PostDescription" field.
func PostDescriptionHasPrefix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasPrefix(FieldPostDescription, v))
}

// PostDescriptionHasSuffix applies the HasSuffix predicate on the "PostDescription" field.
func PostDescriptionHasSuffix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasSuffix(FieldPostDescription, v))
}

// PostDescriptionEqualFold applies the EqualFold predicate on the "PostDescription" field.
func PostDescriptionEqualFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEqualFold(FieldPostDescription, v))
}

// PostDescriptionContainsFold applies the ContainsFold predicate on the "PostDescription" field.
func PostDescriptionContainsFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContainsFold(FieldPostDescription, v))
}

// GroupEQ applies the EQ predicate on the "Group" field.
func GroupEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "Group" field.
func GroupNEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "Group" field.
func GroupIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "Group" field.
func GroupNotIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "Group" field.
func GroupGT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "Group" field.
func GroupGTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "Group" field.
func GroupLT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "Group" field.
func GroupLTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "Group" field.
func GroupContains(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "Group" field.
func GroupHasPrefix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "Group" field.
func GroupHasSuffix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupEqualFold applies the EqualFold predicate on the "Group" field.
func GroupEqualFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "Group" field.
func GroupContainsFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContainsFold(FieldGroup, v))
}

// PayLevelEQ applies the EQ predicate on the "PayLevel" field.
func PayLevelEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldPayLevel, v))
}

// PayLevelNEQ applies the NEQ predicate on the "PayLevel" field.
func PayLevelNEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldPayLevel, v))
}

// PayLevelIn applies the In predicate on the "PayLevel" field.
func PayLevelIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldPayLevel, vs...))
}

// PayLevelNotIn applies the NotIn predicate on the "PayLevel" field.
func PayLevelNotIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldPayLevel, vs...))
}

// PayLevelGT applies the GT predicate on the "PayLevel" field.
func PayLevelGT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldPayLevel, v))
}

// PayLevelGTE applies the GTE predicate on the "PayLevel" field.
func PayLevelGTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldPayLevel, v))
}

// PayLevelLT applies the LT predicate on the "PayLevel" field.
func PayLevelLT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldPayLevel, v))
}

// PayLevelLTE applies the LTE predicate on the "PayLevel" field.
func PayLevelLTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldPayLevel, v))
}

// PayLevelContains applies the Contains predicate on the "PayLevel" field.
func PayLevelContains(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContains(FieldPayLevel, v))
}

// PayLevelHasPrefix applies the HasPrefix predicate on the "PayLevel" field.
func PayLevelHasPrefix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasPrefix(FieldPayLevel, v))
}

// PayLevelHasSuffix applies the HasSuffix predicate on the "PayLevel" field.
func PayLevelHasSuffix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasSuffix(FieldPayLevel, v))
}

// PayLevelEqualFold applies the EqualFold predicate on the "PayLevel" field.
func PayLevelEqualFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEqualFold(FieldPayLevel, v))
}

// PayLevelContainsFold applies the ContainsFold predicate on the "PayLevel" field.
func PayLevelContainsFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContainsFold(FieldPayLevel, v))
}

// ScaleEQ applies the EQ predicate on the "Scale" field.
func ScaleEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldScale, v))
}

// ScaleNEQ applies the NEQ predicate on the "Scale" field.
func ScaleNEQ(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldScale, v))
}

// ScaleIn applies the In predicate on the "Scale" field.
func ScaleIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldIn(FieldScale, vs...))
}

// ScaleNotIn applies the NotIn predicate on the "Scale" field.
func ScaleNotIn(vs ...string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNotIn(FieldScale, vs...))
}

// ScaleGT applies the GT predicate on the "Scale" field.
func ScaleGT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGT(FieldScale, v))
}

// ScaleGTE applies the GTE predicate on the "Scale" field.
func ScaleGTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldGTE(FieldScale, v))
}

// ScaleLT applies the LT predicate on the "Scale" field.
func ScaleLT(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLT(FieldScale, v))
}

// ScaleLTE applies the LTE predicate on the "Scale" field.
func ScaleLTE(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldLTE(FieldScale, v))
}

// ScaleContains applies the Contains predicate on the "Scale" field.
func ScaleContains(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContains(FieldScale, v))
}

// ScaleHasPrefix applies the HasPrefix predicate on the "Scale" field.
func ScaleHasPrefix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasPrefix(FieldScale, v))
}

// ScaleHasSuffix applies the HasSuffix predicate on the "Scale" field.
func ScaleHasSuffix(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldHasSuffix(FieldScale, v))
}

// ScaleEqualFold applies the EqualFold predicate on the "Scale" field.
func ScaleEqualFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEqualFold(FieldScale, v))
}

// ScaleContainsFold applies the ContainsFold predicate on the "Scale" field.
func ScaleContainsFold(v string) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldContainsFold(FieldScale, v))
}

// BaseCadreFlagEQ applies the EQ predicate on the "BaseCadreFlag" field.
func BaseCadreFlagEQ(v bool) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldEQ(FieldBaseCadreFlag, v))
}

// BaseCadreFlagNEQ applies the NEQ predicate on the "BaseCadreFlag" field.
func BaseCadreFlagNEQ(v bool) predicate.EmployeePosts {
	return predicate.EmployeePosts(sql.FieldNEQ(FieldBaseCadreFlag, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmployeePosts) predicate.EmployeePosts {
	return predicate.EmployeePosts(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmployeePosts) predicate.EmployeePosts {
	return predicate.EmployeePosts(func(s *sql.Selector) {
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
func Not(p predicate.EmployeePosts) predicate.EmployeePosts {
	return predicate.EmployeePosts(func(s *sql.Selector) {
		p(s.Not())
	})
}
