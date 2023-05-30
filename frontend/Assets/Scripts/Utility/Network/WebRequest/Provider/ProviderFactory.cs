using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Factory;

namespace Network.WebRequest.Provider
{
    class ProviderFactory
    {
        public static IProvider Generate(IOrder<ProviderType> order)
        {
            IProvider provider = null;
            switch (order.Type)
            {
                case ProviderType.UnityProvider:
                    provider = new UnityProvider(order as UnityProviderOrder);
                    break;
            }
            return provider;
        }
    }

    public enum ProviderType
    {
        UnityProvider,
    }
}

