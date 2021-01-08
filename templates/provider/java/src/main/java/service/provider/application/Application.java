package service.provider.application;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import service.provider.service.Client;
import service.provider.service.ProviderListener;
import service.provider.{{service_name}}.CallbackImpl;

public class Application {

  public Client client;

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr, String chainID, String feeConfig) throws Exception {
    String[] fee = getFee(feeConfig);
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr, chainID, fee[0], fee[1]);
  }

  public void start(String keyName, String password, String keyPath) {
    String address = importKey(keyName, password, password, keyPath);

    ProviderListener providerListener = new ProviderListener();
    providerListener.setOptions(address);
    providerListener.setICallback(new CallbackImpl(keyName, password));
    client.subscribe(providerListener);
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
      System.err.println("Invalid fee cap in config.yaml!");
      System.exit(0);
    }

    return new String[]{amount, denom};
  }
}