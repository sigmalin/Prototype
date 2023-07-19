## Prototype - frontend

<br><br>

# Factory

## 介紹

- Design Pattern Factory 的實現架構

- Factory 的引數必須繼承 [IOrder](IOrder.cs) 類別

- 泛用 Enum 變數 Type 用來對相同功能做分類的

- 所有產品的設定參數寫在繼承 [IOrder](IOrder.cs) 類別

- 可參考 [ProtocolFactory](../Network/WebRequest/Protocol/ProtocolFactory.cs) 、 [ProviderFactory](../Network/WebRequest/Provider/ProviderFactory.cs)