package service.consumer.service;

import service.consumer.types.ServiceOutput;

public interface ICallback {
  
  void onResponse(ServiceOutput req);
}
