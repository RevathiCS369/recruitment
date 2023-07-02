// Code generated by ent, DO NOT EDIT.

package division_choice_pa

import (
	"recruit/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldID, id))
}

// Application applies equality check predicate on the "Application" field. It's identical to ApplicationEQ.
func Application(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldApplication, v))
}

// CadrePrefNo applies equality check predicate on the "CadrePrefNo" field. It's identical to CadrePrefNoEQ.
func CadrePrefNo(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldCadrePrefNo, v))
}

// CadrePrefValue applies equality check predicate on the "CadrePrefValue" field. It's identical to CadrePrefValueEQ.
func CadrePrefValue(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldCadrePrefValue, v))
}

// EmployeeID applies equality check predicate on the "EmployeeID" field. It's identical to EmployeeIDEQ.
func EmployeeID(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldEmployeeID, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "UpdatedBy" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldUpdatedBy, v))
}

// ApplicationEQ applies the EQ predicate on the "Application" field.
func ApplicationEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldApplication, v))
}

// ApplicationNEQ applies the NEQ predicate on the "Application" field.
func ApplicationNEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldApplication, v))
}

// ApplicationIn applies the In predicate on the "Application" field.
func ApplicationIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldApplication, vs...))
}

// ApplicationNotIn applies the NotIn predicate on the "Application" field.
func ApplicationNotIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldApplication, vs...))
}

// ApplicationGT applies the GT predicate on the "Application" field.
func ApplicationGT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldApplication, v))
}

// ApplicationGTE applies the GTE predicate on the "Application" field.
func ApplicationGTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldApplication, v))
}

// ApplicationLT applies the LT predicate on the "Application" field.
func ApplicationLT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldApplication, v))
}

// ApplicationLTE applies the LTE predicate on the "Application" field.
func ApplicationLTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldApplication, v))
}

// ApplicationContains applies the Contains predicate on the "Application" field.
func ApplicationContains(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContains(FieldApplication, v))
}

// ApplicationHasPrefix applies the HasPrefix predicate on the "Application" field.
func ApplicationHasPrefix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasPrefix(FieldApplication, v))
}

// ApplicationHasSuffix applies the HasSuffix predicate on the "Application" field.
func ApplicationHasSuffix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasSuffix(FieldApplication, v))
}

// ApplicationIsNil applies the IsNil predicate on the "Application" field.
func ApplicationIsNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIsNull(FieldApplication))
}

// ApplicationNotNil applies the NotNil predicate on the "Application" field.
func ApplicationNotNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotNull(FieldApplication))
}

// ApplicationEqualFold applies the EqualFold predicate on the "Application" field.
func ApplicationEqualFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEqualFold(FieldApplication, v))
}

// ApplicationContainsFold applies the ContainsFold predicate on the "Application" field.
func ApplicationContainsFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContainsFold(FieldApplication, v))
}

// CadrePrefNoEQ applies the EQ predicate on the "CadrePrefNo" field.
func CadrePrefNoEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldCadrePrefNo, v))
}

// CadrePrefNoNEQ applies the NEQ predicate on the "CadrePrefNo" field.
func CadrePrefNoNEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldCadrePrefNo, v))
}

// CadrePrefNoIn applies the In predicate on the "CadrePrefNo" field.
func CadrePrefNoIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldCadrePrefNo, vs...))
}

// CadrePrefNoNotIn applies the NotIn predicate on the "CadrePrefNo" field.
func CadrePrefNoNotIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldCadrePrefNo, vs...))
}

// CadrePrefNoGT applies the GT predicate on the "CadrePrefNo" field.
func CadrePrefNoGT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldCadrePrefNo, v))
}

// CadrePrefNoGTE applies the GTE predicate on the "CadrePrefNo" field.
func CadrePrefNoGTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldCadrePrefNo, v))
}

// CadrePrefNoLT applies the LT predicate on the "CadrePrefNo" field.
func CadrePrefNoLT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldCadrePrefNo, v))
}

// CadrePrefNoLTE applies the LTE predicate on the "CadrePrefNo" field.
func CadrePrefNoLTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldCadrePrefNo, v))
}

// CadrePrefNoContains applies the Contains predicate on the "CadrePrefNo" field.
func CadrePrefNoContains(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContains(FieldCadrePrefNo, v))
}

// CadrePrefNoHasPrefix applies the HasPrefix predicate on the "CadrePrefNo" field.
func CadrePrefNoHasPrefix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasPrefix(FieldCadrePrefNo, v))
}

// CadrePrefNoHasSuffix applies the HasSuffix predicate on the "CadrePrefNo" field.
func CadrePrefNoHasSuffix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasSuffix(FieldCadrePrefNo, v))
}

// CadrePrefNoEqualFold applies the EqualFold predicate on the "CadrePrefNo" field.
func CadrePrefNoEqualFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEqualFold(FieldCadrePrefNo, v))
}

