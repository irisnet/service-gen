package service.consumer.application;

import com.alibaba.fastjson.JSONObject;
import org.apache.log4j.Logger;
import org.bouncycastle.crypto.CryptoException;

import iservice.sdk.core.ServiceClient;

import service.consumer.service.Client;
import service.consumer.service.ConsumerListener;
import service.consumer.types.InputWithHeader;
import service.consumer.types.ServiceRequest;
import service.consumer.{{service_name}}.CallbackImpl;

import java.io.IOException;

public class Application {

  public Client client;

  public static Logger logger = Logger.getLogger(Application.class);

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr) throws Exception {
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr);
  }

  public void invoke(String keyName, String password, String address, String sender, String input) throws IOException, CryptoException {
    ConsumerListener consumerListener = new ConsumerListener();
    consumerListener.setOptions(address, sender);
    consumerListener.setICallback(new CallbackImpl(keyName, password));

    // TODO
    recoverKey(keyName, password, "cup floor miss diagram salute dream hat secret ladder faith siren floor basic high battle little kitten fall live early police erase rifle asset");
    ServiceClient serviceClient = client.getServiceClient();

    InputWithHeader inputWithHeader = JSONObject.parseObject(input, InputWithHeader.class);
    ServiceRequest serviceRequest = new ServiceRequest(keyName, password, inputWithHeader.getHeader(), inputWithHeader.getBody());
    serviceRequest.addProvider(sender);
    serviceClient.callService(serviceRequest);
    client.subscribe(consumerListener);
  }

  public void addKey(String keyName, String password) {
    client.addKey(keyName, password);
  }

  public void showKey(String keyName) {
    client.showKey(keyName);
  }

  public void importKey(String keyName, String password, String filePath) {
    client.importKey(keyName, password, filePath);
  }

  public void recoverKey(String keyName, String password, String mnemonic) {
    client.recoverKey(keyName, password, mnemonic);
  }
}