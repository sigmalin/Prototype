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
        int timeOut = 5;
        public UnityProvider(UnityProviderOrder order)
        {
            timeOut = order.TimeOut;
        }

        public async Task<Tuple<Result, string>> Get(string api, Dictionary<string, string> header)
        {
            using (UnityWebRequest request = UnityWebRequest.Get(api))
            {
                request.certificateHandler = new SkipCertificate();
                request.SetRequestHeaders(header);
                request.timeout = timeOut;

                await request.SendWebRequest();

                return parseResult(request);
            }                
        }

        public async Task<Tuple<Result, string>> Post(string api, Dictionary<string, string> field, Dictionary<string, string> header)
        {
            WWWForm form = new WWWForm();
            form.SetField(field);

            using (UnityWebRequest request = UnityWebRequest.Post(api, form))
            {
                request.certificateHandler = new SkipCertificate();
                request.SetRequestHeaders(header);
                request.timeout = timeOut;

                await request.SendWebRequest();

                return parseResult(request);
            }
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
        public static void SetRequestHeaders(this UnityWebRequest request, Dictionary<string, string> header)
        {
            if (header == null) return;

            var etor = header.GetEnumerator();
            while (etor.MoveNext())
            {
                request.SetRequestHeader(etor.Current.Key, etor.Current.Value);
            }
        }
    }

    static class WWWFormExtension
    {
        public static void SetField(this WWWForm form, Dictionary<string, string> field)
        {
            if (field == null) return;

            var etor = field.GetEnumerator();
            while (etor.MoveNext())
            {
                form.AddField(etor.Current.Key, etor.Current.Value);
            }
        }
    }

    static class UnityWebRequestAsyncOperationExtension
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

