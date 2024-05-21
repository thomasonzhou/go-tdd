package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	time1 := checkGetSpeed(url1)
	time2 := checkGetSpeed(url2)

	if time1 < time2 {
		return url1
	} else {
		return url2
	}
}

func checkGetSpeed(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	time := time.Since(start)
	return time
}
