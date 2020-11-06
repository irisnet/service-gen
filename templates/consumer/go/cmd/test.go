package main

import (
	"io/ioutil"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/test"
	"github.com/irisnet/service-gen/types"
	"github.com/spf13/cobra"
)

func TestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test",
		Short:   "test comsumer daemon",
		Example: `servicename-sp test`,
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// CreatClient
			testConfig := service.Config{
				ChainID:      types.TestChainID,
				NodeRPCAddr:  types.TestNodeRPCAddr,
				NodeGRPCAddr: types.TestNodeGRPCAddr,
				KeyPath:      types.TestKeyPath,
				KeyName:      types.TestKeyName,
				Fee:          types.TestFee,
				KeyAlgorithm: types.TestKeyAlgorithm,
			}
			serviceClient := service.MakeServiceClientWrapper(testConfig, types.TestPassword)

			// Import key
			keyArmor, err := ioutil.ReadFile("/home/sunny/data/node0/iritacli")
			if err != nil {
				return err
			}

			addr, err := serviceClient.ImportKey("node0", types.TestPassword, string(keyArmor))
			if err != nil {
				return err
			}

			// fmt.Printf("key imported successfully: %s\n", addr)

			// NewTestApp
			logger := common.Logger
			appInstance := app.NewTestApp(serviceClient, test.ResponseCallback, logger)

			// Define service
			err = appInstance.ServiceClient.DefineService()
			if err != nil {
				appInstance.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
				return err
			}

			// Bind service
			err = appInstance.ServiceClient.BindService()
			if err != nil {
				appInstance.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
				return err
			}

			// Subscribe request
			appInstance.ServiceClient.SubscribeServiceRequest(types.TestServiceName, test.RequestCallback)

			// Invoke service and subscribe response
			err = appInstance.ServiceClient.InvokeService(addr)
			if err != nil {
				appInstance.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
				return err
			}

			addr, err = serviceClient.ShowKey(serviceClient.KeyName, serviceClient.Password)
			if err != nil {
				return err
			}

			// Subscribe Service Response
			appInstance.Start(types.TestServiceName, test.ResponseCallback, addr)

			return nil
		},
	}

	return cmd
}
