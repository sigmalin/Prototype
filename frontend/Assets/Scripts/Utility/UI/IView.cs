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
        WaitBinding,
        BingCompleted,
        Open,
        Hide,
        Close,
    }

    internal interface IView
    {
        ViewLayer Layer { get; }
        RectTransform root { get; }
        ViewState state { get; }
        void open();
        void hide();
        void close();
    }

    interface IViewManager
    {
        void SetParent(IView view);
        bool Open(IView view);
        void Close(IView view);
    }
}
