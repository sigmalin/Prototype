using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind
{
    public class BindNode : BindContainer
    {
        public override BindType Type => BindType.Node;

        public override Identifier ID => identifier;

        [SerializeField]
        Identifier identifier = new Identifier();
    }
}
