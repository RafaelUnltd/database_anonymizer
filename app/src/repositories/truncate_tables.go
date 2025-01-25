package repositories

func (r Repository) TruncateTable(tableName string) error {
	return r.db.Exec("DELETE FROM " + tableName).Error
}
