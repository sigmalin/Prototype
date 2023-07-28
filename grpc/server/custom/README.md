## Prototype - gRPC

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- 使用自訂的 Docker image 產生 golang 代碼

- 可自由決定 protobuf 和 code generator 版本

<br><br>

# 操作方式

- 第一次執行時，需要 build docker image

```console
sh build.sh
```

- 將 proto 檔案放在 src 資料夾內，然後執行指令

```console
sh generator.sh
```

- 會將自動生成的 golang 代碼儲存在 grpc 資料夾