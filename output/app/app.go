package app

import (
	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/tianle/service-gen/service"
	"gitlab.bianjie.ai/tianle/service-gen/{{service_name}}"
	"gitlab.bianjie.ai/tianle/service-gen/types"
)

// App represents the provider application
type App struct {
	ServiceClient   service.ServiceClientWrapper
	ServiceCallback service.ServiceCallback
	Logger          *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	serviceClient service.ServiceClientWrapper,
	serviceCallback service.ServiceCallback,
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
	err := app.ServiceClient.SubscribeServiceRequest(types.ServiceName, app.ServiceCallback)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}
