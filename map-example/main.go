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

}
