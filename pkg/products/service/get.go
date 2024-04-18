package service

import (
	"fmt"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
)

func (s *service) Get(page int) (*productsPkg.Result, error) {
	result, err := s.repo.Get(page)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrGetProduct, err)
	}

	return result, nil
}
