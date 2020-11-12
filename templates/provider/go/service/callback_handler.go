package service

import (
	"encoding/json"

	"github.com/irisnet/service-gen/types"
	log "github.com/sirupsen/logrus"
)

// CallbackHandler is processing function of RequestCallback
var CallbackHandler = func(
	reqID string,
	input string,
	requestCallback types.RequestCallback,
	logger *log.Logger,
) (string, string) {
	// Receiving processing results of RequestCallback
	serviceOutput, requestResult := requestCallback(reqID, input)

	// Convert the requestResult to the corresponding error code
	res := convertRequestResult(requestResult)

	response, result := buildResAndOutput(res, serviceOutput)
	logger.Infof("request processed, result: %s, response: %s", result, response)
	return response, result
}

// Convert the requestresult to the corresponding error code
func convertRequestResult(requestResult *types.RequestResult) *types.Result {
	res := types.Result{}
	if requestResult == nil {
		res.Code = 500
		res.Message = "RequestResult is empty."
		return &res
	}

	res = types.Result{Message: requestResult.Message}
	switch requestResult.State {
	case types.Success:
		res.Code = 200
	case types.ClientError:
		res.Code = 400
	case types.ServiceError:
		res.Code = 500
	default:
		res.Code = 500
	}
	return &res
}

func buildResAndOutput(
	res *types.Result,
	serviceOutput *types.ServiceOutput,
) (response, result string) {
	resBz, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	result = string(resBz)

	if res.Code == 200 {
		outputBz, err := json.Marshal(serviceOutput)
		if err != nil {
			panic(err)
		}
		output := types.Response{
			Header: "",
			Body:   string(outputBz),
		}
		responseBz, err := json.Marshal(output)
		if err != nil {
			panic(err)
		}
		response = string(responseBz)
	}

	return response, result
}
