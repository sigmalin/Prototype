using System.Collections.Generic;
using System.Threading.Tasks;
using Network.WebRequest.Provider;
using Network.WebRequest.Response;

namespace Network.WebRequest.Protocol
{
    public interface IProtocol
    {
        Task<T> Get<T>(string api) where T : ServerResponse;

        Task<T> Post<T>(string api, Dictionary<string, string> field) where T : ServerResponse;

        void Inject(IProvider provider);

        void Authorization(string auth);
    }
}

