package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	ch := make(chan int)
	// T2
	for i := 0; i < 10000; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 100000; i++ {
		ch <- i
	}

	close(ch)
}
