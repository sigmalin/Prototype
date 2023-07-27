using System.Collections;
using System.Collections.Generic;
using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.UI;
using Bind.Presenter;
using UI;
using System;
using Services;

namespace Demo.Gamebar
{
    class GamebarPresenter : UiPresenter
    {
        public override string ViewPath => "Gamebar.prefab";

        public override ViewLayer Layer => ViewLayer.Bar;

        BankNode bankNode;

        public GamebarPresenter()
        {
            bankNode = null;
        }

        protected override void onBindingCompleted()
        {
            bankNode = PresenterMaker.Binding<BankNode>(getBindData("bankNode").Target);

            bankNode.Update();
        }
    }
}
