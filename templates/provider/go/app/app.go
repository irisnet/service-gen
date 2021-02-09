package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
)

// App represents the provider application
type App struct {
	ServiceClient   service.ServiceClientWrapper
	RequestCallback types.RequestCallback
	Logger          *log.Logger
}

// NewApp constructs a new App instance
func NewApp(serviceClient service.ServiceClientWrapper) App {
	return App{
		ServiceClient:   serviceClient,
		RequestCallback: {{service_name}}.RequestCallback,
		Logger:          common.Logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	// Subscribe
	err := app.ServiceClient.SubscribeServiceRequest(app.RequestCallback)
	if err != nil {
		app.Logger.Errorf("failed to subscribe service request, err: %s", err.Error())
		return
	}

	select {}
}
