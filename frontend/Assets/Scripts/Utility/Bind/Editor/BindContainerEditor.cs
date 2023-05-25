using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UnityEditor;

namespace Bind.Editor
{
    [CustomEditor(typeof(BindContainer))]
    public class BindContainerEditor : UiBindingEditor
    {
        public override void OnInspectorGUI()
        {
            base.OnInspectorGUI();

            GUILayout.Space(10f);

            drawBinding();

            GUILayout.Space(10f);

            if (GUILayout.Button("Re-bind All Component"))
            {
                (target as BindContainer).Initialize();

                EditorUtility.SetDirty(target);

                EditorUtility.DisplayDialog("Re-bind", "Re-bind All Componet Completed", "OK");
            }

            GUILayout.Space(10f);

            BindHierarchyLabel.showInfo = EditorGUILayout.Toggle("Show Info in Hierarchy", BindHierarchyLabel.showInfo);
        }
    }
}
