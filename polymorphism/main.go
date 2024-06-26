package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

/*
	func Area1(c Circle) float64  {
		return c.Radius * c.Radius * math.Pi
	}

	func Area1(s Square) float64 {
		return s.Length * s.Length
	}
*/
func PrintArea(s Shape) {
	fmt.Println("Area of shape is:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	s := Square{Length: 4}

	PrintArea(c)
	PrintArea(s)
}
