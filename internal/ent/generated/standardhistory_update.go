// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/internal/ent/generated/standardhistory"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// StandardHistoryUpdate is the builder for updating StandardHistory entities.
type StandardHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *StandardHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the StandardHistoryUpdate builder.
func (shu *StandardHistoryUpdate) Where(ps ...predicate.StandardHistory) *StandardHistoryUpdate {
	shu.mutation.Where(ps...)
	return shu
}

// SetUpdatedAt sets the "updated_at" field.
func (shu *StandardHistoryUpdate) SetUpdatedAt(t time.Time) *StandardHistoryUpdate {
	shu.mutation.SetUpdatedAt(t)
	return shu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (shu *StandardHistoryUpdate) ClearUpdatedAt() *StandardHistoryUpdate {
	shu.mutation.ClearUpdatedAt()
	return shu
}

// SetUpdatedBy sets the "updated_by" field.
func (shu *StandardHistoryUpdate) SetUpdatedBy(s string) *StandardHistoryUpdate {
	shu.mutation.SetUpdatedBy(s)
	return shu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableUpdatedBy(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetUpdatedBy(*s)
	}
	return shu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (shu *StandardHistoryUpdate) ClearUpdatedBy() *StandardHistoryUpdate {
	shu.mutation.ClearUpdatedBy()
	return shu
}

// SetDeletedAt sets the "deleted_at" field.
func (shu *StandardHistoryUpdate) SetDeletedAt(t time.Time) *StandardHistoryUpdate {
	shu.mutation.SetDeletedAt(t)
	return shu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableDeletedAt(t *time.Time) *StandardHistoryUpdate {
	if t != nil {
		shu.SetDeletedAt(*t)
	}
	return shu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (shu *StandardHistoryUpdate) ClearDeletedAt() *StandardHistoryUpdate {
	shu.mutation.ClearDeletedAt()
	return shu
}

// SetDeletedBy sets the "deleted_by" field.
func (shu *StandardHistoryUpdate) SetDeletedBy(s string) *StandardHistoryUpdate {
	shu.mutation.SetDeletedBy(s)
	return shu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableDeletedBy(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetDeletedBy(*s)
	}
	return shu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (shu *StandardHistoryUpdate) ClearDeletedBy() *StandardHistoryUpdate {
	shu.mutation.ClearDeletedBy()
	return shu
}

// SetTags sets the "tags" field.
func (shu *StandardHistoryUpdate) SetTags(s []string) *StandardHistoryUpdate {
	shu.mutation.SetTags(s)
	return shu
}

// AppendTags appends s to the "tags" field.
func (shu *StandardHistoryUpdate) AppendTags(s []string) *StandardHistoryUpdate {
	shu.mutation.AppendTags(s)
	return shu
}

// ClearTags clears the value of the "tags" field.
func (shu *StandardHistoryUpdate) ClearTags() *StandardHistoryUpdate {
	shu.mutation.ClearTags()
	return shu
}

// SetName sets the "name" field.
func (shu *StandardHistoryUpdate) SetName(s string) *StandardHistoryUpdate {
	shu.mutation.SetName(s)
	return shu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableName(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetName(*s)
	}
	return shu
}

// SetDescription sets the "description" field.
func (shu *StandardHistoryUpdate) SetDescription(s string) *StandardHistoryUpdate {
	shu.mutation.SetDescription(s)
	return shu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableDescription(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetDescription(*s)
	}
	return shu
}

// ClearDescription clears the value of the "description" field.
func (shu *StandardHistoryUpdate) ClearDescription() *StandardHistoryUpdate {
	shu.mutation.ClearDescription()
	return shu
}

// SetFamily sets the "family" field.
func (shu *StandardHistoryUpdate) SetFamily(s string) *StandardHistoryUpdate {
	shu.mutation.SetFamily(s)
	return shu
}

// SetNillableFamily sets the "family" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableFamily(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetFamily(*s)
	}
	return shu
}

// ClearFamily clears the value of the "family" field.
func (shu *StandardHistoryUpdate) ClearFamily() *StandardHistoryUpdate {
	shu.mutation.ClearFamily()
	return shu
}

// SetStatus sets the "status" field.
func (shu *StandardHistoryUpdate) SetStatus(s string) *StandardHistoryUpdate {
	shu.mutation.SetStatus(s)
	return shu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableStatus(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetStatus(*s)
	}
	return shu
}

// ClearStatus clears the value of the "status" field.
func (shu *StandardHistoryUpdate) ClearStatus() *StandardHistoryUpdate {
	shu.mutation.ClearStatus()
	return shu
}

// SetStandardType sets the "standard_type" field.
func (shu *StandardHistoryUpdate) SetStandardType(s string) *StandardHistoryUpdate {
	shu.mutation.SetStandardType(s)
	return shu
}

// SetNillableStandardType sets the "standard_type" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableStandardType(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetStandardType(*s)
	}
	return shu
}

// ClearStandardType clears the value of the "standard_type" field.
func (shu *StandardHistoryUpdate) ClearStandardType() *StandardHistoryUpdate {
	shu.mutation.ClearStandardType()
	return shu
}

// SetVersion sets the "version" field.
func (shu *StandardHistoryUpdate) SetVersion(s string) *StandardHistoryUpdate {
	shu.mutation.SetVersion(s)
	return shu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableVersion(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetVersion(*s)
	}
	return shu
}

// ClearVersion clears the value of the "version" field.
func (shu *StandardHistoryUpdate) ClearVersion() *StandardHistoryUpdate {
	shu.mutation.ClearVersion()
	return shu
}

// SetPurposeAndScope sets the "purpose_and_scope" field.
func (shu *StandardHistoryUpdate) SetPurposeAndScope(s string) *StandardHistoryUpdate {
	shu.mutation.SetPurposeAndScope(s)
	return shu
}

// SetNillablePurposeAndScope sets the "purpose_and_scope" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillablePurposeAndScope(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetPurposeAndScope(*s)
	}
	return shu
}

