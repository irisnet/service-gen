package service.consumer.types;

import iservice.sdk.entity.Header;

public class InputWithHeader {
  Header header;
  ServiceInput body;

  public Header getHerder() {
    return header;
  }

  public ServiceInput getBody() {
    return body;
  }

  public void setHeader(Header header) {
    this.header = header;
  }

  public void setBody(ServiceInput body) {
    this.body = body;
  }
}
