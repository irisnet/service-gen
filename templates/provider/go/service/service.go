package service

import (
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/types"
	servicesdk "github.com/irisnet/service-sdk-go"
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

// SubscribeServiceRequest wraps service.SubscribeServiceRequest
func (s ServiceClientWrapper) SubscribeServiceRequest(serviceName string, RequestCb types.RequestCallback) error {
	callback := func(reqCtxID, reqID, input string) (output, result string) {
		return CallbackHandler(reqID, input, RequestCb, s.Logger)
	}
	_, err := s.ServiceClient.SubscribeServiceRequest(serviceName, callback, s.BuildBaseTx())
	return err
}

// BuildBaseTx builds a base tx
func (s ServiceClientWrapper) BuildBaseTx() sdkTypes.BaseTx {
	return sdkTypes.BaseTx{
		From:     s.KeyName,
		Password: s.Password,
	}
}
