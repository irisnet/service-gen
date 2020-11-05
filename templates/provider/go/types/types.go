package types

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

type Response struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}
