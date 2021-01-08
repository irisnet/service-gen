package service.consumer.service;

import iservice.sdk.core.AbstractConsumerListener;
import iservice.sdk.entity.options.ConsumerListenerOptions;

import service.consumer.common.Config;
import service.consumer.types.ServiceOutput;

public class ConsumerListener extends AbstractConsumerListener<ServiceOutput> {

	public ConsumerListenerOptions options;
	public ICallback iCallback;

	public void setOptions(String addr, String sender) {
		options = new ConsumerListenerOptions();
		options.setServiceName(Config.ServiceName);
		options.setAddress(addr);
		options.setSender(sender);
	}

	@Override
	public ConsumerListenerOptions getOptions() {
		return options;
	}

	public void setICallback(ICallback iCallback) {
		this.iCallback = iCallback;
	}

	@Override
	protected Class<ServiceOutput> getReqClass() {
		return ServiceOutput.class;
	}

	@Override
	public void onResponse(ServiceOutput req) {
		iCallback.onResponse(req);
	}
}
