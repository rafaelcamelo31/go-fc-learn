package main

import "fmt"

type Vertex struct {
	X, Y int
}

func Struct() {
	v := Vertex{1, 2}

	v.X = 4
	fmt.Println(v.X)
}
