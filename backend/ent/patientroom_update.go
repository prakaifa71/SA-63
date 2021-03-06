// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prakaifa21/app/ent/patientofphysician"
	"github.com/prakaifa21/app/ent/patientroom"
	"github.com/prakaifa21/app/ent/predicate"
)

// PatientroomUpdate is the builder for updating Patientroom entities.
type PatientroomUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientroomMutation
	predicates []predicate.Patientroom
}

// Where adds a new predicate for the builder.
func (pu *PatientroomUpdate) Where(ps ...predicate.Patientroom) *PatientroomUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetTyperoom sets the Typeroom field.
func (pu *PatientroomUpdate) SetTyperoom(s string) *PatientroomUpdate {
	pu.mutation.SetTyperoom(s)
	return pu
}

// AddPatientroomIDs adds the Patientroom edge to Patientofphysician by ids.
func (pu *PatientroomUpdate) AddPatientroomIDs(ids ...int) *PatientroomUpdate {
	pu.mutation.AddPatientroomIDs(ids...)
	return pu
}

// AddPatientroom adds the Patientroom edges to Patientofphysician.
func (pu *PatientroomUpdate) AddPatientroom(p ...*Patientofphysician) *PatientroomUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPatientroomIDs(ids...)
}

// Mutation returns the PatientroomMutation object of the builder.
func (pu *PatientroomUpdate) Mutation() *PatientroomMutation {
	return pu.mutation
}

// RemovePatientroomIDs removes the Patientroom edge to Patientofphysician by ids.
func (pu *PatientroomUpdate) RemovePatientroomIDs(ids ...int) *PatientroomUpdate {
	pu.mutation.RemovePatientroomIDs(ids...)
	return pu
}

// RemovePatientroom removes Patientroom edges to Patientofphysician.
func (pu *PatientroomUpdate) RemovePatientroom(p ...*Patientofphysician) *PatientroomUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePatientroomIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientroomUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientroomUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientroomUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientroomUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientroomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientroom.Table,
			Columns: patientroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientroom.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Typeroom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientroom.FieldTyperoom,
		})
	}
	if nodes := pu.mutation.RemovedPatientroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientroom.PatientroomTable,
			Columns: []string{patientroom.PatientroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientofphysician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PatientroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientroom.PatientroomTable,
			Columns: []string{patientroom.PatientroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientofphysician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientroomUpdateOne is the builder for updating a single Patientroom entity.
type PatientroomUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientroomMutation
}

// SetTyperoom sets the Typeroom field.
func (puo *PatientroomUpdateOne) SetTyperoom(s string) *PatientroomUpdateOne {
	puo.mutation.SetTyperoom(s)
	return puo
}

// AddPatientroomIDs adds the Patientroom edge to Patientofphysician by ids.
func (puo *PatientroomUpdateOne) AddPatientroomIDs(ids ...int) *PatientroomUpdateOne {
	puo.mutation.AddPatientroomIDs(ids...)
	return puo
}

// AddPatientroom adds the Patientroom edges to Patientofphysician.
func (puo *PatientroomUpdateOne) AddPatientroom(p ...*Patientofphysician) *PatientroomUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPatientroomIDs(ids...)
}

// Mutation returns the PatientroomMutation object of the builder.
func (puo *PatientroomUpdateOne) Mutation() *PatientroomMutation {
	return puo.mutation
}

// RemovePatientroomIDs removes the Patientroom edge to Patientofphysician by ids.
func (puo *PatientroomUpdateOne) RemovePatientroomIDs(ids ...int) *PatientroomUpdateOne {
	puo.mutation.RemovePatientroomIDs(ids...)
	return puo
}

// RemovePatientroom removes Patientroom edges to Patientofphysician.
func (puo *PatientroomUpdateOne) RemovePatientroom(p ...*Patientofphysician) *PatientroomUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePatientroomIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientroomUpdateOne) Save(ctx context.Context) (*Patientroom, error) {

	var (
		err  error
		node *Patientroom
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientroomUpdateOne) SaveX(ctx context.Context) *Patientroom {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientroomUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientroomUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientroomUpdateOne) sqlSave(ctx context.Context) (pa *Patientroom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientroom.Table,
			Columns: patientroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientroom.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patientroom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Typeroom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientroom.FieldTyperoom,
		})
	}
	if nodes := puo.mutation.RemovedPatientroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientroom.PatientroomTable,
			Columns: []string{patientroom.PatientroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientofphysician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PatientroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientroom.PatientroomTable,
			Columns: []string{patientroom.PatientroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientofphysician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patientroom{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
