package main

import (
	"fmt"
	"math/cmplx"
)

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var python, java, c = true, false, "maybe"
	fmt.Println(i, j, python, java, c)

	fmt.Printf("Type: %10T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %10T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %10T Value: %v\n", z, z)
}
