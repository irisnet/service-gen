package main

import (
	"github.com/spf13/cobra"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start consumer daemon",
		Example: `{{service_name}}-sc start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			password := getPassword()

			var configPath string

			if len(args) == 0 {
				configPath = common.ConfigPath
			} else {
				configPath = args[0]
			}

			config, err := common.LoadYAMLConfig(configPath)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), password)

			logger := common.Logger

			addr, err := serviceClient.ShowKey(serviceClient.KeyName, serviceClient.Password)
			if err != nil {
				return err
			}

			appInstance := app.NewApp(serviceClient, {{service_name}}.ResponseCallback, logger)
			appInstance.Start(addr, {{service_name}}.ResponseCallback)

			return nil
		},
	}

	return cmd
}
