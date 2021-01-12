package service.provider.{{service_name}};

import com.alibaba.fastjson.JSON;

import service.provider.application.Application;
import service.provider.types.ServiceInput;
import service.provider.types.ServiceOutput;
import service.provider.service.ICallback;
import service.provider.types.ServiceResponse;

public class CallbackImpl implements ICallback {
	public String keyName;
	public String password;

	public CallbackImpl(String keyName, String password) {
		this.keyName = keyName;
		this.password = password;
	}
  
  public ServiceResponse onRequest(ServiceInput req) {
		System.out.println("----------------- Provider -----------------");
		System.out.println("Got request");
		System.out.println(JSON.toJSONString(req));

		// Supplementary service logic...
		ServiceOutput output = new ServiceOutput();

		System.out.println("Sending response");
		ServiceResponse res = new ServiceResponse(this.keyName, this.password);
		res.setBody(output);

		return res;
  }
}