package main

func sum(
	x int,
	y int,
) (int, bool) {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}

func main() {
	result, status := sum(2, 2)
	println(result, status)
}
