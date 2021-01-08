package service.provider.service;

import iservice.sdk.core.AbstractProviderListener;
import iservice.sdk.entity.options.ProviderListenerOptions;

import service.provider.types.ServiceInput;
import service.provider.types.ServiceOutput;
import service.provider.common.Config;
import service.provider.types.ServiceResponse;

public class ProviderListener extends AbstractProviderListener<ServiceInput, ServiceOutput, ServiceResponse> {

	public ProviderListenerOptions options;
	public ICallback iCallback;

	public void setOptions(String addr) {
		options = new ProviderListenerOptions();
		options.setServiceName(Config.ServiceName);
		options.setAddress(addr);
	}

	@Override
	public ProviderListenerOptions getOptions() {
		return options;
	}

	public void setICallback(ICallback iCallback) {
		this.iCallback = iCallback;
	}

	@Override
	protected Class<ServiceInput> getReqClass() {
		return ServiceInput.class;
	}

	@Override
	public ServiceResponse onRequest(ServiceInput req) {
		return this.iCallback.onRequest(req);
	}
}
