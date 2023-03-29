#用 Docker 包裝 Golang 範本

#Server 的安裝環境 設定安裝環境為 builder
FROM golang:1.19-alpine as builder

#設定中繼影像標籤 之後使用指令 docker image prune --filter label=stage=gobuilder -f 手動移除這些中繼影像
LABEL stage=gobuilder

#Dockerfile 作者
MAINTAINER Sigma

#設定 GOROOT 位置(環境預設值)
ENV GOROOT=/usr/local/go

#複製 golang 自訂 package 到 $GOROOT/src 下
COPY lib/         ${GOROOT}/src

#複製 golang 檔案 到 Container 內的指定目錄 (包含 go.mod(必要), go.sum(根據專案產生))
COPY src/         /src

#執行編譯 並設定編譯參數減少輸出檔案容量 將輸出結果命名為 app
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app 



#設定執行環境
FROM alpine

#開放 Port 1234
EXPOSE 1234

#將編譯結果搬移到執行的指定位置
COPY --from=builder /src/app /

#將 ssl 相關檔案複製到指定位置
COPY ssl/ /ssl

#檔案的執行權限 +x = 任何人都可以執行
CMD chmod +x /app

#執行程式
ENTRYPOINT ./app
