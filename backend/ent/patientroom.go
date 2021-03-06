// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/prakaifa21/app/ent/patientroom"
)

// Patientroom is the model entity for the Patientroom schema.
type Patientroom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Typeroom holds the value of the "Typeroom" field.
	Typeroom string `json:"Typeroom,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientroomQuery when eager-loading is set.
	Edges PatientroomEdges `json:"edges"`
}

// PatientroomEdges holds the relations/edges for other nodes in the graph.
type PatientroomEdges struct {
	// Patientroom holds the value of the Patientroom edge.
	Patientroom []*Patientofphysician
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PatientroomOrErr returns the Patientroom value or an error if the edge
// was not loaded in eager-loading.
func (e PatientroomEdges) PatientroomOrErr() ([]*Patientofphysician, error) {
	if e.loadedTypes[0] {
		return e.Patientroom, nil
	}
	return nil, &NotLoadedError{edge: "Patientroom"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patientroom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Typeroom
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patientroom fields.
func (pa *Patientroom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patientroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Typeroom", values[0])
	} else if value.Valid {
		pa.Typeroom = value.String
	}
	return nil
}

// QueryPatientroom queries the Patientroom edge of the Patientroom.
func (pa *Patientroom) QueryPatientroom() *PatientofphysicianQuery {
	return (&PatientroomClient{config: pa.config}).QueryPatientroom(pa)
}

// Update returns a builder for updating this Patientroom.
// Note that, you need to call Patientroom.Unwrap() before calling this method, if this Patientroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patientroom) Update() *PatientroomUpdateOne {
	return (&PatientroomClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patientroom) Unwrap() *Patientroom {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patientroom is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patientroom) String() string {
	var builder strings.Builder
	builder.WriteString("Patientroom(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", Typeroom=")
	builder.WriteString(pa.Typeroom)
	builder.WriteByte(')')
	return builder.String()
}

// Patientrooms is a parsable slice of Patientroom.
type Patientrooms []*Patientroom

func (pa Patientrooms) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
