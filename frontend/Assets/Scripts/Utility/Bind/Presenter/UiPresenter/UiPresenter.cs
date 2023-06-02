using UnityEngine;
using UI;
using System;
using System.Collections.Generic;

namespace Bind.Presenter
{
    public abstract class UiPresenter : Presenter, IView, IViewLoader
    {
        public virtual string ViewPath { get; }

        public virtual ViewLayer Layer { get; }

        public ViewState state => viewState;

        RectTransform IView.root => (RectTransform)(target?.transform);

        List<IDisposable> disposables;

        ViewState viewState;

        public UiPresenter() : base()
        {
            viewState = ViewState.WaitBinding;

            disposables = new List<IDisposable>();
        }

        void clearDisposables()
        {
            for (int i = 0; i < disposables.Count; ++i)
            {
                disposables[i].Dispose();
            }

            disposables.Clear();
        }

        protected override void release()
        {
            clearDisposables();

            base.release();
        }

        public void AddDisposable(IDisposable disposable)
        {
            if (disposable == null) return;

            disposables.Add(disposable);
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
            viewState = ViewState.Open;

            if (target != null)
            {
                onOpenHandle();
            }                
        }

        void IView.hide()
        {
            viewState = ViewState.Hide;

            if (target != null)
            {
                onHideHandle();
            }
        }

        void IView.close()
        {
            viewState = ViewState.Close;

            if (target != null)
            {
                onCloseHandle();
            }
        }

        void IViewLoader.finish()
        {
            viewState = ViewState.BingCompleted;
        }
    }
}

