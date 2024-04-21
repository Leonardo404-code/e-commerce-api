package service

import (
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (s *service) Get(filter *productsPkg.Filter) (*productsPkg.Result, error) {
	result, err := s.repo.Get(filter)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrGetProduct, err)
	}

	return result, nil
}
