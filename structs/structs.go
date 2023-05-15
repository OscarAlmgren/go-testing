package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() (area float64) {
	return r.height * r.width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() (area float64) {
	return (t.base * t.height) / 2
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.height + r.width)
}
