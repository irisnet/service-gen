package {{service_name}}

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/irisnet/service-gen/types"
	"github.com/irisnet/service-gen/common"
)

// RequestCallback provider need to supplementary service logic
func RequestCallback(reqID, input string) (
	output *types.ServiceOutput,
	requestResult *types.RequestResult,
) {
	output = &types.ServiceOutput{}
	requestResult = &types.RequestResult{}
	serviceInput, err := parseInput(input)
	if err != nil {
		requestResult.State = types.ClientError
		requestResult.Message = "failed to parse input"
		return nil, requestResult
	}
	common.Logger.Info("Get request: \n", serviceInput)
	// Supplementary service logic...

	return output, requestResult
}

func parseInput(input string) (serviceInput *types.ServiceInput, err error) {
	serviceInput = &types.ServiceInput{}
	input = gjson.Get(input, "body").String()
	if len(input) == 0 {
		return serviceInput, nil
	}
	err = json.Unmarshal([]byte(input), serviceInput)

	return serviceInput, err
}
