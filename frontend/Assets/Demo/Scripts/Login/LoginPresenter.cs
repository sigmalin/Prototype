using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Bind.Presenter;
using UI;
using System;
using Services;
using Singleton;
using NetworkData.ApiServer.Model.users;

namespace Demo.Login
{
    public class LoginPresenter : UiPresenter
    {
        public override string ViewPath => "Login.prefab";

        public override ViewLayer Layer => ViewLayer.BackGround;

        AccessTokenNode accessToken;
        TouchAreaNode touchArea;

        public LoginPresenter()
        {

        }

        protected override void onBindingCompleted()
        {
            accessToken = PresenterMaker.Binding<AccessTokenNode>(getBindData("accessTokenNode").Target);
            touchArea = PresenterMaker.Binding<TouchAreaNode>(getBindData("touchArea").Target);
            touchArea.setHandle(onTouchScreen);
        }

        void onTouchScreen()
        {
            string token = accessToken.GetAccessToken();
            if (string.IsNullOrEmpty(token))
            {
                signInHandle();
            }
            else
            {
                logInHandle(token);
            }
        }

        async void logInHandle(string token)
        {
            var model = Singleton<LoginModel>.Instance;
            await model.Update(accessToken.GetAccessToken());

            if (model.Data != null)
            {
                setLoginData(model.Data);
            }
        }

        async void signInHandle()
        {
            var model = Singleton<SigninModel>.Instance;
            await model.Update();

            if (model.Data != null)
            {
                accessToken.SetAccessToken(model.Data.Token);
                setLoginData(model.Data.login);
            }
        }

        void setLoginData(login data)
        {
            GameServices.ApiServer.Authorization(data.JWT);

            GameServices.UI.Close<LoginPresenter>();
            GameServices.UI.Open<Entrance.EntrancePresenter>();
        }
    }
}
