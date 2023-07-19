using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Bind.Presenter;
using NetworkData.ApiServer.Model.me;
using UnityEngine;
using UnityEngine.UI;

namespace Demo.Gamebar
{
    class BankNode : Presenter
    {
        BankDataNode coin;
        BankDataNode faith;
        BankDataNode gems;
        BankDataNode treasure;

        public BankNode()
        {
            coin = null;
            faith = null;
            gems = null;
            treasure = null;
        }

        protected override void onBindingCompleted()
        {
            coin = PresenterMaker.Binding<BankDataNode>(getBindData("coinNode").Target);
            faith = PresenterMaker.Binding<BankDataNode>(getBindData("faithNode").Target);
            gems = PresenterMaker.Binding<BankDataNode>(getBindData("gemsNode").Target);
            treasure = PresenterMaker.Binding<BankDataNode>(getBindData("treasureNode").Target);
        }

        public void Update(bank data)
        {
            if (data == null) return;

            coin.SetValue(data.Coin);
            faith.SetValue(data.Faith);
            gems.SetValue(data.Gems);
            treasure.SetValue(data.Tresure);

            LayoutRebuilder.ForceRebuildLayoutImmediate(target.transform as RectTransform);
        }
    }
}
