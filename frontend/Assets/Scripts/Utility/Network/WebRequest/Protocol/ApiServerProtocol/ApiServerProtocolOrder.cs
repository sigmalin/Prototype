using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Factory;

namespace Network.WebRequest.Protocol
{
    class ApiServerProtocolOrder : IOrder<WebProtocolType>
    {
        public WebProtocolType Type => WebProtocolType.ApiServer;

        public string Url { get; private set; }

        public ApiServerProtocolOrder(string url)
        {
            Url = url;
        }
    }
}

