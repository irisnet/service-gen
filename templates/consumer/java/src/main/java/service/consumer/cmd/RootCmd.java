package service.consumer.cmd;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.Parameter;
import com.beust.jcommander.Parameters;

import service.consumer.application.Application;
import service.consumer.common.Config;

public class RootCmd {

  @Parameter(names = { "--help", "-h" }, help = true)
  private static boolean help;

  @Parameters(commandDescription = "Invoke.")
  public static class CommandInvoke {
    @Parameter(names = {"--provider", "-p"}, description = "Provider list to call.", required = true)
    private String providers;

    @Parameter(names = {"--input", "-i"}, description = "input", required = true)
    private String input;

    @Parameter(names = {"--fee-cap", "-f"}, description = "fee cap", required = true)
    private String feeCap;

    @Parameter(names = { "--config", "-c" }, description = "Config path.")
    private String configPath;
  }

  @Parameters(commandDescription = "Add key.")
  public static class CommandAdd {
    @Parameter(names = {"--name", "-n"}, description = "key name", required = true)
    private String keyName;

    @Parameter(names = { "--config", "-c" }, description = "Config path.")
    private String configPath;
  }

  @Parameters(commandDescription = "Show key.")
  public static class CommandShow {
    @Parameter(names = {"--name", "-n"}, description = "key name", required = true)
    private String keyName;

    @Parameter(names = { "--config", "-c" }, description = "Config path.")
    private String configPath;
  }

  @Parameters(commandDescription = "Import Key.")
  public static class CommandImport {
    @Parameter(names = {"--name", "-n"}, description = "key name", required = true)
    private String keyName;

    @Parameter(names = {"--path", "-p"}, description = "key path", required = true)
    public String keyPath;

    @Parameter(names = { "--config", "-c" }, description = "Config path.")
    private String configPath;
  }

  public static void main(String[] args) {
    RootCmd root = new RootCmd();
    CommandInvoke invoke = new CommandInvoke();
    CommandAdd addKey = new CommandAdd();
    CommandShow showKey = new CommandShow();
    CommandImport importKey = new CommandImport();

    JCommander jc = JCommander.newBuilder()
      .addObject(root)
      .addCommand("invoke", invoke)
      .addCommand("add", addKey)
      .addCommand("show", showKey)
      .addCommand("import", importKey)
      .build();

    try {
      jc.parse(args);
    } catch (Exception e) {
      System.out.println(e.getMessage());
      jc.setProgramName("java -jar " + Config.ServiceName + "-sc-jar-with-dependencies.jar");
      jc.usage();
      return;
    }

    if (help) {
      jc.setProgramName("java -jar " + Config.ServiceName + "-sp-jar-with-dependencies.jar");
      jc.usage();
      return;
    }

    switch (jc.getParsedCommand()) {
      case "invoke":
        root.invoke(invoke);
        break;
      case "add":
        root.addKey(addKey);
        break;
      case "show":
        root.showKey(showKey);
        break;
      case "import":
        root.importKey(importKey);
        break;
    }
  }

  public void invoke(CommandInvoke invoke) {
    try {
      Config config = new Config(invoke.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      application.invoke(config.keyName, config.password, config.keyPath, invoke.feeCap, invoke.providers, invoke.input);
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }

  public void addKey(CommandAdd add) {
    try {
      Config config = new Config(add.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      System.out.println("Please enter the password again:\n");
      String pwd = readPassword();
      if (pwd != config.password) {
        System.out.println("The two passwords do not match!");
        return;
      }
      application.addKey(add.keyName, config.password);
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }

  public void showKey(CommandShow show) {
    try {
      Config config = new Config(show.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      application.showKey(show.keyName);
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }

  public void importKey(CommandImport commandImport) {
    try {
      Config config = new Config(commandImport.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      application.importKey(commandImport.keyName, config.password, config.password, commandImport.keyPath);
    } catch (Exception e) {
      Application.logger.error(e.getMessage());
    }
  }

  private String readPassword() {
    System.out.println("Please enter password:");
    char[] password = System.console().readPassword();
    return String.valueOf(password);
  }
}
