package main

import (
	"fmt"

	"github.com/rafaelcamelo31/go-fc-learn/internal/order/entity"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		panic(err)
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println("The final price is:", order.FinalPrice)
}
