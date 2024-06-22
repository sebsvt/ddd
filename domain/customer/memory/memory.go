package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (repo *MemoryRepository) Get(customer_id uuid.UUID) (customer.Customer, error) {
	if cus, has := repo.customers[customer_id]; has {
		return cus, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}
func (repo *MemoryRepository) Add(new_customer customer.Customer) error {
	if repo.customers == nil {
		repo.Lock()
		repo.customers = make(map[uuid.UUID]customer.Customer)
		repo.Unlock()
	}
	if _, has := repo.customers[new_customer.GetID()]; has {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustmer)
	}
	repo.Lock()
	repo.customers[new_customer.GetID()] = new_customer
	repo.Unlock()
	return nil
}
func (repo *MemoryRepository) Update(cus customer.Customer) error {
	if _, has := repo.customers[cus.GetID()]; !has {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	repo.Lock()
	repo.customers[cus.GetID()] = cus
	repo.Unlock()
	return nil
}
