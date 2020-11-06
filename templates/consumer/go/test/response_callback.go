package test

import (
	"encoding/json"
	"fmt"

	"github.com/irisnet/service-gen/types"
	"github.com/tidwall/gjson"
)

// ResponseCallback provider need to supplementary service logic
func ResponseCallback(reqID, output string) {
	serviceOutput := parseOutput(output)
	fmt.Println(serviceOutput.Output)
}

// parseOutput
func parseOutput(output string) (serviceOutput *types.TestServiceOutput) {
	output = gjson.Get(output, "body").String()

	err := json.Unmarshal([]byte(output), &serviceOutput)
	if err != nil {
		panic(err)
	}

	return serviceOutput
}
