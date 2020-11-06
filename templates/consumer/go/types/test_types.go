package types

import "github.com/irisnet/service-sdk-go/types"

type RequestCallback func(reqID, input string) (*TestServiceOutput, *RequestResult)

var (
	TestServiceName        = "test"
	TestServiceDescription = "test service"
	TestServiceTags        = []string{"test"}
	TestAuthorDescription  = "irita test"
	TestSchemas            = `{"input":{"type":"object","properties":{"inpupt":{"type":"string"}}},"output":{"type":"object","properties":{"output":{"type":"string"}}}}`
)

const (
	TestChainID      = "test"
	TestNodeRPCAddr  = "http://localhost:26657"
	TestNodeGRPCAddr = "localhost:9090"
	// TODO
	TestKeyPath      = ""
	TestKeyName      = "test"
	TestFee          = "4point"
	TestKeyAlgorithm = "sm2"
)

var (
	TestDeposit = types.NewDecCoins(types.NewDecCoin("point", types.NewInt(10000)))
	TestPricing = `{"price":"1point"}`
	TestQos     = uint64(50)
)

var (
	TestServiceFeeCap = types.NewDecCoins(types.NewDecCoin("point", types.NewInt(10)))
	TestInput         = `{"input":"hello"}`
	TestTimeout       = int64(100)
	TestRepeated      = true
	TestFrequency     = uint64(100)
	TestTotal         = int64(1)
)

const (
	TestPassword = "12345678"
)

type State int

// Status returned of RequestCallback
const (
	Success = iota
	ClientError
	ServiceError
)

// RequestResult is result of RequestCallback
type RequestResult struct {
	State   State // Use status returned
	Message string
}

// Result of RequestCallback
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

type TestServiceInput struct {
	Input string `json:"input"`
}
type TestServiceOutput struct {
	Output string `json:"output"`
}
