package service.consumer.application;

import com.alibaba.fastjson.JSONObject;
import org.bouncycastle.crypto.CryptoException;

import java.io.IOException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import iservice.sdk.entity.options.TxOptions;
import iservice.sdk.core.ServiceClient;

import service.consumer.service.Client;
import service.consumer.service.ConsumerListener;
import service.consumer.types.InputWithHeader;
import service.consumer.types.ServiceRequest;
import service.consumer.{{service_name}}.CallbackImpl;

public class Application {

  public Client client;

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr, String chainID, String feeConfig) throws Exception {
    String[] fee = getFee(feeConfig);
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr, chainID, fee[0], fee[1]);
  }

  public void invoke(String keyName, String password, String keyPath, String feeCap, String sender, String input) throws IOException, CryptoException {
    String address = importKey(keyName, password, password, keyPath);
    ServiceClient serviceClient = client.getServiceClient();

    InputWithHeader inputWithHeader = JSONObject.parseObject(input, InputWithHeader.class);
    ServiceRequest serviceRequest = new ServiceRequest(keyName, password, inputWithHeader.getHeader(), inputWithHeader.getBody());

    String[] fee = getFee(feeCap);
    serviceRequest.setFee(new TxOptions.Fee(fee[0], fee[1]));
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

  public String[] getFee(String fee) {
    String amount = "";
    String denom = "";

    Matcher matcher1 = Pattern.compile("^[0-9][0-9]*").matcher(fee);
    Matcher matcher2 = Pattern.compile("[a-zA-Z][a-zA-Z0-9]{2,127}$").matcher(fee);
    if (matcher1.find() && matcher2.find()) {
      amount = matcher1.group(0);
      denom = matcher2.group(0);
    } else {
      System.err.println("Invalid fee cap in input or config.yaml!");
      System.exit(0);
    }

    return new String[]{amount, denom};
  }
}