using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Factory;

namespace Network.WebRequest.Provider
{
    class UnityProviderOrder : IOrder<ProviderType>
    {
        public ProviderType Type => ProviderType.UnityProvider;

        public int RetryTimeOut { get; private set; }

        public UnityProviderOrder(int retry)
        {
            RetryTimeOut = retry;
        }
    }
}
