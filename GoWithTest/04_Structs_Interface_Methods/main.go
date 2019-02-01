package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Heigth float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Heigth float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Heigth) * 0.5
}

func main() {

}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Heigth)
}

func Area(r Rectangle) float64 {
	return r.Width * r.Heigth
}
