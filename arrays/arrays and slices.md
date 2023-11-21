# [Arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)



使用t.Run優點
+ 可以為測試案例命名
+ 在一個測試函數裡面進行多個相關的子測試

```golang
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
```

slice只能跟nil比較
+ slice指向的底層的array, len, cap，不會用`==`來比較兩個slice。因為通常是要比較元素是否相同比較有意義，所以會使用 `reflect.DeepEqual`方法。
+ `==`可以跟nil比較，判斷是否經過初始化
```golang
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
```



`go test -cover` => 用來了解測試的覆蓋率，作者提到如果遵循TDD，覆蓋論會接近100%