using System;

namespace AssetLoader
{
    public interface IAssetLoader
    {        
        void UpdateVersion(Action<IDownloadStatus> onProcess, Action onCompleted, Action<Exception> onFailured);
        
        IDisposable Load(string key, Action<Object> callback);
        
        void Release(string key);
    }

    public interface IDownloadStatus
    {
        int TotalFiles { get; }
        int CompleteFiles { get; }
        long DownloadedBytes { get; }
        long FileBytes { get; }
    }
}
