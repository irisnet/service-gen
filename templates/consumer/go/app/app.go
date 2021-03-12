package app

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/irisnet/service-gen/common"
	callback "github.com/irisnet/service-gen/{{service_name}}"
	"github.com/irisnet/service-gen/service"
	servicesdk "github.com/irisnet/service-sdk-go/service"
)

// App represents the provider application
type App struct {
	ServiceClient    service.ServiceClientWrapper
	ResponseCallback servicesdk.InvokeCallback
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

// Invoke providers' service
func (app App) Invoke(invokeConfig servicesdk.InvokeServiceRequest) {
	invokeConfig.Callback = app.ResponseCallback
	reqCtxID, reqID, err := app.ServiceClient.InvokeService(invokeConfig)
	if err != nil {
		app.Logger.Errorf("failed to invoke service request, err: %s \n", err.Error())
		return
	}

	common.Logger.Info("Successfully invoke service.")
	common.Logger.Info("reqCtxID:", reqCtxID)
	common.Logger.Info("reqID:", reqID)

	for {
		queryRequestContextResp, err := app.ServiceClient.ServiceClient.QueryRequestContext(reqCtxID)
		if err != nil {
			common.Logger.Error("lost a block")
		}

		if queryRequestContextResp.ServiceName == "" {
			common.Logger.Info("finished subscribing")
			return
		}

		time.Sleep(5 * time.Second)
	}
}
