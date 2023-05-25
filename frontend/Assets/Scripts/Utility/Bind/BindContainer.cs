using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Bind
{
    public class BindContainer : BindComponent
    {
        public override BindType Type => BindType.Container;

        public override Identifier ID => null;

        public override List<BindData> Binds => bindings;

        [HideInInspector]
        [SerializeField]        
        List<BindData> bindings = new List<BindData>();

        public void Initialize()
        {
            bindings.Clear();
            traceBindData(this.transform);
        }

        void traceBindData(Transform root)
        {
            if (root == null) return;

            int childCount = root.childCount;

            for (int i = 0; i < childCount; ++i)
            {
                var child = root.GetChild(i);
                var component = child.GetComponent<BindComponent>();
                AddBindComponent(component);

                if (component != null && component.Type == BindType.Node)
                {
                    BindNode node = component as BindNode;
                    node.Initialize();
                    continue;
                }
                else
                {
                    traceBindData(child);
                }
            }
        }

        public bool AddBindComponent(BindComponent component)
        {
            if (component == null) return false;
            if (component.Type == BindType.Container || component.Type == BindType.None) return false;

            bindings.Add(new BindData(component));
            return true;
        }

        public List<BindData> GetBindings()
        {
            return bindings;
        }
    }
}
