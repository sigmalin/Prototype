# 介紹

- API Server 使用的 Swagger UI Server 設定

# 架構

- 將 Swagger UI Server 獨立， 處理 API Server 的 API 事件

# CORS

- 因為將 Swagger UI Server 獨立，從 Client 端透過 Swagger UI Server 發送的 API 到 API Sertver 會有跨網域問題
- 需透過 nginx 處理 CORS

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

// @host 127.0.0.1:80
// @schemes http
func main() {
}
```

6.輸入下列指令，產生 Swagger docs 到 backend/swagger/ 目錄

```console
swag init -g cmd/main.go -o ../swagger/doc/
```


    
