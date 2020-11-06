package types

const (
	ServiceName = "servicename"
)

type ResponseCallback func(reqID, output string)

type ServiceInput struct{}
type ServiceOutput struct{}
