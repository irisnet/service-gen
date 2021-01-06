package service.consumer.application;

import com.alibaba.fastjson.JSONObject;
import org.bouncycastle.crypto.CryptoException;

import iservice.sdk.entity.options.TxOptions;
import iservice.sdk.core.ServiceClient;

import service.consumer.service.Client;
import service.consumer.service.ConsumerListener;
import service.consumer.types.InputWithHeader;
import service.consumer.types.ServiceRequest;
import service.consumer.{{service_name}}.CallbackImpl;

import java.io.IOException;

public class Application {

  public Client client;

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr, String chain_id, String fee) throws Exception {
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr, chain_id, fee);
  }

  public void invoke(String keyName, String password, String keyPath, String feeCap, String sender, String input) throws IOException, CryptoException {
    String address = importKey(keyName, password, password, keyPath);
    ServiceClient serviceClient = client.getServiceClient();

    InputWithHeader inputWithHeader = JSONObject.parseObject(input, InputWithHeader.class);
    ServiceRequest serviceRequest = new ServiceRequest(keyName, password, inputWithHeader.getHeader(), inputWithHeader.getBody());
    String amount = feeCap.replaceAll("[^0-9]", "");
    String denom = feeCap.replaceAll( "[^a-z]", "");
    serviceRequest.setFee(new TxOptions.Fee(amount, denom));
    serviceRequest.addProvider(sender);
    serviceClient.callService(serviceRequest);

    ConsumerListener consumerListener = new ConsumerListener();
    consumerListener.setOptions(address, sender);
    consumerListener.setICallback(new CallbackImpl(keyName, password));
    client.subscribe(consumerListener);
  }

  public String addKey(String keyName, String password) {
    return client.addKey(keyName, password);
  }

  public String showKey(String keyName) {
    return client.showKey(keyName);
  }

  public String importKey(String keyName, String password, String keyStorePassword, String filePath) {
    return client.importKey(keyName, password, keyStorePassword, filePath);
  }
}