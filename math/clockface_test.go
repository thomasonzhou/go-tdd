package clockface_test

import (
	clockface "hello/math"
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

	const margin = 1e-9

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			want := c.angle
			got := clockface.SecondsToRadians(c.time)

			if math.Abs(want-got) > margin {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2024, time.October, 14, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
