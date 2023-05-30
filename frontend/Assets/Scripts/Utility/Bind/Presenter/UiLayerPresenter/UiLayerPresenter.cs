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
        Dictionary<ViewLayer, IViewOperator> layers;

        public UiLayerPresenter() : base()
        {
            layers = new Dictionary<ViewLayer, IViewOperator>();
        }

        IViewOperator getViewLayer(ViewLayer layer)
        {
            IViewOperator viewOperator = null;
            layers.TryGetValue(layer, out viewOperator);
            return viewOperator;
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
            IViewOperator viewOperator = getViewLayer(view.Layer);            
            if (viewOperator != null) viewOperator.SetParent(view);
        }

        bool IViewManager.Open(IView view)
        {
            IViewOperator viewOperator = getViewLayer(view.Layer);
            viewOperator?.Open(view);

            return viewOperator != null;
        }

        void IViewManager.Close(IView view)
        {
            IViewOperator viewOperator = getViewLayer(view.Layer);
            viewOperator?.Close(view);
        }

        void IViewManager.Clear(ViewLayer layer)
        {
            IViewOperator viewOperator = getViewLayer(layer);
            viewOperator?.Clear();
        }
    }
}

