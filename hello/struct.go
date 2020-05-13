package main

import "math"

//Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle struct
type Circle struct {
	Radius float64
}

//Shape interface
type Shape interface {
	Area() float64
}

//Perimeter function
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//Area func for rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area func for circle
func (c *Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}
