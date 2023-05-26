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

        RectTransform IView.root => (RectTransform)(target?.transform);

        List<IDisposable> disposables;

        ViewState state;

        public UiPresenter() : base()
        {
            state = ViewState.None;

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

        protected virtual void onOpenHandle()
        {
            target?.SetActive(true);
        }

        public void AddDisposable(IDisposable disposable)
        {
            if (disposable == null) return;

            disposables.Add(disposable);
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

        void IViewLoader.finish()
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

