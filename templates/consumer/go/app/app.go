package app

import (
	"fmt"

	servicesdk "github.com/irisnet/service-sdk-go/service"
	log "github.com/sirupsen/logrus"

	"github.com/irisnet/service-gen/common"
	"github.com/irisnet/service-gen/service"
	callback "github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/types"
)

// App represents the provider application
type App struct {
	ServiceClient    service.ServiceClientWrapper
	ResponseCallback types.ResponseCallback
	Logger           *log.Logger
}

// NewApp constructs a new App instance
func NewApp(serviceClient service.ServiceClientWrapper) App {
	return App{
		ServiceClient:    serviceClient,
		ResponseCallback: callback.ResponseCallback,
		Logger:           common.Logger,
	}
}

func (app App) subscribe(reqCtxID, reqID string) {
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
	invokeConfig.Callback = callback.ResponseCallback
	reqCtxID, reqID, err := app.ServiceClient.InvokeService(invokeConfig)
	if err != nil {
		app.Logger.Errorf("failed to invoke service request, err: %s \n", err.Error())
		return
	}

	fmt.Println("Successfully invoke service.")
	fmt.Println("reqCtxID:", reqCtxID)
	fmt.Println("reqID:", reqID)

	app.subscribe(reqCtxID, reqID)
}
