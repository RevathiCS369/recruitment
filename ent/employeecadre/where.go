// Code generated by ent, DO NOT EDIT.

package employeecadre

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLTE(FieldID, id))
}

// Cadrecode applies equality check predicate on the "cadrecode" field. It's identical to CadrecodeEQ.
func Cadrecode(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldCadrecode, v))
}

// Cadredescription applies equality check predicate on the "cadredescription" field. It's identical to CadredescriptionEQ.
func Cadredescription(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldCadredescription, v))
}

// PayLevel applies equality check predicate on the "PayLevel" field. It's identical to PayLevelEQ.
func PayLevel(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldPayLevel, v))
}

// Scale applies equality check predicate on the "Scale" field. It's identical to ScaleEQ.
func Scale(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldScale, v))
}

// CadrecodeEQ applies the EQ predicate on the "cadrecode" field.
func CadrecodeEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldCadrecode, v))
}

// CadrecodeNEQ applies the NEQ predicate on the "cadrecode" field.
func CadrecodeNEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNEQ(FieldCadrecode, v))
}

// CadrecodeIn applies the In predicate on the "cadrecode" field.
func CadrecodeIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldIn(FieldCadrecode, vs...))
}

// CadrecodeNotIn applies the NotIn predicate on the "cadrecode" field.
func CadrecodeNotIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNotIn(FieldCadrecode, vs...))
}

// CadrecodeGT applies the GT predicate on the "cadrecode" field.
func CadrecodeGT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGT(FieldCadrecode, v))
}

// CadrecodeGTE applies the GTE predicate on the "cadrecode" field.
func CadrecodeGTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGTE(FieldCadrecode, v))
}

// CadrecodeLT applies the LT predicate on the "cadrecode" field.
func CadrecodeLT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLT(FieldCadrecode, v))
}

// CadrecodeLTE applies the LTE predicate on the "cadrecode" field.
func CadrecodeLTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLTE(FieldCadrecode, v))
}

// CadrecodeContains applies the Contains predicate on the "cadrecode" field.
func CadrecodeContains(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContains(FieldCadrecode, v))
}

// CadrecodeHasPrefix applies the HasPrefix predicate on the "cadrecode" field.
func CadrecodeHasPrefix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasPrefix(FieldCadrecode, v))
}

// CadrecodeHasSuffix applies the HasSuffix predicate on the "cadrecode" field.
func CadrecodeHasSuffix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasSuffix(FieldCadrecode, v))
}

// CadrecodeEqualFold applies the EqualFold predicate on the "cadrecode" field.
func CadrecodeEqualFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEqualFold(FieldCadrecode, v))
}

// CadrecodeContainsFold applies the ContainsFold predicate on the "cadrecode" field.
func CadrecodeContainsFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContainsFold(FieldCadrecode, v))
}

// CadredescriptionEQ applies the EQ predicate on the "cadredescription" field.
func CadredescriptionEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldCadredescription, v))
}

// CadredescriptionNEQ applies the NEQ predicate on the "cadredescription" field.
func CadredescriptionNEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNEQ(FieldCadredescription, v))
}

// CadredescriptionIn applies the In predicate on the "cadredescription" field.
func CadredescriptionIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldIn(FieldCadredescription, vs...))
}

// CadredescriptionNotIn applies the NotIn predicate on the "cadredescription" field.
func CadredescriptionNotIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNotIn(FieldCadredescription, vs...))
}

// CadredescriptionGT applies the GT predicate on the "cadredescription" field.
func CadredescriptionGT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGT(FieldCadredescription, v))
}

// CadredescriptionGTE applies the GTE predicate on the "cadredescription" field.
func CadredescriptionGTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGTE(FieldCadredescription, v))
}

// CadredescriptionLT applies the LT predicate on the "cadredescription" field.
func CadredescriptionLT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLT(FieldCadredescription, v))
}

// CadredescriptionLTE applies the LTE predicate on the "cadredescription" field.
func CadredescriptionLTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLTE(FieldCadredescription, v))
}

// CadredescriptionContains applies the Contains predicate on the "cadredescription" field.
func CadredescriptionContains(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContains(FieldCadredescription, v))
}

// CadredescriptionHasPrefix applies the HasPrefix predicate on the "cadredescription" field.
func CadredescriptionHasPrefix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasPrefix(FieldCadredescription, v))
}

// CadredescriptionHasSuffix applies the HasSuffix predicate on the "cadredescription" field.
func CadredescriptionHasSuffix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasSuffix(FieldCadredescription, v))
}

