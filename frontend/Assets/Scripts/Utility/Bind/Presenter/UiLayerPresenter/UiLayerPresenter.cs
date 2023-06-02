using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UI;

namespace Bind.Presenter
{
    public class UiLayerPresenter : Presenter, IViewManager
    {
        Dictionary<ViewLayer, IViewLayerOperator> layers;

        public UiLayerPresenter() : base()
        {
            layers = new Dictionary<ViewLayer, IViewLayerOperator>();
        }

        IViewLayerOperator getViewLayerOperator(ViewLayer layer)
        {
            IViewLayerOperator viewLayerOperator = null;
            layers.TryGetValue(layer, out viewLayerOperator);
            return viewLayerOperator;
        }

        void parseLayer()
        {
            var enumNames = Enum.GetNames(typeof(ViewLayer));
            var enumValues = Enum.GetValues(typeof(ViewLayer)).Cast<ViewLayer>().ToList();
            for (int i = 0; i < enumNames.Length; ++i)
            {
                var data = getBindData(enumNames[i]);
                if (data != null)
                {
                    layers.Add(enumValues[i], new NormalOperator((RectTransform)data.Target.transform));
                }
            }
        }

        protected override void onBindingCompleted()
        {
            parseLayer();
        }

        void IViewManager.SetParent(IView view)
        {
            IViewLayerOperator viewLayerOperator = getViewLayerOperator(view.Layer);            
            if (viewLayerOperator != null) viewLayerOperator.SetParent(view);
        }

        bool IViewManager.Open(IView view)
        {
            if (view.state == ViewState.WaitBinding) return true;
            
            IViewLayerOperator viewLayerOperator = getViewLayerOperator(view.Layer);
            viewLayerOperator?.Open(view);

            return viewLayerOperator != null;
        }

        void IViewManager.Close(IView view)
        {
            IViewLayerOperator viewLayerOperator = getViewLayerOperator(view.Layer);
            viewLayerOperator?.Close(view);
        }

        void IViewManager.Clear(ViewLayer layer)
        {
            IViewLayerOperator viewLayerOperator = getViewLayerOperator(layer);
            viewLayerOperator?.Clear();
        }
    }
}

