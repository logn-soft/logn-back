// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/logn-soft/logn-back/internal/ent/area"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
	"github.com/logn-soft/logn-back/internal/ent/vacancy"
)

// AreaUpdate is the builder for updating Area entities.
type AreaUpdate struct {
	config
	hooks    []Hook
	mutation *AreaMutation
}

// Where appends a list predicates to the AreaUpdate builder.
func (au *AreaUpdate) Where(ps ...predicate.Area) *AreaUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AreaUpdate) SetName(s string) *AreaUpdate {
	au.mutation.SetName(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AreaUpdate) SetCreatedAt(t time.Time) *AreaUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AreaUpdate) SetNillableCreatedAt(t *time.Time) *AreaUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddVacancyIDs adds the "vacancies" edge to the Vacancy entity by IDs.
func (au *AreaUpdate) AddVacancyIDs(ids ...int) *AreaUpdate {
	au.mutation.AddVacancyIDs(ids...)
	return au
}

// AddVacancies adds the "vacancies" edges to the Vacancy entity.
func (au *AreaUpdate) AddVacancies(v ...*Vacancy) *AreaUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.AddVacancyIDs(ids...)
}

// Mutation returns the AreaMutation object of the builder.
func (au *AreaUpdate) Mutation() *AreaMutation {
	return au.mutation
}

// ClearVacancies clears all "vacancies" edges to the Vacancy entity.
func (au *AreaUpdate) ClearVacancies() *AreaUpdate {
	au.mutation.ClearVacancies()
	return au
}

// RemoveVacancyIDs removes the "vacancies" edge to Vacancy entities by IDs.
func (au *AreaUpdate) RemoveVacancyIDs(ids ...int) *AreaUpdate {
	au.mutation.RemoveVacancyIDs(ids...)
	return au
}

// RemoveVacancies removes "vacancies" edges to Vacancy entities.
func (au *AreaUpdate) RemoveVacancies(v ...*Vacancy) *AreaUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.RemoveVacancyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AreaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AreaMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AreaUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AreaUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AreaUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AreaUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := area.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Area.name": %w`, err)}
		}
	}
	return nil
}

func (au *AreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   area.Table,
			Columns: area.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: area.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(area.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(area.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.VacanciesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedVacanciesIDs(); len(nodes) > 0 && !au.mutation.VacanciesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.VacanciesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{area.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AreaUpdateOne is the builder for updating a single Area entity.
type AreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AreaMutation
}

// SetName sets the "name" field.
func (auo *AreaUpdateOne) SetName(s string) *AreaUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AreaUpdateOne) SetCreatedAt(t time.Time) *AreaUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AreaUpdateOne) SetNillableCreatedAt(t *time.Time) *AreaUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddVacancyIDs adds the "vacancies" edge to the Vacancy entity by IDs.
func (auo *AreaUpdateOne) AddVacancyIDs(ids ...int) *AreaUpdateOne {
	auo.mutation.AddVacancyIDs(ids...)
	return auo
}

// AddVacancies adds the "vacancies" edges to the Vacancy entity.
func (auo *AreaUpdateOne) AddVacancies(v ...*Vacancy) *AreaUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.AddVacancyIDs(ids...)
}

// Mutation returns the AreaMutation object of the builder.
func (auo *AreaUpdateOne) Mutation() *AreaMutation {
	return auo.mutation
}

// ClearVacancies clears all "vacancies" edges to the Vacancy entity.
func (auo *AreaUpdateOne) ClearVacancies() *AreaUpdateOne {
	auo.mutation.ClearVacancies()
	return auo
}

// RemoveVacancyIDs removes the "vacancies" edge to Vacancy entities by IDs.
func (auo *AreaUpdateOne) RemoveVacancyIDs(ids ...int) *AreaUpdateOne {
	auo.mutation.RemoveVacancyIDs(ids...)
	return auo
}

// RemoveVacancies removes "vacancies" edges to Vacancy entities.
func (auo *AreaUpdateOne) RemoveVacancies(v ...*Vacancy) *AreaUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.RemoveVacancyIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AreaUpdateOne) Select(field string, fields ...string) *AreaUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Area entity.
func (auo *AreaUpdateOne) Save(ctx context.Context) (*Area, error) {
	return withHooks[*Area, AreaMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AreaUpdateOne) SaveX(ctx context.Context) *Area {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AreaUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AreaUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AreaUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := area.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Area.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AreaUpdateOne) sqlSave(ctx context.Context) (_node *Area, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   area.Table,
			Columns: area.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: area.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Area.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, area.FieldID)
		for _, f := range fields {
			if !area.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != area.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(area.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(area.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.VacanciesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedVacanciesIDs(); len(nodes) > 0 && !auo.mutation.VacanciesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.VacanciesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Area{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{area.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
