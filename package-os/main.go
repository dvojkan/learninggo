package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//filename := ""words.txt""
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var wordCount int

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}

	/*
		fileContents, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	*/
	//words := strings.Fields(string(fileContents))

	// words := strings.Fields(string(fileContents))

	fmt.Printf("found %d words", wordCount)
}
