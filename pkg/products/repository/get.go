package repository

import (
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

// Get implements products.Repository.
func (r repository) Get(page int) (*productsPkg.Result, error) {
	var products []productsPkg.Product

	totalPages := r.db.Find(&products).RowsAffected

	if totalPages == 0 {
		totalPages = 1
	}

	err := r.db.Offset((page - 1) * 10).Limit(10).Order("id desc").Find(&products)
	if err.Error != nil {
		return nil, fmt.Errorf("failed in find products: %v", err.Error)
	}

	return &productsPkg.Result{
		Products: products,
		Total:    int(totalPages),
		Page:     page,
	}, nil
}
