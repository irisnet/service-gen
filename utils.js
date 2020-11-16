const fs = require('fs')
const path = require('path')
const shell = require('child_process')

function CopyDir(src, dst) {
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
      CopyDir(item_path, path.join(dst, item));
    }
  });
}

function ReplaceTemp(url, service_name) {
  let reg = new RegExp("{{service_name}}", 'g');
  let files = fs.readdirSync(url);
  files.forEach(function (file, index) {
    const curUrl = path.join(url, file);
    if (fs.statSync(curUrl).isDirectory()) {
      ReplaceTemp(curUrl, service_name);
    } else {
      let data = fs.readFileSync(curUrl)
      data = data.toString().replace(reg, service_name);
      fs.writeFileSync(curUrl, data)
    }
  });
}

function DeleteDir(url) {
  let files = [];
  files = fs.readdirSync(url);
  files.forEach(function (file, index) {
    const curPath = path.join(url, file);
    if (fs.statSync(curPath).isDirectory()) {
      DeleteDir(curPath);
    } else {
      fs.unlinkSync(curPath);
    }
  });
  fs.rmdirSync(url);
}

function GoParseJson(output_dir, schemas) {
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
}

exports.CopyDir = CopyDir
exports.ReplaceTemp = ReplaceTemp
exports.DeleteDir = DeleteDir
exports.GoParseJson = GoParseJson