// ClearPurposeAndScope clears the value of the "purpose_and_scope" field.
func (shu *StandardHistoryUpdate) ClearPurposeAndScope() *StandardHistoryUpdate {
	shu.mutation.ClearPurposeAndScope()
	return shu
}

// SetBackground sets the "background" field.
func (shu *StandardHistoryUpdate) SetBackground(s string) *StandardHistoryUpdate {
	shu.mutation.SetBackground(s)
	return shu
}

// SetNillableBackground sets the "background" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableBackground(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetBackground(*s)
	}
	return shu
}

// ClearBackground clears the value of the "background" field.
func (shu *StandardHistoryUpdate) ClearBackground() *StandardHistoryUpdate {
	shu.mutation.ClearBackground()
	return shu
}

// SetSatisfies sets the "satisfies" field.
func (shu *StandardHistoryUpdate) SetSatisfies(s string) *StandardHistoryUpdate {
	shu.mutation.SetSatisfies(s)
	return shu
}

// SetNillableSatisfies sets the "satisfies" field if the given value is not nil.
func (shu *StandardHistoryUpdate) SetNillableSatisfies(s *string) *StandardHistoryUpdate {
	if s != nil {
		shu.SetSatisfies(*s)
	}
	return shu
}

// ClearSatisfies clears the value of the "satisfies" field.
func (shu *StandardHistoryUpdate) ClearSatisfies() *StandardHistoryUpdate {
	shu.mutation.ClearSatisfies()
	return shu
}

// SetDetails sets the "details" field.
func (shu *StandardHistoryUpdate) SetDetails(m map[string]interface{}) *StandardHistoryUpdate {
	shu.mutation.SetDetails(m)
	return shu
}

// ClearDetails clears the value of the "details" field.
func (shu *StandardHistoryUpdate) ClearDetails() *StandardHistoryUpdate {
	shu.mutation.ClearDetails()
	return shu
}

