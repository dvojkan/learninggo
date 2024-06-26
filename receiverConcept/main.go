package main

import (
	"fmt"
)

// Define a struct type named "Rectangle"
type Rectangle struct {
	width  float64
	height float64
}

// Method associated with the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// Create an instance of the Rectangle struct
	rect := Rectangle{width: 3, height: 4}

	// Call the Area method on the rect instance
	area := rect.Area()

	// Print the result
	fmt.Println("Area of the rectangle:", area)
}
