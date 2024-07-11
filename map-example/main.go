package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]

	fmt.Println(earth)

	fmt.Println(len(planets))

	fmt.Println(planets[3] == "")

	i := 7

	planets[i] = "Pluto"

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfs[0])

	planets2 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
		"Pluto",
	}

	fmt.Println("There are ", len(planets2), " in planets2 array")

	// iterating through array:
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	fmt.Println("--------------------------")

	for _, planet := range planets2 {
		fmt.Println(planet)
	}

	// how range works:
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("-------------------------- before planetsMarkII")
	planetsMarkII := planets2

	planets2[2] = "whoops"
	fmt.Println(planets2)
	fmt.Println(planetsMarkII)

	fmt.Println("-------------------------- before map example with interface")

	//map example with interface
	// var request map[string]interface{}

	// Initialize the map
	// var request map[string]interface{}
	request := make(map[string]interface{})

	// Add different types of data to the map
	request["username"] = "johndoe"
	request["age"] = 30
	request["balance"] = 100.50
	request["is_active"] = true
	request["preferences"] = map[string]interface{}{
		"language": "Go",
		"theme":    "dark",
	}

	// Print the map
	fmt.Println("Request map:", request)

	// Accessing elements from the map
	username := request["username"].(string)
	age := request["age"].(int)
	balance := request["balance"].(float64)
	isActive := request["is_active"].(bool)
	preferences := request["preferences"].(map[string]interface{})
	language := preferences["language"].(string) //type assertation
	theme := preferences["theme"].(string)       //type assertation

	// Print the accessed elements
	fmt.Println("Username:", username)
	fmt.Println("Age:", age)
	fmt.Println("Balance:", balance)
	fmt.Println("Is Active:", isActive)
	fmt.Println("Language Preference:", language)
	fmt.Println("Theme Preference:", theme)

}
