package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	time1 := measureGetTime(url1)
	time2 := measureGetTime(url2)

	if time1 < time2 {
		return url1
	} else {
		return url2
	}
}

func measureGetTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	time := time.Since(start)
	return time
}
