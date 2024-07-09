package main

import "fmt"

type Address struct {
	City    string
	ZipCode string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "New York",
			ZipCode: "10001",
		},
	}

	// Accessing nested struct values
	fmt.Println(p.Name)            // Output: Alice
	fmt.Println(p.Age)             // Output: 30
	fmt.Println(p.Address.City)    // Output: New York
	fmt.Println(p.Address.ZipCode) // Output: 10001

	//Accessing Values in Slices of Structs
	people := []Person{
		{Name: "Alice",
			Age: 30,
			Address: Address{
				City:    "New York",
				ZipCode: "10001"}},
		{Name: "Bob",
			Age: 25,
			Address: Address{
				City:    "Philly",
				ZipCode: "10001"}},
	}

	// Accessing values in a slice of structs
	fmt.Println(people[0].Name)         // Output: Alice
	fmt.Println(people[1].Age)          // Output: 25
	fmt.Println(people[1].Address.City) // Output: Philly

	//Accessing Values in Maps of Structs
	peopleMap := map[string]Person{
		"Alice": {Name: "Alice",
			Age: 30,
			Address: Address{
				City:    "New York",
				ZipCode: "10001"}},
		"Bob": {Name: "Bob",
			Age: 25,
			Address: Address{
				City:    "Philly",
				ZipCode: "20001"}},
	}

	// Accessing values in a map of structs
	fmt.Println(peopleMap["Alice"].Name)          // Output: Alice
	fmt.Println(peopleMap["Bob"].Address.ZipCode) // Output: 20001

}
