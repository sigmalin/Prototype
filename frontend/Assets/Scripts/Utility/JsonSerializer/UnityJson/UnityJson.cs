using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using UnityEngine;

namespace JsonSerializer
{
    class UnityJson : IJson
    {
        public T Deserialize<T>(string json)
        {
            return JsonUtility.FromJson<T>(json);
        }

        public string Serialize(object obj)
        {
            return JsonUtility.ToJson(obj);
        }
    }
}
