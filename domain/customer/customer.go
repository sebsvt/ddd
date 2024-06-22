package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/entity"
	"github.com/sebsvt/ddd-go/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person      *entity.Person
	products    []*entity.Item
	transaction []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}
