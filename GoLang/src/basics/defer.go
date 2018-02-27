package main

import "fmt"

func greeting() string {
	fmt.Println("Executing greeting()")
	return "world"
}

func main() {
	defer fmt.Println(greeting())

	fmt.Println("hello")
}
