package main

import (
	"strings"
)

// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.

func WordCount(s string) map[string]int {

	stringArray := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range stringArray {
		m[word] += 1
	}

	return m
}
