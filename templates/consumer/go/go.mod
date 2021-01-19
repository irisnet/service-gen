module github.com/irisnet/service-gen

go 1.15

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/service-sdk-go => github.com/GTLiSunnyi/service-sdk-go v1.0.0-rc0.0.20210112025042-ac43e9ccba7d
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.1-dev0.0.20201126055325-2217bc51b6c7
)

require (
	github.com/Workiva/go-datastructures v1.0.52
	github.com/cosmos/cosmos-sdk v0.40.0 // indirect
	github.com/ethereum/go-ethereum v1.9.23
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/irisnet/service-sdk-go v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.6.1
)
