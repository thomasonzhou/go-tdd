package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const (
	radsInCircle    = 2 * math.Pi
	secondsInCircle = 60
	minutesInCircle = 60
	hoursInCircle   = 12
)

func secondsToRadians(t time.Time) float64 {
	return radsInCircle * float64((t.Second())) / secondsInCircle
}

func minutesToRadians(t time.Time) float64 {
	extraAngleFromSeconds := secondsToRadians(t) / minutesInCircle
	return radsInCircle*float64((t.Minute()))/minutesInCircle + extraAngleFromSeconds
}

func hoursToRadians(t time.Time) float64 {
	extraAngleFromMinutes := minutesToRadians(t) / hoursInCircle
	return radsInCircle*float64((t.Hour()%hoursInCircle))/hoursInCircle + extraAngleFromMinutes
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
