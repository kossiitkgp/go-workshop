package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// Try removing the explicit type conversion expressions.
	var z uint = unint(f)
	fmt.Println(x, y, z)
}
