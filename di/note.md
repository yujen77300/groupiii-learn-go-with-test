在Go語言中，依賴注入（Dependency Injection，簡稱DI）是一種設計模式，用於實現控制反轉（Inversion of Control，IoC），以提升程式碼的模組化和可測試性。Go語言中實現依賴注入的關鍵點包括：

介面（Interface）: Go語言透過其強大的介面系統支援依賴注入。你可以定義介面來表示依賴項的行為，然後在函數或結構體中使用這些介面作為參數。

結構體（Struct）: 透過在結構體中包含介面類型的字段，可以將具體的依賴項注入到結構體實例中。這允許在運行時動態更換依賴項。

函數參數: 另一種常見的依賴注入方式是直接在函數參數中傳遞依賴項。這種方式簡單直接，適用於不需要太多依賴管理的場景。


## 依賴注入

為一種設計模式，管理對象之間的依賴關係

> 傳遞「依賴」給使用者，而不是讓使用者直接獲得「依賴」

依賴 => 可以用蓋大樓來比喻，地基 -> 一樓 -> 二樓
注入 => 每一層樓所需的資源（材料、工人等）自動且恰當地提供給建造過程，以軟體開發來看使得這些組件或服務可以在不直接管理這些依賴的情況下正常工作。


在GO中用interface和method來舉例 
+ 介面（Interface）：「咖啡」。它是一個通用的概念，代表了一切可以被當作咖啡的飲品。它有一些通用特性，比如含有咖啡因，熱的，通常放在杯子裡。
+ 實現（Implementation）：拿鐵、美式、卡布奇諾等。這些是「咖啡」介面的具體實現。它們都是咖啡，但每一種都有其獨特的製作方法和口味。
在咖啡廳中，無論想要哪種咖啡（拿鐵、美式還是卡布奇諾），你都只需要說「我要一杯咖啡」。咖啡師（依賴注入容器）會根據你的需求（配置或情境），為你提供正確的咖啡種類。這樣，即使你改變喜歡的咖啡類型，你也不需要改變你的點單方式，因為「咖啡」這個介面始終如一。

依賴注入框架
+ 自動解析依賴關係，省去了手動編寫和維護依賴注入代碼的麻煩
+ 易於維護和重構
+ 例子
    + [uber - dig](https://github.com/uber-go/dig)
        + 透顧reflection在runtime進行依賴注入, 較為強大，但runtime才可以看到錯誤
    + [google - wire](https://github.com/google/wire)
        + go generate產生程式碼, 限制較多, compile時候就可以看到錯誤



## 依賴注入實踐方式
1. constructor function injection
    ```golang
    type UserService struct {
        userRepository *UserRepository
    }

    func NewUserService(userRepository *UserRepository) *UserService {
        return &UserService{userRepository: userRepository}
    }


    ```
2. Attribute injection
    ```golang
    type UserService struct {
    userRepository *UserRepository
    }

    func (u *UserService) SetUserRepository(userRepository *UserRepository) {
        u.userRepository = userRepository
    }
    ```
3. Method injecion
    ```golang
    type UserService struct {}

    func (u *UserService) SaveUser(user *User, userRepository *UserRepository) error {
        return userRepository.SaveUser(user)
    }
    ```


# Ref.

1. [何謂依賴注入](https://medium.com/wenchin-rolls-around/%E6%B7%BA%E5%85%A5%E6%B7%BA%E5%87%BA-dependency-injection-ea672ba033ca)
2. [3種依賴注入的實踐方式](https://tehub.com/a/c0W0jZ5qR8)
3. golang dependency injection影片
4. [Go单测从零到溜系列6—编写可测试的代码](https://www.liwenzhou.com/posts/Go/unit-test-6/)
5. [Mastering Dependency Injection In Golang](https://www.youtube.com/watch?v=UX4XjxWcDB4)
6. [pure function 介紹](https://ithelp.ithome.com.tw/articles/10262576)



# 補充
S.O.L.I.D 
SRP The Single Responsibility Principle 單一責任原則
OCP The Open Closed Principle 開放封閉原則
LSP The Liskov Substitution Principle 里氏替换原則
ISP The Interface Segregation Principle 接口分離原則
DIP The Dependency Inversion Principle 依賴倒置原則

