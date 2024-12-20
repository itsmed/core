// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/file"
)

// File is the model entity for the File schema.
type File struct {
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
	// the name of the file provided in the payload key without the extension
	ProvidedFileName string `json:"provided_file_name,omitempty"`
	// the extension of the file provided
	ProvidedFileExtension string `json:"provided_file_extension,omitempty"`
	// the computed size of the file in the original http request
	ProvidedFileSize int64 `json:"provided_file_size,omitempty"`
	// PersistedFileSize holds the value of the "persisted_file_size" field.
	PersistedFileSize int64 `json:"persisted_file_size,omitempty"`
	// the mime type detected by the system
	DetectedMimeType string `json:"detected_mime_type,omitempty"`
	// the computed md5 hash of the file calculated after we received the contents of the file, but before the file was written to permanent storage
	Md5Hash string `json:"md5_hash,omitempty"`
	// the content type of the HTTP request - may be different than MIME type as multipart-form can transmit multiple files and different types
	DetectedContentType string `json:"detected_content_type,omitempty"`
	// the key parsed out of a multipart-form request; if we allow multiple files to be uploaded we may want our API specifications to require the use of different keys allowing us to perform easier conditional evaluation on the key and what to do with the file based on key
	StoreKey string `json:"store_key,omitempty"`
	// the category type of the file, if any (e.g. evidence, invoice, etc.)
	CategoryType string `json:"category_type,omitempty"`
	// the full URI of the file
	URI string `json:"uri,omitempty"`
	// the storage scheme of the file, e.g. file://, s3://, etc.
	StorageScheme string `json:"storage_scheme,omitempty"`
	// the storage volume of the file which typically will be the organization ID the file belongs to - this is not a literal volume but the overlay file system mapping
	StorageVolume string `json:"storage_volume,omitempty"`
	// the storage path is the second-level directory of the file path, typically the correlating logical object ID the file is associated with; files can be stand alone objects and not always correlated to a logical one, so this path of the tree may be empty
	StoragePath string `json:"storage_path,omitempty"`
	// the contents of the file
	FileContents []byte `json:"file_contents,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Organization holds the value of the organization edge.
	Organization []*Organization `json:"organization,omitempty"`
	// Group holds the value of the group edge.
	Group []*Group `json:"group,omitempty"`
	// Contact holds the value of the contact edge.
	Contact []*Contact `json:"contact,omitempty"`
	// Entity holds the value of the entity edge.
	Entity []*Entity `json:"entity,omitempty"`
	// Usersetting holds the value of the usersetting edge.
	Usersetting []*UserSetting `json:"usersetting,omitempty"`
	// Organizationsetting holds the value of the organizationsetting edge.
	Organizationsetting []*OrganizationSetting `json:"organizationsetting,omitempty"`
	// Template holds the value of the template edge.
	Template []*Template `json:"template,omitempty"`
	// Documentdata holds the value of the documentdata edge.
	Documentdata []*DocumentData `json:"documentdata,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// Program holds the value of the program edge.
	Program []*Program `json:"program,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [11]bool
	// totalCount holds the count of the edges above.
	totalCount [11]map[string]int

	namedUser                map[string][]*User
	namedOrganization        map[string][]*Organization
	namedGroup               map[string][]*Group
	namedContact             map[string][]*Contact
	namedEntity              map[string][]*Entity
	namedUsersetting         map[string][]*UserSetting
	namedOrganizationsetting map[string][]*OrganizationSetting
	namedTemplate            map[string][]*Template
	namedDocumentdata        map[string][]*DocumentData
	namedEvents              map[string][]*Event
	namedProgram             map[string][]*Program
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) OrganizationOrErr() ([]*Organization, error) {
	if e.loadedTypes[1] {
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) GroupOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// ContactOrErr returns the Contact value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) ContactOrErr() ([]*Contact, error) {
	if e.loadedTypes[3] {
		return e.Contact, nil
	}
	return nil, &NotLoadedError{edge: "contact"}
}

// EntityOrErr returns the Entity value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) EntityOrErr() ([]*Entity, error) {
	if e.loadedTypes[4] {
		return e.Entity, nil
	}
	return nil, &NotLoadedError{edge: "entity"}
}

// UsersettingOrErr returns the Usersetting value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) UsersettingOrErr() ([]*UserSetting, error) {
	if e.loadedTypes[5] {
		return e.Usersetting, nil
	}
	return nil, &NotLoadedError{edge: "usersetting"}
}

// OrganizationsettingOrErr returns the Organizationsetting value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) OrganizationsettingOrErr() ([]*OrganizationSetting, error) {
	if e.loadedTypes[6] {
		return e.Organizationsetting, nil
	}
	return nil, &NotLoadedError{edge: "organizationsetting"}
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) TemplateOrErr() ([]*Template, error) {
	if e.loadedTypes[7] {
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// DocumentdataOrErr returns the Documentdata value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) DocumentdataOrErr() ([]*DocumentData, error) {
	if e.loadedTypes[8] {
		return e.Documentdata, nil
	}
	return nil, &NotLoadedError{edge: "documentdata"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[9] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// ProgramOrErr returns the Program value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) ProgramOrErr() ([]*Program, error) {
	if e.loadedTypes[10] {
		return e.Program, nil
	}
	return nil, &NotLoadedError{edge: "program"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldTags, file.FieldFileContents:
			values[i] = new([]byte)
		case file.FieldProvidedFileSize, file.FieldPersistedFileSize:
			values[i] = new(sql.NullInt64)
		case file.FieldID, file.FieldCreatedBy, file.FieldUpdatedBy, file.FieldDeletedBy, file.FieldMappingID, file.FieldProvidedFileName, file.FieldProvidedFileExtension, file.FieldDetectedMimeType, file.FieldMd5Hash, file.FieldDetectedContentType, file.FieldStoreKey, file.FieldCategoryType, file.FieldURI, file.FieldStorageScheme, file.FieldStorageVolume, file.FieldStoragePath:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt, file.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = value.String
			}
		case file.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = value.String
			}
		case file.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				f.DeletedAt = value.Time
			}
		case file.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				f.DeletedBy = value.String
			}
		case file.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				f.MappingID = value.String
			}
		case file.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case file.FieldProvidedFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provided_file_name", values[i])
			} else if value.Valid {
				f.ProvidedFileName = value.String
			}
		case file.FieldProvidedFileExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provided_file_extension", values[i])
			} else if value.Valid {
				f.ProvidedFileExtension = value.String
			}
		case file.FieldProvidedFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field provided_file_size", values[i])
			} else if value.Valid {
				f.ProvidedFileSize = value.Int64
			}
		case file.FieldPersistedFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field persisted_file_size", values[i])
			} else if value.Valid {
				f.PersistedFileSize = value.Int64
			}
		case file.FieldDetectedMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detected_mime_type", values[i])
			} else if value.Valid {
				f.DetectedMimeType = value.String
			}
		case file.FieldMd5Hash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5_hash", values[i])
			} else if value.Valid {
				f.Md5Hash = value.String
			}
		case file.FieldDetectedContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detected_content_type", values[i])
			} else if value.Valid {
				f.DetectedContentType = value.String
			}
		case file.FieldStoreKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_key", values[i])
			} else if value.Valid {
				f.StoreKey = value.String
			}
		case file.FieldCategoryType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category_type", values[i])
			} else if value.Valid {
				f.CategoryType = value.String
			}
		case file.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				f.URI = value.String
			}
		case file.FieldStorageScheme:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storage_scheme", values[i])
			} else if value.Valid {
				f.StorageScheme = value.String
			}
		case file.FieldStorageVolume:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storage_volume", values[i])
			} else if value.Valid {
				f.StorageVolume = value.String
			}
		case file.FieldStoragePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storage_path", values[i])
			} else if value.Valid {
				f.StoragePath = value.String
			}
		case file.FieldFileContents:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field file_contents", values[i])
			} else if value != nil {
				f.FileContents = *value
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the File entity.
func (f *File) QueryUser() *UserQuery {
	return NewFileClient(f.config).QueryUser(f)
}

// QueryOrganization queries the "organization" edge of the File entity.
func (f *File) QueryOrganization() *OrganizationQuery {
	return NewFileClient(f.config).QueryOrganization(f)
}

// QueryGroup queries the "group" edge of the File entity.
func (f *File) QueryGroup() *GroupQuery {
	return NewFileClient(f.config).QueryGroup(f)
}

// QueryContact queries the "contact" edge of the File entity.
func (f *File) QueryContact() *ContactQuery {
	return NewFileClient(f.config).QueryContact(f)
}

// QueryEntity queries the "entity" edge of the File entity.
func (f *File) QueryEntity() *EntityQuery {
	return NewFileClient(f.config).QueryEntity(f)
}

// QueryUsersetting queries the "usersetting" edge of the File entity.
func (f *File) QueryUsersetting() *UserSettingQuery {
	return NewFileClient(f.config).QueryUsersetting(f)
}

// QueryOrganizationsetting queries the "organizationsetting" edge of the File entity.
func (f *File) QueryOrganizationsetting() *OrganizationSettingQuery {
	return NewFileClient(f.config).QueryOrganizationsetting(f)
}

// QueryTemplate queries the "template" edge of the File entity.
func (f *File) QueryTemplate() *TemplateQuery {
	return NewFileClient(f.config).QueryTemplate(f)
}

// QueryDocumentdata queries the "documentdata" edge of the File entity.
func (f *File) QueryDocumentdata() *DocumentDataQuery {
	return NewFileClient(f.config).QueryDocumentdata(f)
}

// QueryEvents queries the "events" edge of the File entity.
func (f *File) QueryEvents() *EventQuery {
	return NewFileClient(f.config).QueryEvents(f)
}

// QueryProgram queries the "program" edge of the File entity.
func (f *File) QueryProgram() *ProgramQuery {
	return NewFileClient(f.config).QueryProgram(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("generated: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(f.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(f.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(f.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(f.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(f.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", f.Tags))
	builder.WriteString(", ")
	builder.WriteString("provided_file_name=")
	builder.WriteString(f.ProvidedFileName)
	builder.WriteString(", ")
	builder.WriteString("provided_file_extension=")
	builder.WriteString(f.ProvidedFileExtension)
	builder.WriteString(", ")
	builder.WriteString("provided_file_size=")
	builder.WriteString(fmt.Sprintf("%v", f.ProvidedFileSize))
	builder.WriteString(", ")
	builder.WriteString("persisted_file_size=")
	builder.WriteString(fmt.Sprintf("%v", f.PersistedFileSize))
	builder.WriteString(", ")
	builder.WriteString("detected_mime_type=")
	builder.WriteString(f.DetectedMimeType)
	builder.WriteString(", ")
	builder.WriteString("md5_hash=")
	builder.WriteString(f.Md5Hash)
	builder.WriteString(", ")
	builder.WriteString("detected_content_type=")
	builder.WriteString(f.DetectedContentType)
	builder.WriteString(", ")
	builder.WriteString("store_key=")
	builder.WriteString(f.StoreKey)
	builder.WriteString(", ")
	builder.WriteString("category_type=")
	builder.WriteString(f.CategoryType)
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(f.URI)
	builder.WriteString(", ")
	builder.WriteString("storage_scheme=")
	builder.WriteString(f.StorageScheme)
	builder.WriteString(", ")
	builder.WriteString("storage_volume=")
	builder.WriteString(f.StorageVolume)
	builder.WriteString(", ")
	builder.WriteString("storage_path=")
	builder.WriteString(f.StoragePath)
	builder.WriteString(", ")
	builder.WriteString("file_contents=")
	builder.WriteString(fmt.Sprintf("%v", f.FileContents))
	builder.WriteByte(')')
	return builder.String()
}

// NamedUser returns the User named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedUser(name string) ([]*User, error) {
	if f.Edges.namedUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedUser(name string, edges ...*User) {
	if f.Edges.namedUser == nil {
		f.Edges.namedUser = make(map[string][]*User)
	}
	if len(edges) == 0 {
		f.Edges.namedUser[name] = []*User{}
	} else {
		f.Edges.namedUser[name] = append(f.Edges.namedUser[name], edges...)
	}
}

// NamedOrganization returns the Organization named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedOrganization(name string) ([]*Organization, error) {
	if f.Edges.namedOrganization == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedOrganization[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedOrganization(name string, edges ...*Organization) {
	if f.Edges.namedOrganization == nil {
		f.Edges.namedOrganization = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		f.Edges.namedOrganization[name] = []*Organization{}
	} else {
		f.Edges.namedOrganization[name] = append(f.Edges.namedOrganization[name], edges...)
	}
}

// NamedGroup returns the Group named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedGroup(name string) ([]*Group, error) {
	if f.Edges.namedGroup == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedGroup[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedGroup(name string, edges ...*Group) {
	if f.Edges.namedGroup == nil {
		f.Edges.namedGroup = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		f.Edges.namedGroup[name] = []*Group{}
	} else {
		f.Edges.namedGroup[name] = append(f.Edges.namedGroup[name], edges...)
	}
}

// NamedContact returns the Contact named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedContact(name string) ([]*Contact, error) {
	if f.Edges.namedContact == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedContact[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedContact(name string, edges ...*Contact) {
	if f.Edges.namedContact == nil {
		f.Edges.namedContact = make(map[string][]*Contact)
	}
	if len(edges) == 0 {
		f.Edges.namedContact[name] = []*Contact{}
	} else {
		f.Edges.namedContact[name] = append(f.Edges.namedContact[name], edges...)
	}
}

// NamedEntity returns the Entity named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedEntity(name string) ([]*Entity, error) {
	if f.Edges.namedEntity == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedEntity[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedEntity(name string, edges ...*Entity) {
	if f.Edges.namedEntity == nil {
		f.Edges.namedEntity = make(map[string][]*Entity)
	}
	if len(edges) == 0 {
		f.Edges.namedEntity[name] = []*Entity{}
	} else {
		f.Edges.namedEntity[name] = append(f.Edges.namedEntity[name], edges...)
	}
}

// NamedUsersetting returns the Usersetting named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedUsersetting(name string) ([]*UserSetting, error) {
	if f.Edges.namedUsersetting == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedUsersetting[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedUsersetting(name string, edges ...*UserSetting) {
	if f.Edges.namedUsersetting == nil {
		f.Edges.namedUsersetting = make(map[string][]*UserSetting)
	}
	if len(edges) == 0 {
		f.Edges.namedUsersetting[name] = []*UserSetting{}
	} else {
		f.Edges.namedUsersetting[name] = append(f.Edges.namedUsersetting[name], edges...)
	}
}

// NamedOrganizationsetting returns the Organizationsetting named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedOrganizationsetting(name string) ([]*OrganizationSetting, error) {
	if f.Edges.namedOrganizationsetting == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedOrganizationsetting[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedOrganizationsetting(name string, edges ...*OrganizationSetting) {
	if f.Edges.namedOrganizationsetting == nil {
		f.Edges.namedOrganizationsetting = make(map[string][]*OrganizationSetting)
	}
	if len(edges) == 0 {
		f.Edges.namedOrganizationsetting[name] = []*OrganizationSetting{}
	} else {
		f.Edges.namedOrganizationsetting[name] = append(f.Edges.namedOrganizationsetting[name], edges...)
	}
}

// NamedTemplate returns the Template named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedTemplate(name string) ([]*Template, error) {
	if f.Edges.namedTemplate == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedTemplate[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedTemplate(name string, edges ...*Template) {
	if f.Edges.namedTemplate == nil {
		f.Edges.namedTemplate = make(map[string][]*Template)
	}
	if len(edges) == 0 {
		f.Edges.namedTemplate[name] = []*Template{}
	} else {
		f.Edges.namedTemplate[name] = append(f.Edges.namedTemplate[name], edges...)
	}
}

// NamedDocumentdata returns the Documentdata named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedDocumentdata(name string) ([]*DocumentData, error) {
	if f.Edges.namedDocumentdata == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedDocumentdata[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedDocumentdata(name string, edges ...*DocumentData) {
	if f.Edges.namedDocumentdata == nil {
		f.Edges.namedDocumentdata = make(map[string][]*DocumentData)
	}
	if len(edges) == 0 {
		f.Edges.namedDocumentdata[name] = []*DocumentData{}
	} else {
		f.Edges.namedDocumentdata[name] = append(f.Edges.namedDocumentdata[name], edges...)
	}
}

// NamedEvents returns the Events named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedEvents(name string) ([]*Event, error) {
	if f.Edges.namedEvents == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedEvents[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedEvents(name string, edges ...*Event) {
	if f.Edges.namedEvents == nil {
		f.Edges.namedEvents = make(map[string][]*Event)
	}
	if len(edges) == 0 {
		f.Edges.namedEvents[name] = []*Event{}
	} else {
		f.Edges.namedEvents[name] = append(f.Edges.namedEvents[name], edges...)
	}
}

// NamedProgram returns the Program named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedProgram(name string) ([]*Program, error) {
	if f.Edges.namedProgram == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedProgram[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedProgram(name string, edges ...*Program) {
	if f.Edges.namedProgram == nil {
		f.Edges.namedProgram = make(map[string][]*Program)
	}
	if len(edges) == 0 {
		f.Edges.namedProgram[name] = []*Program{}
	} else {
		f.Edges.namedProgram[name] = append(f.Edges.namedProgram[name], edges...)
	}
}

// Files is a parsable slice of File.
type Files []*File
