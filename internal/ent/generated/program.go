// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/program"
	"github.com/theopenlane/core/pkg/enums"
)

// Program is the model entity for the Program schema.
type Program struct {
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
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// the organization id that owns the object
	OwnerID string `json:"owner_id,omitempty"`
	// the name of the program
	Name string `json:"name,omitempty"`
	// the description of the program
	Description string `json:"description,omitempty"`
	// the status of the program
	Status enums.ProgramStatus `json:"status,omitempty"`
	// the start date of the period
	StartDate time.Time `json:"start_date,omitempty"`
	// the end date of the period
	EndDate time.Time `json:"end_date,omitempty"`
	// is the program ready for the auditor
	AuditorReady bool `json:"auditor_ready,omitempty"`
	// can the auditor write comments
	AuditorWriteComments bool `json:"auditor_write_comments,omitempty"`
	// can the auditor read comments
	AuditorReadComments bool `json:"auditor_read_comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProgramQuery when eager-loading is set.
	Edges        ProgramEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProgramEdges holds the relations/edges for other nodes in the graph.
type ProgramEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// Controls holds the value of the controls edge.
	Controls []*Control `json:"controls,omitempty"`
	// Subcontrols holds the value of the subcontrols edge.
	Subcontrols []*Subcontrol `json:"subcontrols,omitempty"`
	// Controlobjectives holds the value of the controlobjectives edge.
	Controlobjectives []*ControlObjective `json:"controlobjectives,omitempty"`
	// Policies holds the value of the policies edge.
	Policies []*InternalPolicy `json:"policies,omitempty"`
	// Procedures holds the value of the procedures edge.
	Procedures []*Procedure `json:"procedures,omitempty"`
	// Risks holds the value of the risks edge.
	Risks []*Risk `json:"risks,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Notes holds the value of the notes edge.
	Notes []*Note `json:"notes,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Narratives holds the value of the narratives edge.
	Narratives []*Narrative `json:"narratives,omitempty"`
	// Actionplans holds the value of the actionplans edge.
	Actionplans []*ActionPlan `json:"actionplans,omitempty"`
	// the framework(s) that the program is based on
	Standards []*Standard `json:"standards,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Members holds the value of the members edge.
	Members []*ProgramMembership `json:"members,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [15]bool
	// totalCount holds the count of the edges above.
	totalCount [15]map[string]int

	namedControls          map[string][]*Control
	namedSubcontrols       map[string][]*Subcontrol
	namedControlobjectives map[string][]*ControlObjective
	namedPolicies          map[string][]*InternalPolicy
	namedProcedures        map[string][]*Procedure
	namedRisks             map[string][]*Risk
	namedTasks             map[string][]*Task
	namedNotes             map[string][]*Note
	namedFiles             map[string][]*File
	namedNarratives        map[string][]*Narrative
	namedActionplans       map[string][]*ActionPlan
	namedStandards         map[string][]*Standard
	namedUsers             map[string][]*User
	namedMembers           map[string][]*ProgramMembership
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProgramEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ControlsOrErr returns the Controls value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) ControlsOrErr() ([]*Control, error) {
	if e.loadedTypes[1] {
		return e.Controls, nil
	}
	return nil, &NotLoadedError{edge: "controls"}
}

// SubcontrolsOrErr returns the Subcontrols value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) SubcontrolsOrErr() ([]*Subcontrol, error) {
	if e.loadedTypes[2] {
		return e.Subcontrols, nil
	}
	return nil, &NotLoadedError{edge: "subcontrols"}
}

// ControlobjectivesOrErr returns the Controlobjectives value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) ControlobjectivesOrErr() ([]*ControlObjective, error) {
	if e.loadedTypes[3] {
		return e.Controlobjectives, nil
	}
	return nil, &NotLoadedError{edge: "controlobjectives"}
}

// PoliciesOrErr returns the Policies value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) PoliciesOrErr() ([]*InternalPolicy, error) {
	if e.loadedTypes[4] {
		return e.Policies, nil
	}
	return nil, &NotLoadedError{edge: "policies"}
}

// ProceduresOrErr returns the Procedures value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) ProceduresOrErr() ([]*Procedure, error) {
	if e.loadedTypes[5] {
		return e.Procedures, nil
	}
	return nil, &NotLoadedError{edge: "procedures"}
}

// RisksOrErr returns the Risks value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) RisksOrErr() ([]*Risk, error) {
	if e.loadedTypes[6] {
		return e.Risks, nil
	}
	return nil, &NotLoadedError{edge: "risks"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[7] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// NotesOrErr returns the Notes value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) NotesOrErr() ([]*Note, error) {
	if e.loadedTypes[8] {
		return e.Notes, nil
	}
	return nil, &NotLoadedError{edge: "notes"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[9] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// NarrativesOrErr returns the Narratives value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) NarrativesOrErr() ([]*Narrative, error) {
	if e.loadedTypes[10] {
		return e.Narratives, nil
	}
	return nil, &NotLoadedError{edge: "narratives"}
}

// ActionplansOrErr returns the Actionplans value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) ActionplansOrErr() ([]*ActionPlan, error) {
	if e.loadedTypes[11] {
		return e.Actionplans, nil
	}
	return nil, &NotLoadedError{edge: "actionplans"}
}

// StandardsOrErr returns the Standards value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) StandardsOrErr() ([]*Standard, error) {
	if e.loadedTypes[12] {
		return e.Standards, nil
	}
	return nil, &NotLoadedError{edge: "standards"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[13] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) MembersOrErr() ([]*ProgramMembership, error) {
	if e.loadedTypes[14] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Program) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case program.FieldTags:
			values[i] = new([]byte)
		case program.FieldAuditorReady, program.FieldAuditorWriteComments, program.FieldAuditorReadComments:
			values[i] = new(sql.NullBool)
		case program.FieldID, program.FieldCreatedBy, program.FieldUpdatedBy, program.FieldMappingID, program.FieldDeletedBy, program.FieldOwnerID, program.FieldName, program.FieldDescription, program.FieldStatus:
			values[i] = new(sql.NullString)
		case program.FieldCreatedAt, program.FieldUpdatedAt, program.FieldDeletedAt, program.FieldStartDate, program.FieldEndDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Program fields.
func (pr *Program) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case program.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case program.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case program.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case program.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = value.String
			}
		case program.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pr.UpdatedBy = value.String
			}
		case program.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				pr.MappingID = value.String
			}
		case program.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case program.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pr.DeletedBy = value.String
			}
		case program.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case program.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				pr.OwnerID = value.String
			}
		case program.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case program.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case program.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = enums.ProgramStatus(value.String)
			}
		case program.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pr.StartDate = value.Time
			}
		case program.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				pr.EndDate = value.Time
			}
		case program.FieldAuditorReady:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auditor_ready", values[i])
			} else if value.Valid {
				pr.AuditorReady = value.Bool
			}
		case program.FieldAuditorWriteComments:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auditor_write_comments", values[i])
			} else if value.Valid {
				pr.AuditorWriteComments = value.Bool
			}
		case program.FieldAuditorReadComments:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auditor_read_comments", values[i])
			} else if value.Valid {
				pr.AuditorReadComments = value.Bool
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Program.
// This includes values selected through modifiers, order, etc.
func (pr *Program) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Program entity.
func (pr *Program) QueryOwner() *OrganizationQuery {
	return NewProgramClient(pr.config).QueryOwner(pr)
}

// QueryControls queries the "controls" edge of the Program entity.
func (pr *Program) QueryControls() *ControlQuery {
	return NewProgramClient(pr.config).QueryControls(pr)
}

// QuerySubcontrols queries the "subcontrols" edge of the Program entity.
func (pr *Program) QuerySubcontrols() *SubcontrolQuery {
	return NewProgramClient(pr.config).QuerySubcontrols(pr)
}

// QueryControlobjectives queries the "controlobjectives" edge of the Program entity.
func (pr *Program) QueryControlobjectives() *ControlObjectiveQuery {
	return NewProgramClient(pr.config).QueryControlobjectives(pr)
}

// QueryPolicies queries the "policies" edge of the Program entity.
func (pr *Program) QueryPolicies() *InternalPolicyQuery {
	return NewProgramClient(pr.config).QueryPolicies(pr)
}

// QueryProcedures queries the "procedures" edge of the Program entity.
func (pr *Program) QueryProcedures() *ProcedureQuery {
	return NewProgramClient(pr.config).QueryProcedures(pr)
}

// QueryRisks queries the "risks" edge of the Program entity.
func (pr *Program) QueryRisks() *RiskQuery {
	return NewProgramClient(pr.config).QueryRisks(pr)
}

// QueryTasks queries the "tasks" edge of the Program entity.
func (pr *Program) QueryTasks() *TaskQuery {
	return NewProgramClient(pr.config).QueryTasks(pr)
}

// QueryNotes queries the "notes" edge of the Program entity.
func (pr *Program) QueryNotes() *NoteQuery {
	return NewProgramClient(pr.config).QueryNotes(pr)
}

// QueryFiles queries the "files" edge of the Program entity.
func (pr *Program) QueryFiles() *FileQuery {
	return NewProgramClient(pr.config).QueryFiles(pr)
}

// QueryNarratives queries the "narratives" edge of the Program entity.
func (pr *Program) QueryNarratives() *NarrativeQuery {
	return NewProgramClient(pr.config).QueryNarratives(pr)
}

// QueryActionplans queries the "actionplans" edge of the Program entity.
func (pr *Program) QueryActionplans() *ActionPlanQuery {
	return NewProgramClient(pr.config).QueryActionplans(pr)
}

// QueryStandards queries the "standards" edge of the Program entity.
func (pr *Program) QueryStandards() *StandardQuery {
	return NewProgramClient(pr.config).QueryStandards(pr)
}

// QueryUsers queries the "users" edge of the Program entity.
func (pr *Program) QueryUsers() *UserQuery {
	return NewProgramClient(pr.config).QueryUsers(pr)
}

// QueryMembers queries the "members" edge of the Program entity.
func (pr *Program) QueryMembers() *ProgramMembershipQuery {
	return NewProgramClient(pr.config).QueryMembers(pr)
}

// Update returns a builder for updating this Program.
// Note that you need to call Program.Unwrap() before calling this method if this Program
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Program) Update() *ProgramUpdateOne {
	return NewProgramClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Program entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Program) Unwrap() *Program {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Program is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Program) String() string {
	var builder strings.Builder
	builder.WriteString("Program(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(pr.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(pr.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", pr.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(pr.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(pr.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(pr.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("auditor_ready=")
	builder.WriteString(fmt.Sprintf("%v", pr.AuditorReady))
	builder.WriteString(", ")
	builder.WriteString("auditor_write_comments=")
	builder.WriteString(fmt.Sprintf("%v", pr.AuditorWriteComments))
	builder.WriteString(", ")
	builder.WriteString("auditor_read_comments=")
	builder.WriteString(fmt.Sprintf("%v", pr.AuditorReadComments))
	builder.WriteByte(')')
	return builder.String()
}

// NamedControls returns the Controls named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedControls(name string) ([]*Control, error) {
	if pr.Edges.namedControls == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedControls[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedControls(name string, edges ...*Control) {
	if pr.Edges.namedControls == nil {
		pr.Edges.namedControls = make(map[string][]*Control)
	}
	if len(edges) == 0 {
		pr.Edges.namedControls[name] = []*Control{}
	} else {
		pr.Edges.namedControls[name] = append(pr.Edges.namedControls[name], edges...)
	}
}

// NamedSubcontrols returns the Subcontrols named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedSubcontrols(name string) ([]*Subcontrol, error) {
	if pr.Edges.namedSubcontrols == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedSubcontrols[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedSubcontrols(name string, edges ...*Subcontrol) {
	if pr.Edges.namedSubcontrols == nil {
		pr.Edges.namedSubcontrols = make(map[string][]*Subcontrol)
	}
	if len(edges) == 0 {
		pr.Edges.namedSubcontrols[name] = []*Subcontrol{}
	} else {
		pr.Edges.namedSubcontrols[name] = append(pr.Edges.namedSubcontrols[name], edges...)
	}
}

// NamedControlobjectives returns the Controlobjectives named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedControlobjectives(name string) ([]*ControlObjective, error) {
	if pr.Edges.namedControlobjectives == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedControlobjectives[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedControlobjectives(name string, edges ...*ControlObjective) {
	if pr.Edges.namedControlobjectives == nil {
		pr.Edges.namedControlobjectives = make(map[string][]*ControlObjective)
	}
	if len(edges) == 0 {
		pr.Edges.namedControlobjectives[name] = []*ControlObjective{}
	} else {
		pr.Edges.namedControlobjectives[name] = append(pr.Edges.namedControlobjectives[name], edges...)
	}
}

// NamedPolicies returns the Policies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedPolicies(name string) ([]*InternalPolicy, error) {
	if pr.Edges.namedPolicies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedPolicies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedPolicies(name string, edges ...*InternalPolicy) {
	if pr.Edges.namedPolicies == nil {
		pr.Edges.namedPolicies = make(map[string][]*InternalPolicy)
	}
	if len(edges) == 0 {
		pr.Edges.namedPolicies[name] = []*InternalPolicy{}
	} else {
		pr.Edges.namedPolicies[name] = append(pr.Edges.namedPolicies[name], edges...)
	}
}

// NamedProcedures returns the Procedures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedProcedures(name string) ([]*Procedure, error) {
	if pr.Edges.namedProcedures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedProcedures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedProcedures(name string, edges ...*Procedure) {
	if pr.Edges.namedProcedures == nil {
		pr.Edges.namedProcedures = make(map[string][]*Procedure)
	}
	if len(edges) == 0 {
		pr.Edges.namedProcedures[name] = []*Procedure{}
	} else {
		pr.Edges.namedProcedures[name] = append(pr.Edges.namedProcedures[name], edges...)
	}
}

// NamedRisks returns the Risks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedRisks(name string) ([]*Risk, error) {
	if pr.Edges.namedRisks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedRisks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedRisks(name string, edges ...*Risk) {
	if pr.Edges.namedRisks == nil {
		pr.Edges.namedRisks = make(map[string][]*Risk)
	}
	if len(edges) == 0 {
		pr.Edges.namedRisks[name] = []*Risk{}
	} else {
		pr.Edges.namedRisks[name] = append(pr.Edges.namedRisks[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedTasks(name string) ([]*Task, error) {
	if pr.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedTasks(name string, edges ...*Task) {
	if pr.Edges.namedTasks == nil {
		pr.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		pr.Edges.namedTasks[name] = []*Task{}
	} else {
		pr.Edges.namedTasks[name] = append(pr.Edges.namedTasks[name], edges...)
	}
}

// NamedNotes returns the Notes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedNotes(name string) ([]*Note, error) {
	if pr.Edges.namedNotes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedNotes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedNotes(name string, edges ...*Note) {
	if pr.Edges.namedNotes == nil {
		pr.Edges.namedNotes = make(map[string][]*Note)
	}
	if len(edges) == 0 {
		pr.Edges.namedNotes[name] = []*Note{}
	} else {
		pr.Edges.namedNotes[name] = append(pr.Edges.namedNotes[name], edges...)
	}
}

// NamedFiles returns the Files named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedFiles(name string) ([]*File, error) {
	if pr.Edges.namedFiles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedFiles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedFiles(name string, edges ...*File) {
	if pr.Edges.namedFiles == nil {
		pr.Edges.namedFiles = make(map[string][]*File)
	}
	if len(edges) == 0 {
		pr.Edges.namedFiles[name] = []*File{}
	} else {
		pr.Edges.namedFiles[name] = append(pr.Edges.namedFiles[name], edges...)
	}
}

// NamedNarratives returns the Narratives named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedNarratives(name string) ([]*Narrative, error) {
	if pr.Edges.namedNarratives == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedNarratives[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedNarratives(name string, edges ...*Narrative) {
	if pr.Edges.namedNarratives == nil {
		pr.Edges.namedNarratives = make(map[string][]*Narrative)
	}
	if len(edges) == 0 {
		pr.Edges.namedNarratives[name] = []*Narrative{}
	} else {
		pr.Edges.namedNarratives[name] = append(pr.Edges.namedNarratives[name], edges...)
	}
}

// NamedActionplans returns the Actionplans named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedActionplans(name string) ([]*ActionPlan, error) {
	if pr.Edges.namedActionplans == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedActionplans[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedActionplans(name string, edges ...*ActionPlan) {
	if pr.Edges.namedActionplans == nil {
		pr.Edges.namedActionplans = make(map[string][]*ActionPlan)
	}
	if len(edges) == 0 {
		pr.Edges.namedActionplans[name] = []*ActionPlan{}
	} else {
		pr.Edges.namedActionplans[name] = append(pr.Edges.namedActionplans[name], edges...)
	}
}

// NamedStandards returns the Standards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedStandards(name string) ([]*Standard, error) {
	if pr.Edges.namedStandards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedStandards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedStandards(name string, edges ...*Standard) {
	if pr.Edges.namedStandards == nil {
		pr.Edges.namedStandards = make(map[string][]*Standard)
	}
	if len(edges) == 0 {
		pr.Edges.namedStandards[name] = []*Standard{}
	} else {
		pr.Edges.namedStandards[name] = append(pr.Edges.namedStandards[name], edges...)
	}
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedUsers(name string) ([]*User, error) {
	if pr.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedUsers(name string, edges ...*User) {
	if pr.Edges.namedUsers == nil {
		pr.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		pr.Edges.namedUsers[name] = []*User{}
	} else {
		pr.Edges.namedUsers[name] = append(pr.Edges.namedUsers[name], edges...)
	}
}

// NamedMembers returns the Members named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Program) NamedMembers(name string) ([]*ProgramMembership, error) {
	if pr.Edges.namedMembers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedMembers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Program) appendNamedMembers(name string, edges ...*ProgramMembership) {
	if pr.Edges.namedMembers == nil {
		pr.Edges.namedMembers = make(map[string][]*ProgramMembership)
	}
	if len(edges) == 0 {
		pr.Edges.namedMembers[name] = []*ProgramMembership{}
	} else {
		pr.Edges.namedMembers[name] = append(pr.Edges.namedMembers[name], edges...)
	}
}

// Programs is a parsable slice of Program.
type Programs []*Program