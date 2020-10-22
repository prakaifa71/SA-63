// Code generated by entc, DO NOT EDIT.

package patient

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/prakaifa21/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Patientname applies equality check predicate on the "Patientname" field. It's identical to PatientnameEQ.
func Patientname(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientname), v))
	})
}

// PatientnameEQ applies the EQ predicate on the "Patientname" field.
func PatientnameEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientname), v))
	})
}

// PatientnameNEQ applies the NEQ predicate on the "Patientname" field.
func PatientnameNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientname), v))
	})
}

// PatientnameIn applies the In predicate on the "Patientname" field.
func PatientnameIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientname), v...))
	})
}

// PatientnameNotIn applies the NotIn predicate on the "Patientname" field.
func PatientnameNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientname), v...))
	})
}

// PatientnameGT applies the GT predicate on the "Patientname" field.
func PatientnameGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientname), v))
	})
}

// PatientnameGTE applies the GTE predicate on the "Patientname" field.
func PatientnameGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientname), v))
	})
}

// PatientnameLT applies the LT predicate on the "Patientname" field.
func PatientnameLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientname), v))
	})
}

// PatientnameLTE applies the LTE predicate on the "Patientname" field.
func PatientnameLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientname), v))
	})
}

// PatientnameContains applies the Contains predicate on the "Patientname" field.
func PatientnameContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPatientname), v))
	})
}

// PatientnameHasPrefix applies the HasPrefix predicate on the "Patientname" field.
func PatientnameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPatientname), v))
	})
}

// PatientnameHasSuffix applies the HasSuffix predicate on the "Patientname" field.
func PatientnameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPatientname), v))
	})
}

// PatientnameEqualFold applies the EqualFold predicate on the "Patientname" field.
func PatientnameEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPatientname), v))
	})
}

// PatientnameContainsFold applies the ContainsFold predicate on the "Patientname" field.
func PatientnameContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPatientname), v))
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patientofphysician) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		p(s.Not())
	})
}
