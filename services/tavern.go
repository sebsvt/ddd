package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService   *OrderService
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	log.Printf("Bill the customer: %0.0f", price)
	return nil
}
