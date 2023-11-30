package main

import "fmt"

func sum(
	x int,
	y int,
) int {
	return x + y
}

func main() {
	result := sum(2, 2)
	fmt.Printf("The sum results in %d\n", result)
}
