package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	paymentStruct := []struct {
		id    int
		price float64
	}{
		{id: 1, price: 23.8},
		{id: 2, price: 15.4},
	}
	fmt.Printf("paymentStruct: %v\n", paymentStruct)
}
