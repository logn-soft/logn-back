// Code generated by ent, DO NOT EDIT.

package technology

import (
	"time"
)

const (
	// Label holds the string label denoting the technology type in the database.
	Label = "technology"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeVacancies holds the string denoting the vacancies edge name in mutations.
	EdgeVacancies = "vacancies"
	// EdgeTechnologyLevels holds the string denoting the technology_levels edge name in mutations.
	EdgeTechnologyLevels = "technology_levels"
	// Table holds the table name of the technology in the database.
	Table = "technologies"
	// VacanciesTable is the table that holds the vacancies relation/edge. The primary key declared below.
	VacanciesTable = "technology_levels"
	// VacanciesInverseTable is the table name for the Vacancy entity.
	// It exists in this package in order to avoid circular dependency with the "vacancy" package.
	VacanciesInverseTable = "vacancies"
	// TechnologyLevelsTable is the table that holds the technology_levels relation/edge.
	TechnologyLevelsTable = "technology_levels"
	// TechnologyLevelsInverseTable is the table name for the TechnologyLevel entity.
	// It exists in this package in order to avoid circular dependency with the "technologylevel" package.
	TechnologyLevelsInverseTable = "technology_levels"
	// TechnologyLevelsColumn is the table column denoting the technology_levels relation/edge.
	TechnologyLevelsColumn = "technology_id"
)

// Columns holds all SQL columns for technology fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
}

var (
	// VacanciesPrimaryKey and VacanciesColumn2 are the table columns denoting the
	// primary key for the vacancies relation (M2M).
	VacanciesPrimaryKey = []string{"technology_id", "vacancy_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
