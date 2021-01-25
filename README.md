# Service-gen Help Documentation

- Codegen tool for service providers and consumers.

- We will complete the process. First, consumer invokes service. Second, provider monitors request and sends response. Third, consumer gets the response.

- Dependencies:
  - go project: nodejs & go & git
  - java project: nodejs & java v1.8 & maven & curl

- This "Hello world" example uses "node0"(addr: iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv) as the consumer and provider.

## 1. Download
  ```shell
  git clone https://github.com/irisnet/service-gen.git
  ```

## 2. Generate code project.

  - Create schemas.json for your code project.
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

  - Command to build the project: 
    | name | description | default value | parameter value |
    | :-: | :-: | :-: | :-: |
    | type | Generate consumer's or provider's code | | consumer(c) provider(p) |
    | lang | Select language | | go, java, js |
    | service-name(s) | Service's name |  |  |
    | schemas | Path of schemas | ./schemas.json |  |
    | output-dir(o) | Generate path | ../output |  |
  - Example
    ```shell
    node service-gen.js --type consumer --lang go --service-name hello --schemas schemas.json --output-dir ../consumer
    node service-gen.js --type provider --lang go --service-name hello --schemas schemas.json --output-dir ../provider
    ```

## 3. Get ready

  ### 3.1 Key management

  - Commond to key management
    | commond | description |
    | :-: | :-: |
    | add | New-build key |
    | show | Show information of key |
    | import | Import key |
      
    - You need to put the exported information into a file node0.key, and specify the path of the file in config.yaml.

      ##### 3.1.1 Export node0

        ```shell
        iris testnet --v=1 --chain-id=iris -o=/home/sunny/iris
        iris keys export node0 --home /home/sunny/iris/node0/ iriscli
        ```

      ##### 3.1.2 Import node0

        - Example of go
          ```shell
          hello-sc keys import node0
          hello-sp keys import node0
          ```
          
        - Example of java
          
          java -jar target/hello.sc import node0
          java -jar target/hello.sp import node0

  ### 3.2 Callback function
  - The files that need to be modified are on the floder hello.

  - **consumer**
    - Example of go
      ```go
      func ResponseCallback(reqCtxID, reqID, output string) {
        common.Logger.Infof("Get response: %+v\n", output)
        serviceOutput := parseOutput(output)
        // Supplementary service logic...
        fmt.Println(serviceOutput.Output)
      }
      ```
    
    - Example of java
      ```java
      public void onResponse(ServiceOutput res) {
        System.out.println("----------------- Consumer -----------------");
        System.out.println("Got response: "+ JSON.toJSONString(res));
      }
      ```
    - When you get a response from provider, output will appear on terminal.
  
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
        fmt.Println(serviceInput.Input)
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

    - When you get a request, input will appear on terminal, and you will give a word "hello-world" to consumer.
  
  - Compile your project.

  ### 3.3 Config
  - Note!!!: The configuration file is in the $HOME/.hello-sc for consumer and $HOME/.hello-sp for provider.

  - Configuration parameter:
    | name | description |
    | :-: | :-: |
    | chain_id | Chain id |
    | node_rpc_addr | Node URL |
    | node_grpc_addr | Node GRPC address |
    | key_path | Key path |
    | key_name | Key name |
    | fee | Transaction fee |
    | key_algorithm | Key algorithm |
  
  - Example
    ```yaml
    chain_id: iris
    node_rpc_addr: http://localhost:26657
    node_grpc_addr: http://localhost:9090
    key_path: .keys
    key_name: node0
    fee: 4point
    key_algorithm: sm2
    ```

## 4. Start irisnet.

    ```shell
    iris start --home=/home/sunny/iris/node0/iris
    ```

## 5. Define service
  - Open another terminal.
    ```shell
    iris tx service define \
      --name=hello \
      --description=test \
      --author-description=test \
      --tags=test \
      --schemas=/mnt/d/gocode/src/now/test.json \
      --from=node0 \
      --chain-id=iris \
      -b=block -y \
      --home=/home/sunny/iris/node0/iriscli \
      --fees 10point \
    ```

## 6. Bind service

  ```shell
    iris tx service bind \
      --service-name=hello \
      --deposit=10000point \
      --pricing='{"price":"1point"}' \
      --qos=50 \
      --from=node0 \
      --chain-id=iris \
      -b=block -y \
      --home=/home/sunny/iris/node0/iriscli \
      --options={} \
      --fees 10point \
      --provider=iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
  ```

## 7. Start consumer's subscribe response and provider's subscribe request.
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
    - Example of go
      ```shell
      hello-sc invoke \
        --providers iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
        --fee-cap 1 \
        --input {"header":{},"body":{"input":"hello"}} \
      ```
    
    - Example of java
      ```shell
      java -jar target/hello-sc.jar invoke \
        --providers iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
        --fee-cap 1 \
        --input {"header":{},"body":{"input":"hello"}} \
      ```
