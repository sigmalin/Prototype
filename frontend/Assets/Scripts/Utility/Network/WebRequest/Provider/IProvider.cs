using System.Collections;
using System.Collections.Generic;
using System;
using UnityEngine;
using System.Threading.Tasks;
using Network.WebRequest.Response;

namespace Network.WebRequest.Provider
{
    public interface IProvider
    {
        Task<Tuple<Result, string>> Get(string api, Dictionary<string, string> header);
        Task<Tuple<Result, string>> Post(string api, Dictionary<string, string> field, Dictionary<string, string> header);
    }
}

