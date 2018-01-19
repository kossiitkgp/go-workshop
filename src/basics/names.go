package main

// In Go, a name is exported if it begins with a capital letter.

import (
	"fmt"
	"math"
)

func main() {
	// When importing a package, you can refer only to its exported names.
	// Any "unexported" names are not accessible from outside the package.
	fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
