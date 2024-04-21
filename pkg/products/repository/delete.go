package repository

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (r *repository) Delete(filter *products.Filter) error {
	if err := r.db.Delete(&products.Product{}, filter.ID); err.Error != nil {
		return err.Error
	}

	return nil
}
