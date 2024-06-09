package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const radsInCircle = 2 * math.Pi

func secondsToRadians(t time.Time) float64 {
	return radsInCircle * float64((t.Second())) / 60.0
}

func minutesToRadians(t time.Time) float64 {
	extraAngleFromSeconds := secondsToRadians(t) / 60
	return radsInCircle*float64((t.Minute()))/60.0 + extraAngleFromSeconds
}

func hoursToRadians(t time.Time) float64 {
	extraAngleFromMinutes := minutesToRadians(t) / 12
	return radsInCircle*float64((t.Hour()%12))/12.0 + extraAngleFromMinutes
}

func secondsToPoint(t time.Time) Point {
	radians := secondsToRadians(t)
	return radiansToPoint(radians)
}

func minutesToPoint(t time.Time) Point {
	radians := minutesToRadians(t)
	return radiansToPoint(radians)
}

func hoursToPoint(t time.Time) Point {
	radians := hoursToRadians(t)
	return radiansToPoint(radians)
}

func radiansToPoint(radians float64) Point {
	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}
