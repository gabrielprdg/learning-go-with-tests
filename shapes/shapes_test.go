package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	retangle := Rectangle{12.0, 12.0}
	got := CalcPerimeter(retangle)
	want := 48.0

	if got != want {
		t.Errorf("want %.2f but got %f", want, got)
	}
}

func TestArea(t *testing.T) {

	// table Test -> Basically table test is when i put my test cases in a struct and iterate over this struct, avoiding code repetition
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTest {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("want %g but got %g", tt.hasArea, got)
			}
		})

	}

}
