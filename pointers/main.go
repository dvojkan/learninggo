package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	p.Age++ // Dereference 'p' and increment the Age field - this is automatic dereferencing - different than with x variable.
	// this can be done with more clear, and without automatic dereferencing!!
	(*p).Age++
	// Writing *p.Age++ is incorrect because it implies Age is a pointer, which it is not.
}

func increment(x *float32) {
	*x = *x + 1 // Dereference 'x' and increment the value by 1
}

func main() {
	var a int = 10 // Declare an integer variable
	var b float32 = 10
	var p *int // Declare a pointer to an integer

	p = &a // Reference: get the address of 'a' and assign it to 'p'

	fmt.Println("Value of a:", a)    // Output: 10
	fmt.Println("Address of a:", &a) // Output: memory address of 'a'
	//fmt.Println("Is this possible:", *a)   // Output: No output - it is showing compiler error!!
	fmt.Println("Value of p:", p)          // Output: same memory address of 'a'
	fmt.Println("Value at address p:", *p) // Dereference: Output: 10

	increment(&b) // Pass the address of 'a' to the function

	fmt.Println("After increment:", b) // Output: 11

	person := Person{Name: "Alice", Age: 30}

	fmt.Println("Before birthday:", person) // Output: {Alice 30}

	birthday(&person) // Pass the address of 'person' to the function

	fmt.Println("After birthday:", person) // Output: {Alice 32} -- because I have two ways of showing how to dereference something in birthday function!!

}
