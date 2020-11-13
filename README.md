# Service-gen Help Documentation

- This "Hello world" example uses "node0"(addr: iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv) as the consumer and provider.

## 1.Generate code project.

  - Create schemas.json on your code project.
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
    | type | Generate consumer's or provider's code | | consumer, provider |
    | lang | Select language | | go, java, js(Only support go for the time being) |
    | service-name | Service's name |  |  |
    | schemas | Path of schemas | ./schemas.json |  |
    | output-dir | Generate path | ../output |  |
  - Example
    ```shell
    ./service-gen.sh consumer go hello ../schemas.json ../consumer
    ./service-gen.sh provider go hello ../schemas.json ../provider
    ```

## 2.Get ready

  #### 2.1 Key management
  - Commond to key management
    | commond | description |
    | :-: | :-: |
    | add | New-build key |
    | show | Show information of key |
    | import | Import key |
    
  - You need to export node0, and import to your consumer's and provider's project.

    ###### 2.1.1 Export node0

      ```shell
      iris testnet --v=1 --chain-id=iris -o=/home/sunny/iris
      iris keys export node0 --home /home/sunny/iris/node0/ iriscli
      ```

    ###### 2.1.2 Import node0

      - Create file node.txt, write export to it.

        ```shell
        hello-sc keys import node0 node0.txt
        hello-sp keys import node0 node0.txt
        ```

  #### 2.2 Callback function of hello-world
  - You need to be in the hello folder under the generated code directory, there is a file response_callback.

  - **consumer**
    - Example of golang
      ```go
      func ResponseCallback(reqCtxID, reqID, output string) {
        common.Logger.Infof("Get response: %+v\n", output)
        serviceOutput := parseOutput(output)
        // Supplementary service logic...
        fmt.Println(serviceOutput.Output)
      }
      ```
    - When you get a response from provider, input will appear on terminal.
  
  - **provider**
    - Example of golang
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
    - When you get a request, input will appear on terminal, and you will give a word "hello-world" to consumer.
  
  - Compile your project.

  #### 2.3 Config of hello-world
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
    # service config
    service:
      chain_id: iris
      node_rpc_addr: http://localhost:26657
      node_grpc_addr: localhost:9090
      key_path: .keys
      key_name: node0
      fee: 4stake
      key_algorithm: sm2
    ```

## 3.Start irisnet.
  ```shell
  iris start --home=/home/sunny/iris/node0/iris
  ```

## 4.Define service
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
      --fees 10stake \
    ```

## 5.Bind service

  ```shell
    iris tx service bind \
      --service-name=hello \
      --deposit=10000stake \
      --pricing='{"price":"1stake"}' \
      --qos=50 \
      --from=node0 \
      --chain-id=iris \
      -b=block -y \
      --home=/home/sunny/iris/node0/iriscli \
      --options={} \
      --fees 10stake \
      --provider=iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
  ```

## 6.Start consumer's subscribe response and provider's subscribe request.
  - **provider**(Subscribe service request first.)
    ```shell
    hello-sp start
    ```

  - **consumer**(Invoke and subscribe service response.)
    ```shell
    hello-sc invoke \
      --providers iaa15e06fun0plgm22x480g23qeptxu44s4r7cuskv \
      --fee-cap 1 \
      --input {"header":{},"body":{"input":"hello"}} \
    ```
