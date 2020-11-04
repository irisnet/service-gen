package types

import (
	"encoding/json"
)

const (
	ServiceName = "{{service_name}}"
)

type State int

// Status returned of serviceCallback
const (
	Success = iota
	ClientError
	ServiceError
)

// RequestResult is result of serviceCallback
type RequestResult struct {
	State   State // Use status returned
	Message string
}

// Result of serviceCallback
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Marshal output
func (output *ServiceOutput) Marshal() ([]byte, error) {
	data, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	return data, nil
}