// Mutation returns the StandardHistoryMutation object of the builder.
func (shu *StandardHistoryUpdate) Mutation() *StandardHistoryMutation {
	return shu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (shu *StandardHistoryUpdate) Save(ctx context.Context) (int, error) {
	shu.defaults()
	return withHooks(ctx, shu.sqlSave, shu.mutation, shu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shu *StandardHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := shu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (shu *StandardHistoryUpdate) Exec(ctx context.Context) error {
	_, err := shu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shu *StandardHistoryUpdate) ExecX(ctx context.Context) {
	if err := shu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shu *StandardHistoryUpdate) defaults() {
	if _, ok := shu.mutation.UpdatedAt(); !ok && !shu.mutation.UpdatedAtCleared() {
		v := standardhistory.UpdateDefaultUpdatedAt()
		shu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (shu *StandardHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *StandardHistoryUpdate {
	shu.modifiers = append(shu.modifiers, modifiers...)
	return shu
}

func (shu *StandardHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(standardhistory.Table, standardhistory.Columns, sqlgraph.NewFieldSpec(standardhistory.FieldID, field.TypeString))
	if ps := shu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if shu.mutation.RefCleared() {
		_spec.ClearField(standardhistory.FieldRef, field.TypeString)
	}
	if shu.mutation.CreatedAtCleared() {
		_spec.ClearField(standardhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := shu.mutation.UpdatedAt(); ok {
		_spec.SetField(standardhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if shu.mutation.UpdatedAtCleared() {
		_spec.ClearField(standardhistory.FieldUpdatedAt, field.TypeTime)
	}
	if shu.mutation.CreatedByCleared() {
		_spec.ClearField(standardhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := shu.mutation.UpdatedBy(); ok {
		_spec.SetField(standardhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if shu.mutation.UpdatedByCleared() {
		_spec.ClearField(standardhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := shu.mutation.DeletedAt(); ok {
		_spec.SetField(standardhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if shu.mutation.DeletedAtCleared() {
		_spec.ClearField(standardhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := shu.mutation.DeletedBy(); ok {
		_spec.SetField(standardhistory.FieldDeletedBy, field.TypeString, value)
	}
	if shu.mutation.DeletedByCleared() {
		_spec.ClearField(standardhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := shu.mutation.Tags(); ok {
		_spec.SetField(standardhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := shu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, standardhistory.FieldTags, value)
		})
	}
	if shu.mutation.TagsCleared() {
		_spec.ClearField(standardhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := shu.mutation.Name(); ok {
		_spec.SetField(standardhistory.FieldName, field.TypeString, value)
	}
	if value, ok := shu.mutation.Description(); ok {
		_spec.SetField(standardhistory.FieldDescription, field.TypeString, value)
	}
	if shu.mutation.DescriptionCleared() {
		_spec.ClearField(standardhistory.FieldDescription, field.TypeString)
	}
	if value, ok := shu.mutation.Family(); ok {
		_spec.SetField(standardhistory.FieldFamily, field.TypeString, value)
	}
	if shu.mutation.FamilyCleared() {
		_spec.ClearField(standardhistory.FieldFamily, field.TypeString)
	}
	if value, ok := shu.mutation.Status(); ok {
		_spec.SetField(standardhistory.FieldStatus, field.TypeString, value)
	}
	if shu.mutation.StatusCleared() {
		_spec.ClearField(standardhistory.FieldStatus, field.TypeString)
	}
	if value, ok := shu.mutation.StandardType(); ok {
		_spec.SetField(standardhistory.FieldStandardType, field.TypeString, value)
	}
	if shu.mutation.StandardTypeCleared() {
		_spec.ClearField(standardhistory.FieldStandardType, field.TypeString)
	}
	if value, ok := shu.mutation.Version(); ok {
		_spec.SetField(standardhistory.FieldVersion, field.TypeString, value)
	}
	if shu.mutation.VersionCleared() {
		_spec.ClearField(standardhistory.FieldVersion, field.TypeString)
	}
	if value, ok := shu.mutation.PurposeAndScope(); ok {
		_spec.SetField(standardhistory.FieldPurposeAndScope, field.TypeString, value)
	}
	if shu.mutation.PurposeAndScopeCleared() {
		_spec.ClearField(standardhistory.FieldPurposeAndScope, field.TypeString)
	}
	if value, ok := shu.mutation.Background(); ok {
		_spec.SetField(standardhistory.FieldBackground, field.TypeString, value)
	}
	if shu.mutation.BackgroundCleared() {
		_spec.ClearField(standardhistory.FieldBackground, field.TypeString)
	}
	if value, ok := shu.mutation.Satisfies(); ok {
		_spec.SetField(standardhistory.FieldSatisfies, field.TypeString, value)
	}
	if shu.mutation.SatisfiesCleared() {
		_spec.ClearField(standardhistory.FieldSatisfies, field.TypeString)
	}
	if value, ok := shu.mutation.Details(); ok {
		_spec.SetField(standardhistory.FieldDetails, field.TypeJSON, value)
	}
	if shu.mutation.DetailsCleared() {
		_spec.ClearField(standardhistory.FieldDetails, field.TypeJSON)
	}
	_spec.Node.Schema = shu.schemaConfig.StandardHistory
	ctx = internal.NewSchemaConfigContext(ctx, shu.schemaConfig)
	_spec.AddModifiers(shu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, shu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{standardhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	shu.mutation.done = true
	return n, nil
}

// StandardHistoryUpdateOne is the builder for updating a single StandardHistory entity.
type StandardHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *StandardHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (shuo *StandardHistoryUpdateOne) SetUpdatedAt(t time.Time) *StandardHistoryUpdateOne {
	shuo.mutation.SetUpdatedAt(t)
	return shuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (shuo *StandardHistoryUpdateOne) ClearUpdatedAt() *StandardHistoryUpdateOne {
	shuo.mutation.ClearUpdatedAt()
	return shuo
}

// SetUpdatedBy sets the "updated_by" field.
func (shuo *StandardHistoryUpdateOne) SetUpdatedBy(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetUpdatedBy(s)
	return shuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableUpdatedBy(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetUpdatedBy(*s)
	}
	return shuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (shuo *StandardHistoryUpdateOne) ClearUpdatedBy() *StandardHistoryUpdateOne {
	shuo.mutation.ClearUpdatedBy()
	return shuo
}

// SetDeletedAt sets the "deleted_at" field.
func (shuo *StandardHistoryUpdateOne) SetDeletedAt(t time.Time) *StandardHistoryUpdateOne {
	shuo.mutation.SetDeletedAt(t)
	return shuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *StandardHistoryUpdateOne {
	if t != nil {
		shuo.SetDeletedAt(*t)
	}
	return shuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (shuo *StandardHistoryUpdateOne) ClearDeletedAt() *StandardHistoryUpdateOne {
	shuo.mutation.ClearDeletedAt()
	return shuo
}

// SetDeletedBy sets the "deleted_by" field.
func (shuo *StandardHistoryUpdateOne) SetDeletedBy(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetDeletedBy(s)
	return shuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableDeletedBy(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetDeletedBy(*s)
	}
	return shuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (shuo *StandardHistoryUpdateOne) ClearDeletedBy() *StandardHistoryUpdateOne {
	shuo.mutation.ClearDeletedBy()
	return shuo
}

// SetTags sets the "tags" field.
func (shuo *StandardHistoryUpdateOne) SetTags(s []string) *StandardHistoryUpdateOne {
	shuo.mutation.SetTags(s)
	return shuo
}

// AppendTags appends s to the "tags" field.
func (shuo *StandardHistoryUpdateOne) AppendTags(s []string) *StandardHistoryUpdateOne {
	shuo.mutation.AppendTags(s)
	return shuo
}

// ClearTags clears the value of the "tags" field.
func (shuo *StandardHistoryUpdateOne) ClearTags() *StandardHistoryUpdateOne {
	shuo.mutation.ClearTags()
	return shuo
}

// SetName sets the "name" field.
func (shuo *StandardHistoryUpdateOne) SetName(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetName(s)
	return shuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableName(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetName(*s)
	}
	return shuo
}

// SetDescription sets the "description" field.
func (shuo *StandardHistoryUpdateOne) SetDescription(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetDescription(s)
	return shuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableDescription(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetDescription(*s)
	}
	return shuo
}

// ClearDescription clears the value of the "description" field.
func (shuo *StandardHistoryUpdateOne) ClearDescription() *StandardHistoryUpdateOne {
	shuo.mutation.ClearDescription()
	return shuo
}

// SetFamily sets the "family" field.
func (shuo *StandardHistoryUpdateOne) SetFamily(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetFamily(s)
	return shuo
}

// SetNillableFamily sets the "family" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableFamily(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetFamily(*s)
	}
	return shuo
}

// ClearFamily clears the value of the "family" field.
func (shuo *StandardHistoryUpdateOne) ClearFamily() *StandardHistoryUpdateOne {
	shuo.mutation.ClearFamily()
	return shuo
}

// SetStatus sets the "status" field.
func (shuo *StandardHistoryUpdateOne) SetStatus(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetStatus(s)
	return shuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableStatus(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetStatus(*s)
	}
	return shuo
}

// ClearStatus clears the value of the "status" field.
func (shuo *StandardHistoryUpdateOne) ClearStatus() *StandardHistoryUpdateOne {
	shuo.mutation.ClearStatus()
	return shuo
}

// SetStandardType sets the "standard_type" field.
func (shuo *StandardHistoryUpdateOne) SetStandardType(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetStandardType(s)
	return shuo
}

// SetNillableStandardType sets the "standard_type" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableStandardType(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetStandardType(*s)
	}
	return shuo
}

// ClearStandardType clears the value of the "standard_type" field.
func (shuo *StandardHistoryUpdateOne) ClearStandardType() *StandardHistoryUpdateOne {
	shuo.mutation.ClearStandardType()
	return shuo
}

// SetVersion sets the "version" field.
func (shuo *StandardHistoryUpdateOne) SetVersion(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetVersion(s)
	return shuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableVersion(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetVersion(*s)
	}
	return shuo
}

// ClearVersion clears the value of the "version" field.
func (shuo *StandardHistoryUpdateOne) ClearVersion() *StandardHistoryUpdateOne {
	shuo.mutation.ClearVersion()
	return shuo
}

// SetPurposeAndScope sets the "purpose_and_scope" field.
func (shuo *StandardHistoryUpdateOne) SetPurposeAndScope(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetPurposeAndScope(s)
	return shuo
}

// SetNillablePurposeAndScope sets the "purpose_and_scope" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillablePurposeAndScope(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetPurposeAndScope(*s)
	}
	return shuo
}

// ClearPurposeAndScope clears the value of the "purpose_and_scope" field.
func (shuo *StandardHistoryUpdateOne) ClearPurposeAndScope() *StandardHistoryUpdateOne {
	shuo.mutation.ClearPurposeAndScope()
	return shuo
}

// SetBackground sets the "background" field.
func (shuo *StandardHistoryUpdateOne) SetBackground(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetBackground(s)
	return shuo
}

// SetNillableBackground sets the "background" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableBackground(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetBackground(*s)
	}
	return shuo
}

// ClearBackground clears the value of the "background" field.
func (shuo *StandardHistoryUpdateOne) ClearBackground() *StandardHistoryUpdateOne {
	shuo.mutation.ClearBackground()
	return shuo
}

// SetSatisfies sets the "satisfies" field.
func (shuo *StandardHistoryUpdateOne) SetSatisfies(s string) *StandardHistoryUpdateOne {
	shuo.mutation.SetSatisfies(s)
	return shuo
}

// SetNillableSatisfies sets the "satisfies" field if the given value is not nil.
func (shuo *StandardHistoryUpdateOne) SetNillableSatisfies(s *string) *StandardHistoryUpdateOne {
	if s != nil {
		shuo.SetSatisfies(*s)
	}
	return shuo
}

// ClearSatisfies clears the value of the "satisfies" field.
func (shuo *StandardHistoryUpdateOne) ClearSatisfies() *StandardHistoryUpdateOne {
	shuo.mutation.ClearSatisfies()
	return shuo
}

// SetDetails sets the "details" field.
func (shuo *StandardHistoryUpdateOne) SetDetails(m map[string]interface{}) *StandardHistoryUpdateOne {
	shuo.mutation.SetDetails(m)
	return shuo
}

// ClearDetails clears the value of the "details" field.
func (shuo *StandardHistoryUpdateOne) ClearDetails() *StandardHistoryUpdateOne {
	shuo.mutation.ClearDetails()
	return shuo
}

// Mutation returns the StandardHistoryMutation object of the builder.
func (shuo *StandardHistoryUpdateOne) Mutation() *StandardHistoryMutation {
	return shuo.mutation
}

// Where appends a list predicates to the StandardHistoryUpdate builder.
func (shuo *StandardHistoryUpdateOne) Where(ps ...predicate.StandardHistory) *StandardHistoryUpdateOne {
	shuo.mutation.Where(ps...)
	return shuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (shuo *StandardHistoryUpdateOne) Select(field string, fields ...string) *StandardHistoryUpdateOne {
	shuo.fields = append([]string{field}, fields...)
	return shuo
}

// Save executes the query and returns the updated StandardHistory entity.
func (shuo *StandardHistoryUpdateOne) Save(ctx context.Context) (*StandardHistory, error) {
	shuo.defaults()
	return withHooks(ctx, shuo.sqlSave, shuo.mutation, shuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shuo *StandardHistoryUpdateOne) SaveX(ctx context.Context) *StandardHistory {
	node, err := shuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (shuo *StandardHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := shuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shuo *StandardHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := shuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shuo *StandardHistoryUpdateOne) defaults() {
	if _, ok := shuo.mutation.UpdatedAt(); !ok && !shuo.mutation.UpdatedAtCleared() {
		v := standardhistory.UpdateDefaultUpdatedAt()
		shuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (shuo *StandardHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *StandardHistoryUpdateOne {
	shuo.modifiers = append(shuo.modifiers, modifiers...)
	return shuo
}

func (shuo *StandardHistoryUpdateOne) sqlSave(ctx context.Context) (_node *StandardHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(standardhistory.Table, standardhistory.Columns, sqlgraph.NewFieldSpec(standardhistory.FieldID, field.TypeString))
	id, ok := shuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "StandardHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := shuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, standardhistory.FieldID)
		for _, f := range fields {
			if !standardhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != standardhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := shuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if shuo.mutation.RefCleared() {
		_spec.ClearField(standardhistory.FieldRef, field.TypeString)
	}
	if shuo.mutation.CreatedAtCleared() {
		_spec.ClearField(standardhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := shuo.mutation.UpdatedAt(); ok {
		_spec.SetField(standardhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if shuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(standardhistory.FieldUpdatedAt, field.TypeTime)
	}
	if shuo.mutation.CreatedByCleared() {
		_spec.ClearField(standardhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := shuo.mutation.UpdatedBy(); ok {
		_spec.SetField(standardhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if shuo.mutation.UpdatedByCleared() {
		_spec.ClearField(standardhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := shuo.mutation.DeletedAt(); ok {
		_spec.SetField(standardhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if shuo.mutation.DeletedAtCleared() {
		_spec.ClearField(standardhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := shuo.mutation.DeletedBy(); ok {
		_spec.SetField(standardhistory.FieldDeletedBy, field.TypeString, value)
	}
	if shuo.mutation.DeletedByCleared() {
		_spec.ClearField(standardhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := shuo.mutation.Tags(); ok {
		_spec.SetField(standardhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := shuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, standardhistory.FieldTags, value)
		})
	}
	if shuo.mutation.TagsCleared() {
		_spec.ClearField(standardhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := shuo.mutation.Name(); ok {
		_spec.SetField(standardhistory.FieldName, field.TypeString, value)
	}
	if value, ok := shuo.mutation.Description(); ok {
		_spec.SetField(standardhistory.FieldDescription, field.TypeString, value)
	}
	if shuo.mutation.DescriptionCleared() {
		_spec.ClearField(standardhistory.FieldDescription, field.TypeString)
	}
	if value, ok := shuo.mutation.Family(); ok {
		_spec.SetField(standardhistory.FieldFamily, field.TypeString, value)
	}
	if shuo.mutation.FamilyCleared() {
		_spec.ClearField(standardhistory.FieldFamily, field.TypeString)
	}
	if value, ok := shuo.mutation.Status(); ok {
		_spec.SetField(standardhistory.FieldStatus, field.TypeString, value)
	}
	if shuo.mutation.StatusCleared() {
		_spec.ClearField(standardhistory.FieldStatus, field.TypeString)
	}
	if value, ok := shuo.mutation.StandardType(); ok {
		_spec.SetField(standardhistory.FieldStandardType, field.TypeString, value)
	}
	if shuo.mutation.StandardTypeCleared() {
		_spec.ClearField(standardhistory.FieldStandardType, field.TypeString)
	}
	if value, ok := shuo.mutation.Version(); ok {
		_spec.SetField(standardhistory.FieldVersion, field.TypeString, value)
	}
	if shuo.mutation.VersionCleared() {
		_spec.ClearField(standardhistory.FieldVersion, field.TypeString)
	}
	if value, ok := shuo.mutation.PurposeAndScope(); ok {
		_spec.SetField(standardhistory.FieldPurposeAndScope, field.TypeString, value)
	}
	if shuo.mutation.PurposeAndScopeCleared() {
		_spec.ClearField(standardhistory.FieldPurposeAndScope, field.TypeString)
	}
	if value, ok := shuo.mutation.Background(); ok {
		_spec.SetField(standardhistory.FieldBackground, field.TypeString, value)
	}
	if shuo.mutation.BackgroundCleared() {
		_spec.ClearField(standardhistory.FieldBackground, field.TypeString)
	}
	if value, ok := shuo.mutation.Satisfies(); ok {
		_spec.SetField(standardhistory.FieldSatisfies, field.TypeString, value)
	}
	if shuo.mutation.SatisfiesCleared() {
		_spec.ClearField(standardhistory.FieldSatisfies, field.TypeString)
	}
	if value, ok := shuo.mutation.Details(); ok {
		_spec.SetField(standardhistory.FieldDetails, field.TypeJSON, value)
	}
	if shuo.mutation.DetailsCleared() {
		_spec.ClearField(standardhistory.FieldDetails, field.TypeJSON)
	}
	_spec.Node.Schema = shuo.schemaConfig.StandardHistory
	ctx = internal.NewSchemaConfigContext(ctx, shuo.schemaConfig)
	_spec.AddModifiers(shuo.modifiers...)
	_node = &StandardHistory{config: shuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, shuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{standardhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	shuo.mutation.done = true
	return _node, nil
}