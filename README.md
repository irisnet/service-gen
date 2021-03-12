# Service-gen Help Documentation

- Codegen tool for service providers and consumers.

- Consumer invokes service, Provider sends response, Consumer gets the response.

- Requires:
  - [Node](https://nodejs.org/) 

- This "Hello-world" example uses "node0"(addr: iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv) as the consumer and provider.

## 1. Download

```shell
git clone https://github.com/irisnet/service-gen.git
```

## 2. Generate code project.

  - Prepare schemas.json for your code project.
    - Example
        ```json
        {
          "input": {
            "type": "object",
            "properties": {
              "input": {
                "type": "string"
              }
            }
          },
          "output": {
            "type": "object",
            "properties": {
              "output": {
                "type": "string"
              }
            }
          }
        }
        ```

  - Build parameter: 
  
    | name | description | default value | parameter value |
    | :----: | :----: | :----: | :----: |
    | type | Generate consumer's or provider's code | | consumer(c) provider(p) |
    | lang | Select language | | go, java, js |
    | service-name(s) | Service's name |  |  |
    | schemas | Path of schemas | ./schemas.json |  |
    | output-dir(o) | Generate path | ../output |  |
  
  - Example
    ```shell
    npm install
    node service-gen.js --type consumer --lang go --service-name hello --schemas schemas.json --output ./consumer
    node service-gen.js --type provider --lang go --service-name hello --schemas schemas.json --output ./provider
    ```

## 3. Get ready

  ### 3.1 Config
  - The configuration file is in the config folder.

  - Configuration parameter:
  
    | name | description |
    | :----: | :----: |
    | chain_id | Chain id |
    | node_rpc_addr | Node URL |
    | node_grpc_addr | Node GRPC address |
    | key_path | Key path |
    | key_name | Key name |
    | fee | Transaction fee |
    | key_algorithm | Key algorithm |
  
  - Example
    ```yaml
    service:
        chain_id: irishub-1
        node_rpc_addr: http://localhost:26657
        node_grpc_addr: http://localhost:9090
        key_path: .keys
        key_name: node0
        fee: 40000uiris
        key_algorithm: secp256k1
    ```

  ### 3.2 Callback function
  - The files that need to be modified are on the floder hello.

  - **consumer**
    - Example of go
      ```go
      func ResponseCallback(reqCtxID, reqID, output string) {
      	common.Logger.Infof("Get response: %+v\n", output)
      	serviceOutput := parseOutput(output)
      	// Supplementary service logic...
      	common.Logger.Info(serviceOutput.Output)
      }
      ```
    
    - Example of java
      ```java
      public void onResponse(ServiceOutput res) {
        System.out.println("----------------- Consumer -----------------");
        System.out.println("Got response: "+ JSON.toJSONString(res));
      }
      ```
      
    - When consumer get a response from provider, output will appear on terminal.
  
  - **provider**
    - Example of go
      ```go
      func RequestCallback(reqID, input string) (
      	output *types.ServiceOutput,
      	requestResult *types.RequestResult,
      ) {
      	serviceInput, err := parseInput(input)
      	if err != nil {
      		requestResult = &types.RequestResult{
      			State:   types.ClientError,
      			Message: "failed to parse input",
      		}
      		return nil, requestResult
      	}
      	// Supplementary service logic...
      	common.Logger.Info(serviceInput.Input)
      	var o string = "hello-world"
      	output = &types.ServiceOutput{
      		Output: &o,
      	}
      	requestResult = &types.RequestResult{
      		State:   types.Success,
      		Message: "success",
      	}
      	return output, requestResult
      }
      ```
    
    - Example of java
      ```java
      public ServiceResponse onRequest(ServiceInput req) {
        System.out.println("----------------- Provider -----------------");
        Application.logger.info("Got request:");
        Application.logger.info(JSON.toJSONString(req));

        ServiceOutput serviceOutput = new ServiceOutput();
        serviceOutput.output = "hello-world";

        Application.logger.info("Sending response");
        ServiceResponse res = new ServiceResponse(this.keyName, this.password);
        res.setBody(output);
        
        return res;
      }
      ```

    - When provider get a request, input will appear on provider's terminal, and provider will give a word "hello-world" to consumer.
  
  ### 3.3 Compile project.
  

## 4. [Start irisnet](https://www.irisnet.org/docs/get-started/testnet.html#public-endpoints).

## 5. Key management
    
  - Key parameter:
  
    | commond | description |
    | :----: | :----: |
    | add | New-build key |
    | show | Show information of key |
    | import | Import key |
      
    - You need to put the exported information into a file node0.key, and specify the path of the file in config.yaml.

      ### 5.1 Export node0

        ```shell
        iris keys export node0
        ```

      ### 5.2 Import node0

        - Example of go
          ```shell
          hello-sc keys import node0 node0.key
          hello-sp keys import node0 node0.key
          ```
          
        - Example of java
          ```shell
          java -jar target/hello.sc import node0
          java -jar target/hello.sp import node0
          ```


## 6. Define service
```shell
iris tx service define \
  --name=hello \
  --description=test \
  --author-description=test \
  --tags=test \
  --schemas=/home/sunny/test.json \
  --from=node0 \
  --chain-id=irishub-1 \
  -b=block -y \
  --home=mytestnet/node0/iriscli \
  --keyring-backend file \
  --fees 40000uiris 
```

## 7. Bind service

```shell
iris tx service bind \
  --service-name=hello \
  --deposit=100000000000uiris \
  --pricing='{"price":"1uiris"}' \
  --qos=10 \
  --from=node0 \
  --chain-id=irishub-1 \
  -b=block -y \
  --home=mytestnet/node0/iriscli \
  --keyring-backend file \
  --options={} \
  --fees 40000uiris \
  --provider=iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv
  ```

## 8. Start consumer's subscribe response and provider's subscribe request.
  - **provider**(Subscribe service request first.)
    - Example of go
      ```shell
      hello-sp start
      ```
    
    - Example of java
      ```shell
      java -jar target/hello-sp.jar start
      ```

  - **consumer**(Invoke and subscribe service response.)
  - Note!!!: The cost of the fee-cap should be the largest unit!
    - Example of go
      ```shell
      hello-sc invoke \
        --providers iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
        --fee-cap 0.04iris \
        --input '{"header":{},"body":{"input":"hello"}}' \
        --frequency 11 --timeout 11 --total 1
      ```
    
    - Example of java
      ```shell
      java -jar target/hello-sc.jar invoke \
        --providers iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
        --fee-cap 0.04iris \
        --input {"header":{},"body":{"input":"hello"}}
      ```
     
  - When consumer get response, program will stop soon.
