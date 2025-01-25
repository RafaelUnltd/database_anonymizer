package repositories

func (r Repository) CountTableRecords(tableName string) (int64, error) {
	var total int64

	err := r.db.Table(tableName).
		Count(&total).
		Error

	if err != nil {
		return 0, err
	}

	return total, nil
}
