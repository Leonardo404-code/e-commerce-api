package service

import (
	"github.com/Leonargo404-code/e-commerce/pkg/products"
	"github.com/Leonargo404-code/e-commerce/pkg/products/repository"
)

type service struct {
	repo products.Repository
}

func Must() products.Services {
	return &service{
		repo: repository.Must(),
	}
}
