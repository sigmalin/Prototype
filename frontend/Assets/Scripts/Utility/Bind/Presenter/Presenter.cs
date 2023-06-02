using System.Collections.Generic;
using UnityEngine;

namespace Bind.Presenter
{
    public abstract class Presenter : IPresnter
    {
        Dictionary<string, BindData> table;

        protected GameObject target { private set; get; }

        protected Presenter()
        {
            table = null;
        }

        void parseBindData(List<BindData> list)
        {
            if (list == null) return;

            table = new Dictionary<string, BindData>(list.Count);

            for (int i = 0; i < list.Count; ++i)
            {
                if (table.ContainsKey(list[i].ID.Value) == true)
                {
                    Debug.LogError($"duplicate identifier {list[i].ID.Value} !!");
                }
                else
                {
                    table.Add(list[i].ID.Value, list[i]);
                }
            }
        }

        protected virtual void release()
        {
            if (target != null)
            {
                GameObject.Destroy(target);
                target = null;
            }

            table.Clear();
        }

        protected BindData getBindData(string key)
        {
            BindData data = null;
            if (table == null || table.TryGetValue(key, out data) == false)
            {
                Debug.LogError($"get identifier {key} failured !!");
            }
            return data;
        }

        protected T getBindData<T>(string key) where T : UnityEngine.Component
        {
            BindData data = getBindData(key);
            if (data == null) return null;
            return data.Target?.GetComponent<T>();
        }

        public bool Binding(GameObject go)
        {
            bool isSuccess = true;

            target = go;

            BindComponent componenet = go?.GetComponent<BindComponent>();
            parseBindData(componenet?.Binds);

            try
            {
                onBindingCompleted();
            }
            catch(System.Exception ex)
            {
                Debug.LogError(ex);

                isSuccess = false;
            }

            return isSuccess;
        }

        protected virtual void onBindingCompleted()
        {

        }
    }
}