// CadredescriptionEqualFold applies the EqualFold predicate on the "cadredescription" field.
func CadredescriptionEqualFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEqualFold(FieldCadredescription, v))
}

// CadredescriptionContainsFold applies the ContainsFold predicate on the "cadredescription" field.
func CadredescriptionContainsFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContainsFold(FieldCadredescription, v))
}

// PayLevelEQ applies the EQ predicate on the "PayLevel" field.
func PayLevelEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldPayLevel, v))
}

// PayLevelNEQ applies the NEQ predicate on the "PayLevel" field.
func PayLevelNEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNEQ(FieldPayLevel, v))
}

// PayLevelIn applies the In predicate on the "PayLevel" field.
func PayLevelIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldIn(FieldPayLevel, vs...))
}

// PayLevelNotIn applies the NotIn predicate on the "PayLevel" field.
func PayLevelNotIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNotIn(FieldPayLevel, vs...))
}

// PayLevelGT applies the GT predicate on the "PayLevel" field.
func PayLevelGT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGT(FieldPayLevel, v))
}

// PayLevelGTE applies the GTE predicate on the "PayLevel" field.
func PayLevelGTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGTE(FieldPayLevel, v))
}

// PayLevelLT applies the LT predicate on the "PayLevel" field.
func PayLevelLT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLT(FieldPayLevel, v))
}

// PayLevelLTE applies the LTE predicate on the "PayLevel" field.
func PayLevelLTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLTE(FieldPayLevel, v))
}

// PayLevelContains applies the Contains predicate on the "PayLevel" field.
func PayLevelContains(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContains(FieldPayLevel, v))
}

// PayLevelHasPrefix applies the HasPrefix predicate on the "PayLevel" field.
func PayLevelHasPrefix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasPrefix(FieldPayLevel, v))
}

// PayLevelHasSuffix applies the HasSuffix predicate on the "PayLevel" field.
func PayLevelHasSuffix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasSuffix(FieldPayLevel, v))
}

// PayLevelEqualFold applies the EqualFold predicate on the "PayLevel" field.
func PayLevelEqualFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEqualFold(FieldPayLevel, v))
}

// PayLevelContainsFold applies the ContainsFold predicate on the "PayLevel" field.
func PayLevelContainsFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContainsFold(FieldPayLevel, v))
}

// ScaleEQ applies the EQ predicate on the "Scale" field.
func ScaleEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEQ(FieldScale, v))
}

// ScaleNEQ applies the NEQ predicate on the "Scale" field.
func ScaleNEQ(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNEQ(FieldScale, v))
}

// ScaleIn applies the In predicate on the "Scale" field.
func ScaleIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldIn(FieldScale, vs...))
}

// ScaleNotIn applies the NotIn predicate on the "Scale" field.
func ScaleNotIn(vs ...string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldNotIn(FieldScale, vs...))
}

// ScaleGT applies the GT predicate on the "Scale" field.
func ScaleGT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGT(FieldScale, v))
}

// ScaleGTE applies the GTE predicate on the "Scale" field.
func ScaleGTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldGTE(FieldScale, v))
}

// ScaleLT applies the LT predicate on the "Scale" field.
func ScaleLT(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLT(FieldScale, v))
}

// ScaleLTE applies the LTE predicate on the "Scale" field.
func ScaleLTE(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldLTE(FieldScale, v))
}

// ScaleContains applies the Contains predicate on the "Scale" field.
func ScaleContains(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContains(FieldScale, v))
}

// ScaleHasPrefix applies the HasPrefix predicate on the "Scale" field.
func ScaleHasPrefix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasPrefix(FieldScale, v))
}

// ScaleHasSuffix applies the HasSuffix predicate on the "Scale" field.
func ScaleHasSuffix(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldHasSuffix(FieldScale, v))
}

// ScaleEqualFold applies the EqualFold predicate on the "Scale" field.
func ScaleEqualFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldEqualFold(FieldScale, v))
}

// ScaleContainsFold applies the ContainsFold predicate on the "Scale" field.
func ScaleContainsFold(v string) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(sql.FieldContainsFold(FieldScale, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmployeeCadre) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmployeeCadre) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(func(s *sql.Selector) {
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
func Not(p predicate.EmployeeCadre) predicate.EmployeeCadre {
	return predicate.EmployeeCadre(func(s *sql.Selector) {
		p(s.Not())
	})
}
