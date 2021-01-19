package service.provider.cmd;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.Parameter;
import com.beust.jcommander.Parameters;

import service.provider.application.Application;
import service.provider.common.Config;

public class RootCmd {
  @Parameter(names = { "--help", "-h" }, help = true)
  private static boolean help;

  @Parameters(commandDescription = "Start.")
  public static class CommandStart {
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
    @Parameter(names = {"--name", "-n"}, description = "key name")
    private String keyName;

    @Parameter(names = {"--path", "-p"}, description = "key path")
    public String keyPath;

    @Parameter(names = { "--config", "-c" }, description = "Config path.")
    private String configPath;
  }

  public static void main(String[] args) {
    RootCmd root = new RootCmd();
    CommandStart start = new CommandStart();
    CommandAdd addKey = new CommandAdd();
    CommandShow showKey = new CommandShow();
    CommandImport importKey = new CommandImport();

    JCommander jc = JCommander.newBuilder()
      .addObject(root)
      .addCommand("start", start)
      .addCommand("add", addKey)
      .addCommand("show", showKey)
      .addCommand("import", importKey)
      .build();

    try {
      jc.parse(args);
    } catch (Exception e) {
      System.out.println(e.getMessage());
      jc.setProgramName("java -jar " + Config.ServiceName + "-sp.jar");
      jc.usage();
      return;
    }

    if (help) {
      jc.setProgramName("java -jar " + Config.ServiceName + "-sp.jar");
      jc.usage();
      return;
    }

    switch (jc.getParsedCommand()) {
      case "start":
        root.start(start);
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

  public void start(CommandStart start) {
    try {
      Config config = new Config(start.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      application.start(config.keyName, config.password, config.keyPath);
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }

  public void addKey(CommandAdd add) {
    try {
      Config config = new Config(add.configPath);
      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      System.out.println("Please enter the password again:\n");
      String pwd = Config.readPassword();
      if (!pwd.equals(config.password)) {
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

      if (commandImport.keyName == null) {
        commandImport.keyName = config.keyName;
      }
      if (commandImport.keyPath == null) {
        commandImport.keyPath = config.keyPath;
      }

      Application application = new Application(config.keyAlgorithm, config.nodeRPCAddr, config.nodeGRPCAddr, config.chainID, config.fee);
      application.importKey(commandImport.keyName, config.password, config.password, commandImport.keyPath);
    } catch (Exception e) {
      System.err.println(e.getMessage());
    }
  }
}
