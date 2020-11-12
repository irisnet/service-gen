package main

import (
	"strconv"
	"strings"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
	servicesdk "github.com/irisnet/service-sdk-go/service"
	sdkTypes "github.com/irisnet/service-sdk-go/types"
	"github.com/spf13/cobra"
)

func invokeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "invoke",
		Short:   "Invoke service",
		Example: `{{service_name}}-sc invoke [provider-list] [fee-cap] [input] [timeout] [repeated] [frequency] [total] [config-file]`,
		Args:    cobra.RangeArgs(3, 8),
		RunE: func(cmd *cobra.Command, args []string) error {
			password := getPassword()

			invokeConfig, configPath, err := parameterHandler(args)
			if err != nil {
				return err
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), password)

			logger := common.Logger

			appInstance := app.NewApp(serviceClient, {{service_name}}.ResponseCallback, logger)
			appInstance.Invoke(invokeConfig)

			return nil
		},
	}

	return cmd
}

func parameterHandler(args []string) (servicesdk.InvokeServiceRequest, string, error) {
	providerList := strings.Split(args[0], "/")

	feeCap, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		panic(err)
	}
	serviceFeeCap := sdkTypes.NewDecCoins(sdkTypes.NewDecCoin("point", sdkTypes.NewInt(feeCap)))

	input := args[2]

	invokeConfig := servicesdk.InvokeServiceRequest{
		ServiceName:       types.ServiceName,
		Providers:         providerList,
		ServiceFeeCap:     serviceFeeCap,
		Input:             input,
		Timeout:           types.DefaultTimeout,
		Repeated:          types.DefaultRepeated,
		RepeatedFrequency: types.DefaultFrequency,
		RepeatedTotal:     types.DefaultTotal,
	}

	configPath := common.ConfigPath

	if len(args) == 4 || len(args) == 5 || len(args) == 6 || len(args) == 7 || len(args) == 8 {
		invokeConfig.Timeout, err = strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	if len(args) == 5 || len(args) == 6 || len(args) == 7 || len(args) == 8 {
		invokeConfig.Repeated, err = strconv.ParseBool(args[4])
		if err != nil {
			panic(err)
		}
	}

	if len(args) == 6 || len(args) == 7 || len(args) == 8 {
		repeatedFrequency, err := strconv.Atoi(args[5])
		if err != nil {
			panic(err)
		}
		invokeConfig.RepeatedFrequency = uint64(repeatedFrequency)
	}

	if len(args) == 7 || len(args) == 8 {
		invokeConfig.RepeatedTotal, err = strconv.ParseInt(args[6], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	if len(args) == 8 {
		configPath = args[7]
	}

	return invokeConfig, configPath, nil
}
