## Prototype - API Server 

<br><br>

# 介紹

- 基本 API Server 

<br><br>
    
# 說明

- 使用 [Go Work](#建立工作區方式) 管理專案

- swagger 文件，參考 [swagger](../swagger/README.md)

<br><br>

# 建立工作區方式

 1.專案子目錄下，所有資料夾都產生 .mod 檔案

 ```console
 go mod init [module]
 ```

 2.到專案根目錄，建立 go.work 檔案

  ```console
 go work init
 ```

 3.更新所有專案子目錄

  ```console
 go use -r ./
 ```

<br><br>

# 專案架構

- 參考 [golang-standards/project-layout](https://github.com/golang-standards/project-layout) 

<br><br>

# TODO

- 玩家資訊、郵件、背包