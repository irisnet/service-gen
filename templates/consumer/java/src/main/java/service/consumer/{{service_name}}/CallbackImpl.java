package service.consumer.{{service_name}};

import com.alibaba.fastjson.JSON;

import service.consumer.service.ICallback;
import service.consumer.types.ServiceOutput;

public class CallbackImpl implements ICallback {
	public String keyName;
	public String password;

	public CallbackImpl(String keyName, String password) {
		this.keyName = keyName;
		this.password = password;
	}
  
  public void onResponse(ServiceOutput req) {
		System.out.println("----------------- Consumer -----------------");
		// Supplementary service logic...
		System.out.println("Got response: "+ JSON.toJSONString(req));
  }
}
