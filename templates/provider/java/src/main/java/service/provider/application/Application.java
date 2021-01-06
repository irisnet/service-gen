package service.provider.application;

import service.provider.service.Client;
import service.provider.service.ProviderListener;
import service.provider.{{service_name}}.CallbackImpl;

public class Application {

  public Client client;

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr, String chainID, String fee) throws Exception {
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr, chainID, fee);
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
}