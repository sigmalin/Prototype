using System.Collections.Generic;
using System.Runtime.CompilerServices;
using System.Threading.Tasks;
using System;
using UnityEngine.Networking;
using UnityEngine;
using Network.WebRequest.Response;


namespace Network.WebRequest.Provider
{
    class UnityProvider : IProvider
    {
        int retryTimeOut;

        SkipCertificate skipCertificate;

        public UnityProvider(UnityProviderOrder order)
        {
            retryTimeOut = order.RetryTimeOut;

            skipCertificate = new SkipCertificate();
        }

        public async Task<Tuple<Result, string>> Get(string api)
        {
            UnityWebRequest request = UnityWebRequest.Get(api);
            request.certificateHandler = skipCertificate;

            request.timeout = retryTimeOut;

            await request.SendWebRequest();

            return parseResult(request);
        }

        public async Task<Tuple<Result, string>> Post(string api, Dictionary<string, string> field)
        {
            WWWForm form = new WWWForm();

            var etor = field.GetEnumerator();
            while (etor.MoveNext())
            {
                field.Add(etor.Current.Key, etor.Current.Value);
            }

            UnityWebRequest request = UnityWebRequest.Post(api, form);
            request.certificateHandler = skipCertificate;

            request.timeout = retryTimeOut;

            await request.SendWebRequest();

            return parseResult(request);
        }

        Tuple<Result, string> parseResult(UnityWebRequest request)
        {
            Tuple<Result, string> response = null;

            if (request.result != UnityWebRequest.Result.Success)
            {
                response = Tuple.Create<Result, string>(Result.NetError, request.error);
            }
            else
            {
                response = Tuple.Create<Result, string>(Result.Success, request.downloadHandler.text);
            }

            return response;
        }
    }

    class SkipCertificate : CertificateHandler
    {
        protected override bool ValidateCertificate(byte[] certificateData)
        {
            return true;
        }
    }

    static class UnityWebRequestExtension
    {
        public static TaskAwaiter<object> GetAwaiter(this UnityWebRequestAsyncOperation asyncop)
        {
            var taskCompletionSource = new TaskCompletionSource<object>();
            asyncop.completed += (obj) =>
            {
                taskCompletionSource.SetResult(null);
            };
            return taskCompletionSource.Task.GetAwaiter();
        }
    }
}

