using System.Collections.Generic;
using System.Linq;
using System;
using UnityEngine;
using UnityEditor;

namespace Bind.Editor
{
    public class BindHierarchyLabel
    {
        [InitializeOnLoad]
        static class LabelDisplay
        {
            static Font _font;
            static Font font
            {
                get
                {
                    if (_font == null)
                    {
                        _font = LoadFont("arial");
                    }

                    return _font;
                }
            }

            static LabelDisplay()
            {
                EditorApplication.hierarchyWindowItemOnGUI += HighlightItems;
            }

            /// <summary>
            /// Gets a string of characters mapping for the game object.
            /// </summary>
            static string GetLabel(GameObject target, out GUIStyle style)
            {
                var binding = target.GetComponent<BindComponent>();

                style = null;

                string label = "";

                if (binding != null)
                {
                    var Bindingtype = binding.GetType();

                    if (LabelMapping.componentLabels.ContainsKey(Bindingtype))
                    {
                        label = LabelMapping.componentLabels[Bindingtype](binding);
                    }

                    StyleMapping.componentStyles.TryGetValue(Bindingtype, out style);
                }

                return label;
            }

            /// <summary>
            /// Displays label beside each game object in the Hierarchy panel.
            /// </summary>
            static void HighlightItems(int instanceID, Rect selectionRect)
            {
                if (showInfo == false) return;

                var target = EditorUtility.InstanceIDToObject(instanceID) as GameObject;

                if (target == null)
                    return;

                GUIStyle style;

                var labelString = GetLabel(target, out style);

                if (style != null)
                {
                    GUI.Label(selectionRect, labelString, style);
                }
            }

            /// <summary>
            /// Loads a font from the asset database.
            /// </summary>
            static Font LoadFont(string name)
            {
                var guid = AssetDatabase.FindAssets(name + " t:font").First();
                var path = AssetDatabase.GUIDToAssetPath(guid);
                return AssetDatabase.LoadAssetAtPath<Font>(path);
            }
        }

        static class LabelMapping
        {
            public static readonly Dictionary<Type, Func<BindComponent, string>> componentLabels = new Dictionary<Type, Func<BindComponent, string>>()
            {
                { typeof(BindElement), getElementInfo },
                { typeof(BindNode), getNodeInfo },
                { typeof(BindContainer), getContainerInfo },
            };

            public static string getElementInfo(BindComponent component)
            {
                BindElement element = component as BindElement;

                return $"E ( {element.ID.Value} )";
            }

            public static string getNodeInfo(BindComponent component)
            {
                BindNode node = component as BindNode;

                return $"N ( {node.ID.Value} )";
            }

            public static string getContainerInfo(BindComponent component)
            {
                return "Container";
            }
        }

        static class StyleMapping
        {
            static GUIStyleState yellow = new GUIStyleState() { textColor = Color.yellow };
            static GUIStyleState green = new GUIStyleState() { textColor = Color.green };
            static GUIStyleState blue = new GUIStyleState() { textColor = Color.blue };

            public static readonly Dictionary<Type, GUIStyle> componentStyles = new Dictionary<Type, GUIStyle>()
            {
                { typeof(BindElement), new GUIStyle(EditorStyles.label) { normal = yellow, alignment = TextAnchor.MiddleRight } },
                { typeof(BindNode), new GUIStyle(EditorStyles.label) { normal = green, alignment = TextAnchor.MiddleRight } },
                { typeof(BindContainer), new GUIStyle(EditorStyles.label) { normal = blue, alignment = TextAnchor.MiddleRight } },
            };
        }

        static public bool showInfo = false;
    }    
}
