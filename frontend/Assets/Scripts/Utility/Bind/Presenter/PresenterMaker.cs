using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind.Presenter
{
    public class PresenterMaker
    {
        static public T Binding<T>(GameObject go) where T : IPresnter, new()
        {
            T presenter = new T();

            if (presenter.Binding(go) == false)
            {
                presenter = default(T);
            }

            return presenter;
        }
    }
}
