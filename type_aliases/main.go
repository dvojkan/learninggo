package main

import "fmt"

// Define type alias
type IDKey = string

// This is not the same as defining new type. If I want to
type IDKeyNew string

// there is a difference in defining new type and alias!!

func main() {
	var myKey IDKey
	myKey = "12345"
	fmt.Println(myKey) // Output: 12345

	// Since IDKey is just an alias, you can use it interchangeably with string
	var anotherKey string
	anotherKey = myKey
	fmt.Println(anotherKey) // Output: 12345

	// but since IDKeyNew is new type, it can not be used interchangeably with string
	var myKeyNew IDKeyNew
	myKeyNew = "12345"
	fmt.Println(myKeyNew) // Output: 12345

	// The following would cause a compile-time error because IDKeyNew is a distinct type
	// var anotherKeyNew string
	// anotherKeyNew = myKeyNew // Error: cannot use myKey (type IDKey) as type string in assignment

	// in order to continue working with IDKeyNew and myKeyNew I would need to do type conversion!
	var anotherKeyNew string
	anotherKeyNew = string(myKeyNew) // Error: cannot use myKey (type IDKey) as type string in assignment
	fmt.Println(anotherKeyNew)       // Output: 12345

}
