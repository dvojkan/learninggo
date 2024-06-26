package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	fmt.Println("Sending message...")
	ch <- "Hello, receiver!" // Send a message on the channel
	fmt.Println("Message sent")
}

func receiver(ch chan string) {
	fmt.Println("Waiting for message...")
	msg := <-ch // Receive a message from the channel
	fmt.Println("Received:", msg)
}

func main() {
	// Create an unbuffered channel
	ch := make(chan string)

	// Start sender and receiver goroutines
	go sender(ch)
	go receiver(ch)

	// Sleep for a while to allow goroutines to complete
	time.Sleep(1 * time.Second)
}
