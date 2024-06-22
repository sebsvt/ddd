package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/customer"
	"github.com/sebsvt/ddd-go/domain/customer/memory"
	"github.com/sebsvt/ddd-go/domain/customer/mongo"
	"github.com/sebsvt/ddd-go/domain/product"
	prodmem "github.com/sebsvt/ddd-go/domain/product/memory"
)

// Configuration Section

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customer_repository customer.CustomerRepository
	product_repository  product.ProductRepository
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customer_repository = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithProductRepository(pr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.product_repository = pr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	pr := prodmem.New()
	for _, p := range products {
		pr.Add(p)
	}
	return WithProductRepository(pr)
}

func WithMongoCustomerRepository(connection_string string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connection_string)
		if err != nil {
			return err
		}
		os.customer_repository = cr
		return nil
	}
}

// Business Logic Section

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (srv *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	var price float64
	var products []product.Product

	cust, err := srv.customer_repository.Get(customerID)
	if err != nil {
		return 0, err
	}
	for _, id := range productIDs {
		p, err := srv.product_repository.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", cust.GetID(), len(products))
	return price, nil
}
