package main

import (
	"github.com/spf13/cobra"
	"gitlab.bianjie.ai/tianle/servicegen/app"
	"gitlab.bianjie.ai/tianle/servicegen/common"
	"gitlab.bianjie.ai/tianle/servicegen/service"
	"gitlab.bianjie.ai/tianle/servicegen/{{service_name}}"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `{{service_name}}-sp start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase := readPassphrase(1)

			configFileName := ""

			if len(args) == 0 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[0]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), passphrase)

			logger := common.Logger

			appInstance := app.NewApp(serviceClient, {{service_name}}.ServiceCallback, logger)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
