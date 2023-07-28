FROM golang:1.19

# 安裝 protobuf compiler
RUN apt-get update && apt-get install unzip -y
RUN wget -q -O protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip && unzip protoc.zip
RUN mv /go/bin/protoc /usr/local/bin/protoc && rm protoc.zip


# 安裝 golang 相關 plugin
ENV GO111MODULE=on
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /go/src/app
