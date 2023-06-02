# WebRequest

## 介紹

- Request-response 架構通訊工具

## 原理

- 實現 Request-response 架構通訊工具
- 將功能分成 3 個部分， Protocol、Provider 和 Response 
- Protocol、Provider 分別使用 [ProtocolFactory](./Protocol/ProtocolFactory.cs)、[ProviderFactory](./Provider/ProviderFactory.cs) 產生實體
- Response 根據伺服器回傳內容，自訂繼承 [ServerResponse](./Response/README.md) 結構

## 內容

- 功能說明 :

    | component | 功能 |
    |:-:|:--|
    |Protocol|負責處理對應不特定伺服器的回傳封包(ex.解碼、json 轉換)，轉換成前端處理的統一格式(繼承 [ServerResponse](./Response/README.md) 結構)|
    |Provider|負責處理向伺服器發送 Requst 的功能，目前提供 Get、Post|
    |Response|前端處理的統一格式，需要繼承 [ServerResponse](./Response/README.md) 結構|



- 操作方法 :

    ```cs
    // 宣告方式
    // 宣告連線 Prototype ApiServer
    var apiServer = ProtocolFactory.Generate(new ApiServerProtocolOrder("https://127.0.0.1:1234"));
    // 使用指定的 Request 發送方式
    apiServer.Inject(ProviderFactory.Generate(new UnityProviderOrder(5)));
    ```

    ```cs    
    // 向伺服器發送 request
    // response 內容可參考 Prototype ApiServer
    var response = await apiServer.Get<ApiServerResponse>("");
    ```

- 相關連結 :

    - [ApiServerProtocolOrder](./Protocol/ApiServerProtocol/README.md)
    - [UnityProviderOrder](./Provider/UnityProvider/README.md)
    - [ApiServerResponse](../../../NetworkData/ApiServer/ApiServerResponse/ApiServerResponse.cs)
