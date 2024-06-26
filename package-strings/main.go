package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "let's count some words!"
	//var words []string
	words := strings.Fields(text)
	fmt.Println("Found", len(words), "words")
}
