// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/logn-soft/logn-back/internal/ent/area"
	"github.com/logn-soft/logn-back/internal/ent/vacancy"
)

// AreaCreate is the builder for creating a Area entity.
type AreaCreate struct {
	config
	mutation *AreaMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AreaCreate) SetName(s string) *AreaCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AreaCreate) SetCreatedAt(t time.Time) *AreaCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AreaCreate) SetNillableCreatedAt(t *time.Time) *AreaCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// AddVacancyIDs adds the "vacancies" edge to the Vacancy entity by IDs.
func (ac *AreaCreate) AddVacancyIDs(ids ...int) *AreaCreate {
	ac.mutation.AddVacancyIDs(ids...)
	return ac
}

// AddVacancies adds the "vacancies" edges to the Vacancy entity.
func (ac *AreaCreate) AddVacancies(v ...*Vacancy) *AreaCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ac.AddVacancyIDs(ids...)
}

// Mutation returns the AreaMutation object of the builder.
func (ac *AreaCreate) Mutation() *AreaMutation {
	return ac.mutation
}

// Save creates the Area in the database.
func (ac *AreaCreate) Save(ctx context.Context) (*Area, error) {
	ac.defaults()
	return withHooks[*Area, AreaMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AreaCreate) SaveX(ctx context.Context) *Area {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AreaCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AreaCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AreaCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := area.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AreaCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Area.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := area.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Area.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Area.created_at"`)}
	}
	return nil
}

func (ac *AreaCreate) sqlSave(ctx context.Context) (*Area, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AreaCreate) createSpec() (*Area, *sqlgraph.CreateSpec) {
	var (
		_node = &Area{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: area.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: area.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(area.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(area.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.VacanciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   area.VacanciesTable,
			Columns: area.VacanciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vacancy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AreaCreateBulk is the builder for creating many Area entities in bulk.
type AreaCreateBulk struct {
	config
	builders []*AreaCreate
}

// Save creates the Area entities in the database.
func (acb *AreaCreateBulk) Save(ctx context.Context) ([]*Area, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Area, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AreaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AreaCreateBulk) SaveX(ctx context.Context) []*Area {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AreaCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AreaCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
