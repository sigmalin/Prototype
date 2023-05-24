using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AssetLoader
{
    class AssetLoaderSingletonFactory
    {
        static Dictionary<AssetLoaderType, IAssetLoader> table = new Dictionary<AssetLoaderType, IAssetLoader>();

        static public IAssetLoader GetAssetLoader(AssetLoaderType key)
        {
            IAssetLoader loader = null;
            if (table.TryGetValue(key, out loader)) return loader;

            loader = createAssetLoader(key);
            if (loader != null)
            {
                table.Add(key, loader);
            }

            return loader;
        }

        static IAssetLoader createAssetLoader(AssetLoaderType key)
        {
            IAssetLoader loader = null;

            switch (key)
            {
                case AssetLoaderType.Addressable:
                    loader = new AddressableLoader();
                    break;
            }

            return loader;
        }
    }

    enum AssetLoaderType
    {
        Addressable,
    }
}
