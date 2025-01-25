package structs

import (
	"database_anonymizer/app/src/libs/anonymizer"
	"fmt"
)

type AnonymizationRequest struct {
	InputConnectionInfo  DatabaseConnectionInfo `json:"input_connection_info"`
	OutputConnectionInfo DatabaseConnectionInfo `json:"output_connection_info"`
	AnonymizationRules   Tables                 `json:"anonymization_rules"`
}

func (r *AnonymizationRequest) TableNames() []string {
	var tableNames []string = []string{}
	for tableName := range r.AnonymizationRules {
		tableNames = append(tableNames, tableName)
	}
	return tableNames
}

type DatabaseConnectionInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (r *DatabaseConnectionInfo) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.Host,
		r.Port,
		r.User,
		r.Password,
		r.Database,
	)
}

func (r *DatabaseConnectionInfo) DumpString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		r.User,
		r.Password,
		r.Host,
		r.Port,
		r.Database,
	)
}

type Tables map[string]Table

type Table struct {
	Identifier string       `json:"identifier"`
	Columns    TableColumns `json:"columns"`
}

type TableColumns map[string]ColumnOptions

type ColumnOptions struct {
	Type   anonymizer.AnonimizationType `json:"type"`
	Value  string                       `json:"value"`
	Unique bool                         `json:"unique"`
}
