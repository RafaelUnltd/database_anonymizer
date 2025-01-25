package structs

type Status string

var (
	StatusPending    Status = "PENDING"
	StatusProcessing Status = "PROCESSING"
	StatusFinished   Status = "FINISHED"
	StatusError      Status = "ERROR"
)

type PollingStatus struct {
	Key      string                 `json:"key"`
	Status   Status                 `json:"status"`
	Finished bool                   `json:"finished"`
	Progress map[string]TableStatus `json:"progress"`
}

type TableStatus struct {
	TotalRecords     int `json:"total_records"`
	ProcessedRecords int `json:"processed_records"`
}
