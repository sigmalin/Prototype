using System.Collections;
using System.Collections.Generic;
using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.UI;
using Bind.Presenter;
using UI;
using System;
using Services;

namespace Demo.Messagebox
{
    class MessageboxPresenter : UiPresenter
    {
        public override string ViewPath => "Messagebox.prefab";

        public override ViewLayer Layer => ViewLayer.MessageBox;

        string message;

        Text txtMessage;

        ButtonNode buttonNode;

        Action clickHandle;

        public MessageboxPresenter()
        {
            message = string.Empty;

            buttonNode = null;

            clickHandle = null;
        }

        protected override void onBindingCompleted()
        {
            txtMessage = getBindData<Text>("message");
            txtMessage.text = message;

            buttonNode = PresenterMaker.Binding<ButtonNode>(getBindData("button").Target);
            buttonNode.SetClickHandle(onClick);
        }

        void apply()
        {
            if (state == ViewState.WaitBinding) return;

            txtMessage.text = message;
        }

        void onClick()
        {
            clickHandle?.Invoke();

            GameServices.UI.Close<MessageboxPresenter>();
        }

        public void SetContent(string msg, Action callback)
        {
            message = msg;
            clickHandle = callback;

            if (txtMessage != null)
            {
                txtMessage.text = message;
            }  
        }
    }
}
