package repositories

import "database_anonymizer/app/src/structs"

func (r Repository) SetReplicationRole(role structs.ReplicationRole) error {
	return r.db.Exec("SET session_replication_role = '" + string(role) + "'").Error
}
