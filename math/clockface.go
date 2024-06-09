package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const SecondHandLength = 90

const ClockCenterX = 150
const ClockCenterY = 150

func SecondHand(tm time.Time) Point {
	p := secondsToPoint(tm)
	p = Point{p.X * SecondHandLength, p.Y * SecondHandLength} // Scale
	p = Point{p.X, -p.Y}                                      // Flip vertically
	p = Point{p.X + ClockCenterX, p.Y + ClockCenterY}         // Translate to origin
	return p
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
