// Code generated by ent, DO NOT EDIT.

package controlobjective

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the controlobjective type in the database.
	Label = "control_objective"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldControlObjectiveType holds the string denoting the control_objective_type field in the database.
	FieldControlObjectiveType = "control_objective_type"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldControlNumber holds the string denoting the control_number field in the database.
	FieldControlNumber = "control_number"
	// FieldFamily holds the string denoting the family field in the database.
	FieldFamily = "family"
	// FieldClass holds the string denoting the class field in the database.
	FieldClass = "class"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldMappedFrameworks holds the string denoting the mapped_frameworks field in the database.
	FieldMappedFrameworks = "mapped_frameworks"
	// FieldDetails holds the string denoting the details field in the database.
	FieldDetails = "details"
	// EdgePolicy holds the string denoting the policy edge name in mutations.
	EdgePolicy = "policy"
	// EdgeControls holds the string denoting the controls edge name in mutations.
	EdgeControls = "controls"
	// EdgeProcedures holds the string denoting the procedures edge name in mutations.
	EdgeProcedures = "procedures"
	// EdgeRisks holds the string denoting the risks edge name in mutations.
	EdgeRisks = "risks"
	// EdgeSubcontrols holds the string denoting the subcontrols edge name in mutations.
	EdgeSubcontrols = "subcontrols"
	// EdgeStandard holds the string denoting the standard edge name in mutations.
	EdgeStandard = "standard"
	// EdgeNarratives holds the string denoting the narratives edge name in mutations.
	EdgeNarratives = "narratives"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgePrograms holds the string denoting the programs edge name in mutations.
	EdgePrograms = "programs"
	// Table holds the table name of the controlobjective in the database.
	Table = "control_objectives"
	// PolicyTable is the table that holds the policy relation/edge. The primary key declared below.
	PolicyTable = "internal_policy_controlobjectives"
	// PolicyInverseTable is the table name for the InternalPolicy entity.
	// It exists in this package in order to avoid circular dependency with the "internalpolicy" package.
	PolicyInverseTable = "internal_policies"
	// ControlsTable is the table that holds the controls relation/edge.
	ControlsTable = "controls"
	// ControlsInverseTable is the table name for the Control entity.
	// It exists in this package in order to avoid circular dependency with the "control" package.
	ControlsInverseTable = "controls"
	// ControlsColumn is the table column denoting the controls relation/edge.
	ControlsColumn = "control_objective_controls"
	// ProceduresTable is the table that holds the procedures relation/edge.
	ProceduresTable = "procedures"
	// ProceduresInverseTable is the table name for the Procedure entity.
	// It exists in this package in order to avoid circular dependency with the "procedure" package.
	ProceduresInverseTable = "procedures"
	// ProceduresColumn is the table column denoting the procedures relation/edge.
	ProceduresColumn = "control_objective_procedures"
	// RisksTable is the table that holds the risks relation/edge.
	RisksTable = "risks"
	// RisksInverseTable is the table name for the Risk entity.
	// It exists in this package in order to avoid circular dependency with the "risk" package.
	RisksInverseTable = "risks"
	// RisksColumn is the table column denoting the risks relation/edge.
	RisksColumn = "control_objective_risks"
	// SubcontrolsTable is the table that holds the subcontrols relation/edge.
	SubcontrolsTable = "subcontrols"
	// SubcontrolsInverseTable is the table name for the Subcontrol entity.
	// It exists in this package in order to avoid circular dependency with the "subcontrol" package.
	SubcontrolsInverseTable = "subcontrols"
	// SubcontrolsColumn is the table column denoting the subcontrols relation/edge.
	SubcontrolsColumn = "control_objective_subcontrols"
	// StandardTable is the table that holds the standard relation/edge. The primary key declared below.
	StandardTable = "standard_controlobjectives"
	// StandardInverseTable is the table name for the Standard entity.
	// It exists in this package in order to avoid circular dependency with the "standard" package.
	StandardInverseTable = "standards"
	// NarrativesTable is the table that holds the narratives relation/edge. The primary key declared below.
	NarrativesTable = "control_objective_narratives"
	// NarrativesInverseTable is the table name for the Narrative entity.
	// It exists in this package in order to avoid circular dependency with the "narrative" package.
	NarrativesInverseTable = "narratives"
	// TasksTable is the table that holds the tasks relation/edge. The primary key declared below.
	TasksTable = "control_objective_tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// ProgramsTable is the table that holds the programs relation/edge. The primary key declared below.
	ProgramsTable = "program_controlobjectives"
	// ProgramsInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramsInverseTable = "programs"
)