// CadrePrefNoContainsFold applies the ContainsFold predicate on the "CadrePrefNo" field.
func CadrePrefNoContainsFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContainsFold(FieldCadrePrefNo, v))
}

// CadrePrefValueEQ applies the EQ predicate on the "CadrePrefValue" field.
func CadrePrefValueEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldCadrePrefValue, v))
}

// CadrePrefValueNEQ applies the NEQ predicate on the "CadrePrefValue" field.
func CadrePrefValueNEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldCadrePrefValue, v))
}

// CadrePrefValueIn applies the In predicate on the "CadrePrefValue" field.
func CadrePrefValueIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldCadrePrefValue, vs...))
}

// CadrePrefValueNotIn applies the NotIn predicate on the "CadrePrefValue" field.
func CadrePrefValueNotIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldCadrePrefValue, vs...))
}

// CadrePrefValueGT applies the GT predicate on the "CadrePrefValue" field.
func CadrePrefValueGT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldCadrePrefValue, v))
}

// CadrePrefValueGTE applies the GTE predicate on the "CadrePrefValue" field.
func CadrePrefValueGTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldCadrePrefValue, v))
}

// CadrePrefValueLT applies the LT predicate on the "CadrePrefValue" field.
func CadrePrefValueLT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldCadrePrefValue, v))
}

// CadrePrefValueLTE applies the LTE predicate on the "CadrePrefValue" field.
func CadrePrefValueLTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldCadrePrefValue, v))
}

// CadrePrefValueContains applies the Contains predicate on the "CadrePrefValue" field.
func CadrePrefValueContains(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContains(FieldCadrePrefValue, v))
}

// CadrePrefValueHasPrefix applies the HasPrefix predicate on the "CadrePrefValue" field.
func CadrePrefValueHasPrefix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasPrefix(FieldCadrePrefValue, v))
}

// CadrePrefValueHasSuffix applies the HasSuffix predicate on the "CadrePrefValue" field.
func CadrePrefValueHasSuffix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasSuffix(FieldCadrePrefValue, v))
}

// CadrePrefValueEqualFold applies the EqualFold predicate on the "CadrePrefValue" field.
func CadrePrefValueEqualFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEqualFold(FieldCadrePrefValue, v))
}

// CadrePrefValueContainsFold applies the ContainsFold predicate on the "CadrePrefValue" field.
func CadrePrefValueContainsFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContainsFold(FieldCadrePrefValue, v))
}

// EmployeeIDEQ applies the EQ predicate on the "EmployeeID" field.
func EmployeeIDEQ(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "EmployeeID" field.
func EmployeeIDNEQ(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "EmployeeID" field.
func EmployeeIDIn(vs ...int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "EmployeeID" field.
func EmployeeIDNotIn(vs ...int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// EmployeeIDGT applies the GT predicate on the "EmployeeID" field.
func EmployeeIDGT(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldEmployeeID, v))
}

// EmployeeIDGTE applies the GTE predicate on the "EmployeeID" field.
func EmployeeIDGTE(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldEmployeeID, v))
}

// EmployeeIDLT applies the LT predicate on the "EmployeeID" field.
func EmployeeIDLT(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldEmployeeID, v))
}

// EmployeeIDLTE applies the LTE predicate on the "EmployeeID" field.
func EmployeeIDLTE(v int64) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldEmployeeID, v))
}

// EmployeeIDIsNil applies the IsNil predicate on the "EmployeeID" field.
func EmployeeIDIsNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIsNull(FieldEmployeeID))
}

// EmployeeIDNotNil applies the NotNil predicate on the "EmployeeID" field.
func EmployeeIDNotNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotNull(FieldEmployeeID))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "UpdatedAt" field.
func UpdatedAtIsNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "UpdatedAt" field.
func UpdatedAtNotNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "UpdatedBy" field.
func UpdatedByEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "UpdatedBy" field.
func UpdatedByNEQ(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "UpdatedBy" field.
func UpdatedByIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "UpdatedBy" field.
func UpdatedByNotIn(vs ...string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "UpdatedBy" field.
func UpdatedByGT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "UpdatedBy" field.
func UpdatedByGTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "UpdatedBy" field.
func UpdatedByLT(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "UpdatedBy" field.
func UpdatedByLTE(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "UpdatedBy" field.
func UpdatedByContains(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "UpdatedBy" field.
func UpdatedByHasPrefix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "UpdatedBy" field.
func UpdatedByHasSuffix(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "UpdatedBy" field.
func UpdatedByIsNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "UpdatedBy" field.
func UpdatedByNotNil() predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "UpdatedBy" field.
func UpdatedByEqualFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "UpdatedBy" field.
func UpdatedByContainsFold(v string) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Division_Choice_PA) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Division_Choice_PA) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(func(s *sql.Selector) {
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
func Not(p predicate.Division_Choice_PA) predicate.Division_Choice_PA {
	return predicate.Division_Choice_PA(func(s *sql.Selector) {
		p(s.Not())
	})
}
