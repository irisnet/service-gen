package service.provider.application;

import org.apache.log4j.Logger;

import service.provider.service.Client;
import service.provider.service.ProviderListener;
import service.provider.{{service_name}}.CallbackImpl;

public class Application {

  public Client client;

  public static Logger logger = Logger.getLogger(Application.class);

  public Application(String keyAlgorithm, String nodeRPCAddr, String nodeGRPCAddr) throws Exception {
    client = new Client(keyAlgorithm, nodeRPCAddr, nodeGRPCAddr);
  }

  public void start(String keyName, String password, String address) {
    ProviderListener providerListener = new ProviderListener();
    providerListener.setOptions(address);
    providerListener.setICallback(new CallbackImpl(keyName, password));

    // TODO
    recoverKey(keyName, password, "wash bargain vicious basket blur assist fault august involve quit fit camp eagle supreme chef process auction surge crucial orphan ticket hundred express bike");
    client.subscribe(providerListener);
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