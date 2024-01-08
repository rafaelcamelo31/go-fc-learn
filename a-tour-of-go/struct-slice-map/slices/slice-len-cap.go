package main

import "fmt"

func SliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("default", s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("re-slice", s)

	// Extend its length.
	s = s[:4]
	printSlice("extend", s)

	// Drop its first two values.
	s = s[2:]
	printSlice("drop", s)
}

func printSlice(message string, s []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", message, len(s), cap(s), s)
}
