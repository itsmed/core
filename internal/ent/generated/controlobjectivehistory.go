// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/controlobjectivehistory"
	"github.com/theopenlane/entx/history"
)

// ControlObjectiveHistory is the model entity for the ControlObjectiveHistory schema.
type ControlObjectiveHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation history.OpType `json:"operation,omitempty"`
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
	Details      map[string]interface{} `json:"details,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ControlObjectiveHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case controlobjectivehistory.FieldTags, controlobjectivehistory.FieldDetails:
			values[i] = new([]byte)
		case controlobjectivehistory.FieldOperation:
			values[i] = new(history.OpType)
		case controlobjectivehistory.FieldID, controlobjectivehistory.FieldRef, controlobjectivehistory.FieldCreatedBy, controlobjectivehistory.FieldUpdatedBy, controlobjectivehistory.FieldDeletedBy, controlobjectivehistory.FieldMappingID, controlobjectivehistory.FieldName, controlobjectivehistory.FieldDescription, controlobjectivehistory.FieldStatus, controlobjectivehistory.FieldControlObjectiveType, controlobjectivehistory.FieldVersion, controlobjectivehistory.FieldControlNumber, controlobjectivehistory.FieldFamily, controlobjectivehistory.FieldClass, controlobjectivehistory.FieldSource, controlobjectivehistory.FieldMappedFrameworks:
			values[i] = new(sql.NullString)
		case controlobjectivehistory.FieldHistoryTime, controlobjectivehistory.FieldCreatedAt, controlobjectivehistory.FieldUpdatedAt, controlobjectivehistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ControlObjectiveHistory fields.
func (coh *ControlObjectiveHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case controlobjectivehistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				coh.ID = value.String
			}
		case controlobjectivehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				coh.HistoryTime = value.Time
			}
		case controlobjectivehistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				coh.Ref = value.String
			}
		case controlobjectivehistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				coh.Operation = *value
			}
		case controlobjectivehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				coh.CreatedAt = value.Time
			}
		case controlobjectivehistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				coh.UpdatedAt = value.Time
			}
		case controlobjectivehistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				coh.CreatedBy = value.String
			}
		case controlobjectivehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				coh.UpdatedBy = value.String
			}
		case controlobjectivehistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				coh.DeletedAt = value.Time
			}
		case controlobjectivehistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				coh.DeletedBy = value.String
			}
		case controlobjectivehistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				coh.MappingID = value.String
			}
		case controlobjectivehistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &coh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case controlobjectivehistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				coh.Name = value.String
			}
		case controlobjectivehistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				coh.Description = value.String
			}
		case controlobjectivehistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				coh.Status = value.String
			}
		case controlobjectivehistory.FieldControlObjectiveType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_objective_type", values[i])
			} else if value.Valid {
				coh.ControlObjectiveType = value.String
			}
		case controlobjectivehistory.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				coh.Version = value.String
			}
		case controlobjectivehistory.FieldControlNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_number", values[i])
			} else if value.Valid {
				coh.ControlNumber = value.String
			}
		case controlobjectivehistory.FieldFamily:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family", values[i])
			} else if value.Valid {
				coh.Family = value.String
			}
		case controlobjectivehistory.FieldClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class", values[i])
			} else if value.Valid {
				coh.Class = value.String
			}
		case controlobjectivehistory.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				coh.Source = value.String
			}
		case controlobjectivehistory.FieldMappedFrameworks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapped_frameworks", values[i])
			} else if value.Valid {
				coh.MappedFrameworks = value.String
			}
		case controlobjectivehistory.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &coh.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		default:
			coh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ControlObjectiveHistory.
// This includes values selected through modifiers, order, etc.
func (coh *ControlObjectiveHistory) Value(name string) (ent.Value, error) {
	return coh.selectValues.Get(name)
}

// Update returns a builder for updating this ControlObjectiveHistory.
// Note that you need to call ControlObjectiveHistory.Unwrap() before calling this method if this ControlObjectiveHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (coh *ControlObjectiveHistory) Update() *ControlObjectiveHistoryUpdateOne {
	return NewControlObjectiveHistoryClient(coh.config).UpdateOne(coh)
}

// Unwrap unwraps the ControlObjectiveHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (coh *ControlObjectiveHistory) Unwrap() *ControlObjectiveHistory {
	_tx, ok := coh.config.driver.(*txDriver)
	if !ok {
		panic("generated: ControlObjectiveHistory is not a transactional entity")
	}
	coh.config.driver = _tx.drv
	return coh
}

// String implements the fmt.Stringer.
func (coh *ControlObjectiveHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ControlObjectiveHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", coh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(coh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(coh.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", coh.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(coh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(coh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(coh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(coh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(coh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(coh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(coh.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", coh.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(coh.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(coh.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(coh.Status)
	builder.WriteString(", ")
	builder.WriteString("control_objective_type=")
	builder.WriteString(coh.ControlObjectiveType)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(coh.Version)
	builder.WriteString(", ")
	builder.WriteString("control_number=")
	builder.WriteString(coh.ControlNumber)
	builder.WriteString(", ")
	builder.WriteString("family=")
	builder.WriteString(coh.Family)
	builder.WriteString(", ")
	builder.WriteString("class=")
	builder.WriteString(coh.Class)
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(coh.Source)
	builder.WriteString(", ")
	builder.WriteString("mapped_frameworks=")
	builder.WriteString(coh.MappedFrameworks)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", coh.Details))
	builder.WriteByte(')')
	return builder.String()
}

// ControlObjectiveHistories is a parsable slice of ControlObjectiveHistory.
type ControlObjectiveHistories []*ControlObjectiveHistory