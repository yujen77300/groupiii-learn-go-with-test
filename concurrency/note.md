1. 為何使用 **!reflect.DeepEqual(expectedResults, actualResults)** 和 **if want != got**
    =>映射是引用類型，不能直接使用 == 或 != 符號來比較。
    => reflect.DeepEqual 會進行深度比較，檢查映射中的所有鍵和相對應的值，但性能較差

2. 判斷是否有race condition
    => go test -race

3. 避免多個goroutine訪問和修改同個資源
    => channel & mutex