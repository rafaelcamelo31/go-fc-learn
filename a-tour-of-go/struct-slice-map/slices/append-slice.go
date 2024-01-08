package main

import "fmt"

func AppendSlice() {
	var s []int
	printSlice1(s)

	// append works on nill slices.
	s = append(s, 0)
	printSlice1(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice1(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice1(s)
}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
