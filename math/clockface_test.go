package clockface

import (
	"github.com/thomasonzhou/go-tdd/math/clockface"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.October, 14, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if want != got {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
