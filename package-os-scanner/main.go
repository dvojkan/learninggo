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
	//filename := "words.txt"
	var wordCount int
	if len(os.Args) < 2 {
		log.Println("need to provide at least one filename!")
		os.Exit(1)
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Printf("%s: %v", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)

		//option 1 - custom split function
		scanner.Split(customSplitFunc)

		//option 2 - using ScanWords as split function
		//scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			wordCount += len(words)
		}

		if scanner.Err() != nil {
			log.Printf("scan error %s: %v", filename, scanner.Err())
			file.Close()
			continue
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
		file.Close()
	}

	fmt.Printf("found %d words", wordCount)
}
