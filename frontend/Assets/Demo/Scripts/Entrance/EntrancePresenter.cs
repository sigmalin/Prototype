using System.Collections;
using System.Collections.Generic;
using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.UI;
using Bind.Presenter;
using UI;
using System.Linq;
using Services;
using Singleton;

namespace Demo.Entrance
{
    public class EntrancePresenter : UiPresenter
    {
        public override string ViewPath => "Loading.prefab";

        public override ViewLayer Layer => ViewLayer.System;

        List<UiPresenter> waitPresenters;

        public EntrancePresenter()
        {
            waitPresenters = new List<UiPresenter>();
        }

        protected override void release()
        {
            waitPresenters.Clear();

            waitPresenters = null;

            base.release();
        }

        protected override void onBindingCompleted()
        {
            getUserData();
        }

        async void getUserData()
        {
            openGameBar();

            await waitPresentsBinding();

            GameServices.UI.Close<EntrancePresenter>();
        }

        void openGameBar()
        {
            waitPresenters.Add(
                GameServices.UI.Open<Gamebar.GamebarPresenter>()
                );
        }

        bool isAllPresentersReady()
        {
            return waitPresenters
                .Where(_ => _.state == ViewState.WaitBinding)
                .Any() == false;            
        }

        async Task waitPresentsBinding()
        {
            while (isAllPresentersReady() == false)
                await Task.Yield();
        }

    }
}

