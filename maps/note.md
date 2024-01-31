1. 一個map裡面的key值必須是可以比較的
    + 鍵的類型必須是可比較的：在 Go 中，映射的鍵必須是一種可以進行比較的類型。這意味著你可以使用 == 和 != 運算符來比較兩個鍵。可比較的類型包括基本類型（如 int、string）和一些複合類型（如數組），但不包括切片、映射以及包含切片的結構體，因為這些類型本身不支持比較運算。
    + 為什麼需要可比較的鍵：映射的主要作用是根據鍵來存儲和檢索值。如果不能確定兩個鍵是否相同，就無法正確地檢索或更新映射中的值。例如，當你嘗試獲取或設置映射中的一個元素時，Go 需要能夠判斷你提供的鍵是否已經存在於映射中。



2. map 是 reference type, 不需要指針就可以修改值，其底層的數據結構是hash map

3. 不應該透過`var m map[string]string`初始化一個空的map變數，其並沒有分配記憶體空間，默認值是nil，如果對nil進行操作會有runtime error。建議解決方法為如下，這樣就可以練一個空的hash map指向dictionary
```golang
dictionary = map[string]string{}
// OR
dictionary = make(map[string]string)
```


4. 有兩種for range方式
```golang
func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}

    key := "He"
    if val, ok := elements[key]; ok {
        fmt.Println("The element", key, "is", val)
        return
        // Output:
        // The element He is Helium
    }
    fmt.Println("The element", key, "is not in the map")


    for key , value :=range elements{
        fmt.Println(key)
        fmt.Println(value)
    }
}


```