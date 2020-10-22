package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Patientroom holds the schema definition for the Patientroom entity.
type Patientroom struct {
    ent.Schema
}

// Fields of the Patientroom.
func (Patientroom) Fields() []ent.Field {
    return []ent.Field{
        field.String("Typeroom").
        Unique(),
        
    } 
}

// Edges of the Patientroom.
func (Patientroom) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Patientroom", Patientofphysician.Type),
    }
}
