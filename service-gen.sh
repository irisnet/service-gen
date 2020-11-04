#! /bin/bash
# 1 接收用户输入的参数
type=$1
lang=$2
service_name=$3
schema=$4
output_dir=$5
if [ x"$type" != x"provider" ];then
  echo "目前只支持 provider"
  exit 8
fi
if [ x"$lang" != x"go" ];then
  echo "目前只支持 go 语言"
  exit 8
fi

# 记录 genTool.sh 的绝对路径
genToolPath=$(pwd)

# 记录 schema 的绝对路径
cd ${schema%/*}
schema=${schema##*/}
schema_path=$(pwd)
cd $genToolPath

# 记录 output_dir 的绝对路径
if [ "${output_dir##*/}" = "" ];then
  cd $output_dir
else
  cd ${output_dir%/*}
  mkdir ${output_dir##*/}
  cd ${output_dir##*/}
fi
output_dir=$(pwd)
cd $genToolPath

echo "Complete initialization."

# 2 复制指定模板到指定项目路径
cd templates
cd $type
cd $lang
cp -r * $output_dir

echo "Copy complete"

# 3 修改文件夹名、包名、模板变量
# 进入项目目录
cd $output_dir

# 修改文件夹名
mv {{service_name}} $service_name

# 修改 app.go 中的服务名
cd app
sed -i 's/{{service_name}}/'${service_name}'/g' app.go

# 修改 root.go 中的服务名
cd ../cmd
sed -i 's/{{service_name}}/'${service_name}'/g' root.go

# 修改 start.go 中的服务名
sed -i 's/{{service_name}}/'${service_name}'/g' start.go

# 修改 cbHandler.go 中的服务名
cd ../service
sed -i 's/{{service_name}}/'${service_name}'/g' cbHandler.go

# 修改 serviceCallback.go 中的服务名
cd ../$service_name
sed -i 's/{{service_name}}/'${service_name}'/g' serviceCallback.go

# 修改 types.go 中的服务名
cd ../types
sed -i 's/{{service_name}}/'${service_name}'/g' types.go

# 修改 Makefile 中的服务名
cd ..
sed -i 's/{{service_name}}/'${service_name}'/g' Makefile

echo "Complete the modification."

# 4 安装转换程序，读取 schema，转换为对应的语言结构
cd $output_dir

# 创建临时文件
mkdir .temp
cd .temp
touch ServiceInput.json
touch ServiceOutput.json

if [ "$lang"x = "go"x ];then
  go env -w GOPROXY=https://goproxy.cn
  go get github.com/atombender/go-jsonschema/...
  go build github.com/atombender/go-jsonschema/cmd/gojsonschema

  echo "Parsing JSON..."
  # 解析 input 存入临时文件
  json=$(cat "$schema_path/$schema")
  input=${json%%\"output\"*}
  input=${input%,*}
  echo ${input#*:} > ServiceInput.json

  # 解析 output 存入临时文件
  output=${json%\}*}
  echo ${output#*\"output\":} > ServiceOutput.json

  # 转换为对应类
  input=$(./gojsonschema -p types ServiceInput.json)
  output=$(./gojsonschema -p types ServiceOutput.json)
  
  cd $output_dir/types

  # 把类写入到文件中
  touch input.go
  touch output.go
  echo "$input" >> input.go
  echo "$output" >> output.go

  echo "Parsing JSON complete."

elif [ "$lang"x = "java"x ];then
  echo $lang
else
  echo $lang
fi

# 清理临时文件
cd $output_dir
rm -rf .temp

# 5 使用 curl 安装项目依赖
echo "Installing project dependencies..."

if [ "$lang"x = "go"x ];then
  go mod tidy
elif [ "$lang"x = "java"x ];then
  echo $lang
else
  echo $lang
fi

echo "Complete installation project dependencies"
