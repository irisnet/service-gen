package service

import (
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/test"
	"github.com/irisnet/service-gen/types"
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	sdkTypes "github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
)

// ServiceClientWrapper defines a wrapper for service client
type ServiceClientWrapper struct {
	ChainID      string
	NodeRPCAddr  string
	NodeGRPCAddr string

	KeyPath  string
	KeyName  string
	Password string

	Logger        *log.Logger
	ServiceClient servicesdk.ServiceClient
}

// NewServiceClientWrapper constructs a new ServiceClientWrapper
func NewServiceClientWrapper(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
	keyPath string,
	keyName string,
	password string,
	feeConfig string,
	keyAlgorithm string,
	logger *log.Logger,
) ServiceClientWrapper {
	if len(chainID) == 0 {
		chainID = defaultChainID
	}

	if len(nodeRPCAddr) == 0 {
		nodeRPCAddr = defaultNodeRPCAddr
	}

	if len(nodeGRPCAddr) == 0 {
		nodeGRPCAddr = defaultNodeGRPCAddr
	}

	if len(keyPath) == 0 {
		keyPath = defaultKeyPath
	}

	if len(feeConfig) == 0 {
		feeConfig = defaultFee
	}
	fee, err := sdkTypes.ParseDecCoins(feeConfig)
	if err != nil {
		panic(err)
	}

	if len(keyAlgorithm) == 0 {
		keyAlgorithm = defaultKeyAlgorithm
	}

	config := sdkTypes.ClientConfig{
		NodeURI:  nodeRPCAddr,
		GRPCAddr: nodeGRPCAddr,
		ChainID:  chainID,
		Gas:      defaultGas,
		Fee:      fee,
		Mode:     defaultBroadcastMode,
		Algo:     keyAlgorithm,
		KeyDAO:   store.NewFileDAO(keyPath),
	}

	wrapper := ServiceClientWrapper{
		ChainID:       chainID,
		NodeRPCAddr:   nodeRPCAddr,
		NodeGRPCAddr:  nodeGRPCAddr,
		KeyPath:       keyPath,
		KeyName:       keyName,
		Password:      password,
		Logger:        logger,
		ServiceClient: servicesdk.NewServiceClient(config),
	}

	return wrapper
}

// MakeServiceClientWrapper builds a ServiceClientWrapper from the given config
func MakeServiceClientWrapper(config Config, password string) ServiceClientWrapper {
	return NewServiceClientWrapper(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.KeyPath,
		config.KeyName,
		password,
		config.Fee,
		config.KeyAlgorithm,
		common.Logger,
	)
}

// DefineService wraps service.DefineService
func (s ServiceClientWrapper) DefineService() error {
	defineServiceReq := service.DefineServiceRequest{
		ServiceName:       types.TestServiceName,
		Description:       types.TestServiceDescription,
		Tags:              types.TestServiceTags,
		AuthorDescription: types.TestAuthorDescription,
		Schemas:           types.TestSchemas,
	}

	_, err := s.ServiceClient.DefineService(defineServiceReq, s.BuildBaseTx())
	return err
}

// BindService wraps service.BindService
func (s ServiceClientWrapper) BindService() error {
	bindServiceReq := service.BindServiceRequest{
		ServiceName: types.TestServiceName,
		Deposit:     types.TestDeposit,
		Pricing:     types.TestPricing,
		QoS:         types.TestQos,
	}

	_, err := s.ServiceClient.BindService(bindServiceReq, s.BuildBaseTx())
	return err
}

// SubscribeServiceRequest wraps service.SubscribeServiceRequest
func (s ServiceClientWrapper) SubscribeServiceRequest(serviceName string, requestCb types.RequestCallback) error {
	callback := func(reqCtxID, reqID, input string) (output, result string) {
		return test.CallbackHandler(reqID, input, requestCb, s.Logger)
	}
	_, err := s.ServiceClient.SubscribeServiceRequest(serviceName, callback, s.BuildBaseTx())
	return err
}

// InvokeService wraps service.InvokeService
func (s ServiceClientWrapper) InvokeService(provider string) error {
	invokeServiceReq := service.InvokeServiceRequest{
		ServiceName:       types.TestServiceName,
		Providers:         []string{provider},
		ServiceFeeCap:     types.TestServiceFeeCap,
		Input:             types.TestInput,
		Timeout:           types.TestTimeout,
		Repeated:          types.TestRepeated,
		RepeatedFrequency: types.TestFrequency,
		RepeatedTotal:     types.TestTotal,
	}

	_, err := s.ServiceClient.InvokeService(invokeServiceReq, s.BuildBaseTx())
	return err
}

// SubscribeServiceResponse wraps service.SubscribeServiceResponse
func (s ServiceClientWrapper) SubscribeServiceResponse(serviceName string, responseCallback types.ResponseCallback, consumerAddr string) error {
	builder := sdkTypes.NewEventQueryBuilder().AddCondition(
		sdkTypes.NewCond(
			sdkTypes.EventTypeMessage,
			"service_name",
		).EQ(
			sdkTypes.EventValue(serviceName),
		),
	).AddCondition(
		sdkTypes.NewCond(
			sdkTypes.EventTypeMessage,
			"consumer",
		).EQ(
			sdkTypes.EventValue(consumerAddr),
		),
	)

	callback := func(txs sdkTypes.EventDataTx) {
		events := txs.Result.Events
		println(events)
		// events.GetValue("","reqctxid")
		// responseCallback(reqID, output)
	}

	_, err := s.ServiceClient.SubscribeTx(builder, callback)
	return err
}

// BuildBaseTx builds a base tx
func (s ServiceClientWrapper) BuildBaseTx() sdkTypes.BaseTx {
	return sdkTypes.BaseTx{
		From:     s.KeyName,
		Password: s.Password,
	}
}
