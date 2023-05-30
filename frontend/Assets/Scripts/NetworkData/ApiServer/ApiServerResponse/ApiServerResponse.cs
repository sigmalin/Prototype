using System.Collections.Generic;
using System;
using Network.WebRequest.Response;

namespace NetworkData.ApiServer
{
    public class ApiServerResponse : ServerResponse
    {
        public Error error;
        public ApiServerResponse() : base(Result.Success)
        {
        }
    }

    [Serializable]
    public class Error
    {
        public decimal code;
        public string message;
    }
}
