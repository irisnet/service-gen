package service.provider.service;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.NoSuchElementException;
import java.util.concurrent.TimeUnit;

import iservice.sdk.core.AbstractServiceListener;
import iservice.sdk.core.ServiceClient;
import iservice.sdk.core.ServiceClientFactory;
import iservice.sdk.entity.SignAlgo;
import iservice.sdk.entity.options.ServiceClientOptions;
import iservice.sdk.entity.options.TxOptions;
import iservice.sdk.exception.ServiceSDKException;
import iservice.sdk.module.IKeyService;

public class Client {
    public ServiceClient serviceClient;

    public Client(String algorithm, String rpc, String grpc, String chainID, String amount, String denom) throws URISyntaxException {
        ServiceClientOptions options = new ServiceClientOptions();
        if (algorithm.toUpperCase().equals("SM2")) {
            options.setSignAlgo(SignAlgo.SM2);
        }
        options.setRpcURI(new URI(rpc));
        options.setGrpcURI(new URI(grpc));
        options.setRpcStartTimeout(TimeUnit.SECONDS.toMillis(7));
        this.serviceClient = ServiceClientFactory.getInstance()
                .setOptions(options)
                .getClient();

        TxOptions txOptions = new TxOptions(chainID, amount, denom);
        this.serviceClient.setTxOptions(txOptions);
    }

    public void subscribe(AbstractServiceListener listener) {
        this.serviceClient.addListener(listener);
        this.serviceClient.start();
    }

    public String addKey(String keyName, String password) {
        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.addKey(keyName, password).toString();
        } catch (ServiceSDKException e) {
            System.err.println(e.getMessage());
        }
        System.out.println(address);
        return address;
    }

    public String showKey(String keyName) {
        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.showAddress(keyName);
        } catch (ServiceSDKException e) {
            System.err.println(e.getMessage());
        }
        System.out.println(address);
        return address;
    }

    public String importKey(String keyName, String password, String keyStorePassword, String filePath) {

        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.importFromKeystore(keyName, password, keyStorePassword, filePath);
        } catch (NoSuchElementException e) {
            System.err.println("Wrong password!");
        } catch (ServiceSDKException | IOException e) {
            System.err.println(e.getMessage());
        }
        System.out.println(address);
        return address;
    }
}
