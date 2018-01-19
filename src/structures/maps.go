package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Point1"] = Vertex{3, 4}
	m["Origin"] = Vertex{}
	fmt.Println("Origin:", m["Origin"])
	fmt.Println("Point1:", m["Point1"])

	v, ok := m["Origin"]
	if ok {
		fmt.Println("Origin found in the map")
	}
}
