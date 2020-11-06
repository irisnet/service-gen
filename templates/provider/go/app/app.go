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
	ServiceCallback types.ServiceCallback
	Logger          *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	serviceClient service.ServiceClientWrapper,
	serviceCallback types.ServiceCallback,
	logger *log.Logger,
) App {
	return App{
		ServiceClient:   serviceClient,
		ServiceCallback: {{service_name}}.ServiceCallback,
		Logger:          logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	// Subscribe
	err := app.ServiceClient.SubscribeServiceRequest(types.{{service_name}}, app.ServiceCallback)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}
