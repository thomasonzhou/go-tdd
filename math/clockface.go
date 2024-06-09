package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y int
}

func SecondHand(tm time.Time) Point {

	return Point{150, 60}
}

func SecondsToRadians(tm time.Time) float64 {
	const radsInCircle = 2.0 * math.Pi
	return radsInCircle * float64((tm.Second() % 60)) / 60.0
}
