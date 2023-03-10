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
	"github.com/logn-soft/logn-back/internal/ent/location"
	"github.com/logn-soft/logn-back/internal/ent/technology"
	"github.com/logn-soft/logn-back/internal/ent/technologylevel"
	"github.com/logn-soft/logn-back/internal/ent/vacancy"
)

// VacancyCreate is the builder for creating a Vacancy entity.
type VacancyCreate struct {
	config
	mutation *VacancyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vc *VacancyCreate) SetName(s string) *VacancyCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VacancyCreate) SetDescription(s string) *VacancyCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetIsNegotiate sets the "is_negotiate" field.
func (vc *VacancyCreate) SetIsNegotiate(b bool) *VacancyCreate {
	vc.mutation.SetIsNegotiate(b)
	return vc
}

// SetMinSalary sets the "min_salary" field.
func (vc *VacancyCreate) SetMinSalary(i int) *VacancyCreate {
	vc.mutation.SetMinSalary(i)
	return vc
}

// SetMaxSalary sets the "max_salary" field.
func (vc *VacancyCreate) SetMaxSalary(i int) *VacancyCreate {
	vc.mutation.SetMaxSalary(i)
	return vc
}

// SetIsRemote sets the "is_remote" field.
func (vc *VacancyCreate) SetIsRemote(b bool) *VacancyCreate {
	vc.mutation.SetIsRemote(b)
	return vc
}

// SetViews sets the "views" field.
func (vc *VacancyCreate) SetViews(i int) *VacancyCreate {
	vc.mutation.SetViews(i)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VacancyCreate) SetCreatedAt(t time.Time) *VacancyCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VacancyCreate) SetNillableCreatedAt(t *time.Time) *VacancyCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// AddTechnologyIDs adds the "technologies" edge to the Technology entity by IDs.
func (vc *VacancyCreate) AddTechnologyIDs(ids ...int) *VacancyCreate {
	vc.mutation.AddTechnologyIDs(ids...)
	return vc
}

// AddTechnologies adds the "technologies" edges to the Technology entity.
func (vc *VacancyCreate) AddTechnologies(t ...*Technology) *VacancyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vc.AddTechnologyIDs(ids...)
}

// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
func (vc *VacancyCreate) AddLocationIDs(ids ...int) *VacancyCreate {
	vc.mutation.AddLocationIDs(ids...)
	return vc
}

// AddLocations adds the "locations" edges to the Location entity.
func (vc *VacancyCreate) AddLocations(l ...*Location) *VacancyCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vc.AddLocationIDs(ids...)
}

// AddAreaIDs adds the "areas" edge to the Area entity by IDs.
func (vc *VacancyCreate) AddAreaIDs(ids ...int) *VacancyCreate {
	vc.mutation.AddAreaIDs(ids...)
	return vc
}

// AddAreas adds the "areas" edges to the Area entity.
func (vc *VacancyCreate) AddAreas(a ...*Area) *VacancyCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vc.AddAreaIDs(ids...)
}

// AddTechnologyLevelIDs adds the "technology_levels" edge to the TechnologyLevel entity by IDs.
func (vc *VacancyCreate) AddTechnologyLevelIDs(ids ...int) *VacancyCreate {
	vc.mutation.AddTechnologyLevelIDs(ids...)
	return vc
}

// AddTechnologyLevels adds the "technology_levels" edges to the TechnologyLevel entity.
func (vc *VacancyCreate) AddTechnologyLevels(t ...*TechnologyLevel) *VacancyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return vc.AddTechnologyLevelIDs(ids...)
}

// Mutation returns the VacancyMutation object of the builder.
func (vc *VacancyCreate) Mutation() *VacancyMutation {
	return vc.mutation
}

