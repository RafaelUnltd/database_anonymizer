package repositories

func (r Repository) GetTableColumns(tableName string) ([]string, error) {
	var columns []string

	err := r.db.Table("information_schema.columns").
		Where("table_schema = ?", "public").
		Where("table_name = ?", tableName).
		Pluck("column_name", &columns).
		Error

	if err != nil {
		return []string{}, err
	}

	return columns, nil
}
