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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 9.0}, want: 108.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}

}

func assertFloat64Equal(t *testing.T, f1, f2 float64) {
	if f1 != f2 {
		t.Errorf("got %g want %g", f1, f2)
	}
}
