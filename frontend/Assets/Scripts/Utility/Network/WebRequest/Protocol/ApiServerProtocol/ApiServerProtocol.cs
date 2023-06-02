using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using UnityEngine;
using Network.WebRequest.Provider;
using Network.WebRequest.Response;

namespace Network.WebRequest.Protocol
{
    class ApiServerProtocol : IProtocol
    {
        string url;

        IProvider provider;

        public ApiServerProtocol(ApiServerProtocolOrder order)
        {
            url = order.Url;
        }

        public void Inject(IProvider method)
        {
            provider = method;
        }

        public async Task<T> Get<T>(string api) where T : ServerResponse
        {
            var response = await provider.Get($"{url}/{api}");

            return parseResponse<T>(response);
        }

        public async Task<T> Post<T>(string api, Dictionary<string, string> field) where T : ServerResponse
        {
            var response = await provider.Post($"{url}/{api}", field);

            return parseResponse<T>(response);
        }

        T parseResponse<T>(Tuple<Result, string> response) where T : ServerResponse
        {
            if (response.Item1 != Result.Success)
            {
                return (T)Activator.CreateInstance(typeof(T), response.Item1);
            }

            return parseContent<T>(response.Item2);
        }

        T parseContent<T>(string json) where T : ServerResponse
        {
            return JsonUtility.FromJson<T>(json);
        }
    }
}

