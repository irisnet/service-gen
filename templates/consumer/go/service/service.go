package service

import (
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	sdkTypes "github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
	log "github.com/sirupsen/logrus"

	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/types"
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

// InvokeService wraps service.InvokeService
func (s ServiceClientWrapper) InvokeService(invokeConfig service.InvokeServiceRequest) (string, string, error) {
	reqCtxID, _, err := s.ServiceClient.InvokeService(invokeConfig, s.buildBaseTx())
	if err != nil {
		return "", "", err
	}
	QueryServiceRequestResponse, err := s.ServiceClient.QueryRequestsByReqCtx(reqCtxID, 1)
	if err != nil {
		return "", "", err
	}
	reqID := QueryServiceRequestResponse[0].ID
	return reqCtxID, reqID, nil
}

// SubscribeServiceResponse wraps service.SubscribeServiceResponse
func (s ServiceClientWrapper) SubscribeServiceResponse(
	reqCtxID, reqID,
	consumerAddr string,
	responseCallback types.ResponseCallback,
) error {
	builder := createFilter(reqID, consumerAddr)

	callback := func(txs sdkTypes.EventDataTx) {
		for _, v := range txs.Result.Events {
			if v.GetType() == "service_slash" {
				common.Logger.Info("Illegal event detected")
				return
			}
		}

		serviceResponseResponse, err := s.ServiceClient.QueryServiceResponse(reqID)
		if err != nil {
			common.Logger.Info("fail to query output", err)
			return
		}

		responseCallback(reqCtxID, reqID, serviceResponseResponse.Output)
	}

	_, err := s.ServiceClient.SubscribeTx(builder, callback)
	return err
}

// buildBaseTx builds a base tx
func (s ServiceClientWrapper) buildBaseTx() sdkTypes.BaseTx {
	return sdkTypes.BaseTx{
		From:     s.KeyName,
		Password: s.Password,
	}
}

func createFilter(reqID string, consumerAddr string) (builder *sdkTypes.EventQueryBuilder) {
	return sdkTypes.NewEventQueryBuilder().AddCondition(
		sdkTypes.NewCond(
			sdkTypes.EventTypeMessage,
			"action",
		).EQ(
			sdkTypes.EventValue("respond_service"),
		),
	).AddCondition(
		sdkTypes.NewCond(
			sdkTypes.EventTypeResponseService,
			"request_id",
		).EQ(
			sdkTypes.EventValue(reqID),
		),
	)
}
