using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Bind.Presenter;
using UI;
using UnityEngine;
using UnityEngine.UI;

namespace Demo.Gamebar
{
    class BankDataNode : Presenter
    {
        Text value;

        public BankDataNode()
        {
            value = null;
        }

        protected override void onBindingCompleted()
        {
            value = getBindData<Text>("value");
        }

        public void SetValue(long value)
        {
            this.value.text = value.ToString("N0");
            LayoutRebuilder.ForceRebuildLayoutImmediate(target.transform as RectTransform);
        }
    }
}
