package service.consumer.service;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.concurrent.TimeUnit;

import iservice.sdk.core.AbstractServiceListener;
import iservice.sdk.core.ServiceClient;
import iservice.sdk.core.ServiceClientFactory;
import iservice.sdk.entity.SignAlgo;
import iservice.sdk.entity.options.ServiceClientOptions;
import iservice.sdk.exception.ServiceSDKException;
import iservice.sdk.module.IKeyService;

public class Client {
    public final ServiceClient serviceClient;

    public Client(String algorithm, String rpc, String grpc) throws URISyntaxException {
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
    }

    public void subscribe(AbstractServiceListener listener) {
        this.serviceClient.addListener(listener);
        this.serviceClient.start();
    }

    public void addKey(String keyName, String password) {
        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.addKey(keyName, password).toString();
        } catch (ServiceSDKException e) {
            System.out.println(e.getMessage());
        }
        System.out.println(address);;
    }

    public void showKey(String keyName) {
        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.showAddress(keyName);
        } catch (ServiceSDKException e) {
            System.out.println(e.getMessage());
        }
        System.out.println(address);
    }

    public void importKey(String keyName, String password, String filePath) {

        IKeyService keyService = serviceClient.getKeyService();
        String address = "";
        try {
            address = keyService.importFromKeystore(keyName, password, filePath);
        } catch (IOException | ServiceSDKException e) {
            System.out.println(e.getMessage());
        }
        System.out.println(address);
    }

    public void recoverKey(String keyName, String password, String mnemonic) {
        IKeyService keyService = serviceClient.getKeyService();
        try {
            address = keyService.recoverKey(keyName, password, mnemonic, true, 0, "");
        } catch (ServiceSDKException e) {
            System.out.println(e.getMessage());
        }
    }

    public ServiceClient getServiceClient() {
        return this.serviceClient;
    }
}
