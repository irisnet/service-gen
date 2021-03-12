package {{service_name}}

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/types"
)

// ResponseCallback provider need to supplementary service logic
func ResponseCallback(reqCtxID, reqID, output string) {
	serviceOutput := parseOutput(output)
	outputBytes, _ := json.Marshal(serviceOutput.Output)
	common.Logger.Info("Get response: \n", string(outputBytes))
	// Supplementary service logic...

}

func parseOutput(output string) (serviceOutput *types.ServiceOutput) {
	serviceOutput = &types.ServiceOutput{}
	output = gjson.Get(output, "body").String()
	if output == "" {
		return serviceOutput
	}
	err := json.Unmarshal([]byte(output), serviceOutput)
	if err != nil {
		panic(err)
	}

	return serviceOutput
}
