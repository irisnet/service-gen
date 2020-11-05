package service

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/tianle/service-gen/types"
)

type ServiceCallback func(reqID, input string) (output *types.ServiceOutput, requestResult *types.RequestResult)

// CallbackHandler Servicecallback processing function
var CallbackHandler = func(reqID, input string, serviceCb ServiceCallback, logger *log.Logger) (response, result string) {
	// Receiving serviceCallback processing results
	serviceOutput, requestResult := serviceCb(reqID, input)

	// Convert the requestresult to the corresponding error code
	res := resultConvert(requestResult)

	return MarshalResAndOutput(res, serviceOutput, logger)
}

// Convert the requestresult to the corresponding error code
func resultConvert(requestResult *types.RequestResult) *types.Result {
	res := types.Result{}
	if requestResult == nil {
		res.Code = 500
		res.Message = "The response result is empty."
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

func MarshalResAndOutput(res *types.Result, serviceOutput *types.ServiceOutput, logger *log.Logger) (response, result string) {
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

	logger.Infof("request processed, result: %s, response: %s", result, response)

	return response, result
}
