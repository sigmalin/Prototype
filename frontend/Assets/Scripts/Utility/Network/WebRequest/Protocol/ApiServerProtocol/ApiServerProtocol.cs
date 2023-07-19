using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using UnityEngine;
using Network.WebRequest.Provider;
using Network.WebRequest.Response;
using NetworkData.ApiServer;

namespace Network.WebRequest.Protocol
{
    class ApiServerProtocol : IProtocol
    {
        const string AUTH_KEY = "Authorization";

        string url;

        Dictionary<string, string> header;

        IProvider provider;

        public ApiServerProtocol(ApiServerProtocolOrder order)
        {
            url = order.Url;
            header = new Dictionary<string, string>();
        }

        public void Inject(IProvider method)
        {
            provider = method;
        }

        public void Authorization(string auth)
        {
            if (header.ContainsKey(AUTH_KEY))
            {
                header[AUTH_KEY] = $"Bearer {auth}";
            }
            else
            {
                header.Add(AUTH_KEY, $"Bearer {auth}");
            }
        }

        public async Task<T> Get<T>(string api) where T : ServerResponse
        {
            var response = await provider.Get($"{url}/{api}", header);

            return parseResponse<T>(api, response);
        }

        public async Task<T> Post<T>(string api, Dictionary<string, string> field) where T : ServerResponse
        {
            var response = await provider.Post($"{url}/{api}", field, header);            
            return parseResponse<T>(api, response);
        }

        T parseResponse<T>(string api, Tuple<Result, string> response) where T : ServerResponse
        {
            if (response.Item1 != Result.Success)
            {
                Debug.LogError($"recv api:{api} failure!");

                return (T)Activator.CreateInstance(typeof(T), response.Item1);
            }

            Debug.Log($"recv api:{api} = {response.Item2}");

            return parseContent<T>(response.Item2);
        }

        T parseContent<T>(string json) where T : ServerResponse
        {
            return JsonUtility.FromJson<T>(json);
        }
    }
}

