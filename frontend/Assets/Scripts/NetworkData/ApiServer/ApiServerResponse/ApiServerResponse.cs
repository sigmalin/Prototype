using System.Collections.Generic;
using System;
using Network.WebRequest.Response;

namespace NetworkData.ApiServer
{
    public class ApiServerResponse : ServerResponse
    {
        public Error error;
        public ApiServerResponse(Result result) : base(result)
        {
        }
    }

    [Serializable]
    public class Error
    {
        public int code;
        public string message;
    }
}
