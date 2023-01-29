// Code generated by ent, DO NOT EDIT.

package vacancy

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldDescription, v))
}

// IsNegotiate applies equality check predicate on the "is_negotiate" field. It's identical to IsNegotiateEQ.
func IsNegotiate(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldIsNegotiate, v))
}

// MinSalary applies equality check predicate on the "min_salary" field. It's identical to MinSalaryEQ.
func MinSalary(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldMinSalary, v))
}

// MaxSalary applies equality check predicate on the "max_salary" field. It's identical to MaxSalaryEQ.
func MaxSalary(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldMaxSalary, v))
}

// IsRemote applies equality check predicate on the "is_remote" field. It's identical to IsRemoteEQ.
func IsRemote(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldIsRemote, v))
}

// Views applies equality check predicate on the "views" field. It's identical to ViewsEQ.
func Views(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldViews, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldContainsFold(FieldDescription, v))
}

// IsNegotiateEQ applies the EQ predicate on the "is_negotiate" field.
func IsNegotiateEQ(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldIsNegotiate, v))
}

// IsNegotiateNEQ applies the NEQ predicate on the "is_negotiate" field.
func IsNegotiateNEQ(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldIsNegotiate, v))
}

// MinSalaryEQ applies the EQ predicate on the "min_salary" field.
func MinSalaryEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldMinSalary, v))
}

// MinSalaryNEQ applies the NEQ predicate on the "min_salary" field.
func MinSalaryNEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldMinSalary, v))
}

// MinSalaryIn applies the In predicate on the "min_salary" field.
func MinSalaryIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldMinSalary, vs...))
}

// MinSalaryNotIn applies the NotIn predicate on the "min_salary" field.
func MinSalaryNotIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldMinSalary, vs...))
}

// MinSalaryGT applies the GT predicate on the "min_salary" field.
func MinSalaryGT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldMinSalary, v))
}

// MinSalaryGTE applies the GTE predicate on the "min_salary" field.
func MinSalaryGTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldMinSalary, v))
}

// MinSalaryLT applies the LT predicate on the "min_salary" field.
func MinSalaryLT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldMinSalary, v))
}

// MinSalaryLTE applies the LTE predicate on the "min_salary" field.
func MinSalaryLTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldMinSalary, v))
}

// MaxSalaryEQ applies the EQ predicate on the "max_salary" field.
func MaxSalaryEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldMaxSalary, v))
}

// MaxSalaryNEQ applies the NEQ predicate on the "max_salary" field.
func MaxSalaryNEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldMaxSalary, v))
}

// MaxSalaryIn applies the In predicate on the "max_salary" field.
func MaxSalaryIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldMaxSalary, vs...))
}

// MaxSalaryNotIn applies the NotIn predicate on the "max_salary" field.
func MaxSalaryNotIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldMaxSalary, vs...))
}

// MaxSalaryGT applies the GT predicate on the "max_salary" field.
func MaxSalaryGT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldMaxSalary, v))
}

// MaxSalaryGTE applies the GTE predicate on the "max_salary" field.
func MaxSalaryGTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldMaxSalary, v))
}

// MaxSalaryLT applies the LT predicate on the "max_salary" field.
func MaxSalaryLT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldMaxSalary, v))
}

// MaxSalaryLTE applies the LTE predicate on the "max_salary" field.
func MaxSalaryLTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldMaxSalary, v))
}

// IsRemoteEQ applies the EQ predicate on the "is_remote" field.
func IsRemoteEQ(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldIsRemote, v))
}

// IsRemoteNEQ applies the NEQ predicate on the "is_remote" field.
func IsRemoteNEQ(v bool) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldIsRemote, v))
}

// ViewsEQ applies the EQ predicate on the "views" field.
func ViewsEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldViews, v))
}

// ViewsNEQ applies the NEQ predicate on the "views" field.
func ViewsNEQ(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldViews, v))
}

// ViewsIn applies the In predicate on the "views" field.
func ViewsIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldViews, vs...))
}

// ViewsNotIn applies the NotIn predicate on the "views" field.
func ViewsNotIn(vs ...int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldViews, vs...))
}

// ViewsGT applies the GT predicate on the "views" field.
func ViewsGT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldViews, v))
}

// ViewsGTE applies the GTE predicate on the "views" field.
func ViewsGTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldViews, v))
}

// ViewsLT applies the LT predicate on the "views" field.
func ViewsLT(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldViews, v))
}

// ViewsLTE applies the LTE predicate on the "views" field.
func ViewsLTE(v int) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldViews, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Vacancy {
	return predicate.Vacancy(sql.FieldLTE(FieldCreatedAt, v))
}

// HasTechnologies applies the HasEdge predicate on the "technologies" edge.
func HasTechnologies() predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TechnologiesTable, TechnologiesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTechnologiesWith applies the HasEdge predicate on the "technologies" edge with a given conditions (other predicates).
func HasTechnologiesWith(preds ...predicate.Technology) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TechnologiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TechnologiesTable, TechnologiesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocations applies the HasEdge predicate on the "locations" edge.
func HasLocations() predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LocationsTable, LocationsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationsWith applies the HasEdge predicate on the "locations" edge with a given conditions (other predicates).
func HasLocationsWith(preds ...predicate.Location) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LocationsTable, LocationsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAreas applies the HasEdge predicate on the "areas" edge.
func HasAreas() predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AreasTable, AreasPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAreasWith applies the HasEdge predicate on the "areas" edge with a given conditions (other predicates).
func HasAreasWith(preds ...predicate.Area) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AreasInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AreasTable, AreasPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTechnologyLevels applies the HasEdge predicate on the "technology_levels" edge.
func HasTechnologyLevels() predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TechnologyLevelsTable, TechnologyLevelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTechnologyLevelsWith applies the HasEdge predicate on the "technology_levels" edge with a given conditions (other predicates).
func HasTechnologyLevelsWith(preds ...predicate.TechnologyLevel) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TechnologyLevelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TechnologyLevelsTable, TechnologyLevelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vacancy) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vacancy) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
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
func Not(p predicate.Vacancy) predicate.Vacancy {
	return predicate.Vacancy(func(s *sql.Selector) {
		p(s.Not())
	})
}