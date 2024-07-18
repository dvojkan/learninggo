package main

import "fmt"

// terraform accomplishes nothing
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func terraformByReference(planets *[8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func modifyValue(x *int) {
	*x = *x * 2
}

func main() {
	// this declares the array of strings
	// ... means that the number of element in array should be determined by compiler
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets)
	fmt.Println(planets)

	terraformByReference(&planets)
	fmt.Println(planets)

	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)

	// Define a variable
	num := 10

	fmt.Println("Before modification:", num)

	// Pass the address of the variable to the function
	modifyValue(&num)

	fmt.Println("After modification:", num)

}
