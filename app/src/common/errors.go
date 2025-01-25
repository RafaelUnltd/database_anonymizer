package common

import (
	"fmt"
)

func ErrTableNotFound(tableName string) error {
	return fmt.Errorf("table not found: '%s'", tableName)
}

func ErrColumnNotFound(tableName string, columnName string) error {
	return fmt.Errorf("column not found for table '%s': '%s'", tableName, columnName)
}
