using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.AddressableAssets;
using UnityEngine.AddressableAssets.ResourceLocators;
using UnityEngine.ResourceManagement.AsyncOperations;

namespace AssetLoader
{
    public class AddressableLoader : IAssetLoader
    {
        Dictionary<string, AsyncOperationHandle> resTable;

        public AddressableLoader(int capacity = 128)
        {
            resTable = new Dictionary<string, AsyncOperationHandle>(capacity);
        }

        async public void UpdateVersion(Action<IDownloadStatus> onProcess, Action onCompleted, Action<Exception> onFailured)
        {
            // 功能初始化
            AsyncOperationHandle initHandle = Addressables.InitializeAsync();
            await initHandle.Task;

            // 檢查更新
            var checkHandle = Addressables.CheckForCatalogUpdates(false);
            await checkHandle.Task;

            if (checkHandle.Status != AsyncOperationStatus.Succeeded)
            {
                onFailured?.Invoke(checkHandle.OperationException);
                return;
            }

            AddressableDownloadStatus status = new AddressableDownloadStatus();

            status.TotalFiles = checkHandle.Result.Count;

            if (0 < status.TotalFiles)
            {
                // 取得更新清單
                var updateHandle = Addressables.UpdateCatalogs(checkHandle.Result, true);
                await updateHandle.Task;

                if (updateHandle.Status != AsyncOperationStatus.Succeeded)
                {
                    onFailured?.Invoke(updateHandle.OperationException);
                    return;
                }

                List<IResourceLocator>.Enumerator etor = updateHandle.Result.GetEnumerator();
                while (etor.MoveNext())
                {
                    List<object> keys = new List<object>();
                    keys.AddRange(etor.Current.Keys);

                    // 取得下載文件總大小
                    var sizeHandle = Addressables.GetDownloadSizeAsync(keys.GetEnumerator());
                    await sizeHandle.Task;
                    if (sizeHandle.Status != AsyncOperationStatus.Succeeded)
                    {
                        onFailured?.Invoke(sizeHandle.OperationException);
                        return;
                    }

                    status.FileBytes = sizeHandle.Result;
                    if (0u < status.FileBytes)
                    {
                        // 檔案下載
                        var downloadHandle = Addressables.DownloadDependenciesAsync(keys, true);
                        while (!downloadHandle.IsDone)
                        {
                            status.DownloadedBytes = downloadHandle.GetDownloadStatus().DownloadedBytes;
                            onProcess?.Invoke(status);
                            await Task.Delay(TimeSpan.FromSeconds(Time.deltaTime));
                        }

                        if (downloadHandle.Status != AsyncOperationStatus.Succeeded)
                        {
                            onFailured?.Invoke(downloadHandle.OperationException);
                            return;
                        }

                        onProcess?.Invoke(status);
                    }

                    ++status.CompleteFiles;
                }
            }

            onCompleted?.Invoke();
        }

        public IDisposable Load(string key, Action<object> callback)
        {
            AsyncOperationHandle handle;
            if (resTable.TryGetValue(key, out handle))
            {
                if (handle.IsDone)
                {
                    callback?.Invoke(handle.Result);
                    return null;
                }
            }
            else
            {
                handle = Addressables.LoadAssetAsync<GameObject>(key);
                resTable.Add(key, handle);
            }

            AddressableDisposable disposable = new AddressableDisposable(callback);

            handle.Completed += disposable.Completed;
            return disposable;
        }

        public void Release(string key)
        {
            AsyncOperationHandle handle;
            if (resTable.TryGetValue(key, out handle))
            {
                Addressables.Release(handle);
                resTable.Remove(key);
            }
        }
    }
}
