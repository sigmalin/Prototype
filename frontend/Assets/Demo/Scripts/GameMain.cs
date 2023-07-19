using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Services;
using AssetLoader;
using Singleton;
using UI;
using Network.WebRequest.Protocol;
using Network.WebRequest.Provider;

namespace Demo
{
    public class GameMain : MonoBehaviour
    {
        const string apiServer = "https://127.0.0.1:443";

        void Start()
        {
            initServices();
        }

        void initServices()
        {
            initAssetLoader();

            initUI();

            initServerConnector();
        }

        void initAssetLoader()
        {
            GameServices.AssetLoader = Singleton<AddressableLoader>.Instance;
            GameServices.AssetLoader.UpdateVersion(onProcess, onCompleted, onFailured);
        }

        void initUI()
        {
            GameServices.UI = Singleton<UiManager>.Instance;
            GameServices.UI.InjectAssetLoader(GameServices.AssetLoader);
        }

        void initServerConnector()
        {
            GameServices.ApiServer = ProtocolFactory.Generate(new ApiServerProtocolOrder(apiServer));
            GameServices.ApiServer.Inject(ProviderFactory.Generate(new UnityProviderOrder()));
        }

        void onProcess(IDownloadStatus status)
        {

        }

        void onCompleted()
        {
            prepareUiLayout();
        }

        void onFailured(System.Exception ex)
        {

        }

        void prepareUiLayout()
        {
            GameServices.AssetLoader.Load<GameObject>("UiLayout.prefab", setUiLayout);
        }

        void setUiLayout(object raw)
        {
            GameObject viewLayers = GameObject.Instantiate<GameObject>(raw as GameObject);

            GameServices.UI.InjectViewLayers(viewLayers);

            GameServices.UI.Open<Login.LoginPresenter>();
        }
    }
}

