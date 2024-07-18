package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	// this is slice declaration
	// [] indicates that we are declaring slice, not array
	planets := []string{" Venus ", "Earth ", " Mars"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}
