package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (area float64) {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}
