package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
)

func main() {
	v := Vertex{1, 2}
	v.X, v.Y = 3, 4
	fmt.Printf("%v\n", v)

	p := &v
	// Programmer friendly
	p.X, p.Y = 1, 2
	fmt.Println(v)
}
