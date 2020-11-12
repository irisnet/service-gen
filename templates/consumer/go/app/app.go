package app

import (
	"fmt"

	"github.com/irisnet/service-gen/service"
	"github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
	servicesdk "github.com/irisnet/service-sdk-go/service"
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
		ResponseCallback: {{service_name}}.ResponseCallback,
		Logger:           logger,
	}
}

// Start starts the provider process
func (app App) Start(addr string, responseCallback types.ResponseCallback) {
	addr, err := app.ServiceClient.ShowKey(app.ServiceClient.KeyName, app.ServiceClient.Password)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}
	// Subscribe
	err = app.ServiceClient.SubscribeServiceResponse(addr, responseCallback)
	if err != nil {
		app.Logger.Errorf("failed to subscribe service request listener, err: %s", err.Error())
		return
	}

	select {}
}

// Invoke providers' service
func (app App) Invoke(invokeConfig servicesdk.InvokeServiceRequest) {
	reqCtxID, reqID, err := app.ServiceClient.InvokeService(invokeConfig)
	if err != nil {
		fmt.Printf("failed to invoke service request listener, err: %s \n", err.Error())
		return
	}
	fmt.Println("Successfully invoke service.")
	fmt.Println("reqCtxID:", reqCtxID)
	fmt.Println("reqID:", reqID)
}
