using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Bind.Presenter;
using UI;
using UnityEngine;
using UnityEngine.UI;

namespace Demo.Login
{
    class TouchAreaNode : Presenter
    {
        Button btnTouch;

        Action touchHandle;

        public TouchAreaNode()
        {
            btnTouch = null;
            touchHandle = null;
        }

        protected override void onBindingCompleted()
        {
            btnTouch = target.GetComponent<Button>();
            btnTouch.onClick.AddListener(onTouch);
        }

        public void setHandle(Action handle)
        {
            touchHandle = handle;
        }

        void onTouch()
        {
            touchHandle?.Invoke();
        }
    }
}
