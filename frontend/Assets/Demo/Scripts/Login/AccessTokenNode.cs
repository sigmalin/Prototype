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
    class AccessTokenNode : Presenter
    {
        const string ACCESS_TOKEN_KEY = "access_token";

        Text token;

        public AccessTokenNode()
        {
            token = null;
        }

        protected override void onBindingCompleted()
        {
            token = getBindData<Text>("token");
            SetAccessToken(GetAccessToken());
        }

        public string GetAccessToken()
        {
            return PlayerPrefs.GetString(ACCESS_TOKEN_KEY);
        }

        public void SetAccessToken(string tokenValue)
        {
            PlayerPrefs.SetString(ACCESS_TOKEN_KEY, tokenValue);
            token.text = tokenValue;
            LayoutRebuilder.ForceRebuildLayoutImmediate(target.transform as RectTransform);
        }
    }
}
