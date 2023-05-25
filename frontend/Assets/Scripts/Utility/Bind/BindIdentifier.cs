using System.Collections;
using System.Collections.Generic;
using System;
using UnityEngine;

namespace Bind
{
    [Serializable]
    public class Identifier
    {
        [SerializeField]
        string identifier = string.Empty;

        public string Value => identifier;
    }
}

