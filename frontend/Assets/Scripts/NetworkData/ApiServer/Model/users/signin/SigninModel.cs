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
    public class signin
    {
        public string Token;
        public login login;
    }

    class SigninModel : IModel
    {
        signin data;

        public signin Data => data;

        const string api = "users/signin";

        public async Task Update()
        {
            data = null;

            var response = await GameServices.ApiServer.Get<ApiServerResponse<signin>>(api);
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
