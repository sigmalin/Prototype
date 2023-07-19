## Prototype - frontend

<br>

Translations:

* [繁體中文](README.md)

---

<br>

# Introduction

- Front-end source code, developed in Unity3D

<br><br>

# How to use

- See [Demo](./Assets/Demo/) folder

- Before running, must be setting addressable as below

![Group Name can free，but Addressable Name must the same as photo](../imgs/addressables.png)

- Demo show how to signin (or login) and get user data from Api server

- Demo may be added (modified) at any time as the framework expands.

- The framework uses the model-view-presenter design pattern, and the implementation can be found in [Binding](./Assets/Scripts/Utility/Bind/README.md)

<br><br>

# Frequently questions

- how to implementation MonoBehaviour::Update at presenter component

Considering the performance of [Update](https://blog.unity.com/engine-platform/10000-update-calls), presenter is not implemented as a method for each frame call.
<br>
There are two ways to run the Update function in presenter component

1.Custom Manager that inheritance MonoBehaviour，and call all registered presenter component member function at Update method

2.[UniRx](https://github.com/neuecc/UniRx) (recommond)

<br><br>

# TODO

- Code generate for Presenter and Node

Add a new button at inspector of presenter and node that automatically generates the corresponding program code when clicked.

- Localization

The current idea is to package each language resource separately as it needs to be adapted to all resource loading tools.
<br>
Then modify Presenter::ViewPath so that have different values in different languages.

- Optimize the resource loading and updating process

Currently, addressable is used as a resource loading (updating) tool, but the updating method is completely the basic design of addressable, and it is not possible to achieve additional functions such as dynamic updating of download addresses, encryption of resource files, etc. It is expected to expand the functions in this area.

- Added local read and write example for Server response.

The current framework has a plan for this, and also provides [Model](./Assets/Scripts/Utility/Network/WebRequest/Model/) to implement，However, there is no actual operation in the demo, and we will provide relevant examples based on the expansion of the back-end in the future.

- Manage 3D Objects

Currently, we only provide game UI operation, and the 3D objects will be basically the same as the UI operation, using the model-view-presenter design pattern.
<br>
A [UiManager](./Assets/Scripts/Utility/UI/README.md)-like manager is provided to manage all 3D objects.



