package clockface

import (
	"math"
	"time"
)

// A cartesian coordinate. Inside this package, positive X is right and positive Y is up
type Point struct {
	X, Y float64
}

const (
	radsInCircle    = 2 * math.Pi
	secondsInCircle = 60
	minutesInCircle = 60
	hoursInCircle   = 12
)

func SecondsToRadians(t time.Time) float64 {
	return radsInCircle * float64((t.Second())) / secondsInCircle
}

func MinutesToRadians(t time.Time) float64 {
	extraAngleFromSeconds := SecondsToRadians(t) / minutesInCircle
	return radsInCircle*float64((t.Minute()))/minutesInCircle + extraAngleFromSeconds
}

func HoursToRadians(t time.Time) float64 {
	extraAngleFromMinutes := MinutesToRadians(t) / hoursInCircle
	return radsInCircle*float64((t.Hour()%hoursInCircle))/hoursInCircle + extraAngleFromMinutes
}

func SecondsToPoint(t time.Time) Point {
	radians := SecondsToRadians(t)
	return radiansToPoint(radians)
}

func MinutesToPoint(t time.Time) Point {
	radians := MinutesToRadians(t)
	return radiansToPoint(radians)
}

func HoursToPoint(t time.Time) Point {
	radians := HoursToRadians(t)
	return radiansToPoint(radians)
}

func radiansToPoint(radians float64) Point {
	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}
