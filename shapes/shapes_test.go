package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := Perimeter(rect)
	want := 40.0

	assertFloat64Equal(t, got, want)
}

func TestArea(t *testing.T) {
	rect := Rectangle{12.0, 9.0}
	got := Area(rect)
	want := 108.0
	assertFloat64Equal(t, got, want)
}

func assertFloat64Equal(t *testing.T, f1, f2 float64) {
	if f1 != f2 {
		t.Errorf("got %.2f want %.2f", f1, f2)
	}
}
