package

import "encoding/json"
{{service_name}}

import (
	"encoding/json"

	"github.com/tidwall/gjson"

	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/types"
)

// ResponseCallback provider need to supplementary service logic
func ResponseCallback(reqCtxID, reqID, output string) {
	common.Logger.Infof("Get response: \n", output)
	serviceOutput := parseOutput(output)
	// Supplementary service logic...

}

func parseOutput(output string) (serviceOutput *types.ServiceOutput) {
	output = gjson.Get(output, "body").String()
	serviceOutput = &types.ServiceOutput
	if output == "" {
		return serviceOutput
	}
	err := json.Unmarshal([]byte(output), serviceOutput)
	if err != nil {
		panic(err)
	}

	return serviceOutput
}
