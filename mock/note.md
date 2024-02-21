# Mocking 筆記

## Table of Contents
- [Mocking 筆記](#mocking-筆記)
  - [Table of Contents](#table-of-contents)
  - [Intro to Mock](#intro-to-mock)
  - [Implementation Based on Book](#implementation-based-on-book)
  - [Test Doubles](#test-doubles)
  - [References](#references)


## Intro to Mock

參考 [The Clean Code Blog - When to Mock](https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html)，提到何時該使用mock

+ 測試需要使用很多外部資源，如外部的api呼叫, 資料庫
+ 測試易受外部系統干擾：測試結果可能受到外部系統狀態的影響，例如網路延遲、資料庫狀態等。
+ 測試覆蓋率低，例如需要模擬錯誤的error handling 、一些要執行危險動作的函數(如刪除檔案, 資料表)

什麼情況最好
+ 使用在不同邊界
+ 自行編寫mock，維持測試的簡潔和速度。

## Implementation Based on Book

+ 拆分需求讓軟體可以持續增加功能是很重要的一件事，每次只實現一个最小可行的功能,可以盡快讓軟體上線,也更容易測試和擴展,這是一種好的迭代開發方法。

+ 如果不用返引號，要用雙引號`want := "3\n2\n1\nGo!"`

+ 用依賴注入來代替真實的time.Sleep。代表插入一個mock的time.Sleep函式，這個函式在行為上與真實的time.Sleep相似，但可以被監視和控制。這樣，你就可以在測試中“spy on”（監視）這個模擬的time.Sleep函式，檢查它是否按預期被呼叫，以及被呼叫的次數和時機。



模擬可能被視為有害（"mocking is evil"）：
有些人認為過度使用模擬（即在單元測試中代替真實對象的技術）是有害的。這主要是因為不當使用模擬可能導致程式碼更難維護、理解和測試。

模擬的適當與否取決於使用方式：
+ 與許多軟件開發中的技術一樣，模擬本身不是有害的，它的好壞處取決於如何使用。這就像DRY（Don't Repeat Yourself，不要重複自己）原則一樣，過度或不當使用可能導致問題。

+ 你正在进行的测试需要做太多的事情
 => 把模块分开就会减少测试内容
+ 它的依赖关系太细致
 => 考虑如何将这些依赖项合并到一个有意义的模块中
+ 你的测试过于关注实现细节
 => 專注在测试预期的行为，而不是功能的实现


測試過度關注實現細節：優先測試預期的行為，而不是具體的實現細節。


## Test Doubles

參考notify folder

## References
1. [The Clean Code Blog - When to Mock](https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html)
2. [The Clean Code Blog - The Little Mocker](https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html)
3. [Martin Fowler - Test Doubles](https://martinfowler.com/bliki/TestDouble.html)
4. [Go Test Doubles by Example](https://aminrb.me/blog/go-test-double/)
5. [测试代码的基本结构以及(Test Double)的种类](https://zhuanlan.zhihu.com/p/582029396)
6. [Go 单元测试时尽量用 fake object](https://zhuanlan.zhihu.com/p/624485568)