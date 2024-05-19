package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	assertFloat64Equal(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 9.0)
	want := 108.0
	assertFloat64Equal(t, got, want)
}

func assertFloat64Equal(t *testing.T, f1, f2 float64) {
	if f1 != f2 {
		t.Errorf("got %.2f want %.2f", f1, f2)
	}
}
