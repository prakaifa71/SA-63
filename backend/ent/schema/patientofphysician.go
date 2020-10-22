package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Patientofphysician holds the schema definition for the Patientofphysician entity.
type Patientofphysician struct {
    ent.Schema
}

// Fields of the Patientofphysician.
func (Patientofphysician) Fields() []ent.Field {
    return []ent.Field{
       field.String("Ailment"),
    }
}

// Edges of the Patientofphysician.
func (Patientofphysician) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("Idpatient", Patient.Type).
            Ref("Patient").
            Unique(),
        edge.From("Roomid", Patientroom.Type).
            Ref("Patientroom").
            Unique(),
        edge.From("Idphysicianid", Physician.Type).
            Ref("Physician").
            Unique(),
    }
}
