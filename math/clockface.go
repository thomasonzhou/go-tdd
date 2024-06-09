package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

func secondsToRadians(tm time.Time) float64 {
	const radsInCircle = 2.0 * math.Pi
	return radsInCircle * float64((tm.Second() % 60)) / 60.0
}

func secondsToPoint(tm time.Time) Point {
	radians := secondsToRadians(tm)

	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}
