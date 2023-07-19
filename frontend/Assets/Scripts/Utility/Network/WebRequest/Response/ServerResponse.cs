using System.Collections;
using System.Collections.Generic;
using UnityEngine;


namespace Network.WebRequest.Response
{
    public abstract class ServerResponse
    {
        public Result Result { get; private set; }

        public ServerResponse(Result result)
        {
            Result = result;
        }
    }

    public enum Result
    {
        Success,
        NetError,
    }
}

