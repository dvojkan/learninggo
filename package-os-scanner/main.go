package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func customSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We found a comma, return the data up to and including it.
		return i + 1, data[:i], nil
	}
	// If at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

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

	scanner.Split(customSplitFunc)

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
