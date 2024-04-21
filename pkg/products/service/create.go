package service

import (
	"fmt"
	"strings"
	"time"

	productsPkg "github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/google/uuid"
)

func (s *service) Create(
	product *productsPkg.Product,
) (*productsPkg.Product, error) {
	if product.Name == "" {
		return nil, fmt.Errorf("name fild canot be empty")
	}

	if product.Value == 0.0 {
		return nil, fmt.Errorf("value cannot be zero")
	}

	if product.Units == 0 {
		return nil, fmt.Errorf("you need to have more than one unit to create the product")
	}

	newProduct := &productsPkg.Product{
		ID:          uuid.New(),
		Name:        strings.ToLower(product.Name),
		Description: product.Description,
		Value:       product.Value,
		Units:       product.Units,
		Image:       product.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(newProduct); err != nil {
		return nil, fmt.Errorf("error in create product: %v", err)
	}

	return newProduct, nil
}
