package service

import (
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (s *service) Delete(filter *productsPkg.Filter) error {
	products, err := s.repo.Get(filter)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrGetProduct, err)
	}

	if len(products.Products) == 0 {
		return fmt.Errorf("%w: product not found", ErrProductsNotFound)
	}

	if err := s.repo.Delete(filter); err != nil {
		return fmt.Errorf("%w: %v", ErrDeleteProduct, err)
	}

	return nil
}
