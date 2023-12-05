package main

import "fmt"

// T1
func main() {
	channel := make(chan string)

	// T2
	go func() {
		channel <- "Hello, World!"
	}()

	fmt.Println(<-channel)
}
