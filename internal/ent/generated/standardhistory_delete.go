// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
	"github.com/theopenlane/core/internal/ent/generated/standardhistory"
)

// StandardHistoryDelete is the builder for deleting a StandardHistory entity.
type StandardHistoryDelete struct {
	config
	hooks    []Hook
	mutation *StandardHistoryMutation
}

// Where appends a list predicates to the StandardHistoryDelete builder.
func (shd *StandardHistoryDelete) Where(ps ...predicate.StandardHistory) *StandardHistoryDelete {
	shd.mutation.Where(ps...)
	return shd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (shd *StandardHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, shd.sqlExec, shd.mutation, shd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (shd *StandardHistoryDelete) ExecX(ctx context.Context) int {
	n, err := shd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (shd *StandardHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(standardhistory.Table, sqlgraph.NewFieldSpec(standardhistory.FieldID, field.TypeString))
	_spec.Node.Schema = shd.schemaConfig.StandardHistory
	ctx = internal.NewSchemaConfigContext(ctx, shd.schemaConfig)
	if ps := shd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, shd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	shd.mutation.done = true
	return affected, err
}

// StandardHistoryDeleteOne is the builder for deleting a single StandardHistory entity.
type StandardHistoryDeleteOne struct {
	shd *StandardHistoryDelete
}

// Where appends a list predicates to the StandardHistoryDelete builder.
func (shdo *StandardHistoryDeleteOne) Where(ps ...predicate.StandardHistory) *StandardHistoryDeleteOne {
	shdo.shd.mutation.Where(ps...)
	return shdo
}

// Exec executes the deletion query.
func (shdo *StandardHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := shdo.shd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{standardhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (shdo *StandardHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := shdo.Exec(ctx); err != nil {
		panic(err)
	}
}