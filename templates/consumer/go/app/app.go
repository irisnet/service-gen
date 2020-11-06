package app

import (
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/servicename"
	"github.com/irisnet/service-gen/test"
	"github.com/irisnet/service-gen/types"
	log "github.com/sirupsen/logrus"
)

// App represents the provider application
type App struct {
	ServiceClient    service.ServiceClientWrapper
	ResponseCallback types.ResponseCallback
	Logger           *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	serviceClient service.ServiceClientWrapper,
	ResponseCallback types.ResponseCallback,
	logger *log.Logger,
) App {
	return App{
		ServiceClient:    serviceClient,
		ResponseCallback: servicename.ResponseCallback,
		Logger:           logger,
	}
}

func NewTestApp(
	serviceClient service.ServiceClientWrapper,
	ResponseCallback types.ResponseCallback,
	logger *log.Logger,
) App {
	return App{
		ServiceClient:    serviceClient,
		ResponseCallback: test.ResponseCallback,
		Logger:           logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	// Subscribe
	err := app.ServiceClient.SubscribeServiceResponse(types.ServiceName, app.ResponseCallback)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}
