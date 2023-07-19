using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Network.WebRequest.Model
{
    public interface IModel
    {
        void SaveLocal();
        void LoadLocal();
    }
}
