using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UnityEditor;
using UnityEditorInternal;

namespace Bind.Editor
{
    public class UiBindingEditor : UnityEditor.Editor
    {
        SerializedProperty bindings;

        ReorderableList list;

        GUIStyle nodeStyle;
        GUIStyle elementStyle;

        private void OnEnable()
        {
            nodeStyle = new GUIStyle();
            nodeStyle.normal.textColor = Color.yellow;

            elementStyle = new GUIStyle();
            elementStyle.normal.textColor = Color.green;

            bindings = serializedObject.FindProperty("bindings");

            list = new ReorderableList(serializedObject, bindings, true, true, true, true);
            list.drawElementCallback = drawListItems;
            list.drawHeaderCallback = drawHeader;
        }

        protected void drawBinding()
        {
            serializedObject.Update();

            list.DoLayoutList();
            
            serializedObject.ApplyModifiedProperties();
        }

        GUIStyle getStyle(BindType type)
        {
            GUIStyle style = null;
            StyleMapping.componentStyles.TryGetValue(type, out style);
            return style;
        }

        void drawListItems(Rect rect, int index, bool isActive, bool isFocused)
        {
            SerializedProperty element = list.serializedProperty.GetArrayElementAtIndex(index);

            EditorGUI.LabelField(new Rect(rect.x, rect.y, 60, EditorGUIUtility.singleLineHeight), "Identifier");
                        
            EditorGUI.PropertyField(
                new Rect(rect.x + 60, rect.y, 100, EditorGUIUtility.singleLineHeight),
                element.FindPropertyRelative("identifier").FindPropertyRelative("identifier"),
                GUIContent.none
                );

            element.FindPropertyRelative("identifier").serializedObject.ApplyModifiedProperties();

            BindType type = (BindType)Enum.GetValues(typeof(BindType)).GetValue(element.FindPropertyRelative("type").enumValueIndex);

            EditorGUI.LabelField(new Rect(rect.x + 200, rect.y, 60, EditorGUIUtility.singleLineHeight), $"{type}", getStyle(type));

            EditorGUI.PropertyField(
                new Rect(rect.x + 260, rect.y, 100, EditorGUIUtility.singleLineHeight),
                element.FindPropertyRelative("gameObject"),
                GUIContent.none
                );
        }

        void drawHeader(Rect rect)
        {
            EditorGUI.LabelField(rect, "bindings");
        }
    }

    static class StyleMapping
    {
        static GUIStyleState yellow = new GUIStyleState() { textColor = Color.yellow };
        static GUIStyleState green = new GUIStyleState() { textColor = Color.green };
        static GUIStyleState blue = new GUIStyleState() { textColor = Color.blue };

        public static readonly Dictionary<BindType, GUIStyle> componentStyles = new Dictionary<BindType, GUIStyle>()
        {
            { BindType.Element, new GUIStyle(EditorStyles.label) { normal = yellow } },
            { BindType.Node, new GUIStyle(EditorStyles.label) { normal = green } },
            { BindType.Container, new GUIStyle(EditorStyles.label) { normal = blue } },
        };
    }

}
