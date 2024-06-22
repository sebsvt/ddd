package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/customer"
)

func TestMemoryGetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	cust, err := customer.NewCustomer("Sebastian")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}
	testcases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected: %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryAddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}
	test_cases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Sebastian",
			expectedErr: nil,
		},
	}
	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}
			cust, err := customer.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
