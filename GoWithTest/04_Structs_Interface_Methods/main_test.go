package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("perimeter-h:10 w:10", func(t *testing.T) {
		got := Perimeter(Rectangle{Width: 10, Heigth: 10})
		want := 40.0
		assertEquals(t, got, want)
	})

	t.Run("perimeter-h:20 w:10", func(t *testing.T) {
		got := Perimeter(Rectangle{Width: 20, Heigth: 10})
		want := 60.0
		assertEquals(t, got, want)
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		hasArea  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Heigth: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Heigth: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v; Got: %.2f; want: %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}
}

func assertEquals(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %.2f; want: %.2f", got, want)
	}

}
