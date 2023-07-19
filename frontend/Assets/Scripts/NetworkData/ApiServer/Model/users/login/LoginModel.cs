using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Network.WebRequest.Response;
using Network.WebRequest.Model;
using Services;
using UnityEngine;

namespace NetworkData.ApiServer.Model.users
{
    [Serializable]
    public class login
    {
        public string JWT;
    }

    class LoginModel : IModel
    {
        login data;

        public login Data => data;

        const string api = "users/login";

        public async Task Update(string token)
        {
            data = null;

            Dictionary<string, string> field = new Dictionary<string, string>()
            {
                { "token", token },
            };

            var response = await GameServices.ApiServer.Post<ApiServerResponse<login>>(api, field);
            if (response.Result != Result.Success)
            {
                Debug.LogError($"connect failure, result = {response.Result}");
                return;
            }

            if (response.code != 0)
            {
                Debug.LogError($"server response failure (code = {response.code}), message = {response.message}");
                return;
            }

            data = response.data;
        }

        public void LoadLocal()
        {
        }

        public void SaveLocal()
        {
        }
    }
}
