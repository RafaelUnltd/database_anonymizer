package services

import (
	"database_anonymizer/app/src/common"
	"database_anonymizer/app/src/structs"

	"github.com/samber/lo"
)

func (s Service) ValidateRules(request structs.AnonymizationRequest) error {
	tables, err := s.inputRepository.GetTables()
	if err != nil {
		return err
	}

	for tableName := range request.AnonymizationRules {
		if !lo.Contains(tables, tableName) {
			return common.ErrTableNotFound(tableName)
		}

		columns, err := s.inputRepository.GetTableColumns(tableName)
		if err != nil {
			return err
		}

		for columnName := range request.AnonymizationRules[tableName].Columns {
			if !lo.Contains(columns, columnName) {
				return common.ErrColumnNotFound(tableName, columnName)
			}
		}
	}

	return nil
}
