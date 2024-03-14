
## 簡介
在go學習tdd的時候，
在 Golang 的測試框架中，httptest 的主要作用為幫助測試者模擬 HTTP 請求和回應，更方便地對伺服器API 接口進行測試。它不依賴網路和實際的http伺服器，而是使用 Golang 提供的 net/http 中的功能，模擬出請求和回應的流程。

## Key
```go
func NewServer(handler http.Handler) *Server {
	ts := NewUnstartedServer(handler)
	ts.Start()
	return ts
}
```
接受一個 http.Handler 作為參數，並返回一個新的 httptest.Server 實例。這個伺服器會在本地端口上啟動，並且在測試結束後自動關閉。對於測試需要與 HTTP 伺服器交互的功能非常有用。

http.Handler 是http package中的一個接口，它定義了一個方法 ServeHTTP(ResponseWriter, *Request)。這個方法接受兩個參數：一個 ResponseWriter 和一個指向 *Request 的指針。ResponseWriter 是一個接口，它提供了寫入 HTTP 回應的方法，而 *Request 是一個結構，它包含了 HTTP 請求的所有信息。

http.Handler 接口的目的是讓能夠實現自己的 HTTP 處理器，這些處理器可以被用來處理特定的 HTTP 請求。這對於建立自定義的 HTTP 服務器非常有用，因為它允許你根據請求的內容來決定如何回應。

```go
type Server struct {
	URL      string // base URL of form http://ipaddr:port with no trailing slash
	Listener net.Listener
    // ...
}
```
httptest.Server.URL提供了伺服器的 URL

httptest.Server.Close:
```go
// Close shuts down the server and blocks until all outstanding
// requests on this server have completed.
func (s *Server) Close() {
	s.mu.Lock()
	if !s.closed {
    }
}
```
手動關閉由 NewServer 或 NewTLSServer 創建的伺服器，用於在測試結束後清理資源。

## 實做

在v1版本會透過Racer函數裡面的http.Get來請求url資料，但會有速度慢不可靠的問題


httptest.NewServer 傳入 匿名的http.HandlerFunc 參數，會在本地尋找一個可以監聽的port，在v2版本就可以在輕鬆帶入	time.Sleep(20 * time.Millisecond)來調整測試



## goroutine + select
v3版本
為什麼要一個一個測試呢，可以同時側兩個，且不用關注時間，只要知道哪個比較快

運用select可以允許在多個channel等待，第一個回傳資訊的程式碼就會被執行

可以用BenchmarkRacer來看出彼此的差異

v4超過十秒的案例
在某個「案例」中使用 time.After 來防止系統被永久阻塞


## 結尾
httptest 的注意事项

（1）httptest 不應該在生產環境的 HTTP 或 HTTPS 伺服器，僅用在測試刺目的。

（2）    測試伺服器最後要關閉 關閉一個server，才不會測試套件消耗過多或是持續監聽port，消耗資源

（3）如果使用 httptest 提交到的数据和你所提交的数据不匹配，你将无法得到正确的测试结果。

httptest 是一個非常有用的工具，可以幫助我們模擬 HTTP 請求和回應，從而方便地對服務端的 API 介面進行測試。詳細介紹了 httptest 的基本概念、用法和注意事項，希望能對 Golang 開發人員在接下來的測試工作中有所幫助