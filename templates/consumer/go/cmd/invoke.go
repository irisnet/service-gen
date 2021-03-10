package main

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	servicesdk "github.com/irisnet/service-sdk-go/service"
	sdkTypes "github.com/irisnet/service-sdk-go/types"

	"github.com/irisnet/service-gen/app"
	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/types"
)

const (
	flagProviders = "providers"
	flagFeeCap    = "fee-cap"
	flagInput     = "input"
	flagTimeout   = "timeout"
	flagRepeated  = "repeated"
	flagFrequency = "frequency"
	flagTotal     = "total"
)

func invokeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Invoke service",
		Example: `a
{{service_name}}-sc invoke [config-path] \
	--providers iaa135p42vm5vxrk4rmryn6sqgusm4yqwxmqgm05tn \
	--fee-cap 0.04iris \
	--input '{"header":{},"body":{"input":"hello"}}' \
	--timeout 100 \
	--frequency 110 \
	--total 1`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			password := getPassword()

			providers := viper.GetStringSlice(flagProviders)
			feeCap := viper.GetString(flagFeeCap)
			input := viper.GetString(flagInput)
			timeout := viper.GetInt64(flagTimeout)
			repeated := viper.GetBool(flagRepeated)
			frequency := viper.GetUint64(flagFrequency)
			total := viper.GetInt64(flagTotal)

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

			feecap, err := sdkTypes.ParseDecCoin(feeCap)
			if err != nil {
				return err
			}

			serviceFeeCap := sdkTypes.NewDecCoins(feecap)

			invokeConfig := servicesdk.InvokeServiceRequest{
				ServiceName:       types.ServiceName,
				Providers:         providers,
				ServiceFeeCap:     serviceFeeCap,
				Input:             input,
				Timeout:           timeout,
				Repeated:          repeated,
				RepeatedFrequency: frequency,
				RepeatedTotal:     total,
			}

			serviceClient := service.MakeServiceClientWrapper(service.NewConfig(config), password)

			appInstance := app.NewApp(serviceClient)

			appInstance.Invoke(invokeConfig)

			return nil
		},
	}

	cmd.Flags().AddFlagSet(fsInvoke)
	_ = viper.BindPFlags(fsInvoke)

	_ = cmd.MarkFlagRequired(flagProviders)
	_ = cmd.MarkFlagRequired(flagFeeCap)
	_ = cmd.MarkFlagRequired(flagInput)

	return cmd
}

var (
	fsInvoke = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsInvoke.StringSlice(flagProviders, nil, "Providers that you want to invoke(Use '/' to split).")
	fsInvoke.String(flagFeeCap, "", "fee cap")
	fsInvoke.String(flagInput, "", "input")
	fsInvoke.Int64(flagTimeout, types.DefaultTimeout, "timeout")
	fsInvoke.Bool(flagRepeated, types.DefaultRepeated, "wheather repeat")
	fsInvoke.Uint64(flagFrequency, types.DefaultFrequency, "frequency")
	fsInvoke.Int64(flagTotal, types.DefaultTotal, "total invoke times")
}
