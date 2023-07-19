## Prototype - nginx

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- 反向代理伺服 config 存放位置

<br><br>

# 內容

| Port | 功能 |
|:--|:--|
|[80](#port--80)|提供 http 連線，透過子目錄轉接到指定服務|
|443|提供 https 連線，處理 frontend 送出的 API|


<br><br>

# Port : 80
| 子目錄 | 功能 |
|:-:|:--|
|/swagger/|轉接到 swagger ui 介面|
|/api/|處理 swagger ui 送出的 API|