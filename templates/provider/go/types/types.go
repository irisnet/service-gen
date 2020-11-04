package types

import (
	"encoding/json"
)

const (
	ServiceName = "{{service_name}}"
)

type State int

// 状态变量
const (
	Success = iota
	ClientError
	ServiceError
)

// RequestResult 业务逻辑结果
type RequestResult struct {
	State   State // 使用状态变量
	Message string
}

// Result 逻辑函数处理结果
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Serialize 序列化 output
func (output *ServiceOutput) Marshal() ([]byte, error) {
	data, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	return data, nil
}
