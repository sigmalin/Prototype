using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UnityEditor;

namespace Bind.Editor
{
    [CustomEditor(typeof(BindNode))]
    public class BindNodeEditor : UiBindingEditor
    {
        public override void OnInspectorGUI()
        {
            base.OnInspectorGUI();

            GUILayout.Space(10f);

            drawBinding();                     
        }
    }
}
