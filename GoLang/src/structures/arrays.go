package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array literal

	var s []int = primes[1:4]
	fmt.Println(s)

	// q := []int{2, 3, 5, 7, 11, 13} // slice literal

	var r []int
	if r == nil {
		fmt.Println("The zero value of slices is nil.")
	}

	// Creating slices with make
	// a := make([]int, 5, 10)

	board()
}

func board() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
