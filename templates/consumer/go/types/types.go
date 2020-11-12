package types

const (
	ServiceName = "{{service_name}}"
)

const (
	DefaultTimeout   = 100
	DefaultRepeated  = false
	DefaultFrequency = 110
	DefaultTotal     = 1
)

type ResponseCallback func(reqCtxID, reqID, output string)
