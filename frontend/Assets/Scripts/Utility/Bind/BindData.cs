using System.Collections;
using System.Collections.Generic;
using System;
using UnityEngine;

namespace Bind
{
    [Serializable]
    public class BindData
    {
        [SerializeField]
        private Identifier identifier = null;

        [SerializeField]
        private GameObject gameObject = null;

        [SerializeField]
        private BindType type = BindType.None;

        public Identifier ID => identifier;
        public GameObject Target => gameObject;
        public BindType Type => type;

        public BindData(BindComponent component)
        {
            if (component == null) return;

            identifier = component.ID;
            gameObject = component.gameObject;
            type = component.Type;
        }
    }
}
