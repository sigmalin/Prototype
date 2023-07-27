using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Bind.Presenter;
using UnityEngine;
using UnityEngine.UI;
using Singleton;
using NetworkData.ApiServer.Model;

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

        public void Update()
        {
            var profile = Singleton<Profile>.Instance;

            coin.SetValue(profile.Coin);
            faith.SetValue(profile.Faith);
            gems.SetValue(profile.Gems);
            treasure.SetValue(profile.Treasure);

            LayoutRebuilder.ForceRebuildLayoutImmediate(target.transform as RectTransform);
        }
    }
}
