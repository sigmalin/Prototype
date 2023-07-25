# Prototype

語言:

* [English](README_en.md)

---

## 介紹

- 一個遊戲前、後端的程式框架
- 目標是提供簡單開發、輕易部屬(測試)的設計環境

## 資料夾

- frontend :
  
    - info :

        前端代碼資料夾，使用 Unity3D 實作

    - logs :

        | 日期 | 更新內容 |
        |:-:|:--|
        |2023/02/23|完成與後端通訊功能範本|
        |2023/06/02|代碼重購, 提供 UI 管理, 資源下載 和 網路通訊功能|
        |2023/07/19|更新與後端聯繫的基本操作 Demo 範本|


- backend :

    - info :

        後端代碼資料夾，使用 golang 實作

    - logs :

        | 日期 | 更新內容 |
        |:-:|:--|
        |2023/02/23|完成基礎 API Server, 提供 sql access, session 功能|
        |2023/07/10|代碼重購, 使用 redis 作為 sql 的快取, 建立程式架構規則|
        |2023/07/17|完成對應 API Server 的 Swagger 環境, 取消 session 功能 使用 json web token 做為認證手段|
        |2023/07/31|代碼重購, 使用 mongoDB 作為玩家資料庫, 建立程式架構規則|
    
