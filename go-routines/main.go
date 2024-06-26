package main

import (
	"fmt"
	"time"
)

func DoSomethingSlow() {
	fmt.Println(time.Now().Format(time.RFC3339Nano), "SLOW: maybe I'm doing network stuff?")
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().Format(time.RFC3339Nano), "SLOW: Okay, finally finished!")
}

func process(channel chan string, output chan string) {
	s := <-channel
	s = s + " World"
	output <- s
}

func main() {
	fmt.Println(time.Now().Format(time.RFC3339Nano), "Starting the main task...")
	// Create a goroutine to do something slow.
	go DoSomethingSlow()
	fmt.Println(time.Now().Format(time.RFC3339Nano), "Resuming the main task...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(time.Now().Format(time.RFC3339Nano), "Finished the main task!")
	time.Sleep(1 * time.Second)

	channel := make(chan string)
	output := make(chan string)

	channel <- "Hello, "

	process(channel, output)

	finalString := <-output
	fmt.Println("This string was build concurrently:", finalString)
}
