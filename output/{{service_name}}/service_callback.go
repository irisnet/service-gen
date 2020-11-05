package {{service_name}}

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"gitlab.bianjie.ai/tianle/service-gen/types"
)

// ServiceCallback provider need to supplementary service logic
func ServiceCallback(reqID, input string) (output *types.ServiceOutput, requestResult *types.RequestResult) {
	serviceInput, err := ParseInput(input)
	if err != nil {
		requestResult = &types.RequestResult{
			State:   types.ClientError,
			Message: "failed to parse input",
		}
		return nil, requestResult
	}

	_ = serviceInput
	// Supplementary service logic...
	output = &types.ServiceOutput{}
	requestResult = &types.RequestResult{
		State:   types.Success,
		Message: "success",
	}
	return output, requestResult
}

// ParseInput ParseInput input
func ParseInput(input string) (serviceInput *types.ServiceInput, err error) {
	input = gjson.Get(input, "body").String()

	err = json.Unmarshal([]byte(input), &serviceInput)

	return serviceInput, err
}
