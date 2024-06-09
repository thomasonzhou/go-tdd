package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

func secondsToRadians(t time.Time) float64 {
	const radsInCircle = 2.0 * math.Pi
	return radsInCircle * float64((t.Second() % 60)) / 60.0
}

func minutesToRadians(t time.Time) float64 {
	const radsInCircle = 2.0 * math.Pi
	extraAngleFromSeconds := secondsToRadians(t) / 60
	return radsInCircle*float64((t.Minute()%60))/60.0 + extraAngleFromSeconds
}

func secondsToPoint(t time.Time) Point {
	radians := secondsToRadians(t)

	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}

func minutesToPoint(t time.Time) Point {
	radians := minutesToRadians(t)

	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}
