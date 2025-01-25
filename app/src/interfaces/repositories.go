package interfaces

import (
	"database_anonymizer/app/src/structs"

	"gorm.io/gorm"
)

type RepositoriesInterface interface {
	GetTables() ([]string, error)
	GetTableColumns(tableName string) ([]string, error)
	GetTableRecords(tableName string, limit int, page int) ([]map[string]interface{}, error)
	CountTableRecords(tableName string) (int64, error)
	UpdateTableRecord(tableName string, values map[string]interface{}, identifier string) error
	TruncateTable(tableName string) error
	SetReplicationRole(role structs.ReplicationRole) error
	Insert(transaction *gorm.DB, tableName string, values []map[string]interface{}) error
	BeginTransaction() *gorm.DB
}
