## Prototype - frontend

<br><br>

# 介紹

- 資源管理工具，提供資源更新與載入

<br><br>

# 內容

- IAssetLoader :
  
    - info :

        資源載入操作介面

    - functions :

        | 函數名稱 | 內容 |
        |:-:|:-:|
        |UpdateVersion|資源更新檢查
        |Load|資源載入
        |Release|資源釋放

    - example :

    ```csharp
    // 取得資源載入工具
    IAssetLoader assetLoader = Singleton<AddressableLoader>.Instance;
    
    // 檢查檔案更新
    assetLoader.UpdateVersion(onDownLoadProcess, onDownLoadCompleted, onDownLoadFailure);`

    void onDownLoadProcess(IDownloadStatus status)
    {
        // 下載進度更新處理
    }

    void onDownLoadCompleted()
    {
        // 下載完成處理
    }

    void onDownLoadFailure(System.Exception ex)
    {
        // 下載失敗處理
    }
    ```

    ```cs
    // 檔案載入呼叫
    var disposable = assetLoader.Load(AddressNameStr, (raw) => {
            // TODO
        });
    ```


- AddressableDownloadStatus :
  
    - info :

        資源更新進度結構

    - variants :

        | 變數名稱 | 型態 | 內容 |
        |:-:|:-:|:-:|
        |TotalFiles|int|需要下載的檔案數量
        |CompleteFiles|int|完成下載的檔案數量
        |DownloadedBytes|long|目前檔案下載已完成的大小(byte)
        |FileBytes|long|目前檔案下載的大小(byte)
    
