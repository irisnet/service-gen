package {{service_name}}

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/irisnet/service-gen/types"
)

// RequestCallback provider need to supplementary service logic
func RequestCallback(reqID, input string) (
	output *types.ServiceOutput,
	requestResult *types.RequestResult,
) {
	serviceInput, err := parseInput(input)
	if err != nil {
		requestResult = &types.RequestResult{
			State:   types.ClientError,
			Message: "failed to parse input",
		}
		return nil, requestResult
	}
	// Supplementary service logic...

	return output, requestResult
}

func parseInput(input string) (serviceInput *types.ServiceInput, err error) {
	input = gjson.Get(input, "body").String()

	err = json.Unmarshal([]byte(input), &serviceInput)

	return serviceInput, err
}
