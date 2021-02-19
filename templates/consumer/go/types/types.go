package types

const (
	ServiceName = "{{service_name}}"
)

const (
	DefaultTimeout   = 100
	DefaultRepeated  = false
	DefaultFrequency = 0
	DefaultTotal     = 1
)

const feeRe = "^(\\d+(?:\\.\\d+)?|\\.\\d+)([A-Za-z][A-Za-z0-9/]{2,127])$"

type ResponseCallback func(reqCtxID, reqID, output string)
