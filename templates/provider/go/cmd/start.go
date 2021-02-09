package main

import (
	"github.com/spf13/cobra"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `{{service_name}}-sp start [config-file]`,
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

			appInstance := app.NewApp(serviceClient)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
