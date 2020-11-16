const path = require('path');
const fs = require('fs');
const os = require('os');
const shell = require('child_process');
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
var type = cmd.type
var lang = cmd.lang
var service_name = cmd.serviceName
var schemasPath = cmd.schemas
var output_dir = cmd.output

// Parameter shelling
if (type == "p") type = "provider"
if (type == "c") type = "consumer"

if (type != "provider" && type != "consumer") {
  console.log("Please enter correct type: consumer | provider.");
  return
}

if (lang != "go") {
  console.log("only supported go currently.");
  return
}

if (typeof (service_name) == "undefined") {
  console.log("Please enter service name");
  return
}

if (fs.existsSync(schemasPath) == false) {
  console.log("Schemas not exist.")
  return
}

if (fs.existsSync(output_dir) == false) {
  fs.mkdirSync(output_dir);
}

// Record template path
const template_path = fs.realpathSync('.') + "/templates/" + type + "/" + lang;
// Record config path
let config_path
if (type == "consumer") {
  config_path = os.homedir() + "/." + service_name + "-sc"
} else {
  config_path = os.homedir() + "/." + service_name + "-sp"
}
// Record schemas path
const schemas = require(schemasPath)
console.log("Complete initialization.")

// 2 Copy the specified template to the specified project path
utils.CopyDir(template_path, output_dir);
fs.unlinkSync(output_dir + "/config/config.yaml")
fs.rmdirSync(output_dir + "/config")

// Copy config
utils.CopyDir(template_path + "/config", config_path)

console.log("Complete creating project.")

// 3 Modify template variables
// Modify folder name
fs.mkdirSync(output_dir + "/" + service_name)
if (type == "consumer") {
  fs.renameSync(output_dir + "/{{service_name}}/response_callback.go", output_dir + "/" + service_name + "/response_callback.go")
} else {
  fs.renameSync(output_dir + "/{{service_name}}/request_callback.go", output_dir + "/" + service_name + "/request_callback.go")
}
fs.rmdirSync(output_dir + "/{{service_name}}")

// Modify the service name in the app.go
utils.ReplaceTemp(output_dir, service_name)

console.log("Complete template replacement.")

// 4 Install converter, read schema.json, convert to the corresponding language structure
// Create temporary folder
fs.mkdirSync(output_dir + "/.temp")

if (lang == "go") {
  utils.GoParseJson(output_dir, schemas)
}
console.log("Complete parsing json.")

// Remove temporary folder
utils.DeleteDir(output_dir + "/.temp");

// 5 Installation project dependencies
console.log("Installing project dependencies...")
if (lang == "go") {
  // shell.execSync("cd " + output_dir)
  shell.execSync("cd " + output_dir + " && go mod tidy")
}
console.log("Complete installation project dependencies.")
