// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/program"
	"github.com/theopenlane/core/internal/ent/generated/programmembership"
	"github.com/theopenlane/core/internal/ent/generated/user"
	"github.com/theopenlane/core/pkg/enums"
)

// ProgramMembershipCreate is the builder for creating a ProgramMembership entity.
type ProgramMembershipCreate struct {
	config
	mutation *ProgramMembershipMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pmc *ProgramMembershipCreate) SetCreatedAt(t time.Time) *ProgramMembershipCreate {
	pmc.mutation.SetCreatedAt(t)
	return pmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableCreatedAt(t *time.Time) *ProgramMembershipCreate {
	if t != nil {
		pmc.SetCreatedAt(*t)
	}
	return pmc
}

// SetUpdatedAt sets the "updated_at" field.
func (pmc *ProgramMembershipCreate) SetUpdatedAt(t time.Time) *ProgramMembershipCreate {
	pmc.mutation.SetUpdatedAt(t)
	return pmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableUpdatedAt(t *time.Time) *ProgramMembershipCreate {
	if t != nil {
		pmc.SetUpdatedAt(*t)
	}
	return pmc
}

// SetCreatedBy sets the "created_by" field.
func (pmc *ProgramMembershipCreate) SetCreatedBy(s string) *ProgramMembershipCreate {
	pmc.mutation.SetCreatedBy(s)
	return pmc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableCreatedBy(s *string) *ProgramMembershipCreate {
	if s != nil {
		pmc.SetCreatedBy(*s)
	}
	return pmc
}

// SetUpdatedBy sets the "updated_by" field.
func (pmc *ProgramMembershipCreate) SetUpdatedBy(s string) *ProgramMembershipCreate {
	pmc.mutation.SetUpdatedBy(s)
	return pmc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableUpdatedBy(s *string) *ProgramMembershipCreate {
	if s != nil {
		pmc.SetUpdatedBy(*s)
	}
	return pmc
}

// SetMappingID sets the "mapping_id" field.
func (pmc *ProgramMembershipCreate) SetMappingID(s string) *ProgramMembershipCreate {
	pmc.mutation.SetMappingID(s)
	return pmc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableMappingID(s *string) *ProgramMembershipCreate {
	if s != nil {
		pmc.SetMappingID(*s)
	}
	return pmc
}

// SetDeletedAt sets the "deleted_at" field.
func (pmc *ProgramMembershipCreate) SetDeletedAt(t time.Time) *ProgramMembershipCreate {
	pmc.mutation.SetDeletedAt(t)
	return pmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableDeletedAt(t *time.Time) *ProgramMembershipCreate {
	if t != nil {
		pmc.SetDeletedAt(*t)
	}
	return pmc
}

// SetDeletedBy sets the "deleted_by" field.
func (pmc *ProgramMembershipCreate) SetDeletedBy(s string) *ProgramMembershipCreate {
	pmc.mutation.SetDeletedBy(s)
	return pmc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableDeletedBy(s *string) *ProgramMembershipCreate {
	if s != nil {
		pmc.SetDeletedBy(*s)
	}
	return pmc
}

// SetRole sets the "role" field.
func (pmc *ProgramMembershipCreate) SetRole(e enums.Role) *ProgramMembershipCreate {
	pmc.mutation.SetRole(e)
	return pmc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableRole(e *enums.Role) *ProgramMembershipCreate {
	if e != nil {
		pmc.SetRole(*e)
	}
	return pmc
}

// SetProgramID sets the "program_id" field.
func (pmc *ProgramMembershipCreate) SetProgramID(s string) *ProgramMembershipCreate {
	pmc.mutation.SetProgramID(s)
	return pmc
}

// SetUserID sets the "user_id" field.
func (pmc *ProgramMembershipCreate) SetUserID(s string) *ProgramMembershipCreate {
	pmc.mutation.SetUserID(s)
	return pmc
}

// SetID sets the "id" field.
func (pmc *ProgramMembershipCreate) SetID(s string) *ProgramMembershipCreate {
	pmc.mutation.SetID(s)
	return pmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pmc *ProgramMembershipCreate) SetNillableID(s *string) *ProgramMembershipCreate {
	if s != nil {
		pmc.SetID(*s)
	}
	return pmc
}

// SetProgram sets the "program" edge to the Program entity.
func (pmc *ProgramMembershipCreate) SetProgram(p *Program) *ProgramMembershipCreate {
	return pmc.SetProgramID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pmc *ProgramMembershipCreate) SetUser(u *User) *ProgramMembershipCreate {
	return pmc.SetUserID(u.ID)
}

// Mutation returns the ProgramMembershipMutation object of the builder.
func (pmc *ProgramMembershipCreate) Mutation() *ProgramMembershipMutation {
	return pmc.mutation
}

// Save creates the ProgramMembership in the database.
func (pmc *ProgramMembershipCreate) Save(ctx context.Context) (*ProgramMembership, error) {
	if err := pmc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pmc.sqlSave, pmc.mutation, pmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pmc *ProgramMembershipCreate) SaveX(ctx context.Context) *ProgramMembership {
	v, err := pmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmc *ProgramMembershipCreate) Exec(ctx context.Context) error {
	_, err := pmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmc *ProgramMembershipCreate) ExecX(ctx context.Context) {
	if err := pmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmc *ProgramMembershipCreate) defaults() error {
	if _, ok := pmc.mutation.CreatedAt(); !ok {
		if programmembership.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized programmembership.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := programmembership.DefaultCreatedAt()
		pmc.mutation.SetCreatedAt(v)
	}
	if _, ok := pmc.mutation.UpdatedAt(); !ok {
		if programmembership.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized programmembership.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := programmembership.DefaultUpdatedAt()
		pmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pmc.mutation.MappingID(); !ok {
		if programmembership.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized programmembership.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := programmembership.DefaultMappingID()
		pmc.mutation.SetMappingID(v)
	}
	if _, ok := pmc.mutation.Role(); !ok {
		v := programmembership.DefaultRole
		pmc.mutation.SetRole(v)
	}
	if _, ok := pmc.mutation.ID(); !ok {
		if programmembership.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized programmembership.DefaultID (forgotten import generated/runtime?)")
		}
		v := programmembership.DefaultID()
		pmc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pmc *ProgramMembershipCreate) check() error {
	if _, ok := pmc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "ProgramMembership.mapping_id"`)}
	}
	if _, ok := pmc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`generated: missing required field "ProgramMembership.role"`)}
	}
	if v, ok := pmc.mutation.Role(); ok {
		if err := programmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "ProgramMembership.role": %w`, err)}
		}
	}
	if _, ok := pmc.mutation.ProgramID(); !ok {
		return &ValidationError{Name: "program_id", err: errors.New(`generated: missing required field "ProgramMembership.program_id"`)}
	}
	if _, ok := pmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "ProgramMembership.user_id"`)}
	}
	if len(pmc.mutation.ProgramIDs()) == 0 {
		return &ValidationError{Name: "program", err: errors.New(`generated: missing required edge "ProgramMembership.program"`)}
	}
	if len(pmc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`generated: missing required edge "ProgramMembership.user"`)}
	}
	return nil
}

func (pmc *ProgramMembershipCreate) sqlSave(ctx context.Context) (*ProgramMembership, error) {
	if err := pmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ProgramMembership.ID type: %T", _spec.ID.Value)
		}
	}
	pmc.mutation.id = &_node.ID
	pmc.mutation.done = true
	return _node, nil
}

func (pmc *ProgramMembershipCreate) createSpec() (*ProgramMembership, *sqlgraph.CreateSpec) {
	var (
		_node = &ProgramMembership{config: pmc.config}
		_spec = sqlgraph.NewCreateSpec(programmembership.Table, sqlgraph.NewFieldSpec(programmembership.FieldID, field.TypeString))
	)
	_spec.Schema = pmc.schemaConfig.ProgramMembership
	if id, ok := pmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pmc.mutation.CreatedAt(); ok {
		_spec.SetField(programmembership.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pmc.mutation.UpdatedAt(); ok {
		_spec.SetField(programmembership.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pmc.mutation.CreatedBy(); ok {
		_spec.SetField(programmembership.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pmc.mutation.UpdatedBy(); ok {
		_spec.SetField(programmembership.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pmc.mutation.MappingID(); ok {
		_spec.SetField(programmembership.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := pmc.mutation.DeletedAt(); ok {
		_spec.SetField(programmembership.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pmc.mutation.DeletedBy(); ok {
		_spec.SetField(programmembership.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := pmc.mutation.Role(); ok {
		_spec.SetField(programmembership.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := pmc.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programmembership.ProgramTable,
			Columns: []string{programmembership.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeString),
			},
		}
		edge.Schema = pmc.schemaConfig.ProgramMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProgramID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programmembership.UserTable,
			Columns: []string{programmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = pmc.schemaConfig.ProgramMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProgramMembershipCreateBulk is the builder for creating many ProgramMembership entities in bulk.
type ProgramMembershipCreateBulk struct {
	config
	err      error
	builders []*ProgramMembershipCreate
}

// Save creates the ProgramMembership entities in the database.
func (pmcb *ProgramMembershipCreateBulk) Save(ctx context.Context) ([]*ProgramMembership, error) {
	if pmcb.err != nil {
		return nil, pmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pmcb.builders))
	nodes := make([]*ProgramMembership, len(pmcb.builders))
	mutators := make([]Mutator, len(pmcb.builders))
	for i := range pmcb.builders {
		func(i int, root context.Context) {
			builder := pmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProgramMembershipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmcb *ProgramMembershipCreateBulk) SaveX(ctx context.Context) []*ProgramMembership {
	v, err := pmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmcb *ProgramMembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := pmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmcb *ProgramMembershipCreateBulk) ExecX(ctx context.Context) {
	if err := pmcb.Exec(ctx); err != nil {
		panic(err)
	}
}