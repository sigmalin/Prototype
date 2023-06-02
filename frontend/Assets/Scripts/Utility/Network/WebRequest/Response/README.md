# ServerResponse

## 介紹

- 前端處理伺服器封包的統一格式，所有的變形都需要繼承此類別

- 繼承的子類別都需要提供引數 Result 的建構子(可參考 [ApiServerResponse](../../../../NetworkData/ApiServer/ApiServerResponse/ApiServerResponse.cs))

## 內容

- Result : 前端連線狀況

    | 變數名稱 | 功能 |
    |:-:|:--|
    |Success|連線成功，Response 內容有收到伺服器回應|
    |NetError|連線失敗，Requset 沒有成功送到伺服器|