// Columns holds all SQL columns for controlobjective fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldMappingID,
	FieldTags,
	FieldName,
	FieldDescription,
	FieldStatus,
	FieldControlObjectiveType,
	FieldVersion,
	FieldControlNumber,
	FieldFamily,
	FieldClass,
	FieldSource,
	FieldMappedFrameworks,
	FieldDetails,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "control_objectives"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"control_controlobjectives",
}

var (
	// PolicyPrimaryKey and PolicyColumn2 are the table columns denoting the
	// primary key for the policy relation (M2M).
	PolicyPrimaryKey = []string{"internal_policy_id", "control_objective_id"}
	// StandardPrimaryKey and StandardColumn2 are the table columns denoting the
	// primary key for the standard relation (M2M).
	StandardPrimaryKey = []string{"standard_id", "control_objective_id"}
	// NarrativesPrimaryKey and NarrativesColumn2 are the table columns denoting the
	// primary key for the narratives relation (M2M).
	NarrativesPrimaryKey = []string{"control_objective_id", "narrative_id"}
	// TasksPrimaryKey and TasksColumn2 are the table columns denoting the
	// primary key for the tasks relation (M2M).
	TasksPrimaryKey = []string{"control_objective_id", "task_id"}
	// ProgramsPrimaryKey and ProgramsColumn2 are the table columns denoting the
	// primary key for the programs relation (M2M).
	ProgramsPrimaryKey = []string{"program_id", "control_objective_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/theopenlane/core/internal/ent/generated/runtime"
var (
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the ControlObjective queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByMappingID orders the results by the mapping_id field.
func ByMappingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByControlObjectiveType orders the results by the control_objective_type field.
func ByControlObjectiveType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldControlObjectiveType, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByControlNumber orders the results by the control_number field.
func ByControlNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldControlNumber, opts...).ToFunc()
}

// ByFamily orders the results by the family field.
func ByFamily(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFamily, opts...).ToFunc()
}

// ByClass orders the results by the class field.
func ByClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClass, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByMappedFrameworks orders the results by the mapped_frameworks field.
func ByMappedFrameworks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappedFrameworks, opts...).ToFunc()
}

// ByPolicyCount orders the results by policy count.
func ByPolicyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPolicyStep(), opts...)
	}
}

// ByPolicy orders the results by policy terms.
func ByPolicy(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPolicyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByControlsCount orders the results by controls count.
func ByControlsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newControlsStep(), opts...)
	}
}

// ByControls orders the results by controls terms.
func ByControls(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newControlsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProceduresCount orders the results by procedures count.
func ByProceduresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProceduresStep(), opts...)
	}
}

// ByProcedures orders the results by procedures terms.
func ByProcedures(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProceduresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRisksCount orders the results by risks count.
func ByRisksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRisksStep(), opts...)
	}
}

// ByRisks orders the results by risks terms.
func ByRisks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRisksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubcontrolsCount orders the results by subcontrols count.
func BySubcontrolsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubcontrolsStep(), opts...)
	}
}

// BySubcontrols orders the results by subcontrols terms.
func BySubcontrols(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubcontrolsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStandardCount orders the results by standard count.
func ByStandardCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStandardStep(), opts...)
	}
}

// ByStandard orders the results by standard terms.
func ByStandard(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStandardStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNarrativesCount orders the results by narratives count.
func ByNarrativesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNarrativesStep(), opts...)
	}
}

// ByNarratives orders the results by narratives terms.
func ByNarratives(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNarrativesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProgramsCount orders the results by programs count.
func ByProgramsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramsStep(), opts...)
	}
}

// ByPrograms orders the results by programs terms.
func ByPrograms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPolicyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PolicyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PolicyTable, PolicyPrimaryKey...),
	)
}
func newControlsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ControlsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ControlsTable, ControlsColumn),
	)
}
func newProceduresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProceduresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProceduresTable, ProceduresColumn),
	)
}
func newRisksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RisksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RisksTable, RisksColumn),
	)
}
func newSubcontrolsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubcontrolsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubcontrolsTable, SubcontrolsColumn),
	)
}
func newStandardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StandardInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, StandardTable, StandardPrimaryKey...),
	)
}
func newNarrativesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NarrativesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NarrativesTable, NarrativesPrimaryKey...),
	)
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TasksTable, TasksPrimaryKey...),
	)
}
func newProgramsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProgramsTable, ProgramsPrimaryKey...),
	)
}
