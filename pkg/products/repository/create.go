package repository

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (r *repository) Create(newProduct *products.Product) error {
	if err := r.db.Create(&newProduct); err.Error != nil {
		return err.Error
	}
	return nil
}
