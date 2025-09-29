package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine runs in the background (multiple threads)
	fmt.Println("Goroutines")
	go say("Hello to use Goroutine")
	fmt.Println("Hello!")

	go func(msg string) {
		fmt.Println(msg)
	}("Helloooooo!")

	time.Sleep(2 * time.Second)
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
