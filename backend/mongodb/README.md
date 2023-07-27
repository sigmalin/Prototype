## Prototype - MongoDB

<br>

語言:

* [English](README_en.md)

---

<br>

# 介紹

- MongoDB 架設與使用說明

<br><br>

# Schema 設計規則

- 由於 MongoDB 每筆 document 大小不能超過 16 MB，因此設計規則為固定長度 Field 放在同一個 Document 裡面，可變長度的 Field 就額外使用 Document 包裝起來，然後使用 ```One-to-Squillions``` 設計模式做關聯

ex. 玩家資訊和玩家遊戲貨幣是固定長度，所以將兩者放入同一個 Document 內

```golang
type bankData struct {
	Coin     uint64 `json:"Coin" bson:"Coin"`
	Faith    uint64 `json:"Faith" bson:"Faith"`
	Gems     uint64 `json:"Gems" bson:"Gems"`
	Treasure uint64 `json:"Treasure" bson:"Treasure"`
}

type usersData struct {
	Token      string   `json:"Token" bson:"Token"`
	Name       string   `json:"Name,omitempty" bson:"Name,omitempty"`
	Mail       string   `json:"Mail,omitempty" bson:"Mail,omitempty"`
	CreateTime int64    `json:"CreateTime" bson:"CreateTime"`
	UpdateTime int64    `json:"UpdateTime" bson:"UpdateTime"`
	Bank       bankData `json:"Bank" bson:"Bank"`
}
```

郵件(或背包)是可變長度，所以額外建立一個 Document 來儲存這些資料

```golang
type itemData struct {
	ID     uint `json:"ID" bson:"ID"`
	Amount uint `json:"Amount" bson:"Amount"`
}

type mailData struct {
	Sender  string     `json:"Sender" bson:"Sender"`
	Status  uint       `json:"Status" bson:"Status"`
	Title   uint       `json:"Title" bson:"Title"`
	Content uint       `json:"Content" bson:"Content"`
	Time    uint64     `json:"Time" bson:"Time"`
	Expired uint64     `json:"Expired" bson:"Expired"`
	Items   []itemData `json:"Items,omitempty" bson:"Items,omitempty"`
}

type mailBoxData struct {
	ID    primitive.ObjectID `json:"-" bson:"_id"`
	Mails []mailData         `json:"Mails,omitempty" bson:"Mails,omitempty"`
}
```

<br><br>

# 常見問題

- 啟動時出現 `Operation not permitted` 錯誤訊息

在 [dockerhub](https://hub.docker.com/_/mongo) 的 `Where to Store Data` 說明裡，有提到

```
WARNING (Windows & OS X): When running the Linux-based MongoDB images on Windows and OS X, the file systems used to share between the host system and the Docker container is not compatible with the memory mapped files used by MongoDB (docs.mongodb.org and related jira.mongodb.org bug). This means that it is not possible to run a MongoDB container with the data directory mapped to the host. To persist data between container restarts, we recommend using a local named volume instead (see docker volume create). Alternatively you can use the Windows-based images on Windows.
```

因此如果使用 virtual box，然後將 mongoDB 的資料儲存在 shared folder，將會出現 `Operation not permitted` 錯誤訊息

解決方法: 使用 docker volume 儲存 mongoDB 資料

建立 volume 指令
```console
docker volume create --name <名稱>
```

檢視 volume 指令
```console
docker volume ls
```

刪除 volume 指令
```console
docker volume rm <名稱>
```

- 為什麼不使用 nginx 做反向代理

因為使用 nginx 做反向代理會發生靜態資源存取失敗(*.css,*.js...)，需要對這些靜態資源做轉向，但目前已提供 swagger ui 做反向代理，如果設定資源轉向，會造成 swagger ui 顯示錯誤