const fs = require('fs');
const os = require('os');
const utils = require("./utils.js");
const cmd = require('commander');

cmd
  .version('0.1.0', '-v, --version')
  .option('--type <string>', 'provider(p) or consumer(c)')
  .option('--lang <string>', 'code language')
  .option('-s, --service-name <string>', 'service and package name')
  .option('--schemas [string]', 'path of jsonschema', './schemas.json')
  .option('-o, --output [string]', 'path of output dir', './output')
  .parse(process.argv);

// 1 Receive parameters
let type = cmd.type
let lang = cmd.lang
let service_name = cmd.serviceName
let schemasPath = cmd.schemas
let output_dir = cmd.output

if (type === "p") type = "provider"
if (type === "c") type = "consumer"

if (type !== "provider" && type !== "consumer") {
  console.log("Please enter correct type: consumer | provider.");
  return
}

if (lang !== "go" && lang !== "java" && lang !== "js") {
  console.log("Only support go, java, js.");
  return
}

if (/[A-Z]/g.test(service_name)) {
  console.log("************************\nNote: Uppercase letters are converted to lowercase!\n************************")
  service_name = service_name.toLowerCase()
}
if (typeof (service_name) == "undefined") {
  console.log("Please enter service name");
  return
}

if (fs.existsSync(schemasPath) === false) {
  console.log("Schemas not exist.")
  return
}

if (fs.existsSync(output_dir) === false) {
  fs.mkdirSync(output_dir)
}

// Record template path
let template_path = fs.realpathSync('.') + "/templates/" + type + "/" + lang
// Record config path
let config_path
if (type === "consumer") {
  config_path = os.homedir() + "/." + service_name + "-sc/"
} else {
  config_path = os.homedir() + "/." + service_name + "-sp/"
}
// Record schemas path
const schemas = JSON.parse(fs.readFileSync(schemasPath).toString().trim());
console.log("Complete initialization.")

// 2 Copy the specified template to the specified project path
utils.CopyDir(template_path, output_dir)
if (lang === "java") {
  utils.DeleteDir(output_dir + "/lib")
  fs.mkdirSync(output_dir + "/lib")
  utils.copyFile(template_path + "/lib/service-sdk-1.0-SNAPSHOT-jar-with-dependencies.jar", output_dir + "/lib/service-sdk-1.0-SNAPSHOT-jar-with-dependencies.jar")
}

// 3 Modify template variables
// Modify folder name
if (lang === "go") {
  utils.DeleteDir(output_dir + "/config")
  utils.CopyDir(template_path + "/config", config_path)
  fs.mkdirSync(output_dir + "/" + service_name)
  fs.renameSync(output_dir + "/{{service_name}}/callback.go", output_dir + "/" + service_name + "/callback.go")
  fs.rmdirSync(output_dir + "/{{service_name}}")
} else if (lang === "java") {
  utils.DeleteDir(output_dir + "/src/main/java/service/" + type + "/config")
  utils.CopyDir(template_path + "/src/main/java/service/" + type + "/config", config_path)
  fs.mkdirSync(output_dir + "/src/main/java/service/" + type + "/" + service_name)
  fs.renameSync(output_dir + "/src/main/java/service/" + type + "/{{service_name}}/CallbackImpl.java", output_dir + "/src/main/java/service/" + type + "/" + service_name + "/CallbackImpl.java")
  fs.rmdirSync(output_dir + "/src/main/java/service/" + type + "/{{service_name}}")
}

// Modify the service name
utils.ReplaceTemp(output_dir, "{{service_name}}", service_name)
if (service_name.indexOf("-") !== -1) {
  if (lang === "go") {
    utils.ReplaceTemp(output_dir + "/" + service_name, service_name, service_name.replace(/-/g, "_"))
  }
  if (lang === "java") {
    utils.ReplaceTemp(output_dir + "/src/main/java/service/" + type + service_name, service_name, service_name.replaceAll(/-/g, "_"))
  }
}

console.log("Complete copying config and replacing template.")

// 4 Install converter, read schema.json, convert to the corresponding language structure
// Create temporary folder
fs.mkdirSync(output_dir + "/.temp")

if (lang === "go") {
  utils.GoParseJson(output_dir, schemas)
} else if (lang === "java") {
  utils.JavaParseJson(output_dir, schemas, type)
}
console.log("Complete parsing json.")

// Remove temporary folder
utils.DeleteDir(output_dir + "/.temp");

console.log("Complete installation project dependencies.")
