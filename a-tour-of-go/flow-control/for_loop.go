package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	// and then used as "while loop"
	iterator := 1
	for iterator < 1000 {
		iterator += iterator
	}
	fmt.Println(iterator)
}
