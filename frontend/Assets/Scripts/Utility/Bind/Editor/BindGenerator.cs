using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UnityEditor;

namespace Bind.Editor
{
    class BindGenerator
    {
        [MenuItem("GameObject/Binding/BindingElement", false, 0)]
        public static void generateBindingElement()
        {
            attachBindComponent<BindElement>();
        }

        [MenuItem("GameObject/Binding/BindingNode", false, 1)]
        public static void generateBindingNode()
        {
            attachBindComponent<BindNode>();
        }

        [MenuItem("GameObject/Binding/BindingContainer", false, 2)]
        public static void generateBindingContainer()
        {
            attachBindComponent<BindContainer>();
        }

        static void attachBindComponent<T>() where T : BindComponent
        {
            if (Selection.objects == null) return;

            for (int i = 0; i < Selection.objects.Length; ++i)
            {
                GameObject obj = Selection.objects[i] as GameObject;
                if (obj == null) continue;

                BindComponent component = obj.GetComponent<BindComponent>();
                if (component != null) continue;

                obj.AddComponent<T>();
            }
        }
    }
}
