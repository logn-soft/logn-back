// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/logn-soft/logn-back/internal/ent/vacancy"
)

// Vacancy is the model entity for the Vacancy schema.
type Vacancy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsNegotiate holds the value of the "is_negotiate" field.
	IsNegotiate bool `json:"is_negotiate,omitempty"`
	// MinSalary holds the value of the "min_salary" field.
	MinSalary int `json:"min_salary,omitempty"`
	// MaxSalary holds the value of the "max_salary" field.
	MaxSalary int `json:"max_salary,omitempty"`
	// IsRemote holds the value of the "is_remote" field.
	IsRemote bool `json:"is_remote,omitempty"`
	// Views holds the value of the "views" field.
	Views int `json:"views,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VacancyQuery when eager-loading is set.
	Edges VacancyEdges `json:"edges"`
}

// VacancyEdges holds the relations/edges for other nodes in the graph.
type VacancyEdges struct {
	// Technologies holds the value of the technologies edge.
	Technologies []*Technology `json:"technologies,omitempty"`
	// Locations holds the value of the locations edge.
	Locations []*Location `json:"locations,omitempty"`
	// Areas holds the value of the areas edge.
	Areas []*Area `json:"areas,omitempty"`
	// TechnologyLevels holds the value of the technology_levels edge.
	TechnologyLevels []*TechnologyLevel `json:"technology_levels,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TechnologiesOrErr returns the Technologies value or an error if the edge
// was not loaded in eager-loading.
func (e VacancyEdges) TechnologiesOrErr() ([]*Technology, error) {
	if e.loadedTypes[0] {
		return e.Technologies, nil
	}
	return nil, &NotLoadedError{edge: "technologies"}
}

// LocationsOrErr returns the Locations value or an error if the edge
// was not loaded in eager-loading.
func (e VacancyEdges) LocationsOrErr() ([]*Location, error) {
	if e.loadedTypes[1] {
		return e.Locations, nil
	}
	return nil, &NotLoadedError{edge: "locations"}
}

// AreasOrErr returns the Areas value or an error if the edge
// was not loaded in eager-loading.
func (e VacancyEdges) AreasOrErr() ([]*Area, error) {
	if e.loadedTypes[2] {
		return e.Areas, nil
	}
	return nil, &NotLoadedError{edge: "areas"}
}

// TechnologyLevelsOrErr returns the TechnologyLevels value or an error if the edge
// was not loaded in eager-loading.
func (e VacancyEdges) TechnologyLevelsOrErr() ([]*TechnologyLevel, error) {
	if e.loadedTypes[3] {
		return e.TechnologyLevels, nil
	}
	return nil, &NotLoadedError{edge: "technology_levels"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vacancy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vacancy.FieldIsNegotiate, vacancy.FieldIsRemote:
			values[i] = new(sql.NullBool)
		case vacancy.FieldID, vacancy.FieldMinSalary, vacancy.FieldMaxSalary, vacancy.FieldViews:
			values[i] = new(sql.NullInt64)
		case vacancy.FieldName, vacancy.FieldDescription:
			values[i] = new(sql.NullString)
		case vacancy.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vacancy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vacancy fields.
func (v *Vacancy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vacancy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vacancy.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case vacancy.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				v.Description = value.String
			}
		case vacancy.FieldIsNegotiate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_negotiate", values[i])
			} else if value.Valid {
				v.IsNegotiate = value.Bool
			}
		case vacancy.FieldMinSalary:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_salary", values[i])
			} else if value.Valid {
				v.MinSalary = int(value.Int64)
			}
		case vacancy.FieldMaxSalary:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_salary", values[i])
			} else if value.Valid {
				v.MaxSalary = int(value.Int64)
			}
		case vacancy.FieldIsRemote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_remote", values[i])
			} else if value.Valid {
				v.IsRemote = value.Bool
			}
		case vacancy.FieldViews:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field views", values[i])
			} else if value.Valid {
				v.Views = int(value.Int64)
			}
		case vacancy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTechnologies queries the "technologies" edge of the Vacancy entity.
func (v *Vacancy) QueryTechnologies() *TechnologyQuery {
	return NewVacancyClient(v.config).QueryTechnologies(v)
}

// QueryLocations queries the "locations" edge of the Vacancy entity.
func (v *Vacancy) QueryLocations() *LocationQuery {
	return NewVacancyClient(v.config).QueryLocations(v)
}

// QueryAreas queries the "areas" edge of the Vacancy entity.
func (v *Vacancy) QueryAreas() *AreaQuery {
	return NewVacancyClient(v.config).QueryAreas(v)
}

// QueryTechnologyLevels queries the "technology_levels" edge of the Vacancy entity.
func (v *Vacancy) QueryTechnologyLevels() *TechnologyLevelQuery {
	return NewVacancyClient(v.config).QueryTechnologyLevels(v)
}

// Update returns a builder for updating this Vacancy.
// Note that you need to call Vacancy.Unwrap() before calling this method if this Vacancy
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vacancy) Update() *VacancyUpdateOne {
	return NewVacancyClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vacancy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vacancy) Unwrap() *Vacancy {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vacancy is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vacancy) String() string {
	var builder strings.Builder
	builder.WriteString("Vacancy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(v.Description)
	builder.WriteString(", ")
	builder.WriteString("is_negotiate=")
	builder.WriteString(fmt.Sprintf("%v", v.IsNegotiate))
	builder.WriteString(", ")
	builder.WriteString("min_salary=")
	builder.WriteString(fmt.Sprintf("%v", v.MinSalary))
	builder.WriteString(", ")
	builder.WriteString("max_salary=")
	builder.WriteString(fmt.Sprintf("%v", v.MaxSalary))
	builder.WriteString(", ")
	builder.WriteString("is_remote=")
	builder.WriteString(fmt.Sprintf("%v", v.IsRemote))
	builder.WriteString(", ")
	builder.WriteString("views=")
	builder.WriteString(fmt.Sprintf("%v", v.Views))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Vacancies is a parsable slice of Vacancy.
type Vacancies []*Vacancy

func (v Vacancies) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}