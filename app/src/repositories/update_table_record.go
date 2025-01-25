package repositories

func (r Repository) UpdateTableRecord(tableName string, values map[string]interface{}, identifier string) error {
	err := r.db.Table(tableName).
		Where(identifier+" = ?", values[identifier]).
		Updates(values).
		Error

	if err != nil {
		return err
	}

	return nil
}
