using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace UI
{
    class NormalOperator : IViewOperator
    {
        RectTransform root;

        Stack<IView> stack;

        public NormalOperator(RectTransform canvas, int capacity = 16)
        {
            root = canvas;
            stack = new Stack<IView>(capacity);
        }

        IView currentView()
        {
            IView current = null;
            stack.TryPeek(out current);
            return current;
        }

        void setCurrentView(IView view)
        {
            filiterView(view);

            stack.Push(view);
        }

        void filiterView(IView view)
        {
            Stack<IView> tmp = new Stack<IView>();

            IView item = null;
            while (stack.TryPop(out item))
            {
                if (view != item)
                {
                    tmp.Push(item);
                }
            }

            while (tmp.TryPop(out item))
            {
                stack.Push(item);
            }
        }

        public void SetParent(IView view)
        {
            if (view == null) return;
            
            if (view.root != null)
            {
                view.root.SetParent(root);
                view.root.localPosition = Vector3.zero;
                view.root.localScale = Vector3.one;
            }
        }

        public void Open(IView view)
        {
            if (view == null) return;

            IView current = currentView();
            if (view == current) return;

            current?.hide();

            setCurrentView(view);

            view.open();
        }

        public void Close(IView view)
        {
            if (view == null) return;

            view.close();

            IView current = currentView();
            if (view == current)
            {
                stack.Pop();
                currentView()?.open();
            }
            else
            {
                filiterView(view);
            }
        }

        public void Clear()
        {
            IView view = null;
            while (stack.TryPop(out view))
            {
                view.close();
            }
        }
    }
}
