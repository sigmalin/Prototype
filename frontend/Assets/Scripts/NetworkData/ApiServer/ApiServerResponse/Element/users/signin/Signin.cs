using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Network.WebRequest.Response;
using NetworkData.ApiServer.Model.Element;
using Services;
using UnityEngine;

namespace NetworkData.ApiServer.Element.users
{
    [Serializable]
    public class SigninData
    {
        public string AccessToken;
        public string JsonWebToken;
        public Bank Bank;
    }
}
