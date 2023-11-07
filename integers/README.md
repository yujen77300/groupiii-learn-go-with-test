# Example functions
Go的testing package有提供一個Example方法，可以讓其他開發者了解如何使用你開發的函式，並預期輸出。
且根據下面[原文](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers#examples)敘述，如果撰寫公開的套件，也可自動成為文檔的一部分。
關於如何撰寫公開文檔，請參考[About pkgsite](https://pkg.go.dev/about#adding-a-package)

> If you publish your code with examples to a public URL, you can share the documentation of your code at pkg.go.dev.


+ Examples are compiled (and optionally executed)

    有時候單元測試需要外部網路資源才可執行，但仍想要說明函數的使用功能，這時可以不要撰寫Output comments，讓程式可以complie但不會被執行。

    ```go
    func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
    }
    ```

+ godoec 文件

    `go install golang.org/x/tools/cmd/godoc@latest` => 安裝完後可以執行`godoc -http=:6060`，查看官方文件。


# Ref.
1. [Testable Examples in Go](https://go.dev/blog/examples)
2. [testing package - Examples](https://pkg.go.dev/testing#hdr-Examples)
3. [build your own package](https://www.youtube.com/watch?v=KNHEXOoV-H4)