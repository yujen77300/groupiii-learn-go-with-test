package selectgoroutine

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// 兩個url比賽，返回先回應的url
// v1
// func TestRacer(t *testing.T) {
// 	slowURL := "http://www.facebook.com"
// 	fastURL := "http://www.quii.co.uk"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got '%s', want '%s'", got, want)
// 	}
// }

// ============================
// v2
func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(10 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	defer slowServer.Close()
	defer fastServer.Close()

	// v4 測試10秒以上
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := RacerWithSelect(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

// 建立了一個新的 httptest.Server 實例，並且使用了一個 http.HandlerFunc 來處理 HTTP 請求，建立一個簡單的 HTTP 處理器。函數在接收到請求時會先等待 20 毫秒，然後將 HTTP 狀態碼設置為200並返回。
func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

func BenchmarkRacer(b *testing.B) {
	slowServer := makeDelayedServer(10 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	for i := 0; i < b.N; i++ {
		// Racer(slowURL, fastURL)
		RacerWithSelect(slowURL, fastURL)
	}
}
