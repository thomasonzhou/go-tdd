package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const (
	SecondHandLength = 90
	ClockCenterX     = 150
	ClockCenterY     = 150
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func secondHand(w io.Writer, tm time.Time) {
	p := secondsToPoint(tm)
	p = Point{p.X * SecondHandLength, p.Y * SecondHandLength} // Scale
	p = Point{p.X, -p.Y}                                      // Flip vertically
	p = Point{p.X + ClockCenterX, p.Y + ClockCenterY}         // Translate to origin
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
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
