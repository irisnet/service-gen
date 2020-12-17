package app

import (
	"fmt"

	servicesdk "github.com/irisnet/service-sdk-go/service"
	log "github.com/sirupsen/logrus"

	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
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
		ResponseCallback: {{service_name}}.ResponseCallback,
		Logger:           logger,
	}
}

func (app App) subscribe(reqCtxID, reqID string, responseCallback types.ResponseCallback) {
	addr, err := app.ServiceClient.ShowKey(app.ServiceClient.KeyName, app.ServiceClient.Password)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}
	// Subscribe
	err = app.ServiceClient.SubscribeServiceResponse(reqCtxID, reqID, addr, app.ResponseCallback)
	if err != nil {
		app.Logger.Errorf("failed to subscribe service request, err: %s", err.Error())
		return
	}

	select {}
}

// Invoke providers' service
func (app App) Invoke(invokeConfig servicesdk.InvokeServiceRequest) {
	reqCtxID, reqID := app.ServiceClient.InvokeService(invokeConfig)

	fmt.Println("Successfully invoke service.")
	fmt.Println("reqCtxID:", reqCtxID)
	fmt.Println("reqID:", reqID)

	app.subscribe(reqCtxID, reqID, app.ResponseCallback)
}
