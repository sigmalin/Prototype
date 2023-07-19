using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Bind.Presenter;
using UI;
using UnityEngine;
using UnityEngine.UI;

namespace Demo.Messagebox
{
    class ButtonNode : Presenter
    {
        Text text;

        Button btn;

        Action clickHandle;

        public ButtonNode()
        {
            text = null;
        }

        protected override void onBindingCompleted()
        {
            text = getBindData<Text>("text");

            btn = target.GetComponent<Button>();
            btn.onClick.AddListener(onClick);
        }

        void onClick()
        {
            clickHandle?.Invoke();
        }

        public void SetClickHandle(Action callback)
        {
            clickHandle = callback;
        }

        public void SetText(string text)
        {
            this.text.text = text;
        }
    }
}
