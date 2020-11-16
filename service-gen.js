const path = require('path');
const fs = require('fs');
const os = require('os');
const shell = require('child_process');

// 1 Receive parameters
var arguments = process.argv
var type = arguments[2]
var lang = arguments[3]
var service_name = arguments[4]
var schemasPath = arguments[5]
var output_dir = arguments[6]

// Parameter shelling
if (type != "provider" && type != "consumer") {
  throw new Error("Please enter correct type: consumer | provider.");
}

if (lang != "go") {
  throw new Error("only supported go currently.");
}

if (service_name == "") {
  throw "Please enter service name";
}

if (typeof (schemasPath) == "undefined") {
  schemasPath = "./schemas.json"
} else {
  schemasPath = path.join(__dirname, schemasPath)
}

if (fs.existsSync(schemasPath) == false) {
  throw "Schemas not exist."
}

if (typeof (output_dir) == "undefined") {
  output_dir = "./output"
} else {
  output_dir = path.join(__dirname, output_dir)
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
copyDir(template_path, output_dir);
fs.unlinkSync(output_dir + "/config/config.yaml")
fs.rmdirSync(output_dir + "/config")

// Copy config
copyDir(template_path + "/config", config_path)

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
replaceTemp(output_dir)

console.log("Complete template replacement.")

// 4 Install converter, read schema.json, convert to the corresponding language structure
// Create temporary folder
fs.mkdirSync(output_dir + "/.temp")

if (lang == "go") {
  shell.execSync("go env -w GOPROXY=https://goproxy.cn")
  shell.execSync("go get github.com/atombender/go-jsonschema/...")
  shell.execSync("go build -o " + output_dir + "/.temp github.com/atombender/go-jsonschema/cmd/gojsonschema")

  fs.writeFileSync(output_dir + "/.temp/ServiceInput.json", JSON.stringify(schemas.input))
  fs.writeFileSync(output_dir + "/.temp/ServiceOutput.json", JSON.stringify(schemas.output))
  shell.execSync(path.resolve(fs.realpathSync('.'), output_dir + "/.temp/gojsonschema -p types " + output_dir + "/.temp/ServiceInput.json >> " + output_dir + "/types/input.go"))
  shell.execSync(path.resolve(fs.realpathSync('.'), output_dir + "/.temp/gojsonschema -p types " + output_dir + "/.temp/ServiceOutput.json >> " + output_dir + "/types/output.go"))
  data = fs.readFileSync(output_dir + "/types/input.go")
  data = data.toString().replace(new RegExp("ServiceInputJson", 'g'), "ServiceInput");
  fs.writeFileSync(output_dir + "/types/input.go", data)
  data = fs.readFileSync(output_dir + "/types/output.go")
  data = data.toString().replace(new RegExp("ServiceOutputJson", 'g'), "ServiceOutput");
  fs.writeFileSync(output_dir + "/types/output.go", data)
  console.log("Complete parsing json.")
}

// Remove temporary folder
deleteDir(output_dir + "/.temp");

// 5 Installation project dependencies
console.log("Installing project dependencies...")
if (lang == "go") {
  // shell.execSync("cd " + output_dir)
  shell.execSync("cd " + output_dir + " && go mod tidy")
} else if (lang == "java") {
  console.log("java")
} else {
  console.log("js")
}
console.log("Complete installation project dependencies.")

function copyDir(src, dst) {
  if (fs.existsSync(dst) == false) {
    fs.mkdirSync(dst);
  }
  if (fs.existsSync(src) == false) {
    throw new Error("Path no exist: ", src)
  }
  var dirs = fs.readdirSync(src);
  dirs.forEach(function (item) {
    var item_path = path.join(src, item);
    var temp = fs.statSync(item_path);
    if (temp.isFile()) {
      fs.copyFileSync(item_path, path.join(dst, item));
    } else if (temp.isDirectory()) {
      copyDir(item_path, path.join(dst, item));
    }
  });
}

function replaceTemp(url) {
  let reg = new RegExp("{{service_name}}", 'g');
  let files = fs.readdirSync(url);
  files.forEach(function (file, index) {
    const curUrl = path.join(url, file);
    if (fs.statSync(curUrl).isDirectory()) {
      replaceTemp(curUrl);
    } else {
      let data = fs.readFileSync(curUrl)
      data = data.toString().replace(reg, service_name);
      fs.writeFileSync(curUrl, data)
    }
  });
}

function deleteDir(url) {
  let files = [];
  files = fs.readdirSync(url);
  files.forEach(function (file, index) {
    const curPath = path.join(url, file);
    if (fs.statSync(curPath).isDirectory()) {
      deleteDir(curPath);
    } else {
      fs.unlinkSync(curPath);
    }
  });
  fs.rmdirSync(url);
}
