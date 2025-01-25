package services

import (
	"context"
	"database_anonymizer/app/src/libs/anonymizer"
	"database_anonymizer/app/src/structs"
	"fmt"
	"math"
)

const BATCH_SIZE = 500

func (s Service) Anonymize(ctx context.Context, request structs.AnonymizationRequest, cacheKey string) error {
	tx := s.outputRepository.BeginTransaction()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	if err := tx.Exec("SET session_replication_role = 'replica'").Error; err != nil {
		return fmt.Errorf("erro ao configurar session_replication_role: %v", err)
	}

	for tableName := range request.AnonymizationRules {
		uniqueMap := make(anonymizer.UniqueAttributes)

		totalRecords, err := s.inputRepository.CountTableRecords(tableName)
		if err != nil {
			return err
		}

		fmt.Printf("Anonymizing %s table, total records: %d\n", tableName, totalRecords)

		var attributes []anonymizer.Attribute
		for columnName, columnOptions := range request.AnonymizationRules[tableName].Columns {
			attributes = append(attributes, anonymizer.Attribute{
				Name:   columnName,
				Method: columnOptions.Type,
				Mask:   columnOptions.Value,
				Unique: columnOptions.Unique,
			})
		}

		lastPage := int(math.Ceil(float64(totalRecords) / float64(BATCH_SIZE)))
		for i := 1; i <= lastPage; i++ {
			fmt.Printf("Anonymizing %s table, page %d of %d\n", tableName, i, lastPage)

			records, err := s.inputRepository.GetTableRecords(tableName, BATCH_SIZE, i)
			if err != nil {
				return err
			}

			var anonymizedRecords []map[string]interface{}
			for _, record := range records {
				err := anonymizer.AnonymizeRecord(&record, attributes, &uniqueMap)
				if err != nil {
					return fmt.Errorf("erro ao anonimizar registro: %v", err)
				}
				anonymizedRecords = append(anonymizedRecords, record)
			}

			fmt.Printf("Batch %d: Tentando inserir %d registros\n", i, len(anonymizedRecords))

			err = s.outputRepository.Insert(tx, tableName, anonymizedRecords)
			if err != nil {
				fmt.Printf("Erro ao inserir batch %d: %v\n", i, err)
				return fmt.Errorf("erro ao inserir registros anonimizados: %v", err)
			}

			err = s.updatePollingTableStatus(ctx, cacheKey, tableName, int(totalRecords), len(anonymizedRecords))
			if err != nil {
				return fmt.Errorf("erro ao atualizar status de polling: %v", err)
			}
		}
	}

	if err := tx.Exec("SET session_replication_role = 'origin'").Error; err != nil {
		return fmt.Errorf("erro ao restaurar session_replication_role: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("erro ao fazer commit: %v", err)
	}

	return nil
}

func (s Service) updatePollingTableStatus(ctx context.Context, cacheKey string, tableName string, totalRecords int, processed int) error {
	pollingStatus, err := s.cache.ReadPollingStatus(ctx, cacheKey)
	if err != nil {
		return err
	}

	totalProcessed := processed
	if _, ok := pollingStatus.Progress[tableName]; ok {
		totalProcessed = pollingStatus.Progress[tableName].ProcessedRecords + processed
	}

	pollingStatus.Progress[tableName] = structs.TableStatus{
		TotalRecords:     totalRecords,
		ProcessedRecords: totalProcessed,
	}

	err = s.cache.UpdatePollingStatus(ctx, cacheKey, pollingStatus)
	if err != nil {
		return err
	}

	return nil
}
