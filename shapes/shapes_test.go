package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	got := rect.Perimeter()
	want := 40.0

	assertFloat64Equal(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 9.0}, 108.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertFloat64Equal(t, got, tt.want)
	}

}

func assertFloat64Equal(t *testing.T, f1, f2 float64) {
	if f1 != f2 {
		t.Errorf("got %g want %g", f1, f2)
	}
}
