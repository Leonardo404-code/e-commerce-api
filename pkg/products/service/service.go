package service

import (
	"github.com/Leonargo404-code/e-commerce/pkg/google"
	"github.com/Leonargo404-code/e-commerce/pkg/products"
)

type Service struct {
	Repo    products.Repository
	Storage google.Google
}

func Must(repository products.Repository, storage google.Google) products.Services {
	return &Service{
		Repo:    repository,
		Storage: storage,
	}
}
