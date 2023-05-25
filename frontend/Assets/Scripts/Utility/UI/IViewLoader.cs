using System;

namespace UI
{
    internal interface IViewLoader
    {
        IDisposable Disposable { set; }

        void onFinishHandle();
    }
}
