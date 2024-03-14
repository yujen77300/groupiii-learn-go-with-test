package selectgoroutine

import (
	"fmt"
	"net/http"
	"time"
)

// v1+v2
func Racer(urlA, urlB string) (winner string) {

	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		return urlA
	}

	return urlB
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}

// ============================
// v3
func RacerWithSelect(urlA, urlB string,) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
