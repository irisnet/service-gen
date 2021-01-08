const fs = require('fs')
const path = require('path')
const shell = require('child_process')

function mkdirsSync(dirname) {  
  if (fs.existsSync(dirname)) {  
    return true;
  } else {  
      if (mkdirsSync(path.dirname(dirname))) {  
          fs.mkdirSync(dirname);
          return true;
      }  
  }  
}

function _copy(src, dist) {
  var paths = fs.readdirSync(src)
  paths.forEach(function(p) {
    var _src = src + '/' +p;
    var _dist = dist + '/' +p;
    var stat = fs.statSync(_src)
    if(stat.isFile()) {
      fs.writeFileSync(_dist, fs.readFileSync(_src));
    } else if(stat.isDirectory()) {
      CopyDir(_src, _dist)
    }
  })
}

function CopyDir(src,dist){
  var b = fs.existsSync(dist)
  if(!b){
    mkdirsSync(dist);
  }
  _copy(src,dist);
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
  console.log("downloading go-jsonschema...")
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

function JavaParseJson(output_dir, schemas, type) {
  console.log("downloading jsonschema2pojo...")
  shell.execSync("curl -L https://github.com/joelittlejohn/jsonschema2pojo/releases/download/jsonschema2pojo-1.0.2/jsonschema2pojo-1.0.2.tar.gz " + "-o " + output_dir + "/.temp/jsonschema2pojo-1.0.2.tar.gz")
  shell.execSync("tar -zxf " + output_dir + "/.temp/jsonschema2pojo-1.0.2.tar.gz -C " + output_dir + "/.temp/")

  fs.writeFileSync(output_dir + "/.temp/ServiceInput.json", JSON.stringify(schemas.input))
  fs.writeFileSync(output_dir + "/.temp/ServiceOutput.json", JSON.stringify(schemas.output))

  shell.execSync(path.resolve(fs.realpathSync('.'), output_dir + "/.temp/jsonschema2pojo-1.0.2/bin/jsonschema2pojo --source " + output_dir + "/.temp/ServiceInput.json" + " -tv 1.8 -p service." + type + ".types -t " + output_dir + "/src/main/java/"))
  shell.execSync(path.resolve(fs.realpathSync('.'), output_dir + "/.temp/jsonschema2pojo-1.0.2/bin/jsonschema2pojo --source " + output_dir + "/.temp/ServiceOutput.json" + " -tv 1.8 -p service." + type + ".types -t " + output_dir + "/src/main/java/"))
}

var copyFile = function(srcPath, tarPath, cb) {
  var rs = fs.createReadStream(srcPath)
  rs.on('error', function(err) {
    if (err) {
      console.log('read error', srcPath)
    }
    cb && cb(err)
  })

  var ws = fs.createWriteStream(tarPath)
  ws.on('error', function(err) {
    if (err) {
      console.log('write error', tarPath)
    }
    cb && cb(err)
  })
  ws.on('close', function(ex) {
    cb && cb(ex)
  })

  rs.pipe(ws)
}

exports.copyFile = copyFile
exports.CopyDir = CopyDir
exports.ReplaceTemp = ReplaceTemp
exports.DeleteDir = DeleteDir
exports.GoParseJson = GoParseJson
exports.JavaParseJson = JavaParseJson
