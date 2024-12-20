// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/controlobjective"
)

// ControlObjective is the model entity for the ControlObjective schema.
type ControlObjective struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// the name of the control objective
	Name string `json:"name,omitempty"`
	// description of the control objective
	Description string `json:"description,omitempty"`
	// status of the control objective
	Status string `json:"status,omitempty"`
	// type of the control objective
	ControlObjectiveType string `json:"control_objective_type,omitempty"`
	// version of the control objective
	Version string `json:"version,omitempty"`
	// number of the control objective
	ControlNumber string `json:"control_number,omitempty"`
	// family of the control objective
	Family string `json:"family,omitempty"`
	// class associated with the control objective
	Class string `json:"class,omitempty"`
	// source of the control objective, e.g. framework, template, user-defined, etc.
	Source string `json:"source,omitempty"`
	// mapped frameworks
	MappedFrameworks string `json:"mapped_frameworks,omitempty"`
	// json data including details of the control objective
	Details map[string]interface{} `json:"details,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ControlObjectiveQuery when eager-loading is set.
	Edges                     ControlObjectiveEdges `json:"edges"`
	control_controlobjectives *string
	selectValues              sql.SelectValues
}

// ControlObjectiveEdges holds the relations/edges for other nodes in the graph.
type ControlObjectiveEdges struct {
	// Policy holds the value of the policy edge.
	Policy []*InternalPolicy `json:"policy,omitempty"`
	// Controls holds the value of the controls edge.
	Controls []*Control `json:"controls,omitempty"`
	// Procedures holds the value of the procedures edge.
	Procedures []*Procedure `json:"procedures,omitempty"`
	// Risks holds the value of the risks edge.
	Risks []*Risk `json:"risks,omitempty"`
	// Subcontrols holds the value of the subcontrols edge.
	Subcontrols []*Subcontrol `json:"subcontrols,omitempty"`
	// Standard holds the value of the standard edge.
	Standard []*Standard `json:"standard,omitempty"`
	// Narratives holds the value of the narratives edge.
	Narratives []*Narrative `json:"narratives,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Programs holds the value of the programs edge.
	Programs []*Program `json:"programs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
	// totalCount holds the count of the edges above.
	totalCount [9]map[string]int

	namedPolicy      map[string][]*InternalPolicy
	namedControls    map[string][]*Control
	namedProcedures  map[string][]*Procedure
	namedRisks       map[string][]*Risk
	namedSubcontrols map[string][]*Subcontrol
	namedStandard    map[string][]*Standard
	namedNarratives  map[string][]*Narrative
	namedTasks       map[string][]*Task
	namedPrograms    map[string][]*Program
}

// PolicyOrErr returns the Policy value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) PolicyOrErr() ([]*InternalPolicy, error) {
	if e.loadedTypes[0] {
		return e.Policy, nil
	}
	return nil, &NotLoadedError{edge: "policy"}
}

// ControlsOrErr returns the Controls value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ControlsOrErr() ([]*Control, error) {
	if e.loadedTypes[1] {
		return e.Controls, nil
	}
	return nil, &NotLoadedError{edge: "controls"}
}

// ProceduresOrErr returns the Procedures value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ProceduresOrErr() ([]*Procedure, error) {
	if e.loadedTypes[2] {
		return e.Procedures, nil
	}
	return nil, &NotLoadedError{edge: "procedures"}
}

// RisksOrErr returns the Risks value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) RisksOrErr() ([]*Risk, error) {
	if e.loadedTypes[3] {
		return e.Risks, nil
	}
	return nil, &NotLoadedError{edge: "risks"}
}

// SubcontrolsOrErr returns the Subcontrols value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) SubcontrolsOrErr() ([]*Subcontrol, error) {
	if e.loadedTypes[4] {
		return e.Subcontrols, nil
	}
	return nil, &NotLoadedError{edge: "subcontrols"}
}

// StandardOrErr returns the Standard value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) StandardOrErr() ([]*Standard, error) {
	if e.loadedTypes[5] {
		return e.Standard, nil
	}
	return nil, &NotLoadedError{edge: "standard"}
}

// NarrativesOrErr returns the Narratives value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) NarrativesOrErr() ([]*Narrative, error) {
	if e.loadedTypes[6] {
		return e.Narratives, nil
	}
	return nil, &NotLoadedError{edge: "narratives"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[7] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// ProgramsOrErr returns the Programs value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ProgramsOrErr() ([]*Program, error) {
	if e.loadedTypes[8] {
		return e.Programs, nil
	}
	return nil, &NotLoadedError{edge: "programs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ControlObjective) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case controlobjective.FieldTags, controlobjective.FieldDetails:
			values[i] = new([]byte)
		case controlobjective.FieldID, controlobjective.FieldCreatedBy, controlobjective.FieldUpdatedBy, controlobjective.FieldDeletedBy, controlobjective.FieldMappingID, controlobjective.FieldName, controlobjective.FieldDescription, controlobjective.FieldStatus, controlobjective.FieldControlObjectiveType, controlobjective.FieldVersion, controlobjective.FieldControlNumber, controlobjective.FieldFamily, controlobjective.FieldClass, controlobjective.FieldSource, controlobjective.FieldMappedFrameworks:
			values[i] = new(sql.NullString)
		case controlobjective.FieldCreatedAt, controlobjective.FieldUpdatedAt, controlobjective.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case controlobjective.ForeignKeys[0]: // control_controlobjectives
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ControlObjective fields.
func (co *ControlObjective) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case controlobjective.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				co.ID = value.String
			}
		case controlobjective.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				co.CreatedAt = value.Time
			}
		case controlobjective.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				co.UpdatedAt = value.Time
			}
		case controlobjective.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				co.CreatedBy = value.String
			}
		case controlobjective.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				co.UpdatedBy = value.String
			}
		case controlobjective.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				co.DeletedAt = value.Time
			}
		case controlobjective.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				co.DeletedBy = value.String
			}
		case controlobjective.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				co.MappingID = value.String
			}
		case controlobjective.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &co.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case controlobjective.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				co.Name = value.String
			}
		case controlobjective.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				co.Description = value.String
			}
		case controlobjective.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				co.Status = value.String
			}
		case controlobjective.FieldControlObjectiveType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_objective_type", values[i])
			} else if value.Valid {
				co.ControlObjectiveType = value.String
			}
		case controlobjective.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				co.Version = value.String
			}
		case controlobjective.FieldControlNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_number", values[i])
			} else if value.Valid {
				co.ControlNumber = value.String
			}
		case controlobjective.FieldFamily:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family", values[i])
			} else if value.Valid {
				co.Family = value.String
			}
		case controlobjective.FieldClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class", values[i])
			} else if value.Valid {
				co.Class = value.String
			}
		case controlobjective.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				co.Source = value.String
			}
		case controlobjective.FieldMappedFrameworks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapped_frameworks", values[i])
			} else if value.Valid {
				co.MappedFrameworks = value.String
			}
		case controlobjective.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &co.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		case controlobjective.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_controlobjectives", values[i])
			} else if value.Valid {
				co.control_controlobjectives = new(string)
				*co.control_controlobjectives = value.String
			}
		default:
			co.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ControlObjective.
// This includes values selected through modifiers, order, etc.
func (co *ControlObjective) Value(name string) (ent.Value, error) {
	return co.selectValues.Get(name)
}

// QueryPolicy queries the "policy" edge of the ControlObjective entity.
func (co *ControlObjective) QueryPolicy() *InternalPolicyQuery {
	return NewControlObjectiveClient(co.config).QueryPolicy(co)
}

// QueryControls queries the "controls" edge of the ControlObjective entity.
func (co *ControlObjective) QueryControls() *ControlQuery {
	return NewControlObjectiveClient(co.config).QueryControls(co)
}

// QueryProcedures queries the "procedures" edge of the ControlObjective entity.
func (co *ControlObjective) QueryProcedures() *ProcedureQuery {
	return NewControlObjectiveClient(co.config).QueryProcedures(co)
}

// QueryRisks queries the "risks" edge of the ControlObjective entity.
func (co *ControlObjective) QueryRisks() *RiskQuery {
	return NewControlObjectiveClient(co.config).QueryRisks(co)
}

// QuerySubcontrols queries the "subcontrols" edge of the ControlObjective entity.
func (co *ControlObjective) QuerySubcontrols() *SubcontrolQuery {
	return NewControlObjectiveClient(co.config).QuerySubcontrols(co)
}

// QueryStandard queries the "standard" edge of the ControlObjective entity.
func (co *ControlObjective) QueryStandard() *StandardQuery {
	return NewControlObjectiveClient(co.config).QueryStandard(co)
}

// QueryNarratives queries the "narratives" edge of the ControlObjective entity.
func (co *ControlObjective) QueryNarratives() *NarrativeQuery {
	return NewControlObjectiveClient(co.config).QueryNarratives(co)
}

// QueryTasks queries the "tasks" edge of the ControlObjective entity.
func (co *ControlObjective) QueryTasks() *TaskQuery {
	return NewControlObjectiveClient(co.config).QueryTasks(co)
}

// QueryPrograms queries the "programs" edge of the ControlObjective entity.
func (co *ControlObjective) QueryPrograms() *ProgramQuery {
	return NewControlObjectiveClient(co.config).QueryPrograms(co)
}

// Update returns a builder for updating this ControlObjective.
// Note that you need to call ControlObjective.Unwrap() before calling this method if this ControlObjective
// was returned from a transaction, and the transaction was committed or rolled back.
func (co *ControlObjective) Update() *ControlObjectiveUpdateOne {
	return NewControlObjectiveClient(co.config).UpdateOne(co)
}

// Unwrap unwraps the ControlObjective entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (co *ControlObjective) Unwrap() *ControlObjective {
	_tx, ok := co.config.driver.(*txDriver)
	if !ok {
		panic("generated: ControlObjective is not a transactional entity")
	}
	co.config.driver = _tx.drv
	return co
}

// String implements the fmt.Stringer.
func (co *ControlObjective) String() string {
	var builder strings.Builder
	builder.WriteString("ControlObjective(")
	builder.WriteString(fmt.Sprintf("id=%v, ", co.ID))
	builder.WriteString("created_at=")
	builder.WriteString(co.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(co.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(co.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(co.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(co.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(co.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(co.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", co.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(co.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(co.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(co.Status)
	builder.WriteString(", ")
	builder.WriteString("control_objective_type=")
	builder.WriteString(co.ControlObjectiveType)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(co.Version)
	builder.WriteString(", ")
	builder.WriteString("control_number=")
	builder.WriteString(co.ControlNumber)
	builder.WriteString(", ")
	builder.WriteString("family=")
	builder.WriteString(co.Family)
	builder.WriteString(", ")
	builder.WriteString("class=")
	builder.WriteString(co.Class)
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(co.Source)
	builder.WriteString(", ")
	builder.WriteString("mapped_frameworks=")
	builder.WriteString(co.MappedFrameworks)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", co.Details))
	builder.WriteByte(')')
	return builder.String()
}

// NamedPolicy returns the Policy named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedPolicy(name string) ([]*InternalPolicy, error) {
	if co.Edges.namedPolicy == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedPolicy[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedPolicy(name string, edges ...*InternalPolicy) {
	if co.Edges.namedPolicy == nil {
		co.Edges.namedPolicy = make(map[string][]*InternalPolicy)
	}
	if len(edges) == 0 {
		co.Edges.namedPolicy[name] = []*InternalPolicy{}
	} else {
		co.Edges.namedPolicy[name] = append(co.Edges.namedPolicy[name], edges...)
	}
}

// NamedControls returns the Controls named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedControls(name string) ([]*Control, error) {
	if co.Edges.namedControls == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedControls[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedControls(name string, edges ...*Control) {
	if co.Edges.namedControls == nil {
		co.Edges.namedControls = make(map[string][]*Control)
	}
	if len(edges) == 0 {
		co.Edges.namedControls[name] = []*Control{}
	} else {
		co.Edges.namedControls[name] = append(co.Edges.namedControls[name], edges...)
	}
}

// NamedProcedures returns the Procedures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedProcedures(name string) ([]*Procedure, error) {
	if co.Edges.namedProcedures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedProcedures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedProcedures(name string, edges ...*Procedure) {
	if co.Edges.namedProcedures == nil {
		co.Edges.namedProcedures = make(map[string][]*Procedure)
	}
	if len(edges) == 0 {
		co.Edges.namedProcedures[name] = []*Procedure{}
	} else {
		co.Edges.namedProcedures[name] = append(co.Edges.namedProcedures[name], edges...)
	}
}

// NamedRisks returns the Risks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedRisks(name string) ([]*Risk, error) {
	if co.Edges.namedRisks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedRisks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedRisks(name string, edges ...*Risk) {
	if co.Edges.namedRisks == nil {
		co.Edges.namedRisks = make(map[string][]*Risk)
	}
	if len(edges) == 0 {
		co.Edges.namedRisks[name] = []*Risk{}
	} else {
		co.Edges.namedRisks[name] = append(co.Edges.namedRisks[name], edges...)
	}
}

// NamedSubcontrols returns the Subcontrols named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedSubcontrols(name string) ([]*Subcontrol, error) {
	if co.Edges.namedSubcontrols == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedSubcontrols[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedSubcontrols(name string, edges ...*Subcontrol) {
	if co.Edges.namedSubcontrols == nil {
		co.Edges.namedSubcontrols = make(map[string][]*Subcontrol)
	}
	if len(edges) == 0 {
		co.Edges.namedSubcontrols[name] = []*Subcontrol{}
	} else {
		co.Edges.namedSubcontrols[name] = append(co.Edges.namedSubcontrols[name], edges...)
	}
}

// NamedStandard returns the Standard named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedStandard(name string) ([]*Standard, error) {
	if co.Edges.namedStandard == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedStandard[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedStandard(name string, edges ...*Standard) {
	if co.Edges.namedStandard == nil {
		co.Edges.namedStandard = make(map[string][]*Standard)
	}
	if len(edges) == 0 {
		co.Edges.namedStandard[name] = []*Standard{}
	} else {
		co.Edges.namedStandard[name] = append(co.Edges.namedStandard[name], edges...)
	}
}

// NamedNarratives returns the Narratives named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedNarratives(name string) ([]*Narrative, error) {
	if co.Edges.namedNarratives == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedNarratives[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedNarratives(name string, edges ...*Narrative) {
	if co.Edges.namedNarratives == nil {
		co.Edges.namedNarratives = make(map[string][]*Narrative)
	}
	if len(edges) == 0 {
		co.Edges.namedNarratives[name] = []*Narrative{}
	} else {
		co.Edges.namedNarratives[name] = append(co.Edges.namedNarratives[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedTasks(name string) ([]*Task, error) {
	if co.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedTasks(name string, edges ...*Task) {
	if co.Edges.namedTasks == nil {
		co.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		co.Edges.namedTasks[name] = []*Task{}
	} else {
		co.Edges.namedTasks[name] = append(co.Edges.namedTasks[name], edges...)
	}
}

// NamedPrograms returns the Programs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedPrograms(name string) ([]*Program, error) {
	if co.Edges.namedPrograms == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedPrograms[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedPrograms(name string, edges ...*Program) {
	if co.Edges.namedPrograms == nil {
		co.Edges.namedPrograms = make(map[string][]*Program)
	}
	if len(edges) == 0 {
		co.Edges.namedPrograms[name] = []*Program{}
	} else {
		co.Edges.namedPrograms[name] = append(co.Edges.namedPrograms[name], edges...)
	}
}

// ControlObjectives is a parsable slice of ControlObjective.
type ControlObjectives []*ControlObjective
