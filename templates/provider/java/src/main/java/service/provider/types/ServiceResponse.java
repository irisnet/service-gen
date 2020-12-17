package service.provider.types;

import iservice.sdk.entity.BaseServiceResponse;
import iservice.sdk.entity.Header;

import service.provider.common.Config;

// TODO
public class ServiceResponse extends BaseServiceResponse<ServiceOutput> {
  
  public ServiceOutput body;
  public String keyName;
  public String password;

  public ServiceResponse(String keyName, String password) {
    this.keyName = keyName;
    this.password = password;
  }

  @Override
  public String getKeyName() {
    return this.keyName;
  }

  @Override
  public String getKeyPassword() {
    return this.password;
  }

  @Override
  public String getServiceName() {
    return Config.ServiceName;
  }

  @Override
  public Header getHeader() {
    return new Header();
  }

  @Override
  public ServiceOutput getBody() {
    return this.body;
  }

  public void setBody(ServiceOutput body) {
    this.body = body;
  }
}
