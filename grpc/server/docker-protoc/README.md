## Prototype - gRPC

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- 使用 [gRPC/Protocol Buffer Compiler Containers](https://github.com/namely/docker-protoc) 產生 golang 代碼

- 一個支援生成各種程式語言的 gRPC 代碼生成工具

<br><br>

# 操作方式

- 將 proto 檔案存放在 src 資料夾內，輸入下列指令

```console
sh generator.sh
```

- 會產生 grpc 資料夾，並存放自動生成的 golang 代碼 