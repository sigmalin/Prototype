using System;
using UnityEngine.ResourceManagement.AsyncOperations;

namespace AssetLoader
{
    public class AddressableDisposable : IDisposable
    {
        public bool isDispose { get; private set; } = false;

        Action<object> onCompleted = null;

        public AddressableDisposable(Action<object> callback)
        {
            onCompleted = callback;
        }

        public void Completed(AsyncOperationHandle handle)
        {
            if (isDispose) return;

            if (handle.Status == AsyncOperationStatus.Succeeded)
            {
                onCompleted?.Invoke(handle.Result);
            }
        }

        public void Dispose()
        {
            isDispose = true;
        }
    }
}
