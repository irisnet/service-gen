package main

import (
	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/servicename"
	"github.com/irisnet/service-gen/types"
	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start consumer daemon",
		Example: `servicename-sc start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(1)

			configPath := ""

			if len(args) == 0 {
				configPath = common.ConfigPath
			} else {
				configPath = args[0]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			logger := common.Logger

			addr, err := serviceClient.ShowKey(serviceClient.KeyName, serviceClient.Password)
			if err != nil {
				return err
			}

			appInstance := app.NewApp(serviceClient, servicename.ResponseCallback, logger)
			appInstance.Start(types.ServiceName, servicename.ResponseCallback, addr)

			return nil
		},
	}

	return cmd
}
