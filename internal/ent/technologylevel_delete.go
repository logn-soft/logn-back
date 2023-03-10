// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/logn-soft/logn-back/internal/ent/predicate"
	"github.com/logn-soft/logn-back/internal/ent/technologylevel"
)

// TechnologyLevelDelete is the builder for deleting a TechnologyLevel entity.
type TechnologyLevelDelete struct {
	config
	hooks    []Hook
	mutation *TechnologyLevelMutation
}

// Where appends a list predicates to the TechnologyLevelDelete builder.
func (tld *TechnologyLevelDelete) Where(ps ...predicate.TechnologyLevel) *TechnologyLevelDelete {
	tld.mutation.Where(ps...)
	return tld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tld *TechnologyLevelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TechnologyLevelMutation](ctx, tld.sqlExec, tld.mutation, tld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tld *TechnologyLevelDelete) ExecX(ctx context.Context) int {
	n, err := tld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tld *TechnologyLevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: technologylevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technologylevel.FieldID,
			},
		},
	}
	if ps := tld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tld.mutation.done = true
	return affected, err
}

// TechnologyLevelDeleteOne is the builder for deleting a single TechnologyLevel entity.
type TechnologyLevelDeleteOne struct {
	tld *TechnologyLevelDelete
}

// Where appends a list predicates to the TechnologyLevelDelete builder.
func (tldo *TechnologyLevelDeleteOne) Where(ps ...predicate.TechnologyLevel) *TechnologyLevelDeleteOne {
	tldo.tld.mutation.Where(ps...)
	return tldo
}

// Exec executes the deletion query.
func (tldo *TechnologyLevelDeleteOne) Exec(ctx context.Context) error {
	n, err := tldo.tld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{technologylevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tldo *TechnologyLevelDeleteOne) ExecX(ctx context.Context) {
	if err := tldo.Exec(ctx); err != nil {
		panic(err)
	}
}
