using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using NetworkData.ApiServer.Model.Element;

namespace NetworkData.ApiServer.Model
{
    public class Profile
    {
        Bank bank;

        public long Coin => (bank != null) ? bank.Coin : 0;
        public long Faith => (bank != null) ? bank.Faith : 0;
        public long Gems => (bank != null) ? bank.Gems : 0;
        public long Treasure => (bank != null) ? bank.Treasure : 0;


        public void setBank(Bank bank)
        {
            this.bank = bank;
        }
    }
}
