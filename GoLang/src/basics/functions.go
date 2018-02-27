package main

import "fmt"

// Adding two integers
func add(x int, y int) int {
	return x + y
}

// Swap strings
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swap("Hello", "World"))
}