// Save creates the Vacancy in the database.
func (vc *VacancyCreate) Save(ctx context.Context) (*Vacancy, error) {
	vc.defaults()
	return withHooks[*Vacancy, VacancyMutation](ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VacancyCreate) SaveX(ctx context.Context) *Vacancy {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VacancyCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VacancyCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VacancyCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := vacancy.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VacancyCreate) check() error {
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Vacancy.name"`)}
	}
	if v, ok := vc.mutation.Name(); ok {
		if err := vacancy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vacancy.name": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Vacancy.description"`)}
	}
	if v, ok := vc.mutation.Description(); ok {
		if err := vacancy.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Vacancy.description": %w`, err)}
		}
	}
	if _, ok := vc.mutation.IsNegotiate(); !ok {
		return &ValidationError{Name: "is_negotiate", err: errors.New(`ent: missing required field "Vacancy.is_negotiate"`)}
	}
	if _, ok := vc.mutation.MinSalary(); !ok {
		return &ValidationError{Name: "min_salary", err: errors.New(`ent: missing required field "Vacancy.min_salary"`)}
	}
	if v, ok := vc.mutation.MinSalary(); ok {
		if err := vacancy.MinSalaryValidator(v); err != nil {
			return &ValidationError{Name: "min_salary", err: fmt.Errorf(`ent: validator failed for field "Vacancy.min_salary": %w`, err)}
		}
	}
	if _, ok := vc.mutation.MaxSalary(); !ok {
		return &ValidationError{Name: "max_salary", err: errors.New(`ent: missing required field "Vacancy.max_salary"`)}
	}
	if _, ok := vc.mutation.IsRemote(); !ok {
		return &ValidationError{Name: "is_remote", err: errors.New(`ent: missing required field "Vacancy.is_remote"`)}
	}
	if _, ok := vc.mutation.Views(); !ok {
		return &ValidationError{Name: "views", err: errors.New(`ent: missing required field "Vacancy.views"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Vacancy.created_at"`)}
	}
	return nil
}

func (vc *VacancyCreate) sqlSave(ctx context.Context) (*Vacancy, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VacancyCreate) createSpec() (*Vacancy, *sqlgraph.CreateSpec) {
	var (
		_node = &Vacancy{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vacancy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vacancy.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(vacancy.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(vacancy.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.IsNegotiate(); ok {
		_spec.SetField(vacancy.FieldIsNegotiate, field.TypeBool, value)
		_node.IsNegotiate = value
	}
	if value, ok := vc.mutation.MinSalary(); ok {
		_spec.SetField(vacancy.FieldMinSalary, field.TypeInt, value)
		_node.MinSalary = value
	}
	if value, ok := vc.mutation.MaxSalary(); ok {
		_spec.SetField(vacancy.FieldMaxSalary, field.TypeInt, value)
		_node.MaxSalary = value
	}
	if value, ok := vc.mutation.IsRemote(); ok {
		_spec.SetField(vacancy.FieldIsRemote, field.TypeBool, value)
		_node.IsRemote = value
	}
	if value, ok := vc.mutation.Views(); ok {
		_spec.SetField(vacancy.FieldViews, field.TypeInt, value)
		_node.Views = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(vacancy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := vc.mutation.TechnologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.TechnologiesTable,
			Columns: vacancy.TechnologiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TechnologyLevelCreate{config: vc.config, mutation: newTechnologyLevelMutation(vc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.LocationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.LocationsTable,
			Columns: vacancy.LocationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.AreasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   vacancy.AreasTable,
			Columns: vacancy.AreasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.TechnologyLevelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   vacancy.TechnologyLevelsTable,
			Columns: []string{vacancy.TechnologyLevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologylevel.FieldID,
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

// VacancyCreateBulk is the builder for creating many Vacancy entities in bulk.
type VacancyCreateBulk struct {
	config
	builders []*VacancyCreate
}

// Save creates the Vacancy entities in the database.
func (vcb *VacancyCreateBulk) Save(ctx context.Context) ([]*Vacancy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vacancy, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VacancyMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VacancyCreateBulk) SaveX(ctx context.Context) []*Vacancy {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VacancyCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VacancyCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
