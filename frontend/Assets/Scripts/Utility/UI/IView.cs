using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace UI
{
    public enum ViewLayer
    {
        BackGround,
        Bar,
        FrontGround,
        MessageBox,
        System,
    }

    public enum ViewState
    {
        None,
        Open,
        Hide,
        Close,
    }

    internal interface IView
    {
        ViewLayer Layer { get; }
        RectTransform root { get; }
        void open();
        void hide();
        void close();
    }

    interface IViewManager
    {
        void SetParent(IView view);
        bool Open(IView view);
        void Close(IView view);
        void Clear(ViewLayer layer);
    }
}
