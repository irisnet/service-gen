#! /bin/bash
# 1 Receive parameters
type=$1
lang=$2
service_name=$3
schema=$4
output_dir=$5

# Parameter processing
if [ "$type" != "provider" ] && [ "$type" != "consumer" ];then
  echo "Please enter correct type: consumer | provider."
  exit 8
fi

if [ x"$lang" != x"go" ];then
  echo "only supported golang currently"
  exit 8
fi

# Record genTool.sh's absolute path
genToolPath=$(pwd)

# Record schema's absolute path
cd ${schema%/*}
schema=${schema##*/}
schema_path=$(pwd)
cd $genToolPath

# Record output_dir's absolute path
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

# 2 Copy the specified template to the specified project path
cd templates
cd $type
cd $lang
if [ "$type" == "provider" ];then
  mkdir $HOME/.$service_name-sp/
  cp config/config.yaml $HOME/.$service_name-sp/
else
  mkdir $HOME/.$service_name-sc/
  cp config/config.yaml $HOME/.$service_name-sc/
fi
cp -r * $output_dir

echo "Copy complete"

# 3 Modify template variables
cd $output_dir
rm -rf config

# Modify folder name
mv servicename $service_name

# Modify the service name in the app.go
cd app
sed -i 's/servicename/'${service_name}'/g' app.go

# Modify the service name in the root.go
cd ../cmd
sed -i 's/servicename/'${service_name}'/g' root.go

# Modify the service name in the start.go
sed -i 's/servicename/'${service_name}'/g' start.go

# Modify the service name in the config.go
cd ../common
sed -i 's/servicename/'${service_name}'/g' config.go

# Modify the service name in the types.go
cd ../types
sed -i 's/servicename/'${service_name}'/g' types.go

# Modify the service name in the Makefile
cd ..
sed -i 's/servicename/'${service_name}'/g' Makefile

if [ "$type" == "consumer" ];then
  # Modify the service name in the test.go
  cd cmd
  sed -i 's/servicename/'${service_name}'/g' test.go

  # Modify the service name in the response_callback.go
  cd ../$service_name
  sed -i 's/servicename/'${service_name}'/g' response_callback.go
else
  # Modify the service name in the request_callback.go
  cd $service_name
  sed -i 's/servicename/'${service_name}'/g' request_callback.go
fi

echo "Complete the modification."

# 4 Install converter, read schema.json, convert to the corresponding language structure
cd $output_dir

# Create temporary folder
mkdir .temp
cd .temp
touch ServiceInput.json
touch ServiceOutput.json

if [ "$lang"x = "go"x ];then
  go env -w GOPROXY=https://goproxy.cn
  go get github.com/atombender/go-jsonschema/...
  go build github.com/atombender/go-jsonschema/cmd/gojsonschema

  echo "Parsing JSON..."
  # Parsing input
  json=$(cat "$schema_path/$schema")
  input=${json%%\"output\"*}
  input=${input%,*}
  echo ${input#*:} > ServiceInput.json

  # Parsing output
  output=${json%\}*}
  echo ${output#*\"output\":} > ServiceOutput.json

  # Convert to the corresponding structure
  input=$(./gojsonschema -p types ServiceInput.json)
  output=$(./gojsonschema -p types ServiceOutput.json)
  
  cd $output_dir/types

  # Write to file
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

# Clean up temporary files
cd $output_dir
rm -rf .temp

# 5 Installation project dependencies
echo "Installing project dependencies..."

if [ "$lang"x = "go"x ];then
  go mod tidy
elif [ "$lang"x = "java"x ];then
  echo $lang
else
  echo $lang
fi

echo "Complete installation project dependencies"
