package structures

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Ray float64
}

type Triangle struct {
	Base  float64
	Width float64
}

type Form interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Ray * c.Ray
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Width) * 0.5
}
