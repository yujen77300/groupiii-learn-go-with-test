# 多準備
1. 如果長方形和圓形都要計算面積，函式應該留下，有一些語言可以這樣做(什麼語言可以這樣做)
```
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```
所以要重新定義方法

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

當方法被這種類型的變數呼叫時，資料的引用會透過變數 receiverName 取得。在其他許多程式語言中這些被隱藏起來並且透過 this 來獲得接收者。

2. interface

> interface是一種類型

介面在 Go 這種靜態型別語言中是一種非常強而有力的概念。因為介面可以讓函數接受不同類型的參數並能創造類型安全且高解耦的程式碼。

在 Go 語言中 interface resolution 是隱式的。如果傳入的類型符合介面需要的，則編譯正確。 => 是不是可以說golang interface duck typing的觀念

是duck-type programming的一種體現。介面做的事情就像是定義一個協定（規則），只要一台機器有洗衣服和甩乾的功能，我就稱它為洗衣機。不關心屬性（資料），只關心行為（方法）。


1. 類型定義
alias並不會創造新的類別，只是該類別一個別名而已。type definition就會創建新的型別
```go
// alias declaration
type T1 = T2

//type definition
type T1 T2

//類型定義
type NewInt int

//類型別名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	
	fmt.Printf("type of a:%T", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T", b) //type of b:int
}
```

常見的byte和rune就是alias，例如
`type byte = uint8`, `type rune = int32`

但有時候描述一個事物有多個屬性，go提供了struct可以封裝多個基本數據類型。

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "小熊維尼"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={小熊維尼 北京 18}
	fmt.Printf("p1=%+v\n", p1) //p1={name:小熊維尼 city:北京 age:18}


    // 用new進行實例化，會得到的是指向結構體的指針
    var p2 = new(person)
}
```

結構體的fields是連續存取，優點可以有
+ 內存的訪問效率提高
+ 提高cpu cache的效率
+ 減少記憶體碎片化
[Memory alignment in golang](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/go-memory-align)

```go
func main() {
    type test struct {
        a int8
        b int8
        c int8
        d int8
    }
    n := test{
        1, 2, 3, 4,
    }
    fmt.Printf("n.a %p\n", &n.a)
    fmt.Printf("n.b %p\n", &n.b)
    fmt.Printf("n.c %p\n", &n.c)
    fmt.Printf("n.d %p\n", &n.d)
}
```

4. constructor function
```golang
// Task: 因為struct是值型，如果結構體比較複雜的話，值拷貝性能開銷會比較大，所以該構造函數回傳的是結構體指標類型，可以實際跑benchmark測試看看。
type Person struct {
	name string
	city string
	age  int8
}

func main() {
	pp := newPerson("john", "paris", 20)
	fmt.Println(pp)
	pp.Dream()

}

func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

// 接收者變量：接收者中的参数變量名在命名時，官方建議使用接收者類型名的第一個小寫字母，而不是self、this之類的命名。例如，Person類型的接收者變量應該命名為 p。

// 這個案例是值類型的接收，如果要改變變數，因為有copy一份所以指示改道複製的那一份，原本的值不會變。因此建議如果要改變可以用指針類型。
func (p Person) Dream() {
	fmt.Printf("%s喜歡Go language！\n", p.name)
}

```

5. embedd struct  和實現繼承的方法
```golang
type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "威斯康辛州",
			City:     "密爾瓦基",
		},
	}
	fmt.Printf("user1=%+v\n", user1)
}
```

```golang
// 如果要共享就要用指針
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s會動!\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s會汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "小白",
		},
	}
	d1.wang()
	d1.move()
}

```


6. json的應用
Task :stuct與JSON的應用



7. interface
Task 舉例為什麼要用interface
這次讀書會書上的例子是例如三角形，四邊形，圓形都能計算週長和麵積，我們能不能把它們當成「圖形」來處理呢？
例如銷售、行政、工程師都能計算月薪，我們能不能把他們當成「員工」來處理呢？

go語言的介面在命名時，通常會在單字後面加上er，如有寫入操作的介面叫Writer，有字串功能的介面叫Stringer等
接口名最好要能突顯該介面的類型含意。

> 接口是一種抽象的類型。當你看到一個接口類型的值時，你不知道它是什麼，唯一知道的是通過它的方法能做什麼。

```golang
type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

type Sayer interface{
	Say() string
}

func main() {
	var sayer Sayer

	c := Cat{}
	sayer = c
	fmt.Println("猫:", sayer.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
```


8. 思考: 一個類型可以滿足多個interface的好處

> 只有當有兩個或兩個以上的具體類型必須以相同的方式進行處理時才需要定義介面。不要為了介面而寫介面，那樣只會增加不必要的抽象，導致不必要的運行時損耗。
 Interfaces can be added after the fact if a new idea comes along or for testing—without annotating the original types.

9. 方法是可以繼承
```golang
type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名欄位
    school string
}

type Employee struct {
    Human //匿名欄位
    company string
}

//在 human 上面定義了一個 method
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
    mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
    sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

    mark.SayHi()
    sam.SayHi()
}

```

9. 可以增加的重點
+ value type 進行賦值或到函數裡面會先copy一個新的
+ 遷入interface, 遷入struct


11. benchtest
避免過度使用接口：

過度使用接口可能會導致性能下降。僅在必要的時候使用接口，並在可能的情況下使用具體類型。
使用指針而非值接收器：

當你定義方法時，考慮使用指針接收器而不是值接收器。這樣可以避免在方法調用時進行值的拷貝，提高性能。


go interface源码在runtime\iface.go中。

go的介面由兩種類型實作iface和eface。 iface是包含方法的接口，而eface不包含方法。

