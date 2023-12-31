package main

import "fmt"

func Pointer() {
	i, j := 42, 2701

	p := &i         // point to i (assign address of i)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j (assign address of j)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
