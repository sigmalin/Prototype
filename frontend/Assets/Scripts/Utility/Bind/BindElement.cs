using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind
{

    public class BindElement : BindComponent
    {
        public override BindType Type => BindType.Element;

        public override Identifier ID => identifier;

        public override List<BindData> Binds => null;

        [SerializeField]
        Identifier identifier = new Identifier();
    }
}

