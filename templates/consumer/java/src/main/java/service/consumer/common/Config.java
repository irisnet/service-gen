package service.consumer.common;

import org.yaml.snakeyaml.Yaml;

import java.io.FileInputStream;
import java.io.InputStream;
import java.util.Map;

public class Config {

    // TODO
    public final static String ServiceName = "{{service_name}}";
    private static final String DefaultChainId = "irita";
    private static final String DefaultNodeRPCAddr = "http://localhost:26657";
    private static final String DefaultNodeGRPCAddr = "http://localhost:9090";
    private static final String DefaultKeyName = "sc";
    private static final String DefaultKeyPath = "key.txt";
    // TODO
    private static final String DefaultFee = "4stake";
    private static final String DefaultKeyAlgorithm = "SM2";

    private static String ConfigPath = System.getProperty("user.home") + "/." + ServiceName + "-sc/config.yaml";

    public String chainID;
    public String nodeRPCAddr;
    public String nodeGRPCAddr;
    public String keyName;
    public String keyPath;
    public String address;
    public String fee;
    public String keyAlgorithm;
    public String password;

    public Config(String configPath) throws Exception {

        if (configPath != null) {
            ConfigPath = configPath;
        }

        password = readPassword();

        Map<String, String> configMap = loadYamlConfig();

        chainID = configMap.get("chain_id") != null ? configMap.get("chain_id") : DefaultChainId;
        nodeRPCAddr = configMap.get("node_rpc_addr") != null ? configMap.get("node_rpc_addr") : DefaultNodeRPCAddr;
        nodeGRPCAddr = configMap.get("node_grpc_addr") != null ? configMap.get("node_grpc_addr") : DefaultNodeGRPCAddr;

        keyName = configMap.get("key_name") != null ? configMap.get("key_name") : DefaultKeyName;
        keyPath = configMap.get("key_path") != null ? configMap.get("key_path") : DefaultKeyPath;
        address = configMap.get("address");

        fee = configMap.get("fee") != null ? configMap.get("fee") : DefaultFee;
        keyAlgorithm = configMap.get("key_algorithm") != null ? configMap.get("key_algorithm") : DefaultKeyAlgorithm;
    }

    private Map<String, String> loadYamlConfig() throws Exception {

        InputStream input = new FileInputStream(Config.ConfigPath);
        Yaml yaml = new Yaml();
        return (Map<String, String>) yaml.load(input);
    }

    private String readPassword() {
        System.out.println("Please enter password:");
        char[] password = System.console().readPassword();
        return String.valueOf(password);
    }
}
