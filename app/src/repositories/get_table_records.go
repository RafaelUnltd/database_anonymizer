package repositories

func (r Repository) GetTableRecords(tableName string, limit int, page int) ([]map[string]interface{}, error) {
	var records []map[string]interface{}

	err := r.db.Table(tableName).
		Limit(limit).
		Offset(limit * (page - 1)).
		Order("id ASC").
		Find(&records).
		Error

	if err != nil {
		return []map[string]interface{}{}, err
	}

	return records, nil
}
