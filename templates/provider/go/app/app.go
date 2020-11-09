package app

import (
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
	log "github.com/sirupsen/logrus"
)

// App represents the provider application
type App struct {
	ServiceClient   service.ServiceClientWrapper
	RequestCallback types.RequestCallback
	Logger          *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	serviceClient service.ServiceClientWrapper,
	requestCallback types.RequestCallback,
	logger *log.Logger,
) App {
	return App{
		ServiceClient:   serviceClient,
		RequestCallback: {{service_name}}.RequestCallback,
		Logger:          logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	// Subscribe
	err := app.ServiceClient.SubscribeServiceRequest(types.ServiceName, app.RequestCallback)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}
