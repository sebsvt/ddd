package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/customer"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customer_repository.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}

func TestMongoTavern(t *testing.T) {
	product := init_products(t)
	os, err := NewOrderService(
		WithMongoCustomerRepository("mongodb://root:example@localhost:27017"),
		WithMemoryProductRepository(product),
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}
	cust, err := customer.NewCustomer("Sebastian")
	if err != nil {
		t.Error(err)
	}
	err = os.customer_repository.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		product[0].GetID(),
	}
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
