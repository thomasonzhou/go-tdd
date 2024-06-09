package clockface_test

// import (
// 	"testing"
// 	"time"

// 	clockface "hello/math"
// )

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := time.Date(1337, time.October, 14, 0, 0, 0, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 - 90}
// 	got := clockface.SecondHand(tm)

// 	if want != got {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.October, 14, 0, 0, 30, 0, time.UTC)

// 	want := clockface.Point{X: 150, Y: 150 + 90}
// 	got := clockface.SecondHand(tm)

// 	if want != got {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }
