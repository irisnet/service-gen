package service.provider.service;

import service.provider.types.ServiceInput;
import service.provider.types.ServiceResponse;

public interface ICallback {
  
  ServiceResponse onRequest(ServiceInput req);
}
