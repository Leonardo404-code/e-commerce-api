package service

import (
	"context"
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (s *Service) Delete(filter *productsPkg.Filter) error {
	products, err := s.Repo.Get(filter)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrGetProduct, err)
	}

	if len(products.Products) == 0 {
		return fmt.Errorf("%w: product not found", ErrProductsNotFound)
	}

	if len(products.Products) > 1 {
		return fmt.Errorf("%w: more than 1 product found", ErrProductsConflict)
	}

	productName := products.Products[0].Name

	if err := s.Repo.Delete(filter, s.deleteFromBucket(productName)); err != nil {
		return fmt.Errorf("%w: %v", ErrDeleteProduct, err)
	}

	return nil
}

func (s *Service) deleteFromBucket(imageURL string) func() error {
	return func() error {
		return s.Storage.Delete(context.Background(), imageURL)
	}
}
