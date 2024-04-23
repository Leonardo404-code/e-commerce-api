package service

import (
	"github.com/Leonargo404-code/e-commerce/pkg/google"
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

type service struct {
	repo    products.Repository
	storage google.Google
}

func Must(repository products.Repository, storage google.Google) products.Services {
	return &service{
		repo:    repository,
		storage: storage,
	}
}
