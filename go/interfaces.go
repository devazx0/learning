package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Line struct {
	lenght float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// type assertion
func checkShape(s Shape) {
	c, ok := s.(Circle) //we can use this to assert if a type is what we are expecting.

	fmt.Println(c)
	fmt.Println(ok)

	if ok {
		fmt.Println("It's a circle with radius:", c.Radius)
		return
	}
	fmt.Println("It is not a circle")
}

func findShape(s Shape) {
	switch v := s.(type) {
	case Rect:
		fmt.Println("This is a rect with width of ", v.Width)
	case Circle:
		fmt.Println("This is a circle with radius", v.Radius)
	default:
		fmt.Println("Unknown shape")
	}
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	r := Rect{12.3, 10.4}
	c := Circle{30.8}

	l := Line{20}

	printArea(r)
	printArea(c)
	// printArea(l)
}
