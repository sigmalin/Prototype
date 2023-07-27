using System.Collections.Generic;
using System;
using Network.WebRequest.Response;

namespace NetworkData.ApiServer
{
    [Serializable]
    public class ApiServerResponse<T> : ServerResponse where T : new()
    {
        public int code;
        public string message;
        public T data;

        public ApiServerResponse(Result result) : base(result)
        {

        }

        public override bool IsSuccess()
        {
            return base.IsSuccess() && (code == 0);
        }
    }
}
