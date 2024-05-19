package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := rect.Perimeter()
	want := 40.0

	assertFloat64Equal(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 9.0}
		got := rect.Area()
		want := 108.0
		assertFloat64Equal(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func assertFloat64Equal(t *testing.T, f1, f2 float64) {
	if f1 != f2 {
		t.Errorf("got %.2f want %.2f", f1, f2)
	}
}
