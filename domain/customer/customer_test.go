package customer_test

import (
	"testing"

	"github.com/sebsvt/ddd-go/domain/customer"
)

func TestCustomerNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testcases := []testCase{
		{"Empyt name", "", customer.ErrInvalidPerson},
		{"Valid name", "Sebastian", nil},
	}
	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected: %v but got: %v", tc.expectedErr, err)
			}
		})
	}
}
