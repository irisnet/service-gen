package service

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/tianle/servicegen/types"
)

type ServiceCallback func(reqCtxID, input string) (output *types.ServiceOutput, requestResult *types.RequestResult)

// CbHandler Servicecallback processing function
var CbHandler = func(reqID, input string, serviceCb ServiceCallback, logger *log.Logger) (output, result string) {
	serviceOutput := &types.ServiceOutput{}
	res := &types.Result{
		Code: 200,
	}

	defer func() {
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
			output = fmt.Sprintf(`{"header":{},"body":%s}`, outputBz)
		}

		logger.Infof("request processed, result: %s, output: %s", result, output)
	}()

	// Receiving serviceCallback processing results
	serviceOutput, requestResult := serviceCb(reqID, input)
	// 转换为对应错误码
	res = resultConvert(requestResult)

	return output, result
}

// Convert the requestresult to the corresponding error code
func resultConvert(requestResult *types.RequestResult) *types.Result {
	res := types.Result{Message: requestResult.Message}
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
