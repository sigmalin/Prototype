using AssetLoader;
using UI;
using Network.WebRequest.Protocol;

namespace Services
{
    public class GameServices
    {
        static public IAssetLoader AssetLoader { get; set; }

        static public UiManager UI { get; set; }

        static public IProtocol ApiServer { get; set; }
    }    
}