package product_test

import (
	"testing"

	"github.com/sebsvt/ddd-go/domain/product"
)

func TestProductNewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}
	testcases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			description: "s",
			price:       0,
			expectedErr: product.ErrMissingValues,
		},
		{
			test:        "valid values",
			name:        "Chocolate",
			description: "I eat it everyday",
			price:       1.0,
			expectedErr: nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.NewProduct(tc.name, tc.description, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
