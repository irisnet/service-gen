package main

import (
	"strconv"

	servicesdk "github.com/irisnet/service-sdk-go/service"
	sdkTypes "github.com/irisnet/service-sdk-go/types"
	"github.com/spf13/cobra"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
)

func bindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bind",
		Short:   "bind service",
		Example: `{{service_name}}-sp bind [deposit] [pricing] [qos] [options] [provider] [config-file]`,
		Args:    cobra.RangeArgs(5, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			password := getPassword()

			bindConfig, configPath, err := parameterHandler(args)
			if err != nil {
				return err
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), password)

			logger := common.Logger

			appInstance := app.NewApp(serviceClient, {{service_name}}.RequestCallback, logger)
			appInstance.Bind(bindConfig)

			return nil
		},
	}

	return cmd
}

func parameterHandler(args []string) (servicesdk.BindServiceRequest, string, error) {
	i, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		panic(err)
	}
	deposit := sdkTypes.NewDecCoins(sdkTypes.NewDecCoin("stake", sdkTypes.NewInt(i)))

	j, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}
	qos := uint64(j)

	bindConfig := servicesdk.BindServiceRequest{
		ServiceName: types.ServiceName,
		Deposit:     deposit,
		Pricing:     args[1],
		QoS:         qos,
		Options:     args[3],
		Provider:    args[4],
	}

	configPath := common.ConfigPath

	if len(args) == 6 {
		configPath = args[5]
	}

	return bindConfig, configPath, nil
}
