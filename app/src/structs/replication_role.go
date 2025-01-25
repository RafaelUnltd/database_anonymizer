package structs

type ReplicationRole string

var (
	REPLICA ReplicationRole = "replica"
	ORIGIN  ReplicationRole = "origin"
)
