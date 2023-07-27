using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace JsonSerializer
{
    public interface IJson
    {
        T Deserialize<T>(string json);
        string Serialize(object obj);
    }
}
