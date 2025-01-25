package repositories

import "gorm.io/gorm"

func (r Repository) Insert(transaction *gorm.DB, tableName string, values []map[string]interface{}) error {
	database := r.db
	if transaction != nil {
		database = transaction
	}

	err := database.Table(tableName).
		Create(values).
		Error

	if err != nil {
		return err
	}

	return nil
}
