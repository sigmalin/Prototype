using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Factory;

namespace Network.WebRequest.Protocol
{
    class ProtocolFactory
    {
        public static IProtocol Generate(IOrder<WebProtocolType> order)
        {
            IProtocol protocol = null;
            switch (order.Type)
            {
                case WebProtocolType.ApiServer:
                    protocol = new ApiServerProtocol(order as ApiServerProtocolOrder);
                    break;
            }
            return protocol;
        }
    }

    public enum WebProtocolType
    {
        ApiServer,
    }
}
