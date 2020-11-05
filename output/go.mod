module gitlab.bianjie.ai/tianle/service-gen

go 1.15

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/service-sdk-go => github.com/secret2830/service-sdk-go v0.0.0-20200930025908-91ed6ca17b1b
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703.0.20200920152706-f907f8a9ab6c
)

require (
	github.com/golang/snappy v0.0.2-0.20200707131729-196ae77b8a26 // indirect
	github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/irisnet/service-sdk-go v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200824131525-c12d262b63d8 // indirect
)
