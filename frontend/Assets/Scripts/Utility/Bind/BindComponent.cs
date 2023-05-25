using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind
{
    [DisallowMultipleComponent]
    public abstract class BindComponent : MonoBehaviour
    {
        public virtual BindType Type { get; }

        public virtual Identifier ID { get; }

        public virtual List<BindData> Binds { get; }
    }

    public enum BindType
    {
        None,
        Element,
        Node,
        Container,
    }
}

