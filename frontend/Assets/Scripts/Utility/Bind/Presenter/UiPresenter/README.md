## Prototype - frontend

<br><br>

# UiPresenter

## 介紹

- UI 介面的 Presenter，需搭配 [Prototype UiManager](../../../UI/README.md)

## 說明

- 使用 [Prototype UiManager](../../../UI/README.md) 管理的 UI 介面都需繼承此類別
- 操作的 UI 介面要先設定 [Bind](../../README.md) 資料

## 內容

- 成員變數 :

    | 變數名稱 | 功能 |
    |:-:|:--|
    |ViewPath|UI 介面的 Prefab 位置|
    |Layer|UI 介面的[圖層位置](../../../UI/IView.cs#L7)|
    |state|UI 介面的[狀態](#ui-介面操作)|

- 成員函數 :

    | 函數名稱 | 功能 |
    |:-:|:--|
    |onBindingCompleted|Binding 完成時呼叫|
    |onOpenHandle|UI 顯示時呼叫|
    |onHideHandle|UI 隱藏時呼叫|
    |onCloseHandle|UI 關閉時呼叫|

## 用法

### 宣告
```cs
// 宣告一個管理指定 UI 介面的 Presenter
class TutorialPresenter : UiPresenter
{
    // 設定 UI 介面的 Prefab 位置
    public override string ViewPath => "tutorial.prefab";

    // 設定顯示圖層
    public override ViewLayer Layer => ViewLayer.BackGround;

    // 宣告要操作的 componet
    Text message;

    TutorialPresenter()
    {
        // 成員變數初始化
        message = null;
    }

    protected override void onBindingCompleted()
    {
        // Binding 完成時呼叫
        message = getBindData<Text>("msg_txt");

        // Binding 後就可對物件操作
        message.text = "hello world";
    }

    protected override void onOpenHandle()
    {
        // 介面開啟時呼叫
        base.onOpenHandle();
    }

    protected override void onHideHandle()
    {
        // 介面隱藏時呼叫
        base.onHideHandle();
    }

    protected override void onCloseHandle()
    {
        // 介面關閉時呼叫
        base.onCloseHandle();
    }
}
```

### 操作

- UManager 宣告可參考 [UiManager](../../../UI/README.md)

```cs
var uiManager = Singleton<UiManager>.Instance;
// 開啟 UI介面
uiManager.Open<TutorialPresenter>();
// 關閉 UI介面
uiManager.Close<TutorialPresenter>();
```



## 補充

### UI介面狀態
- 參考 [Enum ViewState](../../../UI/IView.cs#L16)

    | 變數名稱 | 說明 |
    |:-:|:--|
    |WaitBinding|等待與 UI 介面 Binding，通常表示 UI 物件尚未 Loading 完成|
    |BingCompleted|完成與 UI 介面 Binding|
    |Open|UI 介面顯示於畫面中|
    |Hide|UI 介面隱藏|
    |Close|UI 介面關閉，並等待資源釋放|
