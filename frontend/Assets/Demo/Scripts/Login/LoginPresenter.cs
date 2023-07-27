using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Bind.Presenter;
using UI;
using System;
using Services;
using Singleton;
using Network.WebRequest.Response;
using NetworkData.ApiServer;
using NetworkData.ApiServer.Element.users;
using NetworkData.ApiServer.Model;

namespace Demo.Login
{
    public class LoginPresenter : UiPresenter
    {
        public override string ViewPath => "Login.prefab";

        public override ViewLayer Layer => ViewLayer.BackGround;

        const string apiLogin = "users/login";

        const string apiSignin = "users/signin";

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

        void showMessage(string msg, Action callback)
        {
            GameServices.UI.Open<Messagebox.MessageboxPresenter>()
                .SetContent(msg, callback);
        }

        async void logInHandle(string token)
        {
            Dictionary<string, string> field = new Dictionary<string, string>()
            {
                { "token", token },
            };

            var response = await GameServices.ApiServer.Post<ApiServerResponse<LoginData>>(apiLogin, field);
            if (response.IsSuccess() == false)
            {
                showMessage($"Failure! Error Code = {response.code}", null);
                return;
            }

            var model = Singleton<Profile>.Instance;
            model.setBank(response.data.Bank);

            setAuthData(response.data.JsonWebToken);
        }

        async void signInHandle()
        {
            var response = await GameServices.ApiServer.Get<ApiServerResponse<SigninData>>(apiSignin);
            if (response.IsSuccess() == false)
            {
                showMessage($"Failure! Error Code = {response.code}", null);
                return;
            }

            var model = Singleton<Profile>.Instance;
            model.setBank(response.data.Bank);

            setAuthData(response.data.JsonWebToken);

            accessToken.SetAccessToken(response.data.AccessToken);
            setAuthData(response.data.JsonWebToken);
        }

        void setAuthData(string token)
        {
            GameServices.ApiServer.Authorization(token);

            GameServices.UI.Close<LoginPresenter>();
            GameServices.UI.Open<Entrance.EntrancePresenter>();
        }
    }
}
