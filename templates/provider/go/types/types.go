package types

const (
	{{service_name}} = "{{service_name}}"
)

type ServiceCallback func(reqID, input string) (output *ServiceOutput, requestResult *RequestResult)

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

type ServiceInput struct{}
type ServiceOutput struct{}
