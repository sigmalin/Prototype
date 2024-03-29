## Prototype - Swagger

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- API Server 使用的 Swagger UI Server 設定

<br><br>

# 架構

- 將 Swagger UI Server 獨立， 處理 API Server 的 API 事件

<br><br>

# CORS

- 因為將 Swagger UI Server 獨立，從 Client 端透過 Swagger UI Server 發送的 API 到 API Sertver 會有跨網域問題
- 需透過 nginx 處理 CORS

<br><br>

# Swagger UI 文件生成

- 使用 [Swaggo/swag](https://github.com/swaggo/swag)

- 安裝方式 (以 Api Server 為例，使用 VSCODE 編輯)


1.在 Terminal 內，切換目錄到 cmd (main.go 在 cmd 目錄下)

```console
cd cmd
```

2.安裝 [Swaggo/swag](https://github.com/swaggo/swag)

```console
 go get -u github.com/swaggo/swag/cmd/swag 
 go install github.com/swaggo/swag/cmd/swag
```

3.安裝完畢後，可透過下列指令確認是否完成安裝

```console
swag --version
```

4.回到 Api Server 專案根目錄

```console
cd ../
```
 
5.在 main function 和所有 Api handler 入口，依照 [Swaggo/swag](https://github.com/swaggo/swag) 指定格式，寫入註解


```go
// @title Prototype Api Server
// @version 1.0
// @description Standard Api Server

// @contact.name sigma
// @contact.url https://github.com/sigmalin/Prototype

// @host 127.0.0.1:443
// @schemes https
func main() {
}
```

6.輸入下列指令，產生 Swagger docs 到 backend/swagger/ 目錄

```console
swag init -g cmd/main.go -o ../swagger/doc/
```

<br><br>

# Swagger UI 代碼生成

- 使用 [Swagger Codegen](https://github.com/swagger-api/swagger-codegen)

- 使用 docker image 進行轉換

- 轉換後的檔案，存放在 convert 資料夾內

| SHELL 檔案名稱 | 功能 |
|:-:|:--|
|openapi|將 [Swaggo/swag](https://github.com/swaggo/swag) 產生的 openAPI 2.0 轉換成 openAPI 3.0 格式|
|csharp|自動生成 csharp client端檔案|

執行指令

```console
sh openapi.sh
```

<br><br>

# TODO

- 目前產生的 csharp 代碼無法用在 prototype frontend(需安裝額外套件 & 產生的變數名稱不明瞭)，故先保留轉換功能，等之後進行擴充
    
