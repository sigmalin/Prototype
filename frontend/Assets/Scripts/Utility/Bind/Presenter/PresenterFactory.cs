using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind.Presenter
{
    public class PresenterFactory
    {
        static public T Binding<T>(GameObject go) where T : IPresnter, new()
        {
            T presenter = new T();

            presenter.Binding(go);

            return presenter;
        }
    }
}
