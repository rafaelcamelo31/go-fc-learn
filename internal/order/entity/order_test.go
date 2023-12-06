package entity_test

import (
	"testing"

	"github.com/rafaelcamelo31/go-fc-learn/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "123"}

	assert.Error(t, order.IsValid(), "invalid price")
}
