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

	// map example with interface
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

	//Map with pointers
	fmt.Println("\n--------------------------MAP WITH POINTERS------------------------------------")
	// Define the Pocket struct
	type Pocket struct {
		Color string
		Size  int
	}

	// Initialize the map
	pockets := make(map[string]*Pocket)

	// Create some Pocket instances and add them to the map
	// Well I would say that this pocket instance exists out of map, and that we are just storing pointers to these pockets
	pockets["frontLeft"] = &Pocket{Color: "blue", Size: 10}
	// this is interesting to me! How & is used for Pocket instances creation!!! It can be said that this is creating
	// Pocket instance and then take address of it @
	// so this does two things
	// pocket := Pocket{Color: "blue", Size: 10}
	// pocketPointer := &pocket
	// pockets["frontLeft"] = pocketPointer
	// here I stumbled upon to new command so I need to clarify this to me since I do not understand it!!

	pockets["frontRight"] = &Pocket{Color: "red", Size: 12}
	pockets["backLeft"] = &Pocket{Color: "green", Size: 8}
	pockets["backRight"] = &Pocket{Color: "black", Size: 9}

	// Print the map
	fmt.Println("Pockets map:")
	for position, pocket := range pockets {
		fmt.Printf("%s: Color = %s, Size = %d\n", position, (*pocket).Color, (*pocket).Size) //<- this is explicit dereferencing or the same functionality with implicit dereferencing shown in next line
		fmt.Printf("%s: Color = %s, Size = %d\n", position, pocket.Color, pocket.Size)
	}

	// Accessing and modifying elements in the map
	pockets["frontLeft"].Color = "yellow" // what is interesting here is that we have the same syntax as if the Pocket is in the Map not an object outside map - the syntax for changing is the same!!
	// meaning that we have the same syntax for two different cases. Which is little bit confusing for me.
	// this can be written as well
	// (*pockets["frontLeft"]).Color = "yellow"
	fmt.Println("\nUpdated Pockets map:")
	for position, pocket := range pockets {
		fmt.Printf("%s: Color = %s, Size = %d\n", position, pocket.Color, pocket.Size)
	}

	// here I am adding a line to check if the real object to which frontLeft pointer is showin really contains that new value, meaning that the change happened on object not on map!
	// so I need to find an address of this object
	fmt.Println("\nCheck real object!!")
	pointerToFrontLeftObject := pockets["frontLeft"]                                             // I think that I've got the address of the Pocket object which is assigned to "frontLeft" element of the map!
	fmt.Println("Value at address pointerToFrontLeftObject:", (*pointerToFrontLeftObject).Color) //or this is the same just another syntax!!
	fmt.Println("Value at address pointerToFrontLeftObject:", pointerToFrontLeftObject.Color)

	// Adding a new pocket to the map
	pockets["sidePocket"] = &Pocket{Color: "purple", Size: 7}
	fmt.Println("\nAfter adding sidePocket:")
	for position, pocket := range pockets {
		fmt.Printf("%s: Color = %s, Size = %d\n", position, pocket.Color, pocket.Size)
	}

	// Delete an existing pocket from the map!
	// before deletion from map I would like to save the address of sidePocket object on heap!
	pointerToSideObject := pockets["sidePocket"]
	delete(pockets, "sidePocket")
	fmt.Println("\nAfter deleting sidePocket:")
	for position, pocket := range pockets {
		fmt.Printf("%s: Color = %s, Size = %d\n", position, pocket.Color, pocket.Size)
	}
	// in this case I have deleted one element of the map, but that object still exists in the heap, and I will simulate this, by adding it again!!

	// now I am going to add it again!!
	//pockets["sidePocket"] = &Pocket{Color: "purple", Size: 7}
	pockets["sidePocket"] = pointerToSideObject
	fmt.Println("\nAfter adding sidePocket for second time:")
	for position, pocket := range pockets {
		fmt.Printf("%s: Color = %s, Size = %d\n", position, pocket.Color, pocket.Size)
	}

	// after this I am getting the same values fof sidePocket, because it is not deleted - just the reference to it in the map pockets

	// I have learned some new thing! How to destroy the map:
	// Destroys the map
	fmt.Println("\nDestroys the map")
	pockets = nil

	// If I have Pocket instances in that map, and in this case I mean pointers to pocket instances, by destroying the map, they will still exists.
	// GC will delete them only when there are no more references to it, including from local variables.

	//map with slices example
	fmt.Println("\n--------------------------MAP WITH SLICES------------------------------------")
	// I've concluded here that I need to learn more about slices!!

	// Declare a map with string keys and slice of string values
	exampleMap := map[string][]string{
		"fruits":  {"apple", "banana", "cherry", "plum"},
		"colors":  {"red", "green", "blue"},
		"animals": {"cat", "dog", "elephant"},
	}

	// Print the map
	fmt.Println(exampleMap)

	// Accessing elements
	fruits := exampleMap["fruits"]
	fmt.Println("Fruits:", fruits)

	// Updating elements
	exampleMap["fruits"] = append(exampleMap["fruits"], "orange")
	fmt.Println("Updated fruits:", exampleMap["fruits"])

	// Adding a new key-value pair
	exampleMap["countries"] = []string{"USA", "Canada", "Mexico", "Tunisia", "Algeria"}
	fmt.Println("Added countries:", exampleMap["countries"])

	// Deleting a key-value pair
	delete(exampleMap, "colors")
	fmt.Println("Map after deleting colors:", exampleMap)

	// Checking if a key exists
	if animals, ok := exampleMap["animals"]; ok {
		fmt.Println("Animals:", animals)
	} else {
		fmt.Println("No animals found")
	}

	// Iterating over the map
	fmt.Println("Iterating over the map:")
	for key, value := range exampleMap {
		fmt.Println(key, ":", value)
	}

}
