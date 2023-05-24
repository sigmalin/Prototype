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
        string identifer = string.Empty;

        public string Value => identifer;
    }
}

