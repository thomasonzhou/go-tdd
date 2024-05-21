package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string, error error) {
	return ConfigurableRacer(url1, url2, 10*time.Second)
}

func ConfigurableRacer(url1 string, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %q and %q", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
