## Prototype - backend

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- 後端架構，提供遊戲開發所需的後端功能

<br><br>

# 內容

| 資料夾 | 功能 |
|:--|:--|
|api|建構 Api Server 相關代碼|
|bundles|提供 frontend 下載資源的存放位置|
|file|建構 FTP Server 相關代碼，提供 frontend 下載 bundles 資料夾內的檔案服務|
|nginx|反向代理伺服 config 存放位置，解決 cors 問題所設定|
|sql|database 資料(data) 和 欄位設計(schema) 的存放位置|
|ssl|ssl 金鑰的存放位置|
|swagger|swagger 腳本存放位置|

<br><br>

# 使用方式

- 安裝 docker

- 在目錄下輸入 ``` docker-compose up -d ```

<br><br>

# 環境變數 (docker-compose.yml)

- go_api

| 資料夾 | 功能 |
|:--|:--|
|API_PORT|Api Server 開放的 Port|
|SSL_CERTIFICATION|ssl 憑證位置|
|SSL_PRIVATE_KEY|ssl 金鑰位置|
|JWT_SIGNING_KEY|Json web token 的簽名字串|
|JWT_CLAIMS_KEY|Json web token 的存取索引|
|SQL_DRIVER|指定的 SQL Driver|
|SQL_USERNAME|SQL 登入帳號|
|SQL_PASSWORD|SQL 登入密碼|
|SQL_ADDRESS|SQL 網址|
|SQL_PORT|SQL 開放的 Port|
|SQL_DATABASE|存取 SQL 的 table 表|
|SQL_TIMEOUT|每筆 SQL 指令允許的最大等待時間|
|SQL_MAXLIFETIME|每個 SQL 連線可以存活的最大時間(秒)|
|SQL_MAXOPENCONNECT|SQL 最大的連線數|
|SQL_MAXIDLECONNECT|SQL 最大的閒置連線數|
|REDIS_CACHE_ADDRESS|Redis cache 網址|
|REDIS_CACHE_PORT|Redis 開放的 Port|
|REDIS_CACHE_EXPIRATION|每個 Cahae 存活時間(秒)|

<br><br>

- go_api

| 資料夾 | 功能 |
|:--|:--|
|FILE_PORT|資源下載的 Port|

<br><br>

- swagger_ui

| 資料夾 | 功能 |
|:--|:--|
|API_URL|swagger 腳本位置|

<br><br>

# TODO

- 建構 Jenkins，提供開發所需的自動化操作
- 建構 gRPC Server，提供強連線服務
