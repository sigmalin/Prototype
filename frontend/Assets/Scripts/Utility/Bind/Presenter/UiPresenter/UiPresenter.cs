using UnityEngine;
using UI;
using System;

namespace Bind.Presenter
{
    public abstract class UiPresenter : Presenter, IView, IViewLoader
    {
        public virtual string ViewPath { get; }

        public virtual ViewLayer Layer { get; }

        RectTransform IView.root => (RectTransform)(target?.transform);

        IDisposable IViewLoader.Disposable { set => viewDisposable = value; }

        IDisposable viewDisposable;

        ViewState state;

        public UiPresenter() : base()
        {
            state = ViewState.None;
        }

        protected override void release()
        {
            viewDisposable?.Dispose();
            viewDisposable = null;

            base.release();
        }

        protected virtual void onOpenHandle()
        {
            target?.SetActive(true);
        }

        protected virtual void onHideHandle()
        {
            target?.SetActive(false);
        }

        protected virtual void onCloseHandle()
        {
            release();
        }

        void IView.open()
        {
            state = ViewState.Open;

            if (target != null)
            {
                onOpenHandle();
            }                
        }

        void IView.hide()
        {
            state = ViewState.Hide;

            if (target != null)
            {
                onHideHandle();
            }
        }

        void IView.close()
        {
            state = ViewState.Close;

            if (target != null)
            {
                onCloseHandle();
            }
        }

        void IViewLoader.onFinishHandle()
        {
            switch (state)
            {
                case ViewState.Open:
                    onOpenHandle();
                    break;

                case ViewState.Hide:
                    onHideHandle();
                    break;

                case ViewState.Close:
                    onCloseHandle();
                    break;
            }
        }
    }
}

