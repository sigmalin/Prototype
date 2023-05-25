using UnityEngine;

namespace UI
{
    interface IViewOperator
    {
        void SetParent(IView view);
        void Open(IView view);
        void Close(IView view);
        void Clear();
    }
}
