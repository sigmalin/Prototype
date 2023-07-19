using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Network.WebRequest.Response;
using Network.WebRequest.Model;
using Services;
using UnityEngine;

namespace NetworkData.ApiServer.Model.me
{
    [Serializable]
    public class bank
    {
        public long Coin;
        public long Faith;
        public long Gems;
        public long Tresure;
    }

    class BankModel
    {
        bank data;

        public bank Data => data;

        const string api = "me/bank";

        public async Task Update(bool forceUpdate = false)
        {
            if (data != null && forceUpdate == false) return;

            data = null;

            var response = await GameServices.ApiServer.Get<ApiServerResponse<bank>>(api);
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
