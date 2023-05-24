using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind
{

    public class BindElement : BindComponent
    {
        public override BindType Type => BindType.Element;

        public override Identifier ID => identifier;

        [SerializeField]
        Identifier identifier = new Identifier();
    }
}

