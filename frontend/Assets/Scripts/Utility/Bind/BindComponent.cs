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
    }

    public enum BindType
    {
        Element,
        Node,
        Container,
    }
}

