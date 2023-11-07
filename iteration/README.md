# Notes

## Benchmark
```golang
func BenchmarkRepeat(b *testing.B){
	for i:=0 ; i<b.N;i++{
		Repeat("a")
	}
}
```
testing.B這個結構體提供的b.N屬性代表會執行的迭代次數。

> The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

執行 `go test -bench .`的到下面結果
```
goos: windows
goarch: amd64
pkg: github.com/yujen77300/groupiii-learn-go-with-test/iteration
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkRepeat-8       12374604                85.69 ns/op
```  
代表在我的電腦最大cpu數為8，b.N最大值是12374604，每次循環平均需要85.69ns。


## Others
`:=` represent declaring and initializing

`var repeated string`為更明顯的宣告方式，且[Go有zero value的特性](https://exercism.org/tracks/go/concepts/zero-values)，會直接賦予空字串。

+ 在Python裡如果變數沒有宣告會有<span style="color:red">NameError</span>，如果宣告但沒有初始化，會被賦予為None或是零。
+ 在JS裡如果只有宣告但沒有初始化，會被賦予為undefined。

# Ref.
1. [Go语言基准测试(benchmark)三部曲之一：基础篇](https://cloud.tencent.com/developer/article/2216022)