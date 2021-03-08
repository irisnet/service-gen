package types

import "regexp"

const (
	ServiceName = "{{service_name}}"
)

const (
	DefaultTimeout   = 100
	DefaultRepeated  = false
	DefaultFrequency = 0
	DefaultTotal     = 1
)

var FeeReg = regexp.MustCompile("[A-Za-z][A-Za-z0-9/]{2,127}")

type ResponseCallback func(reqCtxID, reqID, output string)
