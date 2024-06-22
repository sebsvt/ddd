package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/product"
)

func TestMemoryProductRepositoryAdd(t *testing.T) {
	repo := New()
	product, err := product.NewProduct("Beer", "Good for me?", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(product)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProductRepositoryGet(t *testing.T) {
	// Add Product before get product
	repo := New()
	prod, err := product.NewProduct("Beer", "Good for me?", 1.99)
	if err != nil {
		t.Error(err)
	}
	repo.Add(prod)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          prod.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	existingProd, err := product.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.products))
	}
}
