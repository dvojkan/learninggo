package main

import (
	"fmt"
)

func process(channel chan string, output chan string) {
	s := <-channel  // Receive a string from the channel
	s = s + "World" // Concatenate "World" to the received string
	output <- s     // Send the result to the output channel
}

func main() {
	// Create channels for communication
	channel := make(chan string)
	output := make(chan string)

	// Call process function synchronously
	go process(channel, output)

	// Send "Hello" to the channel
	channel <- "Hello, "

	// Receive and print the result from the output channel
	fmt.Println(<-output)
}
