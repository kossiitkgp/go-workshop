package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	// Functions can be arguments to other functions
	return fn(3, 4)
}

func main() {
	// Functions can be assigned to variables
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
