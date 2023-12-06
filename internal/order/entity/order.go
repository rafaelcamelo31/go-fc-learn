package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (order *Order) IsValid() error {
	if order.ID == "" {
		return errors.New("invalid id")
	}

	return nil
}
