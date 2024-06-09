package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), math.Pi / 2.0},
		{simpleTime(0, 0, 20), math.Pi * 2.0 / 3.0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := secondsToRadians(c.time)

			if !roughlyEqualFloat(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}
func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (60 * 30))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := minutesToRadians(c.time)

			if !roughlyEqualFloat(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func roughlyEqualFloat(want, got float64) bool {
	const margin = 1e-9
	return math.Abs(want-got) < margin
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.point
			got := secondsToPoint(c.time)

			if !roughlyEqualPoint(want, got) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func roughlyEqualPoint(p1, p2 Point) bool {
	return roughlyEqualFloat(p1.X, p2.X) && roughlyEqualFloat(p1.Y, p2.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2024, time.October, 14, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
