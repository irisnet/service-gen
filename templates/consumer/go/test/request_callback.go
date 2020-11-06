package test

import (
	"encoding/json"

	"github.com/irisnet/service-gen/types"
	"github.com/tidwall/gjson"
)

// RequestCallback provider need to supplementary service logic
func RequestCallback(reqID, input string) (output *types.TestServiceOutput, requestResult *types.RequestResult) {
	serviceInput, err := ParseInput(input)
	if err != nil {
		requestResult = &types.RequestResult{
			State:   types.ClientError,
			Message: "failed to parse input",
		}
		return nil, requestResult
	}

	_ = serviceInput
	output = &types.TestServiceOutput{
		Output: "wrold",
	}
	requestResult = &types.RequestResult{
		State:   types.Success,
		Message: "success",
	}
	return output, requestResult
}

// ParseInput ParseInput input
func ParseInput(input string) (serviceInput *types.TestServiceInput, err error) {
	input = gjson.Get(input, "body").String()

	err = json.Unmarshal([]byte(input), &serviceInput)

	return serviceInput, err
}
