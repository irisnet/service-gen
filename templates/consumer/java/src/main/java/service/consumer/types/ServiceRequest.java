package service.consumer.types;

import com.google.common.collect.Lists;
import cosmos.base.v1beta1.CoinOuterClass;

import java.util.ArrayList;
import java.util.List;

import iservice.sdk.entity.BaseServiceRequest;
import iservice.sdk.entity.Header;
import iservice.sdk.entity.options.TxOptions;

import service.consumer.common.Config;

public class ServiceRequest extends BaseServiceRequest<ServiceInput> {
  public Header header;
  public ServiceInput body;

  public TxOptions.Fee fee;

  public String keyName;
  public String password;
  public List<String> providerList = new ArrayList<>();

  public ServiceRequest(String keyName, String password, Header header, ServiceInput body) {
    this.keyName = keyName;
    this.password = password;
    this.header = header;
    this.body = body;
  }

  public void setFee(TxOptions.Fee fee) {
    this.fee = fee;
  }

  @Override
  public String getKeyName() {
    return keyName;
  }

  @Override
  public String getKeyPassword() {
    return password;
  }

  @Override
  public String getServiceName() {
    return Config.ServiceName;
  }

  @Override
  public List<String> getProviders() {
    return providerList;
  }

  @Override
  public List<CoinOuterClass.Coin> getServiceFeeCap() {
    return Lists.newArrayList(CoinOuterClass.Coin.newBuilder().setAmount(this.fee.amount).setDenom(this.fee.denom).build());
  }

  public void addProvider(String provider) {
    providerList.add(provider);
  }

  @Override
  public Header getHeader() {
    return new Header();
  }

  @Override
  public ServiceInput getBody() {
    return body;
  }
}
