reflect package 可以用來檢查變數的型別

## 基本用法

1、取得變數型別、變數值、該型別的zero value
```go
func main() {
	var num int = 9
	typ := reflect.TypeOf(num)
	val := reflect.ValueOf(num)
	zeroVal := reflect.Zero(typ)

	fmt.Printf("Type: %v\n", typ)           // Type: int
	fmt.Printf("Value: %v\n", val)          // Value: 9
	fmt.Printf("Zero Value: %v\n", zeroVal) // Zero Value: 0
}
```

2、取得struct相關資訊
```go
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alan", Age: 30}
	val := reflect.ValueOf(person)

	nameField := val.FieldByName("Name")
	fmt.Printf("Name: %v\n", nameField) // Name: "Alan"
}
```


## 重要概念
reflect package主要圍繞在type Type interface、type Value struct
因執行的成本較高、程式碼也會較於複雜，建議避免使用。

## 題目

x interface{}：這是一個空接口，可以接受任何類型的值作為參數。
fn func(string)：這是一個函數，接受一個字符串作為參數。
函數的目的是递歸地遍歷結構體 x 中的所有字段，並對於每個字符串字段，調用傳入的函數 fn 進行處理。

這個挑戰的難點在於要遞歸地處理結構體中的字段，因為結構體可能包含其他結構體，所以需要递歸地訪問所有字段。同時，需要判斷字段的類型是否為字符串，如果是字符串類型，則調用傳入的函數進行處理。



## Ref
1. [Reflect：Go 語言的鏡子](https://vocus.cc/article/6508042ffd897800019e3100)
2. [Go 语言设计与实现 - 反射 ](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/)
3. [ openhome - 反射](https://openhome.cc/Gossip/Go/Reflect.html)
4. [ The Go Blog - The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
