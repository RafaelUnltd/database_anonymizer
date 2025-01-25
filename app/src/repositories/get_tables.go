package repositories

func (r Repository) GetTables() ([]string, error) {
	var tables []string

	err := r.db.Table("information_schema.tables").
		Where("table_schema = ?", "public").
		Pluck("table_name", &tables).
		Error

	if err != nil {
		return []string{}, err
	}

	return tables, nil
}
