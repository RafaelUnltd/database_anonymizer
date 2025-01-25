package structs

type HttpErrorResponse struct {
	Tag     interface{} `json:"tag"`
	Message string      `json:"mesage"`
}

type HttpDataResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}
