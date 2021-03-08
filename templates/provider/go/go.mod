module github.com/irisnet/service-gen

go 1.15

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

require (
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/irisnet/service-sdk-go master
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/tidwall/gjson v1.6.1
)
