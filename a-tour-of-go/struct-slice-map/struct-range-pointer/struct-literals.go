package main

import "fmt"

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	v4 = Vertex{Y: 5, X: 3}
	p  = &Vertex{1, 2} // has type *Vertex
)

func StructLiterals() {
	fmt.Println(v1, p, v2, v3, v4)
}
