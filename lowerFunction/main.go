package main

import "fmt"

func lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func main() {

	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	// Using var
	var a int
	a = 5

	// Using short variable declaration
	b := 10

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	println(lower('V'))

}
