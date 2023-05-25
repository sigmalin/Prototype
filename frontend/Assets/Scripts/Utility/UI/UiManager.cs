using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using AssetLoader;
using Bind;
using Bind.Presenter;

namespace UI
{
    public class UiManager
    {
        IAssetLoader assetLoader;

        IViewManager viewManager;

        Dictionary<string, UiPresenter> table;
        
        public UiManager()
        {
            assetLoader = null;

            table = new Dictionary<string, UiPresenter>(128);
        }

        public void InjectAssetLoader(IAssetLoader loader)
        {
            assetLoader = loader;
        }

        public void InjectViewLayers(GameObject viewLayers)
        {
            viewManager = PresenterFactory.Binding<UiLayerPresenter>(viewLayers);
        }

        string getPresenterKey(Type type)
        {
            return type.ToString();
        }

        UiPresenter createPresenter<T>() where T : UiPresenter, new()
        {
            UiPresenter presenter = new T();

            if (assetLoader != null)
            {
                IViewLoader viewLoader = presenter;
                viewLoader.Disposable = assetLoader.Load<GameObject>(
                    presenter.ViewPath, 
                    (raw) => {
                        GameObject view = GameObject.Instantiate<GameObject>(raw as GameObject);
                        presenter.Binding(view);
                        viewLoader.onFinishHandle();

                        viewManager?.SetParent(presenter);
                    });
            }

            return presenter;
        }
        
        public T Open<T>() where T : UiPresenter, new()
        {
            UiPresenter presenter = null;

            if (viewManager == null) return null;

            string key = getPresenterKey(typeof(T));

            if (table.TryGetValue(key, out presenter) == false)
            {
                presenter = createPresenter<T>();
                table.Add(key, presenter);
            }

            if (viewManager.Open(presenter) == false)
            {
                Close<T>();
            }

            return (presenter as T);
        }

        public void Close<T>() where T : UiPresenter, new()
        {
            UiPresenter presenter = null;

            string key = getPresenterKey(typeof(T));

            if (table.TryGetValue(key, out presenter) == false)
                return;

            table.Remove(key);

            viewManager.Close(presenter);
        }

        public void ClearLayer(ViewLayer layer)
        {
            viewManager?.Clear(layer);
        }

    }
}

