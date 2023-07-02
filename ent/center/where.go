// Code generated by ent, DO NOT EDIT.

package center

import (
	"recruit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Center {
	return predicate.Center(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Center {
	return predicate.Center(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Center {
	return predicate.Center(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Center {
	return predicate.Center(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Center {
	return predicate.Center(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Center {
	return predicate.Center(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Center {
	return predicate.Center(sql.FieldLTE(FieldID, id))
}

// NotifyCode applies equality check predicate on the "NotifyCode" field. It's identical to NotifyCodeEQ.
func NotifyCode(v int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldNotifyCode, v))
}

// NodalOfficerCode applies equality check predicate on the "NodalOfficerCode" field. It's identical to NodalOfficerCodeEQ.
func NodalOfficerCode(v int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldNodalOfficerCode, v))
}

// CenterName applies equality check predicate on the "CenterName" field. It's identical to CenterNameEQ.
func CenterName(v string) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldCenterName, v))
}

// NotifyCodeEQ applies the EQ predicate on the "NotifyCode" field.
func NotifyCodeEQ(v int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldNotifyCode, v))
}

// NotifyCodeNEQ applies the NEQ predicate on the "NotifyCode" field.
func NotifyCodeNEQ(v int32) predicate.Center {
	return predicate.Center(sql.FieldNEQ(FieldNotifyCode, v))
}

// NotifyCodeIn applies the In predicate on the "NotifyCode" field.
func NotifyCodeIn(vs ...int32) predicate.Center {
	return predicate.Center(sql.FieldIn(FieldNotifyCode, vs...))
}

// NotifyCodeNotIn applies the NotIn predicate on the "NotifyCode" field.
func NotifyCodeNotIn(vs ...int32) predicate.Center {
	return predicate.Center(sql.FieldNotIn(FieldNotifyCode, vs...))
}

// NotifyCodeIsNil applies the IsNil predicate on the "NotifyCode" field.
func NotifyCodeIsNil() predicate.Center {
	return predicate.Center(sql.FieldIsNull(FieldNotifyCode))
}

// NotifyCodeNotNil applies the NotNil predicate on the "NotifyCode" field.
func NotifyCodeNotNil() predicate.Center {
	return predicate.Center(sql.FieldNotNull(FieldNotifyCode))
}

// NodalOfficerCodeEQ applies the EQ predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeEQ(v int32) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldNodalOfficerCode, v))
}

// NodalOfficerCodeNEQ applies the NEQ predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeNEQ(v int32) predicate.Center {
	return predicate.Center(sql.FieldNEQ(FieldNodalOfficerCode, v))
}

// NodalOfficerCodeIn applies the In predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeIn(vs ...int32) predicate.Center {
	return predicate.Center(sql.FieldIn(FieldNodalOfficerCode, vs...))
}

// NodalOfficerCodeNotIn applies the NotIn predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeNotIn(vs ...int32) predicate.Center {
	return predicate.Center(sql.FieldNotIn(FieldNodalOfficerCode, vs...))
}

// NodalOfficerCodeIsNil applies the IsNil predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeIsNil() predicate.Center {
	return predicate.Center(sql.FieldIsNull(FieldNodalOfficerCode))
}

// NodalOfficerCodeNotNil applies the NotNil predicate on the "NodalOfficerCode" field.
func NodalOfficerCodeNotNil() predicate.Center {
	return predicate.Center(sql.FieldNotNull(FieldNodalOfficerCode))
}

// CenterNameEQ applies the EQ predicate on the "CenterName" field.
func CenterNameEQ(v string) predicate.Center {
	return predicate.Center(sql.FieldEQ(FieldCenterName, v))
}

// CenterNameNEQ applies the NEQ predicate on the "CenterName" field.
func CenterNameNEQ(v string) predicate.Center {
	return predicate.Center(sql.FieldNEQ(FieldCenterName, v))
}

// CenterNameIn applies the In predicate on the "CenterName" field.
func CenterNameIn(vs ...string) predicate.Center {
	return predicate.Center(sql.FieldIn(FieldCenterName, vs...))
}

// CenterNameNotIn applies the NotIn predicate on the "CenterName" field.
func CenterNameNotIn(vs ...string) predicate.Center {
	return predicate.Center(sql.FieldNotIn(FieldCenterName, vs...))
}

// CenterNameGT applies the GT predicate on the "CenterName" field.
func CenterNameGT(v string) predicate.Center {
	return predicate.Center(sql.FieldGT(FieldCenterName, v))
}

// CenterNameGTE applies the GTE predicate on the "CenterName" field.
func CenterNameGTE(v string) predicate.Center {
	return predicate.Center(sql.FieldGTE(FieldCenterName, v))
}

// CenterNameLT applies the LT predicate on the "CenterName" field.
func CenterNameLT(v string) predicate.Center {
	return predicate.Center(sql.FieldLT(FieldCenterName, v))
}

// CenterNameLTE applies the LTE predicate on the "CenterName" field.
func CenterNameLTE(v string) predicate.Center {
	return predicate.Center(sql.FieldLTE(FieldCenterName, v))
}

// CenterNameContains applies the Contains predicate on the "CenterName" field.
func CenterNameContains(v string) predicate.Center {
	return predicate.Center(sql.FieldContains(FieldCenterName, v))
}

// CenterNameHasPrefix applies the HasPrefix predicate on the "CenterName" field.
func CenterNameHasPrefix(v string) predicate.Center {
	return predicate.Center(sql.FieldHasPrefix(FieldCenterName, v))
}

// CenterNameHasSuffix applies the HasSuffix predicate on the "CenterName" field.
func CenterNameHasSuffix(v string) predicate.Center {
	return predicate.Center(sql.FieldHasSuffix(FieldCenterName, v))
}

// CenterNameEqualFold applies the EqualFold predicate on the "CenterName" field.
func CenterNameEqualFold(v string) predicate.Center {
	return predicate.Center(sql.FieldEqualFold(FieldCenterName, v))
}

// CenterNameContainsFold applies the ContainsFold predicate on the "CenterName" field.
func CenterNameContainsFold(v string) predicate.Center {
	return predicate.Center(sql.FieldContainsFold(FieldCenterName, v))
}

// HasApplications applies the HasEdge predicate on the "applications" edge.
func HasApplications() predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ApplicationsTable, ApplicationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicationsWith applies the HasEdge predicate on the "applications" edge with a given conditions (other predicates).
func HasApplicationsWith(preds ...predicate.Application) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := newApplicationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodalOfficer applies the HasEdge predicate on the "nodal_officer" edge.
func HasNodalOfficer() predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodalOfficerTable, NodalOfficerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodalOfficerWith applies the HasEdge predicate on the "nodal_officer" edge with a given conditions (other predicates).
func HasNodalOfficerWith(preds ...predicate.NodalOfficer) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := newNodalOfficerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotification applies the HasEdge predicate on the "notification" edge.
func HasNotification() predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NotificationTable, NotificationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationWith applies the HasEdge predicate on the "notification" edge with a given conditions (other predicates).
func HasNotificationWith(preds ...predicate.Notification) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		step := newNotificationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Center) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Center) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
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
func Not(p predicate.Center) predicate.Center {
	return predicate.Center(func(s *sql.Selector) {
		p(s.Not())
	})
}