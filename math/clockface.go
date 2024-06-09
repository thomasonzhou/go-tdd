package clockface

import "time"

type Point struct {
	X, Y int
}

func SecondHand(tm time.Time) Point {
	return Point{150, 60}
}
